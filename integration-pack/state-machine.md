# Cross-chain state machine (integrator view)

Happy path (EVM Hub ↔ Solana / EVM Bridge) — names are conceptual; bind to real events below.

```
User initiates
  → SourceLocked / HubWithdrawRequested
  → RelayerQuorumGathered          (off-chain EIP-712)
  → DestinationFinalized           (mint or Solana withdraw)
  → OptionalConfirm / Completed
```

## Key EVM events (listen these)

### Bridge (source)

- `Deposit(...)`
- `RequestedWithdrawal(...)`
- `FinalizedWithdrawal(...)`
- `FailedWithdrawal(bytes32 message, uint32 errorCode)`

### BridgeHub (destination)

- `Deposit(...)` — mint side success
- `FailedDeposit(bytes32 message, uint32 errorCode)`
- `Withdraw(...)` — user burn/withdraw request toward another chain
- `WithdrawCompleted(uint256 nonce, bytes32 message)`
- `BridgeSignatureSubmitted(...)`

## Solana

- Program emits withdraw finalization events consumed by relayer → Hub `withdrawConfirm` path.
- Client must not assume relayer timing; UI should key off `message` / `nonce`.

## Failure branches

| Symptom | Likely cause |
|---------|----------------|
| Wrong recovered validator / `UnknownValidator` | Domain separator mismatch (program id / chainId / verifyingContract) |
| Signature verify fails | Empty-sig padding applied on Solana path; or low-S / v mismatch |
| Stuck RequestedWithdrawal | Dispute period / finalizer / locker pause — ops, not frontend invent |
