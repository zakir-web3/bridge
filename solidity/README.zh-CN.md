# 合约

在**源链**部署 `Bridge`（锁定 / 存入），在**目标链**部署 `BridgeHub`（铸造 / 提款）。使用 `deploy.sh` 与 Hardhat。任意提供稳定 RPC 和 chain ID 的 EVM 网络均可；`bsc` / `bscTestnet` 只是便捷预设。

[English](README.md) · [仓库 README](../README.zh-CN.md)

## 前置条件

- Node.js ≥ 22.18.0、npm ≥ 10.9.3
- `jq`（脚本用它解析 JSON 地址）
- 部署账户在目标网络上有足够 Gas
- 稳定的 RPC 地址

## 网络

`NETWORK` 对应 `hardhat.config.ts` 中的名称：

| `NETWORK` | 含义 |
|-----------|------|
| `custom` | 任意 EVM 链。设置 `ETH_RPC_URL` 和 `CHAIN_ID`（默认 `1337`） |
| `bsc` | BSC 主网预设（`chainId = 56`） |
| `bscTestnet` | BSC 测试网预设（`chainId = 97`） |
| `localhost` | 本地 Hardhat / Anvil |

`ETH_RPC_URL` 是当前正在部署的那条链的 RPC，不限于 Ethereum。

## 环境变量

两边合约都必须配置，且必须一致：

- `HOT_ADDRESSES` — 逗号分隔的验证者地址
- `COLD_ADDRESSES` — 逗号分隔的冷备验证者地址（可与热钱包相同）
- `POWERS` — 与地址一一对应的整数权重，例如 `1,1,2`

部署：

- `PRIVATE_KEY` — `0x` 开头的部署私钥
- `ETH_RPC_URL` — 对应 `NETWORK` 的 RPC
- `NETWORK` — 见上表
- `CHAIN_ID` — `custom` 网络必填；注册代币对时表示**源链** chain ID

代币对（调用 `bridgeToken` 时）：

- `TOKEN_ADDRESS` — **源链** ERC-20
- `BRIDGED_TOKEN_ADDRESS` — **目标链** 对应代币
- `BRIDGE_HUB_ADDRESS` — 已部署的 `BridgeHub`

`Bridge` 可选参数（有默认值）：

- `DISPUTE_PERIOD_SECONDS`（默认 `200`）
- `BLOCK_DURATION_MILLIS`（默认 `750`）
- `LOCKER_THRESHOLD`（默认 `1`）

## 安装与编译

```bash
cd solidity
npm install
npm run compile
```

## 配置验证者

```bash
export HOT_ADDRESSES=0xHot1,0xHot2
export COLD_ADDRESSES=0xCold1,0xCold2
export POWERS=1,1
```

热、冷地址与权重长度必须一致。

## 部署 Bridge（源链）

下面以 BSC 预设为例。其他 EVM 链使用 `NETWORK=custom` 并设置 `CHAIN_ID`。

```bash
export NETWORK=bsc
export PRIVATE_KEY=0x你的私钥
export ETH_RPC_URL=https://your-source-rpc

bash ./deploy.sh bridge
```

执行完成后终端会输出 Bridge 合约地址。

测试 Token（仅本地 / 测试网）：

```bash
export NETWORK=bsc
export PRIVATE_KEY=0x你的私钥

npx hardhat --network "$NETWORK" run scripts/deploy-bridge-token.ts
```

## 部署 BridgeHub（目标链）

```bash
export NETWORK=custom
export CHAIN_ID=1337
export PRIVATE_KEY=0x你的私钥
export ETH_RPC_URL=https://your-destination-rpc

bash ./deploy.sh bridgeHub
```

执行完成后终端会输出 BridgeHub 合约地址。

测试 Token（仅本地 / 测试网）：

```bash
export NETWORK=custom
export PRIVATE_KEY=0x你的私钥

npx hardhat --network "$NETWORK" run scripts/deploy-bridge-token.ts
```

## 注册代币对并授予 mint/burn

这里的 `CHAIN_ID` 是代币所在的**源链** ID。

```bash
export NETWORK=custom
export PRIVATE_KEY=0x你的私钥
export CHAIN_ID=56
export TOKEN_ADDRESS=0x源链Token
export BRIDGED_TOKEN_ADDRESS=0x目标链Token
export BRIDGE_HUB_ADDRESS=0xBridgeHub

bash ./deploy.sh bridgeToken
```

## 提款手续费

不设置则免费。每个目标链 token 单独计费。

```bash
export NETWORK=custom
export PRIVATE_KEY=0x你的私钥
export BRIDGE_HUB_ADDRESS=0xBridgeHub
export BRIDGED_TOKEN_ADDRESS=0x目标链Token
export WITHDRAW_FEE=1000000000000000000

bash ./deploy.sh setWithdrawFee
```

## 常见问题

- 验证者地址与权重长度必须一致，否则脚本会中止。
- 请确认部署私钥余额充足，且 `ETH_RPC_URL` 指向 `NETWORK` 对应的链。
- 未安装 `jq` 会导致脚本读取地址失败：`brew install jq`（macOS）。
