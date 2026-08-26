import { ethers } from "hardhat";
import { expect } from "chai";

const UPDATE_VALIDATOR_SET_TYPEHASH = ethers.keccak256(
  ethers.toUtf8Bytes(
    "UpdateValidatorSet(uint64 epoch,address[] hotAddresses,address[] coldAddresses,uint64[] powers)"
  )
);

describe("bridge contract tests", function () {
  it("UPDATE_VALIDATOR_SET_TYPEHASH matches BridgeHub.makeUpdateValidatorSetMessage", async function () {
    const signers = await ethers.getSigners();
    const hot = signers[0];
    const cold = signers[1];

    const BridgeHub = await ethers.getContractFactory("BridgeHub");
    const impl = await BridgeHub.deploy();
    const initData = impl.interface.encodeFunctionData("initialize", [
      [hot.address],
      [cold.address],
      [100n],
    ]);
    const ERC1967Proxy = await ethers.getContractFactory("ERC1967Proxy");
    const proxy = await ERC1967Proxy.deploy(await impl.getAddress(), initData);
    const bridgeHub = BridgeHub.attach(await proxy.getAddress());

    const hotAddresses = [hot.address];
    const coldAddresses = [cold.address];
    const powers = [2n];
    const epoch = 1n;

    const packedHot = ethers.solidityPacked(["address[]"], [hotAddresses]);
    const packedCold = ethers.solidityPacked(["address[]"], [coldAddresses]);
    const packedPowers = ethers.solidityPacked(["uint64[]"], [powers]);
    const expected = ethers.keccak256(
      ethers.AbiCoder.defaultAbiCoder().encode(
        ["bytes32", "uint64", "bytes32", "bytes32", "bytes32"],
        [
          UPDATE_VALIDATOR_SET_TYPEHASH,
          epoch,
          ethers.keccak256(packedHot),
          ethers.keccak256(packedCold),
          ethers.keccak256(packedPowers),
        ]
      )
    );

    expect(
      await bridgeHub.makeUpdateValidatorSetMessage(
        epoch,
        hotAddresses,
        coldAddresses,
        powers
      )
    ).to.equal(expected);
  });
});
