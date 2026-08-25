import { ethers } from "hardhat";
import {
  ConfigurationError,
  handleError,
  PermissionError,
  validateAddress,
} from "./utils/error-handler";

async function main() {
  console.log("开始设置代币提现费用...");

  const [singer] = await ethers.getSigners();
  console.log("操作账户:", singer.address);
  const balance = await ethers.provider.getBalance(singer.address);
  console.log("账户余额:", ethers.formatEther(balance));

  // 从环境变量读取参数（也支持从命令行读取，优先环境变量）
  const bridgeHubAddress = process.env.BRIDGE_HUB_ADDRESS || process.argv[2];
  const tokenAddress = process.env.BRIDGED_TOKEN_ADDRESS || process.argv[3];
  const fee = process.env.WITHDRAW_FEE || process.argv[4];

  if (!fee) {
    throw new ConfigurationError(
      "缺少提现费用参数",
      "请通过环境变量 WITHDRAW_FEE 或命令行参数提供",
      "示例: WITHDRAW_FEE=1000000000000000000 hardhat run scripts/set-withdraw-fee.ts"
    );
  }

  // 验证地址格式
  validateAddress(bridgeHubAddress, "BridgeHub 合约");
  validateAddress(tokenAddress, "代币合约");

  // 打印参数预览
  console.log("\n参数预览:");
  console.log(`  BridgeHub 合约地址: ${bridgeHubAddress}`);
  console.log(`  代币地址: ${tokenAddress}`);
  console.log(`  提现费用: ${fee} (wei)`);

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

    // 检查当前提现费用设置
    console.log("检查当前提现费用设置...");
    const currentFee = await bridgeHub.tokenWithdrawFee(tokenAddress);
    console.log(`当前提现费用: ${currentFee.toString()} (wei)`);

    if (BigInt(currentFee) === BigInt(fee)) {
      console.log("⚠️  提现费用已经是目标值，无需更改");
      return;
    }

    // 执行 setWithdrawFee 调用
    console.log("\n=== 执行 setWithdrawFee ===");
    const tx = await bridgeHub.setWithdrawFee(tokenAddress, BigInt(fee));

    console.log("交易已提交，等待确认...");
    console.log("交易哈希:", tx.hash);

    const receipt = await tx.wait();
    console.log("✅ 交易确认成功!");
    console.log("区块号:", receipt?.blockNumber);
    console.log("Gas 使用量:", receipt?.gasUsed?.toString());

    // 验证设置结果
    console.log("\n=== 验证设置结果 ===");
    const newFee = await bridgeHub.tokenWithdrawFee(tokenAddress);

    if (BigInt(newFee) == BigInt(fee)) {
      console.log("✅ 提现费用设置成功!");
      console.log(
        `代币 ${tokenAddress} 的提现费用已设置为 ${ethers.formatEther(
          BigInt(fee)
        )}`
      );
    } else {
      throw new ConfigurationError(
        "提现费用设置失败",
        `期望: ${BigInt(fee)}, 实际: ${newFee.toString()}`,
        "请检查交易是否成功执行"
      );
    }

    // 显示费用变化
    const feeChange = BigInt(fee) - BigInt(currentFee);
    if (feeChange > 0) {
      console.log(`📈 提现费用增加了 ${ethers.formatEther(feeChange)}`);
    } else if (feeChange < 0) {
      console.log(`📉 提现费用减少了 ${ethers.formatEther(feeChange)}`);
    }
  } catch (error) {
    if (error instanceof ConfigurationError) {
      throw error;
    }
    throw new PermissionError(
      "设置提现费用失败",
      error instanceof Error ? error.message : String(error),
      "请检查网络连接和合约状态，然后重试"
    );
  }
}

main()
  .then(() => {
    console.log("🎉 提现费用设置完成!");
    process.exit(0);
  })
  .catch(handleError);
