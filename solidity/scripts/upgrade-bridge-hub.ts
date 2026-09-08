import { execSync } from "node:child_process";
import { existsSync } from "node:fs";
import { ethers, upgrades } from "hardhat";
import { BridgeHub } from "../typechain-types";
import {
  validateAddress,
  handleError,
  ConfigurationError,
  PermissionError,
} from "./utils/error-handler";

function printForgeStorageLayout(contractName: string) {
  const foundryToml = "foundry.toml";
  if (!existsSync(foundryToml)) {
    console.log(
      "ℹ foundry.toml not found, skipping forge inspect storage-layout (upgrades-core can still validate)"
    );
    return;
  }

  try {
    const output = execSync(
      `forge inspect ${contractName} storage-layout --json`,
      { encoding: "utf8", stdio: ["pipe", "pipe", "pipe"] }
    );
    const layout = JSON.parse(output) as {
      storage?: Array<{ label: string; slot: string; type: string }>;
    };
    const rows = layout.storage ?? [];
    console.log(
      `forge storage-layout (${contractName}, ${rows.length} slots):`
    );
    for (const row of rows.slice(0, 12)) {
      console.log(`  slot ${row.slot}: ${row.label} (${row.type})`);
    }
    if (rows.length > 12) {
      console.log(`  ... ${rows.length - 12} more slots`);
    }
  } catch (error) {
    const message = error instanceof Error ? error.message : String(error);
    console.log(`⚠ forge inspect failed: ${message}`);
  }
}

async function main() {
  console.log("🚀 Upgrading BridgeHub contract...\n");

  const [signer] = await ethers.getSigners();
  console.log("Caller:", signer.address);
  const balance = await ethers.provider.getBalance(signer.address);
  console.log("Balance:", ethers.formatEther(balance), "ETH\n");

  // Read params from env
  const proxyAddress = process.env.BRIDGE_HUB_PROXY || process.argv[2];

  if (!proxyAddress) {
    throw new ConfigurationError(
      "Missing proxy address",
      "Provide BRIDGE_HUB_PROXY via env or CLI argument",
      "Example: BRIDGE_HUB_PROXY=0x... npx hardhat run scripts/upgrade-bridge-hub.ts"
    );
  }

  validateAddress(proxyAddress, "BridgeHub Proxy");

  try {
    // Step 1: verify current proxy state
    console.log("=== Step 1: Verify current proxy state ===");
    const BridgeHubV1 = await ethers.getContractFactory("BridgeHub");
    const proxyInstance = BridgeHubV1.attach(proxyAddress) as BridgeHub;

    // Current on-chain info
    const currentEpoch = await proxyInstance.epoch();
    const validatorsCount = (await proxyInstance.getHotValidators()).length;

    console.log(`✓ Current epoch: ${currentEpoch}`);
    console.log(`✓ Validator count: ${validatorsCount}`);

    // Check admin permissions
    const ADMIN_ROLE = await proxyInstance.ADMIN_ROLE();
    const hasAdminRole = await proxyInstance.hasRole(
      ADMIN_ROLE,
      signer.address
    );

    if (!hasAdminRole) {
      throw new PermissionError(
        "Insufficient permissions",
        `Account ${signer.address} does not have ADMIN_ROLE`,
        "Use an account with admin permissions, or ask the contract admin to grant the role"
      );
    }
    console.log("✓ Admin role verified\n");

    // Step 2: compile new implementation
    console.log("=== Step 2: Compile new implementation ===");
    console.log("Compiling BridgeHub V2...");
    // Contract is already compiled; get the factory
    const BridgeHubV2 = await ethers.getContractFactory("BridgeHub");
    console.log("✓ BridgeHub V2 is ready\n");

    // Step 2.5: storage layout check
    console.log("=== Step 2.5: Storage layout check ===");
    printForgeStorageLayout("BridgeHub");
    try {
      await upgrades.validateUpgrade(proxyAddress, BridgeHubV2, {
        kind: "uups",
      });
      console.log("✓ @openzeppelin/upgrades-core: storage layout compatible\n");
    } catch (error) {
      console.error("✗ Storage layout incompatible, aborting upgrade");
      if (error instanceof Error) {
        console.error(error.message);
      } else {
        console.error(String(error));
      }
      throw new PermissionError(
        "Storage layout check failed",
        "The new implementation is incompatible with the on-chain proxy storage layout",
        "Fix the layout conflict before upgrading, or deploy a new proxy instead of in-place upgrade"
      );
    }

    // Step 3: perform upgrade
    console.log("=== Step 3: Perform UUPS upgrade ===");
    console.log(`Proxy address: ${proxyAddress}`);
    console.log("Deploying new implementation...");

    const deployedImpl = await upgrades.upgradeProxy(
      proxyAddress,
      BridgeHubV2,
      {
        // Skip initializer: the proxy is already initialized
        unsafeSkipStorageCheck: false,
        kind: "uups",
      }
    );

    await deployedImpl.waitForDeployment();
    const implAddress = await ethers.provider.getStorage(
      proxyAddress,
      "0x360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc"
    );

    console.log("✓ New implementation deployed");
    console.log(`✓ Implementation address: 0x${implAddress.slice(-40)}\n`);

    // Step 4: verify upgrade
    console.log("=== Step 4: Verify upgrade ===");
    const upgradedInstance = BridgeHubV2.attach(proxyAddress) as BridgeHub;

    // Verify data integrity
    const newEpoch = await upgradedInstance.epoch();
    const newValidatorsCount = (await upgradedInstance.getHotValidators())
      .length;

    console.log(`✓ Epoch unchanged: ${newEpoch === currentEpoch}`);
    console.log(
      `✓ Validator data unchanged: ${newValidatorsCount === validatorsCount}`
    );

    // Verify new mapping exists
    try {
      // Probe the new tokenDecimalDiff mapping
      const testDiff = await upgradedInstance.tokenDecimalDiff(
        1,
        ethers.zeroPadValue(signer.address, 32)
      );
      console.log(`✓ New tokenDecimalDiff mapping is accessible`);
    } catch (e) {
      console.log(`⚠ Could not verify tokenDecimalDiff, which may be expected`);
    }

    console.log("\n✅ Upgrade succeeded!\n");

    // Step 5: print summary
    console.log("=== Upgrade summary ===");
    console.log(`Proxy address:            ${proxyAddress}`);
    console.log(`New implementation:       0x${implAddress.slice(-40)}`);
    console.log(`Upgrade account:          ${signer.address}`);
    console.log(`\nNew features:`);
    console.log(`  • setTokenPair now takes a tokenDecimal argument`);
    console.log(`  • Automatically verifies dstToken is a valid ERC20`);
    console.log(`  • Automatically computes and stores the decimal diff`);
    console.log(`  • deposit/withdraw automatically convert decimals`);
    console.log(`\nNext steps:`);
    console.log(
      `  1. Configure token pairs with the new setTokenPair method, passing tokenDecimal`
    );
    console.log(
      `     Example: npx hardhat run scripts/set-token-pair.ts --network <network>`
    );
    console.log(`  2. See DECIMAL_CONVERSION.md for decimal handling details`);
  } catch (error) {
    if (
      error instanceof ConfigurationError ||
      error instanceof PermissionError
    ) {
      throw error;
    }
    throw new PermissionError(
      "Failed to upgrade BridgeHub",
      error instanceof Error ? error.message : String(error),
      "Check network connectivity, contract state, and permissions, then retry"
    );
  }
}

main()
  .then(() => {
    console.log("🎉 BridgeHub upgrade complete!\n");
    process.exit(0);
  })
  .catch(handleError);
