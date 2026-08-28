// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.22;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/access/AccessControlUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/utils/PausableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/utils/ReentrancyGuardUpgradeable.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import "@openzeppelin/contracts/token/ERC20/extensions/IERC20Permit.sol";
import "@openzeppelin/contracts/token/ERC20/extensions/IERC20Metadata.sol";
import "./Signature.sol";

struct DepositConfirm {
    bytes32 user;
    address destination;
    bytes32 token;
    uint256 amount;
    uint256 chainId;
    uint64 blockNumber;
    bytes32 txHash;
    uint32 index;
    Signature signature;
}

struct WithdrawRequest {
    address user;
    address destination;
    address token;
    uint256 amount;
    uint256 chainId;
    uint64 nonce;
}

struct WithdrawConfirm {
    address user;
    address destination;
    address token;
    uint256 amount;
    uint256 chainId;
    uint64 nonce;
    Signature signature;
}

struct ValidatorSetUpdate {
    uint64 epoch;
    address[] hotAddresses;
    address[] coldAddresses;
    uint64[] powers;
}

struct WithdrawWithPermit {
    address user;
    address destination;
    address token;
    uint256 amount;
    uint256 chainId;
    uint64 deadline;
    Signature signature;
}

struct CrossChainMessage {
    bytes32 domainSeparator;
    Signature signature;
    bytes messageRawData;
}

struct MessageSignature {
    uint64 totalPower;
    mapping(address => Signature) signatures;
    bytes rawData;
}

// Minimal interface for bridged tokens to mint/burn
interface IBridgedToken {
    function mint(address to, uint256 amount) external;
    function burn(uint256 amount) external;
}

// EIP-712 type hash for deposit message
bytes32 constant DEPOSIT_TYPEHASH = keccak256(
    "Deposit(bytes32 user,address destination,bytes32 token,uint256 amount,uint256 chainId,uint64 blockNumber,bytes32 txHash,uint32 index)"
);

// EIP-712 type hash for withdraw message
bytes32 constant WITHDRAW_TYPEHASH = keccak256(
    "Withdraw(address user,address destination,address token,uint256 amount,uint256 chainId,uint64 nonce)"
);

// EIP-712 type hash for update validator set message
bytes32 constant UPDATE_VALIDATOR_SET_TYPEHASH = keccak256(
    "UpdateValidatorSet(uint64 epoch,address[] hotAddresses,address[] coldAddresses,uint64[] powers)"
);

