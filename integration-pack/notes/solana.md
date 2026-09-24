# Solana notes for integrators

- Transaction version: **v1 only** (no v0 + ALT dual path).
- EIP-712 `verifyingContract` is derived from the **deployed** program id:
  `address(keccak256(program_id)[12:32])`.
- CI may run `anchor keys sync`, so repo-default `C4Yxx…` separator is not safe to hardcode against a synced deploy.
- Full IDL: build in main repo (`cd solana && anchor build`) then copy `target/idl/bridge.json` into this pack as `idl/bridge.json` when you publish.
- Permissionless on-chain withdraw verifies validator quorum with secp256k1 + EIP-712 (see main program `withdraw.rs`).
