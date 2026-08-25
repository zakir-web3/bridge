import { ethers } from "hardhat";
import { BridgeHub } from "../typechain-types";

async function main() {
  console.log("🚀 开始执行 BridgeHub Withdraw 操作...");
  // 获取签名者
  const [signer] = await ethers.getSigners();
  console.log(`👤 签名者地址: ${signer.address}`);

  // 检查账户余额
  const balance = await ethers.provider.getBalance(signer.address);
  console.log(`💰 账户原生代币余额: ${ethers.formatEther(balance)}`);

  // 配置参数
  const BRIDGE_HUB_CONTRACT_ADDRESS = process.env.BRIDGE_HUB_CONTRACT_ADDRESS;
  const TOKEN_ADDRESS = process.env.TOKEN_ADDRESS;
  const DESTINATION_ADDRESS = process.env.DESTINATION_ADDRESS;
  const AMOUNT = process.env.AMOUNT || "1000000000000000000"; // 默认 1 ETH (18位小数)
  const CHAIN_ID = process.env.CHAIN_ID || "56"; // 默认 BSC 主网

  if (!BRIDGE_HUB_CONTRACT_ADDRESS) {
    throw new Error("请设置环境变量 BRIDGE_HUB_CONTRACT_ADDRESS");
  }

  if (!TOKEN_ADDRESS) {
    throw new Error("请设置环境变量 TOKEN_ADDRESS");
  }

  if (!DESTINATION_ADDRESS) {
    throw new Error("请设置环境变量 DESTINATION_ADDRESS");
  }

  console.log("📋 配置信息:");
  console.log(`   BridgeHub 合约地址: ${BRIDGE_HUB_CONTRACT_ADDRESS}`);
  console.log(`   代币地址: ${TOKEN_ADDRESS}`);
  console.log(`   目标地址: ${DESTINATION_ADDRESS}`);
  console.log(`   提款金额: ${AMOUNT} wei`);
  console.log(`   目标链 ID: ${CHAIN_ID}`);

  // 获取 BridgeHub 合约实例
  const bridgeHub = (await ethers.getContractAt(
    "BridgeHub",
    BRIDGE_HUB_CONTRACT_ADDRESS
  )) as BridgeHub;
  console.log("✅ BridgeHub 合约实例创建成功");

  // 获取代币合约实例 (假设是 ERC20)
  const token = await ethers.getContractAt("IERC20", TOKEN_ADDRESS);
  console.log("✅ 代币合约实例创建成功");

  // 检查代币余额
  const tokenBalance = await token.balanceOf(signer.address);
  console.log(`🪙 代币余额: ${tokenBalance} wei`);

  // 检查代币授权额度
  const allowance = await token.allowance(
    signer.address,
    BRIDGE_HUB_CONTRACT_ADDRESS
  );
  console.log(`🔐 当前授权额度: ${allowance} wei`);

  // 检查提款费用
  const withdrawFee = await bridgeHub.tokenWithdrawFee(TOKEN_ADDRESS);
  console.log(`💸 提款费用: ${withdrawFee} wei`);

  // 验证提款金额是否大于费用
  if (BigInt(AMOUNT) <= withdrawFee) {
    throw new Error(
      `提款金额 ${AMOUNT} wei 必须大于提款费用 ${withdrawFee} wei`
    );
  }

  // 检查代币对配置
  const bridgeToken = await bridgeHub.tokenPair(CHAIN_ID, TOKEN_ADDRESS);
  if (bridgeToken === ethers.ZeroAddress) {
    throw new Error(`链 ID ${CHAIN_ID} 上的代币 ${TOKEN_ADDRESS} 未配置代币对`);
  }
  console.log(`🔗 桥接代币地址: ${bridgeToken}`);

  // 如果授权额度不足，先进行授权
  if (allowance < BigInt(AMOUNT)) {
    console.log("⚠️  授权额度不足，开始授权...");

    const approveTx = await token.approve(
      BRIDGE_HUB_CONTRACT_ADDRESS,
      ethers.MaxUint256
    );
    console.log(`📝 授权交易已发送，交易哈希: ${approveTx.hash}`);

    const approveReceipt = await approveTx.wait();
    console.log(`✅ 授权成功，区块号: ${approveReceipt?.blockNumber}`);
  }

  // 执行 withdraw 操作
  console.log("💸 开始执行 withdraw 操作...");

  const withdrawTx = await bridgeHub.withdraw(
    DESTINATION_ADDRESS,
    TOKEN_ADDRESS,
    AMOUNT,
    CHAIN_ID
  );

  console.log(`📝 Withdraw 交易已发送，交易哈希: ${withdrawTx.hash}`);
  console.log("⏳ 等待交易确认...");

  const withdrawReceipt = await withdrawTx.wait();
  console.log(`✅ Withdraw 成功！`);
  console.log(`   区块号: ${withdrawReceipt?.blockNumber}`);
  console.log(`   Gas 使用量: ${withdrawReceipt?.gasUsed?.toString()}`);

  // 获取当前提款 nonce
  const currentWithdrawNonce = await bridgeHub.withdrawNonce();
  console.log(`🔢 当前提款 Nonce: ${currentWithdrawNonce}`);

  console.log("🎉 BridgeHub Withdraw 操作完成！");
  console.log("📝 注意: 提款请求已提交，需要等待验证者确认后才能完成跨链转账");
}

// 错误处理
main().catch((error) => {
  console.error("💥 脚本执行失败:");
  console.error(error);
  process.exitCode = 1;
});
