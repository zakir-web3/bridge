// @ts-nocheck
import { ethers } from "hardhat";
import {
  DeploymentError,
  handleError,
  validateAddressArray,
  validateArrayLength,
  validateNumber,
} from "./utils/error-handler";

async function main() {
  console.log("开始部署 Bridge 合约...");

  const [deployer] = await ethers.getSigners();
  console.log("部署账户:", deployer.address);
  const balance = await ethers.provider.getBalance(deployer.address);
  console.log("账户余额:", ethers.formatEther(balance));

  // 从环境变量读取参数（也支持从命令行读取，优先环境变量）
  const hotAddresses = (process.env.HOT_ADDRESSES || process.argv[2] || "")
    .split(",")
    .map((s) => s.trim())
    .filter((s) => s.length > 0);

  const coldAddresses = (process.env.COLD_ADDRESSES || process.argv[3] || "")
    .split(",")
    .map((s) => s.trim())
    .filter((s) => s.length > 0);

  const powers = (process.env.POWERS || process.argv[4] || "")
    .split(",")
    .map((s) => s.trim())
    .filter((s) => s.length > 0)
    .map((s) => parseInt(s, 10));

  const disputePeriodSeconds = parseInt(
    process.env.DISPUTE_PERIOD_SECONDS || process.argv[5] || "200",
    10
  );
  const blockDurationMillis = parseInt(
    process.env.BLOCK_DURATION_MILLIS || process.argv[6] || "750",
    10
  );
  const lockerThreshold = parseInt(
    process.env.LOCKER_THRESHOLD || process.argv[7] || "1",
    10
  );

  // 使用统一的验证函数
  validateAddressArray(hotAddresses, "热验证者");
  validateAddressArray(coldAddresses, "冷验证者");
  validateArrayLength(
    hotAddresses,
    coldAddresses,
    "热验证者地址",
    "冷验证者地址"
  );
  validateArrayLength(powers, hotAddresses, "权重", "验证者地址");

  // 验证数值参数
  validateNumber(disputePeriodSeconds, "争议期", 1);
  validateNumber(blockDurationMillis, "区块时长", 100);
  validateNumber(lockerThreshold, "锁定阈值", 1);

  // 打印参数预览
  console.log("\n参数预览:");
  for (let i = 0; i < hotAddresses.length; i++) {
    console.log(
      `  验证者 #${i + 1}: hot=${hotAddresses[i]}, cold=${
        coldAddresses[i]
      }, power=${powers[i]}`
    );
  }
  console.log("  disputePeriodSeconds:", disputePeriodSeconds);
  console.log("  blockDurationMillis:", blockDurationMillis);
  console.log("  lockerThreshold:", lockerThreshold);

  try {
    const Bridge = await ethers.getContractFactory("Bridge");
    const bridge = await Bridge.deploy(
      hotAddresses,
      coldAddresses,
      powers,
      disputePeriodSeconds,
      blockDurationMillis,
      lockerThreshold
    );
    await bridge.waitForDeployment();
    const bridgeAddress = await bridge.getAddress();
    console.log("Bridge 部署地址:", bridgeAddress);
    console.log(`DEPLOYED_ADDRESS=${bridgeAddress}`);
    const domainSeparator = await bridge.domainSeparator();
    console.log("Bridge 域分隔符:", domainSeparator);
    console.log("✅ Bridge 部署成功!");
  } catch (error) {
    throw new DeploymentError(
      "Bridge 合约部署失败",
      error instanceof Error ? error.message : String(error),
      "请检查网络连接、账户余额和合约参数，然后重试"
    );
  }
}

main()
  .then(() => {
    console.log("🎉 部署完成!");
    process.exit(0);
  })
  .catch(handleError);
