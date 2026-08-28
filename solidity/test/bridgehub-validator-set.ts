import { ethers } from "hardhat";
import { expect } from "chai";
import { HardhatEthersSigner } from "@nomicfoundation/hardhat-ethers/signers";
import { HDNodeWallet } from "ethers";

const HARDHAT_MNEMONIC =
  "test test test test test test test test test test test junk";

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

describe("BridgeHub validator set updates", function () {
  let bridgeHub: any;
  let admin: HardhatEthersSigner;
  let coldValidator: HardhatEthersSigner;
  let hotA: HardhatEthersSigner;
  let hotB: HardhatEthersSigner;

  async function signUpdateValidatorSet(
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
    bridgeHub = BridgeHub.attach(await proxy.getAddress());
  }

  beforeEach(async function () {
    const signers = await ethers.getSigners();
    admin = signers[0];
    coldValidator = signers[1];
    hotA = signers[3];
    hotB = signers[4];
    await deployBridgeHub();
  });

  it("finalizes a requested validator set update", async function () {
    const update = {
      epoch: 1n,
      hotAddresses: [hotA.address],
      coldAddresses: [coldValidator.address],
      powers: [100n],
    };

    await bridgeHub.updateValidatorSet(
      update.epoch,
      update.hotAddresses,
      update.coldAddresses,
      update.powers
    );

    const sig = await signUpdateValidatorSet(admin, update);
    await bridgeHub.connect(admin).updateValidatorSetConfirm(update, {
      r: sig.r,
      s: sig.s,
      v: sig.v,
    });

    expect(await bridgeHub.epoch()).to.equal(1n);
    expect(await bridgeHub.getHotValidators()).to.deep.equal([hotA.address]);
  });

  it("rejects confirming a stale validator set after a newer epoch is finalized", async function () {
    const staleUpdate = {
      epoch: 5n,
      hotAddresses: [hotA.address],
      coldAddresses: [coldValidator.address],
      powers: [100n],
    };
    const newerUpdate = {
      epoch: 6n,
      hotAddresses: [hotB.address],
      coldAddresses: [coldValidator.address],
      powers: [100n],
    };

    await bridgeHub.updateValidatorSet(
      staleUpdate.epoch,
      staleUpdate.hotAddresses,
      staleUpdate.coldAddresses,
      staleUpdate.powers
    );
    await bridgeHub.updateValidatorSet(
      newerUpdate.epoch,
      newerUpdate.hotAddresses,
      newerUpdate.coldAddresses,
      newerUpdate.powers
    );

    const newerSig = await signUpdateValidatorSet(admin, newerUpdate);
    await bridgeHub.connect(admin).updateValidatorSetConfirm(newerUpdate, {
      r: newerSig.r,
      s: newerSig.s,
      v: newerSig.v,
    });

    expect(await bridgeHub.epoch()).to.equal(6n);
    expect(await bridgeHub.getHotValidators()).to.deep.equal([hotB.address]);

    const staleSig = await signUpdateValidatorSet(admin, staleUpdate);
    await expect(
      bridgeHub.connect(hotB).updateValidatorSetConfirm(staleUpdate, {
        r: staleSig.r,
        s: staleSig.s,
        v: staleSig.v,
      })
    ).to.be.revertedWith("Stale validator set update");

    expect(await bridgeHub.epoch()).to.equal(6n);
    expect(await bridgeHub.getHotValidators()).to.deep.equal([hotB.address]);
  });

  it("rejects confirming the same epoch twice", async function () {
    const update = {
      epoch: 1n,
      hotAddresses: [hotA.address],
      coldAddresses: [coldValidator.address],
      powers: [100n],
    };

    await bridgeHub.updateValidatorSet(
      update.epoch,
      update.hotAddresses,
      update.coldAddresses,
      update.powers
    );

    const sig = await signUpdateValidatorSet(admin, update);
    await bridgeHub.connect(admin).updateValidatorSetConfirm(update, {
      r: sig.r,
      s: sig.s,
      v: sig.v,
    });

    expect(await bridgeHub.epoch()).to.equal(1n);

    await expect(
      bridgeHub.connect(hotA).updateValidatorSetConfirm(update, {
        r: sig.r,
        s: sig.s,
        v: sig.v,
      })
    ).to.be.revertedWith("Stale validator set update");
  });
});
