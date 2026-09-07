// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

import (
	"context"
	"errors"
	"math/big"
	"strings"
	"time"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
	_ = time.Tick
	_ = context.Background
)

// CrossChainMessage is an auto generated low-level Go binding around an user-defined struct.
type CrossChainMessage struct {
	DomainSeparator [32]byte
	Signature       Signature
	MessageRawData  []byte
}

// DepositConfirm is an auto generated low-level Go binding around an user-defined struct.
type DepositConfirm struct {
	User        [32]byte
	Destination common.Address
	Token       [32]byte
	Amount      *big.Int
	ChainId     *big.Int
	BlockNumber uint64
	TxHash      [32]byte
	Index       uint32
	Signature   Signature
}


// ValidatorSetUpdate is an auto generated low-level Go binding around an user-defined struct.
type ValidatorSetUpdate struct {
	Epoch         uint64
	HotAddresses  []common.Address
	ColdAddresses []common.Address
	Powers        []uint64
}

// WithdrawConfirm is an auto generated low-level Go binding around an user-defined struct.
type WithdrawConfirm struct {
	User        common.Address
	Destination [32]byte
	Token       [32]byte
	Amount      *big.Int
	ChainId     *big.Int
	Nonce       uint64
	Signature   Signature
}

// WithdrawWithPermit is an auto generated low-level Go binding around an user-defined struct.
type WithdrawWithPermit struct {
	User        common.Address
	Destination [32]byte
	Token       common.Address
	Amount      *big.Int
	ChainId     *big.Int
	Deadline    uint64
	Signature   Signature
}

