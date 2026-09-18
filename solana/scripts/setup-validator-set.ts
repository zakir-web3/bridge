import * as anchor from "@anchor-lang/core";
import { ethers } from "ethers";
import { getProgram } from "./utils";

const DEFAULT_VALIDATOR = "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266";

async function main() {
  const validatorAddress = process.env.VALIDATOR ?? DEFAULT_VALIDATOR;
  const power = Number(process.env.VALIDATOR_POWER ?? "1");
  const epoch = Number(process.env.VALIDATOR_EPOCH ?? "1");

  const { program } = getProgram();
  const ethAddress = Array.from(ethers.getBytes(validatorAddress));

  const sig = await program.methods
    .setValidatorSet(new anchor.BN(epoch), [
      { ethAddress, power: new anchor.BN(power) },
    ])
    .accounts({})
    .rpc();

  console.log(`VALIDATOR_SET_TX=${sig}`);
  console.log(`VALIDATOR=${validatorAddress}`);
  console.log(`VALIDATOR_POWER=${power}`);
  console.log(`VALIDATOR_EPOCH=${epoch}`);
}

main().catch((err) => {
  console.error(err);
  process.exit(1);
});