contract BridgeHub is
    Initializable,
    UUPSUpgradeable,
    PausableUpgradeable,
    ReentrancyGuardUpgradeable,
    AccessControlUpgradeable
{
    using SafeERC20 for IERC20;

    bytes32 public constant ADMIN_ROLE = keccak256("ADMIN_ROLE");
    bytes32 public constant PAUSER_ROLE = keccak256("PAUSER_ROLE");

    // Deposit: srcToken => dstToken (hub). Withdraw: dstToken => srcToken.
    mapping(uint256 => mapping(bytes32 => bytes32)) public tokenPair;
    // srcToken key stores (srcDecimal - dstDecimal); dstToken key stores the inverse.
    mapping(uint256 => mapping(bytes32 => int8)) public tokenDecimalDiff;

    // Validator membership and power
    address[] public coldValidatorList;
    address[] public hotValidatorList;
    mapping(address => uint64) public validatorPowers;
    uint64 public totalValidatorPower;
    uint64 public epoch;
    ValidatorSetUpdate private pendingValidatorSetUpdate;

    // Global deposit nonce
    uint64 public depositNonce;
    // Global withdraw nonce
    uint64 public withdrawNonce;
    // per-token withdraw fee
    mapping(address => uint256) public tokenWithdrawFee;

    // Message info for signature recovery
    mapping(bytes32 => MessageSignature) private bridgeMessageSignature;
    // Message info for signature recovery
    mapping(bytes32 => MessageSignature) private messageSignature;
    // Processed message record to prevent replay
    mapping(bytes32 => bool) public processedMessages;
    // Pending messages waiting for validator signatures
    bytes32[] public pendingMessages;

    // Expected EIP-712 domain separator for this contract (name="BridgeHub", version="1")
    bytes32 public domainSeparator;

    // Reserve storage slots for future upgrades
    uint256[49] private __gap;

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
        address destination,
        address indexed token,
        uint256 amount,
        uint256 chainId,
        uint64 nonce
    );
    event WithdrawCompleted(uint256 indexed nonce, bytes32 indexed message);

    event WithdrawFeeSet(address indexed token, uint256 fee);
    event FeesClaimed(address indexed token, address to, uint256 amount);

    event BridgeSignatureSubmitted(
        bytes32 indexed message,
        address signer,
        uint64 totalPower
    );
    event MessageStorageCleared(bytes32 indexed message);

    event TokenPairSet(
        uint256 indexed chainId,
        bytes32 indexed srcToken,
        address indexed dstToken,
        uint8 srcTokenDecimal,
        uint8 dstTokenDecimal
    );

    event RequestedValidatorSetUpdate(
        uint64 indexed newEpoch,
        address[] hotAddresses,
        address[] coldAddresses,
        uint64[] powers
    );

    event RemovedValidatorSet(
        uint64 indexed oldEpoch,
        address[] hotAddresses,
        address[] coldAddresses,
        uint64[] powers
    );

    event FinalizedValidatorSetUpdate(
        uint64 indexed newEpoch,
        address[] hotAddresses,
        address[] coldAddresses,
        uint64[] powers
    );

    /// @custom:oz-upgrades-unsafe-allow constructor
    constructor() {
        _disableInitializers();
    }

    modifier onlyValidator() {
        _checkValidator(msg.sender);
        _;
    }

    function _authorizeUpgrade(
        address newImplementation
    ) internal override onlyRole(ADMIN_ROLE) {}

    function initialize(
        address[] memory hotAddresses,
        address[] memory coldAddresses,
        uint64[] memory powers
    ) public initializer {
        __UUPSUpgradeable_init();
        __AccessControl_init();
        __Pausable_init();
        __ReentrancyGuard_init();

        _grantRole(DEFAULT_ADMIN_ROLE, msg.sender);
        _grantRole(ADMIN_ROLE, msg.sender);
        _grantRole(PAUSER_ROLE, msg.sender);

        _updateValidatorSet(0, hotAddresses, coldAddresses, powers);
        _setValidators(0, hotAddresses, coldAddresses, powers);

        // Compute and cache domain separator bound to this proxy address
        domainSeparator = keccak256(
            abi.encode(
                EIP712_DOMAIN_SEPARATOR,
                keccak256(bytes("BridgeHub")),
                keccak256(bytes("1")),
                block.chainid,
                address(this)
            )
        );
    }

    function setWithdrawFee(
        address token,
        uint256 fee
    ) external onlyRole(ADMIN_ROLE) {
        require(token != address(0), "Invalid token address");
        tokenWithdrawFee[token] = fee;
        emit WithdrawFeeSet(token, fee);
    }

    function setTokenPair(
        uint256 chainId,
        bytes32 srcToken,
        uint8 srcTokenDecimal,
        address dstToken
    ) external onlyRole(ADMIN_ROLE) {
        require(chainId != 0, "Invalid chainId");
        require(srcToken != bytes32(0), "Invalid src token");
        require(dstToken != address(0), "Invalid dst token");

        uint8 dstTokenDecimal = IERC20Metadata(dstToken).decimals();
        int8 decimalDiff = int8(srcTokenDecimal) - int8(dstTokenDecimal);
        bytes32 dstTokenBytes32 = _addressToBytes32(dstToken);

        tokenPair[chainId][srcToken] = dstTokenBytes32;
        tokenDecimalDiff[chainId][srcToken] = decimalDiff;

        if (_isEvmTokenAddress(srcToken)) {
            tokenPair[chainId][dstTokenBytes32] = srcToken;
            tokenDecimalDiff[chainId][dstTokenBytes32] = -decimalDiff;
        }

        emit TokenPairSet(
            chainId,
            srcToken,
            dstToken,
            srcTokenDecimal,
            dstTokenDecimal
        );
    }

    function pause() external onlyRole(PAUSER_ROLE) {
        _pause();
    }

    function unpause() external onlyRole(PAUSER_ROLE) {
        _unpause();
    }

    function claimFees(
        address token,
        address to,
        uint256 amount
    ) external onlyRole(ADMIN_ROLE) {
        require(to != address(0), "Invalid recipient address");
        require(amount > 0, "Amount must be positive");

        IERC20(token).safeTransfer(to, amount);
        emit FeesClaimed(token, to, amount);
    }

    function clearMessageStorage(
        bytes32 message
    ) external onlyRole(ADMIN_ROLE) {
        _clearMessageStorage(message);
    }

    function updateValidatorSet(
        uint64 newEpoch,
        address[] calldata hotAddresses,
        address[] calldata coldAddresses,
        uint64[] calldata powers
    ) external onlyRole(ADMIN_ROLE) {
        require(newEpoch > 0, "New epoch must be positive");
        _updateValidatorSet(newEpoch, hotAddresses, coldAddresses, powers);
    }

    function _updateValidatorSet(
        uint64 newEpoch,
        address[] memory hotAddresses,
        address[] memory coldAddresses,
        uint64[] memory powers
    ) private {
        // Validate input parameters
        require(hotAddresses.length > 0, "Empty hot addresses array");
        require(
            hotAddresses.length == coldAddresses.length,
            "Hot and cold addresses length mismatch"
        );
        require(
            hotAddresses.length == powers.length,
            "Hot addresses and powers length mismatch"
        );
        // skip validation for newEpoch == 0, as it is allowed for initial setup
        if (epoch != 0 || newEpoch != 0) {
            require(
                newEpoch > epoch,
                "New epoch must be greater than current epoch"
            );
        }

        // Validate addresses, powers, check duplicates, and calculate total power in one loop
        uint64 totalPower;

        // Validate hot addresses and cold addresses
        for (uint256 i = 0; i < hotAddresses.length; i++) {
            address hotAddr = hotAddresses[i];
            address coldAddr = coldAddresses[i];
            uint64 power = powers[i];

            // Validate individual addresses and power
            require(hotAddr != address(0), "Zero address in hot addresses");
            require(coldAddr != address(0), "Zero address in cold addresses");
            require(power > 0, "Power must be positive");
            require(hotAddr != coldAddr, "Hot and cold addresses must differ");

            // Check for duplicate addresses
            for (uint256 j = 0; j < i; j++) {
                require(hotAddr != hotAddresses[j], "Duplicate hot addresses");
                require(
                    coldAddr != coldAddresses[j],
                    "Duplicate cold addresses"
                );
            }

            totalPower += power;
        }
        require(totalPower > 0, "Total power must be positive");

        bytes32 message = makeUpdateValidatorSetMessage(
            newEpoch,
            hotAddresses,
            coldAddresses,
            powers
        );

        if (newEpoch != 0) {
            // Add to pending messages
            pendingMessages.push(message);
        }

        pendingValidatorSetUpdate = ValidatorSetUpdate({
            epoch: newEpoch,
            hotAddresses: hotAddresses,
            coldAddresses: coldAddresses,
            powers: powers
        });

        emit RequestedValidatorSetUpdate(
            newEpoch,
            hotAddresses,
            coldAddresses,
            powers
        );
    }

    function _setValidators(
        uint64 newEpoch,
        address[] memory hotAddresses,
        address[] memory coldAddresses,
        uint64[] memory powers
    ) private {
        uint64[] memory oldPowers = new uint64[](hotValidatorList.length);
        // Clear old validators
        for (uint256 i = 0; i < hotValidatorList.length; i++) {
            address oldAddr = hotValidatorList[i];
            if (validatorPowers[oldAddr] > 0) {
                oldPowers[i] = validatorPowers[oldAddr];
                validatorPowers[oldAddr] = 0;
            }
        }
        emit RemovedValidatorSet(
            epoch,
            hotValidatorList,
            coldValidatorList,
            oldPowers
        );

        // Update validator lists
        coldValidatorList = coldAddresses;
        hotValidatorList = hotAddresses;

        // Set new validators and compute total power
        uint64 total;
        for (uint256 i = 0; i < hotAddresses.length; i++) {
            address addr = hotAddresses[i];
            uint64 power = powers[i];
            validatorPowers[addr] = power;
            total += power;
        }
        totalValidatorPower = total;
        epoch = newEpoch;

        pendingValidatorSetUpdate = ValidatorSetUpdate({
            epoch: 0,
            hotAddresses: new address[](0),
            coldAddresses: new address[](0),
            powers: new uint64[](0)
        });

        emit FinalizedValidatorSetUpdate(
            newEpoch,
            hotValidatorList,
            coldValidatorList,
            powers
        );
    }

    function getValidators()
        external
        view
        returns (
            uint64 _epoch,
            address[] memory validators,
            uint64[] memory powers
        )
    {
        _epoch = epoch;
        validators = hotValidatorList;
        powers = new uint64[](validators.length);
        for (uint256 i = 0; i < validators.length; i++) {
            powers[i] = validatorPowers[validators[i]];
        }
    }

    function getHotValidators() external view returns (address[] memory) {
        return hotValidatorList;
    }

    function getColdValidators() external view returns (address[] memory) {
        return coldValidatorList;
    }

    function getValidatorSignature(
        bytes32 message,
        address signer
    ) external view returns (Signature memory) {
        return messageSignature[message].signatures[signer];
    }

    function getPendingMessages() external view returns (bytes32[] memory) {
        return pendingMessages;
    }

    function getPendingValidatorSetUpdate()
        external
        view
        returns (
            uint64 _epoch,
            address[] memory hotAddresses,
            address[] memory coldAddresses,
            uint64[] memory powers
        )
    {
        ValidatorSetUpdate memory update = pendingValidatorSetUpdate;
        return (
            update.epoch,
            update.hotAddresses,
            update.coldAddresses,
            update.powers
        );
    }

    function _getOrderedSignatures(
        MessageSignature storage info
    )
        private
        view
        returns (address[] memory signers, Signature[] memory signatures)
    {
        uint256 validatorCount = hotValidatorList.length;
        address[] memory orderedSigners = new address[](validatorCount);
        Signature[] memory orderedSignatures = new Signature[](validatorCount);
        uint256 validSignerCount = 0;

        for (uint256 i = 0; i < validatorCount; i++) {
            address validator = hotValidatorList[i];
            Signature memory sig = info.signatures[validator];

            if (sig.r != 0) {
                orderedSigners[validSignerCount] = validator;
                orderedSignatures[validSignerCount] = sig;
                validSignerCount++;
            }
        }

        assembly {
            mstore(orderedSigners, validSignerCount)
            mstore(orderedSignatures, validSignerCount)
        }

        return (orderedSigners, orderedSignatures);
    }

    function getMessageSignatures(
        bytes32 message
    )
        external
        view
        returns (
            uint64 totalPower,
            address[] memory signers,
            Signature[] memory signatures
        )
    {
        MessageSignature storage info = messageSignature[message];
        (signers, signatures) = _getOrderedSignatures(info);
        return (info.totalPower, signers, signatures);
    }

    function getBridgeValidatorSignature(
        bytes32 message,
        address signer
    ) external view returns (Signature memory) {
        return bridgeMessageSignature[message].signatures[signer];
    }

    function getBridgeMessageSignatures(
        bytes32 message
    )
        external
        view
        returns (
            uint64 totalPower,
            address[] memory signers,
            Signature[] memory signatures,
            bytes memory rawData
        )
    {
        MessageSignature storage info = bridgeMessageSignature[message];
        (signers, signatures) = _getOrderedSignatures(info);
        return (info.totalPower, signers, signatures, info.rawData);
    }

    function makeDepositMessage(
        bytes32 user,
        address destination,
        bytes32 token,
        uint256 amount,
        uint256 chainId,
        uint64 blockNumber,
        bytes32 txHash,
        uint32 index
    ) public pure returns (bytes32) {
        return
            keccak256(
                abi.encode(
                    DEPOSIT_TYPEHASH,
                    user,
                    destination,
                    token,
                    amount,
                    chainId,
                    blockNumber,
                    txHash,
                    index
                )
            );
    }

    function makeUpdateValidatorSetMessage(
        uint64 newEpoch,
        address[] memory hotAddresses,
        address[] memory coldAddresses,
        uint64[] memory powers
    ) public pure returns (bytes32) {
        return
            keccak256(
                abi.encode(
                    UPDATE_VALIDATOR_SET_TYPEHASH,
                    newEpoch,
                    keccak256(abi.encodePacked(hotAddresses)),
                    keccak256(abi.encodePacked(coldAddresses)),
                    keccak256(abi.encodePacked(powers))
                )
            );
    }

    function makeWithdrawMessage(
        address user,
        address destination,
        address token,
        uint256 amount,
        uint256 chainId,
        uint64 nonce
    ) public pure returns (bytes32) {
        return
            keccak256(
                abi.encode(
                    WITHDRAW_TYPEHASH,
                    user,
                    destination,
                    token,
                    amount,
                    chainId,
                    nonce
                )
            );
    }

    function _processSignature(
        bytes32 message,
        Signature memory signature
    ) private returns (bool thresholdReached) {
        require(!processedMessages[message], "Already processed");

        // validate domain separator and recover signer
        address signer = SignatureLib.recoverSigner(
            message,
            signature,
            domainSeparator
        );
        _checkValidator(signer);

        // Check if this validator already signed this message
        require(
            messageSignature[message].signatures[signer].r == 0, // Check if signature exists
            "Validator already signed this message"
        );

        // Record signature
        messageSignature[message].signatures[signer] = signature;

        uint64 newPower = messageSignature[message].totalPower +
            validatorPowers[signer];
        messageSignature[message].totalPower = newPower;

        // Check if threshold reached
        if (3 * newPower > 2 * totalValidatorPower) {
            processedMessages[message] = true;
            return true;
        }
        return false;
    }

    function _checkTokenPair(bytes32 token) private pure {
        require(token != bytes32(0), "Token not found");
    }

    function _addressToBytes32(address token) private pure returns (bytes32) {
        return bytes32(uint256(uint160(token)));
    }

    function _bytes32ToAddress(bytes32 token) private pure returns (address) {
        return address(uint160(uint256(token)));
    }

    function _isEvmTokenAddress(bytes32 token) private pure returns (bool) {
        return uint256(token) <= type(uint160).max;
    }

    function _convertAmount(
        uint256 amount,
        int8 decimalDiff
    ) private pure returns (uint256) {
        if (decimalDiff == 0) return amount;
        if (decimalDiff > 0) {
            // Source has more decimals than destination, need to scale down
            return amount / (10 ** uint8(decimalDiff));
        } else {
            // Source has fewer decimals than destination, need to scale up
            return amount * (10 ** uint8(-decimalDiff));
        }
    }

    function _checkValidator(address signer) private view {
        require(validatorPowers[signer] > 0, "Signer is not a validator");
    }

    function _checkPendingMessage(
        bytes32 message
    ) private view returns (uint256) {
        uint256 index = _getPendingMessageIndex(message);
        require(index > 0, "Message not found");
        return index;
    }

    function _getPendingMessageIndex(
        bytes32 message
    ) private view returns (uint256) {
        uint256 length = pendingMessages.length;
        for (uint256 i = 0; i < length; i++) {
            if (pendingMessages[i] == message) {
                return i + 1;
            }
        }
        return 0;
    }

    function _removePendingMessage(uint256 index) private {
        require(index > 0 && index <= pendingMessages.length, "Invalid index");
        uint256 actualIndex = index - 1;

        if (actualIndex < pendingMessages.length - 1) {
            pendingMessages[actualIndex] = pendingMessages[
                pendingMessages.length - 1
            ];
        }
        pendingMessages.pop();
    }

    function _clearMessageStorage(bytes32 message) private {
        delete messageSignature[message];
        delete bridgeMessageSignature[message];
        emit MessageStorageCleared(message);
    }

    function withdraw(
        address destination,
        address token,
        uint256 amount,
        uint256 chainId
    ) external whenNotPaused nonReentrant {
        _withdraw(msg.sender, destination, token, amount, chainId);
    }

    function withdrawBatchWithPermit(
        WithdrawWithPermit[] calldata withdraws
    ) external whenNotPaused nonReentrant {
        require(withdraws.length > 0, "Empty withdraws array");

        uint64 end = uint64(withdraws.length);
        for (uint64 i = 0; i < end; i++) {
            _withdrawWithPermit(withdraws[i]);
        }
    }

    function _withdrawWithPermit(
        WithdrawWithPermit memory withdrawData
    ) private {
        IERC20Permit(withdrawData.token).permit(
            withdrawData.user,
            address(this),
            withdrawData.amount,
            withdrawData.deadline,
            withdrawData.signature.v,
            bytes32(withdrawData.signature.r),
            bytes32(withdrawData.signature.s)
        );

        _withdraw(
            withdrawData.user,
            withdrawData.destination,
            withdrawData.token,
            withdrawData.amount,
            withdrawData.chainId
        );
    }

    function _withdraw(
        address user,
        address destination,
        address token,
        uint256 amount,
        uint256 chainId
    ) private {
        uint256 fee = tokenWithdrawFee[token];
        require(amount > fee, "Amount must exceed fee");
        require(user != address(0), "Invalid user address");
        require(destination != address(0), "Invalid destination address");
        require(token != address(0), "Invalid token address");
        require(chainId != 0, "Invalid chainId");

        bytes32 bridgedTokenBytes32 = _addressToBytes32(token);
        bytes32 srcTokenBytes32 = tokenPair[chainId][bridgedTokenBytes32];
        _checkTokenPair(srcTokenBytes32);
        address bridgeToken = _bytes32ToAddress(srcTokenBytes32);

        // Transfer tokens from user to contract
        IERC20(token).safeTransferFrom(user, address(this), amount);

        amount = amount - fee;

        IBridgedToken(token).burn(amount);

        uint64 nonce = ++withdrawNonce;

        // Convert amount based on decimal difference for cross-chain message
        int8 decimalDiff = tokenDecimalDiff[chainId][bridgedTokenBytes32];
        uint256 messageAmount = _convertAmount(amount, decimalDiff);

        // Create and store pending withdraw message
        bytes32 message = makeWithdrawMessage(
            user,
            destination,
            bridgeToken,
            messageAmount,
            chainId,
            nonce
        );

        // Add to pending messages
        pendingMessages.push(message);

        emit Withdraw(
            message,
            user,
            destination,
            bridgeToken,
            messageAmount,
            chainId,
            nonce
        );
    }

    function depositConfirm(
        DepositConfirm[] calldata deposits
    ) external whenNotPaused nonReentrant onlyValidator {
        require(deposits.length > 0, "Empty deposits array");

        for (uint256 i = 0; i < deposits.length; i++) {
            DepositConfirm calldata data = deposits[i];
            _depositConfirm(data);
        }
    }

    function _depositConfirm(DepositConfirm memory data) private {
        require(data.user != bytes32(0), "Invalid user");
        require(data.destination != address(0), "Invalid destination address");
        require(data.token != bytes32(0), "Invalid token");
        require(data.amount > 0, "Amount must be >0");
        require(data.chainId != 0, "Invalid chainId");

        // Resolve bridged token to mint by (srcChainId, srcToken)
        bytes32 bridgedTokenBytes32 = tokenPair[data.chainId][data.token];
        _checkTokenPair(bridgedTokenBytes32);
        address bridgedToken = _bytes32ToAddress(bridgedTokenBytes32);

        // Build message for signature recovery
        bytes32 message = makeDepositMessage(
            data.user,
            data.destination,
            data.token,
            data.amount,
            data.chainId,
            data.blockNumber,
            data.txHash,
            data.index
        );

        // Process signature and check threshold
        bool thresholdReached = _processSignature(message, data.signature);

        // If threshold reached, execute deposit logic
        if (thresholdReached) {
            // Convert amount based on decimal difference
            int8 decimalDiff = tokenDecimalDiff[data.chainId][data.token];
            uint256 mintAmount = _convertAmount(data.amount, decimalDiff);

            // mint token to user
            IBridgedToken(bridgedToken).mint(data.destination, mintAmount);

            // Increase global deposit nonce
            depositNonce += 1;

            emit Deposit(
                message,
                data.user,
                data.destination,
                data.token,
                data.amount,
                data.chainId,
                data.blockNumber,
                data.txHash,
                data.index,
                depositNonce
            );
        }
    }

    function withdrawConfirm(
        WithdrawConfirm[] calldata withdrawConfirms
    ) external whenNotPaused nonReentrant onlyValidator {
        require(withdrawConfirms.length > 0, "Empty withdraw confirms array");

        for (uint256 i = 0; i < withdrawConfirms.length; i++) {
            WithdrawConfirm calldata data = withdrawConfirms[i];
            _withdrawConfirm(data);
        }
    }

    function _withdrawConfirm(WithdrawConfirm memory data) private {
        // Build message for signature recovery
        bytes32 message = makeWithdrawMessage(
            data.user,
            data.destination,
            data.token,
            data.amount,
            data.chainId,
            data.nonce
        );

        // Verify that this message was requested
        uint256 index = _checkPendingMessage(message);

        // Process signature and check threshold
        bool thresholdReached = _processSignature(message, data.signature);

        // If threshold reached, execute withdraw logic
        if (thresholdReached) {
            _removePendingMessage(index);
            emit WithdrawCompleted(data.nonce, message);
        }
    }

    function updateValidatorSetConfirm(
        ValidatorSetUpdate calldata validatorSet,
        Signature calldata signature
    ) external whenNotPaused nonReentrant onlyValidator {
        bytes32 message = makeUpdateValidatorSetMessage(
            validatorSet.epoch,
            validatorSet.hotAddresses,
            validatorSet.coldAddresses,
            validatorSet.powers
        );

        // Verify that this message was requested
        uint256 index = _checkPendingMessage(message);

        // Process signature and check threshold
        bool thresholdReached = _processSignature(message, signature);

        // If threshold reached, execute validator update logic
        if (thresholdReached) {
            _removePendingMessage(index);
            _setValidators(
                validatorSet.epoch,
                validatorSet.hotAddresses,
                validatorSet.coldAddresses,
                validatorSet.powers
            );
        }
    }

    function submitBridgeSignatures(
        CrossChainMessage[] calldata items
    ) external whenNotPaused nonReentrant onlyValidator {
        require(items.length > 0, "Empty signatures array");

        for (uint256 i = 0; i < items.length; i++) {
            _submitBridgeSignature(
                items[i].signature,
                items[i].domainSeparator,
                items[i].messageRawData
            );
        }
    }

    function _submitBridgeSignature(
        Signature calldata sig,
        bytes32 _domainSeparator,
        bytes calldata messageRawData
    ) private {
        require(_domainSeparator != domainSeparator, "Invalid domain");
        bytes32 message = keccak256(messageRawData);
        // Verify that this message was requested
        _checkPendingMessage(message);
        address signer = SignatureLib.recoverSigner(
            message,
            sig,
            _domainSeparator
        );
        _checkValidator(signer);
        require(
            bridgeMessageSignature[message].signatures[signer].r == 0,
            "Already signed"
        );

        bridgeMessageSignature[message].signatures[signer] = sig;
        bridgeMessageSignature[message].rawData = messageRawData;

        uint64 newTotalPower = bridgeMessageSignature[message].totalPower +
            validatorPowers[signer];
        bridgeMessageSignature[message].totalPower = newTotalPower;

        emit BridgeSignatureSubmitted(message, signer, newTotalPower);
    }
}
