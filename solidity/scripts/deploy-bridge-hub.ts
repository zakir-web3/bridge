import { ethers } from "hardhat";
import {
  DeploymentError,
  handleError,
  validateAddressArray,
  validateArrayLength,
} from "./utils/error-handler";

async function main() {
  console.log("Deploying BridgeHub contract...");

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

  // Shared validators
  validateAddressArray(hotAddresses, "validator");
  validateAddressArray(coldAddresses, "cold wallet");
  validateArrayLength(hotAddresses, powers, "validator addresses", "powers");
  validateArrayLength(
    hotAddresses,
    coldAddresses,
    "validator addresses",
    "cold wallet addresses"
  );

  // Parameter preview
  console.log("\nParameter preview:");
  for (let i = 0; i < hotAddresses.length; i++) {
    console.log(
      `  Validator #${i + 1}: hot=${hotAddresses[i]}, cold=${
        coldAddresses[i]
      }, power=${powers[i]}`
    );
  }

  try {
    // Deploy BridgeHub implementation + initialize via ERC1967Proxy
    console.log("\n=== Deploying BridgeHub (UUPS upgradeable) ===");
    const BridgeHub = await ethers.getContractFactory("BridgeHub");
    const impl = await BridgeHub.deploy();
    await impl.waitForDeployment();
    const implAddress = await impl.getAddress();
    // console.log("BridgeHub implementation address:", implAddress);

    const initData = BridgeHub.interface.encodeFunctionData("initialize", [
      hotAddresses,
      coldAddresses,
      powers,
    ]);

    const ERC1967Proxy = await ethers.getContractFactory("ERC1967Proxy");
    const proxy = await ERC1967Proxy.deploy(implAddress, initData);
    await proxy.waitForDeployment();
    const bridgeHubAddress = await proxy.getAddress();
    console.log("BridgeHub proxy address:", bridgeHubAddress);
    console.log(`DEPLOYED_ADDRESS=${bridgeHubAddress}`);

    // Attach implementation ABI to the proxy (use any to avoid TS warnings)
    const bridgeHub = BridgeHub.attach(bridgeHubAddress) as any;
    const domainSeparator = await bridgeHub.domainSeparator();
    console.log("BridgeHub domain separator:", domainSeparator);
    console.log("✅ BridgeHub deployed successfully!");
  } catch (error) {
    throw new DeploymentError(
      "Failed to deploy BridgeHub contract",
      error instanceof Error ? error.message : String(error),
      "Check network connectivity and account balance, then retry"
    );
  }
}

main()
  .then(() => {
    console.log("🎉 Deployment complete!");
    process.exit(0);
  })
  .catch(handleError);
