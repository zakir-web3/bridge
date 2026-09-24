# Bridge — Agent Guide

Self-hosted **lock-and-mint** cross-chain bridge: EVM `Bridge` (lock/deposit) + `BridgeHub` (mint/withdraw, UUPS) + Go relayer (dual-chain scan, EIP-712 validator quorum, BadgerDB cache).

For appchain and token-issuer teams with their own validator set. **Not** a general omnichain messaging protocol (LayerZero, Hyperlane) or a retail swap router.

Source-chain `Bridge` is adapted from [Hyperliquid Bridge2.sol](https://github.com/hyperliquid-dex/contracts/blob/audit2/Bridge2.sol).

## When to use this repo

| Need | Use this? |
|------|-----------|
| Self-hosted ERC-20 lock/mint with your validator set | Yes |
| EIP-712 signed deposits and withdrawals | Yes |
| Solana → EVM deposits (SPL lock, EVM mint) | Yes — see `solana/` |
| General omnichain messaging | No — use LayerZero, Hyperlane, etc. |
| Retail swap / liquidity routing | No |

## Key paths

| Area | Path |
|------|------|
| Source-chain contract | `solidity/contracts/Bridge.sol` |
| Destination-chain contract | `solidity/contracts/BridgeHub.sol` |
| Bridged ERC-20 | `solidity/contracts/BridgeERC20.sol` |
| Source-chain relayer | `internal/bridge/` |
| Destination-chain relayer | `internal/bridgehub/` |
| Block scanner / scheduler / cache | `internal/scanner/`, `internal/scheduler/`, `internal/cache/` |
| Generated bindings | `internal/contract/*.sol.go` |
| Handwritten EIP-712 wrappers | `internal/contract/{bridge,bridge_hub,bridge_erc20}.go` |
| Solana source program | `solana/programs/bridge/` (Anchor) |
| Solana scanner | `internal/solana/` |
| Relayer config template | `config.example.toml` → copy to `config.toml` locally |
| Contract deployment | `solidity/deploy.sh` only — do not add separate deploy scripts |
| Operator docs | `docs/README.md` |

## Commands

```bash
make build          # compile relayer to ./bin/bridge
make test           # Go tests
make lint           # golangci-lint
./compile_abi.sh    # regenerate bindings after ABI / EIP-712 changes

cd solidity && npm install && npm test   # contract tests
```

Run relayer: `./bin/bridge` (requires local `config.toml`).

## Hard rules

- **Never** hand-edit `internal/contract/*.sol.go` — run `./compile_abi.sh` from repo root instead.
- **Never** commit `config.toml`, `.env`, private keys, or RPC secrets. **Never** log `priv_key`.
- **Bridge** = source chain (lock/deposit). **BridgeHub** = destination chain (mint/withdraw). Do not swap roles, chain IDs, or token addresses.
- **Do not** bypass the existing scanner / scheduler / cache with a second scan or persistence path.
- Source chain only processes whitelisted `bridge_tokens`.
- `start_block` must match contract deployment height.
- `no_send = true` simulates transactions without broadcasting.
- Finalize withdrawals require `send_finalize_withdrawals = true` **and** `ENABLE_FINALIZE_WITHDRAWALS=true`.

## After ABI or EIP-712 changes

1. Run `./compile_abi.sh` from repo root.
2. Sync handwritten wrappers in `internal/contract/{bridge,bridge_hub,bridge_erc20}.go` with Solidity domain / typehash.
3. Changing only one side breaks cross-chain signature verification.

## Relayer engineering

When editing scanner, scheduler, cache, or RPC code:

- Read existing implementations first; extend them rather than reimplementing.
- Long-lived tasks use `server.SafeGo` + `scheduler.Run` — no bare `go func()`.
- Pass `context` down; RPC calls need timeouts.
- Badger access only through `cache.View` / `cache.Update`.
- New config fields go in `config.example.toml` and corresponding `Validate()`.
- See `.cursor/rules/relayer-engineering.mdc` for logging and metrics conventions.

## Further reading

- [README.md](README.md) — overview and quick start
- [CONTRIBUTING.md](CONTRIBUTING.md) — PR workflow and checks
- [solidity/README.md](solidity/README.md) — contract deployment
- [solana/README.md](solana/README.md) — Solana deposit flow
- [config.example.toml](config.example.toml) — relayer configuration reference

## External / AI integration pack

Integrators (and integrator AIs) should use [`integration-pack/`](integration-pack/) instead of scanning the whole repo.

Refresh generated artifacts after ABI / IDL / EIP-712 changes:

```bash
./scripts/export-integration-pack.sh
```
