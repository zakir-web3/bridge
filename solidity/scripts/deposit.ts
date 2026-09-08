import { ethers } from "hardhat";
import { Bridge } from "../typechain-types";
import { sendTx } from "./utils/send-tx";

async function main() {
  const BRIDGE_CONTRACT_ADDRESS = process.env.BRIDGE_CONTRACT_ADDRESS;
  const TOKEN_ADDRESS = process.env.TOKEN_ADDRESS;
  const DESTINATION_ADDRESS = process.env.DESTINATION_ADDRESS;
  const AMOUNT = process.env.AMOUNT || "1000000000000000000";

  if (!BRIDGE_CONTRACT_ADDRESS) {
    throw new Error("Set environment variable BRIDGE_CONTRACT_ADDRESS");
  }
  if (!TOKEN_ADDRESS) {
    throw new Error("Set environment variable TOKEN_ADDRESS");
  }
  if (!DESTINATION_ADDRESS) {
    throw new Error("Set environment variable DESTINATION_ADDRESS");
  }

  const amount = BigInt(AMOUNT);
  const [signer] = await ethers.getSigners();

  console.log("🚀 Starting Bridge deposit...");
  console.log("📋 Config:");
  console.log(`   Bridge contract: ${BRIDGE_CONTRACT_ADDRESS}`);
  console.log(`   Token address: ${TOKEN_ADDRESS}`);
  console.log(`   Destination: ${DESTINATION_ADDRESS}`);
  console.log(`   Deposit amount: ${amount.toString()} wei`);
  console.log(`👤 Signer: ${signer.address}`);

  const balance = await ethers.provider.getBalance(signer.address);
  console.log(`💰 Account ETH balance: ${ethers.formatEther(balance)} ETH`);

  const bridge = (await ethers.getContractAt(
    "Bridge",
    BRIDGE_CONTRACT_ADDRESS,
    signer
  )) as Bridge;
  const token = await ethers.getContractAt("IERC20", TOKEN_ADDRESS, signer);

  const tokenBalance = await token.balanceOf(signer.address);
  console.log(`🪙 Token balance: ${tokenBalance} wei`);

  const allowance = await token.allowance(
    signer.address,
    BRIDGE_CONTRACT_ADDRESS
  );
  console.log(`🔐 Current allowance: ${allowance} wei`);

  if (allowance < amount) {
    console.log("⚠️  Allowance too low, approving...");
    const approveReceipt = await sendTx(signer, (nonce) =>
      token.approve(BRIDGE_CONTRACT_ADDRESS, ethers.MaxUint256, { nonce })
    );
    console.log(`✅ Approved, block: ${approveReceipt.blockNumber}`);
  }

  console.log("💸 Submitting deposit...");
  const depositReceipt = await sendTx(signer, (nonce) =>
    bridge.deposit(DESTINATION_ADDRESS, TOKEN_ADDRESS, amount, { nonce })
  );
  console.log(`✅ Deposit succeeded!`);
  console.log(`   Block: ${depositReceipt.blockNumber}`);
  console.log(`   Gas used: ${depositReceipt.gasUsed.toString()}`);
  console.log("🎉 Bridge deposit complete!");
}

main().catch((error) => {
  console.error("💥 Script failed:");
  console.error(error);
  process.exit(1);
});
