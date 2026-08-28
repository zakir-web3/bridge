import { ethers } from "hardhat";
import { expect } from "chai";
import { anyValue } from "@nomicfoundation/hardhat-chai-matchers/withArgs";
import { HardhatEthersSigner } from "@nomicfoundation/hardhat-ethers/signers";

const SRC_CHAIN_ID = 1337n;

function addressToBytes32(address: string) {
  return ethers.zeroPadValue(address, 32);
}

describe("BridgeHub Decimal Conversion", function () {
  let bridgeHub: any;
  let srcToken: any;
  let bridgedToken: any;
  let admin: HardhatEthersSigner;
  let user: HardhatEthersSigner;

  async function signDeposit(deposit: {
    user: string;
    destination: string;
    token: string;
    amount: bigint;
    chainId: bigint;
    blockNumber: bigint;
    txHash: string;
    index: number;
  }) {
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
    const raw = await admin.signTypedData(domain, types, deposit);
    return ethers.Signature.from(raw);
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

  it("initializes epoch 0 and the hot validator set", async function () {
    expect(await bridgeHub.epoch()).to.equal(0n);
    expect(await bridgeHub.getHotValidators()).to.deep.equal([admin.address]);
  });

  it("returns 0 for an unconfigured tokenDecimalDiff", async function () {
    expect(
      await bridgeHub.tokenDecimalDiff(
        SRC_CHAIN_ID,
        addressToBytes32(admin.address)
      )
    ).to.equal(0n);
  });

  it("stores -12 on src pair when pairing a 6-decimal token with an 18-decimal bridged token", async function () {
    const src = await srcToken.getAddress();
    const bridged = await bridgedToken.getAddress();
    await bridgeHub.setTokenPair(
      SRC_CHAIN_ID,
      addressToBytes32(src),
      6,
      bridged
    );

    expect(
      await bridgeHub.tokenDecimalDiff(SRC_CHAIN_ID, addressToBytes32(src))
    ).to.equal(-12n);
  });

  it("rejects setTokenPair when the src token pair already exists", async function () {
    await pairTokens(6);

    await expect(
      bridgeHub.setTokenPair(
        SRC_CHAIN_ID,
        addressToBytes32(await srcToken.getAddress()),
        6,
        await bridgedToken.getAddress()
      )
    ).to.be.revertedWith("Token pair already set");
  });

  it("rejects setTokenPair when src decimals exceed int8-safe range", async function () {
    await expect(
      bridgeHub.setTokenPair(
        SRC_CHAIN_ID,
        addressToBytes32(await srcToken.getAddress()),
        78,
        await bridgedToken.getAddress()
      )
    ).to.be.revertedWith("Invalid src decimals");
  });

  it("stores 12 on dst pair for withdraw decimal conversion", async function () {
    const src = await srcToken.getAddress();
    const bridged = await bridgedToken.getAddress();
    await bridgeHub.setTokenPair(
      SRC_CHAIN_ID,
      addressToBytes32(src),
      6,
      bridged
    );

    expect(
      await bridgeHub.tokenDecimalDiff(SRC_CHAIN_ID, addressToBytes32(bridged))
    ).to.equal(12n);
  });

  it("mints 100e18 on depositConfirm of 100e6", async function () {
    await pairTokens(6);

    const amount = 100_000_000n;
    const deposit = {
      user: addressToBytes32(user.address),
      destination: user.address,
      token: addressToBytes32(await srcToken.getAddress()),
      amount,
      chainId: SRC_CHAIN_ID,
      blockNumber: 1n,
      txHash: ethers.ZeroHash,
      index: 0,
    };
    const sig = await signDeposit(deposit);

    await bridgeHub.depositConfirm([
      {
        ...deposit,
        signature: { r: sig.r, s: sig.s, v: sig.v },
      },
    ]);

    expect(await bridgedToken.balanceOf(user.address)).to.equal(
      100_000_000_000_000_000_000n
    );
  });

  it("writes 100e6 into the withdraw message when burning 100e18", async function () {
    await pairTokens(6);
    await bridgedToken.mint(user.address, 100_000_000_000_000_000_000n);
    await bridgedToken
      .connect(user)
      .approve(await bridgeHub.getAddress(), 100_000_000_000_000_000_000n);

    await expect(
      bridgeHub
        .connect(user)
        .withdraw(
          user.address,
          await bridgedToken.getAddress(),
          100_000_000_000_000_000_000n,
          SRC_CHAIN_ID
        )
    )
      .to.emit(bridgeHub, "Withdraw")
      .withArgs(
        anyValue,
        user.address,
        user.address,
        await srcToken.getAddress(),
        100_000_000n,
        SRC_CHAIN_ID,
        1n
      );
  });
});
