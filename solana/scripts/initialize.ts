import { getProgram, configPda } from "./utils";

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
    return;
  }

  const sig = await program.methods.initialize().accounts({}).rpc();

  const cfg = await program.account.bridgeConfig.fetch(config);
  console.log("initialize succeeded");
  console.log("  signature:", sig);
  console.log("  config:", config.toBase58());
  console.log("  admin:", cfg.admin.toBase58());
  console.log("  paused:", cfg.paused);
}

main().catch((err) => {
  console.error(err);
  process.exit(1);
});
