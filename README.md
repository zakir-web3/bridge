# Bridge 跨链桥接服务

一个基于 Go 和 Solidity 的跨链桥接服务，支持 EVM 兼容区块链之间的资产转移。

## 快速开始

### 环境要求

- Go 1.25+
- Node.js 22.18.0+
- npm 10.9.3+

## 项目结构

```
bridge/
├── main.go                 # 主程序入口
├── server.go              # 服务启动逻辑
├── config.go              # 配置结构定义
├── .config.toml           # 配置文件模板
├── Makefile               # 构建脚本
├── internal/              # 内部包
│   ├── bridge/            # 桥接逻辑
│   ├── bridgehub/         # 桥接中心逻辑
│   ├── scanner/           # 区块扫描器
│   ├── scheduler/         # 任务调度器
│   └── cache/             # 缓存管理
└── solidity/              # 智能合约
    ├── contracts/         # Solidity 合约源码
    ├── scripts/           # 合约部署脚本
    └── deploy.sh          # 合约部署脚本
```

## 配置说明

主要配置文件 `config.toml` 包含：

- **日志配置**: 日志级别和格式
- **网络配置**: RPC 地址和链 ID
- **合约配置**: BridgeHub 和 Bridge 合约地址
- **扫描配置**: 区块扫描间隔和延迟
- **交易配置**: Gas 限制和费用设置

## 支持的网络

- **源链**: BSC 等 EVM 兼容链
- **目标链**: 自定义 EVM 兼容链

## 开发

### 构建

```bash
make build
```

### 安装

```bash
make install
```

### 代码检查

```bash
make lint
make format
```
