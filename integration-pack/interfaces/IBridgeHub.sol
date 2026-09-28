// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.20;

/// @notice Minimal integrator-facing surface of BridgeHub (trimmed from main @ ab82cd1).
interface IBridgeHub {
    event Deposit(
        bytes32 indexed message,
        bytes32 indexed user,
        address destination,
        bytes32 indexed token,
        uint256 amount,
        uint256 chainId,
        uint64 blockNumber,
        bytes32 txHash,
        uint32 index,
        uint64 nonce
    );

    event FailedDeposit(bytes32 message, uint32 errorCode);

    event Withdraw(
        bytes32 indexed message,
        address indexed user,
        bytes32 destination,
        bytes32 indexed token,
        uint256 amount,
        uint256 chainId,
        uint64 nonce
    );

    event WithdrawCompleted(uint256 indexed nonce, bytes32 indexed message);

    function withdraw(
        bytes32 destination,
        address token,
        uint256 amount,
        uint256 chainId
    ) external;
}
