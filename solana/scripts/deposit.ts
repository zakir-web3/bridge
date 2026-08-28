import * as anchor from "@coral-xyz/anchor";
import {
  createAssociatedTokenAccountIdempotent,
  getAssociatedTokenAddressSync,
} from "@solana/spl-token";
import { PublicKey } from "@solana/web3.js";
import { getProgram, vaultAddresses } from "./utils";

function usage(): never {
  console.error(
    "usage: npm run deposit -- <MINT_PUBKEY> <AMOUNT> <EVM_DESTINATION>"
  );
  console.error(
    "   or: MINT=<MINT_PUBKEY> AMOUNT=<AMOUNT> DESTINATION=<EVM_ADDRESS> npm run deposit"
  );
  console.error("  AMOUNT is in SPL smallest units (e.g. 1000000 = 1 token with 6 decimals)");
  console.error("  DESTINATION is a 20-byte EVM address (with or without 0x prefix)");
  process.exit(1);
}

function parseMint(): PublicKey {
  const mintArg = process.argv[2] ?? process.env.MINT;
  if (!mintArg) {
    usage();
  }

  try {
    return new PublicKey(mintArg);
  } catch {
    console.error("invalid mint pubkey:", mintArg);
    process.exit(1);
  }
}

function parseAmount(): anchor.BN {
  const amountArg = process.argv[3] ?? process.env.AMOUNT;
  if (!amountArg) {
    usage();
  }

  if (!/^\d+$/.test(amountArg)) {
    console.error("invalid amount:", amountArg);
    process.exit(1);
  }

  const amount = new anchor.BN(amountArg);
  if (amount.lte(new anchor.BN(0))) {
    console.error("amount must be greater than zero");
    process.exit(1);
  }

  return amount;
}

function parseDestination(): number[] {
  const destinationArg = process.argv[4] ?? process.env.DESTINATION;
  if (!destinationArg) {
    usage();
  }

  const hex = destinationArg.startsWith("0x")
    ? destinationArg.slice(2)
    : destinationArg;

  if (!/^[0-9a-fA-F]{40}$/.test(hex)) {
    console.error("invalid EVM destination:", destinationArg);
    console.error("expected 20-byte hex address (40 hex chars)");
    process.exit(1);
  }

  return Array.from(Buffer.from(hex, "hex"));
}

async function main() {
  const mint = parseMint();
  const amount = parseAmount();
  const destination = parseDestination();

  const { provider, program } = getProgram();
  const user = (provider.wallet as anchor.Wallet).publicKey;
  const { vaultAuthorityPda } = vaultAddresses(mint, program.programId);

  const userTokenAccount = getAssociatedTokenAddressSync(mint, user);
  const vaultTokenAccount = getAssociatedTokenAddressSync(
    mint,
    vaultAuthorityPda,
    true
  );

  const userAtaInfo = await provider.connection.getAccountInfo(userTokenAccount);
  if (!userAtaInfo) {
    await createAssociatedTokenAccountIdempotent(
      provider.connection,
      (provider.wallet as anchor.Wallet).payer,
      mint,
      user
    );
    console.log("created user ATA:", userTokenAccount.toBase58());
  }

  const userBalance = await provider.connection.getTokenAccountBalance(
    userTokenAccount
  );
  if (BigInt(userBalance.value.amount) < BigInt(amount.toString())) {
    console.error("insufficient SPL balance");
    console.error("  user ATA:", userTokenAccount.toBase58());
    console.error("  balance:", userBalance.value.amount);
    console.error("  required:", amount.toString());
    process.exit(1);
  }

  const sig = await program.methods
    .deposit(destination, amount)
    .accounts({ mint })
    .rpc();

  const vaultBalance = await provider.connection.getTokenAccountBalance(
    vaultTokenAccount
  );

  console.log("deposit succeeded");
  console.log("  signature:", sig);
  console.log("  user:", user.toBase58());
  console.log("  mint:", mint.toBase58());
  console.log("  amount:", amount.toString());
  console.log(
    "  destination:",
    "0x" + Buffer.from(destination).toString("hex")
  );
  console.log("  vault_token_account:", vaultTokenAccount.toBase58());
  console.log("  vault_balance:", vaultBalance.value.amount);
}

main().catch((err) => {
  console.error(err);
  process.exit(1);
});
