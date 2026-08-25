import { ethers } from "hardhat";
import {
  validateAddress,
  handleError,
  PermissionError,
} from "./utils/error-handler";

async function main() {
  console.log("开始设置 ERC20 的 Minter 权限...");

  const [signer] = await ethers.getSigners();
  console.log("调用账户:", signer.address);

  const erc20Address = process.env.ERC20_ADDRESS || process.argv[2];
  const minerAddress = process.env.MINER_ADDRESS || process.argv[3];

  // 使用统一的验证函数
  validateAddress(erc20Address!, "ERC20合约");
  validateAddress(minerAddress!, "Minter授权");

  // 仅支持 true，合约内对 false 会直接 revert
  const isActive = true;

  try {
    const erc20 = await ethers.getContractAt("BridgeERC20", erc20Address!);

    console.log(
      `调用 setMiner(miner=${minerAddress}, isActive=${isActive})...`
    );
    const tx = await (erc20 as any).setMiner(minerAddress!, isActive);
    console.log("交易哈希:", tx.hash);
    await tx.wait();

    console.log("✅ 完成: ERC20 的 MINTER_ROLE 权限已更新");
  } catch (error) {
    throw new PermissionError(
      "设置 Minter 权限失败",
      error instanceof Error ? error.message : String(error),
      "请检查账户权限和合约状态，然后重试"
    );
  }
}

main()
  .then(() => {
    console.log("🎉 操作完成!");
    process.exit(0);
  })
  .catch(handleError);
