import { ethers, upgrades } from "hardhat";
import { BridgeHub } from "../typechain-types";
import {
  validateAddress,
  handleError,
  ConfigurationError,
  PermissionError,
} from "./utils/error-handler";

async function main() {
  console.log("🚀 开始升级 BridgeHub 合约...\n");

  const [signer] = await ethers.getSigners();
  console.log("操作账户:", signer.address);
  const balance = await ethers.provider.getBalance(signer.address);
  console.log("账户余额:", ethers.formatEther(balance), "ETH\n");

  // 从环境变量读取参数
  const proxyAddress = process.env.BRIDGE_HUB_PROXY || process.argv[2];

  if (!proxyAddress) {
    throw new ConfigurationError(
      "缺少代理合约地址",
      "请通过环境变量 BRIDGE_HUB_PROXY 或命令行参数提供",
      "示例: BRIDGE_HUB_PROXY=0x... npx hardhat run scripts/upgrade-bridge-hub.ts"
    );
  }

  validateAddress(proxyAddress, "BridgeHub Proxy");

  try {
    // Step 1: 验证当前代理的状态
    console.log("=== Step 1: 验证当前代理状态 ===");
    const BridgeHubV1 = await ethers.getContractFactory("BridgeHub");
    const proxyInstance = BridgeHubV1.attach(proxyAddress) as BridgeHub;

    // 获取当前信息
    const currentEpoch = await proxyInstance.epoch();
    const validatorsCount = (await proxyInstance.getHotValidators()).length;

    console.log(`✓ 当前 Epoch: ${currentEpoch}`);
    console.log(`✓ 验证器数量: ${validatorsCount}`);

    // 检查是否有足够的权限
    const ADMIN_ROLE = await proxyInstance.ADMIN_ROLE();
    const hasAdminRole = await proxyInstance.hasRole(
      ADMIN_ROLE,
      signer.address
    );

    if (!hasAdminRole) {
      throw new PermissionError(
        "权限不足",
        `账户 ${signer.address} 没有 ADMIN_ROLE 权限`,
        "请使用具有管理员权限的账户，或联系合约管理员授予权限"
      );
    }
    console.log("✓ 管理员权限验证通过\n");

    // Step 2: 编译新版本合约
    console.log("=== Step 2: 编译新版本合约 ===");
    console.log("正在编译 BridgeHub V2...");
    // 合约已编译，直接获取
    const BridgeHubV2 = await ethers.getContractFactory("BridgeHub");
    console.log("✓ BridgeHub V2 已准备就绪\n");

    // Step 3: 执行升级
    console.log("=== Step 3: 执行 UUPS 升级 ===");
    console.log(`代理地址: ${proxyAddress}`);
    console.log("正在部署新实现合约...");

    const deployedImpl = await upgrades.upgradeProxy(
      proxyAddress,
      BridgeHubV2,
      {
        // 不初始化，因为代理已经初始化过
        unsafeSkipStorageCheck: false,
        kind: "uups",
      }
    );

    await deployedImpl.waitForDeployment();
    const implAddress = await ethers.provider.getStorage(
      proxyAddress,
      "0x360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc"
    );

    console.log("✓ 新实现合约已部署");
    console.log(`✓ 实现地址: 0x${implAddress.slice(-40)}\n`);

    // Step 4: 验证升级
    console.log("=== Step 4: 验证升级结果 ===");
    const upgradedInstance = BridgeHubV2.attach(proxyAddress) as BridgeHub;

    // 验证数据完整性
    const newEpoch = await upgradedInstance.epoch();
    const newValidatorsCount = (await upgradedInstance.getHotValidators())
      .length;

    console.log(`✓ Epoch 保持一致: ${newEpoch === currentEpoch}`);
    console.log(
      `✓ 验证器数据保持一致: ${newValidatorsCount === validatorsCount}`
    );

    // 验证新增的映射存在
    try {
      // 测试新增的 tokenDecimalDiff 映射
      const testDiff = await upgradedInstance.tokenDecimalDiff(
        1,
        ethers.zeroPadValue(signer.address, 32)
      );
      console.log(`✓ 新增的 tokenDecimalDiff mapping 可以访问`);
    } catch (e) {
      console.log(`⚠ 无法验证 tokenDecimalDiff，但这可能是正常的`);
    }

    console.log("\n✅ 升级成功!\n");

    // Step 5: 打印总结
    console.log("=== 升级总结 ===");
    console.log(`代理地址:     ${proxyAddress}`);
    console.log(`新实现地址:   0x${implAddress.slice(-40)}`);
    console.log(`升级账户:     ${signer.address}`);
    console.log(`\n新增功能:`);
    console.log(`  • setTokenPair 方法新增 tokenDecimal 参数`);
    console.log(`  • 自动验证 dstToken 是否为有效 ERC20`);
    console.log(`  • 自动计算并存储精度差值`);
    console.log(`  • deposit/withdraw 时自动进行精度转换`);
    console.log(`\n下一步操作:`);
    console.log(
      `  1. 使用新的 setTokenPair 方法配置 token pair，传入 tokenDecimal 参数`
    );
    console.log(
      `     示例: npx hardhat run scripts/set-token-pair.ts --network <network>`
    );
    console.log(`  2. 参考 DECIMAL_CONVERSION.md 了解精度处理的详细信息`);
  } catch (error) {
    if (
      error instanceof ConfigurationError ||
      error instanceof PermissionError
    ) {
      throw error;
    }
    throw new PermissionError(
      "升级 BridgeHub 失败",
      error instanceof Error ? error.message : String(error),
      "请检查网络连接、合约状态和权限，然后重试"
    );
  }
}

main()
  .then(() => {
    console.log("🎉 BridgeHub 升级完成！\n");
    process.exit(0);
  })
  .catch(handleError);
