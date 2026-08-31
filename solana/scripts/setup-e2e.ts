import * as anchor from "@anchor-lang/core";
import {
  createAssociatedTokenAccount,
  createMint,
  getAssociatedTokenAddressSync,
  mintTo,
} from "@solana/spl-token";
import { configPda, getProgram, vaultAddresses } from "./utils";

const TOKEN_DECIMALS = 6;
const MINT_AMOUNT = 1_000_000_000; // 1000 tokens with 6 decimals

async function main() {
  const { provider, program } = getProgram();
  const admin = provider.wallet as anchor.Wallet;
  const config = configPda(program.programId);

  const existingConfig = await provider.connection.getAccountInfo(config);
  if (!existingConfig) {
    const sig = await program.methods.initialize().accounts({}).rpc();
    console.log("initialized bridge config:", sig);
  }

  const mint = await createMint(
    provider.connection,
    admin.payer,
    admin.publicKey,
    null,
    TOKEN_DECIMALS
  );

  const { vaultStatePda } = vaultAddresses(mint, program.programId);
  const existingVault = await provider.connection.getAccountInfo(vaultStatePda);
  if (!existingVault) {
    const sig = await program.methods
      .initializeVault()
      .accounts({ mint })
      .rpc();
    console.log("initialized vault:", sig);
  }

  const userTokenAccount = getAssociatedTokenAddressSync(
    mint,
    admin.publicKey
  );
  await createAssociatedTokenAccount(
    provider.connection,
    admin.payer,
    mint,
    admin.publicKey
  );
  await mintTo(
    provider.connection,
    admin.payer,
    mint,
    userTokenAccount,
    admin.publicKey,
    MINT_AMOUNT
  );

  const mintBytes32 =
    "0x" + Buffer.from(mint.toBytes()).toString("hex");

  console.log(`PROGRAM_ID=${program.programId.toBase58()}`);
  console.log(`MINT=${mint.toBase58()}`);
  console.log(`SRC_TOKEN_BYTES32=${mintBytes32}`);
  console.log(`TOKEN_DECIMAL=${TOKEN_DECIMALS}`);
}

main().catch((err) => {
  console.error(err);
  process.exit(1);
});
