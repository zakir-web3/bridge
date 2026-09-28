# Zakir Bridge — AI Integration Guide

**Pack:** `integration-pack@0.1.0-preview`  
**Protocol:** self-hosted lock-and-mint（EVM Bridge / BridgeHub + Solana program + Go relayer，EIP-712 validator quorum）

## Who this is for

Use this pack when integrating a **frontend**, **another contract**, or **another service** with Zakir Bridge.

## Do NOT

- Do not clone or analyze the full `zakir-web3/bridge` monorepo unless the user explicitly asks for protocol changes.
- Do not invent EIP-712 domain fields, type hashes, or Solana transaction versions.
- Do not fill empty validator signatures on Solana withdraw paths (EVM empty-sig padding must not be applied).
- Do not use Solana Transaction v0 + ALT；**v1 only**.
- Do not hardcode Solana `DOMAIN_SEPARATOR` to a fixed program id；derive verifyingContract from the **deployed** program id.

## Roles (do not swap)

| Role | Chain artifact | Meaning |
|------|----------------|---------|
| Source lock | EVM `Bridge` or Solana program | User locks / deposits |
| Destination mint/burn | EVM `BridgeHub` | Mint bridged ERC-20 / burn to withdraw |
| Relayer | Go binary | Scans events, gathers validator EIP-712 quorum, submits |

## Canonical constants

- Solana canonical `chainId`: `900001`
- Bridge EIP-712 domain: `name = "Bridge"`, `version = "1"`
- BridgeHub EIP-712 domain: `name = "BridgeHub"`, `version = "1"`
- Solana verifyingContract: `address(keccak256(program_id_bytes)[12:32])`

See `eip712/` for types and golden vectors.

## Integration surfaces

1. **EVM app / frontend:** call `Bridge` deposit/withdraw flows; listen Hub `Deposit` / `Withdraw` / `WithdrawCompleted`; poll or index by `message` / `nonce`.
2. **EVM contract:** depend on `interfaces/IBridge.sol` + `IBridgeHub.sol`; never re-implement signature recovery differently from `Signature.sol`.
3. **Solana client:** build withdraw ix with **Transaction v1 only**; domain separator from deployed program id (see vectors).
4. **Ops / indexer:** prefer events in `state-machine.md` over scraping storage.

## Required reading order for an AI

1. This file  
2. `state-machine.md`  
3. `eip712/domains.json` + `eip712/types.json` + `eip712/vectors-solana-withdraw.json`  
4. Matching task card in `task-cards.md`  
5. `addresses.example.json` (filled by deployer)

## Success criteria

- Digests match golden vectors before any mainnet/testnet send.
- Events used for UX match `state-machine.md` field names.
- Solana txs are version **legacy/v1** (not v0+ALT).
