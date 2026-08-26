# Contracts

Deploy `Bridge` on the **source** chain (lock / deposit) and `BridgeHub` on the **custom** chain (mint / withdraw). Use `deploy.sh` with Hardhat. Any EVM network with a stable RPC and a known chain ID works; `bsc` and `bscTestnet` are convenience presets.

[中文](README.zh-CN.md) · [Root README](../README.md)

## Prerequisites

- Node.js ≥ 22.18.0, npm ≥ 10.9.3
- `jq` (the scripts parse JSON addresses)
- A deployer key with enough gas on the target network
- A stable RPC URL

## Networks

Set `NETWORK` to a name from `hardhat.config.ts`:

| `NETWORK` | Meaning |
|-----------|---------|
| `custom` | Any EVM chain. Set `ETH_RPC_URL` and `CHAIN_ID` (default `1337`) |
| `bsc` | BSC mainnet preset (`chainId = 56`) |
| `bscTestnet` | BSC testnet preset (`chainId = 97`) |
| `localhost` | Local Hardhat / Anvil node |

`ETH_RPC_URL` is the RPC for whichever network you are deploying to, not Ethereum-only.

## Environment variables

Required for both contracts (must match on both chains):

- `HOT_ADDRESSES` — comma-separated validator addresses
- `COLD_ADDRESSES` — comma-separated cold validator addresses (may equal hot)
- `POWERS` — comma-separated integer weights, one per address (for example `1,1,2`)

Deployer:

- `PRIVATE_KEY` — `0x`-prefixed deployer key
- `ETH_RPC_URL` — RPC for `NETWORK`
- `NETWORK` — see the table above
- `CHAIN_ID` — required for `custom`, and when registering a token pair (source-chain ID)

Token pair (when calling `bridgeToken`):

- `TOKEN_ADDRESS` — ERC-20 on the **source** chain
- `BRIDGED_TOKEN_ADDRESS` — corresponding token on the **custom** chain
- `BRIDGE_HUB_ADDRESS` — deployed `BridgeHub`

`Bridge` parameters (optional, with defaults):

- `DISPUTE_PERIOD_SECONDS` (default `200`)
- `BLOCK_DURATION_MILLIS` (default `750`)
- `LOCKER_THRESHOLD` (default `1`)

## Install and compile

```bash
cd solidity
npm install
npm run compile
```

## Configure validators

```bash
export HOT_ADDRESSES=0xHot1,0xHot2
export COLD_ADDRESSES=0xCold1,0xCold2
export POWERS=1,1
```

Hot, cold, and power lists must be the same length.

## Deploy Bridge (source chain)

Example using the BSC preset. For any other EVM chain, use `NETWORK=custom` and set `CHAIN_ID`.

```bash
export NETWORK=bsc
export PRIVATE_KEY=0xYOUR_KEY
export ETH_RPC_URL=https://your-source-rpc

bash ./deploy.sh bridge
```

The script prints the `Bridge` address.

Test token (local / testnet only):

```bash
export NETWORK=bsc
export PRIVATE_KEY=0xYOUR_KEY

npx hardhat --network "$NETWORK" run scripts/deploy-bridge-token.ts
```

## Deploy BridgeHub (custom chain)

```bash
export NETWORK=custom
export CHAIN_ID=1337
export PRIVATE_KEY=0xYOUR_KEY
export ETH_RPC_URL=https://your-custom-rpc

bash ./deploy.sh bridgeHub
```

The script prints the `BridgeHub` address.

Test token (local / testnet only):

```bash
export NETWORK=custom
export PRIVATE_KEY=0xYOUR_KEY

npx hardhat --network "$NETWORK" run scripts/deploy-bridge-token.ts
```

## Register a token pair and grant mint/burn

`CHAIN_ID` here is the **source** chain ID the token originates from.

```bash
export NETWORK=custom
export PRIVATE_KEY=0xYOUR_KEY
export CHAIN_ID=56
export TOKEN_ADDRESS=0xSourceChainToken
export BRIDGED_TOKEN_ADDRESS=0xCustomChainToken
export BRIDGE_HUB_ADDRESS=0xBridgeHub

bash ./deploy.sh bridgeToken
```

## Withdrawal fee

Unset means free. Fees are per custom-chain token.

```bash
export NETWORK=custom
export PRIVATE_KEY=0xYOUR_KEY
export BRIDGE_HUB_ADDRESS=0xBridgeHub
export BRIDGED_TOKEN_ADDRESS=0xCustomChainToken
export WITHDRAW_FEE=1000000000000000000

bash ./deploy.sh setWithdrawFee
```

## Troubleshooting

- Validator address and power list lengths must match or the script exits.
- Confirm the deployer balance and that `ETH_RPC_URL` points at the network in `NETWORK`.
- Missing `jq` causes address parsing to fail (`brew install jq` on macOS).
