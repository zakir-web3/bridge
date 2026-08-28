import * as anchor from "@coral-xyz/anchor";
import { Program } from "@coral-xyz/anchor";
import { PublicKey } from "@solana/web3.js";
import { Bridge } from "../target/types/bridge";

export function getProgram(): {
  provider: anchor.AnchorProvider;
  program: Program<Bridge>;
} {
  const provider = anchor.AnchorProvider.env();
  anchor.setProvider(provider);
  const program = anchor.workspace.Bridge as Program<Bridge>;
  return { provider, program };
}

export function configPda(programId: PublicKey): PublicKey {
  const [pda] = PublicKey.findProgramAddressSync(
    [Buffer.from("config")],
    programId
  );
  return pda;
}

export function vaultAddresses(mint: PublicKey, programId: PublicKey) {
  const [vaultStatePda] = PublicKey.findProgramAddressSync(
    [Buffer.from("vault_state"), mint.toBuffer()],
    programId
  );
  const [vaultAuthorityPda] = PublicKey.findProgramAddressSync(
    [Buffer.from("vault"), mint.toBuffer()],
    programId
  );
  return { vaultStatePda, vaultAuthorityPda };
}
