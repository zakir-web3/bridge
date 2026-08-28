import { ethers } from "hardhat";
import { expect } from "chai";
import { anyValue } from "@nomicfoundation/hardhat-chai-matchers/withArgs";
import { HardhatEthersSigner } from "@nomicfoundation/hardhat-ethers/signers";

const SRC_CHAIN_ID = 1337n;

function addressToBytes32(address: string) {
  return ethers.zeroPadValue(address, 32);
}

describe("BridgeHub core flows", function () {
  let bridgeHub: any;
  let srcToken: any;
  let bridgedToken: any;
  let admin: HardhatEthersSigner;
  let user: HardhatEthersSigner;

  async function signDeposit(
    signer: HardhatEthersSigner,
    deposit: {
      user: string;
      destination: string;
      token: string;
      amount: bigint;
      chainId: bigint;
      blockNumber: bigint;
      txHash: string;
      index: number;
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
      Deposit: [
        { name: "user", type: "bytes32" },
        { name: "destination", type: "address" },
        { name: "token", type: "bytes32" },
        { name: "amount", type: "uint256" },
        { name: "chainId", type: "uint256" },
        { name: "blockNumber", type: "uint64" },
        { name: "txHash", type: "bytes32" },
        { name: "index", type: "uint32" },
      ],
    };
    return ethers.Signature.from(
      await signer.signTypedData(domain, types, deposit)
    );
  }

  async function pairTokens(srcDecimals: number) {
    const src = await srcToken.getAddress();
    const bridged = await bridgedToken.getAddress();
    const srcBytes32 = addressToBytes32(src);
    await bridgeHub.setTokenPair(
      SRC_CHAIN_ID,
      srcBytes32,
      srcDecimals,
      bridged
    );
  }

  beforeEach(async function () {
    const signers = await ethers.getSigners();
    admin = signers[0];
    user = signers[2];
    const coldValidator = signers[1];

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

    const ERC20Mock = await ethers.getContractFactory("ERC20Mock");
    srcToken = await ERC20Mock.deploy("USD Coin", "USDC", 6);

    const BridgeERC20 = await ethers.getContractFactory("BridgeERC20");
    bridgedToken = await BridgeERC20.deploy("Bridged USDC", "USDC.b");
    await bridgedToken.setMiner(await bridgeHub.getAddress(), true);
  });

  it("rejects depositConfirm from a non-validator", async function () {
    await pairTokens(6);
    const deposit = {
      user: addressToBytes32(user.address),
      destination: user.address,
      token: addressToBytes32(await srcToken.getAddress()),
      amount: 100_000_000n,
      chainId: SRC_CHAIN_ID,
      blockNumber: 1n,
      txHash: ethers.ZeroHash,
      index: 0,
    };
    const sig = await signDeposit(admin, deposit);

    await expect(
      bridgeHub
        .connect(user)
        .depositConfirm([
          { ...deposit, signature: { r: sig.r, s: sig.s, v: sig.v } },
        ])
    ).to.be.revertedWith("Signer is not a validator");
  });

  it("rejects a second depositConfirm for the same message", async function () {
    await pairTokens(6);
    const deposit = {
      user: addressToBytes32(user.address),
      destination: user.address,
      token: addressToBytes32(await srcToken.getAddress()),
      amount: 100_000_000n,
      chainId: SRC_CHAIN_ID,
      blockNumber: 1n,
      txHash: ethers.ZeroHash,
      index: 0,
    };
    const sig = await signDeposit(admin, deposit);
    const payload = [
      { ...deposit, signature: { r: sig.r, s: sig.s, v: sig.v } },
    ];

    await bridgeHub.depositConfirm(payload);
    await expect(bridgeHub.depositConfirm(payload)).to.be.revertedWith(
      "Already processed"
    );
  });

  it("rejects withdraw when the token pair is missing", async function () {
    await bridgedToken.mint(user.address, 1_000_000_000_000_000_000n);
    await bridgedToken
      .connect(user)
      .approve(await bridgeHub.getAddress(), 1_000_000_000_000_000_000n);

    await expect(
      bridgeHub
        .connect(user)
        .withdraw(
          user.address,
          await bridgedToken.getAddress(),
          1_000_000_000_000_000_000n,
          SRC_CHAIN_ID
        )
    ).to.be.revertedWith("Token not found");
  });

  it("subtracts withdraw fee before writing the cross-chain amount", async function () {
    await pairTokens(6);
    const amount = 100_000_000_000_000_000_000n;
    const fee = 1_000_000_000_000_000_000n;
    await bridgedToken.mint(user.address, amount);
    await bridgedToken
      .connect(user)
      .approve(await bridgeHub.getAddress(), amount);
    await bridgeHub.setWithdrawFee(await bridgedToken.getAddress(), fee);

    await expect(
      bridgeHub
        .connect(user)
        .withdraw(
          user.address,
          await bridgedToken.getAddress(),
          fee,
          SRC_CHAIN_ID
        )
    ).to.be.revertedWith("Amount must exceed fee");

    await expect(
      bridgeHub
        .connect(user)
        .withdraw(
          user.address,
          await bridgedToken.getAddress(),
          amount,
          SRC_CHAIN_ID
        )
    )
      .to.emit(bridgeHub, "Withdraw")
      .withArgs(
        anyValue,
        user.address,
        user.address,
        await srcToken.getAddress(),
        99_000_000n,
        SRC_CHAIN_ID,
        1n
      );
  });

  it("blocks withdraw while paused", async function () {
    await pairTokens(6);
    await bridgeHub.pause();

    await expect(
      bridgeHub
        .connect(user)
        .withdraw(
          user.address,
          await bridgedToken.getAddress(),
          1n,
          SRC_CHAIN_ID
        )
    ).to.be.revertedWithCustomError(bridgeHub, "EnforcedPause");
  });

  it("keeps amount unchanged when source and destination decimals match", async function () {
    const ERC20Mock = await ethers.getContractFactory("ERC20Mock");
    srcToken = await ERC20Mock.deploy("USDT", "USDT", 18);
    await pairTokens(18);

    const amount = 1_000_000_000_000_000_000n;
    const deposit = {
      user: addressToBytes32(user.address),
      destination: user.address,
      token: addressToBytes32(await srcToken.getAddress()),
      amount,
      chainId: SRC_CHAIN_ID,
      blockNumber: 1n,
      txHash: ethers.ZeroHash,
      index: 1,
    };
    const sig = await signDeposit(admin, deposit);

    await bridgeHub.depositConfirm([
      { ...deposit, signature: { r: sig.r, s: sig.s, v: sig.v } },
    ]);
    expect(await bridgedToken.balanceOf(user.address)).to.equal(amount);
  });
});
