// SPDX-License-Identifier: Apache-2.0

/*
    This bridge contract runs on Arbitrum, operating alongside the Hyperliquid L1.
    The only asset for now is USDC, though the logic extends to any other ERC20 token on Arbitrum.
    The L1 runs tendermint consensus, with validator set updates happening at the end of each epoch.
    Epoch duration TBD, but likely somewhere between 1 day and 1 week.
    "Bridge2" is to distinguish from the legacy Bridge contract.

    Validators:
      Each validator has a hot (in memory) and cold wallet.
      Automated withdrawals and validator set updates are approved by 2/3 of the validator power,
      signed by hot wallets.
      For additional security, withdrawals and validator set updates are pending for a dispute period.
      During this period, any locker may lock the bridge (preventing further withdrawals or updates).
      To unlock the bridge, a quorum of cold wallet signatures is required.

    Validator set updates:
      The active validators sign a hash of the new validator set and powers on the L1.
      This contract checks those signatures, and updates the hash of the active validator set.
      The active validators' L1 stake is still locked for at least one more epoch (unbonding period),
      and the new validators will slash the old ones' stake if they do not properly generate the validator set update signatures.
      The validator set change is pending for a period of time for the lockers to dispute the change.

    Withdrawals:
      The validators sign withdrawals on the L1, which are batched and sent to batchedRequestWithdrawals()
      This contract checks the signatures, and then creates a pending withdrawal which can be disputed for a period of time.
      After the dispute period has elapsed (measured in both time and blocks), a second transaction can be sent to finalize the withdrawal and release the USDC.

    Deposits:
      The validators on the L1 listen for and sign DepositEvent events emitted by this contract,
      crediting the L1 with the equivalent USDC. No additional work needs to be done on this contract.

    Signatures:
      For withdrawals and validator set updates, the signatures are sent to the bridge contract
      in the same order as the active validator set, i.e. signing validators should be a subsequence
      of active validators.

    Lockers:
      These addresses are approved by the validators to lock the contract if submitted signatures do not match
      the locker's view of the L1. Once locked, only a quorum of cold wallet validator signatures can unlock the bridge.
      This dispute period is used for both withdrawals and validator set updates.
      L1 operation will automatically register all validator hot addresses as lockers.
      Adding a locker requires hot wallet quorum, and removing requires cold wallet quorum.

    Finalizers:
      These addresses are approved by the validators to finalize withdrawals and validator set updates.
      While not strictly necessary due to the locking mechanism, this adds an additional layer of security without sacrificing functionality.
      Even if locking transactions are censored (which should be economically infeasible), this still requires attackers to control a finalizer private key.
      L1 operation will eventually register all validator hot addresses as finalizers,
      though there may be an intermediate phase where finalizers are a subset of trusted validators.
      Adding a finalizer requires hot wallet quorum, and removing requires cold wallet quorum.

    Unlocking:
      When the bridge is unlocked, a new validator set is atomically set and finalized.
      This is safe because the unlocking message is signed by a quorum of validator cold wallets.

    The L1 will ensure the following, though neither is required by the smart contract:
      1. The order of active validators are ordered in decreasing order of power.
      2. The validators are unique.

    On epoch changes, the L1 will ensure that new signatures are generated for unclaimed withdrawals
    for any validators that have changed.

    This bridge contract assumes there will be 20-30 validators on the L1, so signature sets fit in a single tx.
*/

pragma solidity ^0.8.20;

import "@openzeppelin/contracts/utils/Pausable.sol";
import "@openzeppelin/contracts/utils/ReentrancyGuard.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import "@openzeppelin/contracts/token/ERC20/extensions/IERC20Permit.sol";
import "./Signature.sol";

// EIP-712 type hashes
bytes32 constant REQUEST_WITHDRAWAL_TYPEHASH = keccak256(
    "Withdraw(address user,address destination,address token,uint256 amount,uint256 chainId,uint64 nonce)"
);

bytes32 constant UPDATE_VALIDATOR_SET_TYPEHASH = keccak256(
    "UpdateValidatorSet(uint64 epoch,address[] hotAddresses,address[] coldAddresses,uint64[] powers)"
);

