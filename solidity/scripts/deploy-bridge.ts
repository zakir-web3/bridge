// @ts-nocheck
import { ethers } from "hardhat";
import {
  DeploymentError,
  handleError,
  validateAddressArray,
  validateArrayLength,
  validateNumber,
} from "./utils/error-handler";

async function main() {
  console.log("Deploying Bridge contract...");

  const [deployer] = await ethers.getSigners();
  console.log("Deployer:", deployer.address);
  const balance = await ethers.provider.getBalance(deployer.address);
  console.log("Balance:", ethers.formatEther(balance));

  // Read params from env (CLI args also supported; env takes precedence)
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

  const disputePeriodSeconds = parseInt(
    process.env.DISPUTE_PERIOD_SECONDS || process.argv[5] || "200",
    10
  );
  const blockDurationMillis = parseInt(
    process.env.BLOCK_DURATION_MILLIS || process.argv[6] || "750",
    10
  );
  const lockerThreshold = parseInt(
    process.env.LOCKER_THRESHOLD || process.argv[7] || "1",
    10
  );

  // Shared validators
  validateAddressArray(hotAddresses, "hot validator");
  validateAddressArray(coldAddresses, "cold validator");
  validateArrayLength(
    hotAddresses,
    coldAddresses,
    "hot validator addresses",
    "cold validator addresses"
  );
  validateArrayLength(powers, hotAddresses, "powers", "validator addresses");

  // Numeric params
  validateNumber(disputePeriodSeconds, "dispute period", 1);
  validateNumber(blockDurationMillis, "block duration", 100);
  validateNumber(lockerThreshold, "locker threshold", 1);

  // Parameter preview
  console.log("\nParameter preview:");
  for (let i = 0; i < hotAddresses.length; i++) {
    console.log(
      `  Validator #${i + 1}: hot=${hotAddresses[i]}, cold=${
        coldAddresses[i]
      }, power=${powers[i]}`
    );
  }
  console.log("  disputePeriodSeconds:", disputePeriodSeconds);
  console.log("  blockDurationMillis:", blockDurationMillis);
  console.log("  lockerThreshold:", lockerThreshold);

  try {
    const Bridge = await ethers.getContractFactory("Bridge");
    const bridge = await Bridge.deploy(
      hotAddresses,
      coldAddresses,
      powers,
      disputePeriodSeconds,
      blockDurationMillis,
      lockerThreshold
    );
    await bridge.waitForDeployment();
    const bridgeAddress = await bridge.getAddress();
    console.log("Bridge deployed at:", bridgeAddress);
    console.log(`DEPLOYED_ADDRESS=${bridgeAddress}`);
    const domainSeparator = await bridge.domainSeparator();
    console.log("Bridge domain separator:", domainSeparator);
    console.log("✅ Bridge deployed successfully!");
  } catch (error) {
    throw new DeploymentError(
      "Failed to deploy Bridge contract",
      error instanceof Error ? error.message : String(error),
      "Check network connectivity, account balance, and contract parameters, then retry"
    );
  }
}

main()
  .then(() => {
    console.log("🎉 Deployment complete!");
    process.exit(0);
  })
  .catch(handleError);
