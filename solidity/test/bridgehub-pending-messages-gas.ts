import { ethers } from "hardhat";
import { expect } from "chai";
import { HardhatEthersSigner } from "@nomicfoundation/hardhat-ethers/signers";
import { HDNodeWallet } from "ethers";

const SRC_CHAIN_ID = 1337n;
const WITHDRAW_AMOUNT = 1_000_000_000_000_000_000n;
const HARDHAT_MNEMONIC =
  "test test test test test test test test test test test junk";

/** Approximate per-chain block gas limits for reference in output. */
const BLOCK_GAS_LIMITS = {
  ethereum: 30_000_000n,
  arbitrum: 32_000_000n,
  bsc: 140_000_000n,
};

function addressToBytes32(address: string) {
  return ethers.zeroPadValue(address, 32);
}

function walletFor(signer: HardhatEthersSigner): HDNodeWallet {
  for (let i = 0; i < 20; i++) {
    const wallet = HDNodeWallet.fromPhrase(
      HARDHAT_MNEMONIC,
      undefined,
      `m/44'/60'/0'/0/${i}`
    );
    if (wallet.address.toLowerCase() === signer.address.toLowerCase()) {
      return wallet;
    }
  }
  throw new Error(`No hardhat wallet for ${signer.address}`);
}