bytes32 constant MODIFY_LOCKER_TYPEHASH = keccak256(
    "ModifyLocker(address locker,bool isLocker,uint64 nonce)"
);

bytes32 constant MODIFY_FINALIZER_TYPEHASH = keccak256(
    "ModifyFinalizer(address finalizer,bool isFinalizer,uint64 nonce)"
);

bytes32 constant CHANGE_DISPUTE_PERIOD_TYPEHASH = keccak256(
    "ChangeDisputePeriod(uint64 newDisputePeriodSeconds,uint64 nonce)"
);

bytes32 constant INVALIDATE_WITHDRAWALS_TYPEHASH = keccak256(
    "InvalidateWithdrawals(bytes32[] messages,uint64 nonce)"
);

bytes32 constant CHANGE_BLOCK_DURATION_TYPEHASH = keccak256(
    "ChangeBlockDuration(uint64 newBlockDurationMillis,uint64 nonce)"
);

bytes32 constant CHANGE_LOCKER_THRESHOLD_TYPEHASH = keccak256(
    "ChangeLockerThreshold(uint64 newLockerThreshold,uint64 nonce)"
);

bytes32 constant UNLOCK_TYPEHASH = keccak256(
    "Unlock(uint64 epoch,address[] hotAddresses,address[] coldAddresses,uint64[] powers,uint64 nonce)"
);

bytes32 constant VALIDATORS_SET_TYPEHASH = keccak256(
    "ValidatorSet(uint64 epoch,address[] validators,uint64[] powers)"
);

struct ValidatorSet {
    uint64 epoch;
    address[] validators;
    uint64[] powers;
}

struct ValidatorSetUpdateRequest {
    uint64 epoch;
    address[] hotAddresses;
    address[] coldAddresses;
    uint64[] powers;
}

struct PendingValidatorSetUpdate {
    uint64 epoch;
    uint64 totalValidatorPower;
    uint64 updateTime;
    uint64 updateBlockNumber;
    uint64 nValidators;
    bytes32 hotValidatorSetHash;
    bytes32 coldValidatorSetHash;
}

struct Withdrawal {
    address user;
    address destination;
    address token;
    uint256 amount;
    uint256 chainId;
    uint64 nonce;
    uint64 requestedTime;
    uint64 requestedBlockNumber;
    bytes32 message;
}

struct WithdrawalRequest {
    address user;
    address destination;
    address token;
    uint256 amount;
    uint256 chainId;
    uint64 nonce;
    Signature[] signatures;
}

struct DepositWithPermit {
    address user;
    address destination;
    address token;
    uint256 amount;
    uint64 deadline;
    Signature signature;
}

