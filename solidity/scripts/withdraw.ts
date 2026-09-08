import { ethers } from "hardhat";
import { BridgeHub } from "../typechain-types";
import { sendTx } from "./utils/send-tx";

async function main() {
  console.log("🚀 Starting BridgeHub withdraw...");
  const [signer] = await ethers.getSigners();
  console.log(`👤 Signer: ${signer.address}`);

  const balance = await ethers.provider.getBalance(signer.address);
  console.log(`💰 Native token balance: ${ethers.formatEther(balance)}`);

  const BRIDGE_HUB_CONTRACT_ADDRESS = process.env.BRIDGE_HUB_CONTRACT_ADDRESS;
  const TOKEN_ADDRESS = process.env.TOKEN_ADDRESS;
  const DESTINATION_ADDRESS = process.env.DESTINATION_ADDRESS;
  const AMOUNT = process.env.AMOUNT || "1000000000000000000";
  const CHAIN_ID = process.env.TOKEN_CHAIN_ID || process.env.CHAIN_ID || "56";

  if (!BRIDGE_HUB_CONTRACT_ADDRESS) {
    throw new Error("Set environment variable BRIDGE_HUB_CONTRACT_ADDRESS");
  }
  if (!TOKEN_ADDRESS) {
    throw new Error("Set environment variable TOKEN_ADDRESS");
  }
  if (!DESTINATION_ADDRESS) {
    throw new Error("Set environment variable DESTINATION_ADDRESS");
  }

  const amount = BigInt(AMOUNT);
  const chainId = BigInt(CHAIN_ID);

  console.log("📋 Config:");
  console.log(`   BridgeHub contract: ${BRIDGE_HUB_CONTRACT_ADDRESS}`);
  console.log(`   Token address: ${TOKEN_ADDRESS}`);
  console.log(`   Destination: ${DESTINATION_ADDRESS}`);
  console.log(`   Withdraw amount: ${amount.toString()} wei`);
  console.log(`   Destination chain ID: ${chainId.toString()}`);

  const bridgeHub = (await ethers.getContractAt(
    "BridgeHub",
    BRIDGE_HUB_CONTRACT_ADDRESS,
    signer
  )) as BridgeHub;
  const token = await ethers.getContractAt("IERC20", TOKEN_ADDRESS, signer);

  const tokenBalance = await token.balanceOf(signer.address);
  console.log(`🪙 Token balance: ${tokenBalance} wei`);

  const allowance = await token.allowance(
    signer.address,
    BRIDGE_HUB_CONTRACT_ADDRESS
  );
  console.log(`🔐 Current allowance: ${allowance} wei`);

  const withdrawFee = await bridgeHub.tokenWithdrawFee(TOKEN_ADDRESS);
  console.log(`💸 Withdraw fee: ${withdrawFee} wei`);
  if (amount <= withdrawFee) {
    throw new Error(
      `Withdraw amount ${amount.toString()} wei must be greater than the withdraw fee ${withdrawFee.toString()} wei`
    );
  }

  const bridgeTokenBytes32 = await bridgeHub.tokenPair(
    chainId,
    ethers.zeroPadValue(TOKEN_ADDRESS, 32)
  );
  if (bridgeTokenBytes32 === ethers.ZeroHash) {
    throw new Error(
      `Token ${TOKEN_ADDRESS} on chain ID ${chainId.toString()} has no token pair configured`
    );
  }
  const bridgeToken = ethers.getAddress("0x" + bridgeTokenBytes32.slice(26));
  console.log(`🔗 Bridged token address: ${bridgeToken}`);

  if (allowance < amount) {
    console.log("⚠️  Allowance too low, approving...");
    const approveReceipt = await sendTx(signer, (nonce) =>
      token.approve(BRIDGE_HUB_CONTRACT_ADDRESS, ethers.MaxUint256, { nonce })
    );
    console.log(`✅ Approved, block: ${approveReceipt.blockNumber}`);
  }

  console.log("💸 Submitting withdraw...");
  const withdrawReceipt = await sendTx(signer, (nonce) =>
    bridgeHub.withdraw(
      ethers.zeroPadValue(DESTINATION_ADDRESS, 32),
      TOKEN_ADDRESS,
      amount,
      chainId,
      { nonce }
    )
  );
  console.log(`✅ Withdraw succeeded!`);
  console.log(`   Block: ${withdrawReceipt.blockNumber}`);
  console.log(`   Gas used: ${withdrawReceipt.gasUsed.toString()}`);

  const currentWithdrawNonce = await bridgeHub.withdrawNonce();
  console.log(`🔢 Current withdraw nonce: ${currentWithdrawNonce}`);
  console.log("🎉 BridgeHub withdraw complete!");
  console.log("📝 Note: the withdraw request is submitted; wait for validator confirmation to complete the cross-chain transfer");
}

main().catch((error) => {
  console.error("💥 Script failed:");
  console.error(error);
  process.exit(1);
});
