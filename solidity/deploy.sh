#!/bin/bash

set -e

echo "🚀 Starting contract deployment..."

# Check required environment variables
check_env_var() {
  if [ -z "${!1}" ]; then
    echo "❌ Error: environment variable $1 is not set"
    exit 1
  fi
}

# Check network environment variables
check_env_var "ETH_RPC_URL"
check_env_var "PRIVATE_KEY"

# Set defaults
export COLD_ADDRESSES=${COLD_ADDRESSES:-"0x0000000000000000000000000000000000000001"}
export HOT_ADDRESSES=${HOT_ADDRESSES:-"0x0000000000000000000000000000000000000002"}
export POWERS=${POWERS:-"1"}
export DISPUTE_PERIOD_SECONDS=${DISPUTE_PERIOD_SECONDS:-"200"}
export BLOCK_DURATION_MILLIS=${BLOCK_DURATION_MILLIS:-"750"}
export LOCKER_THRESHOLD=${LOCKER_THRESHOLD:-"1"}

# Network selection
NETWORK=${NETWORK:-"bscTestnet"}
#CHAIN_ID=$(echo "console.log(hre.network.config.chainId)" | npx hardhat console --network $NETWORK)

echo "📋 Deployment config:"
echo "  Network: $NETWORK"
echo "  RPC URL: $ETH_RPC_URL"

function echo_validator_set() {
    echo ""
    echo "🔧 Validator addresses:"
    echo "  Hot: $HOT_ADDRESSES"
    echo "  Cold: $COLD_ADDRESSES"
    echo "  Powers: $POWERS"
}

if [ "$1" == "bridge" ]; then
  echo "  Dispute period: ${DISPUTE_PERIOD_SECONDS}s"
  echo "  Block duration: ${BLOCK_DURATION_MILLIS}ms"
  echo "  Locker threshold: $LOCKER_THRESHOLD"
  echo_validator_set
  npx hardhat run ./scripts/deploy-bridge.ts --network $NETWORK
fi

if [ "$1" == "bridgeHub" ]; then
  echo_validator_set
  npx hardhat run ./scripts/deploy-bridge-hub.ts --network $NETWORK
fi

if [ "$1" == "bridgeToken" ]; then
  echo "  CHAIN ID: $CHAIN_ID"
  echo "  TOKEN_CHAIN_ID: ${TOKEN_CHAIN_ID:-$CHAIN_ID}"
  echo "  Token address: $TOKEN_ADDRESS"
  echo "  Bridged token address: $BRIDGED_TOKEN_ADDRESS"
  echo "  Bridge Hub address: $BRIDGE_HUB_ADDRESS"
  npx hardhat run scripts/set-token-pair.ts --network "$NETWORK"
  export ERC20_ADDRESS=$BRIDGED_TOKEN_ADDRESS
  export MINER_ADDRESS=$BRIDGE_HUB_ADDRESS
  npx hardhat run scripts/set-erc20-miner.ts --network "$NETWORK"
fi

if [ "$1" == "setWithdrawFee" ]; then
  echo "  Bridge Hub address: $BRIDGE_HUB_ADDRESS"
  echo "  Bridged token address: $BRIDGED_TOKEN_ADDRESS"
  echo "  Fee: $WITHDRAW_FEE"
  npx hardhat run scripts/set-withdraw-fee.ts --network "$NETWORK"
fi
