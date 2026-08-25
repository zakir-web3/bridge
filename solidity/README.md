# 合约

## 主网部署指南

以下步骤基于 `deploy.sh` 与 Hardhat 配置，指导在主网（例如 BSC 主网或自定义目标链）部署 `BridgeHub` 与 `Bridge` 合约。

### 前置条件

- 已安装 Node.js ≥ v22.18.0、npm ≥ 10.9.3
- 已安装 jq（脚本中用于解析 JSON）
- 部署账户私钥中主网有充足 Gas 资产（例如 BSC 需 BNB）
- 可用的 RPC 地址（建议使用自有或稳定的服务商）

### 环境变量（必填/常用）

- `ETH_RPC_URL`: 目标主网 RPC URL
- `PRIVATE_KEY`: 部署私钥（0x 开头的原始私钥）
- `NETWORK`: 目标网络标识，支持：
  - `bsc`（BSC 主网，chainId=56）
  - `destination`（自定义目标链，chainId 通过 `CHAIN_ID` 环境变量配置，默认 1337）
- 验证者配置（部署 `BridgeHub` 和 `Bridge` 均会用到，且必须相同）：
  - `HOT_ADDRESSES`: 逗号分隔的验证者地址列表，例如 `0xabc,...`
  - `COLD_ADDRESSES`: 逗号分隔的冷备验证者地址列表（如无需可与 HOT 相同）
  - `POWERS`: 与地址一一对应的权重，逗号分隔的整数，例如 `1,1,2`
- Token 地址（在 `BridgeHub` 中添加跨链 Token 关系）：
  - `TOKEN_ADDRESS`: BSC 链 Token 合约地址
  - `BRIDGED_TOKEN_ADDRESS`: 自有链 Token 合约地址
- 业务参数（部署 `Bridge` 用）：
  - `DISPUTE_PERIOD_SECONDS`（默认 200）
  - `BLOCK_DURATION_MILLIS`（默认 750）
  - `LOCKER_THRESHOLD`（默认 1）

### 一次性安装依赖与编译

```bash
cd solidity
npm install
npm run compile
```

### 配置验证者

```bash
export HOT_ADDRESSES=0xHot1,0xHot2
export COLD_ADDRESSES=0xCold1,0xCold2
export POWERS=1,1
```

### 部署 Bridge

```bash
export NETWORK=bsc
export PRIVATE_KEY=0x你的私钥

bash ./deploy.sh bridge
```

执行完成后终端会输出 Bridge 合约地址

#### 部署测试 Token，正式环境不需要部署

```bash
export NETWORK=bsc
export PRIVATE_KEY=0x你的私钥

npx hardhat --network "$NETWORK" run scripts/deploy-bridge-token.ts
```

### 部署 BridgeHub

```bash
export NETWORK=destination
export PRIVATE_KEY=0x你的私钥

bash ./deploy.sh bridgeHub
```

执行完成后终端会输出 BridgeHub 合约地址

#### 部署测试 Token，正式环境不需要部署

```bash
export NETWORK=destination
export PRIVATE_KEY=0x你的私钥

npx hardhat --network "$NETWORK" run scripts/deploy-bridge-token.ts
```

#### 添加跨链 token 关系,并将自有链 token 的 mint/burn 权限授权给 BridgeHub 合约

```bash
export NETWORK=destination
export PRIVATE_KEY=0x你的私钥
export CHAIN_ID=56 # BSC 主网 chainId
export TOKEN_ADDRESS=0xBSC链 USDT Token 地址
export BRIDGED_TOKEN_ADDRESS=0x自有链 USDT Token 地址
export BRIDGE_HUB_ADDRESS=0x部署完成的 BridgeHub 地址

bash ./deploy.sh bridgeToken
```

#### 设置跨链提款手续费

> 不设置则默认免费，每个 token 都是独立的提款手续费

```bash
export NETWORK=destination
export PRIVATE_KEY=0x你的私钥
export BRIDGE_HUB_ADDRESS=0x部署完成的 BridgeHub 地址
export BRIDGED_TOKEN_ADDRESS=0x自有链 USDT Token 地址
export WITHDRAW_FEE=1000000000000000000 # 1 USDT
bash ./deploy.sh setWithdrawFee
```

### 常见问题

- 验证者地址与权重长度必须一致，否则脚本会中止。
- 请确认部署私钥余额充足，且 `ETH_RPC_URL` 指向目标主网。
- 未安装 `jq` 会导致脚本读取地址失败，请先安装：`brew install jq`（macOS）。