contract Bridge is Pausable, ReentrancyGuard {
    using SafeERC20 for IERC20;

    bytes32 public hotValidatorSetHash;
    bytes32 public coldValidatorSetHash;
    PendingValidatorSetUpdate public pendingValidatorSetUpdate;

    mapping(bytes32 => bool) public usedMessages;
    mapping(address => bool) public lockers;
    address[] private lockersVotingLock;
    uint64 public lockerThreshold;

    mapping(address => bool) public finalizers;
    uint64 public epoch;
    uint64 public totalValidatorPower;
    uint64 public disputePeriodSeconds;
    // Need higher resolution than seconds for Arbitrum.
    uint64 public blockDurationMillis;

    mapping(bytes32 => Withdrawal) public requestedWithdrawals;
    mapping(bytes32 => bool) public finalizedWithdrawals;
    mapping(bytes32 => bool) public withdrawalsInvalidated;

    bytes32 public domainSeparator;

    event Deposit(
        address indexed user,
        address destination,
        address token,
        uint256 amount
    );

    event RequestedWithdrawal(
        bytes32 message,
        address indexed user,
        address destination,
        address token,
        uint256 amount,
        uint256 chainId,
        uint64 nonce,
        uint64 requestedTime,
        uint64 requestedBlockNumber
    );

    event FinalizedWithdrawal(
        bytes32 message,
        address indexed user,
        address destination,
        address token,
        uint256 amount,
        uint64 nonce
    );

    event RequestedValidatorSetUpdate(
        uint64 epoch,
        bytes32 hotValidatorSetHash,
        bytes32 coldValidatorSetHash,
        uint64 updateTime
    );

    event FinalizedValidatorSetUpdate(
        uint64 epoch,
        bytes32 hotValidatorSetHash,
        bytes32 coldValidatorSetHash
    );

    event FailedWithdrawal(bytes32 message, uint32 errorCode);
    event ModifiedLocker(address indexed locker, bool isLocker);
    event ModifiedFinalizer(address indexed finalizer, bool isFinalizer);
    event ChangedDisputePeriodSeconds(uint64 newDisputePeriodSeconds);
    event ChangedBlockDurationMillis(uint64 newBlockDurationMillis);
    event ChangedLockerThreshold(uint64 newLockerThreshold);
    event InvalidatedWithdrawal(Withdrawal withdrawal);

    // We could have the deployer initialize separately so that all function args in this file can be calldata.
    // However, calldata does not seem cheaper than memory on Arbitrum, so not a big deal for now.
    constructor(
        address[] memory hotAddresses,
        address[] memory coldAddresses,
        uint64[] memory powers,
        uint64 _disputePeriodSeconds,
        uint64 _blockDurationMillis,
        uint64 _lockerThreshold
    ) {
        domainSeparator = keccak256(
            abi.encode(
                EIP712_DOMAIN_SEPARATOR,
                keccak256("Bridge"),
                keccak256("1"),
                block.chainid,
                address(this)
            )
        );
        totalValidatorPower = checkNewValidatorPowers(powers);

        require(
            hotAddresses.length == coldAddresses.length,
            "Hot and cold validator sets length mismatch"
        );

        ValidatorSet memory hotValidatorSet;
        hotValidatorSet = ValidatorSet({
            epoch: 0,
            validators: hotAddresses,
            powers: powers
        });
        bytes32 newHotValidatorSetHash = makeValidatorSetHash(hotValidatorSet);
        hotValidatorSetHash = newHotValidatorSetHash;

        ValidatorSet memory coldValidatorSet;
        coldValidatorSet = ValidatorSet({
            epoch: 0,
            validators: coldAddresses,
            powers: powers
        });
        bytes32 newColdValidatorSetHash = makeValidatorSetHash(
            coldValidatorSet
        );
        coldValidatorSetHash = newColdValidatorSetHash;

        disputePeriodSeconds = _disputePeriodSeconds;
        blockDurationMillis = _blockDurationMillis;
        lockerThreshold = _lockerThreshold;
        addLockersAndFinalizers(hotAddresses);

        emit RequestedValidatorSetUpdate(
            0,
            hotValidatorSetHash,
            coldValidatorSetHash,
            uint64(block.timestamp)
        );

        pendingValidatorSetUpdate = PendingValidatorSetUpdate({
            epoch: 0,
            totalValidatorPower: totalValidatorPower,
            updateTime: 0,
            updateBlockNumber: getCurBlockNumber(),
            hotValidatorSetHash: hotValidatorSetHash,
            coldValidatorSetHash: coldValidatorSetHash,
            nValidators: uint64(hotAddresses.length)
        });

        emit FinalizedValidatorSetUpdate(
            0,
            hotValidatorSetHash,
            coldValidatorSetHash
        );
    }

    function addLockersAndFinalizers(address[] memory addresses) private {
        uint64 end = uint64(addresses.length);
        for (uint64 idx; idx < end; idx++) {
            address _address = addresses[idx];
            lockers[_address] = true;
            finalizers[_address] = true;
        }
    }

    // A utility function to make a checkpoint of the validator set supplied.
    // The checkpoint is the hash of all the validators, the powers and the epoch.
    function makeValidatorSetHash(
        ValidatorSet memory validatorSet
    ) private pure returns (bytes32) {
        require(
            validatorSet.validators.length == validatorSet.powers.length,
            "Malformed validator set"
        );

        bytes32 checkpoint = keccak256(
            abi.encode(
                VALIDATORS_SET_TYPEHASH,
                validatorSet.epoch,
                keccak256(abi.encodePacked(validatorSet.validators)),
                keccak256(abi.encodePacked(validatorSet.powers))
            )
        );
        return checkpoint;
    }

    function requestWithdrawal(
        address user,
        address destination,
        address token,
        uint256 amount,
        uint256 chainId,
        uint64 nonce,
        ValidatorSet calldata hotValidatorSet,
        Signature[] memory signatures
    ) private {
        bytes32 message = keccak256(
            abi.encode(
                REQUEST_WITHDRAWAL_TYPEHASH,
                user,
                destination,
                token,
                amount,
                chainId,
                nonce
            )
        );
        if (!isValidWithdrawal(message)) {
            emit FailedWithdrawal(message, 5);
            return;
        }
        Withdrawal memory withdrawal = Withdrawal({
            user: user,
            destination: destination,
            token: token,
            amount: amount,
            chainId: chainId,
            nonce: nonce,
            requestedTime: uint64(block.timestamp),
            requestedBlockNumber: getCurBlockNumber(),
            message: message
        });
        if (requestedWithdrawals[message].requestedTime != 0) {
            emit FailedWithdrawal(message, 0);
            return;
        }
        checkValidatorSignatures(
            message,
            hotValidatorSet,
            signatures,
            hotValidatorSetHash
        );
        requestedWithdrawals[message] = withdrawal;
        emit RequestedWithdrawal(
            withdrawal.message,
            withdrawal.user,
            withdrawal.destination,
            withdrawal.token,
            withdrawal.amount,
            withdrawal.chainId,
            withdrawal.nonce,
            withdrawal.requestedTime,
            withdrawal.requestedBlockNumber
        );
    }

    // An external function anyone can call to withdraw amount from the bridge by providing valid signatures
    // from the active L1 validators.
    function batchedRequestWithdrawals(
        WithdrawalRequest[] memory withdrawalRequests,
        ValidatorSet calldata hotValidatorSet
    ) external nonReentrant whenNotPaused {
        uint64 end = uint64(withdrawalRequests.length);
        for (uint64 idx; idx < end; idx++) {
            WithdrawalRequest memory withdrawalRequest = withdrawalRequests[
                idx
            ];
            requestWithdrawal(
                withdrawalRequest.user,
                withdrawalRequest.destination,
                withdrawalRequest.token,
                withdrawalRequest.amount,
                withdrawalRequest.chainId,
                withdrawalRequest.nonce,
                hotValidatorSet,
                withdrawalRequest.signatures
            );
        }
    }

    function finalizeWithdrawal(bytes32 message) private {
        if (!isValidWithdrawal(message)) {
            emit FailedWithdrawal(message, 5);
            return;
        }

        if (finalizedWithdrawals[message]) {
            emit FailedWithdrawal(message, 1);
            return;
        }

        Withdrawal memory withdrawal = requestedWithdrawals[message];
        if (withdrawal.requestedTime == 0) {
            emit FailedWithdrawal(message, 2);
            return;
        }

        uint32 errorCode = getDisputePeriodErrorCode(
            withdrawal.requestedTime,
            withdrawal.requestedBlockNumber
        );

        if (errorCode != 0) {
            emit FailedWithdrawal(message, errorCode);
            return;
        }

        finalizedWithdrawals[message] = true;
        IERC20(withdrawal.token).safeTransfer(
            withdrawal.destination,
            withdrawal.amount
        );
        emit FinalizedWithdrawal(
            withdrawal.message,
            withdrawal.user,
            withdrawal.destination,
            withdrawal.token,
            withdrawal.amount,
            withdrawal.nonce
        );
    }

    function batchedFinalizeWithdrawals(
        bytes32[] calldata messages
    ) external nonReentrant whenNotPaused {
        checkFinalizer(msg.sender);
        uint64 end = uint64(messages.length);
        for (uint64 idx; idx < end; idx++) {
            finalizeWithdrawal(messages[idx]);
        }
    }

    function isValidWithdrawal(bytes32 message) private view returns (bool) {
        return !withdrawalsInvalidated[message];
    }

    function getCurBlockNumber() private view returns (uint64) {
        return uint64(block.number);
    }

    // Returns 0 if no error
    function getDisputePeriodErrorCode(
        uint64 time,
        uint64 blockNumber
    ) private view returns (uint32) {
        bool enoughTimePassed = block.timestamp > time + disputePeriodSeconds;
        if (!enoughTimePassed) {
            return 3;
        }

        uint64 curBlockNumber = getCurBlockNumber();

        bool enoughBlocksPassed = (curBlockNumber - blockNumber) *
            blockDurationMillis >
            1000 * disputePeriodSeconds;
        if (!enoughBlocksPassed) {
            return 4;
        }

        return 0;
    }

    // Utility function that verifies the signatures supplied and checks that the validators have reached quorum.
    function checkValidatorSignatures(
        bytes32 message,
        ValidatorSet memory activeValidatorSet, // Active set of all L1 validators
        Signature[] memory signatures,
        bytes32 validatorSetHash
    ) private view {
        require(
            makeValidatorSetHash(activeValidatorSet) == validatorSetHash,
            "Supplied active validators and powers do not match the active checkpoint"
        );

        uint64 nSignatures = uint64(signatures.length);
        require(nSignatures > 0, "Signers empty");
        uint64 cumulativePower;
        uint64 signatureIdx;
        uint64 end = uint64(activeValidatorSet.validators.length);

        for (
            uint64 activeValidatorSetIdx;
            activeValidatorSetIdx < end;
            activeValidatorSetIdx++
        ) {
            address signer = SignatureLib.recoverSigner(
                message,
                signatures[signatureIdx],
                domainSeparator
            );
            if (
                signer == activeValidatorSet.validators[activeValidatorSetIdx]
            ) {
                uint64 power = activeValidatorSet.powers[activeValidatorSetIdx];
                cumulativePower += power;

                if (3 * cumulativePower > 2 * totalValidatorPower) {
                    break;
                }

                signatureIdx += 1;
                if (signatureIdx >= nSignatures) {
                    break;
                }
            }
        }

        require(
            3 * cumulativePower > 2 * totalValidatorPower,
            "Submitted validator set signatures do not have enough power"
        );
    }

    function checkMessageNotUsed(bytes32 message) private {
        require(!usedMessages[message], "message already used");
        usedMessages[message] = true;
    }

    // This function updates the validator set by checking that the active validators have signed
    // off on the new validator set
    function updateValidatorSet(
        ValidatorSetUpdateRequest memory newValidatorSet,
        ValidatorSet memory activeHotValidatorSet,
        Signature[] memory signatures
    ) external whenNotPaused {
        require(
            makeValidatorSetHash(activeHotValidatorSet) == hotValidatorSetHash,
            "Supplied active validators and powers do not match checkpoint"
        );

        bytes32 message = keccak256(
            abi.encode(
                UPDATE_VALIDATOR_SET_TYPEHASH,
                newValidatorSet.epoch,
                keccak256(abi.encodePacked(newValidatorSet.hotAddresses)),
                keccak256(abi.encodePacked(newValidatorSet.coldAddresses)),
                keccak256(abi.encodePacked(newValidatorSet.powers))
            )
        );

        updateValidatorSetInner(
            newValidatorSet,
            activeHotValidatorSet,
            signatures,
            message,
            false
        );
    }

    function updateValidatorSetInner(
        ValidatorSetUpdateRequest memory newValidatorSet,
        ValidatorSet memory activeValidatorSet,
        Signature[] memory signatures,
        bytes32 message,
        bool useColdValidatorSet
    ) private {
        require(
            newValidatorSet.hotAddresses.length ==
                newValidatorSet.coldAddresses.length,
            "New hot and cold validator sets length mismatch"
        );

        require(
            newValidatorSet.hotAddresses.length ==
                newValidatorSet.powers.length,
            "New validator set and powers length mismatch"
        );

        require(
            newValidatorSet.epoch > activeValidatorSet.epoch,
            "New validator set epoch must be greater than the active epoch"
        );

        uint64 newTotalValidatorPower = checkNewValidatorPowers(
            newValidatorSet.powers
        );

        bytes32 validatorSetHash;
        if (useColdValidatorSet) {
            validatorSetHash = coldValidatorSetHash;
        } else {
            validatorSetHash = hotValidatorSetHash;
        }

        checkValidatorSignatures(
            message,
            activeValidatorSet,
            signatures,
            validatorSetHash
        );

        ValidatorSet memory newHotValidatorSet;
        newHotValidatorSet = ValidatorSet({
            epoch: newValidatorSet.epoch,
            validators: newValidatorSet.hotAddresses,
            powers: newValidatorSet.powers
        });
        bytes32 newHotValidatorSetHash = makeValidatorSetHash(
            newHotValidatorSet
        );

        ValidatorSet memory newColdValidatorSet;
        newColdValidatorSet = ValidatorSet({
            epoch: newValidatorSet.epoch,
            validators: newValidatorSet.coldAddresses,
            powers: newValidatorSet.powers
        });
        bytes32 newColdValidatorSetHash = makeValidatorSetHash(
            newColdValidatorSet
        );

        uint64 updateTime = uint64(block.timestamp);
        pendingValidatorSetUpdate = PendingValidatorSetUpdate({
            epoch: newValidatorSet.epoch,
            totalValidatorPower: newTotalValidatorPower,
            updateTime: updateTime,
            updateBlockNumber: getCurBlockNumber(),
            hotValidatorSetHash: newHotValidatorSetHash,
            coldValidatorSetHash: newColdValidatorSetHash,
            nValidators: uint64(newHotValidatorSet.validators.length)
        });

        emit RequestedValidatorSetUpdate(
            newValidatorSet.epoch,
            newHotValidatorSetHash,
            newColdValidatorSetHash,
            updateTime
        );
    }

    function finalizeValidatorSetUpdate() external nonReentrant whenNotPaused {
        checkFinalizer(msg.sender);

        require(
            pendingValidatorSetUpdate.updateTime != 0,
            "Pending validator set update already finalized"
        );

        uint32 errorCode = getDisputePeriodErrorCode(
            pendingValidatorSetUpdate.updateTime,
            pendingValidatorSetUpdate.updateBlockNumber
        );
        require(errorCode == 0, "Still in dispute period");

        finalizeValidatorSetUpdateInner();
    }

    function finalizeValidatorSetUpdateInner() private {
        hotValidatorSetHash = pendingValidatorSetUpdate.hotValidatorSetHash;
        coldValidatorSetHash = pendingValidatorSetUpdate.coldValidatorSetHash;
        epoch = pendingValidatorSetUpdate.epoch;
        totalValidatorPower = pendingValidatorSetUpdate.totalValidatorPower;
        pendingValidatorSetUpdate.updateTime = 0;

        emit FinalizedValidatorSetUpdate(
            epoch,
            pendingValidatorSetUpdate.hotValidatorSetHash,
            pendingValidatorSetUpdate.coldValidatorSetHash
        );
    }

    function modifyLocker(
        address locker,
        bool _isLocker,
        uint64 nonce,
        ValidatorSet calldata activeValidatorSet,
        Signature[] memory signatures
    ) external {
        bytes32 message = keccak256(
            abi.encode(MODIFY_LOCKER_TYPEHASH, locker, _isLocker, nonce)
        );

        bytes32 validatorSetHash;
        if (_isLocker) {
            validatorSetHash = hotValidatorSetHash;
        } else {
            validatorSetHash = coldValidatorSetHash;
        }

        checkMessageNotUsed(message);
        checkValidatorSignatures(
            message,
            activeValidatorSet,
            signatures,
            validatorSetHash
        );
        if (lockers[locker] && !_isLocker && !paused()) {
            removeLockerVote(locker);
        }
        lockers[locker] = _isLocker;
        emit ModifiedLocker(locker, _isLocker);
    }

    function modifyFinalizer(
        address finalizer,
        bool _isFinalizer,
        uint64 nonce,
        ValidatorSet calldata activeValidatorSet,
        Signature[] memory signatures
    ) external {
        bytes32 message = keccak256(
            abi.encode(
                MODIFY_FINALIZER_TYPEHASH,
                finalizer,
                _isFinalizer,
                nonce
            )
        );

        bytes32 validatorSetHash;
        if (_isFinalizer) {
            validatorSetHash = hotValidatorSetHash;
        } else {
            validatorSetHash = coldValidatorSetHash;
        }

        checkMessageNotUsed(message);
        checkValidatorSignatures(
            message,
            activeValidatorSet,
            signatures,
            validatorSetHash
        );
        finalizers[finalizer] = _isFinalizer;
        emit ModifiedFinalizer(finalizer, _isFinalizer);
    }

    function checkFinalizer(address finalizer) private view {
        require(finalizers[finalizer], "Sender is not a finalizer");
    }

    // This function checks that the total power of the new validator set is greater than zero.
    function checkNewValidatorPowers(
        uint64[] memory powers
    ) private pure returns (uint64) {
        uint64 cumulativePower;
        for (uint64 i; i < powers.length; i++) {
            cumulativePower += powers[i];
        }

        require(
            cumulativePower > 0,
            "Submitted validator powers must be greater than zero"
        );
        return cumulativePower;
    }

    function changeDisputePeriodSeconds(
        uint64 newDisputePeriodSeconds,
        uint64 nonce,
        ValidatorSet memory activeColdValidatorSet,
        Signature[] memory signatures
    ) external {
        bytes32 message = keccak256(
            abi.encode(
                CHANGE_DISPUTE_PERIOD_TYPEHASH,
                newDisputePeriodSeconds,
                nonce
            )
        );
        checkMessageNotUsed(message);
        checkValidatorSignatures(
            message,
            activeColdValidatorSet,
            signatures,
            coldValidatorSetHash
        );

        disputePeriodSeconds = newDisputePeriodSeconds;
        emit ChangedDisputePeriodSeconds(newDisputePeriodSeconds);
    }

    function invalidateWithdrawals(
        bytes32[] memory messages,
        uint64 nonce,
        ValidatorSet memory activeColdValidatorSet,
        Signature[] memory signatures
    ) external {
        bytes32 message = keccak256(
            abi.encode(
                INVALIDATE_WITHDRAWALS_TYPEHASH,
                keccak256(abi.encodePacked(messages)),
                nonce
            )
        );

        checkMessageNotUsed(message);
        checkValidatorSignatures(
            message,
            activeColdValidatorSet,
            signatures,
            coldValidatorSetHash
        );

        uint64 end = uint64(messages.length);
        for (uint64 idx; idx < end; idx++) {
            withdrawalsInvalidated[messages[idx]] = true;
            emit InvalidatedWithdrawal(requestedWithdrawals[messages[idx]]);
        }
    }

    function changeBlockDurationMillis(
        uint64 newBlockDurationMillis,
        uint64 nonce,
        ValidatorSet memory activeColdValidatorSet,
        Signature[] memory signatures
    ) external {
        bytes32 message = keccak256(
            abi.encode(
                CHANGE_BLOCK_DURATION_TYPEHASH,
                newBlockDurationMillis,
                nonce
            )
        );

        checkMessageNotUsed(message);
        checkValidatorSignatures(
            message,
            activeColdValidatorSet,
            signatures,
            coldValidatorSetHash
        );

        blockDurationMillis = newBlockDurationMillis;
        emit ChangedBlockDurationMillis(newBlockDurationMillis);
    }

    function changeLockerThreshold(
        uint64 newLockerThreshold,
        uint64 nonce,
        ValidatorSet memory activeColdValidatorSet,
        Signature[] memory signatures
    ) external {
        bytes32 message = keccak256(
            abi.encode(
                CHANGE_LOCKER_THRESHOLD_TYPEHASH,
                newLockerThreshold,
                nonce
            )
        );

        checkMessageNotUsed(message);
        checkValidatorSignatures(
            message,
            activeColdValidatorSet,
            signatures,
            coldValidatorSetHash
        );

        lockerThreshold = newLockerThreshold;
        if (uint64(lockersVotingLock.length) >= lockerThreshold && !paused()) {
            _pause();
        }
        emit ChangedLockerThreshold(newLockerThreshold);
    }

    function getLockersVotingLock() external view returns (address[] memory) {
        return lockersVotingLock;
    }

    function isVotingLock(address locker) public view returns (bool) {
        uint64 length = uint64(lockersVotingLock.length);
        for (uint64 i = 0; i < length; i++) {
            if (lockersVotingLock[i] == locker) {
                return true;
            }
        }
        return false;
    }

    function voteEmergencyLock() external {
        require(
            lockers[msg.sender],
            "Sender is not authorized to lock smart contract"
        );
        require(
            !isVotingLock(msg.sender),
            "Locker already voted for emergency lock"
        );
        lockersVotingLock.push(msg.sender);
        if (uint64(lockersVotingLock.length) >= lockerThreshold && !paused()) {
            _pause();
        }
    }

    function unvoteEmergencyLock() external whenNotPaused {
        require(
            lockers[msg.sender],
            "Sender is not authorized to lock smart contract"
        );
        require(
            isVotingLock(msg.sender),
            "Locker is not currently voting for emergency lock"
        );
        removeLockerVote(msg.sender);
    }

    function removeLockerVote(address locker) private whenNotPaused {
        require(
            lockers[locker],
            "Address is not authorized to lock smart contract"
        );
        uint64 length = uint64(lockersVotingLock.length);
        for (uint64 i = 0; i < length; i++) {
            if (lockersVotingLock[i] == locker) {
                lockersVotingLock[i] = lockersVotingLock[length - 1];
                lockersVotingLock.pop();
                break;
            }
        }
    }

    function emergencyUnlock(
        ValidatorSetUpdateRequest memory newValidatorSet,
        ValidatorSet calldata activeColdValidatorSet,
        Signature[] calldata signatures,
        uint64 nonce
    ) external whenPaused {
        bytes32 message = keccak256(
            abi.encode(
                UNLOCK_TYPEHASH,
                newValidatorSet.epoch,
                keccak256(abi.encodePacked(newValidatorSet.hotAddresses)),
                keccak256(abi.encodePacked(newValidatorSet.coldAddresses)),
                keccak256(abi.encodePacked(newValidatorSet.powers)),
                nonce
            )
        );

        checkMessageNotUsed(message);
        updateValidatorSetInner(
            newValidatorSet,
            activeColdValidatorSet,
            signatures,
            message,
            true
        );
        finalizeValidatorSetUpdateInner();
        delete lockersVotingLock;
        _unpause();
    }

    function depositWithPermit(
        address user,
        address destination,
        address token,
        uint256 amount,
        uint64 deadline,
        Signature memory signature
    ) private {
        address spender = address(this);
        IERC20Permit(token).permit(
            user,
            spender,
            amount,
            deadline,
            signature.v,
            bytes32(signature.r),
            bytes32(signature.s)
        );

        IERC20(token).safeTransferFrom(user, spender, amount);
        if (destination != user) {
            emit Deposit(user, destination, token, amount);
        }
    }

    function batchedDepositWithPermit(
        DepositWithPermit[] memory deposits
    ) external nonReentrant whenNotPaused {
        uint64 end = uint64(deposits.length);
        for (uint64 idx; idx < end; idx++) {
            depositWithPermit(
                deposits[idx].user,
                deposits[idx].destination,
                deposits[idx].token,
                deposits[idx].amount,
                deposits[idx].deadline,
                deposits[idx].signature
            );
        }
    }

    function deposit(
        address destination,
        address token,
        uint256 amount
    ) external nonReentrant whenNotPaused {
        require(amount > 0, "Deposit amount must be greater than zero");
        require(destination != address(0), "Invalid destination address");

        IERC20(token).safeTransferFrom(msg.sender, address(this), amount);
        if (destination != msg.sender) {
            emit Deposit(msg.sender, destination, token, amount);
        }
    }
}
