// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.20;

/// @notice Minimal integrator-facing surface of source-chain Bridge (trimmed from main @ ab82cd1).
interface IBridge {
    event Deposit(
        address indexed user,
        address destination,
        address token,
        uint256 amount
    );

    event RequestedWithdrawal(
        bytes32 message,
        address indexed user,
        bytes32 destination,
        bytes32 token,
        uint256 amount,
        uint256 chainId,
        uint64 nonce,
        uint64 requestedTime,
        uint64 requestedBlockNumber
    );

    event FinalizedWithdrawal(
        bytes32 message,
        address indexed user,
        bytes32 destination,
        bytes32 token,
        uint256 amount,
        uint64 nonce
    );

    event FailedWithdrawal(bytes32 message, uint32 errorCode);
}
