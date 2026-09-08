import { ethers } from "hardhat";
import {
  validateAddress,
  handleError,
  ConfigurationError,
  PermissionError,
} from "./utils/error-handler";

function addressToBytes32(address: string) {
  return ethers.zeroPadValue(address, 32);
}

async function main() {
  console.log("Setting token pair...");

  const [singer] = await ethers.getSigners();
  console.log("Caller:", singer.address);
  const balance = await ethers.provider.getBalance(singer.address);
  console.log("Balance:", ethers.formatEther(balance));

  const bridgeHubAddress = process.env.BRIDGE_HUB_ADDRESS || process.argv[2];
  const chainId =
    process.env.TOKEN_CHAIN_ID || process.argv[3] || process.env.CHAIN_ID;
  const tokenAddress = process.env.TOKEN_ADDRESS || process.argv[4];
  const bridgedTokenAddress =
    process.env.BRIDGED_TOKEN_ADDRESS || process.argv[5];
  const tokenDecimal = process.env.TOKEN_DECIMAL || process.argv[6];
  const pairMode = process.env.PAIR_MODE || "evm";

  if (!chainId) {
    throw new ConfigurationError(
      "Missing chain ID",
      "Provide TOKEN_CHAIN_ID / CHAIN_ID via env or CLI argument",
      "Example: TOKEN_CHAIN_ID=56 hardhat run scripts/set-token-pair.ts"
    );
  }

  if (!tokenDecimal) {
    throw new ConfigurationError(
      "Missing token decimals",
      "Provide TOKEN_DECIMAL via env or CLI argument",
      "Example: TOKEN_DECIMAL=6 hardhat run scripts/set-token-pair.ts"
    );
  }

  validateAddress(bridgeHubAddress, "BridgeHub contract");
  validateAddress(bridgedTokenAddress, "bridged token");

  const chainIdNum = parseInt(chainId, 10);
  if (isNaN(chainIdNum) || chainIdNum <= 0) {
    throw new ConfigurationError(
      "Invalid chain ID",
      `Chain ID: ${chainId}`,
      "Please provide a chain ID greater than 0"
    );
  }

  const tokenDecimalNum = parseInt(tokenDecimal, 10);
  if (isNaN(tokenDecimalNum) || tokenDecimalNum < 0 || tokenDecimalNum > 18) {
    throw new ConfigurationError(
      "Invalid token decimals",
      `Decimals: ${tokenDecimal}`,
      "Please provide a decimals value between 0 and 18"
    );
  }

  let srcTokenBytes32: string;
  if (pairMode === "solana") {
    const srcTokenHex = process.env.SRC_TOKEN_BYTES32 || tokenAddress;
    if (
      !srcTokenHex ||
      !srcTokenHex.startsWith("0x") ||
      srcTokenHex.length !== 66
    ) {
      throw new ConfigurationError(
        "Invalid Solana mint bytes32",
        `SRC_TOKEN_BYTES32: ${srcTokenHex}`,
        "Provide 32-byte hex, e.g. 0x..."
      );
    }
    srcTokenBytes32 = srcTokenHex;
    console.log("\nParameter preview (Solana src):");
    console.log(`  Source SPL mint (bytes32): ${srcTokenBytes32}`);
  } else {
    validateAddress(tokenAddress, "source token");
    srcTokenBytes32 = addressToBytes32(tokenAddress);
    console.log("\nParameter preview (EVM src):");
    console.log(`  Source token address: ${tokenAddress}`);
  }

  console.log(`  BridgeHub contract: ${bridgeHubAddress}`);
  console.log(`  Chain ID: ${chainIdNum}`);
  console.log(`  Source token decimals: ${tokenDecimalNum}`);
  console.log(`  Bridged token address: ${bridgedTokenAddress}`);

  try {
    console.log("\n=== Connecting to BridgeHub ===");
    const BridgeHub = await ethers.getContractFactory("BridgeHub");
    const bridgeHub = BridgeHub.attach(bridgeHubAddress) as any;

    console.log("Checking admin role...");
    const ADMIN_ROLE = await bridgeHub.ADMIN_ROLE();
    const hasAdminRole = await bridgeHub.hasRole(ADMIN_ROLE, singer.address);

    if (!hasAdminRole) {
      throw new ConfigurationError(
        "Insufficient permissions",
        `Account ${singer.address} does not have ADMIN_ROLE`,
        "Use an account with admin permissions"
      );
    }
    console.log("✅ Admin role verified");

    const currentBridgedToken = await bridgeHub.tokenPair(
      chainIdNum,
      srcTokenBytes32
    );

    if (currentBridgedToken !== ethers.ZeroHash) {
      throw new Error(
        `Source token ${srcTokenBytes32} on chain ${chainIdNum} is already mapped to ${currentBridgedToken}`
      );
    }

    const bridgedTokenBytes32 = addressToBytes32(bridgedTokenAddress);
    const reverseToken = await bridgeHub.tokenPair(chainIdNum, bridgedTokenBytes32);
    if (reverseToken !== ethers.ZeroHash) {
      throw new ConfigurationError(
        "Bridged token withdraw mapping already exists",
        `Bridged token on chain ${chainIdNum} is already mapped to ${reverseToken}`,
        "Check whether this token pair is already configured"
      );
    }

    console.log("\n=== Calling setTokenPair ===");
    const tx = await bridgeHub.setTokenPair(
      chainIdNum,
      srcTokenBytes32,
      tokenDecimalNum,
      bridgedTokenAddress
    );

    console.log("Transaction submitted, waiting for confirmation...");
    console.log("Tx hash:", tx.hash);

    const receipt = await tx.wait();
    console.log("✅ Transaction confirmed!");
    console.log("Block number:", receipt?.blockNumber);

    const remoteToken = await bridgeHub.tokenPair(
      chainIdNum,
      bridgedTokenBytes32
    );
    if (remoteToken.toLowerCase() !== srcTokenBytes32.toLowerCase()) {
      throw new ConfigurationError(
        "Failed to set withdraw token pair",
        `Expected: ${srcTokenBytes32}, actual: ${remoteToken}`,
        "Check whether the transaction executed successfully"
      );
    }
    console.log("✅ Withdraw mapping configured automatically");

    const newBridgedToken = await bridgeHub.tokenPair(
      chainIdNum,
      srcTokenBytes32
    );
    if (
      newBridgedToken.toLowerCase() !==
      ethers.zeroPadValue(bridgedTokenAddress, 32).toLowerCase()
    ) {
      throw new ConfigurationError(
        "Failed to set token pair",
        `Expected: ${bridgedTokenAddress}, actual: ${newBridgedToken}`,
        "Check whether the transaction executed successfully"
      );
    }
    console.log("✅ Token pair set successfully!");
  } catch (error) {
    if (error instanceof ConfigurationError) {
      throw error;
    }
    throw new PermissionError(
      "Failed to set token pair",
      error instanceof Error ? error.message : String(error),
      "Check network connectivity and contract state, then retry"
    );
  }
}

main()
  .then(() => {
    console.log("🎉 Token pair setup complete!");
    process.exit(0);
  })
  .catch(handleError);
