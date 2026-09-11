import * as anchor from "@anchor-lang/core";
import { Program } from "@anchor-lang/core";
import {
  createAssociatedTokenAccount,
  createMint,
  getAssociatedTokenAddressSync,
  mintTo,
  getAccount,
} from "@solana/spl-token";
import { Keypair, PublicKey, SystemProgram } from "@solana/web3.js";
import { expect } from "chai";
import { ethers } from "ethers";
import { Bridge } from "../target/types/bridge";

// ---------------------------------------------------------------------------
// Helpers
// ---------------------------------------------------------------------------

const CHAIN_ID = 900001;
const NONCES_PER_PAGE = 8192;

interface ValidatorInfo {
  wallet: ethers.Wallet;
  ethAddress: number[]; // 20 bytes
  power: anchor.BN;
}

/** Create a test validator from a deterministic private key. */
function makeValidator(seed: number, power: number): ValidatorInfo {
  const privHex = ethers.zeroPadValue(ethers.toBeHex(seed), 32);
  const wallet = new ethers.Wallet(privHex);
  const addrBytes = Array.from(ethers.getBytes(wallet.address));
  return { wallet, ethAddress: addrBytes, power: new anchor.BN(power) };
}

/** Build EIP-712 domain matching the on-chain BridgeConfig. */
function bridgeDomain(chainId: number, verifyingContract: string) {
  return {
    name: "Bridge",
    version: "1",
    chainId,
    verifyingContract,
  };
}

const WITHDRAW_TYPES = {
  Withdraw: [
    { name: "user", type: "address" },
    { name: "destination", type: "bytes32" },
    { name: "token", type: "bytes32" },
    { name: "amount", type: "uint256" },
    { name: "chainId", type: "uint256" },
    { name: "nonce", type: "uint64" },
  ],
};

/** Derive the verifyingContract pseudo-address from a Solana program ID. */
function deriveVerifyingContract(programId: PublicKey): string {
  const hash = ethers.keccak256(programId.toBytes());
  return "0x" + hash.slice(-40); // last 20 bytes
}

/** Pad a 20-byte EVM address into bytes32 (left-pad with 12 zeros). */
function addressToBytes32(addr: string): string {
  return ethers.zeroPadValue(addr, 32);
}

/** Sign an EIP-712 Withdraw message, returning { sig, recovery_id } for Anchor. */
async function signWithdraw(
  wallet: ethers.Wallet,
  domain: ReturnType<typeof bridgeDomain>,
  user: string,
  destination: Uint8Array | Buffer,
  token: Uint8Array | Buffer,
  amount: bigint,
  chainId: number,
  nonce: number
): Promise<{ sig: number[]; recoveryId: number }> {
  const value = {
    user,
    destination: ethers.hexlify(destination),
    token: ethers.hexlify(token),
    amount: amount.toString(),
    chainId,
    nonce,
  };

  const digest = ethers.TypedDataEncoder.hash(domain, WITHDRAW_TYPES, value);
  const sigRaw = wallet.signingKey.sign(digest);

  const r = ethers.getBytes(sigRaw.r);
  const s = ethers.getBytes(sigRaw.s);
  const sig = Array.from(Buffer.concat([Buffer.from(r), Buffer.from(s)]));
  const recoveryId = sigRaw.v - 27;

  return { sig, recoveryId };
}

/** Find the NoncePage PDA for a given nonce. */
function findNoncePagePda(
  nonce: number,
  programId: PublicKey
): [PublicKey, number] {
  const page = Math.floor(nonce / NONCES_PER_PAGE);
  const pageBytes = Buffer.alloc(8);
  pageBytes.writeBigUInt64LE(BigInt(page));
  return PublicKey.findProgramAddressSync(
    [Buffer.from("nonce_page"), pageBytes],
    programId
  );
}

// ---------------------------------------------------------------------------
// Tests
// ---------------------------------------------------------------------------

