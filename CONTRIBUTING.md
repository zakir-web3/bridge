# Contributing

Thank you for contributing. This repository is a self-hosted lock-and-mint bridge: Solidity contracts plus a Go relayer. Please discuss significant changes in an issue before opening a pull request.

[中文文档](README.zh-CN.md) · [Security](SECURITY.md)

## Development setup

| Tool | Version |
|------|---------|
| Go | 1.25+ |
| Node.js | 22.18.0+ |
| npm | 10.9.3+ |

```bash
git clone https://github.com/zakir-web3/bridge.git
cd bridge
make build
make test
```

Contracts:

```bash
cd solidity
npm install
npm test
```

## What to change — and what not to

- **Do** edit `solidity/contracts/` and the handwritten wrappers in `internal/contract/{bridge,bridge_hub,bridge_erc20}.go` together when you change ABI or EIP-712 types. After an ABI change, run `./compile_abi.sh` from the repo root, then keep those wrappers in sync with the Solidity domain / typehash. Changing only one side will break cross-chain signature verification.
- **Do not** hand-edit generated files `internal/contract/*.sol.go`. They are produced by `./compile_abi.sh`.
- **Do not** commit `config.toml`, `.env`, private keys, or RPC secrets. Copy `.config.toml` locally. Never log `priv_key`.
- **Do not** swap Bridge / BridgeHub roles, chain IDs, or token addresses.
- **Do not** bypass the existing scanner, scheduler, or cache with a second scan or persistence path.

Use `no_send = true` for dry-run relayer tests. Sending finalize withdrawals requires both `send_finalize_withdrawals = true` and `ENABLE_FINALIZE_WITHDRAWALS=true`. `start_block` must match the contract deployment height.

## Checks before you push

```bash
make test
make lint
cd solidity && npm test
```

If you touch contracts or bindings, also run `./compile_abi.sh` and include the regenerated files in the same PR.

## Pull requests

- Keep the change focused. One concern per PR.
- Match the existing commit style: a short sentence that explains **why**, with a period (for example, `Add Go tests for the scanner retry path.`).
- Link the related issue.
- For contract or relayer behavior changes, say how you tested (unit tests, `no_send`, which networks).

### Review and merge policy (`main`)

All pull requests targeting `main` need **one approving review** and **passing CI** (`Go`, `Solidity`, `Automerge`) before they can merge. After approval, the Automerge workflow enables GitHub auto-merge; the PR merges automatically once every required check is green. Do not merge manually unless you are bypassing policy for an emergency.

Maintainers: see [.github/BRANCH_PROTECTION.md](.github/BRANCH_PROTECTION.md) for the one-time repository ruleset and **Allow auto-merge** setup.

## Security

Do not open a public issue for fund-loss, signature forgery, replay, or privilege-escalation reports. Use [GitHub Security Advisories](https://github.com/zakir-web3/bridge/security/advisories/new). See [SECURITY.md](SECURITY.md).