// BridgeHubMetaData contains all meta data concerning the BridgeHub contract.
var BridgeHubMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AddressInsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EnforcedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExpectedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedInnerCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"message\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"totalPower\",\"type\":\"uint64\"}],\"name\":\"BridgeSignatureSubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"message\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"user\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"destination\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"token\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"blockNumber\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"txHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"message\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"errorCode\",\"type\":\"uint32\"}],\"name\":\"FailedDeposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"FeesClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"newEpoch\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"hotAddresses\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"coldAddresses\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint64[]\",\"name\":\"powers\",\"type\":\"uint64[]\"}],\"name\":\"FinalizedValidatorSetUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"message\",\"type\":\"bytes32\"}],\"name\":\"MessageStorageCleared\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"oldEpoch\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"hotAddresses\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"coldAddresses\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint64[]\",\"name\":\"powers\",\"type\":\"uint64[]\"}],\"name\":\"RemovedValidatorSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"newEpoch\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"hotAddresses\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"coldAddresses\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint64[]\",\"name\":\"powers\",\"type\":\"uint64[]\"}],\"name\":\"RequestedValidatorSetUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"srcToken\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"dstToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"srcTokenDecimal\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"dstTokenDecimal\",\"type\":\"uint8\"}],\"name\":\"TokenPairSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"message\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"destination\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"token\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"message\",\"type\":\"bytes32\"}],\"name\":\"WithdrawCompleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"WithdrawFeeSet\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PAUSER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"claimFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"message\",\"type\":\"bytes32\"}],\"name\":\"clearMessageStorage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"coldValidatorList\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"user\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destination\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"token\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"blockNumber\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"txHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"r\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"s\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"}],\"internalType\":\"structSignature\",\"name\":\"signature\",\"type\":\"tuple\"}],\"internalType\":\"structDepositConfirm[]\",\"name\":\"deposits\",\"type\":\"tuple[]\"}],\"name\":\"depositConfirm\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"depositNonce\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"domainSeparator\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"epoch\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"message\",\"type\":\"bytes32\"}],\"name\":\"getBridgeMessageSignatures\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"totalPower\",\"type\":\"uint64\"},{\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"r\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"s\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"}],\"internalType\":\"structSignature[]\",\"name\":\"signatures\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"rawData\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"message\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"getBridgeValidatorSignature\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"r\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"s\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"}],\"internalType\":\"structSignature\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getColdValidators\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getHotValidators\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"message\",\"type\":\"bytes32\"}],\"name\":\"getMessageSignatures\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"totalPower\",\"type\":\"uint64\"},{\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"r\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"s\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"}],\"internalType\":\"structSignature[]\",\"name\":\"signatures\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPendingMessages\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPendingValidatorSetUpdate\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"_epoch\",\"type\":\"uint64\"},{\"internalType\":\"address[]\",\"name\":\"hotAddresses\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"coldAddresses\",\"type\":\"address[]\"},{\"internalType\":\"uint64[]\",\"name\":\"powers\",\"type\":\"uint64[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"message\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"getValidatorSignature\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"r\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"s\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"}],\"internalType\":\"structSignature\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getValidators\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"_epoch\",\"type\":\"uint64\"},{\"internalType\":\"address[]\",\"name\":\"validators\",\"type\":\"address[]\"},{\"internalType\":\"uint64[]\",\"name\":\"powers\",\"type\":\"uint64[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"hotValidatorList\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"hotAddresses\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"coldAddresses\",\"type\":\"address[]\"},{\"internalType\":\"uint64[]\",\"name\":\"powers\",\"type\":\"uint64[]\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"user\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destination\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"token\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"blockNumber\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"txHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"}],\"name\":\"makeDepositMessage\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"newEpoch\",\"type\":\"uint64\"},{\"internalType\":\"address[]\",\"name\":\"hotAddresses\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"coldAddresses\",\"type\":\"address[]\"},{\"internalType\":\"uint64[]\",\"name\":\"powers\",\"type\":\"uint64[]\"}],\"name\":\"makeUpdateValidatorSetMessage\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"destination\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"token\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"name\":\"makeWithdrawMessage\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"pendingMessages\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"processedMessages\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"srcToken\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"srcTokenDecimal\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"dstToken\",\"type\":\"address\"}],\"name\":\"setTokenPair\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"setWithdrawFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"domainSeparator\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"r\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"s\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"}],\"internalType\":\"structSignature\",\"name\":\"signature\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"messageRawData\",\"type\":\"bytes\"}],\"internalType\":\"structCrossChainMessage[]\",\"name\":\"items\",\"type\":\"tuple[]\"}],\"name\":\"submitBridgeSignatures\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"tokenDecimalDiff\",\"outputs\":[{\"internalType\":\"int8\",\"name\":\"\",\"type\":\"int8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"tokenPair\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"tokenWithdrawFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalValidatorPower\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"newEpoch\",\"type\":\"uint64\"},{\"internalType\":\"address[]\",\"name\":\"hotAddresses\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"coldAddresses\",\"type\":\"address[]\"},{\"internalType\":\"uint64[]\",\"name\":\"powers\",\"type\":\"uint64[]\"}],\"name\":\"updateValidatorSet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"epoch\",\"type\":\"uint64\"},{\"internalType\":\"address[]\",\"name\":\"hotAddresses\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"coldAddresses\",\"type\":\"address[]\"},{\"internalType\":\"uint64[]\",\"name\":\"powers\",\"type\":\"uint64[]\"}],\"internalType\":\"structValidatorSetUpdate\",\"name\":\"validatorSet\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"r\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"s\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"}],\"internalType\":\"structSignature\",\"name\":\"signature\",\"type\":\"tuple\"}],\"name\":\"updateValidatorSetConfirm\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"validatorPowers\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"destination\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"destination\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"deadline\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"r\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"s\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"}],\"internalType\":\"structSignature\",\"name\":\"signature\",\"type\":\"tuple\"}],\"internalType\":\"structWithdrawWithPermit[]\",\"name\":\"withdraws\",\"type\":\"tuple[]\"}],\"name\":\"withdrawBatchWithPermit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"destination\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"token\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"r\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"s\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"}],\"internalType\":\"structSignature\",\"name\":\"signature\",\"type\":\"tuple\"}],\"internalType\":\"structWithdrawConfirm[]\",\"name\":\"withdrawConfirms\",\"type\":\"tuple[]\"}],\"name\":\"withdrawConfirm\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawNonce\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60a0604052306080523480156200001557600080fd5b506200002062000026565b620000da565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000900460ff1615620000775760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160401b0390811614620000d75780546001600160401b0319166001600160401b0390811782556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50565b608051615e94620001046000396000818161392f015281816139580152613aa90152615e946000f3fe6080604052600436106102ff5760003560e01c80638456cb5911610190578063c4366158116100dc578063e404d8ce11610095578063f61aa09d1161006f578063f61aa09d146109e2578063f698da2514610a18578063f8156a6e14610a2e578063fa734b1314610a4e57600080fd5b8063e404d8ce1461097e578063e63ab1e91461099e578063ee0791b7146109c057600080fd5b8063c4366158146108ba578063c7bd5b2e146108cf578063ceca23f1146108fe578063d547741f1461091e578063d83ee1e81461093e578063de35f5cb1461095e57600080fd5b8063ad3cb1cc11610149578063b8a4e15111610123578063b8a4e15114610833578063b9d5ca9e1461085a578063ba14c40d1461087a578063bb54608d1461089a57600080fd5b8063ad3cb1cc146107af578063b73124b2146107ed578063b7ab4db51461080f57600080fd5b80638456cb59146106d657806388ba16ab146106eb578063900cf0cf1461071b57806391d148541461075a578063a217fddf1461077a578063a8ff00751461078f57600080fd5b80633ba9613b1161024f57806352d1902d1161020857806364f8b391116101e257806364f8b391146106545780636e4bc0aa1461067457806375b238fc1461069457806378918bb8146106b657600080fd5b806352d1902d146105fa57806354e0fb211461060f5780635c975abb1461062f57600080fd5b80633ba9613b146105425780633f4ba83a14610562578063446c2c9a14610577578063456b07f9146105a75780634dac5132146105c75780634f1ef286146105e757600080fd5b80631dd82536116102bc5780632922e6e5116102965780632922e6e5146104b55780632b90c338146104d55780632f2ff15d1461050257806336568abe1461052257600080fd5b80631dd82536146104505780631ebe566514610475578063248a9ca31461049557600080fd5b806301ffc9a71461030457806303d3f5ee1461033957806306b1399f1461037457806313a4cc83146103c157806317c6365d146103f65780631c09502014610418575b600080fd5b34801561031057600080fd5b5061032461031f366004614a1b565b610a6e565b60405190151581526020015b60405180910390f35b34801561034557600080fd5b50610366610354366004614a61565b600b6020526000908152604090205481565b604051908152602001610330565b34801561038057600080fd5b506103ae61038f366004614a7c565b6001602090815260009283526040808420909152908252812054900b81565b60405160009190910b8152602001610330565b3480156103cd57600080fd5b506103666103dc366004614a7c565b600060208181529281526040808220909352908152205481565b34801561040257600080fd5b50610416610411366004614a9e565b610aa5565b005b34801561042457600080fd5b50610438610433366004614afc565b610d37565b6040516001600160a01b039091168152602001610330565b34801561045c57600080fd5b50610465610d61565b6040516103309493929190614b94565b34801561048157600080fd5b50610416610490366004614be7565b610ef5565b3480156104a157600080fd5b506103666104b0366004614afc565b610fb3565b3480156104c157600080fd5b506103666104d0366004614afc565b610fd5565b3480156104e157600080fd5b506104f56104f0366004614c5c565b610ff6565b6040516103309190614c88565b34801561050e57600080fd5b5061041661051d366004614c5c565b611070565b34801561052e57600080fd5b5061041661053d366004614c5c565b611092565b34801561054e57600080fd5b5061041661055d366004614d0e565b6110ca565b34801561056e57600080fd5b506104166111e4565b34801561058357600080fd5b50610597610592366004614afc565b611207565b6040516103309493929190614e5a565b3480156105b357600080fd5b506104166105c2366004615032565b6112dc565b3480156105d357600080fd5b506104166105e23660046150fe565b61152a565b6104166105f536600461513f565b6115d9565b34801561060657600080fd5b506103666115f4565b34801561061b57600080fd5b5061041661062a3660046151e4565b611611565b34801561063b57600080fd5b50600080516020615dff8339815191525460ff16610324565b34801561066057600080fd5b5061036661066f36600461520e565b6116d1565b34801561068057600080fd5b506104f561068f366004614c5c565b6117c5565b3480156106a057600080fd5b50610366600080516020615e3f83398151915281565b3480156106c257600080fd5b506104166106d13660046150fe565b61183f565b3480156106e257600080fd5b5061041661190d565b3480156106f757600080fd5b50610324610706366004614afc565b600e6020526000908152604090205460ff1681565b34801561072757600080fd5b5060055461074290600160401b90046001600160401b031681565b6040516001600160401b039091168152602001610330565b34801561076657600080fd5b50610324610775366004614c5c565b61192d565b34801561078657600080fd5b50610366600081565b34801561079b57600080fd5b506104166107aa366004614afc565b611965565b3480156107bb57600080fd5b506107e0604051806040016040528060058152602001640352e302e360dc1b81525081565b60405161033091906152a6565b3480156107f957600080fd5b50610802611986565b60405161033091906152b9565b34801561081b57600080fd5b506108246119e8565b604051610330939291906152cc565b34801561083f57600080fd5b50600a5461074290600160401b90046001600160401b031681565b34801561086657600080fd5b5061041661087536600461530f565b611b3d565b34801561088657600080fd5b50610438610895366004614afc565b611e62565b3480156108a657600080fd5b506103666108b5366004615357565b611e72565b3480156108c657600080fd5b50610802611ef9565b3480156108db57600080fd5b506108ef6108ea366004614afc565b611f59565b604051610330939291906153af565b34801561090a57600080fd5b506104166109193660046153e3565b611f8c565b34801561092a57600080fd5b50610416610939366004614c5c565b6120aa565b34801561094a57600080fd5b50610366610959366004615433565b6120c6565b34801561096a57600080fd5b50600a54610742906001600160401b031681565b34801561098a57600080fd5b506104166109993660046154a7565b612162565b3480156109aa57600080fd5b50610366600080516020615dbf83398151915281565b3480156109cc57600080fd5b506109d5612196565b60405161033091906154e2565b3480156109ee57600080fd5b506107426109fd366004614a61565b6004602052600090815260409020546001600160401b031681565b348015610a2457600080fd5b5061036660105481565b348015610a3a57600080fd5b50600554610742906001600160401b031681565b348015610a5a57600080fd5b50610416610a69366004615526565b6121ed565b60006001600160e01b03198216637965db0b60e01b1480610a9f57506301ffc9a760e01b6001600160e01b03198316145b92915050565b610aad6122e5565b610ab5612318565b610abe33612350565b600554600160401b90046001600160401b0316610ade602084018461555b565b6001600160401b031611610b395760405162461bcd60e51b815260206004820152601a60248201527f5374616c652076616c696461746f72207365742075706461746500000000000060448201526064015b60405180910390fd5b6000610c0d610b4b602085018561555b565b610b586020860186615576565b80806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250610b97925050506040870187615576565b80806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250610bd6925050506060880188615576565b808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152506116d192505050565b90506000610c1a826123be565b90506000610c3683610c3136879003870187615624565b612410565b90508015610d1957610c4782612605565b610d19610c57602087018761555b565b610c646020880188615576565b80806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250610ca3925050506040890189615576565b80806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250610ce29250505060608a018a615576565b808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152506126ef92505050565b505050610d336001600080516020615e1f83398151915255565b5050565b60028181548110610d4757600080fd5b6000918252602090912001546001600160a01b0316905081565b60408051608081018252600680546001600160401b03168252600780548451602082810282018101909652818152600095606095869586958995929491938086019390830182828015610ddd57602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311610dbf575b5050505050815260200160028201805480602002602001604051908101604052809291908181526020018280548015610e3f57602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311610e21575b5050505050815260200160038201805480602002602001604051908101604052809291908181526020018280548015610ec957602002820191906000526020600020906000905b82829054906101000a90046001600160401b03166001600160401b031681526020019060080190602082600701049283019260010382029150808411610e865790505b505050919092525050815160208301516040840151606090940151919990985092965094509092505050565b610efd6122e5565b610f05612318565b610f0e33612350565b80610f525760405162461bcd60e51b8152602060048201526014602482015273456d707479206465706f7369747320617272617960601b6044820152606401610b30565b60005b81811015610f9b5736838383818110610f7057610f70615640565b905061016002019050610f9281803603810190610f8d9190615656565b612ad8565b50600101610f55565b50610d336001600080516020615e1f83398151915255565b6000908152600080516020615ddf833981519152602052604090206001015490565b600f8181548110610fe557600080fd5b600091825260209091200154905081565b61101d60405180606001604052806000815260200160008152602001600060ff1681525090565b506000828152600d602090815260408083206001600160a01b03851684526001908101835292819020815160608101835281548152938101549284019290925260029091015460ff169082015292915050565b61107982610fb3565b61108281612ee6565b61108c8383612ef0565b50505050565b6001600160a01b03811633146110bb5760405163334bd91960e11b815260040160405180910390fd5b6110c58282612f9c565b505050565b600080516020615e3f8339815191526110e281612ee6565b6000886001600160401b03161161113b5760405162461bcd60e51b815260206004820152601a60248201527f4e65772065706f6368206d75737420626520706f7369746976650000000000006044820152606401610b30565b6111da8888888080602002602001604051908101604052809392919081815260200183836020028082843760009201919091525050604080516020808c0282810182019093528b82529093508b92508a91829185019084908082843760009201919091525050604080516020808b0282810182019093528a82529093508a92508991829185019084908082843760009201919091525061301892505050565b5050505050505050565b600080516020615dbf8339815191526111fc81612ee6565b611204613650565b50565b6000818152600c6020526040812060609081908190611225816136b0565b82546002840180549397509195506001600160401b03169186918691819061124c906156ea565b80601f0160208091040260200160405190810160405280929190818152602001828054611278906156ea565b80156112c55780601f1061129a576101008083540402835291602001916112c5565b820191906000526020600020905b8154815290600101906020018083116112a857829003601f168201915b505050505090509450945094509450509193509193565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a008054600160401b810460ff1615906001600160401b03166000811580156113215750825b90506000826001600160401b0316600114801561133d5750303b155b90508115801561134b575080155b156113695760405163f92ee8a960e01b815260040160405180910390fd5b845467ffffffffffffffff19166001178555831561139357845460ff60401b1916600160401b1785555b61139b613864565b6113a3613864565b6113ab61386c565b6113b361387c565b6113be600033612ef0565b506113d7600080516020615e3f83398151915233612ef0565b506113f0600080516020615dbf83398151915233612ef0565b506113fe6000898989613018565b61140b60008989896126ef565b6040805180820182526009815268213934b233b2a43ab160b91b6020918201528151808301835260018152603160f81b9082015281517f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f818301527f4c4f9c18a13e136e05f5178f806d2f407f435220d27e3fd1cf60052217fef7e4818401527fc89efdaa54c0f20c7adf612882df0950f5a951637e0307cdcb4c672f298b8bc660608201524660808201523060a0808301919091528351808303909101815260c0909101909252815191012060105583156111da57845460ff60401b19168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15050505050505050565b6115326122e5565b61153a612318565b61154333612350565b806115905760405162461bcd60e51b815260206004820152601d60248201527f456d70747920776974686472617720636f6e6669726d732061727261790000006044820152606401610b30565b60005b81811015610f9b57368383838181106115ae576115ae615640565b9050610120020190506115d0818036038101906115cb9190615724565b61388c565b50600101611593565b6115e1613924565b6115ea826139c9565b610d3382826139e1565b60006115fe613a9e565b50600080516020615d9f83398151915290565b600080516020615e3f83398151915261162981612ee6565b6001600160a01b0383166116775760405162461bcd60e51b8152602060048201526015602482015274496e76616c696420746f6b656e206164647265737360581b6044820152606401610b30565b6001600160a01b0383166000818152600b602052604090819020849055517f64f4bc88c432868175af31fe623ff706067afe070c40f869f28e3daed63f8243906116c49085815260200190565b60405180910390a2505050565b60007fcd26826da4f5c0e82ef8057ecacd8931dfb36167a70c820505f10826298cd05e8585604051602001611706919061579c565b604051602081830303815290604052805190602001208560405160200161172d919061579c565b604051602081830303815290604052805190602001208560405160200161175491906157db565b604051602081830303815290604052805190602001206040516020016117a59594939291909485526001600160401b0393909316602085015260408401919091526060830152608082015260a00190565b604051602081830303815290604052805190602001209050949350505050565b6117ec60405180606001604052806000815260200160008152602001600060ff1681525090565b506000828152600c602090815260408083206001600160a01b03851684526001908101835292819020815160608101835281548152938101549284019290925260029091015460ff169082015292915050565b6118476122e5565b61184f612318565b806118945760405162461bcd60e51b8152602060048201526015602482015274456d7074792077697468647261777320617272617960581b6044820152606401610b30565b8060005b816001600160401b0316816001600160401b031610156118f4576118ec8484836001600160401b03168181106118d0576118d0615640565b905061012002018036038101906118e7919061580e565b613ae7565b600101611898565b5050610d336001600080516020615e1f83398151915255565b600080516020615dbf83398151915261192581612ee6565b611204613bb9565b6000918252600080516020615ddf833981519152602090815260408084206001600160a01b0393909316845291905290205460ff1690565b600080516020615e3f83398151915261197d81612ee6565b610d3382613c02565b606060038054806020026020016040519081016040528092919081815260200182805480156119de57602002820191906000526020600020905b81546001600160a01b031681526001909101906020018083116119c0575b5050505050905090565b6000606080600560089054906101000a90046001600160401b031692506003805480602002602001604051908101604052809291908181526020018280548015611a5b57602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311611a3d575b5050505050915081516001600160401b03811115611a7b57611a7b614ea2565b604051908082528060200260200182016040528015611aa4578160200160208202803683370190505b50905060005b8251811015611b375760046000848381518110611ac957611ac9615640565b60200260200101516001600160a01b03166001600160a01b0316815260200190815260200160002060009054906101000a90046001600160401b0316828281518110611b1757611b17615640565b6001600160401b0390921660209283029190910190910152600101611aaa565b50909192565b600080516020615e3f833981519152611b5581612ee6565b84600003611b755760405162461bcd60e51b8152600401610b309061586f565b83611bb65760405162461bcd60e51b815260206004820152601160248201527024b73b30b634b21039b931903a37b5b2b760791b6044820152606401610b30565b6001600160a01b038216611c005760405162461bcd60e51b815260206004820152601160248201527024b73b30b634b2103239ba103a37b5b2b760791b6044820152606401610b30565b60008581526020818152604080832087845290915290205415611c5e5760405162461bcd60e51b8152602060048201526016602482015275151bdad95b881c185a5c88185b1c9958591e481cd95d60521b6044820152606401610b30565b604d60ff84161115611ca95760405162461bcd60e51b8152602060048201526014602482015273496e76616c69642073726320646563696d616c7360601b6044820152606401610b30565b6000826001600160a01b031663313ce5676040518163ffffffff1660e01b8152600401602060405180830381865afa158015611ce9573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611d0d9190615898565b9050604d60ff82161115611d5a5760405162461bcd60e51b8152602060048201526014602482015273496e76616c69642064737420646563696d616c7360601b6044820152606401610b30565b6000611d6682866158cb565b905060006001600160a01b0385166000898152602081815260408083208b845282528083208490558b8352600182528083208b84529091529020805460ff191660ff851617905590506001600160a01b038711611e0c576000888152602081815260408083208484529091529020879055611de0826158ec565b60008981526001602090815260408083208584529091529020805460ff191660ff929092169190911790555b6040805160ff8089168252851660208201526001600160a01b0387169189918b917f44466fd39f128be1926fbbf7b1314fcddfb506aa13ede292765d9ea429b5a31f910160405180910390a45050505050505050565b60038181548110610d4757600080fd5b604080517fae7dee9fe1cf9016b724a74908236e5f57a1789d05c09f2625fa9baee0cde49d6020808301919091526001600160a01b0398909816818301526060810196909652608086019490945260a085019290925260c08401526001600160401b031660e080840191909152815180840390910181526101009092019052805191012090565b606060028054806020026020016040519081016040528092919081815260200182805480156119de576020028201919060005260206000209081546001600160a01b031681526001909101906020018083116119c0575050505050905090565b6000818152600d602052604081206060908190611f75816136b0565b91546001600160401b031696909550909350915050565b600080516020615e3f833981519152611fa481612ee6565b6001600160a01b038316611ffa5760405162461bcd60e51b815260206004820152601960248201527f496e76616c696420726563697069656e742061646472657373000000000000006044820152606401610b30565b6000821161204a5760405162461bcd60e51b815260206004820152601760248201527f416d6f756e74206d75737420626520706f7369746976650000000000000000006044820152606401610b30565b61205e6001600160a01b0385168484613c8a565b604080516001600160a01b038581168252602082018590528616917ffe3464cd748424446c37877c28ce5b700222c5bc9f90d908afcc4e5cb22707ff910160405180910390a250505050565b6120b382610fb3565b6120bc81612ee6565b61108c8383612f9c565b604080517fe9dc4d5901ee3d0ceb0faa2d345b09188f1d7bfd6cd34ea09ed7ba75ff57391b6020808301919091528183019a909a526001600160a01b03989098166060890152608088019690965260a087019490945260c08601929092526001600160401b031660e085015261010084015263ffffffff1661012080840191909152815180840390910181526101409092019052805191012090565b61216a6122e5565b612172612318565b61217f3385858585613ce9565b61108c6001600080516020615e1f83398151915255565b6060600f8054806020026020016040519081016040528092919081815260200182805480156119de57602002820191906000526020600020905b8154815260200190600101908083116121d0575050505050905090565b6121f56122e5565b6121fd612318565b61220633612350565b8061224c5760405162461bcd60e51b8152602060048201526016602482015275456d707479207369676e61747572657320617272617960501b6044820152606401610b30565b60005b81811015610f9b576122dd83838381811061226c5761226c615640565b905060200281019061227e919061590a565b60200184848481811061229357612293615640565b90506020028101906122a5919061590a565b358585858181106122b8576122b8615640565b90506020028101906122ca919061590a565b6122d890608081019061592a565b61404c565b60010161224f565b600080516020615dff8339815191525460ff16156123165760405163d93c066560e01b815260040160405180910390fd5b565b600080516020615e1f83398151915280546001190161234a57604051633ee5aeb560e01b815260040160405180910390fd5b60029055565b6001600160a01b0381166000908152600460205260409020546001600160401b03166112045760405162461bcd60e51b815260206004820152601960248201527f5369676e6572206973206e6f7420612076616c696461746f72000000000000006044820152606401610b30565b6000806123ca83614245565b905060008111610a9f5760405162461bcd60e51b815260206004820152601160248201527013595cdcd859d9481b9bdd08199bdd5b99607a1b6044820152606401610b30565b6000828152600e602052604081205460ff16156124635760405162461bcd60e51b8152602060048201526011602482015270105b1c9958591e481c1c9bd8d95cdcd959607a1b6044820152606401610b30565b6000612472848460105461429e565b905061247d81612350565b6000848152600d602090815260408083206001600160a01b0385168452600101909152902054156124fe5760405162461bcd60e51b815260206004820152602560248201527f56616c696461746f7220616c7265616479207369676e65642074686973206d65604482015264737361676560d81b6064820152608401610b30565b6000848152600d602081815260408084206001600160a01b038616855260018082018452828620895181558985015191810191909155888301516002909101805460ff191660ff90921691909117905560048352908420548885529290915254612574916001600160401b039081169116615970565b6000868152600d60205260409020805467ffffffffffffffff19166001600160401b03838116919091179091556005549192506125b391166002615990565b6001600160401b03166125c7826003615990565b6001600160401b031611156125fa575050506000828152600e60205260409020805460ff19166001908117909155610a9f565b506000949350505050565b6000811180156126175750600f548111155b6126535760405162461bcd60e51b815260206004820152600d60248201526c092dcecc2d8d2c840d2dcc8caf609b1b6044820152606401610b30565b60006126606001836159bb565b600f54909150612672906001906159bb565b8110156126c457600f8054612689906001906159bb565b8154811061269957612699615640565b9060005260206000200154600f82815481106126b7576126b7615640565b6000918252602090912001555b600f8054806126d5576126d56159ce565b600190038181906000526020600020016000905590555050565b6003546000906001600160401b0381111561270c5761270c614ea2565b604051908082528060200260200182016040528015612735578160200160208202803683370190505b50905060005b6003548110156128115760006003828154811061275a5761275a615640565b60009182526020808320909101546001600160a01b031680835260049091526040909120549091506001600160401b031615612808576001600160a01b03811660009081526004602052604090205483516001600160401b03909116908490849081106127c9576127c9615640565b6001600160401b039092166020928302919091018201526001600160a01b0382166000908152600490915260409020805467ffffffffffffffff191690555b5060010161273b565b50600560089054906101000a90046001600160401b03166001600160401b03167f7078dd7d8f69886fa49f17529fbd86d6b24846959cd85dcfb4e990205537dc5d600360028460405161286693929190615a23565b60405180910390a282516128819060029060208601906148b8565b5083516128959060039060208701906148b8565b506000805b855181101561292c5760008682815181106128b7576128b7615640565b6020026020010151905060008583815181106128d5576128d5615640565b6020908102919091018101516001600160a01b038416600090815260049092526040909120805467ffffffffffffffff19166001600160401b03831617905590506129208185615970565b9350505060010161289a565b50600580546001600160401b03888116600160401b026fffffffffffffffffffffffffffffffff199092169084161717905560408051608081019091526000808252602082019060405190808252806020026020018201604052801561299c578160200160208202803683370190505b50815260200160006040519080825280602002602001820160405280156129cd578160200160208202803683370190505b50815260200160006040519080825280602002602001820160405280156129fe578160200160208202803683370190505b50905280516006805467ffffffffffffffff19166001600160401b039092169190911781556020808301518051612a399260079201906148b8565b5060408201518051612a559160028401916020909101906148b8565b5060608201518051612a7191600384019160209091019061491d565b50905050856001600160401b03167f7fb1406cb8c05384e1f39b879d591d98e7be9ffa67cf31cdbaa10b3aebf7157c6003600286604051612ab493929190615a23565b60405180910390a2505050505050565b6001600080516020615e1f83398151915255565b8051612b155760405162461bcd60e51b815260206004820152600c60248201526b24b73b30b634b2103ab9b2b960a11b6044820152606401610b30565b60208101516001600160a01b0316612b6f5760405162461bcd60e51b815260206004820152601b60248201527f496e76616c69642064657374696e6174696f6e206164647265737300000000006044820152606401610b30565b6040810151612bb05760405162461bcd60e51b815260206004820152600d60248201526c24b73b30b634b2103a37b5b2b760991b6044820152606401610b30565b6000816060015111612bf85760405162461bcd60e51b81526020600482015260116024820152700416d6f756e74206d757374206265203e3607c1b6044820152606401610b30565b8060800151600003612c1c5760405162461bcd60e51b8152600401610b309061586f565b60808101516000908152602081815260408083208185015184529091529020548015801590612c575750612c57816001600160a01b03101590565b612c9b5760405162461bcd60e51b815260206004820152601560248201527424b73b30b634b210313934b233b2b2103a37b5b2b760591b6044820152606401610b30565b60008190506000612cd2846000015185602001518660400151876060015188608001518960a001518a60c001518b60e001516120c6565b90506000612ce582866101000151612410565b90508015612edf5760808501516000908152600160209081526040808320818901518452909152812054606087015190820b9190612d239083614467565b905060008111612d6d5760405162461bcd60e51b8152602060048201526015602482015274135a5b9d08185b5bdd5b9d081d1bdbc81cdb585b1b605a1b6044820152606401610b30565b60208701516040516340c10f1960e01b81526001600160a01b03918216600482015260248101839052908616906340c10f1990604401600060405180830381600087803b158015612dbd57600080fd5b505af1158015612dd1573d6000803e3d6000fd5b5050600a805460019350909150600090612df59084906001600160401b0316615970565b92506101000a8154816001600160401b0302191690836001600160401b0316021790555086604001518760000151857fbaa0634355881c40ba9bac876b01a2b891bcaea1c37a12eba82f5bcf88d048bc8a602001518b606001518c608001518d60a001518e60c001518f60e00151600a60009054906101000a90046001600160401b0316604051612ed497969594939291906001600160a01b03979097168752602087019590955260408601939093526001600160401b039182166060860152608085015263ffffffff9190911660a08401521660c082015260e00190565b60405180910390a450505b5050505050565b61120481336144c1565b6000600080516020615ddf833981519152612f0b848461192d565b612f8b576000848152602082815260408083206001600160a01b03871684529091529020805460ff19166001179055612f413390565b6001600160a01b0316836001600160a01b0316857f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a46001915050610a9f565b6000915050610a9f565b5092915050565b6000600080516020615ddf833981519152612fb7848461192d565b15612f8b576000848152602082815260408083206001600160a01b0387168085529252808320805460ff1916905551339287917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a46001915050610a9f565b60008351116130695760405162461bcd60e51b815260206004820152601960248201527f456d70747920686f7420616464726573736573206172726179000000000000006044820152606401610b30565b81518351146130c95760405162461bcd60e51b815260206004820152602660248201527f486f7420616e6420636f6c6420616464726573736573206c656e677468206d696044820152650e6dac2e8c6d60d31b6064820152608401610b30565b805183511461312b5760405162461bcd60e51b815260206004820152602860248201527f486f742061646472657373657320616e6420706f77657273206c656e677468206044820152670dad2e6dac2e8c6d60c31b6064820152608401610b30565b600554600160401b90046001600160401b031615158061315357506001600160401b03841615155b156131d2576005546001600160401b03600160401b9091048116908516116131d25760405162461bcd60e51b815260206004820152602c60248201527f4e65772065706f6368206d7573742062652067726561746572207468616e206360448201526b0eae4e4cadce840cae0dec6d60a31b6064820152608401610b30565b6000805b84518110156134d25760008582815181106131f3576131f3615640565b60200260200101519050600085838151811061321157613211615640565b60200260200101519050600085848151811061322f5761322f615640565b6020026020010151905060006001600160a01b0316836001600160a01b03160361329b5760405162461bcd60e51b815260206004820152601d60248201527f5a65726f206164647265737320696e20686f74206164647265737365730000006044820152606401610b30565b6001600160a01b0382166132f15760405162461bcd60e51b815260206004820152601e60248201527f5a65726f206164647265737320696e20636f6c642061646472657373657300006044820152606401610b30565b6000816001600160401b0316116133435760405162461bcd60e51b8152602060048201526016602482015275506f776572206d75737420626520706f73697469766560501b6044820152606401610b30565b816001600160a01b0316836001600160a01b0316036133af5760405162461bcd60e51b815260206004820152602260248201527f486f7420616e6420636f6c6420616464726573736573206d757374206469666660448201526132b960f11b6064820152608401610b30565b60005b848110156134b6578881815181106133cc576133cc615640565b60200260200101516001600160a01b0316846001600160a01b0316036134345760405162461bcd60e51b815260206004820152601760248201527f4475706c696361746520686f74206164647265737365730000000000000000006044820152606401610b30565b87818151811061344657613446615640565b60200260200101516001600160a01b0316836001600160a01b0316036134ae5760405162461bcd60e51b815260206004820152601860248201527f4475706c696361746520636f6c642061646472657373657300000000000000006044820152606401610b30565b6001016133b2565b506134c18186615970565b945050600190920191506131d69050565b506000816001600160401b03161161352c5760405162461bcd60e51b815260206004820152601c60248201527f546f74616c20706f776572206d75737420626520706f736974697665000000006044820152606401610b30565b600061353a868686866116d1565b90506001600160401b0386161561358157600f80546001810182556000919091527f8d1108e10bcb7c27dddfc02ed9d693a074039d026cf4ea4240b40f7d581ac802018190555b604080516080810182526001600160401b0388168082526020808301899052928201879052606082018690526006805467ffffffffffffffff191690911781558751919290916135d791600791908a01906148b8565b50604082015180516135f39160028401916020909101906148b8565b506060820151805161360f91600384019160209091019061491d565b50905050856001600160401b03167ff389db8d301520921c1c0eeb8eaf6790791f8916b65e8e93424a7ec831edd69e868686604051612ab493929190615a5c565b6136586144fa565b600080516020615dff833981519152805460ff191681557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b6040516001600160a01b03909116815260200160405180910390a150565b60035460609081906000816001600160401b038111156136d2576136d2614ea2565b6040519080825280602002602001820160405280156136fb578160200160208202803683370190505b5090506000826001600160401b0381111561371857613718614ea2565b60405190808252806020026020018201604052801561377057816020015b61375d60405180606001604052806000815260200160008152602001600060ff1681525090565b8152602001906001900390816137365790505b5090506000805b848110156138535760006003828154811061379457613794615640565b60009182526020808320909101546001600160a01b031680835260018c810183526040938490208451606081018652815480825292820154948101949094526002015460ff16938301939093529250901561384957818685815181106137fc576137fc615640565b60200260200101906001600160a01b031690816001600160a01b0316815250508085858151811061382f5761382f615640565b6020026020010181905250838061384590615a81565b9450505b5050600101613777565b508083528152909590945092505050565b61231661452a565b61387461452a565b612316614573565b61388461452a565b612316614594565b60006138b4826000015183602001518460400151856060015186608001518760a00151611e72565b905060006138c1826123be565b905060006138d3838560c00151612410565b9050801561108c576138e482612605565b60a084015160405184916001600160401b0316907fda5f7a28c1ede043b7194dba8741d889157152b30393f445fac51654abe02fae90600090a350505050565b306001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614806139ab57507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031661399f600080516020615d9f833981519152546001600160a01b031690565b6001600160a01b031614155b156123165760405163703e46dd60e11b815260040160405180910390fd5b600080516020615e3f833981519152610d3381612ee6565b816001600160a01b03166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa925050508015613a3b575060408051601f3d908101601f19168201909252613a3891810190615a9a565b60015b613a6357604051634c9c8ce360e01b81526001600160a01b0383166004820152602401610b30565b600080516020615d9f8339815191528114613a9457604051632a87526960e21b815260048101829052602401610b30565b6110c5838361459c565b306001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146123165760405163703e46dd60e11b815260040160405180910390fd5b6040818101518251606084015160a085015160c0860151808601518151602090920151965163d505accf60e01b81526001600160a01b03958616600482015230602482015260448101949094526001600160401b03909216606484015260ff909116608483015260a482015260c4810193909352169063d505accf9060e401600060405180830381600087803b158015613b8057600080fd5b505af1158015613b94573d6000803e3d6000fd5b5050505061120481600001518260200151836040015184606001518560800151613ce9565b613bc16122e5565b600080516020615dff833981519152805460ff191660011781557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a25833613692565b6000818152600d60205260408120805467ffffffffffffffff1916815590613c2d60028301826149d0565b50506000818152600c60205260408120805467ffffffffffffffff1916815590613c5a60028301826149d0565b505060405181907fa5e9cb0cafdc12b93fa2ff5814a2a2f6ef3e98f1097c8073778b085111058f0090600090a250565b6040516001600160a01b038381166024830152604482018390526110c591859182169063a9059cbb906064015b604051602081830303815290604052915060e01b6020820180516001600160e01b0383818316178352505050506145f2565b6001600160a01b0383166000908152600b6020526040902054808311613d4a5760405162461bcd60e51b8152602060048201526016602482015275416d6f756e74206d757374206578636565642066656560501b6044820152606401610b30565b6001600160a01b038616613d975760405162461bcd60e51b8152602060048201526014602482015273496e76616c69642075736572206164647265737360601b6044820152606401610b30565b84613dda5760405162461bcd60e51b815260206004820152601360248201527224b73b30b634b2103232b9ba34b730ba34b7b760691b6044820152606401610b30565b6001600160a01b038416613e285760405162461bcd60e51b8152602060048201526015602482015274496e76616c696420746f6b656e206164647265737360581b6044820152606401610b30565b81600003613e485760405162461bcd60e51b8152600401610b309061586f565b6000828152602081815260408083206001600160a01b03881680855292529091205480613ea95760405162461bcd60e51b815260206004820152600f60248201526e151bdad95b881b9bdd08199bdd5b99608a1b6044820152606401610b30565b613ebe6001600160a01b038716893088614655565b613ec883866159bb565b604051630852cd8d60e31b8152600481018290529095506001600160a01b038716906342966c6890602401600060405180830381600087803b158015613f0d57600080fd5b505af1158015613f21573d6000803e3d6000fd5b505050506000600a600881819054906101000a90046001600160401b0316613f4890615ab3565b82546001600160401b038083166101009490940a9384029302191691909117909155600086815260016020908152604080832087845290915281205491925090810b90613f958883614467565b90506000613fa78c8c87858c89611e72565b600f80546001810182556000919091527f8d1108e10bcb7c27dddfc02ed9d693a074039d026cf4ea4240b40f7d581ac80201819055604080518d8152602081018590529081018a90526001600160401b038616606082015290915085906001600160a01b038e169083907fddf7473863baaba91098e18c11ea6e972fb72f38f768f5e5d3e984207414ff139060800160405180910390a4505050505050505050505050565b601054830361408e5760405162461bcd60e51b815260206004820152600e60248201526d24b73b30b634b2103237b6b0b4b760911b6044820152606401610b30565b600082826040516140a0929190615ad9565b604051809103902090506140b3816123be565b5060006140cf826140c936899003890189615624565b8761429e565b90506140da81612350565b6000828152600c602090815260408083206001600160a01b03851684526001019091529020541561413e5760405162461bcd60e51b815260206004820152600e60248201526d105b1c9958591e481cda59db995960921b6044820152606401610b30565b6000828152600c602090815260408083206001600160a01b03851684526001019091529020869061416f8282615ae9565b50506000828152600c6020526040902060020161418d848683615b6c565b506001600160a01b038116600090815260046020908152604080832054858452600c9092528220546141cb916001600160401b039081169116615970565b6000848152600c6020908152604091829020805467ffffffffffffffff19166001600160401b03851690811790915582516001600160a01b03871681529182015291925084917fc36d3c025c0f565ea8451b940f66ac343971b6b6d515ea860aed580c6408af2b910160405180910390a250505050505050565b600f54600090815b818110156142945783600f828154811061426957614269615640565b90600052602060002001540361428c57614284816001615c2b565b949350505050565b60010161424d565b5060009392505050565b815160009081036142f15760405162461bcd60e51b815260206004820152601b60248201527f496e76616c6964207369676e6174757265202772272076616c756500000000006044820152606401610b30565b82602001516000036143455760405162461bcd60e51b815260206004820152601b60248201527f496e76616c6964207369676e6174757265202773272076616c756500000000006044820152606401610b30565b60405161190160f01b6020820152602281018390526042810185905260009060620160408051601f19818403018152828252805160209182012087830151885189840151600080885296909401948590529195506001936143c19387939193845260ff9290921660208401526040830152606082015260800190565b6020604051602081039080840390855afa1580156143e3573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b03811661445c5760405162461bcd60e51b815260206004820152602d60248201527f496e76616c6964207369676e61747572652c207265636f76657265642074686560448201526c207a65726f206164647265737360981b6064820152608401610b30565b9150505b9392505050565b60008160000b60000361447b575081610a9f565b60008260000b13156144a35761449282600a615d22565b61449c9084615d31565b9050610a9f565b6144ac826158ec565b6144b790600a615d22565b61449c9084615d53565b6144cb828261192d565b610d335760405163e2517d3f60e01b81526001600160a01b038216600482015260248101839052604401610b30565b600080516020615dff8339815191525460ff1661231657604051638dfc202b60e01b815260040160405180910390fd5b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0054600160401b900460ff1661231657604051631afcd79f60e31b815260040160405180910390fd5b61457b61452a565b600080516020615dff833981519152805460ff19169055565b612ac461452a565b6145a58261468e565b6040516001600160a01b038316907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b90600090a28051156145ea576110c582826146f3565b610d33614769565b60006146076001600160a01b03841683614788565b9050805160001415801561462c57508080602001905181019061462a9190615d6a565b155b156110c557604051635274afe760e01b81526001600160a01b0384166004820152602401610b30565b6040516001600160a01b03848116602483015283811660448301526064820183905261108c9186918216906323b872dd90608401613cb7565b806001600160a01b03163b6000036146c457604051634c9c8ce360e01b81526001600160a01b0382166004820152602401610b30565b600080516020615d9f83398151915280546001600160a01b0319166001600160a01b0392909216919091179055565b6060600080846001600160a01b0316846040516147109190615d8c565b600060405180830381855af49150503d806000811461474b576040519150601f19603f3d011682016040523d82523d6000602084013e614750565b606091505b5091509150614760858383614796565b95945050505050565b34156123165760405163b398979f60e01b815260040160405180910390fd5b6060614460838360006147f2565b6060826147ab576147a68261488f565b614460565b81511580156147c257506001600160a01b0384163b155b156147eb57604051639996b31560e01b81526001600160a01b0385166004820152602401610b30565b5080614460565b6060814710156148175760405163cd78605960e01b8152306004820152602401610b30565b600080856001600160a01b031684866040516148339190615d8c565b60006040518083038185875af1925050503d8060008114614870576040519150601f19603f3d011682016040523d82523d6000602084013e614875565b606091505b5091509150614885868383614796565b9695505050505050565b80511561489f5780518082602001fd5b604051630a12f52160e11b815260040160405180910390fd5b82805482825590600052602060002090810192821561490d579160200282015b8281111561490d57825182546001600160a01b0319166001600160a01b039091161782556020909201916001909101906148d8565b50614919929150614a06565b5090565b8280548282559060005260206000209060030160049004810192821561490d5791602002820160005b8382111561499057835183826101000a8154816001600160401b0302191690836001600160401b031602179055509260200192600801602081600701049283019260010302614946565b80156149c35782816101000a8154906001600160401b030219169055600801602081600701049283019260010302614990565b5050614919929150614a06565b5080546149dc906156ea565b6000825580601f106149ec575050565b601f01602090049060005260206000209081019061120491905b5b808211156149195760008155600101614a07565b600060208284031215614a2d57600080fd5b81356001600160e01b03198116811461446057600080fd5b80356001600160a01b0381168114614a5c57600080fd5b919050565b600060208284031215614a7357600080fd5b61446082614a45565b60008060408385031215614a8f57600080fd5b50508035926020909101359150565b6000808284036080811215614ab257600080fd5b83356001600160401b03811115614ac857600080fd5b840160808187031215614ada57600080fd5b92506060601f1982011215614aee57600080fd5b506020830190509250929050565b600060208284031215614b0e57600080fd5b5035919050565b60008151808452602080850194506020840160005b83811015614b4f5781516001600160a01b031687529582019590820190600101614b2a565b509495945050505050565b60008151808452602080850194506020840160005b83811015614b4f5781516001600160401b031687529582019590820190600101614b6f565b6001600160401b0385168152608060208201526000614bb66080830186614b15565b8281036040840152614bc88186614b15565b90508281036060840152614bdc8185614b5a565b979650505050505050565b60008060208385031215614bfa57600080fd5b82356001600160401b0380821115614c1157600080fd5b818501915085601f830112614c2557600080fd5b813581811115614c3457600080fd5b86602061016083028501011115614c4a57600080fd5b60209290920196919550909350505050565b60008060408385031215614c6f57600080fd5b82359150614c7f60208401614a45565b90509250929050565b815181526020808301519082015260408083015160ff169082015260608101610a9f565b80356001600160401b0381168114614a5c57600080fd5b60008083601f840112614cd557600080fd5b5081356001600160401b03811115614cec57600080fd5b6020830191508360208260051b8501011115614d0757600080fd5b9250929050565b60008060008060008060006080888a031215614d2957600080fd5b614d3288614cac565b965060208801356001600160401b0380821115614d4e57600080fd5b614d5a8b838c01614cc3565b909850965060408a0135915080821115614d7357600080fd5b614d7f8b838c01614cc3565b909650945060608a0135915080821115614d9857600080fd5b50614da58a828b01614cc3565b989b979a50959850939692959293505050565b60008151808452602080850194506020840160005b83811015614b4f57614df7878351805182526020808201519083015260409081015160ff16910152565b6060969096019590820190600101614dcd565b60005b83811015614e25578181015183820152602001614e0d565b50506000910152565b60008151808452614e46816020860160208601614e0a565b601f01601f19169290920160200192915050565b6001600160401b0385168152608060208201526000614e7c6080830186614b15565b8281036040840152614e8e8186614db8565b90508281036060840152614bdc8185614e2e565b634e487b7160e01b600052604160045260246000fd5b60405161012081016001600160401b0381118282101715614edb57614edb614ea2565b60405290565b60405160e081016001600160401b0381118282101715614edb57614edb614ea2565b604051601f8201601f191681016001600160401b0381118282101715614f2b57614f2b614ea2565b604052919050565b60006001600160401b03821115614f4c57614f4c614ea2565b5060051b60200190565b600082601f830112614f6757600080fd5b81356020614f7c614f7783614f33565b614f03565b8083825260208201915060208460051b870101935086841115614f9e57600080fd5b602086015b84811015614fc157614fb481614a45565b8352918301918301614fa3565b509695505050505050565b600082601f830112614fdd57600080fd5b81356020614fed614f7783614f33565b8083825260208201915060208460051b87010193508684111561500f57600080fd5b602086015b84811015614fc15761502581614cac565b8352918301918301615014565b60008060006060848603121561504757600080fd5b83356001600160401b038082111561505e57600080fd5b61506a87838801614f56565b9450602086013591508082111561508057600080fd5b61508c87838801614f56565b935060408601359150808211156150a257600080fd5b506150af86828701614fcc565b9150509250925092565b60008083601f8401126150cb57600080fd5b5081356001600160401b038111156150e257600080fd5b60208301915083602061012083028501011115614d0757600080fd5b6000806020838503121561511157600080fd5b82356001600160401b0381111561512757600080fd5b615133858286016150b9565b90969095509350505050565b6000806040838503121561515257600080fd5b61515b83614a45565b91506020808401356001600160401b038082111561517857600080fd5b818601915086601f83011261518c57600080fd5b81358181111561519e5761519e614ea2565b6151b0601f8201601f19168501614f03565b915080825287848285010111156151c657600080fd5b80848401858401376000848284010152508093505050509250929050565b600080604083850312156151f757600080fd5b61520083614a45565b946020939093013593505050565b6000806000806080858703121561522457600080fd5b61522d85614cac565b935060208501356001600160401b038082111561524957600080fd5b61525588838901614f56565b9450604087013591508082111561526b57600080fd5b61527788838901614f56565b9350606087013591508082111561528d57600080fd5b5061529a87828801614fcc565b91505092959194509250565b6020815260006144606020830184614e2e565b6020815260006144606020830184614b15565b6001600160401b03841681526060602082015260006152ee6060830185614b15565b82810360408401526148858185614b5a565b60ff8116811461120457600080fd5b6000806000806080858703121561532557600080fd5b8435935060208501359250604085013561533e81615300565b915061534c60608601614a45565b905092959194509250565b60008060008060008060c0878903121561537057600080fd5b61537987614a45565b9550602087013594506040870135935060608701359250608087013591506153a360a08801614cac565b90509295509295509295565b6001600160401b03841681526060602082015260006153d16060830185614b15565b82810360408401526148858185614db8565b6000806000606084860312156153f857600080fd5b61540184614a45565b925061540f60208501614a45565b9150604084013590509250925092565b803563ffffffff81168114614a5c57600080fd5b600080600080600080600080610100898b03121561545057600080fd5b8835975061546060208a01614a45565b965060408901359550606089013594506080890135935061548360a08a01614cac565b925060c0890135915061549860e08a0161541f565b90509295985092959890939650565b600080600080608085870312156154bd57600080fd5b843593506154cd60208601614a45565b93969395505050506040820135916060013590565b6020808252825182820181905260009190848201906040850190845b8181101561551a578351835292840192918401916001016154fe565b50909695505050505050565b6000806020838503121561553957600080fd5b82356001600160401b0381111561554f57600080fd5b61513385828601614cc3565b60006020828403121561556d57600080fd5b61446082614cac565b6000808335601e1984360301811261558d57600080fd5b8301803591506001600160401b038211156155a757600080fd5b6020019150600581901b3603821315614d0757600080fd5b6000606082840312156155d157600080fd5b604051606081018181106001600160401b03821117156155f3576155f3614ea2565b80604052508091508235815260208301356020820152604083013561561781615300565b6040919091015292915050565b60006060828403121561563657600080fd5b61446083836155bf565b634e487b7160e01b600052603260045260246000fd5b6000610160828403121561566957600080fd5b615671614eb8565b8235815261568160208401614a45565b60208201526040830135604082015260608301356060820152608083013560808201526156b060a08401614cac565b60a082015260c083013560c08201526156cb60e0840161541f565b60e08201526101006156df858286016155bf565b908201529392505050565b600181811c908216806156fe57607f821691505b60208210810361571e57634e487b7160e01b600052602260045260246000fd5b50919050565b6000610120828403121561573757600080fd5b61573f614ee1565b61574883614a45565b81526020830135602082015260408301356040820152606083013560608201526080830135608082015261577e60a08401614cac565b60a08201526157908460c085016155bf565b60c08201529392505050565b815160009082906020808601845b838110156157cf5781516001600160a01b0316855293820193908201906001016157aa565b50929695505050505050565b815160009082906020808601845b838110156157cf5781516001600160401b0316855293820193908201906001016157e9565b6000610120828403121561582157600080fd5b615829614ee1565b61583283614a45565b81526020830135602082015261584a60408401614a45565b6040820152606083013560608201526080830135608082015261577e60a08401614cac565b6020808252600f908201526e125b9d985b1a590818da185a5b9259608a1b604082015260600190565b6000602082840312156158aa57600080fd5b815161446081615300565b634e487b7160e01b600052601160045260246000fd5b600082810b9082900b03607f198112607f82131715610a9f57610a9f6158b5565b600081810b60808101615901576159016158b5565b60000392915050565b60008235609e1983360301811261592057600080fd5b9190910192915050565b6000808335601e1984360301811261594157600080fd5b8301803591506001600160401b0382111561595b57600080fd5b602001915036819003821315614d0757600080fd5b6001600160401b03818116838216019080821115612f9557612f956158b5565b6001600160401b038181168382160280821691908281146159b3576159b36158b5565b505092915050565b81810381811115610a9f57610a9f6158b5565b634e487b7160e01b600052603160045260246000fd5b600081548084526020808501945083600052602060002060005b83811015614b4f5781546001600160a01b0316875295820195600191820191016159fe565b606081526000615a3660608301866159e4565b8281036020840152615a4881866159e4565b905082810360408401526148858185614b5a565b606081526000615a6f6060830186614b15565b8281036020840152615a488186614b15565b600060018201615a9357615a936158b5565b5060010190565b600060208284031215615aac57600080fd5b5051919050565b60006001600160401b03808316818103615acf57615acf6158b5565b6001019392505050565b8183823760009101908152919050565b8135815560208201356001820155600281016040830135615b0981615300565b815460ff191660ff919091161790555050565b601f8211156110c5576000816000526020600020601f850160051c81016020861015615b455750805b601f850160051c820191505b81811015615b6457828155600101615b51565b505050505050565b6001600160401b03831115615b8357615b83614ea2565b615b9783615b9183546156ea565b83615b1c565b6000601f841160018114615bcb5760008515615bb35750838201355b600019600387901b1c1916600186901b178355612edf565b600083815260209020601f19861690835b82811015615bfc5786850135825560209485019460019092019101615bdc565b5086821015615c195760001960f88860031b161c19848701351681555b505060018560011b0183555050505050565b80820180821115610a9f57610a9f6158b5565b600181815b80851115615c79578160001904821115615c5f57615c5f6158b5565b80851615615c6c57918102915b93841c9390800290615c43565b509250929050565b600082615c9057506001610a9f565b81615c9d57506000610a9f565b8160018114615cb35760028114615cbd57615cd9565b6001915050610a9f565b60ff841115615cce57615cce6158b5565b50506001821b610a9f565b5060208310610133831016604e8410600b8410161715615cfc575081810a610a9f565b615d068383615c3e565b8060001904821115615d1a57615d1a6158b5565b029392505050565b600061446060ff841683615c81565b600082615d4e57634e487b7160e01b600052601260045260246000fd5b500490565b8082028115828204841417610a9f57610a9f6158b5565b600060208284031215615d7c57600080fd5b8151801515811461446057600080fd5b60008251615920818460208701614e0a56fe360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800cd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033009b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f00a49807205ce4d355092ef5a8a18f56e8913cf4a201fbe287825b095693c21775a2646970667358221220251788f62ec199856ae1f14d171187337defcef993f820a165474a7ed39e6e5864736f6c63430008160033",
}

