#!/usr/bin/env bash
# Solana → EVM deposit end-to-end test (local only).
#
# Flow:
#   1. Start solana-test-validator and a Hardhat hub node
#   2. Deploy Solana bridge program + SPL mint, BridgeHub + bridged ERC-20 on hub
#   3. Start relayer (Solana scanner + hub minter)
#   4. Deposit SPL on Solana → relayer mints bridged ERC-20 on hub
#
# Usage: ./scripts/solana-e2e.sh
# Requires: anchor, solana CLI, solana-test-validator, node, npm, make, curl.
set -euo pipefail

ROOT="$(cd "$(dirname "$0")/.." && pwd)"
# shellcheck source=scripts/e2e-common.sh
source "$ROOT/scripts/e2e-common.sh"
SOLIDITY="$ROOT/solidity"
SOLANA="$ROOT/solana"
WORK="$ROOT/tmp/solana-e2e"

# --- Config (all overridable via env) ---
HUB_PORT="${HUB_PORT:-8546}"
SOLANA_RPC_PORT="${SOLANA_RPC_PORT:-8899}"
HUB_RPC="http://127.0.0.1:${HUB_PORT}"
SOLANA_RPC="http://127.0.0.1:${SOLANA_RPC_PORT}"
HUB_CHAIN_ID="${HUB_CHAIN_ID:-1337}"
# Must match the chain id registered in BridgeHub for the Solana source token.
SOLANA_CHAIN_ID="${SOLANA_CHAIN_ID:-900001}"
MINE_INTERVAL="${HARDHAT_MINE_INTERVAL:-1000}"

# Hardhat account #0 is the validator/relayer (public test key, local/CI only).
PRIVATE_KEY="${PRIVATE_KEY:-0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80}"
VALIDATOR="${VALIDATOR:-0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266}"
COLD_VALIDATOR="${COLD_VALIDATOR:-0x70997970C51812dc3A010C7d01b50e0d17dc79C8}"
# Hardhat account #2 receives the bridged ERC-20 on the hub chain.
USER_PRIVATE_KEY="${USER_PRIVATE_KEY:-0x5de4111afa1a4b94908f83103eb1f1706367c2e68ca870fc3fb9a804cdab365a}"

# Short dispute window (only used for BridgeHub deploy params in this script).
DISPUTE_PERIOD_SECONDS="${DISPUTE_PERIOD_SECONDS:-2}"
BLOCK_DURATION_MILLIS="${BLOCK_DURATION_MILLIS:-750}"
DEPOSIT_AMOUNT="${DEPOSIT_AMOUNT:-1000000}" # 1 SPL token with 6 decimals

export HARDHAT_DISABLE_TELEMETRY_PROMPT="${HARDHAT_DISABLE_TELEMETRY_PROMPT:-true}"

