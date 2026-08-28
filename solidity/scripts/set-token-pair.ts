import { ethers } from "hardhat";
import {
  validateAddress,
  handleError,
  ConfigurationError,
  PermissionError,
} from "./utils/error-handler";

function addressToBytes32(address: string) {
  return ethers.zeroPadValue(address, 32);
}

async function main() {
  console.log("开始设置 Token Pair...");

  const [singer] = await ethers.getSigners();
  console.log("操作账户:", singer.address);
  const balance = await ethers.provider.getBalance(singer.address);
  console.log("账户余额:", ethers.formatEther(balance));

  const bridgeHubAddress = process.env.BRIDGE_HUB_ADDRESS || process.argv[2];
  const chainId =
    process.env.TOKEN_CHAIN_ID || process.argv[3] || process.env.CHAIN_ID;
  const tokenAddress = process.env.TOKEN_ADDRESS || process.argv[4];
  const bridgedTokenAddress =
    process.env.BRIDGED_TOKEN_ADDRESS || process.argv[5];
  const tokenDecimal = process.env.TOKEN_DECIMAL || process.argv[6];
  const pairMode = process.env.PAIR_MODE || "evm";

  if (!chainId) {
    throw new ConfigurationError(
      "缺少链 ID",
      "请通过环境变量 TOKEN_CHAIN_ID / CHAIN_ID 或命令行参数提供",
      "示例: TOKEN_CHAIN_ID=56 hardhat run scripts/set-token-pair.ts"
    );
  }

  if (!tokenDecimal) {
    throw new ConfigurationError(
      "缺少 Token 精度",
      "请通过环境变量 TOKEN_DECIMAL 或命令行参数提供",
      "示例: TOKEN_DECIMAL=6 hardhat run scripts/set-token-pair.ts"
    );
  }

  validateAddress(bridgeHubAddress, "BridgeHub 合约");
  validateAddress(bridgedTokenAddress, "桥接代币");

  const chainIdNum = parseInt(chainId, 10);
  if (isNaN(chainIdNum) || chainIdNum <= 0) {
    throw new ConfigurationError(
      "无效的链 ID",
      `链 ID: ${chainId}`,
      "请提供大于 0 的有效链 ID"
    );
  }

  const tokenDecimalNum = parseInt(tokenDecimal, 10);
  if (isNaN(tokenDecimalNum) || tokenDecimalNum < 0 || tokenDecimalNum > 18) {
    throw new ConfigurationError(
      "无效的 Token 精度",
      `精度: ${tokenDecimal}`,
      "请提供 0 到 18 之间的有效精度值"
    );
  }

  let srcTokenBytes32: string;
  if (pairMode === "solana") {
    const srcTokenHex = process.env.SRC_TOKEN_BYTES32 || tokenAddress;
    if (
      !srcTokenHex ||
      !srcTokenHex.startsWith("0x") ||
      srcTokenHex.length !== 66
    ) {
      throw new ConfigurationError(
        "无效的 Solana mint bytes32",
        `SRC_TOKEN_BYTES32: ${srcTokenHex}`,
        "请提供 32 字节的 hex，例如 0x..."
      );
    }
    srcTokenBytes32 = srcTokenHex;
    console.log("\n参数预览 (Solana src):");
    console.log(`  源链 SPL mint (bytes32): ${srcTokenBytes32}`);
  } else {
    validateAddress(tokenAddress, "源链代币");
    srcTokenBytes32 = addressToBytes32(tokenAddress);
    console.log("\n参数预览 (EVM src):");
    console.log(`  源链代币地址: ${tokenAddress}`);
  }

  console.log(`  BridgeHub 合约地址: ${bridgeHubAddress}`);
  console.log(`  链 ID: ${chainIdNum}`);
  console.log(`  源链代币精度: ${tokenDecimalNum}`);
  console.log(`  桥接代币地址: ${bridgedTokenAddress}`);

  try {
    console.log("\n=== 连接到 BridgeHub 合约 ===");
    const BridgeHub = await ethers.getContractFactory("BridgeHub");
    const bridgeHub = BridgeHub.attach(bridgeHubAddress) as any;

    console.log("检查管理员权限...");
    const ADMIN_ROLE = await bridgeHub.ADMIN_ROLE();
    const hasAdminRole = await bridgeHub.hasRole(ADMIN_ROLE, singer.address);

    if (!hasAdminRole) {
      throw new ConfigurationError(
        "权限不足",
        `账户 ${singer.address} 没有 ADMIN_ROLE 权限`,
        "请使用具有管理员权限的账户"
      );
    }
    console.log("✅ 权限验证通过");

    const currentBridgedToken = await bridgeHub.tokenPair(
      chainIdNum,
      srcTokenBytes32
    );

    if (currentBridgedToken !== ethers.ZeroHash) {
      throw new Error(
        `链 ${chainIdNum} 上的源 token ${srcTokenBytes32} 已经映射到 ${currentBridgedToken}`
      );
    }

    console.log("\n=== 执行 setTokenPair ===");
    const tx = await bridgeHub.setTokenPair(
      chainIdNum,
      srcTokenBytes32,
      tokenDecimalNum,
      bridgedTokenAddress
    );

    console.log("交易已提交，等待确认...");
    console.log("交易哈希:", tx.hash);

    const receipt = await tx.wait();
    console.log("✅ 交易确认成功!");
    console.log("区块号:", receipt?.blockNumber);

    if (pairMode === "evm") {
      const remoteToken = await bridgeHub.tokenPair(
        chainIdNum,
        ethers.zeroPadValue(bridgedTokenAddress, 32)
      );
      if (remoteToken.toLowerCase() !== srcTokenBytes32.toLowerCase()) {
        throw new ConfigurationError(
          "Withdraw token pair 设置失败",
          `期望: ${srcTokenBytes32}, 实际: ${remoteToken}`,
          "请检查交易是否成功执行"
        );
      }
      console.log("✅ Withdraw 映射已自动配置");
    }

    const newBridgedToken = await bridgeHub.tokenPair(
      chainIdNum,
      srcTokenBytes32
    );
    if (
      newBridgedToken.toLowerCase() !==
      ethers.zeroPadValue(bridgedTokenAddress, 32).toLowerCase()
    ) {
      throw new ConfigurationError(
        "Token Pair 设置失败",
        `期望: ${bridgedTokenAddress}, 实际: ${newBridgedToken}`,
        "请检查交易是否成功执行"
      );
    }
    console.log("✅ Token Pair 设置成功!");
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
