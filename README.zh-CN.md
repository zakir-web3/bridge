# Bridge

[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)
[![Release](https://img.shields.io/github/v/release/zakir-web3/bridge)](https://github.com/zakir-web3/bridge/releases)

自托管的 EVM **lock-and-mint** 跨链桥：用户在源链 `Bridge` 锁定 ERC-20，Go 中继服务收集你自己的验证者集合的 EIP-712 签名，在目标链 `BridgeHub` 铸币。

面向自己运营验证者的应用链和 Token 发行方。这不是 LayerZero / Hyperlane 那种通用消息协议，也不是面向散户的兑换路由。

源链 `Bridge` 合约参考自 [Hyperliquid Bridge2.sol](https://github.com/hyperliquid-dex/contracts/blob/audit2/Bridge2.sol)。

[English](README.md) · [文档中心](docs/README.md) · [贡献指南](CONTRIBUTING.md)

## 概述

Bridge 通过一对合约连接两条 EVM 链：

- **`Bridge`** — 部署在**源链**，用户在此锁定或存入代币
- **`BridgeHub`** — 部署在**目标链**，在此铸造桥接代币并发起提款

中继服务在两侧运行区块扫描器，收集验证者签名并将确认信息 relay 到另一条链。只要提供稳定的 RPC 端点和正确的 chain ID，任何 EVM 网络均可接入 — 包括 Ethereum、BSC、Polygon、Arbitrum、Base 或自定义链。

```
  源链 (EVM)                            目标链 (EVM)
 ┌─────────────────────┐               ┌─────────────────────┐
 │      Bridge.sol     │               │   BridgeHub.sol     │
 │  锁定 / 存入        │               │  铸造 / 提款        │
 └──────────┬──────────┘               └──────────┬──────────┘
            │                                     │
            └──────────►  中继服务  ◄──────────────┘
                    (扫描 · 签名 · 提交)
```

## 特性

- **EVM 通用** — 支持所有 EVM 兼容链，在 `config.toml` 中配置 RPC 地址和 chain ID 即可
- **ERC-20 支持** — 通过白名单代币实现存入锁定/铸造、提款销毁/释放
- **验证者 quorum** — 跨链消息由配置的验证者集合通过 EIP-712 签名授权
- **争议期** — 源链上的提款进入 pending 状态，经过争议期后方可最终确认
- **可靠扫描** — 区块进度持久化到 BadgerDB，支持可配置的确认深度和重试逻辑
- **灵活 Gas 控制** — 支持 fee history、gas 上限，以及 dry-run（`no_send`）模式

## 环境要求

| 工具 | 版本 |
|------|------|
| Go | 1.25+ |
| Node.js | 22.18.0+ |
| npm | 10.9.3+ |

## 快速开始

### 1. 克隆并构建

```bash
git clone https://github.com/zakir-web3/bridge.git
cd bridge
make build
```

### 2. 部署合约

在源链部署 `Bridge`，在目标链部署 `BridgeHub`。完整部署流程见 [solidity/README.md](solidity/README.md)。

### 3. 配置中继服务

复制配置模板并填入实际值：

```bash
cp .config.toml config.toml
```

主要配置项：

| 配置段 | 说明 |
|--------|------|
| `priv_key` | 用于签名的验证者私钥 |
| `[bridge]` | 源链 RPC、chain ID、合约地址和代币列表 |
| `[bridge_hub]` | 目标链 RPC、chain ID 和 BridgeHub 地址 |
| `source` | BadgerDB 路径，用于持久化扫描进度 |

`[bridge]` 和 `[bridge_hub]` 均支持任意 EVM 链 — 将 `node_url` 和 `chain_id` 设置为对应网络即可。

### 4. 运行

```bash
./bin/bridge
```

服务会启动两个区块扫描器（每条链各一个），并在 `send_finalize_withdrawals = true` 时可选运行提款最终确认任务。

## 文档

| 文档 | 内容 |
|------|------|
| [docs/README.md](docs/README.md) | 运维、存款/提款、钱包集成 |
| [solidity/README.md](solidity/README.md) | 合约部署（`Bridge`、`BridgeHub`、代币对） |
| [CONTRIBUTING.md](CONTRIBUTING.md) | 如何构建、测试和提交 PR |
| [SECURITY.md](SECURITY.md) | 漏洞报告 |

## 配置说明

`config.toml` 控制日志、缓存、网络端点、合约地址、扫描间隔、区块确认深度和交易 gas 参数。完整注释模板见 [.config.toml](.config.toml)。

常用选项：

- **`bridge_tokens`** — 源链上需要监控的 ERC-20 合约地址
- **`block_delay`** — 处理区块前等待的确认数
- **`clear_cache`** — 重置扫描进度，从 `start_block` 重新开始
- **`no_send`** — 模拟交易而不广播（适用于 dry run）
- **`ENABLE_FINALIZE_WITHDRAWALS=true`** — 启用提款最终确认任务的环境变量开关

## 项目结构

```
bridge/
├── main.go                 # CLI 入口
├── server.go               # 服务编排
├── config.go               # 配置类型定义
├── .config.toml            # 配置模板
├── docs/                   # 运维与钱包集成指南
├── internal/
│   ├── bridge/             # 源链逻辑
│   ├── bridgehub/          # 目标链逻辑
│   ├── scanner/            # 区块范围扫描器
│   ├── scheduler/          # 定时任务调度
│   ├── cache/              # BadgerDB 区块缓存
│   ├── contract/           # 生成的合约绑定
│   └── evm/                # RPC 客户端和账户管理
└── solidity/
    ├── contracts/          # Bridge、BridgeHub、BridgeERC20
    ├── scripts/            # Hardhat 部署脚本
    └── deploy.sh           # 部署辅助脚本
```

## 开发

```bash
make build    # 编译到 ./bin/bridge
make install  # 安装到 $GOPATH/bin
make lint     # 运行 golangci-lint
make format   # 自动格式化 Go 代码
```

合约开发：

```bash
cd solidity
npm install
npm run compile
```

## 参考

源链 [`Bridge`](solidity/contracts/Bridge.sol) 合约参考自 [hyperliquid-dex/contracts](https://github.com/hyperliquid-dex/contracts) 仓库中的 [Hyperliquid Bridge2.sol](https://github.com/hyperliquid-dex/contracts/blob/audit2/Bridge2.sol)（Apache-2.0 许可证），并沿用了上游设计中的验证者 quorum、争议期与提款最终确认机制。

## 安全

本项目涉及跨链资产转移。生产部署前请阅读 [SECURITY.md](SECURITY.md)，并通过 GitHub Security Advisories 报告漏洞，而非公开 issue。

## 许可证

[MIT](LICENSE)

## 贡献

见 [CONTRIBUTING.md](CONTRIBUTING.md)。较大改动请先开 issue 讨论后再提交 pull request。
