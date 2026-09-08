import { ethers } from "hardhat";
import { Bridge } from "../typechain-types";

async function main() {
  // Config
  const BRIDGE_CONTRACT_ADDRESS = process.env.BRIDGE_CONTRACT_ADDRESS;
  const PRIVATE_KEY = process.env.PRIVATE_KEY;
  const WITHDRAWAL_MESSAGES = process.env.WITHDRAWAL_MESSAGES; // optional comma-separated withdrawal messages

  if (!BRIDGE_CONTRACT_ADDRESS) {
    throw new Error("Set environment variable BRIDGE_CONTRACT_ADDRESS");
  }

  if (!PRIVATE_KEY) {
    throw new Error("Set environment variable PRIVATE_KEY");
  }

  console.log("🚀 Starting batched finalize withdrawals...");
  console.log("📋 Config:");
  console.log(`   Bridge contract: ${BRIDGE_CONTRACT_ADDRESS}`);

  // Wallet
  const [signer] = await ethers.getSigners();
  console.log(`👤 Validator: ${signer.address}`);

  // Network
  const network = await ethers.provider.getNetwork();
  console.log(`🌐 Network: ${network.name} (Chain ID: ${network.chainId})`);

  // Balance
  const balance = await ethers.provider.getBalance(signer.address);
  console.log(`💰 Account ETH balance: ${ethers.formatEther(balance)} ETH`);

  if (balance === 0n) {
    throw new Error("Insufficient balance to pay gas");
  }

  // Bridge contract instance
  const bridge = (await ethers.getContractAt(
    "Bridge",
    BRIDGE_CONTRACT_ADDRESS
  )) as Bridge;
  console.log("✅ Bridge contract instance created");

  // Check finalizer role
  const isFinalizer = await bridge.finalizers(signer.address);
  if (!isFinalizer) {
    throw new Error(
      `Address ${signer.address} is not a finalizer and cannot finalize`
    );
  }
  console.log("✅ Validator identity confirmed: has finalizer permission");

  // Dispute period
  const disputePeriodSeconds = await bridge.disputePeriodSeconds();
  const blockDurationMillis = await bridge.blockDurationMillis();
  console.log(
    `⏰ Dispute period: ${disputePeriodSeconds}s, block duration: ${blockDurationMillis}ms`
  );

  // Pending withdrawal messages
  let withdrawalMessages: string[] = [];

  if (WITHDRAWAL_MESSAGES) {
    // Use caller-provided messages
    withdrawalMessages = WITHDRAWAL_MESSAGES.split(",").map((msg) =>
      msg.trim()
    );
    console.log(`📝 Using provided withdrawal messages: ${withdrawalMessages.length}`);
  } else {
    // No messages provided; the contract has no query for all pending withdrawals
    console.log("🔍 Querying pending withdrawal requests...");

    // Note: the contract has no direct query for all pending withdrawals.
    // Callers should supply the messages to finalize.
    console.log(
      "⚠️  No withdrawal messages provided; set the WITHDRAWAL_MESSAGES env var"
    );
    console.log("   Format: WITHDRAWAL_MESSAGES=0x123...,0x456...");
    console.log("   Or collect withdrawal messages from contract events in a block explorer:");
    throw new Error("Please provide withdrawal messages to finalize");
  }

  if (withdrawalMessages.length === 0) {
    throw new Error("No withdrawal messages provided");
  }

  // Validate message format and status
  for (let i = 0; i < withdrawalMessages.length; i++) {
    const message = withdrawalMessages[i];
    if (!ethers.isHexString(message, 32)) {
      throw new Error(`Withdrawal message ${i + 1} has invalid format: ${message}`);
    }

    // Skip already finalized withdrawals
    const isFinalized = await bridge.finalizedWithdrawals(message);
    if (isFinalized) {
      console.log(`⚠️  Withdrawal message ${i + 1} already finalized: ${message}`);
      withdrawalMessages.splice(i, 1);
      i--;
      continue;
    }

    // Fetch withdrawal details
    try {
      const withdrawal = await bridge.requestedWithdrawals(message);
      if (withdrawal.requestedTime === 0n) {
        console.log(`⚠️  Withdrawal message ${i + 1} does not exist: ${message}`);
        withdrawalMessages.splice(i, 1);
        i--;
        continue;
      }

      // Check dispute period has elapsed
      const currentTime = BigInt(Math.floor(Date.now() / 1000));
      const disputePeriodSeconds = await bridge.disputePeriodSeconds();
      const timePassed = currentTime - withdrawal.requestedTime;

      if (timePassed < disputePeriodSeconds) {
        const remainingTime = disputePeriodSeconds - timePassed;
        console.log(
          `⚠️  Withdrawal message ${
            i + 1
          } still in dispute period, wait ${remainingTime}s: ${message}`
        );
        withdrawalMessages.splice(i, 1);
        i--;
        continue;
      }

      console.log(`📋 Withdrawal ${i + 1} details:`);
      console.log(`   User: ${withdrawal.user}`);
      console.log(`   Destination: ${withdrawal.destination}`);
      console.log(`   Token: ${withdrawal.token}`);
      console.log(`   Amount: ${withdrawal.amount} wei`);
      console.log(`   Chain ID: ${withdrawal.chainId}`);
      console.log(`   Nonce: ${withdrawal.nonce}`);
      console.log(
        `   Requested at: ${new Date(
          Number(withdrawal.requestedTime) * 1000
        ).toISOString()}`
      );
    } catch (error) {
      console.log(`⚠️  Could not fetch details for withdrawal message ${i + 1}: ${message}`);
      console.log(`   Error: ${error}`);
    }
  }

  if (withdrawalMessages.length === 0) {
    console.log("✅ All withdrawal requests are already finalized");
    return;
  }

  console.log(`📋 Pending withdrawals to finalize: ${withdrawalMessages.length}`);
  withdrawalMessages.forEach((msg, index) => {
    console.log(`   ${index + 1}. ${msg}`);
  });

  // Execute batched finalize
  console.log("💸 Finalizing withdrawals...");

  const finalizeTx = await bridge.batchedFinalizeWithdrawals(
    withdrawalMessages
  );

  console.log(`📝 Finalize tx sent, hash: ${finalizeTx.hash}`);
  console.log("⏳ Waiting for confirmation...");

  const receipt = await finalizeTx.wait();
  console.log(`✅ Finalize tx succeeded!`);
  console.log(`   Block: ${receipt?.blockNumber}`);
  console.log(`   Gas used: ${receipt?.gasUsed?.toString()}`);
  console.log(
    `   Gas price: ${ethers.formatUnits(receipt?.gasPrice || 0, "gwei")} Gwei`
  );
  console.log(
    `   Total cost: ${ethers.formatEther(
      (receipt?.gasUsed || 0n) * (receipt?.gasPrice || 0n)
    )} ETH`
  );

  // Verify finalized status
  console.log("🔍 Verifying finalized status...");
  for (let i = 0; i < withdrawalMessages.length; i++) {
    const message = withdrawalMessages[i];
    const isFinalized = await bridge.finalizedWithdrawals(message);
    if (isFinalized) {
      console.log(`   ✅ Withdrawal ${i + 1} finalized: ${message}`);
    } else {
      console.log(`   ❌ Withdrawal ${i + 1} not finalized: ${message}`);
    }
  }

  console.log("🎉 Finalize withdrawals complete!");
  console.log("📝 Note: eligible withdrawal requests are finalized; funds have been transferred to the destination");
}

// Error handling
main().catch((error) => {
  console.error("💥 Script failed:");
  console.error(error);
  process.exitCode = 1;
});
