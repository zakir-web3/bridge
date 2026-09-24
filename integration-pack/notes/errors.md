# Common errors (integrator-facing)

| Code / name | Meaning | What to check |
|-------------|---------|---------------|
| UnknownValidator / 6009 / 0x1779 | Recovered signer not in validator set | Domain separator / program id / chainId 900001 |
| Invalid signature / zero address | Bad v,r,s or digest | Low-S, v=27/28, no empty-sig padding on Solana path |
| Token not found | Hub tokenPair missing | `setTokenPair` / addresses table |
| Amount must exceed fee | Hub withdraw fee | `tokenWithdrawFee` |
