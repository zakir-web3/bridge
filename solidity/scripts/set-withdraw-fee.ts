import { ethers } from "hardhat";
import {
  ConfigurationError,
  handleError,
  PermissionError,
  validateAddress,
} from "./utils/error-handler";

async function main() {
  console.log("Setting token withdraw fee...");

  const [singer] = await ethers.getSigners();
  console.log("Caller:", singer.address);
  const balance = await ethers.provider.getBalance(singer.address);
  console.log("Balance:", ethers.formatEther(balance));

  // Read params from env (CLI args also supported; env takes precedence)
  const bridgeHubAddress = process.env.BRIDGE_HUB_ADDRESS || process.argv[2];
  const tokenAddress = process.env.BRIDGED_TOKEN_ADDRESS || process.argv[3];
  const fee = process.env.WITHDRAW_FEE || process.argv[4];

  if (!fee) {
    throw new ConfigurationError(
      "Missing withdraw fee",
      "Provide WITHDRAW_FEE via env or CLI argument",
      "Example: WITHDRAW_FEE=1000000000000000000 hardhat run scripts/set-withdraw-fee.ts"
    );
  }

  // Validate address format
  validateAddress(bridgeHubAddress, "BridgeHub contract");
  validateAddress(tokenAddress, "token contract");

  // Parameter preview
  console.log("\nParameter preview:");
  console.log(`  BridgeHub contract: ${bridgeHubAddress}`);
  console.log(`  Token address: ${tokenAddress}`);
  console.log(`  Withdraw fee: ${fee} (wei)`);

  try {
    // Connect to BridgeHub
    console.log("\n=== Connecting to BridgeHub ===");
    const BridgeHub = await ethers.getContractFactory("BridgeHub");
    const bridgeHub = BridgeHub.attach(bridgeHubAddress) as any;

    // Check ADMIN_ROLE
    console.log("Checking admin role...");
    const ADMIN_ROLE = await bridgeHub.ADMIN_ROLE();
    const hasAdminRole = await bridgeHub.hasRole(ADMIN_ROLE, singer.address);

    if (!hasAdminRole) {
      throw new ConfigurationError(
        "Insufficient permissions",
        `Account ${singer.address} does not have ADMIN_ROLE`,
        "Use an account with admin permissions, or ask the contract admin to grant the role"
      );
    }
    console.log("✅ Admin role verified");

    // Check current withdraw fee
    console.log("Checking current withdraw fee...");
    const currentFee = await bridgeHub.tokenWithdrawFee(tokenAddress);
    console.log(`Current withdraw fee: ${currentFee.toString()} (wei)`);

    if (BigInt(currentFee) === BigInt(fee)) {
      console.log("⚠️  Withdraw fee is already the target value, no change needed");
      return;
    }

    // Call setWithdrawFee
    console.log("\n=== Calling setWithdrawFee ===");
    const tx = await bridgeHub.setWithdrawFee(tokenAddress, BigInt(fee));

    console.log("Transaction submitted, waiting for confirmation...");
    console.log("Tx hash:", tx.hash);

    const receipt = await tx.wait();
    console.log("✅ Transaction confirmed!");
    console.log("Block number:", receipt?.blockNumber);
    console.log("Gas used:", receipt?.gasUsed?.toString());

    // Verify result
    console.log("\n=== Verifying result ===");
    const newFee = await bridgeHub.tokenWithdrawFee(tokenAddress);

    if (BigInt(newFee) == BigInt(fee)) {
      console.log("✅ Withdraw fee set successfully!");
      console.log(
        `Withdraw fee for token ${tokenAddress} is now ${ethers.formatEther(
          BigInt(fee)
        )}`
      );
    } else {
      throw new ConfigurationError(
        "Failed to set withdraw fee",
        `Expected: ${BigInt(fee)}, actual: ${newFee.toString()}`,
        "Check whether the transaction executed successfully"
      );
    }

    // Show fee change
    const feeChange = BigInt(fee) - BigInt(currentFee);
    if (feeChange > 0) {
      console.log(`📈 Withdraw fee increased by ${ethers.formatEther(feeChange)}`);
    } else if (feeChange < 0) {
      console.log(`📉 Withdraw fee decreased by ${ethers.formatEther(feeChange)}`);
    }
  } catch (error) {
    if (error instanceof ConfigurationError) {
      throw error;
    }
    throw new PermissionError(
      "Failed to set withdraw fee",
      error instanceof Error ? error.message : String(error),
      "Check network connectivity and contract state, then retry"
    );
  }
}

main()
  .then(() => {
    console.log("🎉 Withdraw fee setup complete!");
    process.exit(0);
  })
  .catch(handleError);
