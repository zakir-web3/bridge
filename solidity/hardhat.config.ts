import { HardhatUserConfig } from "hardhat/config";
import "hardhat-dependency-compiler";
import "hardhat-gas-reporter";
import "@nomicfoundation/hardhat-ethers";
import "@nomicfoundation/hardhat-verify";
import "@nomicfoundation/hardhat-chai-matchers";
import "@typechain/hardhat";
import "dotenv/config";

const config: HardhatUserConfig = {
  defaultNetwork: "hardhat",
  networks: {
    localhost: {
      url: `${process.env.LOCAL_URL || "http://127.0.0.1:9545"}`,
    },
    destination: {
      url: `${process.env.ETH_RPC_URL || "http://127.0.0.1:9545"}`,
      chainId: parseInt(process.env.CHAIN_ID || "1337", 10),
      accounts: process.env.PRIVATE_KEY ? [process.env.PRIVATE_KEY] : undefined,
    },
    bsc: {
      url: `${
        process.env.ETH_RPC_URL || "https://bsc-mainnet.public.blastapi.io"
      }`,
      chainId: 56,
      accounts: process.env.PRIVATE_KEY ? [process.env.PRIVATE_KEY] : undefined,
    },
    bscTestnet: {
      url: `${
        process.env.ETH_RPC_URL || "https://bsc-testnet-dataseed.bnbchain.org"
      }`,
      chainId: 97,
      accounts: process.env.PRIVATE_KEY ? [process.env.PRIVATE_KEY] : undefined,
    },
  },
  solidity: {
    compilers: [
      {
        version: "0.8.20",
        settings: {
          optimizer: {
            enabled: true,
            runs: 200,
          },
        },
      },
      {
        version: "0.8.22",
        settings: {
          optimizer: {
            enabled: true,
            runs: 200,
          },
        },
      },
    ],
  },
  etherscan: {
    apiKey: `${process.env.ETHERSCAN_API_KEY}`,
  },
  dependencyCompiler: {
    paths: [
      "@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol",
      "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol",
    ],
  },
  gasReporter: {
    enabled: true,
    gasPrice: 1,
  },
};

export default config;
