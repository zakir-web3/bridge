# Solana 桥接程序

Anchor 程序，用于 **Solana → EVM** 存款：用户在 Solana 锁 SPL；Go relayer 在 EVM `BridgeHub` 上签 EIP-712 并铸造桥接 ERC-20。

[English](README.md) · [根目录 README](../README.md)

## 在跨链中的位置

```
用户 (Solana)  →  deposit(SPL, evm_destination)  →  vault PDA
Relayer        →  扫描 DepositEvent  →  BridgeHub.depositConfirm (EVM)
```

本程序只负责 **托管 SPL** 和发出事件。验证人签名校验在 EVM `BridgeHub` 上完成，不在 Solana 链上验签。

## 环境要求

- [Rust](https://www.rust-lang.org/tools/install)
- [Solana CLI](https://docs.solanalabs.com/cli/install)
- [Anchor 1.1.2](https://www.anchor-lang.com/docs/installation)（与 `Anchor.toml` 一致）
- Node.js（运行 `anchor test`）

钱包与集群在 `Anchor.toml` 的 `provider` 中配置。覆盖 RPC：

```bash
export ANCHOR_PROVIDER_URL=https://api.devnet.solana.com
export ANCHOR_WALLET=~/.config/solana/id.json
```

## 安装依赖

```bash
cd solana
npm install
```

## 编译

```bash
anchor build
```

IDL 与 TypeScript 类型生成在 `target/idl`、`target/types`。

## 测试（本地验证节点）

自动启动本地 validator、部署程序并运行测试：

```bash
anchor test
```

测试文件：`tests/bridge.ts`（`initialize` → `initialize_vault` → `deposit`）。

## 部署

```bash
# localnet（需先运行 solana-test-validator）
anchor deploy

# devnet
anchor deploy --provider.cluster devnet
```

部署后确认 Program ID（若使用默认 keypair 部署，与 `Anchor.toml` 中一致）。

## 链上初始化

配置 RPC 与钱包（或使用 `Anchor.toml` 的 `provider`）：

```bash
export ANCHOR_PROVIDER_URL=https://api.devnet.solana.com
export ANCHOR_WALLET=~/.config/solana/id.json
```

### 1. initialize

每个部署调用一次，创建 `config` PDA（`seeds = ["config"]`），设置 `admin`、`paused = false`。

```bash
npm run initialize
# 或：anchor deploy --provider.cluster devnet && npm run initialize
```

### 2. initialize_vault（每个 SPL mint）

每种要跨的 SPL mint 在用户 deposit 前必须先初始化 vault。每个 mint 单独执行一次：

- `vault_state` PDA：`seeds = ["vault_state", mint]`
- `vault_authority` PDA：`seeds = ["vault", mint]`
- `vault_token_account`：`vault_authority` 的 ATA

```bash
npm run initialize-vault -- <MINT_PUBKEY>
# 或：MINT=<MINT_PUBKEY> npm run initialize-vault
```

账户组装参考 `scripts/initialize-vault.ts` 与 `tests/bridge.ts`。

### 3. deposit

指令 `deposit(destination: [u8; 20], amount: u64)`：

- `destination` — 20 字节 EVM 地址（`BridgeHub` 铸币收款地址）
- 通过 `transfer_checked` 将 SPL 从用户 ATA 转入 vault ATA
- 发出 `DepositEvent { user, destination, mint, amount, slot }`

## 发送测试 deposit

在 devnet 完成程序部署与 vault 初始化后：

```bash
npm run deposit -- <MINT_PUBKEY> <AMOUNT> <EVM_DESTINATION>
# 或：MINT=<MINT_PUBKEY> AMOUNT=<AMOUNT> DESTINATION=0x... npm run deposit
```

`AMOUNT` 为 SPL 最小单位（例如 6 位小数时 `1000000` = 1 个 token）。`DESTINATION` 为 20 字节 EVM 收款地址。

脚本以 `ANCHOR_WALLET` 作为存款用户；若用户 ATA 不存在会自动创建。

## Relayer 配置

在 `config.toml` 中配置 `[solana_bridge]`（模板见 `.config.toml`）：

- `chain_id` — 如 devnet 用 `900001`（须与 EVM `setTokenPair` 一致）
- `program_id` — 部署后的程序 ID
- `bridge_mints` — SPL mint 白名单

Relayer 解析日志（`Program data:` base64），组装 EIP-712 `Deposit`（含 `blockNumber`、`txHash`、`index`），调用 `BridgeHub.depositConfirm`。

完整步骤见 `internal/solana/` 与 `scripts/solana-e2e.sh`。

## 指令说明

| 指令               | 调用方 | 作用                         |
| ------------------ | ------ | ---------------------------- |
| `initialize`       | admin  | 创建全局 config PDA          |
| `initialize_vault` | admin  | 为某个 SPL mint 创建 vault   |
| `deposit`          | user   | 锁 SPL，指定 EVM destination |
| `pause` / `unpause`| admin  | 紧急暂停 deposit             |

## 目录结构

```
solana/
  Anchor.toml
  programs/bridge/src/lib.rs
  scripts/initialize.ts
  scripts/initialize-vault.ts
  scripts/deposit.ts
  tests/bridge.ts
  package.json
```

## Solana chainId（EVM 配置）

在 relayer 与 `setTokenPair` 中使用固定的 `uint256`：

| 网络     | 建议 `chain_id` |
| -------- | ----------------- |
| Devnet   | `900001`          |
| Mainnet  | `1399811149`      |

须与 Solana 配置、EIP-712 `Deposit.chainId`、`setTokenPair` 全链路一致。
