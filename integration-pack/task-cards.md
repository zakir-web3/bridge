# Task cards（直接粘贴给对接方 AI）

把对应卡片整段复制到 Cursor，并 `@` 本 pack 目录。

---

## Card A — Next.js：Hub withdraw UI

```
只用 ~/bridge-integration 包，不要分析 bridge 主仓。
在现有 Next.js 里加 BridgeHub.withdraw 调用：
- ABI/接口以 interfaces/IBridgeHub.sol 为准
- 地址从 addresses.example.json（已填的那份）读取
- 成功后监听 Withdraw / WithdrawCompleted
- 先用 eip712 向量验证你对 bytes32 destination/token 的编码理解
验收：本地类型检查通过；事件名与 state-machine.md 一致。
```

---

## Card B — Solidity：外部合约监听 Deposit

```
只用 integration pack。
写一个最小合约：监听 IBridgeHub.Deposit，把 message/nonce 记到自己的映射。
禁止重新实现 EIP-712；禁止调用未在 interfaces/ 出现的 admin 函数。
验收：能编译；事件参数顺序与 IBridgeHub.sol 一致。
```

---

## Card C — Solana withdraw 客户端

```
只用 integration pack + notes/solana.md。
实现：根据 deployed program id 推导 verifyingContract 与 domainSeparator；
对 vectors-solana-withdraw.json 里 alt program 例子算出相同 digest。
交易版本只能 v1。
验收：domainSeparator 与 digest 与 JSON 完全一致。
```