describe("BridgeHub pendingMessages gas benchmark", function () {
  this.timeout(600_000);

  let admin: HardhatEthersSigner;
  let coldValidator: HardhatEthersSigner;
  let user: HardhatEthersSigner;

  async function deployBridgeHub() {
    const BridgeHub = await ethers.getContractFactory("BridgeHub");
    const impl = await BridgeHub.deploy();
    const initData = impl.interface.encodeFunctionData("initialize", [
      [admin.address],
      [coldValidator.address],
      [100n],
    ]);
    const ERC1967Proxy = await ethers.getContractFactory("ERC1967Proxy");
    const proxy = await ERC1967Proxy.deploy(await impl.getAddress(), initData);
    const bridgeHub: any = BridgeHub.attach(await proxy.getAddress());

    const ERC20Mock = await ethers.getContractFactory("ERC20Mock");
    const srcToken = await ERC20Mock.deploy("Token", "TKN", 18);

    const BridgeERC20 = await ethers.getContractFactory("BridgeERC20");
    const bridgedToken = await BridgeERC20.deploy("Bridged TKN", "TKN.b");
    await bridgedToken.setMiner(await bridgeHub.getAddress(), true);

    await bridgeHub.setTokenPair(
      SRC_CHAIN_ID,
      addressToBytes32(await srcToken.getAddress()),
      18,
      await bridgedToken.getAddress()
    );

    return { bridgeHub, srcToken, bridgedToken };
  }

  async function signWithdraw(
    bridgeHub: any,
    signer: HardhatEthersSigner,
    withdraw: {
      user: string;
      destination: string;
      token: string;
      amount: bigint;
      chainId: bigint;
      nonce: bigint;
    }
  ) {
    const network = await ethers.provider.getNetwork();
    const domain = {
      name: "BridgeHub",
      version: "1",
      chainId: network.chainId,
      verifyingContract: await bridgeHub.getAddress(),
    };
    const types = {
      Withdraw: [
        { name: "user", type: "address" },
        { name: "destination", type: "address" },
        { name: "token", type: "address" },
        { name: "amount", type: "uint256" },
        { name: "chainId", type: "uint256" },
        { name: "nonce", type: "uint64" },
      ],
    };
    return ethers.Signature.from(
      await signer.signTypedData(domain, types, withdraw)
    );
  }

  async function signUpdateValidatorSet(
    bridgeHub: any,
    signer: HardhatEthersSigner,
    update: {
      epoch: bigint;
      hotAddresses: string[];
      coldAddresses: string[];
      powers: bigint[];
    }
  ) {
    const domainSeparator = await bridgeHub.domainSeparator();
    const messageHash = await bridgeHub.makeUpdateValidatorSetMessage(
      update.epoch,
      update.hotAddresses,
      update.coldAddresses,
      update.powers
    );
    const digest = ethers.keccak256(
      ethers.concat(["0x1901", domainSeparator, messageHash])
    );
    return ethers.Signature.from(walletFor(signer).signingKey.sign(digest));
  }

  async function fillPendingWithdraws(
    bridgeHub: any,
    bridgedToken: any,
    count: number
  ) {
    const bridged = await bridgedToken.getAddress();
    await bridgedToken.mint(user.address, WITHDRAW_AMOUNT * BigInt(count));
    await bridgedToken
      .connect(user)
      .approve(await bridgeHub.getAddress(), WITHDRAW_AMOUNT * BigInt(count));

    for (let i = 0; i < count; i++) {
      await bridgeHub
        .connect(user)
        .withdraw(user.address, bridged, WITHDRAW_AMOUNT, SRC_CHAIN_ID);
    }
  }

  async function fillPendingValidatorSetUpdates(bridgeHub: any, count: number) {
    for (let i = 0; i < count; i++) {
      await bridgeHub.updateValidatorSet(
        BigInt(i + 1),
        [admin.address],
        [coldValidator.address],
        [100n]
      );
    }
  }

  async function measureWithdrawConfirmGas(
    bridgeHub: any,
    srcToken: any,
    targetNonce: bigint
  ) {
    const token = await srcToken.getAddress();
    const sig = await signWithdraw(bridgeHub, admin, {
      user: user.address,
      destination: user.address,
      token,
      amount: WITHDRAW_AMOUNT,
      chainId: SRC_CHAIN_ID,
      nonce: targetNonce,
    });

    const tx = await bridgeHub.connect(admin).withdrawConfirm([
      {
        user: user.address,
        destination: user.address,
        token,
        amount: WITHDRAW_AMOUNT,
        chainId: SRC_CHAIN_ID,
        nonce: targetNonce,
        signature: { r: sig.r, s: sig.s, v: sig.v },
      },
    ]);
    const receipt = await tx.wait();
    return BigInt(receipt!.gasUsed);
  }

  async function measureValidatorSetConfirmGas(
    bridgeHub: any,
    targetEpoch: bigint
  ) {
    const update = {
      epoch: targetEpoch,
      hotAddresses: [admin.address],
      coldAddresses: [coldValidator.address],
      powers: [100n],
    };
    const sig = await signUpdateValidatorSet(bridgeHub, admin, update);
    const tx = await bridgeHub.connect(admin).updateValidatorSetConfirm(update, {
      r: sig.r,
      s: sig.s,
      v: sig.v,
    });
    const receipt = await tx.wait();
    return BigInt(receipt!.gasUsed);
  }

  type GasRow = {
    pendingCount: number;
    withdrawConfirmFirst: bigint;
    withdrawConfirmLast: bigint;
    scanDelta: bigint;
    validatorSetConfirmLast: bigint;
  };

  beforeEach(async function () {
    const signers = await ethers.getSigners();
    admin = signers[0];
    coldValidator = signers[1];
    user = signers[2];
  });

  it("reports withdrawConfirm / updateValidatorSetConfirm gas vs pendingMessages length", async function () {
    const sizes = [1, 10, 25, 50, 100, 200, 500];
    const rows: GasRow[] = [];

    for (const n of sizes) {
      let ctx = await deployBridgeHub();
      await fillPendingWithdraws(ctx.bridgeHub, ctx.bridgedToken, n);
      expect(await ctx.bridgeHub.getPendingMessages()).to.have.length(n);
      const lastGas = await measureWithdrawConfirmGas(
        ctx.bridgeHub,
        ctx.srcToken,
        BigInt(n)
      );

      ctx = await deployBridgeHub();
      await fillPendingWithdraws(ctx.bridgeHub, ctx.bridgedToken, n);
      const firstGas = await measureWithdrawConfirmGas(
        ctx.bridgeHub,
        ctx.srcToken,
        1n
      );

      ctx = await deployBridgeHub();
      await fillPendingValidatorSetUpdates(ctx.bridgeHub, n);
      expect(await ctx.bridgeHub.getPendingMessages()).to.have.length(n);
      const validatorGas = await measureValidatorSetConfirmGas(
        ctx.bridgeHub,
        BigInt(n)
      );

      rows.push({
        pendingCount: n,
        withdrawConfirmFirst: firstGas,
        withdrawConfirmLast: lastGas,
        scanDelta: lastGas - firstGas,
        validatorSetConfirmLast: validatorGas,
      });
    }

    const baseline = rows[0];
    const largest = rows[rows.length - 1];
    const perItem =
      largest.pendingCount > baseline.pendingCount
        ? (largest.withdrawConfirmLast - baseline.withdrawConfirmLast) /
          BigInt(largest.pendingCount - baseline.pendingCount)
        : 0n;

    console.log("\n=== BridgeHub pendingMessages gas benchmark ===\n");
    console.log(
      "pending | withdrawConfirm (1st) | withdrawConfirm (last) | scan Δ | validatorSetConfirm (last)"
    );
    for (const row of rows) {
      console.log(
        `${String(row.pendingCount).padStart(7)} | ${String(row.withdrawConfirmFirst).padStart(21)} | ${String(row.withdrawConfirmLast).padStart(22)} | ${String(row.scanDelta).padStart(6)} | ${String(row.validatorSetConfirmLast).padStart(27)}`
      );
    }

    console.log("\n--- Analysis ---");
    console.log(
      `Baseline (pending=1, last): ${baseline.withdrawConfirmLast} gas`
    );
    console.log(
      `At pending=${largest.pendingCount}, last confirm: ${largest.withdrawConfirmLast} gas (+${((Number(largest.withdrawConfirmLast) / Number(baseline.withdrawConfirmLast) - 1) * 100).toFixed(0)}%)`
    );
    console.log(
      `Approx. linear scan cost: ~${perItem} gas per extra pending message`
    );

    const thresholds = [
      { label: "+50% vs baseline", factor: 1.5 },
      { label: "+100% vs baseline (2x)", factor: 2.0 },
      { label: "+500% vs baseline (6x)", factor: 6.0 },
    ];
    const base = Number(baseline.withdrawConfirmLast);
    for (const t of thresholds) {
      const target = base * t.factor;
      const hit = rows.find((r) => Number(r.withdrawConfirmLast) >= target);
      console.log(
        `${t.label} (${Math.round(target).toLocaleString()} gas): ${hit ? `~${hit.pendingCount} pending` : "not reached in this run"}`
      );
    }

    if (perItem > 0n) {
      console.log("\n--- Extrapolated pending count (withdrawConfirm last) ---");
      for (const [chain, limit] of Object.entries(BLOCK_GAS_LIMITS)) {
        for (const pct of [10, 33, 50]) {
          const budget = (limit * BigInt(pct)) / 100n;
          const extra =
            budget > baseline.withdrawConfirmLast
              ? (budget - baseline.withdrawConfirmLast) / perItem
              : 0n;
          const total = BigInt(baseline.pendingCount) + extra;
          console.log(
            `${chain} ${pct}% block (${budget.toLocaleString()} gas): ~${total.toLocaleString()} pending`
          );
        }
      }
    }

    expect(largest.withdrawConfirmLast).to.be.greaterThan(
      baseline.withdrawConfirmLast
    );
    expect(largest.scanDelta).to.be.greaterThan(baseline.scanDelta);
  });
});
