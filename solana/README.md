# Solana Bridge Program

Anchor program for **Solana → EVM** deposits: users lock SPL tokens on Solana; the Go relayer signs EIP-712 on EVM `BridgeHub` and mints bridged ERC-20.

[中文](README.zh-CN.md) · [Root README](../README.md)

## Role in the bridge

```
User (Solana)  →  deposit(SPL, evm_destination)  →  vault PDA
Relayer        →  scan DepositEvent  →  BridgeHub.depositConfirm (EVM)
```

This program only **custodies SPL** and emits events. Validator signatures are verified on EVM `BridgeHub`, not on-chain here.

## Prerequisites

- [Rust](https://www.rust-lang.org/tools/install)
- [Solana CLI](https://docs.solanalabs.com/cli/install)
- [Anchor 0.30.1](https://www.anchor-lang.com/docs/installation) (matches `Anchor.toml`)
- Node.js (for `anchor test`)

Wallet and cluster are configured in `Anchor.toml` (`provider.wallet`, `provider.cluster`). Override RPC with:

```bash
export ANCHOR_PROVIDER_URL=https://api.devnet.solana.com
export ANCHOR_WALLET=~/.config/solana/id.json
```

## Install dependencies

```bash
cd solana
npm install
```

## Build

```bash
anchor build
```

IDL and types are generated under `target/idl` and `target/types`.

## Test (local validator)

Starts a local validator, deploys the program, and runs TypeScript tests:

```bash
anchor test
```

Tests live in `tests/bridge.ts` (`initialize` → `initialize_vault` → `deposit`).

## Deploy

```bash
# localnet (solana-test-validator must be running)
anchor deploy

# devnet
anchor deploy --provider.cluster devnet
```

After deploy, record the program ID (or use the one in `Anchor.toml` if you deployed with the default keypair).

## On-chain setup

### 1. Initialize config

From tests or your own client, call `initialize` once per deployment. Creates the `config` PDA (`seeds = ["config"]`) with `admin` and `paused = false`.

### 2. Initialize vault per SPL mint

Each bridged SPL mint needs `initialize_vault` before users can deposit:

- `vault_state` PDA: `seeds = ["vault_state", mint]`
- `vault_authority` PDA: `seeds = ["vault", mint]`
- `vault_token_account`: ATA owned by `vault_authority`

See `tests/bridge.ts` for account wiring.

### 3. User deposit

Instruction `deposit(destination: [u8; 20], amount: u64)`:

- `destination` — 20-byte EVM address (where `BridgeHub` mints bridged token)
- Transfers SPL from user ATA to vault ATA via `transfer_checked`
- Emits `DepositEvent { user, destination, mint, amount, slot }`

## Send a test deposit (Go)

From the repo root, after program + vault are set up on devnet:

```bash
export SOLANA_RPC_URL=https://api.devnet.solana.com
export MINT=<SPL_mint_pubkey>
export SOLANA_KEYPAIR=~/.config/solana/id.json
export DESTINATION=0x1234567890123456789012345678901234567890
export AMOUNT=100000

go run ./cmd/solana-deposit
```

## Relayer integration

Configure `[solana_bridge]` in `config.toml` (see `.config.toml` template):

- `chain_id` — e.g. `900001` for devnet (must match EVM `setSrcTokenPair`)
- `program_id` — deployed program ID
- `bridge_mints` — SPL mint whitelist

Relayer parses logs (`Program data:` base64), builds EIP-712 `Deposit` with `slot`, `txSigHash`, `instructionIndex`, and calls `BridgeHub.depositConfirm`.

See `internal/solana/` and `scripts/solana-e2e.sh` for the full checklist.

## Program instructions

| Instruction        | Who        | Purpose                                      |
| ------------------ | ---------- | -------------------------------------------- |
| `initialize`       | admin      | Create global config PDA                     |
| `initialize_vault` | admin      | Create vault for one SPL mint                |
| `deposit`          | user       | Lock SPL; specify EVM `destination`          |
| `pause` / `unpause`| admin      | Emergency pause deposits                     |

## Directory layout

```
solana/
  Anchor.toml
  programs/bridge/src/lib.rs   # program logic
  tests/bridge.ts                # integration tests
  package.json
```

## Solana chain ID (EVM config)

Use a fixed `uint256` in relayer + `setSrcTokenPair`:

| Network | Suggested `chain_id` |
| ------- | -------------------- |
| Devnet  | `900001`             |
| Mainnet | `1399811149`         |

Must be consistent across Solana config, EIP-712 `Deposit.chainId`, and BridgeHub `setSrcTokenPair`.
