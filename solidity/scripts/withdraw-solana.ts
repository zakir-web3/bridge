import { ethers } from "hardhat";
import { BridgeHub } from "../typechain-types";
import { sendTx } from "./utils/send-tx";

async function main() {
  console.log("Starting BridgeHub withdraw to Solana...");
  const [signer] = await ethers.getSigners();
  console.log(`Signer: ${signer.address}`);

  const bridgeHubAddress = process.env.BRIDGE_HUB_CONTRACT_ADDRESS;
  const tokenAddress = process.env.TOKEN_ADDRESS;
  const amountRaw = process.env.AMOUNT;
  const chainIdRaw = process.env.TOKEN_CHAIN_ID ?? process.env.CHAIN_ID;

  const destination = process.env.DESTINATION_BYTES32;

  if (!bridgeHubAddress) {
    throw new Error("Set BRIDGE_HUB_CONTRACT_ADDRESS");
  }
  if (!tokenAddress) {
    throw new Error("Set TOKEN_ADDRESS");
  }
  if (!amountRaw) {
    throw new Error("Set AMOUNT");
  }
  if (!chainIdRaw) {
    throw new Error("Set TOKEN_CHAIN_ID or CHAIN_ID");
  }
  if (!destination) {
    throw new Error("Set DESTINATION_BYTES32 (32-byte hex destination pubkey)");
  }

  const amount = BigInt(amountRaw);
  const chainId = BigInt(chainIdRaw);

  console.log("Config:");
  console.log(`  BridgeHub: ${bridgeHubAddress}`);
  console.log(`  Token: ${tokenAddress}`);
  console.log(`  Destination (bytes32): ${destination}`);
  console.log(`  Amount: ${amount.toString()}`);
  console.log(`  Source chain ID: ${chainId.toString()}`);

  const bridgeHub = (await ethers.getContractAt(
    "BridgeHub",
    bridgeHubAddress,
    signer
  )) as BridgeHub;
  const token = await ethers.getContractAt("IERC20", tokenAddress, signer);

  const withdrawFee = await bridgeHub.tokenWithdrawFee(tokenAddress);
  if (amount <= withdrawFee) {
    throw new Error(
      `Amount ${amount} must exceed withdraw fee ${withdrawFee}`
    );
  }

  const allowance = await token.allowance(signer.address, bridgeHubAddress);
  if (allowance < amount) {
    console.log("Approving BridgeHub...");
    await sendTx(signer, (nonce) =>
      token.approve(bridgeHubAddress, ethers.MaxUint256, { nonce })
    );
  }

  console.log("Submitting withdraw...");
  const receipt = await sendTx(signer, (nonce) =>
    bridgeHub.withdraw(destination, tokenAddress, amount, chainId, { nonce })
  );

  const withdrawNonce = await bridgeHub.withdrawNonce();
  console.log(`Withdraw tx mined in block ${receipt.blockNumber}`);
  console.log(`WITHDRAW_NONCE=${withdrawNonce.toString()}`);
  console.log(`DESTINATION_BYTES32=${destination}`);
}

main().catch((error) => {
  console.error(error);
  process.exit(1);
});
