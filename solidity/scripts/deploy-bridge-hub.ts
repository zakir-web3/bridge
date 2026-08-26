import { ethers } from "hardhat";
import {
  DeploymentError,
  handleError,
  validateAddressArray,
  validateArrayLength,
} from "./utils/error-handler";

async function main() {
  console.log("开始部署 BridgeHub 合约...");

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

  // 使用统一的验证函数
  validateAddressArray(hotAddresses, "验证者");
  validateAddressArray(coldAddresses, "冷钱包");
  validateArrayLength(hotAddresses, powers, "验证者地址", "权重");
  validateArrayLength(hotAddresses, coldAddresses, "验证者地址", "冷钱包地址");

  // 打印参数预览
  console.log("\n参数预览:");
  for (let i = 0; i < hotAddresses.length; i++) {
    console.log(
      `  验证者 #${i + 1}: hot=${hotAddresses[i]}, cold=${
        coldAddresses[i]
      }, power=${powers[i]}`
    );
  }

  try {
    // 部署 BridgeHub 实现合约 + 通过 ERC1967Proxy 代理进行初始化
    console.log("\n=== 部署 BridgeHub (UUPS 可升级) ===");
    const BridgeHub = await ethers.getContractFactory("BridgeHub");
    const impl = await BridgeHub.deploy();
    await impl.waitForDeployment();
    const implAddress = await impl.getAddress();
    // console.log("BridgeHub 实现合约地址:", implAddress);

    const initData = BridgeHub.interface.encodeFunctionData("initialize", [
      hotAddresses,
      coldAddresses,
      powers,
    ]);

    const ERC1967Proxy = await ethers.getContractFactory("ERC1967Proxy");
    const proxy = await ERC1967Proxy.deploy(implAddress, initData);
    await proxy.waitForDeployment();
    const bridgeHubAddress = await proxy.getAddress();
    console.log("BridgeHub 代理地址:", bridgeHubAddress);
    console.log(`DEPLOYED_ADDRESS=${bridgeHubAddress}`);

    // 通过实现合约的 ABI 连接到代理地址（为避免 TS 类型告警，使用 any）
    const bridgeHub = BridgeHub.attach(bridgeHubAddress) as any;
    const domainSeparator = await bridgeHub.domainSeparator();
    console.log("BridgeHub 域分隔符:", domainSeparator);
    console.log("✅ BridgeHub 部署成功!");
  } catch (error) {
    throw new DeploymentError(
      "BridgeHub 合约部署失败",
      error instanceof Error ? error.message : String(error),
      "请检查网络连接和账户余额，然后重试"
    );
  }
}

main()
  .then(() => {
    console.log("🎉 部署完成!");
    process.exit(0);
  })
  .catch(handleError);