// BridgeHubABI is the input ABI used to generate the binding from.
// Deprecated: Use BridgeHubMetaData.ABI instead.
var BridgeHubABI = BridgeHubMetaData.ABI

// BridgeHubBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BridgeHubMetaData.Bin instead.
var BridgeHubBin = BridgeHubMetaData.Bin

// DeployBridgeHub deploys a new Ethereum contract, binding an instance of BridgeHub to it.
func DeployBridgeHub(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BridgeHub, error) {
	parsed, err := BridgeHubMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BridgeHubBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BridgeHub{BridgeHubCaller: BridgeHubCaller{contract: contract}, BridgeHubTransactor: BridgeHubTransactor{contract: contract}, BridgeHubFilterer: BridgeHubFilterer{contract: contract}}, nil
}

// BridgeHub is an auto generated Go binding around an Ethereum contract.
type BridgeHub struct {
	BridgeHubCaller     // Read-only binding to the contract
	BridgeHubTransactor // Write-only binding to the contract
	BridgeHubFilterer   // Log filterer for contract events
}

// BridgeHubCaller is an auto generated read-only Go binding around an Ethereum contract.
type BridgeHubCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeHubTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BridgeHubTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeHubFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BridgeHubFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeHubSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BridgeHubSession struct {
	Contract     *BridgeHub        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BridgeHubCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BridgeHubCallerSession struct {
	Contract *BridgeHubCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// BridgeHubTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BridgeHubTransactorSession struct {
	Contract     *BridgeHubTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// BridgeHubRaw is an auto generated low-level Go binding around an Ethereum contract.
type BridgeHubRaw struct {
	Contract *BridgeHub // Generic contract binding to access the raw methods on
}

// BridgeHubCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BridgeHubCallerRaw struct {
	Contract *BridgeHubCaller // Generic read-only contract binding to access the raw methods on
}

// BridgeHubTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BridgeHubTransactorRaw struct {
	Contract *BridgeHubTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBridgeHub creates a new instance of BridgeHub, bound to a specific deployed contract.
func NewBridgeHub(address common.Address, backend bind.ContractBackend) (*BridgeHub, error) {
	contract, err := bindBridgeHub(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BridgeHub{BridgeHubCaller: BridgeHubCaller{contract: contract}, BridgeHubTransactor: BridgeHubTransactor{contract: contract}, BridgeHubFilterer: BridgeHubFilterer{contract: contract}}, nil
}

// NewBridgeHubCaller creates a new read-only instance of BridgeHub, bound to a specific deployed contract.
func NewBridgeHubCaller(address common.Address, caller bind.ContractCaller) (*BridgeHubCaller, error) {
	contract, err := bindBridgeHub(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeHubCaller{contract: contract}, nil
}

// NewBridgeHubTransactor creates a new write-only instance of BridgeHub, bound to a specific deployed contract.
func NewBridgeHubTransactor(address common.Address, transactor bind.ContractTransactor) (*BridgeHubTransactor, error) {
	contract, err := bindBridgeHub(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeHubTransactor{contract: contract}, nil
}

// NewBridgeHubFilterer creates a new log filterer instance of BridgeHub, bound to a specific deployed contract.
func NewBridgeHubFilterer(address common.Address, filterer bind.ContractFilterer) (*BridgeHubFilterer, error) {
	contract, err := bindBridgeHub(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BridgeHubFilterer{contract: contract}, nil
}

// bindBridgeHub binds a generic wrapper to an already deployed contract.
func bindBridgeHub(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BridgeHubMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BridgeHub *BridgeHubRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BridgeHub.Contract.BridgeHubCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BridgeHub *BridgeHubRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeHub.Contract.BridgeHubTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BridgeHub *BridgeHubRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BridgeHub.Contract.BridgeHubTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BridgeHub *BridgeHubCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BridgeHub.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BridgeHub *BridgeHubTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeHub.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BridgeHub *BridgeHubTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BridgeHub.Contract.contract.Transact(opts, method, params...)
}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_BridgeHub *BridgeHubCaller) ADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BridgeHub.contract.Call(opts, &out, "ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_BridgeHub *BridgeHubSession) ADMINROLE() ([32]byte, error) {
	return _BridgeHub.Contract.ADMINROLE(&_BridgeHub.CallOpts)
}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_BridgeHub *BridgeHubCallerSession) ADMINROLE() ([32]byte, error) {
	return _BridgeHub.Contract.ADMINROLE(&_BridgeHub.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_BridgeHub *BridgeHubCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BridgeHub.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_BridgeHub *BridgeHubSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _BridgeHub.Contract.DEFAULTADMINROLE(&_BridgeHub.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_BridgeHub *BridgeHubCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _BridgeHub.Contract.DEFAULTADMINROLE(&_BridgeHub.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_BridgeHub *BridgeHubCaller) PAUSERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BridgeHub.contract.Call(opts, &out, "PAUSER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_BridgeHub *BridgeHubSession) PAUSERROLE() ([32]byte, error) {
	return _BridgeHub.Contract.PAUSERROLE(&_BridgeHub.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_BridgeHub *BridgeHubCallerSession) PAUSERROLE() ([32]byte, error) {
	return _BridgeHub.Contract.PAUSERROLE(&_BridgeHub.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_BridgeHub *BridgeHubCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BridgeHub.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_BridgeHub *BridgeHubSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _BridgeHub.Contract.UPGRADEINTERFACEVERSION(&_BridgeHub.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_BridgeHub *BridgeHubCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _BridgeHub.Contract.UPGRADEINTERFACEVERSION(&_BridgeHub.CallOpts)
}

// ColdValidatorList is a free data retrieval call binding the contract method 0x1c095020.
//
// Solidity: function coldValidatorList(uint256 ) view returns(address)
func (_BridgeHub *BridgeHubCaller) ColdValidatorList(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _BridgeHub.contract.Call(opts, &out, "coldValidatorList", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ColdValidatorList is a free data retrieval call binding the contract method 0x1c095020.
//
// Solidity: function coldValidatorList(uint256 ) view returns(address)
func (_BridgeHub *BridgeHubSession) ColdValidatorList(arg0 *big.Int) (common.Address, error) {
	return _BridgeHub.Contract.ColdValidatorList(&_BridgeHub.CallOpts, arg0)
}

// ColdValidatorList is a free data retrieval call binding the contract method 0x1c095020.
//
// Solidity: function coldValidatorList(uint256 ) view returns(address)
func (_BridgeHub *BridgeHubCallerSession) ColdValidatorList(arg0 *big.Int) (common.Address, error) {
	return _BridgeHub.Contract.ColdValidatorList(&_BridgeHub.CallOpts, arg0)
}

// DepositNonce is a free data retrieval call binding the contract method 0xde35f5cb.
//
// Solidity: function depositNonce() view returns(uint64)
func (_BridgeHub *BridgeHubCaller) DepositNonce(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _BridgeHub.contract.Call(opts, &out, "depositNonce")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// DepositNonce is a free data retrieval call binding the contract method 0xde35f5cb.
//
// Solidity: function depositNonce() view returns(uint64)
func (_BridgeHub *BridgeHubSession) DepositNonce() (uint64, error) {
	return _BridgeHub.Contract.DepositNonce(&_BridgeHub.CallOpts)
}

// DepositNonce is a free data retrieval call binding the contract method 0xde35f5cb.
//
// Solidity: function depositNonce() view returns(uint64)
func (_BridgeHub *BridgeHubCallerSession) DepositNonce() (uint64, error) {
	return _BridgeHub.Contract.DepositNonce(&_BridgeHub.CallOpts)
}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_BridgeHub *BridgeHubCaller) DomainSeparator(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BridgeHub.contract.Call(opts, &out, "domainSeparator")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_BridgeHub *BridgeHubSession) DomainSeparator() ([32]byte, error) {
	return _BridgeHub.Contract.DomainSeparator(&_BridgeHub.CallOpts)
}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_BridgeHub *BridgeHubCallerSession) DomainSeparator() ([32]byte, error) {
	return _BridgeHub.Contract.DomainSeparator(&_BridgeHub.CallOpts)
}

// Epoch is a free data retrieval call binding the contract method 0x900cf0cf.
//
// Solidity: function epoch() view returns(uint64)
func (_BridgeHub *BridgeHubCaller) Epoch(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _BridgeHub.contract.Call(opts, &out, "epoch")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// Epoch is a free data retrieval call binding the contract method 0x900cf0cf.
//
// Solidity: function epoch() view returns(uint64)
func (_BridgeHub *BridgeHubSession) Epoch() (uint64, error) {
	return _BridgeHub.Contract.Epoch(&_BridgeHub.CallOpts)
}

// Epoch is a free data retrieval call binding the contract method 0x900cf0cf.
//
// Solidity: function epoch() view returns(uint64)
func (_BridgeHub *BridgeHubCallerSession) Epoch() (uint64, error) {
	return _BridgeHub.Contract.Epoch(&_BridgeHub.CallOpts)
}

// GetBridgeMessageSignatures is a free data retrieval call binding the contract method 0x446c2c9a.
//
// Solidity: function getBridgeMessageSignatures(bytes32 message) view returns(uint64 totalPower, address[] signers, (uint256,uint256,uint8)[] signatures, bytes rawData)
func (_BridgeHub *BridgeHubCaller) GetBridgeMessageSignatures(opts *bind.CallOpts, message [32]byte) (struct {
	TotalPower uint64
	Signers    []common.Address
	Signatures []Signature
	RawData    []byte
}, error) {
	var out []interface{}
	err := _BridgeHub.contract.Call(opts, &out, "getBridgeMessageSignatures", message)

	outstruct := new(struct {
		TotalPower uint64
		Signers    []common.Address
		Signatures []Signature
		RawData    []byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TotalPower = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.Signers = *abi.ConvertType(out[1], new([]common.Address)).(*[]common.Address)
	outstruct.Signatures = *abi.ConvertType(out[2], new([]Signature)).(*[]Signature)
	outstruct.RawData = *abi.ConvertType(out[3], new([]byte)).(*[]byte)

	return *outstruct, err

}

// GetBridgeMessageSignatures is a free data retrieval call binding the contract method 0x446c2c9a.
//
// Solidity: function getBridgeMessageSignatures(bytes32 message) view returns(uint64 totalPower, address[] signers, (uint256,uint256,uint8)[] signatures, bytes rawData)
func (_BridgeHub *BridgeHubSession) GetBridgeMessageSignatures(message [32]byte) (struct {
	TotalPower uint64
	Signers    []common.Address
	Signatures []Signature
	RawData    []byte
}, error) {
	return _BridgeHub.Contract.GetBridgeMessageSignatures(&_BridgeHub.CallOpts, message)
}

// GetBridgeMessageSignatures is a free data retrieval call binding the contract method 0x446c2c9a.
//
// Solidity: function getBridgeMessageSignatures(bytes32 message) view returns(uint64 totalPower, address[] signers, (uint256,uint256,uint8)[] signatures, bytes rawData)
func (_BridgeHub *BridgeHubCallerSession) GetBridgeMessageSignatures(message [32]byte) (struct {
	TotalPower uint64
	Signers    []common.Address
	Signatures []Signature
	RawData    []byte
}, error) {
	return _BridgeHub.Contract.GetBridgeMessageSignatures(&_BridgeHub.CallOpts, message)
}

// GetBridgeValidatorSignature is a free data retrieval call binding the contract method 0x6e4bc0aa.
//
// Solidity: function getBridgeValidatorSignature(bytes32 message, address signer) view returns((uint256,uint256,uint8))
func (_BridgeHub *BridgeHubCaller) GetBridgeValidatorSignature(opts *bind.CallOpts, message [32]byte, signer common.Address) (Signature, error) {
	var out []interface{}
	err := _BridgeHub.contract.Call(opts, &out, "getBridgeValidatorSignature", message, signer)

	if err != nil {
		return *new(Signature), err
	}

	out0 := *abi.ConvertType(out[0], new(Signature)).(*Signature)

	return out0, err

}

// GetBridgeValidatorSignature is a free data retrieval call binding the contract method 0x6e4bc0aa.
//
// Solidity: function getBridgeValidatorSignature(bytes32 message, address signer) view returns((uint256,uint256,uint8))
func (_BridgeHub *BridgeHubSession) GetBridgeValidatorSignature(message [32]byte, signer common.Address) (Signature, error) {
	return _BridgeHub.Contract.GetBridgeValidatorSignature(&_BridgeHub.CallOpts, message, signer)
}

// GetBridgeValidatorSignature is a free data retrieval call binding the contract method 0x6e4bc0aa.
//
// Solidity: function getBridgeValidatorSignature(bytes32 message, address signer) view returns((uint256,uint256,uint8))
func (_BridgeHub *BridgeHubCallerSession) GetBridgeValidatorSignature(message [32]byte, signer common.Address) (Signature, error) {
	return _BridgeHub.Contract.GetBridgeValidatorSignature(&_BridgeHub.CallOpts, message, signer)
}

// GetColdValidators is a free data retrieval call binding the contract method 0xc4366158.
//
// Solidity: function getColdValidators() view returns(address[])
func (_BridgeHub *BridgeHubCaller) GetColdValidators(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _BridgeHub.contract.Call(opts, &out, "getColdValidators")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetColdValidators is a free data retrieval call binding the contract method 0xc4366158.
//
// Solidity: function getColdValidators() view returns(address[])
func (_BridgeHub *BridgeHubSession) GetColdValidators() ([]common.Address, error) {
	return _BridgeHub.Contract.GetColdValidators(&_BridgeHub.CallOpts)
}

// GetColdValidators is a free data retrieval call binding the contract method 0xc4366158.
//
// Solidity: function getColdValidators() view returns(address[])
func (_BridgeHub *BridgeHubCallerSession) GetColdValidators() ([]common.Address, error) {
	return _BridgeHub.Contract.GetColdValidators(&_BridgeHub.CallOpts)
}

// GetHotValidators is a free data retrieval call binding the contract method 0xb73124b2.
//
// Solidity: function getHotValidators() view returns(address[])
func (_BridgeHub *BridgeHubCaller) GetHotValidators(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _BridgeHub.contract.Call(opts, &out, "getHotValidators")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetHotValidators is a free data retrieval call binding the contract method 0xb73124b2.
//
// Solidity: function getHotValidators() view returns(address[])
func (_BridgeHub *BridgeHubSession) GetHotValidators() ([]common.Address, error) {
	return _BridgeHub.Contract.GetHotValidators(&_BridgeHub.CallOpts)
}

// GetHotValidators is a free data retrieval call binding the contract method 0xb73124b2.
//
// Solidity: function getHotValidators() view returns(address[])
func (_BridgeHub *BridgeHubCallerSession) GetHotValidators() ([]common.Address, error) {
	return _BridgeHub.Contract.GetHotValidators(&_BridgeHub.CallOpts)
}

// GetMessageSignatures is a free data retrieval call binding the contract method 0xc7bd5b2e.
//
// Solidity: function getMessageSignatures(bytes32 message) view returns(uint64 totalPower, address[] signers, (uint256,uint256,uint8)[] signatures)
func (_BridgeHub *BridgeHubCaller) GetMessageSignatures(opts *bind.CallOpts, message [32]byte) (struct {
	TotalPower uint64
	Signers    []common.Address
	Signatures []Signature
}, error) {
	var out []interface{}
	err := _BridgeHub.contract.Call(opts, &out, "getMessageSignatures", message)

	outstruct := new(struct {
		TotalPower uint64
		Signers    []common.Address
		Signatures []Signature
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TotalPower = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.Signers = *abi.ConvertType(out[1], new([]common.Address)).(*[]common.Address)
	outstruct.Signatures = *abi.ConvertType(out[2], new([]Signature)).(*[]Signature)

	return *outstruct, err

}

// GetMessageSignatures is a free data retrieval call binding the contract method 0xc7bd5b2e.
//
// Solidity: function getMessageSignatures(bytes32 message) view returns(uint64 totalPower, address[] signers, (uint256,uint256,uint8)[] signatures)
func (_BridgeHub *BridgeHubSession) GetMessageSignatures(message [32]byte) (struct {
	TotalPower uint64
	Signers    []common.Address
	Signatures []Signature
}, error) {
	return _BridgeHub.Contract.GetMessageSignatures(&_BridgeHub.CallOpts, message)
}

// GetMessageSignatures is a free data retrieval call binding the contract method 0xc7bd5b2e.
//
// Solidity: function getMessageSignatures(bytes32 message) view returns(uint64 totalPower, address[] signers, (uint256,uint256,uint8)[] signatures)
func (_BridgeHub *BridgeHubCallerSession) GetMessageSignatures(message [32]byte) (struct {
	TotalPower uint64
	Signers    []common.Address
	Signatures []Signature
}, error) {
	return _BridgeHub.Contract.GetMessageSignatures(&_BridgeHub.CallOpts, message)
}

// GetPendingMessages is a free data retrieval call binding the contract method 0xee0791b7.
//
// Solidity: function getPendingMessages() view returns(bytes32[])
func (_BridgeHub *BridgeHubCaller) GetPendingMessages(opts *bind.CallOpts) ([][32]byte, error) {
	var out []interface{}
	err := _BridgeHub.contract.Call(opts, &out, "getPendingMessages")

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetPendingMessages is a free data retrieval call binding the contract method 0xee0791b7.
//
// Solidity: function getPendingMessages() view returns(bytes32[])
func (_BridgeHub *BridgeHubSession) GetPendingMessages() ([][32]byte, error) {
	return _BridgeHub.Contract.GetPendingMessages(&_BridgeHub.CallOpts)
}

// GetPendingMessages is a free data retrieval call binding the contract method 0xee0791b7.
//
// Solidity: function getPendingMessages() view returns(bytes32[])
func (_BridgeHub *BridgeHubCallerSession) GetPendingMessages() ([][32]byte, error) {
	return _BridgeHub.Contract.GetPendingMessages(&_BridgeHub.CallOpts)
}

// GetPendingValidatorSetUpdate is a free data retrieval call binding the contract method 0x1dd82536.
//
// Solidity: function getPendingValidatorSetUpdate() view returns(uint64 _epoch, address[] hotAddresses, address[] coldAddresses, uint64[] powers)
func (_BridgeHub *BridgeHubCaller) GetPendingValidatorSetUpdate(opts *bind.CallOpts) (struct {
	Epoch         uint64
	HotAddresses  []common.Address
	ColdAddresses []common.Address
	Powers        []uint64
}, error) {
	var out []interface{}
	err := _BridgeHub.contract.Call(opts, &out, "getPendingValidatorSetUpdate")

	outstruct := new(struct {
		Epoch         uint64
		HotAddresses  []common.Address
		ColdAddresses []common.Address
		Powers        []uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Epoch = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.HotAddresses = *abi.ConvertType(out[1], new([]common.Address)).(*[]common.Address)
	outstruct.ColdAddresses = *abi.ConvertType(out[2], new([]common.Address)).(*[]common.Address)
	outstruct.Powers = *abi.ConvertType(out[3], new([]uint64)).(*[]uint64)

	return *outstruct, err

}

// GetPendingValidatorSetUpdate is a free data retrieval call binding the contract method 0x1dd82536.
//
// Solidity: function getPendingValidatorSetUpdate() view returns(uint64 _epoch, address[] hotAddresses, address[] coldAddresses, uint64[] powers)
func (_BridgeHub *BridgeHubSession) GetPendingValidatorSetUpdate() (struct {
	Epoch         uint64
	HotAddresses  []common.Address
	ColdAddresses []common.Address
	Powers        []uint64
}, error) {
	return _BridgeHub.Contract.GetPendingValidatorSetUpdate(&_BridgeHub.CallOpts)
}

// GetPendingValidatorSetUpdate is a free data retrieval call binding the contract method 0x1dd82536.
//
// Solidity: function getPendingValidatorSetUpdate() view returns(uint64 _epoch, address[] hotAddresses, address[] coldAddresses, uint64[] powers)
func (_BridgeHub *BridgeHubCallerSession) GetPendingValidatorSetUpdate() (struct {
	Epoch         uint64
	HotAddresses  []common.Address
	ColdAddresses []common.Address
	Powers        []uint64
}, error) {
	return _BridgeHub.Contract.GetPendingValidatorSetUpdate(&_BridgeHub.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_BridgeHub *BridgeHubCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _BridgeHub.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_BridgeHub *BridgeHubSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _BridgeHub.Contract.GetRoleAdmin(&_BridgeHub.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_BridgeHub *BridgeHubCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _BridgeHub.Contract.GetRoleAdmin(&_BridgeHub.CallOpts, role)
}

// GetValidatorSignature is a free data retrieval call binding the contract method 0x2b90c338.
//
// Solidity: function getValidatorSignature(bytes32 message, address signer) view returns((uint256,uint256,uint8))
func (_BridgeHub *BridgeHubCaller) GetValidatorSignature(opts *bind.CallOpts, message [32]byte, signer common.Address) (Signature, error) {
	var out []interface{}
	err := _BridgeHub.contract.Call(opts, &out, "getValidatorSignature", message, signer)

	if err != nil {
		return *new(Signature), err
	}

	out0 := *abi.ConvertType(out[0], new(Signature)).(*Signature)

	return out0, err

}

// GetValidatorSignature is a free data retrieval call binding the contract method 0x2b90c338.
//
// Solidity: function getValidatorSignature(bytes32 message, address signer) view returns((uint256,uint256,uint8))
func (_BridgeHub *BridgeHubSession) GetValidatorSignature(message [32]byte, signer common.Address) (Signature, error) {
	return _BridgeHub.Contract.GetValidatorSignature(&_BridgeHub.CallOpts, message, signer)
}

// GetValidatorSignature is a free data retrieval call binding the contract method 0x2b90c338.
//
// Solidity: function getValidatorSignature(bytes32 message, address signer) view returns((uint256,uint256,uint8))
func (_BridgeHub *BridgeHubCallerSession) GetValidatorSignature(message [32]byte, signer common.Address) (Signature, error) {
	return _BridgeHub.Contract.GetValidatorSignature(&_BridgeHub.CallOpts, message, signer)
}

// GetValidators is a free data retrieval call binding the contract method 0xb7ab4db5.
//
// Solidity: function getValidators() view returns(uint64 _epoch, address[] validators, uint64[] powers)
func (_BridgeHub *BridgeHubCaller) GetValidators(opts *bind.CallOpts) (struct {
	Epoch      uint64
	Validators []common.Address
	Powers     []uint64
}, error) {
	var out []interface{}
	err := _BridgeHub.contract.Call(opts, &out, "getValidators")

	outstruct := new(struct {
		Epoch      uint64
		Validators []common.Address
		Powers     []uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Epoch = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.Validators = *abi.ConvertType(out[1], new([]common.Address)).(*[]common.Address)
	outstruct.Powers = *abi.ConvertType(out[2], new([]uint64)).(*[]uint64)

	return *outstruct, err

}

// GetValidators is a free data retrieval call binding the contract method 0xb7ab4db5.
//
// Solidity: function getValidators() view returns(uint64 _epoch, address[] validators, uint64[] powers)
func (_BridgeHub *BridgeHubSession) GetValidators() (struct {
	Epoch      uint64
	Validators []common.Address
	Powers     []uint64
}, error) {
	return _BridgeHub.Contract.GetValidators(&_BridgeHub.CallOpts)
}

// GetValidators is a free data retrieval call binding the contract method 0xb7ab4db5.
//
// Solidity: function getValidators() view returns(uint64 _epoch, address[] validators, uint64[] powers)
func (_BridgeHub *BridgeHubCallerSession) GetValidators() (struct {
	Epoch      uint64
	Validators []common.Address
	Powers     []uint64
}, error) {
	return _BridgeHub.Contract.GetValidators(&_BridgeHub.CallOpts)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_BridgeHub *BridgeHubCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _BridgeHub.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_BridgeHub *BridgeHubSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _BridgeHub.Contract.HasRole(&_BridgeHub.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_BridgeHub *BridgeHubCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _BridgeHub.Contract.HasRole(&_BridgeHub.CallOpts, role, account)
}

// HotValidatorList is a free data retrieval call binding the contract method 0xba14c40d.
//
// Solidity: function hotValidatorList(uint256 ) view returns(address)
func (_BridgeHub *BridgeHubCaller) HotValidatorList(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _BridgeHub.contract.Call(opts, &out, "hotValidatorList", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// HotValidatorList is a free data retrieval call binding the contract method 0xba14c40d.
//
// Solidity: function hotValidatorList(uint256 ) view returns(address)
func (_BridgeHub *BridgeHubSession) HotValidatorList(arg0 *big.Int) (common.Address, error) {
	return _BridgeHub.Contract.HotValidatorList(&_BridgeHub.CallOpts, arg0)
}

// HotValidatorList is a free data retrieval call binding the contract method 0xba14c40d.
//
// Solidity: function hotValidatorList(uint256 ) view returns(address)
func (_BridgeHub *BridgeHubCallerSession) HotValidatorList(arg0 *big.Int) (common.Address, error) {
	return _BridgeHub.Contract.HotValidatorList(&_BridgeHub.CallOpts, arg0)
}

// MakeDepositMessage is a free data retrieval call binding the contract method 0xd83ee1e8.
//
// Solidity: function makeDepositMessage(bytes32 user, address destination, bytes32 token, uint256 amount, uint256 chainId, uint64 blockNumber, bytes32 txHash, uint32 index) pure returns(bytes32)
func (_BridgeHub *BridgeHubCaller) MakeDepositMessage(opts *bind.CallOpts, user [32]byte, destination common.Address, token [32]byte, amount *big.Int, chainId *big.Int, blockNumber uint64, txHash [32]byte, index uint32) ([32]byte, error) {
	var out []interface{}
	err := _BridgeHub.contract.Call(opts, &out, "makeDepositMessage", user, destination, token, amount, chainId, blockNumber, txHash, index)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MakeDepositMessage is a free data retrieval call binding the contract method 0xd83ee1e8.
//
// Solidity: function makeDepositMessage(bytes32 user, address destination, bytes32 token, uint256 amount, uint256 chainId, uint64 blockNumber, bytes32 txHash, uint32 index) pure returns(bytes32)
func (_BridgeHub *BridgeHubSession) MakeDepositMessage(user [32]byte, destination common.Address, token [32]byte, amount *big.Int, chainId *big.Int, blockNumber uint64, txHash [32]byte, index uint32) ([32]byte, error) {
	return _BridgeHub.Contract.MakeDepositMessage(&_BridgeHub.CallOpts, user, destination, token, amount, chainId, blockNumber, txHash, index)
}

// MakeDepositMessage is a free data retrieval call binding the contract method 0xd83ee1e8.
//
// Solidity: function makeDepositMessage(bytes32 user, address destination, bytes32 token, uint256 amount, uint256 chainId, uint64 blockNumber, bytes32 txHash, uint32 index) pure returns(bytes32)
func (_BridgeHub *BridgeHubCallerSession) MakeDepositMessage(user [32]byte, destination common.Address, token [32]byte, amount *big.Int, chainId *big.Int, blockNumber uint64, txHash [32]byte, index uint32) ([32]byte, error) {
	return _BridgeHub.Contract.MakeDepositMessage(&_BridgeHub.CallOpts, user, destination, token, amount, chainId, blockNumber, txHash, index)
}

// MakeUpdateValidatorSetMessage is a free data retrieval call binding the contract method 0x64f8b391.
//
// Solidity: function makeUpdateValidatorSetMessage(uint64 newEpoch, address[] hotAddresses, address[] coldAddresses, uint64[] powers) pure returns(bytes32)
func (_BridgeHub *BridgeHubCaller) MakeUpdateValidatorSetMessage(opts *bind.CallOpts, newEpoch uint64, hotAddresses []common.Address, coldAddresses []common.Address, powers []uint64) ([32]byte, error) {
	var out []interface{}
	err := _BridgeHub.contract.Call(opts, &out, "makeUpdateValidatorSetMessage", newEpoch, hotAddresses, coldAddresses, powers)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MakeUpdateValidatorSetMessage is a free data retrieval call binding the contract method 0x64f8b391.
//
// Solidity: function makeUpdateValidatorSetMessage(uint64 newEpoch, address[] hotAddresses, address[] coldAddresses, uint64[] powers) pure returns(bytes32)
func (_BridgeHub *BridgeHubSession) MakeUpdateValidatorSetMessage(newEpoch uint64, hotAddresses []common.Address, coldAddresses []common.Address, powers []uint64) ([32]byte, error) {
	return _BridgeHub.Contract.MakeUpdateValidatorSetMessage(&_BridgeHub.CallOpts, newEpoch, hotAddresses, coldAddresses, powers)
}

// MakeUpdateValidatorSetMessage is a free data retrieval call binding the contract method 0x64f8b391.
//
// Solidity: function makeUpdateValidatorSetMessage(uint64 newEpoch, address[] hotAddresses, address[] coldAddresses, uint64[] powers) pure returns(bytes32)
func (_BridgeHub *BridgeHubCallerSession) MakeUpdateValidatorSetMessage(newEpoch uint64, hotAddresses []common.Address, coldAddresses []common.Address, powers []uint64) ([32]byte, error) {
	return _BridgeHub.Contract.MakeUpdateValidatorSetMessage(&_BridgeHub.CallOpts, newEpoch, hotAddresses, coldAddresses, powers)
}

// MakeWithdrawMessage is a free data retrieval call binding the contract method 0xbb54608d.
//
// Solidity: function makeWithdrawMessage(address user, bytes32 destination, bytes32 token, uint256 amount, uint256 chainId, uint64 nonce) pure returns(bytes32)
func (_BridgeHub *BridgeHubCaller) MakeWithdrawMessage(opts *bind.CallOpts, user common.Address, destination [32]byte, token [32]byte, amount *big.Int, chainId *big.Int, nonce uint64) ([32]byte, error) {
	var out []interface{}
	err := _BridgeHub.contract.Call(opts, &out, "makeWithdrawMessage", user, destination, token, amount, chainId, nonce)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MakeWithdrawMessage is a free data retrieval call binding the contract method 0xbb54608d.
//
// Solidity: function makeWithdrawMessage(address user, bytes32 destination, bytes32 token, uint256 amount, uint256 chainId, uint64 nonce) pure returns(bytes32)
func (_BridgeHub *BridgeHubSession) MakeWithdrawMessage(user common.Address, destination [32]byte, token [32]byte, amount *big.Int, chainId *big.Int, nonce uint64) ([32]byte, error) {
	return _BridgeHub.Contract.MakeWithdrawMessage(&_BridgeHub.CallOpts, user, destination, token, amount, chainId, nonce)
}

// MakeWithdrawMessage is a free data retrieval call binding the contract method 0xbb54608d.
//
// Solidity: function makeWithdrawMessage(address user, bytes32 destination, bytes32 token, uint256 amount, uint256 chainId, uint64 nonce) pure returns(bytes32)
func (_BridgeHub *BridgeHubCallerSession) MakeWithdrawMessage(user common.Address, destination [32]byte, token [32]byte, amount *big.Int, chainId *big.Int, nonce uint64) ([32]byte, error) {
	return _BridgeHub.Contract.MakeWithdrawMessage(&_BridgeHub.CallOpts, user, destination, token, amount, chainId, nonce)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_BridgeHub *BridgeHubCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _BridgeHub.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_BridgeHub *BridgeHubSession) Paused() (bool, error) {
	return _BridgeHub.Contract.Paused(&_BridgeHub.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_BridgeHub *BridgeHubCallerSession) Paused() (bool, error) {
	return _BridgeHub.Contract.Paused(&_BridgeHub.CallOpts)
}

// PendingMessages is a free data retrieval call binding the contract method 0x2922e6e5.
//
// Solidity: function pendingMessages(uint256 ) view returns(bytes32)
func (_BridgeHub *BridgeHubCaller) PendingMessages(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _BridgeHub.contract.Call(opts, &out, "pendingMessages", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PendingMessages is a free data retrieval call binding the contract method 0x2922e6e5.
//
// Solidity: function pendingMessages(uint256 ) view returns(bytes32)
func (_BridgeHub *BridgeHubSession) PendingMessages(arg0 *big.Int) ([32]byte, error) {
	return _BridgeHub.Contract.PendingMessages(&_BridgeHub.CallOpts, arg0)
}

// PendingMessages is a free data retrieval call binding the contract method 0x2922e6e5.
//
// Solidity: function pendingMessages(uint256 ) view returns(bytes32)
func (_BridgeHub *BridgeHubCallerSession) PendingMessages(arg0 *big.Int) ([32]byte, error) {
	return _BridgeHub.Contract.PendingMessages(&_BridgeHub.CallOpts, arg0)
}

// ProcessedMessages is a free data retrieval call binding the contract method 0x88ba16ab.
//
// Solidity: function processedMessages(bytes32 ) view returns(bool)
func (_BridgeHub *BridgeHubCaller) ProcessedMessages(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _BridgeHub.contract.Call(opts, &out, "processedMessages", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ProcessedMessages is a free data retrieval call binding the contract method 0x88ba16ab.
//
// Solidity: function processedMessages(bytes32 ) view returns(bool)
func (_BridgeHub *BridgeHubSession) ProcessedMessages(arg0 [32]byte) (bool, error) {
	return _BridgeHub.Contract.ProcessedMessages(&_BridgeHub.CallOpts, arg0)
}

// ProcessedMessages is a free data retrieval call binding the contract method 0x88ba16ab.
//
// Solidity: function processedMessages(bytes32 ) view returns(bool)
func (_BridgeHub *BridgeHubCallerSession) ProcessedMessages(arg0 [32]byte) (bool, error) {
	return _BridgeHub.Contract.ProcessedMessages(&_BridgeHub.CallOpts, arg0)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_BridgeHub *BridgeHubCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BridgeHub.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_BridgeHub *BridgeHubSession) ProxiableUUID() ([32]byte, error) {
	return _BridgeHub.Contract.ProxiableUUID(&_BridgeHub.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_BridgeHub *BridgeHubCallerSession) ProxiableUUID() ([32]byte, error) {
	return _BridgeHub.Contract.ProxiableUUID(&_BridgeHub.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_BridgeHub *BridgeHubCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _BridgeHub.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_BridgeHub *BridgeHubSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _BridgeHub.Contract.SupportsInterface(&_BridgeHub.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_BridgeHub *BridgeHubCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _BridgeHub.Contract.SupportsInterface(&_BridgeHub.CallOpts, interfaceId)
}

// TokenDecimalDiff is a free data retrieval call binding the contract method 0x06b1399f.
//
// Solidity: function tokenDecimalDiff(uint256 , bytes32 ) view returns(int8)
func (_BridgeHub *BridgeHubCaller) TokenDecimalDiff(opts *bind.CallOpts, arg0 *big.Int, arg1 [32]byte) (int8, error) {
	var out []interface{}
	err := _BridgeHub.contract.Call(opts, &out, "tokenDecimalDiff", arg0, arg1)

	if err != nil {
		return *new(int8), err
	}

	out0 := *abi.ConvertType(out[0], new(int8)).(*int8)

	return out0, err

}

// TokenDecimalDiff is a free data retrieval call binding the contract method 0x06b1399f.
//
// Solidity: function tokenDecimalDiff(uint256 , bytes32 ) view returns(int8)
func (_BridgeHub *BridgeHubSession) TokenDecimalDiff(arg0 *big.Int, arg1 [32]byte) (int8, error) {
	return _BridgeHub.Contract.TokenDecimalDiff(&_BridgeHub.CallOpts, arg0, arg1)
}

// TokenDecimalDiff is a free data retrieval call binding the contract method 0x06b1399f.
//
// Solidity: function tokenDecimalDiff(uint256 , bytes32 ) view returns(int8)
func (_BridgeHub *BridgeHubCallerSession) TokenDecimalDiff(arg0 *big.Int, arg1 [32]byte) (int8, error) {
	return _BridgeHub.Contract.TokenDecimalDiff(&_BridgeHub.CallOpts, arg0, arg1)
}

// TokenPair is a free data retrieval call binding the contract method 0x13a4cc83.
//
// Solidity: function tokenPair(uint256 , bytes32 ) view returns(bytes32)
func (_BridgeHub *BridgeHubCaller) TokenPair(opts *bind.CallOpts, arg0 *big.Int, arg1 [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _BridgeHub.contract.Call(opts, &out, "tokenPair", arg0, arg1)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TokenPair is a free data retrieval call binding the contract method 0x13a4cc83.
//
// Solidity: function tokenPair(uint256 , bytes32 ) view returns(bytes32)
func (_BridgeHub *BridgeHubSession) TokenPair(arg0 *big.Int, arg1 [32]byte) ([32]byte, error) {
	return _BridgeHub.Contract.TokenPair(&_BridgeHub.CallOpts, arg0, arg1)
}

// TokenPair is a free data retrieval call binding the contract method 0x13a4cc83.
//
// Solidity: function tokenPair(uint256 , bytes32 ) view returns(bytes32)
func (_BridgeHub *BridgeHubCallerSession) TokenPair(arg0 *big.Int, arg1 [32]byte) ([32]byte, error) {
	return _BridgeHub.Contract.TokenPair(&_BridgeHub.CallOpts, arg0, arg1)
}

// TokenWithdrawFee is a free data retrieval call binding the contract method 0x03d3f5ee.
//
// Solidity: function tokenWithdrawFee(address ) view returns(uint256)
func (_BridgeHub *BridgeHubCaller) TokenWithdrawFee(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BridgeHub.contract.Call(opts, &out, "tokenWithdrawFee", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenWithdrawFee is a free data retrieval call binding the contract method 0x03d3f5ee.
//
// Solidity: function tokenWithdrawFee(address ) view returns(uint256)
func (_BridgeHub *BridgeHubSession) TokenWithdrawFee(arg0 common.Address) (*big.Int, error) {
	return _BridgeHub.Contract.TokenWithdrawFee(&_BridgeHub.CallOpts, arg0)
}

// TokenWithdrawFee is a free data retrieval call binding the contract method 0x03d3f5ee.
//
// Solidity: function tokenWithdrawFee(address ) view returns(uint256)
func (_BridgeHub *BridgeHubCallerSession) TokenWithdrawFee(arg0 common.Address) (*big.Int, error) {
	return _BridgeHub.Contract.TokenWithdrawFee(&_BridgeHub.CallOpts, arg0)
}

// TotalValidatorPower is a free data retrieval call binding the contract method 0xf8156a6e.
//
// Solidity: function totalValidatorPower() view returns(uint64)
func (_BridgeHub *BridgeHubCaller) TotalValidatorPower(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _BridgeHub.contract.Call(opts, &out, "totalValidatorPower")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// TotalValidatorPower is a free data retrieval call binding the contract method 0xf8156a6e.
//
// Solidity: function totalValidatorPower() view returns(uint64)
func (_BridgeHub *BridgeHubSession) TotalValidatorPower() (uint64, error) {
	return _BridgeHub.Contract.TotalValidatorPower(&_BridgeHub.CallOpts)
}

// TotalValidatorPower is a free data retrieval call binding the contract method 0xf8156a6e.
//
// Solidity: function totalValidatorPower() view returns(uint64)
func (_BridgeHub *BridgeHubCallerSession) TotalValidatorPower() (uint64, error) {
	return _BridgeHub.Contract.TotalValidatorPower(&_BridgeHub.CallOpts)
}

// ValidatorPowers is a free data retrieval call binding the contract method 0xf61aa09d.
//
// Solidity: function validatorPowers(address ) view returns(uint64)
func (_BridgeHub *BridgeHubCaller) ValidatorPowers(opts *bind.CallOpts, arg0 common.Address) (uint64, error) {
	var out []interface{}
	err := _BridgeHub.contract.Call(opts, &out, "validatorPowers", arg0)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// ValidatorPowers is a free data retrieval call binding the contract method 0xf61aa09d.
//
// Solidity: function validatorPowers(address ) view returns(uint64)
func (_BridgeHub *BridgeHubSession) ValidatorPowers(arg0 common.Address) (uint64, error) {
	return _BridgeHub.Contract.ValidatorPowers(&_BridgeHub.CallOpts, arg0)
}

// ValidatorPowers is a free data retrieval call binding the contract method 0xf61aa09d.
//
// Solidity: function validatorPowers(address ) view returns(uint64)
func (_BridgeHub *BridgeHubCallerSession) ValidatorPowers(arg0 common.Address) (uint64, error) {
	return _BridgeHub.Contract.ValidatorPowers(&_BridgeHub.CallOpts, arg0)
}

// WithdrawNonce is a free data retrieval call binding the contract method 0xb8a4e151.
//
// Solidity: function withdrawNonce() view returns(uint64)
func (_BridgeHub *BridgeHubCaller) WithdrawNonce(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _BridgeHub.contract.Call(opts, &out, "withdrawNonce")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// WithdrawNonce is a free data retrieval call binding the contract method 0xb8a4e151.
//
// Solidity: function withdrawNonce() view returns(uint64)
func (_BridgeHub *BridgeHubSession) WithdrawNonce() (uint64, error) {
	return _BridgeHub.Contract.WithdrawNonce(&_BridgeHub.CallOpts)
}

// WithdrawNonce is a free data retrieval call binding the contract method 0xb8a4e151.
//
// Solidity: function withdrawNonce() view returns(uint64)
func (_BridgeHub *BridgeHubCallerSession) WithdrawNonce() (uint64, error) {
	return _BridgeHub.Contract.WithdrawNonce(&_BridgeHub.CallOpts)
}

// ClaimFees is a paid mutator transaction binding the contract method 0xceca23f1.
//
// Solidity: function claimFees(address token, address to, uint256 amount) returns()
func (_BridgeHub *BridgeHubTransactor) ClaimFees(opts *bind.TransactOpts, token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BridgeHub.contract.Transact(opts, "claimFees", token, to, amount)
}

// ClaimFees is a paid mutator transaction binding the contract method 0xceca23f1.
//
// Solidity: function claimFees(address token, address to, uint256 amount) returns()
func (_BridgeHub *BridgeHubSession) ClaimFees(token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BridgeHub.Contract.ClaimFees(&_BridgeHub.TransactOpts, token, to, amount)
}

// ClaimFees is a paid mutator transaction binding the contract method 0xceca23f1.
//
// Solidity: function claimFees(address token, address to, uint256 amount) returns()
func (_BridgeHub *BridgeHubTransactorSession) ClaimFees(token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BridgeHub.Contract.ClaimFees(&_BridgeHub.TransactOpts, token, to, amount)
}

// ClearMessageStorage is a paid mutator transaction binding the contract method 0xa8ff0075.
//
// Solidity: function clearMessageStorage(bytes32 message) returns()
func (_BridgeHub *BridgeHubTransactor) ClearMessageStorage(opts *bind.TransactOpts, message [32]byte) (*types.Transaction, error) {
	return _BridgeHub.contract.Transact(opts, "clearMessageStorage", message)
}

// ClearMessageStorage is a paid mutator transaction binding the contract method 0xa8ff0075.
//
// Solidity: function clearMessageStorage(bytes32 message) returns()
func (_BridgeHub *BridgeHubSession) ClearMessageStorage(message [32]byte) (*types.Transaction, error) {
	return _BridgeHub.Contract.ClearMessageStorage(&_BridgeHub.TransactOpts, message)
}

// ClearMessageStorage is a paid mutator transaction binding the contract method 0xa8ff0075.
//
// Solidity: function clearMessageStorage(bytes32 message) returns()
func (_BridgeHub *BridgeHubTransactorSession) ClearMessageStorage(message [32]byte) (*types.Transaction, error) {
	return _BridgeHub.Contract.ClearMessageStorage(&_BridgeHub.TransactOpts, message)
}

// DepositConfirm is a paid mutator transaction binding the contract method 0x1ebe5665.
//
// Solidity: function depositConfirm((bytes32,address,bytes32,uint256,uint256,uint64,bytes32,uint32,(uint256,uint256,uint8))[] deposits) returns()
func (_BridgeHub *BridgeHubTransactor) DepositConfirm(opts *bind.TransactOpts, deposits []DepositConfirm) (*types.Transaction, error) {
	return _BridgeHub.contract.Transact(opts, "depositConfirm", deposits)
}

// DepositConfirm is a paid mutator transaction binding the contract method 0x1ebe5665.
//
// Solidity: function depositConfirm((bytes32,address,bytes32,uint256,uint256,uint64,bytes32,uint32,(uint256,uint256,uint8))[] deposits) returns()
func (_BridgeHub *BridgeHubSession) DepositConfirm(deposits []DepositConfirm) (*types.Transaction, error) {
	return _BridgeHub.Contract.DepositConfirm(&_BridgeHub.TransactOpts, deposits)
}

// DepositConfirm is a paid mutator transaction binding the contract method 0x1ebe5665.
//
// Solidity: function depositConfirm((bytes32,address,bytes32,uint256,uint256,uint64,bytes32,uint32,(uint256,uint256,uint8))[] deposits) returns()
func (_BridgeHub *BridgeHubTransactorSession) DepositConfirm(deposits []DepositConfirm) (*types.Transaction, error) {
	return _BridgeHub.Contract.DepositConfirm(&_BridgeHub.TransactOpts, deposits)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_BridgeHub *BridgeHubTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BridgeHub.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_BridgeHub *BridgeHubSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BridgeHub.Contract.GrantRole(&_BridgeHub.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_BridgeHub *BridgeHubTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BridgeHub.Contract.GrantRole(&_BridgeHub.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0x456b07f9.
//
// Solidity: function initialize(address[] hotAddresses, address[] coldAddresses, uint64[] powers) returns()
func (_BridgeHub *BridgeHubTransactor) Initialize(opts *bind.TransactOpts, hotAddresses []common.Address, coldAddresses []common.Address, powers []uint64) (*types.Transaction, error) {
	return _BridgeHub.contract.Transact(opts, "initialize", hotAddresses, coldAddresses, powers)
}

// Initialize is a paid mutator transaction binding the contract method 0x456b07f9.
//
// Solidity: function initialize(address[] hotAddresses, address[] coldAddresses, uint64[] powers) returns()
func (_BridgeHub *BridgeHubSession) Initialize(hotAddresses []common.Address, coldAddresses []common.Address, powers []uint64) (*types.Transaction, error) {
	return _BridgeHub.Contract.Initialize(&_BridgeHub.TransactOpts, hotAddresses, coldAddresses, powers)
}

// Initialize is a paid mutator transaction binding the contract method 0x456b07f9.
//
// Solidity: function initialize(address[] hotAddresses, address[] coldAddresses, uint64[] powers) returns()
func (_BridgeHub *BridgeHubTransactorSession) Initialize(hotAddresses []common.Address, coldAddresses []common.Address, powers []uint64) (*types.Transaction, error) {
	return _BridgeHub.Contract.Initialize(&_BridgeHub.TransactOpts, hotAddresses, coldAddresses, powers)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_BridgeHub *BridgeHubTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeHub.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_BridgeHub *BridgeHubSession) Pause() (*types.Transaction, error) {
	return _BridgeHub.Contract.Pause(&_BridgeHub.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_BridgeHub *BridgeHubTransactorSession) Pause() (*types.Transaction, error) {
	return _BridgeHub.Contract.Pause(&_BridgeHub.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_BridgeHub *BridgeHubTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _BridgeHub.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_BridgeHub *BridgeHubSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _BridgeHub.Contract.RenounceRole(&_BridgeHub.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_BridgeHub *BridgeHubTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _BridgeHub.Contract.RenounceRole(&_BridgeHub.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_BridgeHub *BridgeHubTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BridgeHub.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_BridgeHub *BridgeHubSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BridgeHub.Contract.RevokeRole(&_BridgeHub.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_BridgeHub *BridgeHubTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BridgeHub.Contract.RevokeRole(&_BridgeHub.TransactOpts, role, account)
}

// SetTokenPair is a paid mutator transaction binding the contract method 0xb9d5ca9e.
//
// Solidity: function setTokenPair(uint256 chainId, bytes32 srcToken, uint8 srcTokenDecimal, address dstToken) returns()
func (_BridgeHub *BridgeHubTransactor) SetTokenPair(opts *bind.TransactOpts, chainId *big.Int, srcToken [32]byte, srcTokenDecimal uint8, dstToken common.Address) (*types.Transaction, error) {
	return _BridgeHub.contract.Transact(opts, "setTokenPair", chainId, srcToken, srcTokenDecimal, dstToken)
}

// SetTokenPair is a paid mutator transaction binding the contract method 0xb9d5ca9e.
//
// Solidity: function setTokenPair(uint256 chainId, bytes32 srcToken, uint8 srcTokenDecimal, address dstToken) returns()
func (_BridgeHub *BridgeHubSession) SetTokenPair(chainId *big.Int, srcToken [32]byte, srcTokenDecimal uint8, dstToken common.Address) (*types.Transaction, error) {
	return _BridgeHub.Contract.SetTokenPair(&_BridgeHub.TransactOpts, chainId, srcToken, srcTokenDecimal, dstToken)
}

// SetTokenPair is a paid mutator transaction binding the contract method 0xb9d5ca9e.
//
// Solidity: function setTokenPair(uint256 chainId, bytes32 srcToken, uint8 srcTokenDecimal, address dstToken) returns()
func (_BridgeHub *BridgeHubTransactorSession) SetTokenPair(chainId *big.Int, srcToken [32]byte, srcTokenDecimal uint8, dstToken common.Address) (*types.Transaction, error) {
	return _BridgeHub.Contract.SetTokenPair(&_BridgeHub.TransactOpts, chainId, srcToken, srcTokenDecimal, dstToken)
}

// SetWithdrawFee is a paid mutator transaction binding the contract method 0x54e0fb21.
//
// Solidity: function setWithdrawFee(address token, uint256 fee) returns()
func (_BridgeHub *BridgeHubTransactor) SetWithdrawFee(opts *bind.TransactOpts, token common.Address, fee *big.Int) (*types.Transaction, error) {
	return _BridgeHub.contract.Transact(opts, "setWithdrawFee", token, fee)
}

// SetWithdrawFee is a paid mutator transaction binding the contract method 0x54e0fb21.
//
// Solidity: function setWithdrawFee(address token, uint256 fee) returns()
func (_BridgeHub *BridgeHubSession) SetWithdrawFee(token common.Address, fee *big.Int) (*types.Transaction, error) {
	return _BridgeHub.Contract.SetWithdrawFee(&_BridgeHub.TransactOpts, token, fee)
}

// SetWithdrawFee is a paid mutator transaction binding the contract method 0x54e0fb21.
//
// Solidity: function setWithdrawFee(address token, uint256 fee) returns()
func (_BridgeHub *BridgeHubTransactorSession) SetWithdrawFee(token common.Address, fee *big.Int) (*types.Transaction, error) {
	return _BridgeHub.Contract.SetWithdrawFee(&_BridgeHub.TransactOpts, token, fee)
}

// SubmitBridgeSignatures is a paid mutator transaction binding the contract method 0xfa734b13.
//
// Solidity: function submitBridgeSignatures((bytes32,(uint256,uint256,uint8),bytes)[] items) returns()
func (_BridgeHub *BridgeHubTransactor) SubmitBridgeSignatures(opts *bind.TransactOpts, items []CrossChainMessage) (*types.Transaction, error) {
	return _BridgeHub.contract.Transact(opts, "submitBridgeSignatures", items)
}

// SubmitBridgeSignatures is a paid mutator transaction binding the contract method 0xfa734b13.
//
// Solidity: function submitBridgeSignatures((bytes32,(uint256,uint256,uint8),bytes)[] items) returns()
func (_BridgeHub *BridgeHubSession) SubmitBridgeSignatures(items []CrossChainMessage) (*types.Transaction, error) {
	return _BridgeHub.Contract.SubmitBridgeSignatures(&_BridgeHub.TransactOpts, items)
}

// SubmitBridgeSignatures is a paid mutator transaction binding the contract method 0xfa734b13.
//
// Solidity: function submitBridgeSignatures((bytes32,(uint256,uint256,uint8),bytes)[] items) returns()
func (_BridgeHub *BridgeHubTransactorSession) SubmitBridgeSignatures(items []CrossChainMessage) (*types.Transaction, error) {
	return _BridgeHub.Contract.SubmitBridgeSignatures(&_BridgeHub.TransactOpts, items)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_BridgeHub *BridgeHubTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeHub.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_BridgeHub *BridgeHubSession) Unpause() (*types.Transaction, error) {
	return _BridgeHub.Contract.Unpause(&_BridgeHub.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_BridgeHub *BridgeHubTransactorSession) Unpause() (*types.Transaction, error) {
	return _BridgeHub.Contract.Unpause(&_BridgeHub.TransactOpts)
}

// UpdateValidatorSet is a paid mutator transaction binding the contract method 0x3ba9613b.
//
// Solidity: function updateValidatorSet(uint64 newEpoch, address[] hotAddresses, address[] coldAddresses, uint64[] powers) returns()
func (_BridgeHub *BridgeHubTransactor) UpdateValidatorSet(opts *bind.TransactOpts, newEpoch uint64, hotAddresses []common.Address, coldAddresses []common.Address, powers []uint64) (*types.Transaction, error) {
	return _BridgeHub.contract.Transact(opts, "updateValidatorSet", newEpoch, hotAddresses, coldAddresses, powers)
}

// UpdateValidatorSet is a paid mutator transaction binding the contract method 0x3ba9613b.
//
// Solidity: function updateValidatorSet(uint64 newEpoch, address[] hotAddresses, address[] coldAddresses, uint64[] powers) returns()
func (_BridgeHub *BridgeHubSession) UpdateValidatorSet(newEpoch uint64, hotAddresses []common.Address, coldAddresses []common.Address, powers []uint64) (*types.Transaction, error) {
	return _BridgeHub.Contract.UpdateValidatorSet(&_BridgeHub.TransactOpts, newEpoch, hotAddresses, coldAddresses, powers)
}

// UpdateValidatorSet is a paid mutator transaction binding the contract method 0x3ba9613b.
//
// Solidity: function updateValidatorSet(uint64 newEpoch, address[] hotAddresses, address[] coldAddresses, uint64[] powers) returns()
func (_BridgeHub *BridgeHubTransactorSession) UpdateValidatorSet(newEpoch uint64, hotAddresses []common.Address, coldAddresses []common.Address, powers []uint64) (*types.Transaction, error) {
	return _BridgeHub.Contract.UpdateValidatorSet(&_BridgeHub.TransactOpts, newEpoch, hotAddresses, coldAddresses, powers)
}

// UpdateValidatorSetConfirm is a paid mutator transaction binding the contract method 0x17c6365d.
//
// Solidity: function updateValidatorSetConfirm((uint64,address[],address[],uint64[]) validatorSet, (uint256,uint256,uint8) signature) returns()
func (_BridgeHub *BridgeHubTransactor) UpdateValidatorSetConfirm(opts *bind.TransactOpts, validatorSet ValidatorSetUpdate, signature Signature) (*types.Transaction, error) {
	return _BridgeHub.contract.Transact(opts, "updateValidatorSetConfirm", validatorSet, signature)
}

// UpdateValidatorSetConfirm is a paid mutator transaction binding the contract method 0x17c6365d.
//
// Solidity: function updateValidatorSetConfirm((uint64,address[],address[],uint64[]) validatorSet, (uint256,uint256,uint8) signature) returns()
func (_BridgeHub *BridgeHubSession) UpdateValidatorSetConfirm(validatorSet ValidatorSetUpdate, signature Signature) (*types.Transaction, error) {
	return _BridgeHub.Contract.UpdateValidatorSetConfirm(&_BridgeHub.TransactOpts, validatorSet, signature)
}

// UpdateValidatorSetConfirm is a paid mutator transaction binding the contract method 0x17c6365d.
//
// Solidity: function updateValidatorSetConfirm((uint64,address[],address[],uint64[]) validatorSet, (uint256,uint256,uint8) signature) returns()
func (_BridgeHub *BridgeHubTransactorSession) UpdateValidatorSetConfirm(validatorSet ValidatorSetUpdate, signature Signature) (*types.Transaction, error) {
	return _BridgeHub.Contract.UpdateValidatorSetConfirm(&_BridgeHub.TransactOpts, validatorSet, signature)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_BridgeHub *BridgeHubTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _BridgeHub.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_BridgeHub *BridgeHubSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _BridgeHub.Contract.UpgradeToAndCall(&_BridgeHub.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_BridgeHub *BridgeHubTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _BridgeHub.Contract.UpgradeToAndCall(&_BridgeHub.TransactOpts, newImplementation, data)
}

// Withdraw is a paid mutator transaction binding the contract method 0xe404d8ce.
//
// Solidity: function withdraw(bytes32 destination, address token, uint256 amount, uint256 chainId) returns()
func (_BridgeHub *BridgeHubTransactor) Withdraw(opts *bind.TransactOpts, destination [32]byte, token common.Address, amount *big.Int, chainId *big.Int) (*types.Transaction, error) {
	return _BridgeHub.contract.Transact(opts, "withdraw", destination, token, amount, chainId)
}

// Withdraw is a paid mutator transaction binding the contract method 0xe404d8ce.
//
// Solidity: function withdraw(bytes32 destination, address token, uint256 amount, uint256 chainId) returns()
func (_BridgeHub *BridgeHubSession) Withdraw(destination [32]byte, token common.Address, amount *big.Int, chainId *big.Int) (*types.Transaction, error) {
	return _BridgeHub.Contract.Withdraw(&_BridgeHub.TransactOpts, destination, token, amount, chainId)
}

// Withdraw is a paid mutator transaction binding the contract method 0xe404d8ce.
//
// Solidity: function withdraw(bytes32 destination, address token, uint256 amount, uint256 chainId) returns()
func (_BridgeHub *BridgeHubTransactorSession) Withdraw(destination [32]byte, token common.Address, amount *big.Int, chainId *big.Int) (*types.Transaction, error) {
	return _BridgeHub.Contract.Withdraw(&_BridgeHub.TransactOpts, destination, token, amount, chainId)
}

// WithdrawBatchWithPermit is a paid mutator transaction binding the contract method 0x78918bb8.
//
// Solidity: function withdrawBatchWithPermit((address,bytes32,address,uint256,uint256,uint64,(uint256,uint256,uint8))[] withdraws) returns()
func (_BridgeHub *BridgeHubTransactor) WithdrawBatchWithPermit(opts *bind.TransactOpts, withdraws []WithdrawWithPermit) (*types.Transaction, error) {
	return _BridgeHub.contract.Transact(opts, "withdrawBatchWithPermit", withdraws)
}

// WithdrawBatchWithPermit is a paid mutator transaction binding the contract method 0x78918bb8.
//
// Solidity: function withdrawBatchWithPermit((address,bytes32,address,uint256,uint256,uint64,(uint256,uint256,uint8))[] withdraws) returns()
func (_BridgeHub *BridgeHubSession) WithdrawBatchWithPermit(withdraws []WithdrawWithPermit) (*types.Transaction, error) {
	return _BridgeHub.Contract.WithdrawBatchWithPermit(&_BridgeHub.TransactOpts, withdraws)
}

// WithdrawBatchWithPermit is a paid mutator transaction binding the contract method 0x78918bb8.
//
// Solidity: function withdrawBatchWithPermit((address,bytes32,address,uint256,uint256,uint64,(uint256,uint256,uint8))[] withdraws) returns()
func (_BridgeHub *BridgeHubTransactorSession) WithdrawBatchWithPermit(withdraws []WithdrawWithPermit) (*types.Transaction, error) {
	return _BridgeHub.Contract.WithdrawBatchWithPermit(&_BridgeHub.TransactOpts, withdraws)
}

// WithdrawConfirm is a paid mutator transaction binding the contract method 0x4dac5132.
//
// Solidity: function withdrawConfirm((address,bytes32,bytes32,uint256,uint256,uint64,(uint256,uint256,uint8))[] withdrawConfirms) returns()
func (_BridgeHub *BridgeHubTransactor) WithdrawConfirm(opts *bind.TransactOpts, withdrawConfirms []WithdrawConfirm) (*types.Transaction, error) {
	return _BridgeHub.contract.Transact(opts, "withdrawConfirm", withdrawConfirms)
}

// WithdrawConfirm is a paid mutator transaction binding the contract method 0x4dac5132.
//
// Solidity: function withdrawConfirm((address,bytes32,bytes32,uint256,uint256,uint64,(uint256,uint256,uint8))[] withdrawConfirms) returns()
func (_BridgeHub *BridgeHubSession) WithdrawConfirm(withdrawConfirms []WithdrawConfirm) (*types.Transaction, error) {
	return _BridgeHub.Contract.WithdrawConfirm(&_BridgeHub.TransactOpts, withdrawConfirms)
}

// WithdrawConfirm is a paid mutator transaction binding the contract method 0x4dac5132.
//
// Solidity: function withdrawConfirm((address,bytes32,bytes32,uint256,uint256,uint64,(uint256,uint256,uint8))[] withdrawConfirms) returns()
func (_BridgeHub *BridgeHubTransactorSession) WithdrawConfirm(withdrawConfirms []WithdrawConfirm) (*types.Transaction, error) {
	return _BridgeHub.Contract.WithdrawConfirm(&_BridgeHub.TransactOpts, withdrawConfirms)
}

// BridgeHubBridgeSignatureSubmittedIterator is returned from FilterBridgeSignatureSubmitted and is used to iterate over the raw logs and unpacked data for BridgeSignatureSubmitted events raised by the BridgeHub contract.
type BridgeHubBridgeSignatureSubmittedIterator struct {
	Event *BridgeHubBridgeSignatureSubmitted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BridgeHubBridgeSignatureSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeHubBridgeSignatureSubmitted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BridgeHubBridgeSignatureSubmitted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BridgeHubBridgeSignatureSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeHubBridgeSignatureSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeHubBridgeSignatureSubmitted represents a BridgeSignatureSubmitted event raised by the BridgeHub contract.
type BridgeHubBridgeSignatureSubmitted struct {
	Message    [32]byte
	Signer     common.Address
	TotalPower uint64
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterBridgeSignatureSubmitted is a free log retrieval operation binding the contract event 0xc36d3c025c0f565ea8451b940f66ac343971b6b6d515ea860aed580c6408af2b.
//
// Solidity: event BridgeSignatureSubmitted(bytes32 indexed message, address signer, uint64 totalPower)
func (_BridgeHub *BridgeHubFilterer) FilterBridgeSignatureSubmitted(opts *bind.FilterOpts, message [][32]byte) (*BridgeHubBridgeSignatureSubmittedIterator, error) {

	var messageRule []interface{}
	for _, messageItem := range message {
		messageRule = append(messageRule, messageItem)
	}

	logs, sub, err := _BridgeHub.contract.FilterLogs(opts, "BridgeSignatureSubmitted", messageRule)
	if err != nil {
		return nil, err
	}
	return &BridgeHubBridgeSignatureSubmittedIterator{contract: _BridgeHub.contract, event: "BridgeSignatureSubmitted", logs: logs, sub: sub}, nil
}

// WatchBridgeSignatureSubmitted is a free log subscription operation binding the contract event 0xc36d3c025c0f565ea8451b940f66ac343971b6b6d515ea860aed580c6408af2b.
//
// Solidity: event BridgeSignatureSubmitted(bytes32 indexed message, address signer, uint64 totalPower)
func (_BridgeHub *BridgeHubFilterer) WatchBridgeSignatureSubmitted(opts *bind.WatchOpts, sink chan<- *BridgeHubBridgeSignatureSubmitted, message [][32]byte) (event.Subscription, error) {

	var messageRule []interface{}
	for _, messageItem := range message {
		messageRule = append(messageRule, messageItem)
	}

	logs, sub, err := _BridgeHub.contract.WatchLogs(opts, "BridgeSignatureSubmitted", messageRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeHubBridgeSignatureSubmitted)
				if err := _BridgeHub.contract.UnpackLog(event, "BridgeSignatureSubmitted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBridgeSignatureSubmitted is a log parse operation binding the contract event 0xc36d3c025c0f565ea8451b940f66ac343971b6b6d515ea860aed580c6408af2b.
//
// Solidity: event BridgeSignatureSubmitted(bytes32 indexed message, address signer, uint64 totalPower)
func (_BridgeHub *BridgeHubFilterer) ParseBridgeSignatureSubmitted(log types.Log) (*BridgeHubBridgeSignatureSubmitted, error) {
	event := new(BridgeHubBridgeSignatureSubmitted)
	if err := _BridgeHub.contract.UnpackLog(event, "BridgeSignatureSubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeHubDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the BridgeHub contract.
type BridgeHubDepositIterator struct {
	Event *BridgeHubDeposit // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BridgeHubDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeHubDeposit)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BridgeHubDeposit)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BridgeHubDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeHubDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeHubDeposit represents a Deposit event raised by the BridgeHub contract.
type BridgeHubDeposit struct {
	Message     [32]byte
	User        [32]byte
	Destination common.Address
	Token       [32]byte
	Amount      *big.Int
	ChainId     *big.Int
	BlockNumber uint64
	TxHash      [32]byte
	Index       uint32
	Nonce       uint64
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xbaa0634355881c40ba9bac876b01a2b891bcaea1c37a12eba82f5bcf88d048bc.
//
// Solidity: event Deposit(bytes32 indexed message, bytes32 indexed user, address destination, bytes32 indexed token, uint256 amount, uint256 chainId, uint64 blockNumber, bytes32 txHash, uint32 index, uint64 nonce)
func (_BridgeHub *BridgeHubFilterer) FilterDeposit(opts *bind.FilterOpts, message [][32]byte, user [][32]byte, token [][32]byte) (*BridgeHubDepositIterator, error) {

	var messageRule []interface{}
	for _, messageItem := range message {
		messageRule = append(messageRule, messageItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _BridgeHub.contract.FilterLogs(opts, "Deposit", messageRule, userRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &BridgeHubDepositIterator{contract: _BridgeHub.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xbaa0634355881c40ba9bac876b01a2b891bcaea1c37a12eba82f5bcf88d048bc.
//
// Solidity: event Deposit(bytes32 indexed message, bytes32 indexed user, address destination, bytes32 indexed token, uint256 amount, uint256 chainId, uint64 blockNumber, bytes32 txHash, uint32 index, uint64 nonce)
func (_BridgeHub *BridgeHubFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *BridgeHubDeposit, message [][32]byte, user [][32]byte, token [][32]byte) (event.Subscription, error) {

	var messageRule []interface{}
	for _, messageItem := range message {
		messageRule = append(messageRule, messageItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _BridgeHub.contract.WatchLogs(opts, "Deposit", messageRule, userRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeHubDeposit)
				if err := _BridgeHub.contract.UnpackLog(event, "Deposit", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDeposit is a log parse operation binding the contract event 0xbaa0634355881c40ba9bac876b01a2b891bcaea1c37a12eba82f5bcf88d048bc.
//
// Solidity: event Deposit(bytes32 indexed message, bytes32 indexed user, address destination, bytes32 indexed token, uint256 amount, uint256 chainId, uint64 blockNumber, bytes32 txHash, uint32 index, uint64 nonce)
func (_BridgeHub *BridgeHubFilterer) ParseDeposit(log types.Log) (*BridgeHubDeposit, error) {
	event := new(BridgeHubDeposit)
	if err := _BridgeHub.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeHubFailedDepositIterator is returned from FilterFailedDeposit and is used to iterate over the raw logs and unpacked data for FailedDeposit events raised by the BridgeHub contract.
type BridgeHubFailedDepositIterator struct {
	Event *BridgeHubFailedDeposit // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BridgeHubFailedDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeHubFailedDeposit)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BridgeHubFailedDeposit)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BridgeHubFailedDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeHubFailedDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeHubFailedDeposit represents a FailedDeposit event raised by the BridgeHub contract.
type BridgeHubFailedDeposit struct {
	Message   [32]byte
	ErrorCode uint32
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterFailedDeposit is a free log retrieval operation binding the contract event 0x36dbe133addd106b8352f32d63d68f6121f6f3a4da2bd029732b370ee54afedd.
//
// Solidity: event FailedDeposit(bytes32 message, uint32 errorCode)
func (_BridgeHub *BridgeHubFilterer) FilterFailedDeposit(opts *bind.FilterOpts) (*BridgeHubFailedDepositIterator, error) {

	logs, sub, err := _BridgeHub.contract.FilterLogs(opts, "FailedDeposit")
	if err != nil {
		return nil, err
	}
	return &BridgeHubFailedDepositIterator{contract: _BridgeHub.contract, event: "FailedDeposit", logs: logs, sub: sub}, nil
}

// WatchFailedDeposit is a free log subscription operation binding the contract event 0x36dbe133addd106b8352f32d63d68f6121f6f3a4da2bd029732b370ee54afedd.
//
// Solidity: event FailedDeposit(bytes32 message, uint32 errorCode)
func (_BridgeHub *BridgeHubFilterer) WatchFailedDeposit(opts *bind.WatchOpts, sink chan<- *BridgeHubFailedDeposit) (event.Subscription, error) {

	logs, sub, err := _BridgeHub.contract.WatchLogs(opts, "FailedDeposit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeHubFailedDeposit)
				if err := _BridgeHub.contract.UnpackLog(event, "FailedDeposit", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseFailedDeposit is a log parse operation binding the contract event 0x36dbe133addd106b8352f32d63d68f6121f6f3a4da2bd029732b370ee54afedd.
//
// Solidity: event FailedDeposit(bytes32 message, uint32 errorCode)
func (_BridgeHub *BridgeHubFilterer) ParseFailedDeposit(log types.Log) (*BridgeHubFailedDeposit, error) {
	event := new(BridgeHubFailedDeposit)
	if err := _BridgeHub.contract.UnpackLog(event, "FailedDeposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeHubFeesClaimedIterator is returned from FilterFeesClaimed and is used to iterate over the raw logs and unpacked data for FeesClaimed events raised by the BridgeHub contract.
type BridgeHubFeesClaimedIterator struct {
	Event *BridgeHubFeesClaimed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BridgeHubFeesClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeHubFeesClaimed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BridgeHubFeesClaimed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BridgeHubFeesClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeHubFeesClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeHubFeesClaimed represents a FeesClaimed event raised by the BridgeHub contract.
type BridgeHubFeesClaimed struct {
	Token  common.Address
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterFeesClaimed is a free log retrieval operation binding the contract event 0xfe3464cd748424446c37877c28ce5b700222c5bc9f90d908afcc4e5cb22707ff.
//
// Solidity: event FeesClaimed(address indexed token, address to, uint256 amount)
func (_BridgeHub *BridgeHubFilterer) FilterFeesClaimed(opts *bind.FilterOpts, token []common.Address) (*BridgeHubFeesClaimedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _BridgeHub.contract.FilterLogs(opts, "FeesClaimed", tokenRule)
	if err != nil {
		return nil, err
	}
	return &BridgeHubFeesClaimedIterator{contract: _BridgeHub.contract, event: "FeesClaimed", logs: logs, sub: sub}, nil
}

// WatchFeesClaimed is a free log subscription operation binding the contract event 0xfe3464cd748424446c37877c28ce5b700222c5bc9f90d908afcc4e5cb22707ff.
//
// Solidity: event FeesClaimed(address indexed token, address to, uint256 amount)
func (_BridgeHub *BridgeHubFilterer) WatchFeesClaimed(opts *bind.WatchOpts, sink chan<- *BridgeHubFeesClaimed, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _BridgeHub.contract.WatchLogs(opts, "FeesClaimed", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeHubFeesClaimed)
				if err := _BridgeHub.contract.UnpackLog(event, "FeesClaimed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseFeesClaimed is a log parse operation binding the contract event 0xfe3464cd748424446c37877c28ce5b700222c5bc9f90d908afcc4e5cb22707ff.
//
// Solidity: event FeesClaimed(address indexed token, address to, uint256 amount)
func (_BridgeHub *BridgeHubFilterer) ParseFeesClaimed(log types.Log) (*BridgeHubFeesClaimed, error) {
	event := new(BridgeHubFeesClaimed)
	if err := _BridgeHub.contract.UnpackLog(event, "FeesClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeHubFinalizedValidatorSetUpdateIterator is returned from FilterFinalizedValidatorSetUpdate and is used to iterate over the raw logs and unpacked data for FinalizedValidatorSetUpdate events raised by the BridgeHub contract.
type BridgeHubFinalizedValidatorSetUpdateIterator struct {
	Event *BridgeHubFinalizedValidatorSetUpdate // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BridgeHubFinalizedValidatorSetUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeHubFinalizedValidatorSetUpdate)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BridgeHubFinalizedValidatorSetUpdate)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BridgeHubFinalizedValidatorSetUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeHubFinalizedValidatorSetUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeHubFinalizedValidatorSetUpdate represents a FinalizedValidatorSetUpdate event raised by the BridgeHub contract.
type BridgeHubFinalizedValidatorSetUpdate struct {
	NewEpoch      uint64
	HotAddresses  []common.Address
	ColdAddresses []common.Address
	Powers        []uint64
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterFinalizedValidatorSetUpdate is a free log retrieval operation binding the contract event 0x7fb1406cb8c05384e1f39b879d591d98e7be9ffa67cf31cdbaa10b3aebf7157c.
//
// Solidity: event FinalizedValidatorSetUpdate(uint64 indexed newEpoch, address[] hotAddresses, address[] coldAddresses, uint64[] powers)
func (_BridgeHub *BridgeHubFilterer) FilterFinalizedValidatorSetUpdate(opts *bind.FilterOpts, newEpoch []uint64) (*BridgeHubFinalizedValidatorSetUpdateIterator, error) {

	var newEpochRule []interface{}
	for _, newEpochItem := range newEpoch {
		newEpochRule = append(newEpochRule, newEpochItem)
	}

	logs, sub, err := _BridgeHub.contract.FilterLogs(opts, "FinalizedValidatorSetUpdate", newEpochRule)
	if err != nil {
		return nil, err
	}
	return &BridgeHubFinalizedValidatorSetUpdateIterator{contract: _BridgeHub.contract, event: "FinalizedValidatorSetUpdate", logs: logs, sub: sub}, nil
}

// WatchFinalizedValidatorSetUpdate is a free log subscription operation binding the contract event 0x7fb1406cb8c05384e1f39b879d591d98e7be9ffa67cf31cdbaa10b3aebf7157c.
//
// Solidity: event FinalizedValidatorSetUpdate(uint64 indexed newEpoch, address[] hotAddresses, address[] coldAddresses, uint64[] powers)
func (_BridgeHub *BridgeHubFilterer) WatchFinalizedValidatorSetUpdate(opts *bind.WatchOpts, sink chan<- *BridgeHubFinalizedValidatorSetUpdate, newEpoch []uint64) (event.Subscription, error) {

	var newEpochRule []interface{}
	for _, newEpochItem := range newEpoch {
		newEpochRule = append(newEpochRule, newEpochItem)
	}

	logs, sub, err := _BridgeHub.contract.WatchLogs(opts, "FinalizedValidatorSetUpdate", newEpochRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeHubFinalizedValidatorSetUpdate)
				if err := _BridgeHub.contract.UnpackLog(event, "FinalizedValidatorSetUpdate", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseFinalizedValidatorSetUpdate is a log parse operation binding the contract event 0x7fb1406cb8c05384e1f39b879d591d98e7be9ffa67cf31cdbaa10b3aebf7157c.
//
// Solidity: event FinalizedValidatorSetUpdate(uint64 indexed newEpoch, address[] hotAddresses, address[] coldAddresses, uint64[] powers)
func (_BridgeHub *BridgeHubFilterer) ParseFinalizedValidatorSetUpdate(log types.Log) (*BridgeHubFinalizedValidatorSetUpdate, error) {
	event := new(BridgeHubFinalizedValidatorSetUpdate)
	if err := _BridgeHub.contract.UnpackLog(event, "FinalizedValidatorSetUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeHubInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the BridgeHub contract.
type BridgeHubInitializedIterator struct {
	Event *BridgeHubInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BridgeHubInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeHubInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BridgeHubInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BridgeHubInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeHubInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeHubInitialized represents a Initialized event raised by the BridgeHub contract.
type BridgeHubInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_BridgeHub *BridgeHubFilterer) FilterInitialized(opts *bind.FilterOpts) (*BridgeHubInitializedIterator, error) {

	logs, sub, err := _BridgeHub.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &BridgeHubInitializedIterator{contract: _BridgeHub.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_BridgeHub *BridgeHubFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *BridgeHubInitialized) (event.Subscription, error) {

	logs, sub, err := _BridgeHub.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeHubInitialized)
				if err := _BridgeHub.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_BridgeHub *BridgeHubFilterer) ParseInitialized(log types.Log) (*BridgeHubInitialized, error) {
	event := new(BridgeHubInitialized)
	if err := _BridgeHub.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeHubMessageStorageClearedIterator is returned from FilterMessageStorageCleared and is used to iterate over the raw logs and unpacked data for MessageStorageCleared events raised by the BridgeHub contract.
type BridgeHubMessageStorageClearedIterator struct {
	Event *BridgeHubMessageStorageCleared // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BridgeHubMessageStorageClearedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeHubMessageStorageCleared)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BridgeHubMessageStorageCleared)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BridgeHubMessageStorageClearedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeHubMessageStorageClearedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeHubMessageStorageCleared represents a MessageStorageCleared event raised by the BridgeHub contract.
type BridgeHubMessageStorageCleared struct {
	Message [32]byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMessageStorageCleared is a free log retrieval operation binding the contract event 0xa5e9cb0cafdc12b93fa2ff5814a2a2f6ef3e98f1097c8073778b085111058f00.
//
// Solidity: event MessageStorageCleared(bytes32 indexed message)
func (_BridgeHub *BridgeHubFilterer) FilterMessageStorageCleared(opts *bind.FilterOpts, message [][32]byte) (*BridgeHubMessageStorageClearedIterator, error) {

	var messageRule []interface{}
	for _, messageItem := range message {
		messageRule = append(messageRule, messageItem)
	}

	logs, sub, err := _BridgeHub.contract.FilterLogs(opts, "MessageStorageCleared", messageRule)
	if err != nil {
		return nil, err
	}
	return &BridgeHubMessageStorageClearedIterator{contract: _BridgeHub.contract, event: "MessageStorageCleared", logs: logs, sub: sub}, nil
}

// WatchMessageStorageCleared is a free log subscription operation binding the contract event 0xa5e9cb0cafdc12b93fa2ff5814a2a2f6ef3e98f1097c8073778b085111058f00.
//
// Solidity: event MessageStorageCleared(bytes32 indexed message)
func (_BridgeHub *BridgeHubFilterer) WatchMessageStorageCleared(opts *bind.WatchOpts, sink chan<- *BridgeHubMessageStorageCleared, message [][32]byte) (event.Subscription, error) {

	var messageRule []interface{}
	for _, messageItem := range message {
		messageRule = append(messageRule, messageItem)
	}

	logs, sub, err := _BridgeHub.contract.WatchLogs(opts, "MessageStorageCleared", messageRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeHubMessageStorageCleared)
				if err := _BridgeHub.contract.UnpackLog(event, "MessageStorageCleared", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMessageStorageCleared is a log parse operation binding the contract event 0xa5e9cb0cafdc12b93fa2ff5814a2a2f6ef3e98f1097c8073778b085111058f00.
//
// Solidity: event MessageStorageCleared(bytes32 indexed message)
func (_BridgeHub *BridgeHubFilterer) ParseMessageStorageCleared(log types.Log) (*BridgeHubMessageStorageCleared, error) {
	event := new(BridgeHubMessageStorageCleared)
	if err := _BridgeHub.contract.UnpackLog(event, "MessageStorageCleared", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeHubPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the BridgeHub contract.
type BridgeHubPausedIterator struct {
	Event *BridgeHubPaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BridgeHubPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeHubPaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BridgeHubPaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BridgeHubPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeHubPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeHubPaused represents a Paused event raised by the BridgeHub contract.
type BridgeHubPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_BridgeHub *BridgeHubFilterer) FilterPaused(opts *bind.FilterOpts) (*BridgeHubPausedIterator, error) {

	logs, sub, err := _BridgeHub.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &BridgeHubPausedIterator{contract: _BridgeHub.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_BridgeHub *BridgeHubFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *BridgeHubPaused) (event.Subscription, error) {

	logs, sub, err := _BridgeHub.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeHubPaused)
				if err := _BridgeHub.contract.UnpackLog(event, "Paused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_BridgeHub *BridgeHubFilterer) ParsePaused(log types.Log) (*BridgeHubPaused, error) {
	event := new(BridgeHubPaused)
	if err := _BridgeHub.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeHubRemovedValidatorSetIterator is returned from FilterRemovedValidatorSet and is used to iterate over the raw logs and unpacked data for RemovedValidatorSet events raised by the BridgeHub contract.
type BridgeHubRemovedValidatorSetIterator struct {
	Event *BridgeHubRemovedValidatorSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BridgeHubRemovedValidatorSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeHubRemovedValidatorSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BridgeHubRemovedValidatorSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BridgeHubRemovedValidatorSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeHubRemovedValidatorSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeHubRemovedValidatorSet represents a RemovedValidatorSet event raised by the BridgeHub contract.
type BridgeHubRemovedValidatorSet struct {
	OldEpoch      uint64
	HotAddresses  []common.Address
	ColdAddresses []common.Address
	Powers        []uint64
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterRemovedValidatorSet is a free log retrieval operation binding the contract event 0x7078dd7d8f69886fa49f17529fbd86d6b24846959cd85dcfb4e990205537dc5d.
//
// Solidity: event RemovedValidatorSet(uint64 indexed oldEpoch, address[] hotAddresses, address[] coldAddresses, uint64[] powers)
func (_BridgeHub *BridgeHubFilterer) FilterRemovedValidatorSet(opts *bind.FilterOpts, oldEpoch []uint64) (*BridgeHubRemovedValidatorSetIterator, error) {

	var oldEpochRule []interface{}
	for _, oldEpochItem := range oldEpoch {
		oldEpochRule = append(oldEpochRule, oldEpochItem)
	}

	logs, sub, err := _BridgeHub.contract.FilterLogs(opts, "RemovedValidatorSet", oldEpochRule)
	if err != nil {
		return nil, err
	}
	return &BridgeHubRemovedValidatorSetIterator{contract: _BridgeHub.contract, event: "RemovedValidatorSet", logs: logs, sub: sub}, nil
}

// WatchRemovedValidatorSet is a free log subscription operation binding the contract event 0x7078dd7d8f69886fa49f17529fbd86d6b24846959cd85dcfb4e990205537dc5d.
//
// Solidity: event RemovedValidatorSet(uint64 indexed oldEpoch, address[] hotAddresses, address[] coldAddresses, uint64[] powers)
func (_BridgeHub *BridgeHubFilterer) WatchRemovedValidatorSet(opts *bind.WatchOpts, sink chan<- *BridgeHubRemovedValidatorSet, oldEpoch []uint64) (event.Subscription, error) {

	var oldEpochRule []interface{}
	for _, oldEpochItem := range oldEpoch {
		oldEpochRule = append(oldEpochRule, oldEpochItem)
	}

	logs, sub, err := _BridgeHub.contract.WatchLogs(opts, "RemovedValidatorSet", oldEpochRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeHubRemovedValidatorSet)
				if err := _BridgeHub.contract.UnpackLog(event, "RemovedValidatorSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRemovedValidatorSet is a log parse operation binding the contract event 0x7078dd7d8f69886fa49f17529fbd86d6b24846959cd85dcfb4e990205537dc5d.
//
// Solidity: event RemovedValidatorSet(uint64 indexed oldEpoch, address[] hotAddresses, address[] coldAddresses, uint64[] powers)
func (_BridgeHub *BridgeHubFilterer) ParseRemovedValidatorSet(log types.Log) (*BridgeHubRemovedValidatorSet, error) {
	event := new(BridgeHubRemovedValidatorSet)
	if err := _BridgeHub.contract.UnpackLog(event, "RemovedValidatorSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeHubRequestedValidatorSetUpdateIterator is returned from FilterRequestedValidatorSetUpdate and is used to iterate over the raw logs and unpacked data for RequestedValidatorSetUpdate events raised by the BridgeHub contract.
type BridgeHubRequestedValidatorSetUpdateIterator struct {
	Event *BridgeHubRequestedValidatorSetUpdate // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BridgeHubRequestedValidatorSetUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeHubRequestedValidatorSetUpdate)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BridgeHubRequestedValidatorSetUpdate)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BridgeHubRequestedValidatorSetUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeHubRequestedValidatorSetUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeHubRequestedValidatorSetUpdate represents a RequestedValidatorSetUpdate event raised by the BridgeHub contract.
type BridgeHubRequestedValidatorSetUpdate struct {
	NewEpoch      uint64
	HotAddresses  []common.Address
	ColdAddresses []common.Address
	Powers        []uint64
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterRequestedValidatorSetUpdate is a free log retrieval operation binding the contract event 0xf389db8d301520921c1c0eeb8eaf6790791f8916b65e8e93424a7ec831edd69e.
//
// Solidity: event RequestedValidatorSetUpdate(uint64 indexed newEpoch, address[] hotAddresses, address[] coldAddresses, uint64[] powers)
func (_BridgeHub *BridgeHubFilterer) FilterRequestedValidatorSetUpdate(opts *bind.FilterOpts, newEpoch []uint64) (*BridgeHubRequestedValidatorSetUpdateIterator, error) {

	var newEpochRule []interface{}
	for _, newEpochItem := range newEpoch {
		newEpochRule = append(newEpochRule, newEpochItem)
	}

	logs, sub, err := _BridgeHub.contract.FilterLogs(opts, "RequestedValidatorSetUpdate", newEpochRule)
	if err != nil {
		return nil, err
	}
	return &BridgeHubRequestedValidatorSetUpdateIterator{contract: _BridgeHub.contract, event: "RequestedValidatorSetUpdate", logs: logs, sub: sub}, nil
}

// WatchRequestedValidatorSetUpdate is a free log subscription operation binding the contract event 0xf389db8d301520921c1c0eeb8eaf6790791f8916b65e8e93424a7ec831edd69e.
//
// Solidity: event RequestedValidatorSetUpdate(uint64 indexed newEpoch, address[] hotAddresses, address[] coldAddresses, uint64[] powers)
func (_BridgeHub *BridgeHubFilterer) WatchRequestedValidatorSetUpdate(opts *bind.WatchOpts, sink chan<- *BridgeHubRequestedValidatorSetUpdate, newEpoch []uint64) (event.Subscription, error) {

	var newEpochRule []interface{}
	for _, newEpochItem := range newEpoch {
		newEpochRule = append(newEpochRule, newEpochItem)
	}

	logs, sub, err := _BridgeHub.contract.WatchLogs(opts, "RequestedValidatorSetUpdate", newEpochRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeHubRequestedValidatorSetUpdate)
				if err := _BridgeHub.contract.UnpackLog(event, "RequestedValidatorSetUpdate", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRequestedValidatorSetUpdate is a log parse operation binding the contract event 0xf389db8d301520921c1c0eeb8eaf6790791f8916b65e8e93424a7ec831edd69e.
//
// Solidity: event RequestedValidatorSetUpdate(uint64 indexed newEpoch, address[] hotAddresses, address[] coldAddresses, uint64[] powers)
func (_BridgeHub *BridgeHubFilterer) ParseRequestedValidatorSetUpdate(log types.Log) (*BridgeHubRequestedValidatorSetUpdate, error) {
	event := new(BridgeHubRequestedValidatorSetUpdate)
	if err := _BridgeHub.contract.UnpackLog(event, "RequestedValidatorSetUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeHubRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the BridgeHub contract.
type BridgeHubRoleAdminChangedIterator struct {
	Event *BridgeHubRoleAdminChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BridgeHubRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeHubRoleAdminChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BridgeHubRoleAdminChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BridgeHubRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeHubRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeHubRoleAdminChanged represents a RoleAdminChanged event raised by the BridgeHub contract.
type BridgeHubRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_BridgeHub *BridgeHubFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*BridgeHubRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _BridgeHub.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &BridgeHubRoleAdminChangedIterator{contract: _BridgeHub.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_BridgeHub *BridgeHubFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *BridgeHubRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _BridgeHub.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeHubRoleAdminChanged)
				if err := _BridgeHub.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_BridgeHub *BridgeHubFilterer) ParseRoleAdminChanged(log types.Log) (*BridgeHubRoleAdminChanged, error) {
	event := new(BridgeHubRoleAdminChanged)
	if err := _BridgeHub.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeHubRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the BridgeHub contract.
type BridgeHubRoleGrantedIterator struct {
	Event *BridgeHubRoleGranted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BridgeHubRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeHubRoleGranted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BridgeHubRoleGranted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BridgeHubRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeHubRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeHubRoleGranted represents a RoleGranted event raised by the BridgeHub contract.
type BridgeHubRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_BridgeHub *BridgeHubFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*BridgeHubRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _BridgeHub.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &BridgeHubRoleGrantedIterator{contract: _BridgeHub.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_BridgeHub *BridgeHubFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *BridgeHubRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _BridgeHub.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeHubRoleGranted)
				if err := _BridgeHub.contract.UnpackLog(event, "RoleGranted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_BridgeHub *BridgeHubFilterer) ParseRoleGranted(log types.Log) (*BridgeHubRoleGranted, error) {
	event := new(BridgeHubRoleGranted)
	if err := _BridgeHub.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeHubRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the BridgeHub contract.
type BridgeHubRoleRevokedIterator struct {
	Event *BridgeHubRoleRevoked // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BridgeHubRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeHubRoleRevoked)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BridgeHubRoleRevoked)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BridgeHubRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeHubRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeHubRoleRevoked represents a RoleRevoked event raised by the BridgeHub contract.
type BridgeHubRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_BridgeHub *BridgeHubFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*BridgeHubRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _BridgeHub.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &BridgeHubRoleRevokedIterator{contract: _BridgeHub.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_BridgeHub *BridgeHubFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *BridgeHubRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _BridgeHub.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeHubRoleRevoked)
				if err := _BridgeHub.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_BridgeHub *BridgeHubFilterer) ParseRoleRevoked(log types.Log) (*BridgeHubRoleRevoked, error) {
	event := new(BridgeHubRoleRevoked)
	if err := _BridgeHub.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeHubTokenPairSetIterator is returned from FilterTokenPairSet and is used to iterate over the raw logs and unpacked data for TokenPairSet events raised by the BridgeHub contract.
type BridgeHubTokenPairSetIterator struct {
	Event *BridgeHubTokenPairSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BridgeHubTokenPairSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeHubTokenPairSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BridgeHubTokenPairSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BridgeHubTokenPairSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeHubTokenPairSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeHubTokenPairSet represents a TokenPairSet event raised by the BridgeHub contract.
type BridgeHubTokenPairSet struct {
	ChainId         *big.Int
	SrcToken        [32]byte
	DstToken        common.Address
	SrcTokenDecimal uint8
	DstTokenDecimal uint8
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterTokenPairSet is a free log retrieval operation binding the contract event 0x44466fd39f128be1926fbbf7b1314fcddfb506aa13ede292765d9ea429b5a31f.
//
// Solidity: event TokenPairSet(uint256 indexed chainId, bytes32 indexed srcToken, address indexed dstToken, uint8 srcTokenDecimal, uint8 dstTokenDecimal)
func (_BridgeHub *BridgeHubFilterer) FilterTokenPairSet(opts *bind.FilterOpts, chainId []*big.Int, srcToken [][32]byte, dstToken []common.Address) (*BridgeHubTokenPairSetIterator, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}
	var srcTokenRule []interface{}
	for _, srcTokenItem := range srcToken {
		srcTokenRule = append(srcTokenRule, srcTokenItem)
	}
	var dstTokenRule []interface{}
	for _, dstTokenItem := range dstToken {
		dstTokenRule = append(dstTokenRule, dstTokenItem)
	}

	logs, sub, err := _BridgeHub.contract.FilterLogs(opts, "TokenPairSet", chainIdRule, srcTokenRule, dstTokenRule)
	if err != nil {
		return nil, err
	}
	return &BridgeHubTokenPairSetIterator{contract: _BridgeHub.contract, event: "TokenPairSet", logs: logs, sub: sub}, nil
}

// WatchTokenPairSet is a free log subscription operation binding the contract event 0x44466fd39f128be1926fbbf7b1314fcddfb506aa13ede292765d9ea429b5a31f.
//
// Solidity: event TokenPairSet(uint256 indexed chainId, bytes32 indexed srcToken, address indexed dstToken, uint8 srcTokenDecimal, uint8 dstTokenDecimal)
func (_BridgeHub *BridgeHubFilterer) WatchTokenPairSet(opts *bind.WatchOpts, sink chan<- *BridgeHubTokenPairSet, chainId []*big.Int, srcToken [][32]byte, dstToken []common.Address) (event.Subscription, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}
	var srcTokenRule []interface{}
	for _, srcTokenItem := range srcToken {
		srcTokenRule = append(srcTokenRule, srcTokenItem)
	}
	var dstTokenRule []interface{}
	for _, dstTokenItem := range dstToken {
		dstTokenRule = append(dstTokenRule, dstTokenItem)
	}

	logs, sub, err := _BridgeHub.contract.WatchLogs(opts, "TokenPairSet", chainIdRule, srcTokenRule, dstTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeHubTokenPairSet)
				if err := _BridgeHub.contract.UnpackLog(event, "TokenPairSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTokenPairSet is a log parse operation binding the contract event 0x44466fd39f128be1926fbbf7b1314fcddfb506aa13ede292765d9ea429b5a31f.
//
// Solidity: event TokenPairSet(uint256 indexed chainId, bytes32 indexed srcToken, address indexed dstToken, uint8 srcTokenDecimal, uint8 dstTokenDecimal)
func (_BridgeHub *BridgeHubFilterer) ParseTokenPairSet(log types.Log) (*BridgeHubTokenPairSet, error) {
	event := new(BridgeHubTokenPairSet)
	if err := _BridgeHub.contract.UnpackLog(event, "TokenPairSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeHubUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the BridgeHub contract.
type BridgeHubUnpausedIterator struct {
	Event *BridgeHubUnpaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BridgeHubUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeHubUnpaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BridgeHubUnpaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BridgeHubUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeHubUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeHubUnpaused represents a Unpaused event raised by the BridgeHub contract.
type BridgeHubUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_BridgeHub *BridgeHubFilterer) FilterUnpaused(opts *bind.FilterOpts) (*BridgeHubUnpausedIterator, error) {

	logs, sub, err := _BridgeHub.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &BridgeHubUnpausedIterator{contract: _BridgeHub.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_BridgeHub *BridgeHubFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *BridgeHubUnpaused) (event.Subscription, error) {

	logs, sub, err := _BridgeHub.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeHubUnpaused)
				if err := _BridgeHub.contract.UnpackLog(event, "Unpaused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_BridgeHub *BridgeHubFilterer) ParseUnpaused(log types.Log) (*BridgeHubUnpaused, error) {
	event := new(BridgeHubUnpaused)
	if err := _BridgeHub.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeHubUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the BridgeHub contract.
type BridgeHubUpgradedIterator struct {
	Event *BridgeHubUpgraded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BridgeHubUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeHubUpgraded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BridgeHubUpgraded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BridgeHubUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeHubUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeHubUpgraded represents a Upgraded event raised by the BridgeHub contract.
type BridgeHubUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_BridgeHub *BridgeHubFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*BridgeHubUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _BridgeHub.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &BridgeHubUpgradedIterator{contract: _BridgeHub.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_BridgeHub *BridgeHubFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *BridgeHubUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _BridgeHub.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeHubUpgraded)
				if err := _BridgeHub.contract.UnpackLog(event, "Upgraded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_BridgeHub *BridgeHubFilterer) ParseUpgraded(log types.Log) (*BridgeHubUpgraded, error) {
	event := new(BridgeHubUpgraded)
	if err := _BridgeHub.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeHubWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the BridgeHub contract.
type BridgeHubWithdrawIterator struct {
	Event *BridgeHubWithdraw // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BridgeHubWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeHubWithdraw)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BridgeHubWithdraw)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BridgeHubWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeHubWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeHubWithdraw represents a Withdraw event raised by the BridgeHub contract.
type BridgeHubWithdraw struct {
	Message     [32]byte
	User        common.Address
	Destination [32]byte
	Token       [32]byte
	Amount      *big.Int
	ChainId     *big.Int
	Nonce       uint64
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0xddf7473863baaba91098e18c11ea6e972fb72f38f768f5e5d3e984207414ff13.
//
// Solidity: event Withdraw(bytes32 indexed message, address indexed user, bytes32 destination, bytes32 indexed token, uint256 amount, uint256 chainId, uint64 nonce)
func (_BridgeHub *BridgeHubFilterer) FilterWithdraw(opts *bind.FilterOpts, message [][32]byte, user []common.Address, token [][32]byte) (*BridgeHubWithdrawIterator, error) {

	var messageRule []interface{}
	for _, messageItem := range message {
		messageRule = append(messageRule, messageItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _BridgeHub.contract.FilterLogs(opts, "Withdraw", messageRule, userRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &BridgeHubWithdrawIterator{contract: _BridgeHub.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0xddf7473863baaba91098e18c11ea6e972fb72f38f768f5e5d3e984207414ff13.
//
// Solidity: event Withdraw(bytes32 indexed message, address indexed user, bytes32 destination, bytes32 indexed token, uint256 amount, uint256 chainId, uint64 nonce)
func (_BridgeHub *BridgeHubFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *BridgeHubWithdraw, message [][32]byte, user []common.Address, token [][32]byte) (event.Subscription, error) {

	var messageRule []interface{}
	for _, messageItem := range message {
		messageRule = append(messageRule, messageItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _BridgeHub.contract.WatchLogs(opts, "Withdraw", messageRule, userRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeHubWithdraw)
				if err := _BridgeHub.contract.UnpackLog(event, "Withdraw", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseWithdraw is a log parse operation binding the contract event 0xddf7473863baaba91098e18c11ea6e972fb72f38f768f5e5d3e984207414ff13.
//
// Solidity: event Withdraw(bytes32 indexed message, address indexed user, bytes32 destination, bytes32 indexed token, uint256 amount, uint256 chainId, uint64 nonce)
func (_BridgeHub *BridgeHubFilterer) ParseWithdraw(log types.Log) (*BridgeHubWithdraw, error) {
	event := new(BridgeHubWithdraw)
	if err := _BridgeHub.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeHubWithdrawCompletedIterator is returned from FilterWithdrawCompleted and is used to iterate over the raw logs and unpacked data for WithdrawCompleted events raised by the BridgeHub contract.
type BridgeHubWithdrawCompletedIterator struct {
	Event *BridgeHubWithdrawCompleted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BridgeHubWithdrawCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeHubWithdrawCompleted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BridgeHubWithdrawCompleted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BridgeHubWithdrawCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeHubWithdrawCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeHubWithdrawCompleted represents a WithdrawCompleted event raised by the BridgeHub contract.
type BridgeHubWithdrawCompleted struct {
	Nonce   *big.Int
	Message [32]byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWithdrawCompleted is a free log retrieval operation binding the contract event 0xda5f7a28c1ede043b7194dba8741d889157152b30393f445fac51654abe02fae.
//
// Solidity: event WithdrawCompleted(uint256 indexed nonce, bytes32 indexed message)
func (_BridgeHub *BridgeHubFilterer) FilterWithdrawCompleted(opts *bind.FilterOpts, nonce []*big.Int, message [][32]byte) (*BridgeHubWithdrawCompletedIterator, error) {

	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}
	var messageRule []interface{}
	for _, messageItem := range message {
		messageRule = append(messageRule, messageItem)
	}

	logs, sub, err := _BridgeHub.contract.FilterLogs(opts, "WithdrawCompleted", nonceRule, messageRule)
	if err != nil {
		return nil, err
	}
	return &BridgeHubWithdrawCompletedIterator{contract: _BridgeHub.contract, event: "WithdrawCompleted", logs: logs, sub: sub}, nil
}

// WatchWithdrawCompleted is a free log subscription operation binding the contract event 0xda5f7a28c1ede043b7194dba8741d889157152b30393f445fac51654abe02fae.
//
// Solidity: event WithdrawCompleted(uint256 indexed nonce, bytes32 indexed message)
func (_BridgeHub *BridgeHubFilterer) WatchWithdrawCompleted(opts *bind.WatchOpts, sink chan<- *BridgeHubWithdrawCompleted, nonce []*big.Int, message [][32]byte) (event.Subscription, error) {

	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}
	var messageRule []interface{}
	for _, messageItem := range message {
		messageRule = append(messageRule, messageItem)
	}

	logs, sub, err := _BridgeHub.contract.WatchLogs(opts, "WithdrawCompleted", nonceRule, messageRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeHubWithdrawCompleted)
				if err := _BridgeHub.contract.UnpackLog(event, "WithdrawCompleted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseWithdrawCompleted is a log parse operation binding the contract event 0xda5f7a28c1ede043b7194dba8741d889157152b30393f445fac51654abe02fae.
//
// Solidity: event WithdrawCompleted(uint256 indexed nonce, bytes32 indexed message)
func (_BridgeHub *BridgeHubFilterer) ParseWithdrawCompleted(log types.Log) (*BridgeHubWithdrawCompleted, error) {
	event := new(BridgeHubWithdrawCompleted)
	if err := _BridgeHub.contract.UnpackLog(event, "WithdrawCompleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeHubWithdrawFeeSetIterator is returned from FilterWithdrawFeeSet and is used to iterate over the raw logs and unpacked data for WithdrawFeeSet events raised by the BridgeHub contract.
type BridgeHubWithdrawFeeSetIterator struct {
	Event *BridgeHubWithdrawFeeSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BridgeHubWithdrawFeeSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeHubWithdrawFeeSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BridgeHubWithdrawFeeSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BridgeHubWithdrawFeeSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeHubWithdrawFeeSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeHubWithdrawFeeSet represents a WithdrawFeeSet event raised by the BridgeHub contract.
type BridgeHubWithdrawFeeSet struct {
	Token common.Address
	Fee   *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterWithdrawFeeSet is a free log retrieval operation binding the contract event 0x64f4bc88c432868175af31fe623ff706067afe070c40f869f28e3daed63f8243.
//
// Solidity: event WithdrawFeeSet(address indexed token, uint256 fee)
func (_BridgeHub *BridgeHubFilterer) FilterWithdrawFeeSet(opts *bind.FilterOpts, token []common.Address) (*BridgeHubWithdrawFeeSetIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _BridgeHub.contract.FilterLogs(opts, "WithdrawFeeSet", tokenRule)
	if err != nil {
		return nil, err
	}
	return &BridgeHubWithdrawFeeSetIterator{contract: _BridgeHub.contract, event: "WithdrawFeeSet", logs: logs, sub: sub}, nil
}

// WatchWithdrawFeeSet is a free log subscription operation binding the contract event 0x64f4bc88c432868175af31fe623ff706067afe070c40f869f28e3daed63f8243.
//
// Solidity: event WithdrawFeeSet(address indexed token, uint256 fee)
func (_BridgeHub *BridgeHubFilterer) WatchWithdrawFeeSet(opts *bind.WatchOpts, sink chan<- *BridgeHubWithdrawFeeSet, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _BridgeHub.contract.WatchLogs(opts, "WithdrawFeeSet", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeHubWithdrawFeeSet)
				if err := _BridgeHub.contract.UnpackLog(event, "WithdrawFeeSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseWithdrawFeeSet is a log parse operation binding the contract event 0x64f4bc88c432868175af31fe623ff706067afe070c40f869f28e3daed63f8243.
//
// Solidity: event WithdrawFeeSet(address indexed token, uint256 fee)
func (_BridgeHub *BridgeHubFilterer) ParseWithdrawFeeSet(log types.Log) (*BridgeHubWithdrawFeeSet, error) {
	event := new(BridgeHubWithdrawFeeSet)
	if err := _BridgeHub.contract.UnpackLog(event, "WithdrawFeeSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
