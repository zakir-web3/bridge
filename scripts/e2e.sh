#!/usr/bin/env bash
# Dual local Hardhat e2e: deploy Bridge + BridgeHub, run the relayer, then
# deposit on the source chain and withdraw back through BridgeHub.
set -euo pipefail

ROOT="$(cd "$(dirname "$0")/.." && pwd)"
SOLIDITY="$ROOT/solidity"
WORK="$ROOT/tmp/e2e"

SOURCE_PORT="${SOURCE_PORT:-8545}"
HUB_PORT="${HUB_PORT:-8546}"
SOURCE_CHAIN_ID="${SOURCE_CHAIN_ID:-31337}"
HUB_CHAIN_ID="${HUB_CHAIN_ID:-1337}"
SOURCE_RPC="http://127.0.0.1:${SOURCE_PORT}"
HUB_RPC="http://127.0.0.1:${HUB_PORT}"
MINE_INTERVAL="${HARDHAT_MINE_INTERVAL:-1000}"

# Hardhat account #0 is the validator/relayer (public test key, local/CI only).
PRIVATE_KEY="${PRIVATE_KEY:-0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80}"
VALIDATOR="${VALIDATOR:-0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266}"
COLD_VALIDATOR="${COLD_VALIDATOR:-0x70997970C51812dc3A010C7d01b50e0d17dc79C8}"
# Hardhat account #2 is the depositor/withdrawer so it does not share nonces with the relayer.
USER_PRIVATE_KEY="${USER_PRIVATE_KEY:-0x5de4111afa1a4b94908f83103eb1f1706367c2e68ca870fc3fb9a804cdab365a}"

DISPUTE_PERIOD_SECONDS="${DISPUTE_PERIOD_SECONDS:-2}"
BLOCK_DURATION_MILLIS="${BLOCK_DURATION_MILLIS:-750}"
AMOUNT="${AMOUNT:-1000000000000000000}"

