import { ethers } from "hardhat";
import { BridgeHub } from "../typechain-types";
import { sendTx } from "./utils/send-tx";

async function main() {
  console.log("🚀 开始执行 BridgeHub Withdraw 操作...");
  const [signer] = await ethers.getSigners();
  console.log(`👤 签名者地址: ${signer.address}`);

  const balance = await ethers.provider.getBalance(signer.address);
  console.log(`💰 账户原生代币余额: ${ethers.formatEther(balance)}`);

  const BRIDGE_HUB_CONTRACT_ADDRESS = process.env.BRIDGE_HUB_CONTRACT_ADDRESS;
  const TOKEN_ADDRESS = process.env.TOKEN_ADDRESS;
  const DESTINATION_ADDRESS = process.env.DESTINATION_ADDRESS;
  const AMOUNT = process.env.AMOUNT || "1000000000000000000";
  const CHAIN_ID =
    process.env.TOKEN_CHAIN_ID || process.env.CHAIN_ID || "56";

  if (!BRIDGE_HUB_CONTRACT_ADDRESS) {
    throw new Error("请设置环境变量 BRIDGE_HUB_CONTRACT_ADDRESS");
  }
  if (!TOKEN_ADDRESS) {
    throw new Error("请设置环境变量 TOKEN_ADDRESS");
  }
  if (!DESTINATION_ADDRESS) {
    throw new Error("请设置环境变量 DESTINATION_ADDRESS");
  }

  const amount = BigInt(AMOUNT);
  const chainId = BigInt(CHAIN_ID);

  console.log("📋 配置信息:");
  console.log(`   BridgeHub 合约地址: ${BRIDGE_HUB_CONTRACT_ADDRESS}`);
  console.log(`   代币地址: ${TOKEN_ADDRESS}`);
  console.log(`   目标地址: ${DESTINATION_ADDRESS}`);
  console.log(`   提款金额: ${amount.toString()} wei`);
  console.log(`   目标链 ID: ${chainId.toString()}`);

  const bridgeHub = (await ethers.getContractAt(
    "BridgeHub",
    BRIDGE_HUB_CONTRACT_ADDRESS,
    signer
  )) as BridgeHub;
  const token = await ethers.getContractAt("IERC20", TOKEN_ADDRESS, signer);

  const tokenBalance = await token.balanceOf(signer.address);
  console.log(`🪙 代币余额: ${tokenBalance} wei`);

  const allowance = await token.allowance(
    signer.address,
    BRIDGE_HUB_CONTRACT_ADDRESS
  );
  console.log(`🔐 当前授权额度: ${allowance} wei`);

  const withdrawFee = await bridgeHub.tokenWithdrawFee(TOKEN_ADDRESS);
  console.log(`💸 提款费用: ${withdrawFee} wei`);
  if (amount <= withdrawFee) {
    throw new Error(
      `提款金额 ${amount.toString()} wei 必须大于提款费用 ${withdrawFee.toString()} wei`
    );
  }

  const bridgeToken = await bridgeHub.tokenPair(chainId, TOKEN_ADDRESS);
  if (bridgeToken === ethers.ZeroAddress) {
    throw new Error(
      `链 ID ${chainId.toString()} 上的代币 ${TOKEN_ADDRESS} 未配置代币对`
    );
  }
  console.log(`🔗 桥接代币地址: ${bridgeToken}`);

  if (allowance < amount) {
    console.log("⚠️  授权额度不足，开始授权...");
    const approveReceipt = await sendTx(signer, (nonce) =>
      token.approve(BRIDGE_HUB_CONTRACT_ADDRESS, ethers.MaxUint256, { nonce })
    );
    console.log(`✅ 授权成功，区块号: ${approveReceipt.blockNumber}`);
  }

  console.log("💸 开始执行 withdraw 操作...");
  const withdrawReceipt = await sendTx(signer, (nonce) =>
    bridgeHub.withdraw(DESTINATION_ADDRESS, TOKEN_ADDRESS, amount, chainId, {
      nonce,
    })
  );
  console.log(`✅ Withdraw 成功！`);
  console.log(`   区块号: ${withdrawReceipt.blockNumber}`);
  console.log(`   Gas 使用量: ${withdrawReceipt.gasUsed.toString()}`);

  const currentWithdrawNonce = await bridgeHub.withdrawNonce();
  console.log(`🔢 当前提款 Nonce: ${currentWithdrawNonce}`);
  console.log("🎉 BridgeHub Withdraw 操作完成！");
  console.log("📝 注意: 提款请求已提交，需要等待验证者确认后才能完成跨链转账");
}

main().catch((error) => {
  console.error("💥 脚本执行失败:");
  console.error(error);
  process.exit(1);
});
