#!/bin/bash

set -e

echo "🚀 开始部署合约..."

# 检查必要的环境变量
check_env_var() {
  if [ -z "${!1}" ]; then
    echo "❌ 错误: 环境变量 $1 未设置"
    exit 1
  fi
}

# 检查网络环境变量
check_env_var "ETH_RPC_URL"
check_env_var "PRIVATE_KEY"

# 设置默认值
export COLD_ADDRESSES=${COLD_ADDRESSES:-"0x0000000000000000000000000000000000000001"}
export HOT_ADDRESSES=${HOT_ADDRESSES:-"0x0000000000000000000000000000000000000002"}
export POWERS=${POWERS:-"1"}
export DISPUTE_PERIOD_SECONDS=${DISPUTE_PERIOD_SECONDS:-"200"}
export BLOCK_DURATION_MILLIS=${BLOCK_DURATION_MILLIS:-"750"}
export LOCKER_THRESHOLD=${LOCKER_THRESHOLD:-"1"}

# 网络选择
NETWORK=${NETWORK:-"bscTestnet"}
#CHAIN_ID=$(echo "console.log(hre.network.config.chainId)" | npx hardhat console --network $NETWORK)

echo "📋 部署配置:"
echo "  网络: $NETWORK"
echo "  RPC URL: $ETH_RPC_URL"

function echo_validator_set() {
    echo ""
    echo "🔧 验证者地址:"
    echo "  Hot: $HOT_ADDRESSES"
    echo "  Cold: $COLD_ADDRESSES"
    echo "  Powers: $POWERS"
}

if [ "$1" == "bridge" ]; then
  echo "  争议期: ${DISPUTE_PERIOD_SECONDS}秒"
  echo "  区块时长: ${BLOCK_DURATION_MILLIS}毫秒"
  echo "  锁定阈值: $LOCKER_THRESHOLD"
  echo_validator_set
  npx hardhat run ./scripts/deploy-bridge.ts --network $NETWORK
fi

if [ "$1" == "bridgeHub" ]; then
  echo_validator_set
  npx hardhat run ./scripts/deploy-bridge-hub.ts --network $NETWORK
fi

if [ "$1" == "bridgeToken" ]; then
  echo "  CHAIN ID: $CHAIN_ID"
  echo "  Token 地址: $TOKEN_ADDRESS"
  echo "  Bridged Token 地址: $BRIDGED_TOKEN_ADDRESS"
  echo "  Bridge Hub 地址: $BRIDGE_HUB_ADDRESS"
  npx hardhat run scripts/set-token-pair.ts --network "$NETWORK"
  export ERC20_ADDRESS=$BRIDGED_TOKEN_ADDRESS
  export MINER_ADDRESS=$BRIDGE_HUB_ADDRESS
  npx hardhat run scripts/set-erc20-miner.ts --network "$NETWORK"
fi

if [ "$1" == "setWithdrawFee" ]; then
  echo "  Bridge Hub 地址: $BRIDGE_HUB_ADDRESS"
  echo "  Bridged Token 地址: $BRIDGED_TOKEN_ADDRESS"
  echo "  Fee: $WITHDRAW_FEE"
  npx hardhat run scripts/set-withdraw-fee.ts --network "$NETWORK"
fi