# --- Workspace & cleanup ---
mkdir -p "$WORK"
rm -rf "$WORK"/*
: >"$WORK/hub.log"
: >"$WORK/validator.log"
: >"$WORK/relayer.log"

VALIDATOR_PID=""
HUB_PID=""
RELAYER_PID=""

cleanup() {
  local code=$?
  e2e_stop_pid "$RELAYER_PID"
  e2e_stop_pid "$HUB_PID"
  e2e_stop_pid "$VALIDATOR_PID"
  e2e_free_port "$HUB_PORT"
  e2e_free_port "$SOLANA_RPC_PORT"
  if [[ $code -ne 0 ]]; then
    echo "---- solana-test-validator ----"
    tail -n 60 "$WORK/validator.log" || true
    echo "---- hub hardhat (keys stripped) ----"
    grep -vi "private key" "$WORK/hub.log" | tail -n 40 || true
    echo "---- relayer ----"
    tail -n 120 "$WORK/relayer.log" || true
  fi
}
trap cleanup EXIT

# --- Helpers ---
port_in_use() {
  e2e_port_in_use "$1"
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

wait_for_solana() {
  local i
  for i in $(seq 1 60); do
    if solana cluster-version --url "$SOLANA_RPC" >/dev/null 2>&1; then
      return 0
    fi
    sleep 1
  done
  echo "Solana RPC $SOLANA_RPC did not become ready" >&2
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

extract_kv() {
  local file="$1"
  local key="$2"
  local val
  val="$(grep -E "^${key}=" "$file" | tail -n 1 | cut -d= -f2- | tr -d '[:space:]')"
  if [[ -z "$val" ]]; then
    echo "failed to parse ${key} from $file" >&2
    cat "$file" >&2
    return 1
  fi
  echo "$val"
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
    echo "[solana-e2e] ${label}: balance=${bal} expected=${expected}"
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

require_cmd() {
  if ! command -v "$1" >/dev/null 2>&1; then
    echo "missing required command: $1" >&2
    exit 1
  fi
}

# --- Preflight ---
require_cmd anchor
require_cmd solana
require_cmd solana-test-validator
require_cmd solana-keygen

if port_in_use "$HUB_PORT"; then
  echo "port $HUB_PORT is already in use" >&2
  exit 1
fi
if port_in_use "$SOLANA_RPC_PORT"; then
  echo "port $SOLANA_RPC_PORT is already in use" >&2
  exit 1
fi

if [[ ! -d "$SOLIDITY/node_modules" ]]; then
  (cd "$SOLIDITY" && npm ci)
fi
if [[ ! -d "$SOLANA/node_modules" ]]; then
  (cd "$SOLANA" && npm ci)
fi

USER_ADDRESS="$(address_from_key "$USER_PRIVATE_KEY")"
echo "==> user ${USER_ADDRESS}"

echo "==> compile EVM contracts"
(cd "$SOLIDITY" && npm run compile)

echo "==> build relayer"
make -C "$ROOT" build

# --- Solana local chain ---
echo "==> start solana-test-validator (port $SOLANA_RPC_PORT)"
solana-test-validator \
  --ledger "$WORK/validator-ledger" \
  --reset \
  --quiet \
  --rpc-port "$SOLANA_RPC_PORT" >"$WORK/validator.log" 2>&1 &
VALIDATOR_PID=$!
wait_for_solana

# Fresh keypair for this run; funded via airdrop on the local validator.
SOLANA_WALLET="$WORK/solana-wallet.json"
solana-keygen new -o "$SOLANA_WALLET" --no-bip39-passphrase --force >/dev/null
SOLANA_PUBKEY="$(solana address -k "$SOLANA_WALLET")"
solana config set --url "$SOLANA_RPC" --keypair "$SOLANA_WALLET" >/dev/null

for _ in $(seq 1 10); do
  if solana airdrop 100 "$SOLANA_PUBKEY" --url "$SOLANA_RPC" >/dev/null 2>&1; then
    break
  fi
  sleep 1
done

SOL_BALANCE="$(solana balance "$SOLANA_PUBKEY" --url "$SOLANA_RPC" 2>/dev/null | awk '{print $1}')"
if [[ -z "$SOL_BALANCE" ]] || awk "BEGIN { exit !($SOL_BALANCE < 1) }"; then
  echo "failed to fund Solana wallet ${SOLANA_PUBKEY} (balance=${SOL_BALANCE:-0})" >&2
  exit 1
fi
echo "==> solana wallet ${SOLANA_PUBKEY} balance ${SOL_BALANCE} SOL"

export ANCHOR_PROVIDER_URL="$SOLANA_RPC"
export ANCHOR_WALLET="$SOLANA_WALLET"

# Deploy the Anchor program and create SPL mint + program vault (setup-e2e.ts).
echo "==> build and deploy Solana program"
(
  cd "$SOLANA"
  export CARGO_TARGET_DIR="$SOLANA/target"
  rm -f target/deploy/bridge-upgrade-buffer.json
  anchor keys sync
  anchor build
  solana program deploy target/deploy/bridge.so \
    --program-id target/deploy/bridge-keypair.json \
    --url "$SOLANA_RPC" \
    --keypair "$SOLANA_WALLET"
) | tee "$WORK/solana-deploy.log"

PROGRAM_ID="$(solana address -k "$SOLANA/target/deploy/bridge-keypair.json")"
if [[ -z "$PROGRAM_ID" ]]; then
  echo "failed to read program id from keypair" >&2
  exit 1
fi
echo "==> program id ${PROGRAM_ID}"

echo "==> setup SPL mint and vault"
(
  cd "$SOLANA"
  export ANCHOR_PROVIDER_URL="$SOLANA_RPC"
  export ANCHOR_WALLET="$SOLANA_WALLET"
  npx ts-node scripts/setup-e2e.ts
) | tee "$WORK/solana-setup.log"

MINT="$(extract_kv "$WORK/solana-setup.log" "MINT")"
SRC_TOKEN_BYTES32="$(extract_kv "$WORK/solana-setup.log" "SRC_TOKEN_BYTES32")"
TOKEN_DECIMAL="$(extract_kv "$WORK/solana-setup.log" "TOKEN_DECIMAL")"
echo "==> mint ${MINT}"

# --- EVM hub chain ---
echo "==> start hub Hardhat (chain $HUB_CHAIN_ID port $HUB_PORT)"
(
  cd "$SOLIDITY"
  HARDHAT_CHAIN_ID="$HUB_CHAIN_ID" HARDHAT_MINE_INTERVAL="$MINE_INTERVAL" \
    exec npx hardhat node --hostname 127.0.0.1 --port "$HUB_PORT"
) >"$WORK/hub.log" 2>&1 &
HUB_PID=$!
wait_for_rpc "$HUB_RPC" "$(to_chain_hex "$HUB_CHAIN_ID")"

echo "==> deploy bridged ERC-20 on hub"
(
  cd "$SOLIDITY"
  TOKEN_NAME="Bridged SPL" TOKEN_SYMBOL="bSPL" \
    ETH_RPC_URL="$HUB_RPC" CHAIN_ID="$HUB_CHAIN_ID" \
    npx hardhat run scripts/deploy-bridge-token.ts --network custom
) | tee "$WORK/dest-token.log"
DEST_TOKEN="$(extract_address "$WORK/dest-token.log")"

echo "==> deploy BridgeHub"
(
  cd "$SOLIDITY"
  export PRIVATE_KEY
  export HOT_ADDRESSES="$VALIDATOR"
  export COLD_ADDRESSES="$COLD_VALIDATOR"
  export POWERS="1"
  export DISPUTE_PERIOD_SECONDS
  export BLOCK_DURATION_MILLIS
  export LOCKER_THRESHOLD=1
  export NETWORK=custom
  ETH_RPC_URL="$HUB_RPC" CHAIN_ID="$HUB_CHAIN_ID" bash ./deploy.sh bridgeHub
) | tee "$WORK/bridgehub.log"
BRIDGE_HUB_ADDRESS="$(extract_address "$WORK/bridgehub.log")"

# Solana source tokens are identified by bytes32 mint pubkey, not an EVM address.
echo "==> register Solana token pair"
(
  cd "$SOLIDITY"
  export PRIVATE_KEY
  export NETWORK=custom
  PAIR_MODE=solana \
    SRC_TOKEN_BYTES32="$SRC_TOKEN_BYTES32" \
    TOKEN_ADDRESS="$SRC_TOKEN_BYTES32" \
    BRIDGED_TOKEN_ADDRESS="$DEST_TOKEN" \
    BRIDGE_HUB_ADDRESS="$BRIDGE_HUB_ADDRESS" \
    ETH_RPC_URL="$HUB_RPC" \
    CHAIN_ID="$HUB_CHAIN_ID" \
    TOKEN_CHAIN_ID="$SOLANA_CHAIN_ID" \
    TOKEN_DECIMAL="$TOKEN_DECIMAL" \
    bash ./deploy.sh bridgeToken
) | tee "$WORK/pair.log"

# Scale SPL amount (6 dec) to hub ERC-20 amount (18 dec) for balance assertion.
EXPECTED_HUB_BALANCE="$(
  cd "$SOLIDITY" && NODE_PATH="$SOLIDITY/node_modules" \
    AMOUNT="$DEPOSIT_AMOUNT" TOKEN_DECIMAL="$TOKEN_DECIMAL" node -e '
      const amount = BigInt(process.env.AMOUNT);
      const srcDec = BigInt(process.env.TOKEN_DECIMAL);
      const destDec = 18n;
      const scale = 10n ** (destDec - srcDec);
      console.log((amount * scale).toString());
    '
)"

# [bridge] is a no-op placeholder here; Solana deposits are handled via [solana].
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
bridge_address = ""
bridge_tokens = []
fee_history_block_count = 10
fee_history_reward_percentiles = [50.0, 65.0, 80.0]
max_gas_tip_cap = "1000000000"
max_gas_fee_cap = "5000000000"
increase_percentile = "120"
max_gas_limit = 5000000
no_send = false
gas_tip_cap = "1000000000"
gas_fee_cap = "1000000000"
send_finalize_withdrawals = false
finalize_withdrawals_interval = "30s"
max_withdraw_amount = "100000000000000000000"
request_withdraw_gas_limit = 0
enable_request_withdraw = false

[solana]
node_url = "${SOLANA_RPC}"
chain_id = "${SOLANA_CHAIN_ID}"
program_id = "${PROGRAM_ID}"
bridge_mints = [
    "${MINT}",
]
max_retries = 5
base_delay = "100ms"
max_delay = "5s"
backoff_rate = 2.0
interval = "500ms"
start_slot = 0
slot_interval = 1000
slot_delay = 0
clear_cache = false
EOF

echo "==> start relayer"
(
  cd "$WORK"
  exec "$ROOT/bin/bridge"
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

# --- Solana deposit → hub mint ---
echo "==> deposit on Solana"
(
  cd "$SOLANA"
  export ANCHOR_PROVIDER_URL="$SOLANA_RPC"
  export ANCHOR_WALLET="$SOLANA_WALLET"
  MINT="$MINT" AMOUNT="$DEPOSIT_AMOUNT" DESTINATION="$USER_ADDRESS" \
    npx ts-node scripts/deposit.ts
) | tee "$WORK/deposit.log"

# Give the relayer a moment to observe the slot before polling hub balance.
sleep 2

wait_for_balance "hub mint after Solana deposit" "$HUB_RPC" "$DEST_TOKEN" "$USER_ADDRESS" "$EXPECTED_HUB_BALANCE" 120

echo "==> solana e2e passed"
echo "ProgramID=${PROGRAM_ID}"
echo "Mint=${MINT}"
echo "BridgeHub=${BRIDGE_HUB_ADDRESS}"
echo "DestToken=${DEST_TOKEN}"
echo "User=${USER_ADDRESS}"
