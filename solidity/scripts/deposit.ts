import { ethers } from "hardhat";
import { Bridge } from "../typechain-types";

async function main() {
  // 配置参数
  const BRIDGE_CONTRACT_ADDRESS = process.env.BRIDGE_CONTRACT_ADDRESS;
  const TOKEN_ADDRESS = process.env.TOKEN_ADDRESS;
  const DESTINATION_ADDRESS = process.env.DESTINATION_ADDRESS;
  const AMOUNT = process.env.AMOUNT || "1000000000000000000";

  if (!BRIDGE_CONTRACT_ADDRESS) {
    throw new Error("请设置环境变量 BRIDGE_CONTRACT_ADDRESS");
  }

  if (!TOKEN_ADDRESS) {
    throw new Error("请设置环境变量 TOKEN_ADDRESS");
  }

  if (!DESTINATION_ADDRESS) {
    throw new Error("请设置环境变量 DESTINATION_ADDRESS");
  }

  console.log("🚀 开始执行 Bridge Deposit 操作...");
  console.log("📋 配置信息:");
  console.log(`   Bridge 合约地址: ${BRIDGE_CONTRACT_ADDRESS}`);
  console.log(`   代币地址: ${TOKEN_ADDRESS}`);
  console.log(`   目标地址: ${DESTINATION_ADDRESS}`);
  console.log(`   存款金额: ${AMOUNT} wei`);

  // 获取签名者
  const [signer] = await ethers.getSigners();
  console.log(`👤 签名者地址: ${signer.address}`);

  // 检查账户余额
  const balance = await ethers.provider.getBalance(signer.address);
  console.log(`💰 账户 ETH 余额: ${ethers.formatEther(balance)} ETH`);

  // 获取 Bridge 合约实例
  const bridge = (await ethers.getContractAt(
    "Bridge",
    BRIDGE_CONTRACT_ADDRESS
  )) as Bridge;
  console.log("✅ Bridge 合约实例创建成功");

  // 获取代币合约实例 (假设是 ERC20)
  const token = await ethers.getContractAt("IERC20", TOKEN_ADDRESS);
  console.log("✅ 代币合约实例创建成功");

  // 检查代币余额
  const tokenBalance = await token.balanceOf(signer.address);
  console.log(`🪙 代币余额: ${tokenBalance} wei`);

  // 检查代币授权额度
  const allowance = await token.allowance(
    signer.address,
    BRIDGE_CONTRACT_ADDRESS
  );
  console.log(`🔐 当前授权额度: ${allowance} wei`);

  // 如果授权额度不足，先进行授权
  if (allowance < BigInt(AMOUNT)) {
    console.log("⚠️  授权额度不足，开始授权...");

    const approveTx = await token.approve(
      BRIDGE_CONTRACT_ADDRESS,
      ethers.MaxUint256
    );
    console.log(`📝 授权交易已发送，交易哈希: ${approveTx.hash}`);

    const approveReceipt = await approveTx.wait();
    console.log(`✅ 授权成功，区块号: ${approveReceipt?.blockNumber}`);
  }

  // 执行 deposit 操作
  console.log("💸 开始执行 deposit 操作...");

  const depositTx = await bridge.deposit(
    DESTINATION_ADDRESS,
    TOKEN_ADDRESS,
    AMOUNT
  );

  console.log(`📝 Deposit 交易已发送，交易哈希: ${depositTx.hash}`);
  console.log("⏳ 等待交易确认...");

  const depositReceipt = await depositTx.wait();
  console.log(`✅ Deposit 成功！`);
  console.log(`   区块号: ${depositReceipt?.blockNumber}`);
  console.log(`   Gas 使用量: ${depositReceipt?.gasUsed?.toString()}`);
  console.log("🎉 Bridge Deposit 操作完成！");
}

// 错误处理
main().catch((error) => {
  console.error("💥 脚本执行失败:");
  console.error(error);
  process.exitCode = 1;
});
