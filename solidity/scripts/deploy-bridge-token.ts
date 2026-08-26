// @ts-nocheck
import { ethers } from "hardhat";
import { DeploymentError, handleError } from "./utils/error-handler";

async function main() {
  console.log("开始部署 BridgeERC20 合约...");

  const [deployer] = await ethers.getSigners();
  console.log("部署账户:", deployer.address);

  const name = process.env.TOKEN_NAME || process.argv[2] || "USDT Token";
  const symbol = process.env.TOKEN_SYMBOL || process.argv[3] || "USDT";

  console.log(`Token 配置: name=${name}, symbol=${symbol}`);

  try {
    const BridgeERC20 = await ethers.getContractFactory("BridgeERC20");
    const token = await BridgeERC20.deploy(name, symbol);
    await token.waitForDeployment();
    const tokenAddress = await token.getAddress();
    console.log("✅ BridgeERC20 部署地址:", tokenAddress);
    console.log(`DEPLOYED_ADDRESS=${tokenAddress}`);

    const mintTo = process.env.MINT_TO;
    if (mintTo) {
      const mintAmount = BigInt(
        process.env.MINT_AMOUNT || "1000000000000000000"
      );
      const mintTx = await token.mint(mintTo, mintAmount);
      await mintTx.wait();
      console.log(`✅ 已向 ${mintTo} 铸造 ${mintAmount.toString()} wei`);
    }
  } catch (error) {
    throw new DeploymentError(
      "BridgeERC20 合约部署失败",
      error instanceof Error ? error.message : String(error),
      "请检查网络连接和账户余额，然后重试"
    );
  }
}

main()
  .then(() => {
    console.log("🎉 部署完成!");
    process.exit(0);
  })
  .catch(handleError);
