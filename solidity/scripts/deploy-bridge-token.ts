// @ts-nocheck
import { ethers } from "hardhat";
import { DeploymentError, handleError } from "./utils/error-handler";

async function main() {
  console.log("Deploying BridgeERC20 contract...");

  const [deployer] = await ethers.getSigners();
  console.log("Deployer:", deployer.address);

  const name = process.env.TOKEN_NAME || process.argv[2] || "USDT Token";
  const symbol = process.env.TOKEN_SYMBOL || process.argv[3] || "USDT";

  console.log(`Token config: name=${name}, symbol=${symbol}`);

  try {
    const BridgeERC20 = await ethers.getContractFactory("BridgeERC20");
    const token = await BridgeERC20.deploy(name, symbol);
    await token.waitForDeployment();
    const tokenAddress = await token.getAddress();
    console.log("✅ BridgeERC20 deployed at:", tokenAddress);
    console.log(`DEPLOYED_ADDRESS=${tokenAddress}`);

    const mintTo = process.env.MINT_TO;
    if (mintTo) {
      const mintAmount = BigInt(
        process.env.MINT_AMOUNT || "1000000000000000000"
      );
      const mintTx = await token.mint(mintTo, mintAmount);
      await mintTx.wait();
      console.log(`✅ Minted ${mintAmount.toString()} wei to ${mintTo}`);
    }
  } catch (error) {
    throw new DeploymentError(
      "Failed to deploy BridgeERC20 contract",
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