mkdir -p "$WORK"
rm -rf "$WORK"/*
: >"$WORK/source.log"
: >"$WORK/hub.log"
: >"$WORK/relayer.log"

SOURCE_PID=""
HUB_PID=""
RELAYER_PID=""

cleanup() {
  local code=$?
  if [[ -n "${RELAYER_PID}" ]] && kill -0 "$RELAYER_PID" 2>/dev/null; then
    kill "$RELAYER_PID" 2>/dev/null || true
    wait "$RELAYER_PID" 2>/dev/null || true
  fi
  if [[ -n "${HUB_PID}" ]] && kill -0 "$HUB_PID" 2>/dev/null; then
    kill "$HUB_PID" 2>/dev/null || true
    wait "$HUB_PID" 2>/dev/null || true
  fi
  if [[ -n "${SOURCE_PID}" ]] && kill -0 "$SOURCE_PID" 2>/dev/null; then
    kill "$SOURCE_PID" 2>/dev/null || true
    wait "$SOURCE_PID" 2>/dev/null || true
  fi
  if [[ $code -ne 0 ]]; then
    echo "---- source hardhat (keys stripped) ----"
    grep -vi "private key" "$WORK/source.log" | tail -n 40 || true
    echo "---- hub hardhat (keys stripped) ----"
    grep -vi "private key" "$WORK/hub.log" | tail -n 40 || true
    echo "---- relayer ----"
    tail -n 120 "$WORK/relayer.log" || true
  fi
}
trap cleanup EXIT

port_in_use() {
  local port="$1"
  if command -v lsof >/dev/null 2>&1; then
    lsof -iTCP:"$port" -sTCP:LISTEN >/dev/null 2>&1
  else
    (echo >/dev/tcp/127.0.0.1/"$port") >/dev/null 2>&1
  fi
}

wait_for_rpc() {
  local url="$1"
  local chain_hex="$2"
  local i
  for i in $(seq 1 60); do
    if curl -sf -X POST "$url" \
      -H 'content-type: application/json' \
      --data '{"jsonrpc":"2.0","id":1,"method":"eth_chainId","params":[]}' \
      | grep -qi "$chain_hex"; then
      return 0
    fi
    sleep 1
  done
  echo "RPC $url did not become ready (want chain id $chain_hex)" >&2
  return 1
}

extract_address() {
  local file="$1"
  local addr
  addr="$(grep -E '^DEPLOYED_ADDRESS=' "$file" | tail -n 1 | cut -d= -f2 | tr -d '[:space:]')"
  if [[ -z "$addr" ]]; then
    echo "failed to parse DEPLOYED_ADDRESS from $file" >&2
    cat "$file" >&2
    return 1
  fi
  echo "$addr"
}

to_chain_hex() {
  printf '0x%x' "$1"
}

address_from_key() {
  local key="$1"
  (cd "$SOLIDITY" && NODE_PATH="$SOLIDITY/node_modules" PRIVATE_KEY="$key" node -e '
    const { ethers } = require("ethers");
    console.log(new ethers.Wallet(process.env.PRIVATE_KEY).address);
  ')
}

token_balance() {
  local rpc="$1"
  local token="$2"
  local account="$3"
  (cd "$SOLIDITY" && NODE_PATH="$SOLIDITY/node_modules" \
    RPC="$rpc" TOKEN="$token" ACCOUNT="$account" node -e '
      const { ethers } = require("ethers");
      (async () => {
        const provider = new ethers.JsonRpcProvider(process.env.RPC);
        const erc20 = new ethers.Contract(
          process.env.TOKEN,
          ["function balanceOf(address) view returns (uint256)"],
          provider
        );
        console.log((await erc20.balanceOf(process.env.ACCOUNT)).toString());
      })().catch((err) => {
        console.error(err);
        process.exit(1);
      });
    ')
}

wait_for_balance() {
  local label="$1"
  local rpc="$2"
  local token="$3"
  local account="$4"
  local expected="$5"
  local timeout_secs="$6"
  local start now bal
  start="$(date +%s)"
  while true; do
    bal="$(token_balance "$rpc" "$token" "$account")"
    echo "[e2e] ${label}: balance=${bal} expected=${expected}"
    if [[ "$bal" == "$expected" ]]; then
      return 0
    fi
    now="$(date +%s)"
    if (( now - start >= timeout_secs )); then
      echo "timeout waiting for ${label}" >&2
      return 1
    fi
    sleep 1
  done
}

if port_in_use "$SOURCE_PORT"; then
  echo "port $SOURCE_PORT is already in use" >&2
  exit 1
fi
if port_in_use "$HUB_PORT"; then
  echo "port $HUB_PORT is already in use" >&2
  exit 1
fi

if [[ ! -d "$SOLIDITY/node_modules" ]]; then
  (cd "$SOLIDITY" && npm ci)
fi

echo "==> compile contracts"
(cd "$SOLIDITY" && npm run compile)

USER_ADDRESS="$(address_from_key "$USER_PRIVATE_KEY")"
echo "==> user ${USER_ADDRESS}"

echo "==> build relayer"
make -C "$ROOT" build

echo "==> start source Hardhat (chain $SOURCE_CHAIN_ID port $SOURCE_PORT)"
(
  cd "$SOLIDITY"
  HARDHAT_CHAIN_ID="$SOURCE_CHAIN_ID" HARDHAT_MINE_INTERVAL="$MINE_INTERVAL" \
    npx hardhat node --hostname 127.0.0.1 --port "$SOURCE_PORT"
) >"$WORK/source.log" 2>&1 &
SOURCE_PID=$!
wait_for_rpc "$SOURCE_RPC" "$(to_chain_hex "$SOURCE_CHAIN_ID")"

echo "==> start hub Hardhat (chain $HUB_CHAIN_ID port $HUB_PORT)"
(
  cd "$SOLIDITY"
  HARDHAT_CHAIN_ID="$HUB_CHAIN_ID" HARDHAT_MINE_INTERVAL="$MINE_INTERVAL" \
    npx hardhat node --hostname 127.0.0.1 --port "$HUB_PORT"
) >"$WORK/hub.log" 2>&1 &
HUB_PID=$!
wait_for_rpc "$HUB_RPC" "$(to_chain_hex "$HUB_CHAIN_ID")"
echo "==> both nodes ready"

export PRIVATE_KEY
export HOT_ADDRESSES="$VALIDATOR"
export COLD_ADDRESSES="$COLD_VALIDATOR"
export POWERS="1"
export DISPUTE_PERIOD_SECONDS
export BLOCK_DURATION_MILLIS
export LOCKER_THRESHOLD=1
export NETWORK=custom
export TOKEN_DECIMAL=18

deploy() {
  local rpc="$1"
  local chain_id="$2"
  local target="$3"
  local out="$4"
  (
    cd "$SOLIDITY"
    ETH_RPC_URL="$rpc" CHAIN_ID="$chain_id" bash ./deploy.sh "$target"
  ) | tee "$out"
}

echo "==> deploy source token"
(
  cd "$SOLIDITY"
  TOKEN_NAME="Source USD" TOKEN_SYMBOL="sUSD" \
    MINT_TO="$USER_ADDRESS" MINT_AMOUNT="$AMOUNT" \
    ETH_RPC_URL="$SOURCE_RPC" CHAIN_ID="$SOURCE_CHAIN_ID" \
    npx hardhat run scripts/deploy-bridge-token.ts --network custom
) | tee "$WORK/source-token.log"
SOURCE_TOKEN="$(extract_address "$WORK/source-token.log")"

echo "==> deploy dest token"
(
  cd "$SOLIDITY"
  TOKEN_NAME="Bridged USD" TOKEN_SYMBOL="bUSD" \
    ETH_RPC_URL="$HUB_RPC" CHAIN_ID="$HUB_CHAIN_ID" \
    npx hardhat run scripts/deploy-bridge-token.ts --network custom
) | tee "$WORK/dest-token.log"
DEST_TOKEN="$(extract_address "$WORK/dest-token.log")"

echo "==> deploy Bridge"
deploy "$SOURCE_RPC" "$SOURCE_CHAIN_ID" bridge "$WORK/bridge.log"
BRIDGE_ADDRESS="$(extract_address "$WORK/bridge.log")"

echo "==> deploy BridgeHub"
deploy "$HUB_RPC" "$HUB_CHAIN_ID" bridgeHub "$WORK/bridgehub.log"
BRIDGE_HUB_ADDRESS="$(extract_address "$WORK/bridgehub.log")"

echo "==> register token pair"
(
  cd "$SOLIDITY"
  TOKEN_ADDRESS="$SOURCE_TOKEN" \
    BRIDGED_TOKEN_ADDRESS="$DEST_TOKEN" \
    BRIDGE_HUB_ADDRESS="$BRIDGE_HUB_ADDRESS" \
    ETH_RPC_URL="$HUB_RPC" \
    CHAIN_ID="$HUB_CHAIN_ID" \
    TOKEN_CHAIN_ID="$SOURCE_CHAIN_ID" \
    bash ./deploy.sh bridgeToken
) | tee "$WORK/pair.log"

echo "==> write relayer config"
cat >"$WORK/config.toml" <<EOF
log_level = "info"
log_format = "console"
source = "./bridge.db"
priv_key = "${PRIVATE_KEY}"

[bridge_hub]
node_url = "${HUB_RPC}"
chain_id = "${HUB_CHAIN_ID}"
max_retries = 5
base_delay = "100ms"
max_delay = "5s"
backoff_rate = 2.0
interval = "1s"
start_block = 1
block_interval = 100
block_delay = 0
clear_cache = false
bridge_hub_address = "${BRIDGE_HUB_ADDRESS}"
fee_history_block_count = 10
fee_history_reward_percentiles = [50.0, 65.0, 80.0]
max_gas_tip_cap = "1000000000"
max_gas_fee_cap = "5000000000"
increase_percentile = "120"
max_gas_limit = 5000000
no_send = false
gas_tip_cap = "1000000000"
gas_fee_cap = "1000000000"

[bridge]
node_url = "${SOURCE_RPC}"
chain_id = "${SOURCE_CHAIN_ID}"
max_retries = 5
base_delay = "100ms"
max_delay = "5s"
backoff_rate = 2.0
interval = "1s"
start_block = 1
block_interval = 100
block_delay = 0
clear_cache = false
bridge_address = "${BRIDGE_ADDRESS}"
bridge_tokens = [
    "${SOURCE_TOKEN}",
]
fee_history_block_count = 10
fee_history_reward_percentiles = [50.0, 65.0, 80.0]
max_gas_tip_cap = "1000000000"
max_gas_fee_cap = "5000000000"
increase_percentile = "120"
max_gas_limit = 5000000
no_send = false
gas_tip_cap = "1000000000"
gas_fee_cap = "1000000000"
send_finalize_withdrawals = true
finalize_withdrawals_interval = "2s"
max_withdraw_amount = "100000000000000000000"
request_withdraw_gas_limit = 0
enable_request_withdraw = true
EOF

echo "==> start relayer"
(
  cd "$WORK"
  ENABLE_FINALIZE_WITHDRAWALS=true exec "$ROOT/bin/bridge"
) >"$WORK/relayer.log" 2>&1 &
RELAYER_PID=$!

ready=0
for _ in $(seq 1 30); do
  if grep -q "All services started successfully" "$WORK/relayer.log"; then
    ready=1
    break
  fi
  if ! kill -0 "$RELAYER_PID" 2>/dev/null; then
    echo "relayer exited before becoming ready" >&2
    exit 1
  fi
  sleep 1
done
if [[ "$ready" -ne 1 ]]; then
  echo "relayer did not start in time" >&2
  exit 1
fi
echo "==> relayer ready"

echo "==> deposit"
(
  cd "$SOLIDITY"
  PRIVATE_KEY="$USER_PRIVATE_KEY" \
    ETH_RPC_URL="$SOURCE_RPC" \
    CHAIN_ID="$SOURCE_CHAIN_ID" \
    BRIDGE_CONTRACT_ADDRESS="$BRIDGE_ADDRESS" \
    TOKEN_ADDRESS="$SOURCE_TOKEN" \
    DESTINATION_ADDRESS="$USER_ADDRESS" \
    AMOUNT="$AMOUNT" \
    npx hardhat run scripts/deposit.ts --network custom
)
wait_for_balance "dest mint after deposit" "$HUB_RPC" "$DEST_TOKEN" "$USER_ADDRESS" "$AMOUNT" 90

source_locked="$(token_balance "$SOURCE_RPC" "$SOURCE_TOKEN" "$USER_ADDRESS")"
if [[ "$source_locked" != "0" ]]; then
  echo "source token should be locked in Bridge, got ${source_locked}" >&2
  exit 1
fi

echo "==> withdraw"
(
  cd "$SOLIDITY"
  PRIVATE_KEY="$USER_PRIVATE_KEY" \
    ETH_RPC_URL="$HUB_RPC" \
    CHAIN_ID="$HUB_CHAIN_ID" \
    TOKEN_CHAIN_ID="$SOURCE_CHAIN_ID" \
    BRIDGE_HUB_CONTRACT_ADDRESS="$BRIDGE_HUB_ADDRESS" \
    TOKEN_ADDRESS="$DEST_TOKEN" \
    DESTINATION_ADDRESS="$USER_ADDRESS" \
    AMOUNT="$AMOUNT" \
    npx hardhat run scripts/withdraw.ts --network custom
)
wait_for_balance "dest burn after withdraw" "$HUB_RPC" "$DEST_TOKEN" "$USER_ADDRESS" "0" 90
wait_for_balance "source unlock after finalize" "$SOURCE_RPC" "$SOURCE_TOKEN" "$USER_ADDRESS" "$AMOUNT" 120

echo "==> e2e passed"
echo "Bridge=${BRIDGE_ADDRESS}"
echo "BridgeHub=${BRIDGE_HUB_ADDRESS}"
echo "SourceToken=${SOURCE_TOKEN}"
echo "DestToken=${DEST_TOKEN}"
