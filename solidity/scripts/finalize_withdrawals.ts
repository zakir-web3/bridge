import { ethers } from "hardhat";
import { Bridge } from "../typechain-types";

async function main() {
  // 配置参数
  const BRIDGE_CONTRACT_ADDRESS = process.env.BRIDGE_CONTRACT_ADDRESS;
  const PRIVATE_KEY = process.env.PRIVATE_KEY;
  const WITHDRAWAL_MESSAGES = process.env.WITHDRAWAL_MESSAGES; // 可选的提款消息数组，用逗号分隔

  if (!BRIDGE_CONTRACT_ADDRESS) {
    throw new Error("请设置环境变量 BRIDGE_CONTRACT_ADDRESS");
  }

  if (!PRIVATE_KEY) {
    throw new Error("请设置环境变量 PRIVATE_KEY");
  }

  console.log("🚀 开始执行批量完成提款操作...");
  console.log("📋 配置信息:");
  console.log(`   Bridge 合约地址: ${BRIDGE_CONTRACT_ADDRESS}`);

  // 创建钱包实例
  const [signer] = await ethers.getSigners();
  console.log(`👤 验证人地址: ${signer.address}`);

  // 获取网络提供者
  const network = await ethers.provider.getNetwork();
  console.log(`🌐 当前网络: ${network.name} (Chain ID: ${network.chainId})`);

  // 检查账户余额
  const balance = await ethers.provider.getBalance(signer.address);
  console.log(`💰 账户 ETH 余额: ${ethers.formatEther(balance)} ETH`);

  if (balance === 0n) {
    throw new Error("账户余额不足，无法支付 gas 费用");
  }

  // 获取 Bridge 合约实例
  const bridge = (await ethers.getContractAt(
    "Bridge",
    BRIDGE_CONTRACT_ADDRESS
  )) as Bridge;
  console.log("✅ Bridge 合约实例创建成功");

  // 检查是否为 finalizer
  const isFinalizer = await bridge.finalizers(signer.address);
  if (!isFinalizer) {
    throw new Error(
      `地址 ${signer.address} 不是 finalizer，无法执行 finalize 操作`
    );
  }
  console.log("✅ 验证人身份确认：具有 finalizer 权限");

  // 获取争议期信息
  const disputePeriodSeconds = await bridge.disputePeriodSeconds();
  const blockDurationMillis = await bridge.blockDurationMillis();
  console.log(
    `⏰ 争议期配置: ${disputePeriodSeconds} 秒, 区块时长: ${blockDurationMillis} 毫秒`
  );

  // 获取待完成的提款消息
  let withdrawalMessages: string[] = [];

  if (WITHDRAWAL_MESSAGES) {
    // 如果提供了具体的提款消息，使用提供的消息
    withdrawalMessages = WITHDRAWAL_MESSAGES.split(",").map((msg) =>
      msg.trim()
    );
    console.log(`📝 使用提供的提款消息: ${withdrawalMessages.length} 条`);
  } else {
    // 如果没有提供具体消息，尝试获取所有待完成的提款
    console.log("🔍 正在查询待完成的提款请求...");

    // 注意：由于合约没有直接的查询接口来获取所有待完成的提款
    // 我们建议用户手动提供需要完成的提款消息
    console.log(
      "⚠️  未提供具体的提款消息，请设置 WITHDRAWAL_MESSAGES 环境变量"
    );
    console.log("   格式: WITHDRAWAL_MESSAGES=0x123...,0x456...");
    console.log("   或者从区块链浏览器中的合约事件源获取提款消息:");
    throw new Error("请提供需要完成的提款消息");
  }

  if (withdrawalMessages.length === 0) {
    throw new Error("没有提供任何提款消息");
  }

  // 验证消息格式和状态
  for (let i = 0; i < withdrawalMessages.length; i++) {
    const message = withdrawalMessages[i];
    if (!ethers.isHexString(message, 32)) {
      throw new Error(`提款消息 ${i + 1} 格式错误: ${message}`);
    }

    // 检查提款是否已经完成
    const isFinalized = await bridge.finalizedWithdrawals(message);
    if (isFinalized) {
      console.log(`⚠️  提款消息 ${i + 1} 已经完成: ${message}`);
      withdrawalMessages.splice(i, 1);
      i--;
      continue;
    }

    // 获取提款详细信息
    try {
      const withdrawal = await bridge.requestedWithdrawals(message);
      if (withdrawal.requestedTime === 0n) {
        console.log(`⚠️  提款消息 ${i + 1} 不存在: ${message}`);
        withdrawalMessages.splice(i, 1);
        i--;
        continue;
      }

      // 检查争议期是否已过
      const currentTime = BigInt(Math.floor(Date.now() / 1000));
      const disputePeriodSeconds = await bridge.disputePeriodSeconds();
      const timePassed = currentTime - withdrawal.requestedTime;

      if (timePassed < disputePeriodSeconds) {
        const remainingTime = disputePeriodSeconds - timePassed;
        console.log(
          `⚠️  提款消息 ${
            i + 1
          } 争议期未过，还需等待 ${remainingTime} 秒: ${message}`
        );
        withdrawalMessages.splice(i, 1);
        i--;
        continue;
      }

      console.log(`📋 提款 ${i + 1} 详情:`);
      console.log(`   用户地址: ${withdrawal.user}`);
      console.log(`   目标地址: ${withdrawal.destination}`);
      console.log(`   代币地址: ${withdrawal.token}`);
      console.log(`   金额: ${withdrawal.amount} wei`);
      console.log(`   链 ID: ${withdrawal.chainId}`);
      console.log(`   Nonce: ${withdrawal.nonce}`);
      console.log(
        `   请求时间: ${new Date(
          Number(withdrawal.requestedTime) * 1000
        ).toISOString()}`
      );
    } catch (error) {
      console.log(`⚠️  无法获取提款消息 ${i + 1} 的详细信息: ${message}`);
      console.log(`   错误: ${error}`);
    }
  }

  if (withdrawalMessages.length === 0) {
    console.log("✅ 所有提款请求都已经完成");
    return;
  }

  console.log(`📋 待完成的提款请求: ${withdrawalMessages.length} 条`);
  withdrawalMessages.forEach((msg, index) => {
    console.log(`   ${index + 1}. ${msg}`);
  });

  // 执行批量完成提款操作
  console.log("💸 开始执行提款操作...");

  const finalizeTx = await bridge.batchedFinalizeWithdrawals(
    withdrawalMessages
  );

  console.log(`📝 提款交易已发送，交易哈希: ${finalizeTx.hash}`);
  console.log("⏳ 等待交易确认...");

  const receipt = await finalizeTx.wait();
  console.log(`✅ 提款交易成功！`);
  console.log(`   区块号: ${receipt?.blockNumber}`);
  console.log(`   Gas 使用量: ${receipt?.gasUsed?.toString()}`);
  console.log(
    `   Gas 价格: ${ethers.formatUnits(receipt?.gasPrice || 0, "gwei")} Gwei`
  );
  console.log(
    `   总费用: ${ethers.formatEther(
      (receipt?.gasUsed || 0n) * (receipt?.gasPrice || 0n)
    )} ETH`
  );

  // 验证提款状态
  console.log("🔍 验证提款完成状态...");
  for (let i = 0; i < withdrawalMessages.length; i++) {
    const message = withdrawalMessages[i];
    const isFinalized = await bridge.finalizedWithdrawals(message);
    if (isFinalized) {
      console.log(`   ✅ 提款 ${i + 1} 已完成: ${message}`);
    } else {
      console.log(`   ❌ 提款 ${i + 1} 未完成: ${message}`);
    }
  }

  console.log("🎉 提款操作完成！");
  console.log("📝 注意: 所有符合条件的提款请求已完成，资金已转移到目标地址");
}

// 错误处理
main().catch((error) => {
  console.error("💥 脚本执行失败:");
  console.error(error);
  process.exitCode = 1;
});
