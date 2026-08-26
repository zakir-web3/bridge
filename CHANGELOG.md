# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [1.0.0] - 2026-08-26

First public release of the self-hosted EVM lock-and-mint bridge.

### Added

- **Solidity contracts** — `Bridge` (source chain lock/deposit), `BridgeHub` (destination chain mint/withdraw, UUPS upgradeable), and `BridgeERC20` wrapped tokens
- **Go relayer** — dual-chain block scanners, EIP-712 validator quorum signing, BadgerDB scan progress cache, and optional withdrawal finalizer
- **Deployment tooling** — Hardhat scripts and `solidity/deploy.sh` for contract deployment and upgrades
- **ABI pipeline** — `compile_abi.sh` to generate Go bindings from Solidity ABIs
- **Tests** — Go unit tests and Hardhat contract tests (decimal conversion, deposit/withdraw flows)
- **CI** — GitHub Actions workflow for Go and Solidity on every push and pull request
- **Documentation** — README (English and Chinese), contributing guide, security policy, and operator configuration template (`.config.toml`)

### Notes

- Source-chain `Bridge` is adapted from [Hyperliquid Bridge2.sol](https://github.com/hyperliquid-dex/contracts/blob/audit2/Bridge2.sol) (Apache-2.0).
- Review [SECURITY.md](SECURITY.md) before deploying to production.

[1.0.0]: https://github.com/zakir-web3/bridge/releases/tag/v1.0.0
