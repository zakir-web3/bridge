import { ethers } from "hardhat";
import {
  validateAddress,
  handleError,
  ConfigurationError,
  PermissionError,
} from "./utils/error-handler";

async function main() {
  console.log("开始设置 Token Pair...");

  const [singer] = await ethers.getSigners();
  console.log("操作账户:", singer.address);
  const balance = await ethers.provider.getBalance(singer.address);
  console.log("账户余额:", ethers.formatEther(balance));

  // 从环境变量读取参数（也支持从命令行读取，优先环境变量）
  const bridgeHubAddress = process.env.BRIDGE_HUB_ADDRESS || process.argv[2];
  const chainId = process.env.CHAIN_ID || process.argv[3];
  const tokenAddress = process.env.TOKEN_ADDRESS || process.argv[4];
  const bridgedTokenAddress =
    process.env.BRIDGED_TOKEN_ADDRESS || process.argv[5];
  const tokenDecimal = process.env.TOKEN_DECIMAL || process.argv[6];

  if (!chainId) {
    throw new ConfigurationError(
      "缺少链 ID",
      "请通过环境变量 CHAIN_ID 或命令行参数提供",
      "示例: CHAIN_ID=1 hardhat run scripts/set-token-pair.ts"
    );
  }

  if (!tokenDecimal) {
    throw new ConfigurationError(
      "缺少 Token 精度",
      "请通过环境变量 TOKEN_DECIMAL 或命令行参数提供",
      "示例: TOKEN_DECIMAL=6 hardhat run scripts/set-token-pair.ts"
    );
  }

  // 验证地址格式
  validateAddress(bridgeHubAddress, "BridgeHub 合约");
  validateAddress(tokenAddress, "源链代币");
  validateAddress(bridgedTokenAddress, "桥接代币");

  // 验证链 ID
  const chainIdNum = parseInt(chainId, 10);
  if (isNaN(chainIdNum) || chainIdNum <= 0) {
    throw new ConfigurationError(
      "无效的链 ID",
      `链 ID: ${chainId}`,
      "请提供大于 0 的有效链 ID"
    );
  }

  // 验证 token 精度
  const tokenDecimalNum = parseInt(tokenDecimal, 10);
  if (isNaN(tokenDecimalNum) || tokenDecimalNum < 0 || tokenDecimalNum > 18) {
    throw new ConfigurationError(
      "无效的 Token 精度",
      `精度: ${tokenDecimal}`,
      "请提供 0 到 18 之间的有效精度值"
    );
  }

  // 打印参数预览
  console.log("\n参数预览:");
  console.log(`  BridgeHub 合约地址: ${bridgeHubAddress}`);
  console.log(`  链 ID: ${chainIdNum}`);
  console.log(`  源链代币地址: ${tokenAddress}`);
  console.log(`  源链代币精度: ${tokenDecimalNum}`);
  console.log(`  桥接代币地址: ${bridgedTokenAddress}`);

  try {
    // 连接到 BridgeHub 合约
    console.log("\n=== 连接到 BridgeHub 合约 ===");
    const BridgeHub = await ethers.getContractFactory("BridgeHub");
    const bridgeHub = BridgeHub.attach(bridgeHubAddress) as any;

    // 检查调用者是否有 ADMIN_ROLE
    console.log("检查管理员权限...");
    const ADMIN_ROLE = await bridgeHub.ADMIN_ROLE();
    const hasAdminRole = await bridgeHub.hasRole(ADMIN_ROLE, singer.address);

    if (!hasAdminRole) {
      throw new ConfigurationError(
        "权限不足",
        `账户 ${singer.address} 没有 ADMIN_ROLE 权限`,
        "请使用具有管理员权限的账户，或联系合约管理员授予权限"
      );
    }
    console.log("✅ 权限验证通过");

    // 检查当前 token pair 设置
    console.log("检查当前 token pair 设置...");
    const currentBridgedToken = await bridgeHub.tokenPair(
      chainIdNum,
      tokenAddress
    );

    if (currentBridgedToken !== ethers.ZeroAddress) {
      throw new Error(
        `链 ${chainIdNum} 上的代币 ${tokenAddress} 已经映射到 ${currentBridgedToken}`
      );
    }
    console.log("✅ 当前没有映射关系，可以安全设置");

    // 执行 setTokenPair 调用
    console.log("\n=== 执行 setTokenPair ===");
    const tx = await bridgeHub.setTokenPair(
      chainIdNum,
      tokenAddress,
      bridgedTokenAddress,
      tokenDecimalNum
    );

    console.log("交易已提交，等待确认...");
    console.log("交易哈希:", tx.hash);

    const receipt = await tx.wait();
    console.log("✅ 交易确认成功!");
    console.log("区块号:", receipt?.blockNumber);
    console.log("Gas 使用量:", receipt?.gasUsed?.toString());

    // 验证设置结果
    console.log("\n=== 验证设置结果 ===");
    const newBridgedToken = await bridgeHub.tokenPair(chainIdNum, tokenAddress);

    if (newBridgedToken === bridgedTokenAddress) {
      console.log("✅ Token Pair 设置成功!");
      console.log(
        `链 ${chainIdNum} 上的代币 ${tokenAddress} 已映射到 ${bridgedTokenAddress}`
      );

      // 验证精度差值存储
      try {
        const decimalDiff = await bridgeHub.tokenDecimalDiff(
          chainIdNum,
          tokenAddress
        );
        console.log(`✅ 精度差值已存储: ${decimalDiff}`);
      } catch (e) {
        console.log(
          "⚠ 无法验证精度差值，但 Token Pair 已设置"
        );
      }
    } else {
      throw new ConfigurationError(
        "Token Pair 设置失败",
        `期望: ${bridgedTokenAddress}, 实际: ${newBridgedToken}`,
        "请检查交易是否成功执行"
      );
    }
  } catch (error) {
    if (error instanceof ConfigurationError) {
      throw error;
    }
    throw new PermissionError(
      "设置 Token Pair 失败",
      error instanceof Error ? error.message : String(error),
      "请检查网络连接和合约状态，然后重试"
    );
  }
}

main()
  .then(() => {
    console.log("🎉 Token Pair 设置完成!");
    process.exit(0);
  })
  .catch(handleError);
