# Bridge

[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)
[![Release](https://img.shields.io/github/v/release/zakir-web3/bridge)](https://github.com/zakir-web3/bridge/releases)

Self-hosted **lock-and-mint** bridge for ERC-20 tokens between EVM chains. Users lock on a source-chain `Bridge`; a Go relayer collects EIP-712 signatures from your validator set and mints on a destination-chain `BridgeHub`.

This is for appchain and token-issuer teams who operate their own validators. It is not a general-purpose messaging protocol (LayerZero, Hyperlane) and not a retail swap router.

The source-chain `Bridge` contract is adapted from [Hyperliquid Bridge2.sol](https://github.com/hyperliquid-dex/contracts/blob/audit2/Bridge2.sol).

[中文文档](README.zh-CN.md) · [Docs](docs/README.md) · [Contributing](CONTRIBUTING.md)

## Overview

Bridge connects two EVM chains through a pair of contracts:

- **`Bridge`** — deployed on the **source chain** where users lock or deposit tokens
- **`BridgeHub`** — deployed on the **destination chain** where bridged tokens are minted and withdrawals are initiated

The relayer runs block scanners on both sides, collects validator signatures, and relays confirmations across chains. Any EVM network with a stable RPC endpoint and a configured chain ID can be used — Ethereum, BSC, Polygon, Arbitrum, Base, or a custom chain.

```
  Source chain (EVM)                    Destination chain (EVM)
 ┌─────────────────────┐               ┌─────────────────────┐
 │      Bridge.sol     │               │   BridgeHub.sol     │
 │  lock / deposit     │               │  mint / withdraw    │
 └──────────┬──────────┘               └──────────┬──────────┘
            │                                     │
            └──────────►  Relayer  ◄──────────────┘
                    (scan · sign · submit)
```

## Features

- **EVM-agnostic** — works with any EVM-compatible chain; configure RPC URLs and chain IDs in `config.toml`
- **ERC-20 support** — bridge whitelisted tokens via lock-and-mint on deposit, burn-and-release on withdrawal
- **Validator quorum** — cross-chain messages are authorized with EIP-712 signatures from a configured validator set
- **Dispute period** — withdrawals on the source chain enter a pending state before finalization
- **Resilient scanning** — block progress is persisted in BadgerDB with configurable confirmation depth and retry logic
- **Flexible gas control** — fee history, gas caps, and a dry-run (`no_send`) mode for testing

## Prerequisites

| Tool | Version |
|------|---------|
| Go | 1.25+ |
| Node.js | 22.18.0+ |
| npm | 10.9.3+ |

## Quick Start

### 1. Clone and build

```bash
git clone https://github.com/zakir-web3/bridge.git
cd bridge
make build
```

### 2. Deploy contracts

Deploy `Bridge` on the source chain and `BridgeHub` on the destination chain. See [solidity/README.md](solidity/README.md) for the full deployment guide.

### 3. Configure the relayer

Copy the template and fill in your values:

```bash
cp .config.toml config.toml
```

Key settings:

| Section | Purpose |
|---------|---------|
| `priv_key` | Validator private key used for signing |
| `[bridge]` | Source chain RPC, chain ID, contract address, and token list |
| `[bridge_hub]` | Destination chain RPC, chain ID, and BridgeHub address |
| `source` | BadgerDB path for persisting scan progress |

Both `[bridge]` and `[bridge_hub]` accept any EVM chain — set `node_url` and `chain_id` to match your networks.

### 4. Run

```bash
./bin/bridge
```

The service starts two block scanners (one per chain) and optionally a withdrawal finalizer when `send_finalize_withdrawals = true`.

## Documentation

| Doc | Contents |
|-----|----------|
| [docs/README.md](docs/README.md) | Operator, deposit/withdraw, and wallet integration guides (Chinese) |
| [solidity/README.md](solidity/README.md) | Contract deployment (`Bridge`, `BridgeHub`, token pairs) |
| [CONTRIBUTING.md](CONTRIBUTING.md) | How to build, test, and open a pull request |
| [SECURITY.md](SECURITY.md) | Vulnerability reporting |

## Configuration Reference

`config.toml` controls logging, caching, network endpoints, contract addresses, scan intervals, block confirmation depth, and transaction gas parameters. See [.config.toml](.config.toml) for a fully annotated template.

Notable options:

- **`bridge_tokens`** — ERC-20 contract addresses to watch on the source chain
- **`block_delay`** — number of confirmations before processing a block
- **`clear_cache`** — reset scan progress and restart from `start_block`
- **`no_send`** — simulate transactions without broadcasting (useful for dry runs)
- **`ENABLE_FINALIZE_WITHDRAWALS=true`** — environment variable gate for the withdrawal finalizer task

## Project Structure

```
bridge/
├── main.go                 # CLI entry point
├── server.go               # Service orchestration
├── config.go               # Configuration types
├── .config.toml            # Configuration template
├── docs/                   # Operator and wallet guides
├── internal/
│   ├── bridge/             # Source-chain logic
│   ├── bridgehub/          # Destination-chain logic
│   ├── scanner/            # Block range scanner
│   ├── scheduler/          # Periodic task runner
│   ├── cache/              # BadgerDB block cache
│   ├── contract/           # Generated contract bindings
│   └── evm/                # RPC client and account manager
└── solidity/
    ├── contracts/          # Bridge, BridgeHub, BridgeERC20
    ├── scripts/            # Hardhat deployment scripts
    └── deploy.sh           # Deployment helper
```

## Development

```bash
make build    # compile to ./bin/bridge
make install  # install to $GOPATH/bin
make lint     # run golangci-lint
make format   # auto-format Go code
```

Contract development:

```bash
cd solidity
npm install
npm run compile
```

## References

The source-chain [`Bridge`](solidity/contracts/Bridge.sol) contract is adapted from [Hyperliquid Bridge2.sol](https://github.com/hyperliquid-dex/contracts/blob/audit2/Bridge2.sol) in the [hyperliquid-dex/contracts](https://github.com/hyperliquid-dex/contracts) repository (Apache-2.0). It inherits the validator quorum, dispute period, and withdrawal finalization model from the upstream design.

## Security

This project handles cross-chain asset transfers. Review [SECURITY.md](SECURITY.md) before deploying to production, and report vulnerabilities through GitHub Security Advisories rather than public issues.

## License

[MIT](LICENSE)

## Contributing

See [CONTRIBUTING.md](CONTRIBUTING.md). Open an issue to discuss significant changes before submitting a pull request.
