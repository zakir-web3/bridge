import * as anchor from "@anchor-lang/core";
import { getProgram, configPda } from "./utils";

const CHAIN_ID = Number(process.env.SOLANA_CHAIN_ID ?? "900001");

async function main() {
  const { provider, program } = getProgram();
  const config = configPda(program.programId);

  const existing = await provider.connection.getAccountInfo(config);
  if (existing) {
    const cfg = await program.account.bridgeConfig.fetch(config);
    console.log("config already initialized");
    console.log("  config:", config.toBase58());
    console.log("  admin:", cfg.admin.toBase58());
    console.log("  paused:", cfg.paused);
    console.log("  chainId:", cfg.chainId.toNumber());
    return;
  }

  const sig = await program.methods
    .initialize(new anchor.BN(CHAIN_ID))
    .accounts({})
    .rpc();

  const cfg = await program.account.bridgeConfig.fetch(config);
  console.log("initialize succeeded");
  console.log("  signature:", sig);
  console.log("  config:", config.toBase58());
  console.log("  admin:", cfg.admin.toBase58());
  console.log("  paused:", cfg.paused);
  console.log("  chainId:", cfg.chainId.toNumber());
}

main().catch((err) => {
  console.error(err);
  process.exit(1);
});
