import { ethers } from "hardhat";
import {
  validateAddress,
  handleError,
  PermissionError,
} from "./utils/error-handler";

async function main() {
  console.log("Setting ERC20 minter role...");

  const [signer] = await ethers.getSigners();
  console.log("Caller:", signer.address);

  const erc20Address = process.env.ERC20_ADDRESS || process.argv[2];
  const minerAddress = process.env.MINER_ADDRESS || process.argv[3];

  // Shared validators
  validateAddress(erc20Address!, "ERC20 contract");
  validateAddress(minerAddress!, "minter");

  // Only true is supported; the contract reverts on false
  const isActive = true;

  try {
    const erc20 = await ethers.getContractAt("BridgeERC20", erc20Address!);

    console.log(
      `Calling setMiner(miner=${minerAddress}, isActive=${isActive})...`
    );
    const tx = await (erc20 as any).setMiner(minerAddress!, isActive);
    console.log("Tx hash:", tx.hash);
    await tx.wait();

    console.log("✅ Done: ERC20 MINTER_ROLE updated");
  } catch (error) {
    throw new PermissionError(
      "Failed to set minter role",
      error instanceof Error ? error.message : String(error),
      "Check account permissions and contract state, then retry"
    );
  }
}

main()
  .then(() => {
    console.log("🎉 Operation complete!");
    process.exit(0);
  })
  .catch(handleError);