describe("bridge", () => {
  const provider = anchor.AnchorProvider.env();
  anchor.setProvider(provider);

  const program = anchor.workspace.Bridge as Program<Bridge>;
  const admin = provider.wallet as anchor.Wallet;

  const [configPda] = PublicKey.findProgramAddressSync(
    [Buffer.from("config")],
    program.programId
  );

  const [validatorSetPda] = PublicKey.findProgramAddressSync(
    [Buffer.from("validator_set")],
    program.programId
  );

  let mint: PublicKey;
  let vaultStatePda: PublicKey;
  let vaultAuthorityPda: PublicKey;
  let vaultTokenAccount: PublicKey;
  let userTokenAccount: PublicKey;

  const destination = Buffer.alloc(20);
  destination.fill(0xab);

  // Withdraw-specific setup
  const verifyingContract = deriveVerifyingContract(program.programId);
  const domain = bridgeDomain(CHAIN_ID, verifyingContract);

  // Powers: v1+v2 = 200, total = 250. Quorum: 3*200 = 600 > 2*250 = 500 ✓
  // v1 alone: 3*100 = 300 > 500 ✗ (for insufficient quorum test)
  const v1 = makeValidator(101, 100);
  const v2 = makeValidator(102, 100);
  const v3 = makeValidator(103, 50);
  const validators = [v1, v2, v3]; // total_power = 250

  let withdrawRecipient: Keypair;
  let withdrawRecipientAta: PublicKey;

  const DEPOSIT_AMOUNT = 1_000_000; // 1 token at 6 decimals
  const WITHDRAW_AMOUNT = 100_000n; // 0.1 token

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
      DEPOSIT_AMOUNT
    );

    withdrawRecipient = Keypair.generate();
  });

  // ---- Deposit flow (existing) ----

  it("initializes config with chain_id", async () => {
    await program.methods
      .initialize(new anchor.BN(CHAIN_ID))
      .accounts({})
      .rpc();

    const config = await program.account.bridgeConfig.fetch(configPda);
    expect(config.admin.toBase58()).to.equal(admin.publicKey.toBase58());
    expect(config.paused).to.equal(false);
    expect(config.withdrawPaused).to.equal(false);
    expect(config.chainId.toNumber()).to.equal(CHAIN_ID);
    expect(config.verifyingContract.length).to.equal(20);
    expect(config.domainSeparator.length).to.equal(32);
  });

  it("initializes vault for mint", async () => {
    await program.methods.initializeVault().accounts({ mint }).rpc();

    const vaultState = await program.account.vaultState.fetch(vaultStatePda);
    expect(vaultState.mint.toBase58()).to.equal(mint.toBase58());
  });

  it("deposits SPL into vault", async () => {
    const amount = new anchor.BN(DEPOSIT_AMOUNT);

    await program.methods
      .deposit(Array.from(destination), amount)
      .accounts({ mint })
      .rpc();

    const vaultBalance = await provider.connection.getTokenAccountBalance(
      vaultTokenAccount
    );
    expect(vaultBalance.value.amount).to.equal(DEPOSIT_AMOUNT.toString());
  });

  // ---- Admin: validator set ----

  it("sets validator set (epoch 1)", async () => {
    await program.methods
      .setValidatorSet(
        new anchor.BN(1),
        validators.map((v) => ({
          ethAddress: v.ethAddress,
          power: v.power,
        }))
      )
      .accounts({})
      .rpc();

    const vs = await program.account.validatorSet.fetch(validatorSetPda);
    expect(vs.epoch.toNumber()).to.equal(1);
    expect(vs.totalPower.toNumber()).to.equal(250);
    expect(vs.validators.length).to.equal(3);
  });

  it("rejects epoch that is not strictly increasing", async () => {
    try {
      await program.methods
        .setValidatorSet(
          new anchor.BN(1),
          validators.map((v) => ({
            ethAddress: v.ethAddress,
            power: v.power,
          }))
        )
        .accounts({})
        .rpc();
      expect.fail("should have thrown");
    } catch (err: any) {
      expect(err.toString()).to.include("EpochNotIncreasing");
    }
  });

  // ---- Admin: set_withdraw_paused ----

  it("pauses and unpauses withdrawals", async () => {
    await program.methods.setWithdrawPaused(true).accounts({}).rpc();
    let config = await program.account.bridgeConfig.fetch(configPda);
    expect(config.withdrawPaused).to.equal(true);

    await program.methods.setWithdrawPaused(false).accounts({}).rpc();
    config = await program.account.bridgeConfig.fetch(configPda);
    expect(config.withdrawPaused).to.equal(false);
  });

  // ---- Withdraw: happy path ----

  it("withdraws with valid 2-of-3 quorum", async () => {
    const nonce = 0;
    const destinationBytes = Array.from(withdrawRecipient.publicKey.toBytes());
    const tokenBytes32 = addressToBytes32(
      ethers.hexlify(mint.toBytes())
    );
    const tokenBuf = ethers.getBytes(tokenBytes32);
    const destBytes32 = ethers.hexlify(
      withdrawRecipient.publicKey.toBytes()
    );
    const destBuf = ethers.getBytes(destBytes32);
    const user = v1.wallet.address;

    const sig1 = await signWithdraw(
      v1.wallet,
      domain,
      user,
      destBuf,
      tokenBuf,
      WITHDRAW_AMOUNT,
      CHAIN_ID,
      nonce
    );
    const sig2 = await signWithdraw(
      v2.wallet,
      domain,
      user,
      destBuf,
      tokenBuf,
      WITHDRAW_AMOUNT,
      CHAIN_ID,
      nonce
    );

    const [noncePagePda] = findNoncePagePda(nonce, program.programId);

    await program.methods
      .withdraw(
        Array.from(ethers.getBytes(user)),
        destinationBytes,
        new anchor.BN(WITHDRAW_AMOUNT.toString()),
        new anchor.BN(nonce),
        [
          { sig: sig1.sig, recoveryId: sig1.recoveryId },
          { sig: sig2.sig, recoveryId: sig2.recoveryId },
        ]
      )
      .accounts({
        mint,
        destinationOwner: withdrawRecipient.publicKey,
        noncePage: noncePagePda,
      })
      .rpc();

    // Verify tokens transferred
    const destAta = getAssociatedTokenAddressSync(
      mint,
      withdrawRecipient.publicKey
    );
    const balance = await provider.connection.getTokenAccountBalance(destAta);
    expect(balance.value.amount).to.equal(WITHDRAW_AMOUNT.toString());

    // Verify vault balance decreased
    const vaultBalance = await provider.connection.getTokenAccountBalance(
      vaultTokenAccount
    );
    expect(vaultBalance.value.amount).to.equal(
      (BigInt(DEPOSIT_AMOUNT) - WITHDRAW_AMOUNT).toString()
    );
  });

  // ---- Withdraw: nonce replay ----

  it("rejects nonce replay", async () => {
    const nonce = 0; // same as above
    const destBuf = ethers.getBytes(
      ethers.hexlify(withdrawRecipient.publicKey.toBytes())
    );
    const tokenBuf = ethers.getBytes(
      addressToBytes32(ethers.hexlify(mint.toBytes()))
    );
    const user = v1.wallet.address;

    const sig1 = await signWithdraw(
      v1.wallet,
      domain,
      user,
      destBuf,
      tokenBuf,
      WITHDRAW_AMOUNT,
      CHAIN_ID,
      nonce
    );
    const sig2 = await signWithdraw(
      v2.wallet,
      domain,
      user,
      destBuf,
      tokenBuf,
      WITHDRAW_AMOUNT,
      CHAIN_ID,
      nonce
    );

    const [noncePagePda] = findNoncePagePda(nonce, program.programId);

    try {
      await program.methods
        .withdraw(
          Array.from(ethers.getBytes(user)),
          Array.from(withdrawRecipient.publicKey.toBytes()),
          new anchor.BN(WITHDRAW_AMOUNT.toString()),
          new anchor.BN(nonce),
          [
            { sig: sig1.sig, recoveryId: sig1.recoveryId },
            { sig: sig2.sig, recoveryId: sig2.recoveryId },
          ]
        )
        .accounts({
          mint,
          destinationOwner: withdrawRecipient.publicKey,
          noncePage: noncePagePda,
        })
        .rpc();
      expect.fail("should have thrown");
    } catch (err: any) {
      expect(err.toString()).to.include("NonceAlreadyUsed");
    }
  });

  // ---- Withdraw: cross page boundary ----

  it("handles nonce on a different page", async () => {
    const nonce = NONCES_PER_PAGE; // first nonce of page 1
    const recipient2 = Keypair.generate();
    const destBuf = ethers.getBytes(
      ethers.hexlify(recipient2.publicKey.toBytes())
    );
    const tokenBuf = ethers.getBytes(
      addressToBytes32(ethers.hexlify(mint.toBytes()))
    );
    const user = v1.wallet.address;

    const sig1 = await signWithdraw(
      v1.wallet,
      domain,
      user,
      destBuf,
      tokenBuf,
      WITHDRAW_AMOUNT,
      CHAIN_ID,
      nonce
    );
    const sig2 = await signWithdraw(
      v2.wallet,
      domain,
      user,
      destBuf,
      tokenBuf,
      WITHDRAW_AMOUNT,
      CHAIN_ID,
      nonce
    );

    const [noncePagePda] = findNoncePagePda(nonce, program.programId);

    await program.methods
      .withdraw(
        Array.from(ethers.getBytes(user)),
        Array.from(recipient2.publicKey.toBytes()),
        new anchor.BN(WITHDRAW_AMOUNT.toString()),
        new anchor.BN(nonce),
        [
          { sig: sig1.sig, recoveryId: sig1.recoveryId },
          { sig: sig2.sig, recoveryId: sig2.recoveryId },
        ]
      )
      .accounts({
        mint,
        destinationOwner: recipient2.publicKey,
        noncePage: noncePagePda,
      })
      .rpc();

    const destAta = getAssociatedTokenAddressSync(
      mint,
      recipient2.publicKey
    );
    const balance = await provider.connection.getTokenAccountBalance(destAta);
    expect(balance.value.amount).to.equal(WITHDRAW_AMOUNT.toString());
  });

  // ---- Withdraw: insufficient quorum ----

  it("rejects withdraw with insufficient quorum (1-of-3)", async () => {
    const nonce = 1;
    const recipient3 = Keypair.generate();
    const destBuf = ethers.getBytes(
      ethers.hexlify(recipient3.publicKey.toBytes())
    );
    const tokenBuf = ethers.getBytes(
      addressToBytes32(ethers.hexlify(mint.toBytes()))
    );
    const user = v1.wallet.address;

    const sig1 = await signWithdraw(
      v1.wallet,
      domain,
      user,
      destBuf,
      tokenBuf,
      WITHDRAW_AMOUNT,
      CHAIN_ID,
      nonce
    );

    const [noncePagePda] = findNoncePagePda(nonce, program.programId);

    try {
      await program.methods
        .withdraw(
          Array.from(ethers.getBytes(user)),
          Array.from(recipient3.publicKey.toBytes()),
          new anchor.BN(WITHDRAW_AMOUNT.toString()),
          new anchor.BN(nonce),
          [{ sig: sig1.sig, recoveryId: sig1.recoveryId }]
        )
        .accounts({
          mint,
          destinationOwner: recipient3.publicKey,
          noncePage: noncePagePda,
        })
        .rpc();
      expect.fail("should have thrown");
    } catch (err: any) {
      expect(err.toString()).to.include("InsufficientQuorum");
    }
  });

  // ---- Withdraw: duplicate signature ----

  it("rejects duplicate validator signatures", async () => {
    const nonce = 2;
    const recipient4 = Keypair.generate();
    const destBuf = ethers.getBytes(
      ethers.hexlify(recipient4.publicKey.toBytes())
    );
    const tokenBuf = ethers.getBytes(
      addressToBytes32(ethers.hexlify(mint.toBytes()))
    );
    const user = v1.wallet.address;

    const sig1 = await signWithdraw(
      v1.wallet,
      domain,
      user,
      destBuf,
      tokenBuf,
      WITHDRAW_AMOUNT,
      CHAIN_ID,
      nonce
    );

    const [noncePagePda] = findNoncePagePda(nonce, program.programId);

    try {
      await program.methods
        .withdraw(
          Array.from(ethers.getBytes(user)),
          Array.from(recipient4.publicKey.toBytes()),
          new anchor.BN(WITHDRAW_AMOUNT.toString()),
          new anchor.BN(nonce),
          [
            { sig: sig1.sig, recoveryId: sig1.recoveryId },
            { sig: sig1.sig, recoveryId: sig1.recoveryId },
          ]
        )
        .accounts({
          mint,
          destinationOwner: recipient4.publicKey,
          noncePage: noncePagePda,
        })
        .rpc();
      expect.fail("should have thrown");
    } catch (err: any) {
      expect(err.toString()).to.include("DuplicateSignature");
    }
  });

  // ---- Withdraw: unknown validator ----

  it("rejects signature from unknown validator", async () => {
    const nonce = 3;
    const unknownV = makeValidator(999, 100);
    const recipient5 = Keypair.generate();
    const destBuf = ethers.getBytes(
      ethers.hexlify(recipient5.publicKey.toBytes())
    );
    const tokenBuf = ethers.getBytes(
      addressToBytes32(ethers.hexlify(mint.toBytes()))
    );
    const user = unknownV.wallet.address;

    const sig1 = await signWithdraw(
      unknownV.wallet,
      domain,
      user,
      destBuf,
      tokenBuf,
      WITHDRAW_AMOUNT,
      CHAIN_ID,
      nonce
    );
    const sig2 = await signWithdraw(
      v1.wallet,
      domain,
      user,
      destBuf,
      tokenBuf,
      WITHDRAW_AMOUNT,
      CHAIN_ID,
      nonce
    );

    const [noncePagePda] = findNoncePagePda(nonce, program.programId);

    try {
      await program.methods
        .withdraw(
          Array.from(ethers.getBytes(user)),
          Array.from(recipient5.publicKey.toBytes()),
          new anchor.BN(WITHDRAW_AMOUNT.toString()),
          new anchor.BN(nonce),
          [
            { sig: sig1.sig, recoveryId: sig1.recoveryId },
            { sig: sig2.sig, recoveryId: sig2.recoveryId },
          ]
        )
        .accounts({
          mint,
          destinationOwner: recipient5.publicKey,
          noncePage: noncePagePda,
        })
        .rpc();
      expect.fail("should have thrown");
    } catch (err: any) {
      expect(err.toString()).to.include("UnknownValidator");
    }
  });

  // ---- Withdraw: destination mismatch ----

  it("rejects when destination does not match destination_owner", async () => {
    const nonce = 4;
    const wrongRecipient = Keypair.generate();
    const actualRecipient = Keypair.generate();
    const destBuf = ethers.getBytes(
      ethers.hexlify(wrongRecipient.publicKey.toBytes())
    );
    const tokenBuf = ethers.getBytes(
      addressToBytes32(ethers.hexlify(mint.toBytes()))
    );
    const user = v1.wallet.address;

    const sig1 = await signWithdraw(
      v1.wallet,
      domain,
      user,
      destBuf,
      tokenBuf,
      WITHDRAW_AMOUNT,
      CHAIN_ID,
      nonce
    );
    const sig2 = await signWithdraw(
      v2.wallet,
      domain,
      user,
      destBuf,
      tokenBuf,
      WITHDRAW_AMOUNT,
      CHAIN_ID,
      nonce
    );

    const [noncePagePda] = findNoncePagePda(nonce, program.programId);

    try {
      await program.methods
        .withdraw(
          Array.from(ethers.getBytes(user)),
          // destination bytes point to wrongRecipient
          Array.from(wrongRecipient.publicKey.toBytes()),
          new anchor.BN(WITHDRAW_AMOUNT.toString()),
          new anchor.BN(nonce),
          [
            { sig: sig1.sig, recoveryId: sig1.recoveryId },
            { sig: sig2.sig, recoveryId: sig2.recoveryId },
          ]
        )
        .accounts({
          mint,
          // but account is actualRecipient
          destinationOwner: actualRecipient.publicKey,
          noncePage: noncePagePda,
        })
        .rpc();
      expect.fail("should have thrown");
    } catch (err: any) {
      expect(err.toString()).to.include("DestinationMismatch");
    }
  });

  // ---- Withdraw: paused ----

  it("rejects withdraw when withdraw_paused is true", async () => {
    await program.methods.setWithdrawPaused(true).accounts({}).rpc();

    const nonce = 5;
    const recipient6 = Keypair.generate();
    const destBuf = ethers.getBytes(
      ethers.hexlify(recipient6.publicKey.toBytes())
    );
    const tokenBuf = ethers.getBytes(
      addressToBytes32(ethers.hexlify(mint.toBytes()))
    );
    const user = v1.wallet.address;

    const sig1 = await signWithdraw(
      v1.wallet,
      domain,
      user,
      destBuf,
      tokenBuf,
      WITHDRAW_AMOUNT,
      CHAIN_ID,
      nonce
    );
    const sig2 = await signWithdraw(
      v2.wallet,
      domain,
      user,
      destBuf,
      tokenBuf,
      WITHDRAW_AMOUNT,
      CHAIN_ID,
      nonce
    );

    const [noncePagePda] = findNoncePagePda(nonce, program.programId);

    try {
      await program.methods
        .withdraw(
          Array.from(ethers.getBytes(user)),
          Array.from(recipient6.publicKey.toBytes()),
          new anchor.BN(WITHDRAW_AMOUNT.toString()),
          new anchor.BN(nonce),
          [
            { sig: sig1.sig, recoveryId: sig1.recoveryId },
            { sig: sig2.sig, recoveryId: sig2.recoveryId },
          ]
        )
        .accounts({
          mint,
          destinationOwner: recipient6.publicKey,
          noncePage: noncePagePda,
        })
        .rpc();
      expect.fail("should have thrown");
    } catch (err: any) {
      expect(err.toString()).to.include("WithdrawPaused");
    }

    // Restore
    await program.methods.setWithdrawPaused(false).accounts({}).rpc();
  });

  // ---- Withdraw: zero amount ----

  it("rejects zero amount withdraw", async () => {
    const nonce = 6;
    const recipient7 = Keypair.generate();
    const destBuf = ethers.getBytes(
      ethers.hexlify(recipient7.publicKey.toBytes())
    );
    const tokenBuf = ethers.getBytes(
      addressToBytes32(ethers.hexlify(mint.toBytes()))
    );
    const user = v1.wallet.address;

    const sig1 = await signWithdraw(
      v1.wallet,
      domain,
      user,
      destBuf,
      tokenBuf,
      0n,
      CHAIN_ID,
      nonce
    );
    const sig2 = await signWithdraw(
      v2.wallet,
      domain,
      user,
      destBuf,
      tokenBuf,
      0n,
      CHAIN_ID,
      nonce
    );

    const [noncePagePda] = findNoncePagePda(nonce, program.programId);

    try {
      await program.methods
        .withdraw(
          Array.from(ethers.getBytes(user)),
          Array.from(recipient7.publicKey.toBytes()),
          new anchor.BN(0),
          new anchor.BN(nonce),
          [
            { sig: sig1.sig, recoveryId: sig1.recoveryId },
            { sig: sig2.sig, recoveryId: sig2.recoveryId },
          ]
        )
        .accounts({
          mint,
          destinationOwner: recipient7.publicKey,
          noncePage: noncePagePda,
        })
        .rpc();
      expect.fail("should have thrown");
    } catch (err: any) {
      expect(err.toString()).to.include("ZeroAmount");
    }
  });

  // ---- Withdraw: high-S signature ----

  it("rejects high-S signature value", async () => {
    const nonce = 7;
    const recipient8 = Keypair.generate();
    const destBuf = ethers.getBytes(
      ethers.hexlify(recipient8.publicKey.toBytes())
    );
    const tokenBuf = ethers.getBytes(
      addressToBytes32(ethers.hexlify(mint.toBytes()))
    );
    const user = v1.wallet.address;

    const sig1 = await signWithdraw(
      v1.wallet,
      domain,
      user,
      destBuf,
      tokenBuf,
      WITHDRAW_AMOUNT,
      CHAIN_ID,
      nonce
    );

    // Flip S to high-S: S' = N - S where N is the secp256k1 order
    const N = BigInt(
      "0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFEBAAEDCE6AF48A03BBFD25E8CD0364141"
    );
    const sBytes = Buffer.from(sig1.sig.slice(32, 64));
    const s = BigInt("0x" + sBytes.toString("hex"));
    const highS = N - s;
    const highSHex = highS.toString(16).padStart(64, "0");
    const highSSig = [...sig1.sig.slice(0, 32), ...Buffer.from(highSHex, "hex")];
    // Recovery ID flips when S is negated
    const flippedRecoveryId = sig1.recoveryId ^ 1;

    const sig2 = await signWithdraw(
      v2.wallet,
      domain,
      user,
      destBuf,
      tokenBuf,
      WITHDRAW_AMOUNT,
      CHAIN_ID,
      nonce
    );

    const [noncePagePda] = findNoncePagePda(nonce, program.programId);

    try {
      await program.methods
        .withdraw(
          Array.from(ethers.getBytes(user)),
          Array.from(recipient8.publicKey.toBytes()),
          new anchor.BN(WITHDRAW_AMOUNT.toString()),
          new anchor.BN(nonce),
          [
            { sig: highSSig, recoveryId: flippedRecoveryId },
            { sig: sig2.sig, recoveryId: sig2.recoveryId },
          ]
        )
        .accounts({
          mint,
          destinationOwner: recipient8.publicKey,
          noncePage: noncePagePda,
        })
        .rpc();
      expect.fail("should have thrown");
    } catch (err: any) {
      expect(err.toString()).to.include("HighSValue");
    }
  });
});
