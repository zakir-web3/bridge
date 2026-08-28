import * as anchor from "@coral-xyz/anchor";
import { Program } from "@coral-xyz/anchor";
import {
  createAssociatedTokenAccount,
  createMint,
  getAssociatedTokenAddressSync,
  mintTo,
} from "@solana/spl-token";
import { PublicKey } from "@solana/web3.js";
import { expect } from "chai";
import { Bridge } from "../target/types/bridge";

describe("bridge", () => {
  const provider = anchor.AnchorProvider.env();
  anchor.setProvider(provider);

  const program = anchor.workspace.Bridge as Program<Bridge>;
  const admin = provider.wallet as anchor.Wallet;

  const [configPda] = PublicKey.findProgramAddressSync(
    [Buffer.from("config")],
    program.programId
  );

  let mint: PublicKey;
  let vaultStatePda: PublicKey;
  let vaultAuthorityPda: PublicKey;
  let vaultTokenAccount: PublicKey;
  let userTokenAccount: PublicKey;

  const destination = Buffer.alloc(20);
  destination.fill(0xab);

  before(async () => {
    mint = await createMint(
      provider.connection,
      admin.payer,
      admin.publicKey,
      null,
      6
    );

    [vaultStatePda] = PublicKey.findProgramAddressSync(
      [Buffer.from("vault_state"), mint.toBuffer()],
      program.programId
    );
    [vaultAuthorityPda] = PublicKey.findProgramAddressSync(
      [Buffer.from("vault"), mint.toBuffer()],
      program.programId
    );
    vaultTokenAccount = getAssociatedTokenAddressSync(
      mint,
      vaultAuthorityPda,
      true
    );
    userTokenAccount = getAssociatedTokenAddressSync(mint, admin.publicKey);

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
      1_000_000
    );
  });

  it("initializes config", async () => {
    await program.methods
      .initialize()
      .accounts({})
      .rpc();

    const config = await program.account.bridgeConfig.fetch(configPda);
    expect(config.admin.toBase58()).to.equal(admin.publicKey.toBase58());
    expect(config.paused).to.equal(false);
  });

  it("initializes vault for mint", async () => {
    await program.methods
      .initializeVault()
      .accounts({
        mint,
      })
      .rpc();

    const vaultState = await program.account.vaultState.fetch(vaultStatePda);
    expect(vaultState.mint.toBase58()).to.equal(mint.toBase58());
  });

  it("deposits SPL into vault", async () => {
    const amount = new anchor.BN(100_000);

    await program.methods
      .deposit(Array.from(destination), amount)
      .accounts({
        mint,
      })
      .rpc();

    const vaultBalance = await provider.connection.getTokenAccountBalance(
      vaultTokenAccount
    );
    expect(vaultBalance.value.amount).to.equal("100000");
  });
});
