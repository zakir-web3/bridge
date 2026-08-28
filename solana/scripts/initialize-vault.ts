import { PublicKey } from "@solana/web3.js";
import { getProgram, vaultAddresses } from "./utils";

function parseMintArg(): PublicKey {
  const mintArg = process.argv[2] ?? process.env.MINT;
  if (!mintArg) {
    console.error("usage: npm run initialize-vault -- <MINT_PUBKEY>");
    console.error("   or: MINT=<MINT_PUBKEY> npm run initialize-vault");
    process.exit(1);
  }

  try {
    return new PublicKey(mintArg);
  } catch {
    console.error("invalid mint pubkey:", mintArg);
    process.exit(1);
  }
}

async function main() {
  const mint = parseMintArg();
  const { provider, program } = getProgram();
  const { vaultStatePda, vaultAuthorityPda } = vaultAddresses(
    mint,
    program.programId
  );

  const existing = await provider.connection.getAccountInfo(vaultStatePda);
  if (existing) {
    const vaultState = await program.account.vaultState.fetch(vaultStatePda);
    console.log("vault already initialized for mint");
    console.log("  mint:", mint.toBase58());
    console.log("  vault_state:", vaultStatePda.toBase58());
    console.log("  vault_authority:", vaultAuthorityPda.toBase58());
    console.log("  stored mint:", vaultState.mint.toBase58());
    return;
  }

  const sig = await program.methods
    .initializeVault()
    .accounts({ mint })
    .rpc();

  const vaultState = await program.account.vaultState.fetch(vaultStatePda);
  console.log("initialize_vault succeeded");
  console.log("  signature:", sig);
  console.log("  mint:", mint.toBase58());
  console.log("  vault_state:", vaultStatePda.toBase58());
  console.log("  vault_authority:", vaultAuthorityPda.toBase58());
  console.log("  stored mint:", vaultState.mint.toBase58());
}

main().catch((err) => {
  console.error(err);
  process.exit(1);
});
