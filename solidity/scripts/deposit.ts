import { ethers } from "hardhat";
import { Bridge } from "../typechain-types";
import { sendTx } from "./utils/send-tx";

async function main() {
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

  const amount = BigInt(AMOUNT);
  const [signer] = await ethers.getSigners();

  console.log("🚀 开始执行 Bridge Deposit 操作...");
  console.log("📋 配置信息:");
  console.log(`   Bridge 合约地址: ${BRIDGE_CONTRACT_ADDRESS}`);
  console.log(`   代币地址: ${TOKEN_ADDRESS}`);
  console.log(`   目标地址: ${DESTINATION_ADDRESS}`);
  console.log(`   存款金额: ${amount.toString()} wei`);
  console.log(`👤 签名者地址: ${signer.address}`);

  const balance = await ethers.provider.getBalance(signer.address);
  console.log(`💰 账户 ETH 余额: ${ethers.formatEther(balance)} ETH`);

  const bridge = (await ethers.getContractAt(
    "Bridge",
    BRIDGE_CONTRACT_ADDRESS,
    signer
  )) as Bridge;
  const token = await ethers.getContractAt("IERC20", TOKEN_ADDRESS, signer);

  const tokenBalance = await token.balanceOf(signer.address);
  console.log(`🪙 代币余额: ${tokenBalance} wei`);

  const allowance = await token.allowance(
    signer.address,
    BRIDGE_CONTRACT_ADDRESS
  );
  console.log(`🔐 当前授权额度: ${allowance} wei`);

  if (allowance < amount) {
    console.log("⚠️  授权额度不足，开始授权...");
    const approveReceipt = await sendTx(signer, (nonce) =>
      token.approve(BRIDGE_CONTRACT_ADDRESS, ethers.MaxUint256, { nonce })
    );
    console.log(`✅ 授权成功，区块号: ${approveReceipt.blockNumber}`);
  }

  console.log("💸 开始执行 deposit 操作...");
  const depositReceipt = await sendTx(signer, (nonce) =>
    bridge.deposit(DESTINATION_ADDRESS, TOKEN_ADDRESS, amount, { nonce })
  );
  console.log(`✅ Deposit 成功！`);
  console.log(`   区块号: ${depositReceipt.blockNumber}`);
  console.log(`   Gas 使用量: ${depositReceipt.gasUsed.toString()}`);
  console.log("🎉 Bridge Deposit 操作完成！");
}

main().catch((error) => {
  console.error("💥 脚本执行失败:");
  console.error(error);
  process.exit(1);
});
