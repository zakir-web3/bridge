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
	Destination common.Address
	Token       common.Address
	Amount      *big.Int
	ChainId     *big.Int
	Nonce       uint64
	Signature   Signature
}

// WithdrawWithPermit is an auto generated low-level Go binding around an user-defined struct.
type WithdrawWithPermit struct {
	User        common.Address
	Destination common.Address
	Token       common.Address
	Amount      *big.Int
	ChainId     *big.Int
	Deadline    uint64
	Signature   Signature
}

// BridgeHubMetaData contains all meta data concerning the BridgeHub contract.
var BridgeHubMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AddressInsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EnforcedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExpectedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedInnerCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"message\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"totalPower\",\"type\":\"uint64\"}],\"name\":\"BridgeSignatureSubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"message\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"user\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"destination\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"token\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"blockNumber\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"txHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"message\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"errorCode\",\"type\":\"uint32\"}],\"name\":\"FailedDeposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"FeesClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"newEpoch\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"hotAddresses\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"coldAddresses\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint64[]\",\"name\":\"powers\",\"type\":\"uint64[]\"}],\"name\":\"FinalizedValidatorSetUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"message\",\"type\":\"bytes32\"}],\"name\":\"MessageStorageCleared\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"oldEpoch\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"hotAddresses\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"coldAddresses\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint64[]\",\"name\":\"powers\",\"type\":\"uint64[]\"}],\"name\":\"RemovedValidatorSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"newEpoch\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"hotAddresses\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"coldAddresses\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint64[]\",\"name\":\"powers\",\"type\":\"uint64[]\"}],\"name\":\"RequestedValidatorSetUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"srcToken\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"dstToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"srcTokenDecimal\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"dstTokenDecimal\",\"type\":\"uint8\"}],\"name\":\"TokenPairSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"message\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"destination\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"message\",\"type\":\"bytes32\"}],\"name\":\"WithdrawCompleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"WithdrawFeeSet\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PAUSER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"claimFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"message\",\"type\":\"bytes32\"}],\"name\":\"clearMessageStorage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"coldValidatorList\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"user\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destination\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"token\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"blockNumber\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"txHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"r\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"s\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"}],\"internalType\":\"structSignature\",\"name\":\"signature\",\"type\":\"tuple\"}],\"internalType\":\"structDepositConfirm[]\",\"name\":\"deposits\",\"type\":\"tuple[]\"}],\"name\":\"depositConfirm\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"depositNonce\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"domainSeparator\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"epoch\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"message\",\"type\":\"bytes32\"}],\"name\":\"getBridgeMessageSignatures\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"totalPower\",\"type\":\"uint64\"},{\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"r\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"s\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"}],\"internalType\":\"structSignature[]\",\"name\":\"signatures\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"rawData\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"message\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"getBridgeValidatorSignature\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"r\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"s\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"}],\"internalType\":\"structSignature\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getColdValidators\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getHotValidators\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"message\",\"type\":\"bytes32\"}],\"name\":\"getMessageSignatures\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"totalPower\",\"type\":\"uint64\"},{\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"r\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"s\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"}],\"internalType\":\"structSignature[]\",\"name\":\"signatures\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPendingMessages\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPendingValidatorSetUpdate\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"_epoch\",\"type\":\"uint64\"},{\"internalType\":\"address[]\",\"name\":\"hotAddresses\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"coldAddresses\",\"type\":\"address[]\"},{\"internalType\":\"uint64[]\",\"name\":\"powers\",\"type\":\"uint64[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"message\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"getValidatorSignature\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"r\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"s\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"}],\"internalType\":\"structSignature\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getValidators\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"_epoch\",\"type\":\"uint64\"},{\"internalType\":\"address[]\",\"name\":\"validators\",\"type\":\"address[]\"},{\"internalType\":\"uint64[]\",\"name\":\"powers\",\"type\":\"uint64[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"hotValidatorList\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"hotAddresses\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"coldAddresses\",\"type\":\"address[]\"},{\"internalType\":\"uint64[]\",\"name\":\"powers\",\"type\":\"uint64[]\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"user\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destination\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"token\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"blockNumber\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"txHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"}],\"name\":\"makeDepositMessage\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"newEpoch\",\"type\":\"uint64\"},{\"internalType\":\"address[]\",\"name\":\"hotAddresses\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"coldAddresses\",\"type\":\"address[]\"},{\"internalType\":\"uint64[]\",\"name\":\"powers\",\"type\":\"uint64[]\"}],\"name\":\"makeUpdateValidatorSetMessage\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"destination\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"name\":\"makeWithdrawMessage\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"pendingMessages\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"processedMessages\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"srcToken\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"srcTokenDecimal\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"dstToken\",\"type\":\"address\"}],\"name\":\"setTokenPair\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"setWithdrawFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"domainSeparator\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"r\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"s\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"}],\"internalType\":\"structSignature\",\"name\":\"signature\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"messageRawData\",\"type\":\"bytes\"}],\"internalType\":\"structCrossChainMessage[]\",\"name\":\"items\",\"type\":\"tuple[]\"}],\"name\":\"submitBridgeSignatures\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"tokenDecimalDiff\",\"outputs\":[{\"internalType\":\"int8\",\"name\":\"\",\"type\":\"int8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"tokenPair\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"tokenWithdrawFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalValidatorPower\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"newEpoch\",\"type\":\"uint64\"},{\"internalType\":\"address[]\",\"name\":\"hotAddresses\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"coldAddresses\",\"type\":\"address[]\"},{\"internalType\":\"uint64[]\",\"name\":\"powers\",\"type\":\"uint64[]\"}],\"name\":\"updateValidatorSet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"epoch\",\"type\":\"uint64\"},{\"internalType\":\"address[]\",\"name\":\"hotAddresses\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"coldAddresses\",\"type\":\"address[]\"},{\"internalType\":\"uint64[]\",\"name\":\"powers\",\"type\":\"uint64[]\"}],\"internalType\":\"structValidatorSetUpdate\",\"name\":\"validatorSet\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"r\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"s\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"}],\"internalType\":\"structSignature\",\"name\":\"signature\",\"type\":\"tuple\"}],\"name\":\"updateValidatorSetConfirm\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"validatorPowers\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"destination\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"destination\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"deadline\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"r\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"s\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"}],\"internalType\":\"structSignature\",\"name\":\"signature\",\"type\":\"tuple\"}],\"internalType\":\"structWithdrawWithPermit[]\",\"name\":\"withdraws\",\"type\":\"tuple[]\"}],\"name\":\"withdrawBatchWithPermit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"destination\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"r\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"s\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"}],\"internalType\":\"structSignature\",\"name\":\"signature\",\"type\":\"tuple\"}],\"internalType\":\"structWithdrawConfirm[]\",\"name\":\"withdrawConfirms\",\"type\":\"tuple[]\"}],\"name\":\"withdrawConfirm\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawNonce\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60a0604052306080523480156200001557600080fd5b506200002062000026565b620000da565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000900460ff1615620000775760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160401b0390811614620000d75780546001600160401b0319166001600160401b0390811782556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50565b608051615e4462000104600039600081816139220152818161394b0152613a9c0152615e446000f3fe6080604052600436106102ff5760003560e01c80638456cb5911610190578063c4366158116100dc578063de35f5cb11610095578063f61aa09d1161006f578063f61aa09d146109e2578063f698da2514610a18578063f8156a6e14610a2e578063fa734b1314610a4e57600080fd5b8063de35f5cb1461097e578063e63ab1e91461099e578063ee0791b7146109c057600080fd5b8063c4366158146108ba578063c7bd5b2e146108cf578063c8921965146108fe578063ceca23f11461091e578063d547741f1461093e578063d83ee1e81461095e57600080fd5b8063a8ff007511610149578063b7ab4db511610123578063b7ab4db51461082f578063b8a4e15114610853578063b9d5ca9e1461087a578063ba14c40d1461089a57600080fd5b8063a8ff0075146107af578063ad3cb1cc146107cf578063b73124b21461080d57600080fd5b80638456cb59146106d657806388ba16ab146106eb578063900cf0cf1461071b57806391d148541461075a5780639adc25d81461077a578063a217fddf1461079a57600080fd5b806336568abe1161024f57806352d1902d1161020857806364f8b391116101e257806364f8b391146106545780636e4bc0aa1461067457806375b238fc146106945780637bfe950c146106b657600080fd5b806352d1902d146105fa57806354e0fb211461060f5780635c975abb1461062f57600080fd5b806336568abe146105425780633ba9613b146105625780633f4ba83a14610582578063446c2c9a14610597578063456b07f9146105c75780634f1ef286146105e757600080fd5b80631c095020116102bc578063248a9ca311610296578063248a9ca3146104b55780632922e6e5146104d55780632b90c338146104f55780632f2ff15d1461052257600080fd5b80631c095020146104385780631dd82536146104705780631ebe56651461049557600080fd5b806301ffc9a71461030457806303d3f5ee1461033957806306b1399f1461037457806313a4cc83146103c157806316a8dfee146103f657806317c6365d14610418575b600080fd5b34801561031057600080fd5b5061032461031f3660046149ee565b610a6e565b60405190151581526020015b60405180910390f35b34801561034557600080fd5b50610366610354366004614a34565b600b6020526000908152604090205481565b604051908152602001610330565b34801561038057600080fd5b506103ae61038f366004614a4f565b6001602090815260009283526040808420909152908252812054900b81565b60405160009190910b8152602001610330565b3480156103cd57600080fd5b506103666103dc366004614a4f565b600060208181529281526040808220909352908152205481565b34801561040257600080fd5b50610416610411366004614abd565b610aa5565b005b34801561042457600080fd5b50610416610433366004614afe565b610b7c565b34801561044457600080fd5b50610458610453366004614b5c565b610e05565b6040516001600160a01b039091168152602001610330565b34801561047c57600080fd5b50610485610e2f565b6040516103309493929190614bf4565b3480156104a157600080fd5b506104166104b0366004614c47565b610fc3565b3480156104c157600080fd5b506103666104d0366004614b5c565b611081565b3480156104e157600080fd5b506103666104f0366004614b5c565b6110a3565b34801561050157600080fd5b50610515610510366004614cbc565b6110c4565b6040516103309190614ce8565b34801561052e57600080fd5b5061041661053d366004614cbc565b61113e565b34801561054e57600080fd5b5061041661055d366004614cbc565b611160565b34801561056e57600080fd5b5061041661057d366004614d67565b611198565b34801561058e57600080fd5b506104166112b2565b3480156105a357600080fd5b506105b76105b2366004614b5c565b6112d5565b6040516103309493929190614eb3565b3480156105d357600080fd5b506104166105e2366004615069565b6113aa565b6104166105f53660046150f0565b6115f8565b34801561060657600080fd5b50610366611613565b34801561061b57600080fd5b5061041661062a366004615195565b611630565b34801561063b57600080fd5b50600080516020615daf8339815191525460ff16610324565b34801561066057600080fd5b5061036661066f3660046151bf565b6116f0565b34801561068057600080fd5b5061051561068f366004614cbc565b6117e4565b3480156106a057600080fd5b50610366600080516020615def83398151915281565b3480156106c257600080fd5b506104166106d1366004615257565b61185e565b3480156106e257600080fd5b50610416611892565b3480156106f757600080fd5b50610324610706366004614b5c565b600e6020526000908152604090205460ff1681565b34801561072757600080fd5b5060055461074290600160401b90046001600160401b031681565b6040516001600160401b039091168152602001610330565b34801561076657600080fd5b50610324610775366004614cbc565b6118b2565b34801561078657600080fd5b50610366610795366004615299565b6118ea565b3480156107a657600080fd5b50610366600081565b3480156107bb57600080fd5b506104166107ca366004614b5c565b611972565b3480156107db57600080fd5b50610800604051806040016040528060058152602001640352e302e360dc1b81525081565b60405161033091906152ff565b34801561081957600080fd5b50610822611993565b6040516103309190615312565b34801561083b57600080fd5b506108446119f5565b60405161033093929190615325565b34801561085f57600080fd5b50600a5461074290600160401b90046001600160401b031681565b34801561088657600080fd5b50610416610895366004615368565b611b4a565b3480156108a657600080fd5b506104586108b5366004614b5c565b611e6f565b3480156108c657600080fd5b50610822611e7f565b3480156108db57600080fd5b506108ef6108ea366004614b5c565b611edf565b604051610330939291906153b0565b34801561090a57600080fd5b50610416610919366004614abd565b611f12565b34801561092a57600080fd5b506104166109393660046153e4565b611fc1565b34801561094a57600080fd5b50610416610959366004614cbc565b6120df565b34801561096a57600080fd5b50610366610979366004615434565b6120fb565b34801561098a57600080fd5b50600a54610742906001600160401b031681565b3480156109aa57600080fd5b50610366600080516020615d6f83398151915281565b3480156109cc57600080fd5b506109d5612197565b60405161033091906154a8565b3480156109ee57600080fd5b506107426109fd366004614a34565b6004602052600090815260409020546001600160401b031681565b348015610a2457600080fd5b5061036660105481565b348015610a3a57600080fd5b50600554610742906001600160401b031681565b348015610a5a57600080fd5b50610416610a693660046154ec565b6121ee565b60006001600160e01b03198216637965db0b60e01b1480610a9f57506301ffc9a760e01b6001600160e01b03198316145b92915050565b610aad6122e6565b610ab5612319565b80610aff5760405162461bcd60e51b8152602060048201526015602482015274456d7074792077697468647261777320617272617960581b60448201526064015b60405180910390fd5b8060005b816001600160401b0316816001600160401b03161015610b5f57610b578484836001600160401b0316818110610b3b57610b3b615521565b90506101200201803603810190610b529190615642565b612351565b600101610b03565b5050610b786001600080516020615dcf83398151915255565b5050565b610b846122e6565b610b8c612319565b610b9533612437565b600554600160401b90046001600160401b0316610bb5602084018461565f565b6001600160401b031611610c0b5760405162461bcd60e51b815260206004820152601a60248201527f5374616c652076616c696461746f7220736574207570646174650000000000006044820152606401610af6565b6000610cdf610c1d602085018561565f565b610c2a602086018661567a565b80806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250610c6992505050604087018761567a565b80806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250610ca892505050606088018861567a565b808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152506116f092505050565b90506000610cec826124a5565b90506000610d0883610d03368790038701876156c3565b6124f7565b90508015610deb57610d19826126ec565b610deb610d29602087018761565f565b610d36602088018861567a565b80806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250610d7592505050604089018961567a565b80806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250610db49250505060608a018a61567a565b808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152506127d692505050565b505050610b786001600080516020615dcf83398151915255565b60028181548110610e1557600080fd5b6000918252602090912001546001600160a01b0316905081565b60408051608081018252600680546001600160401b03168252600780548451602082810282018101909652818152600095606095869586958995929491938086019390830182828015610eab57602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311610e8d575b5050505050815260200160028201805480602002602001604051908101604052809291908181526020018280548015610f0d57602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311610eef575b5050505050815260200160038201805480602002602001604051908101604052809291908181526020018280548015610f9757602002820191906000526020600020906000905b82829054906101000a90046001600160401b03166001600160401b031681526020019060080190602082600701049283019260010382029150808411610f545790505b505050919092525050815160208301516040840151606090940151919990985092965094509092505050565b610fcb6122e6565b610fd3612319565b610fdc33612437565b806110205760405162461bcd60e51b8152602060048201526014602482015273456d707479206465706f7369747320617272617960601b6044820152606401610af6565b60005b81811015611069573683838381811061103e5761103e615521565b9050610160020190506110608180360381019061105b91906156df565b612bab565b50600101611023565b50610b786001600080516020615dcf83398151915255565b6000908152600080516020615d8f833981519152602052604090206001015490565b600f81815481106110b357600080fd5b600091825260209091200154905081565b6110eb60405180606001604052806000815260200160008152602001600060ff1681525090565b506000828152600d602090815260408083206001600160a01b03851684526001908101835292819020815160608101835281548152938101549284019290925260029091015460ff169082015292915050565b61114782611081565b61115081612f71565b61115a8383612f7b565b50505050565b6001600160a01b03811633146111895760405163334bd91960e11b815260040160405180910390fd5b6111938282613027565b505050565b600080516020615def8339815191526111b081612f71565b6000886001600160401b0316116112095760405162461bcd60e51b815260206004820152601a60248201527f4e65772065706f6368206d75737420626520706f7369746976650000000000006044820152606401610af6565b6112a88888888080602002602001604051908101604052809392919081815260200183836020028082843760009201919091525050604080516020808c0282810182019093528b82529093508b92508a91829185019084908082843760009201919091525050604080516020808b0282810182019093528a82529093508a9250899182918501908490808284376000920191909152506130a392505050565b5050505050505050565b600080516020615d6f8339815191526112ca81612f71565b6112d26136db565b50565b6000818152600c60205260408120606090819081906112f38161373b565b82546002840180549397509195506001600160401b03169186918691819061131a90615773565b80601f016020809104026020016040519081016040528092919081815260200182805461134690615773565b80156113935780601f1061136857610100808354040283529160200191611393565b820191906000526020600020905b81548152906001019060200180831161137657829003601f168201915b505050505090509450945094509450509193509193565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a008054600160401b810460ff1615906001600160401b03166000811580156113ef5750825b90506000826001600160401b0316600114801561140b5750303b155b905081158015611419575080155b156114375760405163f92ee8a960e01b815260040160405180910390fd5b845467ffffffffffffffff19166001178555831561146157845460ff60401b1916600160401b1785555b6114696138ef565b6114716138ef565b6114796138f7565b611481613907565b61148c600033612f7b565b506114a5600080516020615def83398151915233612f7b565b506114be600080516020615d6f83398151915233612f7b565b506114cc60008989896130a3565b6114d960008989896127d6565b6040805180820182526009815268213934b233b2a43ab160b91b6020918201528151808301835260018152603160f81b9082015281517f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f818301527f4c4f9c18a13e136e05f5178f806d2f407f435220d27e3fd1cf60052217fef7e4818401527fc89efdaa54c0f20c7adf612882df0950f5a951637e0307cdcb4c672f298b8bc660608201524660808201523060a0808301919091528351808303909101815260c0909101909252815191012060105583156112a857845460ff60401b19168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15050505050505050565b611600613917565b611609826139bc565b610b7882826139d4565b600061161d613a91565b50600080516020615d4f83398151915290565b600080516020615def83398151915261164881612f71565b6001600160a01b0383166116965760405162461bcd60e51b8152602060048201526015602482015274496e76616c696420746f6b656e206164647265737360581b6044820152606401610af6565b6001600160a01b0383166000818152600b602052604090819020849055517f64f4bc88c432868175af31fe623ff706067afe070c40f869f28e3daed63f8243906116e39085815260200190565b60405180910390a2505050565b60007fcd26826da4f5c0e82ef8057ecacd8931dfb36167a70c820505f10826298cd05e858560405160200161172591906157ad565b604051602081830303815290604052805190602001208560405160200161174c91906157ad565b604051602081830303815290604052805190602001208560405160200161177391906157ec565b604051602081830303815290604052805190602001206040516020016117c49594939291909485526001600160401b0393909316602085015260408401919091526060830152608082015260a00190565b604051602081830303815290604052805190602001209050949350505050565b61180b60405180606001604052806000815260200160008152602001600060ff1681525090565b506000828152600c602090815260408083206001600160a01b03851684526001908101835292819020815160608101835281548152938101549284019290925260029091015460ff169082015292915050565b6118666122e6565b61186e612319565b61187b3385858585613ada565b61115a6001600080516020615dcf83398151915255565b600080516020615d6f8339815191526118aa81612f71565b6112d2613e57565b6000918252600080516020615d8f833981519152602090815260408084206001600160a01b0393909316845291905290205460ff1690565b604080517fed44fa5a448edcc9a97caee522159268a5089c8700fe54f678309d7a73c9f6ec6020808301919091526001600160a01b0398891682840152968816606082015294909616608085015260a084019290925260c08301526001600160401b031660e08083019190915283518083039091018152610100909101909252815191012090565b600080516020615def83398151915261198a81612f71565b610b7882613ea0565b606060038054806020026020016040519081016040528092919081815260200182805480156119eb57602002820191906000526020600020905b81546001600160a01b031681526001909101906020018083116119cd575b5050505050905090565b6000606080600560089054906101000a90046001600160401b031692506003805480602002602001604051908101604052809291908181526020018280548015611a6857602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311611a4a575b5050505050915081516001600160401b03811115611a8857611a88614efb565b604051908082528060200260200182016040528015611ab1578160200160208202803683370190505b50905060005b8251811015611b445760046000848381518110611ad657611ad6615521565b60200260200101516001600160a01b03166001600160a01b0316815260200190815260200160002060009054906101000a90046001600160401b0316828281518110611b2457611b24615521565b6001600160401b0390921660209283029190910190910152600101611ab7565b50909192565b600080516020615def833981519152611b6281612f71565b84600003611b825760405162461bcd60e51b8152600401610af69061581f565b83611bc35760405162461bcd60e51b815260206004820152601160248201527024b73b30b634b21039b931903a37b5b2b760791b6044820152606401610af6565b6001600160a01b038216611c0d5760405162461bcd60e51b815260206004820152601160248201527024b73b30b634b2103239ba103a37b5b2b760791b6044820152606401610af6565b60008581526020818152604080832087845290915290205415611c6b5760405162461bcd60e51b8152602060048201526016602482015275151bdad95b881c185a5c88185b1c9958591e481cd95d60521b6044820152606401610af6565b604d60ff84161115611cb65760405162461bcd60e51b8152602060048201526014602482015273496e76616c69642073726320646563696d616c7360601b6044820152606401610af6565b6000826001600160a01b031663313ce5676040518163ffffffff1660e01b8152600401602060405180830381865afa158015611cf6573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611d1a9190615848565b9050604d60ff82161115611d675760405162461bcd60e51b8152602060048201526014602482015273496e76616c69642064737420646563696d616c7360601b6044820152606401610af6565b6000611d73828661587b565b905060006001600160a01b0385166000898152602081815260408083208b845282528083208490558b8352600182528083208b84529091529020805460ff191660ff851617905590506001600160a01b038711611e19576000888152602081815260408083208484529091529020879055611ded8261589c565b60008981526001602090815260408083208584529091529020805460ff191660ff929092169190911790555b6040805160ff8089168252851660208201526001600160a01b0387169189918b917f44466fd39f128be1926fbbf7b1314fcddfb506aa13ede292765d9ea429b5a31f910160405180910390a45050505050505050565b60038181548110610e1557600080fd5b606060028054806020026020016040519081016040528092919081815260200182805480156119eb576020028201919060005260206000209081546001600160a01b031681526001909101906020018083116119cd575050505050905090565b6000818152600d602052604081206060908190611efb8161373b565b91546001600160401b031696909550909350915050565b611f1a6122e6565b611f22612319565b611f2b33612437565b80611f785760405162461bcd60e51b815260206004820152601d60248201527f456d70747920776974686472617720636f6e6669726d732061727261790000006044820152606401610af6565b60005b818110156110695736838383818110611f9657611f96615521565b905061012002019050611fb881803603810190611fb39190615642565b613f28565b50600101611f7b565b600080516020615def833981519152611fd981612f71565b6001600160a01b03831661202f5760405162461bcd60e51b815260206004820152601960248201527f496e76616c696420726563697069656e742061646472657373000000000000006044820152606401610af6565b6000821161207f5760405162461bcd60e51b815260206004820152601760248201527f416d6f756e74206d75737420626520706f7369746976650000000000000000006044820152606401610af6565b6120936001600160a01b0385168484613fc0565b604080516001600160a01b038581168252602082018590528616917ffe3464cd748424446c37877c28ce5b700222c5bc9f90d908afcc4e5cb22707ff910160405180910390a250505050565b6120e882611081565b6120f181612f71565b61115a8383613027565b604080517fe9dc4d5901ee3d0ceb0faa2d345b09188f1d7bfd6cd34ea09ed7ba75ff57391b6020808301919091528183019a909a526001600160a01b03989098166060890152608088019690965260a087019490945260c08601929092526001600160401b031660e085015261010084015263ffffffff1661012080840191909152815180840390910181526101409092019052805191012090565b6060600f8054806020026020016040519081016040528092919081815260200182805480156119eb57602002820191906000526020600020905b8154815260200190600101908083116121d1575050505050905090565b6121f66122e6565b6121fe612319565b61220733612437565b8061224d5760405162461bcd60e51b8152602060048201526016602482015275456d707479207369676e61747572657320617272617960501b6044820152606401610af6565b60005b81811015611069576122de83838381811061226d5761226d615521565b905060200281019061227f91906158ba565b60200184848481811061229457612294615521565b90506020028101906122a691906158ba565b358585858181106122b9576122b9615521565b90506020028101906122cb91906158ba565b6122d99060808101906158da565b61401f565b600101612250565b600080516020615daf8339815191525460ff16156123175760405163d93c066560e01b815260040160405180910390fd5b565b600080516020615dcf83398151915280546001190161234b57604051633ee5aeb560e01b815260040160405180910390fd5b60029055565b6040818101518251606084015160a085015160c0860151808601518151602090920151965163d505accf60e01b81526001600160a01b03958616600482015230602482015260448101949094526001600160401b03909216606484015260ff909116608483015260a482015260c4810193909352169063d505accf9060e401600060405180830381600087803b1580156123ea57600080fd5b505af11580156123fe573d6000803e3d6000fd5b505050506112d281600001518260200151836040015184606001518560800151613ada565b6001600080516020615dcf83398151915255565b6001600160a01b0381166000908152600460205260409020546001600160401b03166112d25760405162461bcd60e51b815260206004820152601960248201527f5369676e6572206973206e6f7420612076616c696461746f72000000000000006044820152606401610af6565b6000806124b183614218565b905060008111610a9f5760405162461bcd60e51b815260206004820152601160248201527013595cdcd859d9481b9bdd08199bdd5b99607a1b6044820152606401610af6565b6000828152600e602052604081205460ff161561254a5760405162461bcd60e51b8152602060048201526011602482015270105b1c9958591e481c1c9bd8d95cdcd959607a1b6044820152606401610af6565b60006125598484601054614271565b905061256481612437565b6000848152600d602090815260408083206001600160a01b0385168452600101909152902054156125e55760405162461bcd60e51b815260206004820152602560248201527f56616c696461746f7220616c7265616479207369676e65642074686973206d65604482015264737361676560d81b6064820152608401610af6565b6000848152600d602081815260408084206001600160a01b038616855260018082018452828620895181558985015191810191909155888301516002909101805460ff191660ff9092169190911790556004835290842054888552929091525461265b916001600160401b039081169116615920565b6000868152600d60205260409020805467ffffffffffffffff19166001600160401b038381169190911790915560055491925061269a91166002615940565b6001600160401b03166126ae826003615940565b6001600160401b031611156126e1575050506000828152600e60205260409020805460ff19166001908117909155610a9f565b506000949350505050565b6000811180156126fe5750600f548111155b61273a5760405162461bcd60e51b815260206004820152600d60248201526c092dcecc2d8d2c840d2dcc8caf609b1b6044820152606401610af6565b600061274760018361596b565b600f549091506127599060019061596b565b8110156127ab57600f80546127709060019061596b565b8154811061278057612780615521565b9060005260206000200154600f828154811061279e5761279e615521565b6000918252602090912001555b600f8054806127bc576127bc61597e565b600190038181906000526020600020016000905590555050565b6003546000906001600160401b038111156127f3576127f3614efb565b60405190808252806020026020018201604052801561281c578160200160208202803683370190505b50905060005b6003548110156128f85760006003828154811061284157612841615521565b60009182526020808320909101546001600160a01b031680835260049091526040909120549091506001600160401b0316156128ef576001600160a01b03811660009081526004602052604090205483516001600160401b03909116908490849081106128b0576128b0615521565b6001600160401b039092166020928302919091018201526001600160a01b0382166000908152600490915260409020805467ffffffffffffffff191690555b50600101612822565b50600560089054906101000a90046001600160401b03166001600160401b03167f7078dd7d8f69886fa49f17529fbd86d6b24846959cd85dcfb4e990205537dc5d600360028460405161294d939291906159d3565b60405180910390a2825161296890600290602086019061488b565b50835161297c90600390602087019061488b565b506000805b8551811015612a1357600086828151811061299e5761299e615521565b6020026020010151905060008583815181106129bc576129bc615521565b6020908102919091018101516001600160a01b038416600090815260049092526040909120805467ffffffffffffffff19166001600160401b0383161790559050612a078185615920565b93505050600101612981565b50600580546001600160401b03888116600160401b026fffffffffffffffffffffffffffffffff1990921690841617179055604080516080810190915260008082526020820190604051908082528060200260200182016040528015612a83578160200160208202803683370190505b5081526020016000604051908082528060200260200182016040528015612ab4578160200160208202803683370190505b5081526020016000604051908082528060200260200182016040528015612ae5578160200160208202803683370190505b50905280516006805467ffffffffffffffff19166001600160401b039092169190911781556020808301518051612b2092600792019061488b565b5060408201518051612b3c91600284019160209091019061488b565b5060608201518051612b589160038401916020909101906148f0565b50905050856001600160401b03167f7fb1406cb8c05384e1f39b879d591d98e7be9ffa67cf31cdbaa10b3aebf7157c6003600286604051612b9b939291906159d3565b60405180910390a2505050505050565b8051612be85760405162461bcd60e51b815260206004820152600c60248201526b24b73b30b634b2103ab9b2b960a11b6044820152606401610af6565b60208101516001600160a01b0316612c425760405162461bcd60e51b815260206004820152601b60248201527f496e76616c69642064657374696e6174696f6e206164647265737300000000006044820152606401610af6565b6040810151612c835760405162461bcd60e51b815260206004820152600d60248201526c24b73b30b634b2103a37b5b2b760991b6044820152606401610af6565b6000816060015111612ccb5760405162461bcd60e51b81526020600482015260116024820152700416d6f756e74206d757374206265203e3607c1b6044820152606401610af6565b8060800151600003612cef5760405162461bcd60e51b8152600401610af69061581f565b60808101516000908152602081815260408083208185015184529091529020548015801590612d2a5750612d2a816001600160a01b03101590565b612d6e5760405162461bcd60e51b815260206004820152601560248201527424b73b30b634b210313934b233b2b2103a37b5b2b760591b6044820152606401610af6565b60008190506000612da5846000015185602001518660400151876060015188608001518960a001518a60c001518b60e001516120fb565b90506000612db8828661010001516124f7565b90508015612f6a5760808501516000908152600160209081526040808320818901518452909152812054606087015190820b9190612df6908361443a565b60208801516040516340c10f1960e01b81526001600160a01b039182166004820152602481018390529192508616906340c10f1990604401600060405180830381600087803b158015612e4857600080fd5b505af1158015612e5c573d6000803e3d6000fd5b5050600a805460019350909150600090612e809084906001600160401b0316615920565b92506101000a8154816001600160401b0302191690836001600160401b0316021790555086604001518760000151857fbaa0634355881c40ba9bac876b01a2b891bcaea1c37a12eba82f5bcf88d048bc8a602001518b606001518c608001518d60a001518e60c001518f60e00151600a60009054906101000a90046001600160401b0316604051612f5f97969594939291906001600160a01b03979097168752602087019590955260408601939093526001600160401b039182166060860152608085015263ffffffff9190911660a08401521660c082015260e00190565b60405180910390a450505b5050505050565b6112d28133614494565b6000600080516020615d8f833981519152612f9684846118b2565b613016576000848152602082815260408083206001600160a01b03871684529091529020805460ff19166001179055612fcc3390565b6001600160a01b0316836001600160a01b0316857f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a46001915050610a9f565b6000915050610a9f565b5092915050565b6000600080516020615d8f83398151915261304284846118b2565b15613016576000848152602082815260408083206001600160a01b0387168085529252808320805460ff1916905551339287917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a46001915050610a9f565b60008351116130f45760405162461bcd60e51b815260206004820152601960248201527f456d70747920686f7420616464726573736573206172726179000000000000006044820152606401610af6565b81518351146131545760405162461bcd60e51b815260206004820152602660248201527f486f7420616e6420636f6c6420616464726573736573206c656e677468206d696044820152650e6dac2e8c6d60d31b6064820152608401610af6565b80518351146131b65760405162461bcd60e51b815260206004820152602860248201527f486f742061646472657373657320616e6420706f77657273206c656e677468206044820152670dad2e6dac2e8c6d60c31b6064820152608401610af6565b600554600160401b90046001600160401b03161515806131de57506001600160401b03841615155b1561325d576005546001600160401b03600160401b90910481169085161161325d5760405162461bcd60e51b815260206004820152602c60248201527f4e65772065706f6368206d7573742062652067726561746572207468616e206360448201526b0eae4e4cadce840cae0dec6d60a31b6064820152608401610af6565b6000805b845181101561355d57600085828151811061327e5761327e615521565b60200260200101519050600085838151811061329c5761329c615521565b6020026020010151905060008584815181106132ba576132ba615521565b6020026020010151905060006001600160a01b0316836001600160a01b0316036133265760405162461bcd60e51b815260206004820152601d60248201527f5a65726f206164647265737320696e20686f74206164647265737365730000006044820152606401610af6565b6001600160a01b03821661337c5760405162461bcd60e51b815260206004820152601e60248201527f5a65726f206164647265737320696e20636f6c642061646472657373657300006044820152606401610af6565b6000816001600160401b0316116133ce5760405162461bcd60e51b8152602060048201526016602482015275506f776572206d75737420626520706f73697469766560501b6044820152606401610af6565b816001600160a01b0316836001600160a01b03160361343a5760405162461bcd60e51b815260206004820152602260248201527f486f7420616e6420636f6c6420616464726573736573206d757374206469666660448201526132b960f11b6064820152608401610af6565b60005b848110156135415788818151811061345757613457615521565b60200260200101516001600160a01b0316846001600160a01b0316036134bf5760405162461bcd60e51b815260206004820152601760248201527f4475706c696361746520686f74206164647265737365730000000000000000006044820152606401610af6565b8781815181106134d1576134d1615521565b60200260200101516001600160a01b0316836001600160a01b0316036135395760405162461bcd60e51b815260206004820152601860248201527f4475706c696361746520636f6c642061646472657373657300000000000000006044820152606401610af6565b60010161343d565b5061354c8186615920565b945050600190920191506132619050565b506000816001600160401b0316116135b75760405162461bcd60e51b815260206004820152601c60248201527f546f74616c20706f776572206d75737420626520706f736974697665000000006044820152606401610af6565b60006135c5868686866116f0565b90506001600160401b0386161561360c57600f80546001810182556000919091527f8d1108e10bcb7c27dddfc02ed9d693a074039d026cf4ea4240b40f7d581ac802018190555b604080516080810182526001600160401b0388168082526020808301899052928201879052606082018690526006805467ffffffffffffffff1916909117815587519192909161366291600791908a019061488b565b506040820151805161367e91600284019160209091019061488b565b506060820151805161369a9160038401916020909101906148f0565b50905050856001600160401b03167ff389db8d301520921c1c0eeb8eaf6790791f8916b65e8e93424a7ec831edd69e868686604051612b9b93929190615a0c565b6136e36144cd565b600080516020615daf833981519152805460ff191681557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b6040516001600160a01b03909116815260200160405180910390a150565b60035460609081906000816001600160401b0381111561375d5761375d614efb565b604051908082528060200260200182016040528015613786578160200160208202803683370190505b5090506000826001600160401b038111156137a3576137a3614efb565b6040519080825280602002602001820160405280156137fb57816020015b6137e860405180606001604052806000815260200160008152602001600060ff1681525090565b8152602001906001900390816137c15790505b5090506000805b848110156138de5760006003828154811061381f5761381f615521565b60009182526020808320909101546001600160a01b031680835260018c810183526040938490208451606081018652815480825292820154948101949094526002015460ff1693830193909352925090156138d4578186858151811061388757613887615521565b60200260200101906001600160a01b031690816001600160a01b031681525050808585815181106138ba576138ba615521565b602002602001018190525083806138d090615a31565b9450505b5050600101613802565b508083528152909590945092505050565b6123176144fd565b6138ff6144fd565b612317614546565b61390f6144fd565b612317614567565b306001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016148061399e57507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316613992600080516020615d4f833981519152546001600160a01b031690565b6001600160a01b031614155b156123175760405163703e46dd60e11b815260040160405180910390fd5b600080516020615def833981519152610b7881612f71565b816001600160a01b03166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa925050508015613a2e575060408051601f3d908101601f19168201909252613a2b91810190615a4a565b60015b613a5657604051634c9c8ce360e01b81526001600160a01b0383166004820152602401610af6565b600080516020615d4f8339815191528114613a8757604051632a87526960e21b815260048101829052602401610af6565b611193838361456f565b306001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146123175760405163703e46dd60e11b815260040160405180910390fd5b6001600160a01b0383166000908152600b6020526040902054808311613b3b5760405162461bcd60e51b8152602060048201526016602482015275416d6f756e74206d757374206578636565642066656560501b6044820152606401610af6565b6001600160a01b038616613b885760405162461bcd60e51b8152602060048201526014602482015273496e76616c69642075736572206164647265737360601b6044820152606401610af6565b6001600160a01b038516613bde5760405162461bcd60e51b815260206004820152601b60248201527f496e76616c69642064657374696e6174696f6e206164647265737300000000006044820152606401610af6565b6001600160a01b038416613c2c5760405162461bcd60e51b8152602060048201526015602482015274496e76616c696420746f6b656e206164647265737360581b6044820152606401610af6565b81600003613c4c5760405162461bcd60e51b8152600401610af69061581f565b6000828152602081815260408083206001600160a01b03881680855292529091205480613cad5760405162461bcd60e51b815260206004820152600f60248201526e151bdad95b881b9bdd08199bdd5b99608a1b6044820152606401610af6565b80613cc36001600160a01b0388168a30896145c5565b613ccd848761596b565b604051630852cd8d60e31b8152600481018290529096506001600160a01b038816906342966c6890602401600060405180830381600087803b158015613d1257600080fd5b505af1158015613d26573d6000803e3d6000fd5b505050506000600a600881819054906101000a90046001600160401b0316613d4d90615a63565b82546001600160401b038083166101009490940a9384029302191691909117909155600087815260016020908152604080832088845290915281205491925090810b90613d9a898361443a565b90506000613dac8d8d87858d896118ea565b600f80546001810182556000919091527f8d1108e10bcb7c27dddfc02ed9d693a074039d026cf4ea4240b40f7d581ac80201819055604080516001600160a01b038f81168252602082018690529181018c90526001600160401b038716606082015291925080871691908f169083907f7ba97cc0a03c9de2b8d972e5a42d8867d1c3483652b679f5b8b8b5b62a4ef5769060800160405180910390a450505050505050505050505050565b613e5f6122e6565b600080516020615daf833981519152805460ff191660011781557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a2583361371d565b6000818152600d60205260408120805467ffffffffffffffff1916815590613ecb60028301826149a3565b50506000818152600c60205260408120805467ffffffffffffffff1916815590613ef860028301826149a3565b505060405181907fa5e9cb0cafdc12b93fa2ff5814a2a2f6ef3e98f1097c8073778b085111058f0090600090a250565b6000613f50826000015183602001518460400151856060015186608001518760a001516118ea565b90506000613f5d826124a5565b90506000613f6f838560c001516124f7565b9050801561115a57613f80826126ec565b60a084015160405184916001600160401b0316907fda5f7a28c1ede043b7194dba8741d889157152b30393f445fac51654abe02fae90600090a350505050565b6040516001600160a01b0383811660248301526044820183905261119391859182169063a9059cbb906064015b604051602081830303815290604052915060e01b6020820180516001600160e01b0383818316178352505050506145fe565b60105483036140615760405162461bcd60e51b815260206004820152600e60248201526d24b73b30b634b2103237b6b0b4b760911b6044820152606401610af6565b60008282604051614073929190615a89565b60405180910390209050614086816124a5565b5060006140a28261409c368990038901896156c3565b87614271565b90506140ad81612437565b6000828152600c602090815260408083206001600160a01b0385168452600101909152902054156141115760405162461bcd60e51b815260206004820152600e60248201526d105b1c9958591e481cda59db995960921b6044820152606401610af6565b6000828152600c602090815260408083206001600160a01b0385168452600101909152902086906141428282615a99565b50506000828152600c60205260409020600201614160848683615b1c565b506001600160a01b038116600090815260046020908152604080832054858452600c90925282205461419e916001600160401b039081169116615920565b6000848152600c6020908152604091829020805467ffffffffffffffff19166001600160401b03851690811790915582516001600160a01b03871681529182015291925084917fc36d3c025c0f565ea8451b940f66ac343971b6b6d515ea860aed580c6408af2b910160405180910390a250505050505050565b600f54600090815b818110156142675783600f828154811061423c5761423c615521565b90600052602060002001540361425f57614257816001615bdb565b949350505050565b600101614220565b5060009392505050565b815160009081036142c45760405162461bcd60e51b815260206004820152601b60248201527f496e76616c6964207369676e6174757265202772272076616c756500000000006044820152606401610af6565b82602001516000036143185760405162461bcd60e51b815260206004820152601b60248201527f496e76616c6964207369676e6174757265202773272076616c756500000000006044820152606401610af6565b60405161190160f01b6020820152602281018390526042810185905260009060620160408051601f19818403018152828252805160209182012087830151885189840151600080885296909401948590529195506001936143949387939193845260ff9290921660208401526040830152606082015260800190565b6020604051602081039080840390855afa1580156143b6573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b03811661442f5760405162461bcd60e51b815260206004820152602d60248201527f496e76616c6964207369676e61747572652c207265636f76657265642074686560448201526c207a65726f206164647265737360981b6064820152608401610af6565b9150505b9392505050565b60008160000b60000361444e575081610a9f565b60008260000b13156144765761446582600a615cd2565b61446f9084615ce1565b9050610a9f565b61447f8261589c565b61448a90600a615cd2565b61446f9084615d03565b61449e82826118b2565b610b785760405163e2517d3f60e01b81526001600160a01b038216600482015260248101839052604401610af6565b600080516020615daf8339815191525460ff1661231757604051638dfc202b60e01b815260040160405180910390fd5b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0054600160401b900460ff1661231757604051631afcd79f60e31b815260040160405180910390fd5b61454e6144fd565b600080516020615daf833981519152805460ff19169055565b6124236144fd565b61457882614661565b6040516001600160a01b038316907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b90600090a28051156145bd5761119382826146c6565b610b7861473c565b6040516001600160a01b03848116602483015283811660448301526064820183905261115a9186918216906323b872dd90608401613fed565b60006146136001600160a01b0384168361475b565b905080516000141580156146385750808060200190518101906146369190615d1a565b155b1561119357604051635274afe760e01b81526001600160a01b0384166004820152602401610af6565b806001600160a01b03163b60000361469757604051634c9c8ce360e01b81526001600160a01b0382166004820152602401610af6565b600080516020615d4f83398151915280546001600160a01b0319166001600160a01b0392909216919091179055565b6060600080846001600160a01b0316846040516146e39190615d3c565b600060405180830381855af49150503d806000811461471e576040519150601f19603f3d011682016040523d82523d6000602084013e614723565b606091505b5091509150614733858383614769565b95945050505050565b34156123175760405163b398979f60e01b815260040160405180910390fd5b6060614433838360006147c5565b60608261477e5761477982614862565b614433565b815115801561479557506001600160a01b0384163b155b156147be57604051639996b31560e01b81526001600160a01b0385166004820152602401610af6565b5080614433565b6060814710156147ea5760405163cd78605960e01b8152306004820152602401610af6565b600080856001600160a01b031684866040516148069190615d3c565b60006040518083038185875af1925050503d8060008114614843576040519150601f19603f3d011682016040523d82523d6000602084013e614848565b606091505b5091509150614858868383614769565b9695505050505050565b8051156148725780518082602001fd5b604051630a12f52160e11b815260040160405180910390fd5b8280548282559060005260206000209081019282156148e0579160200282015b828111156148e057825182546001600160a01b0319166001600160a01b039091161782556020909201916001909101906148ab565b506148ec9291506149d9565b5090565b828054828255906000526020600020906003016004900481019282156148e05791602002820160005b8382111561496357835183826101000a8154816001600160401b0302191690836001600160401b031602179055509260200192600801602081600701049283019260010302614919565b80156149965782816101000a8154906001600160401b030219169055600801602081600701049283019260010302614963565b50506148ec9291506149d9565b5080546149af90615773565b6000825580601f106149bf575050565b601f0160209004906000526020600020908101906112d291905b5b808211156148ec57600081556001016149da565b600060208284031215614a0057600080fd5b81356001600160e01b03198116811461443357600080fd5b80356001600160a01b0381168114614a2f57600080fd5b919050565b600060208284031215614a4657600080fd5b61443382614a18565b60008060408385031215614a6257600080fd5b50508035926020909101359150565b60008083601f840112614a8357600080fd5b5081356001600160401b03811115614a9a57600080fd5b60208301915083602061012083028501011115614ab657600080fd5b9250929050565b60008060208385031215614ad057600080fd5b82356001600160401b03811115614ae657600080fd5b614af285828601614a71565b90969095509350505050565b6000808284036080811215614b1257600080fd5b83356001600160401b03811115614b2857600080fd5b840160808187031215614b3a57600080fd5b92506060601f1982011215614b4e57600080fd5b506020830190509250929050565b600060208284031215614b6e57600080fd5b5035919050565b60008151808452602080850194506020840160005b83811015614baf5781516001600160a01b031687529582019590820190600101614b8a565b509495945050505050565b60008151808452602080850194506020840160005b83811015614baf5781516001600160401b031687529582019590820190600101614bcf565b6001600160401b0385168152608060208201526000614c166080830186614b75565b8281036040840152614c288186614b75565b90508281036060840152614c3c8185614bba565b979650505050505050565b60008060208385031215614c5a57600080fd5b82356001600160401b0380821115614c7157600080fd5b818501915085601f830112614c8557600080fd5b813581811115614c9457600080fd5b86602061016083028501011115614caa57600080fd5b60209290920196919550909350505050565b60008060408385031215614ccf57600080fd5b82359150614cdf60208401614a18565b90509250929050565b815181526020808301519082015260408083015160ff169082015260608101610a9f565b80356001600160401b0381168114614a2f57600080fd5b60008083601f840112614d3557600080fd5b5081356001600160401b03811115614d4c57600080fd5b6020830191508360208260051b8501011115614ab657600080fd5b60008060008060008060006080888a031215614d8257600080fd5b614d8b88614d0c565b965060208801356001600160401b0380821115614da757600080fd5b614db38b838c01614d23565b909850965060408a0135915080821115614dcc57600080fd5b614dd88b838c01614d23565b909650945060608a0135915080821115614df157600080fd5b50614dfe8a828b01614d23565b989b979a50959850939692959293505050565b60008151808452602080850194506020840160005b83811015614baf57614e50878351805182526020808201519083015260409081015160ff16910152565b6060969096019590820190600101614e26565b60005b83811015614e7e578181015183820152602001614e66565b50506000910152565b60008151808452614e9f816020860160208601614e63565b601f01601f19169290920160200192915050565b6001600160401b0385168152608060208201526000614ed56080830186614b75565b8281036040840152614ee78186614e11565b90508281036060840152614c3c8185614e87565b634e487b7160e01b600052604160045260246000fd5b60405161012081016001600160401b0381118282101715614f3457614f34614efb565b60405290565b604051601f8201601f191681016001600160401b0381118282101715614f6257614f62614efb565b604052919050565b60006001600160401b03821115614f8357614f83614efb565b5060051b60200190565b600082601f830112614f9e57600080fd5b81356020614fb3614fae83614f6a565b614f3a565b8083825260208201915060208460051b870101935086841115614fd557600080fd5b602086015b84811015614ff857614feb81614a18565b8352918301918301614fda565b509695505050505050565b600082601f83011261501457600080fd5b81356020615024614fae83614f6a565b8083825260208201915060208460051b87010193508684111561504657600080fd5b602086015b84811015614ff85761505c81614d0c565b835291830191830161504b565b60008060006060848603121561507e57600080fd5b83356001600160401b038082111561509557600080fd5b6150a187838801614f8d565b945060208601359150808211156150b757600080fd5b6150c387838801614f8d565b935060408601359150808211156150d957600080fd5b506150e686828701615003565b9150509250925092565b6000806040838503121561510357600080fd5b61510c83614a18565b91506020808401356001600160401b038082111561512957600080fd5b818601915086601f83011261513d57600080fd5b81358181111561514f5761514f614efb565b615161601f8201601f19168501614f3a565b9150808252878482850101111561517757600080fd5b80848401858401376000848284010152508093505050509250929050565b600080604083850312156151a857600080fd5b6151b183614a18565b946020939093013593505050565b600080600080608085870312156151d557600080fd5b6151de85614d0c565b935060208501356001600160401b03808211156151fa57600080fd5b61520688838901614f8d565b9450604087013591508082111561521c57600080fd5b61522888838901614f8d565b9350606087013591508082111561523e57600080fd5b5061524b87828801615003565b91505092959194509250565b6000806000806080858703121561526d57600080fd5b61527685614a18565b935061528460208601614a18565b93969395505050506040820135916060013590565b60008060008060008060c087890312156152b257600080fd5b6152bb87614a18565b95506152c960208801614a18565b94506152d760408801614a18565b935060608701359250608087013591506152f360a08801614d0c565b90509295509295509295565b6020815260006144336020830184614e87565b6020815260006144336020830184614b75565b6001600160401b03841681526060602082015260006153476060830185614b75565b82810360408401526148588185614bba565b60ff811681146112d257600080fd5b6000806000806080858703121561537e57600080fd5b8435935060208501359250604085013561539781615359565b91506153a560608601614a18565b905092959194509250565b6001600160401b03841681526060602082015260006153d26060830185614b75565b82810360408401526148588185614e11565b6000806000606084860312156153f957600080fd5b61540284614a18565b925061541060208501614a18565b9150604084013590509250925092565b803563ffffffff81168114614a2f57600080fd5b600080600080600080600080610100898b03121561545157600080fd5b8835975061546160208a01614a18565b965060408901359550606089013594506080890135935061548460a08a01614d0c565b925060c0890135915061549960e08a01615420565b90509295985092959890939650565b6020808252825182820181905260009190848201906040850190845b818110156154e0578351835292840192918401916001016154c4565b50909695505050505050565b600080602083850312156154ff57600080fd5b82356001600160401b0381111561551557600080fd5b614af285828601614d23565b634e487b7160e01b600052603260045260246000fd5b60006060828403121561554957600080fd5b604051606081018181106001600160401b038211171561556b5761556b614efb565b80604052508091508235815260208301356020820152604083013561558f81615359565b6040919091015292915050565b600061012082840312156155af57600080fd5b60405160e081018181106001600160401b03821117156155d1576155d1614efb565b6040529050806155e083614a18565b81526155ee60208401614a18565b60208201526155ff60408401614a18565b6040820152606083013560608201526080830135608082015261562460a08401614d0c565b60a08201526156368460c08501615537565b60c08201525092915050565b6000610120828403121561565557600080fd5b614433838361559c565b60006020828403121561567157600080fd5b61443382614d0c565b6000808335601e1984360301811261569157600080fd5b8301803591506001600160401b038211156156ab57600080fd5b6020019150600581901b3603821315614ab657600080fd5b6000606082840312156156d557600080fd5b6144338383615537565b600061016082840312156156f257600080fd5b6156fa614f11565b8235815261570a60208401614a18565b602082015260408301356040820152606083013560608201526080830135608082015261573960a08401614d0c565b60a082015260c083013560c082015261575460e08401615420565b60e082015261010061576885828601615537565b908201529392505050565b600181811c9082168061578757607f821691505b6020821081036157a757634e487b7160e01b600052602260045260246000fd5b50919050565b815160009082906020808601845b838110156157e05781516001600160a01b0316855293820193908201906001016157bb565b50929695505050505050565b815160009082906020808601845b838110156157e05781516001600160401b0316855293820193908201906001016157fa565b6020808252600f908201526e125b9d985b1a590818da185a5b9259608a1b604082015260600190565b60006020828403121561585a57600080fd5b815161443381615359565b634e487b7160e01b600052601160045260246000fd5b600082810b9082900b03607f198112607f82131715610a9f57610a9f615865565b600081810b608081016158b1576158b1615865565b60000392915050565b60008235609e198336030181126158d057600080fd5b9190910192915050565b6000808335601e198436030181126158f157600080fd5b8301803591506001600160401b0382111561590b57600080fd5b602001915036819003821315614ab657600080fd5b6001600160401b0381811683821601908082111561302057613020615865565b6001600160401b0381811683821602808216919082811461596357615963615865565b505092915050565b81810381811115610a9f57610a9f615865565b634e487b7160e01b600052603160045260246000fd5b600081548084526020808501945083600052602060002060005b83811015614baf5781546001600160a01b0316875295820195600191820191016159ae565b6060815260006159e66060830186615994565b82810360208401526159f88186615994565b905082810360408401526148588185614bba565b606081526000615a1f6060830186614b75565b82810360208401526159f88186614b75565b600060018201615a4357615a43615865565b5060010190565b600060208284031215615a5c57600080fd5b5051919050565b60006001600160401b03808316818103615a7f57615a7f615865565b6001019392505050565b8183823760009101908152919050565b8135815560208201356001820155600281016040830135615ab981615359565b815460ff191660ff919091161790555050565b601f821115611193576000816000526020600020601f850160051c81016020861015615af55750805b601f850160051c820191505b81811015615b1457828155600101615b01565b505050505050565b6001600160401b03831115615b3357615b33614efb565b615b4783615b418354615773565b83615acc565b6000601f841160018114615b7b5760008515615b635750838201355b600019600387901b1c1916600186901b178355612f6a565b600083815260209020601f19861690835b82811015615bac5786850135825560209485019460019092019101615b8c565b5086821015615bc95760001960f88860031b161c19848701351681555b505060018560011b0183555050505050565b80820180821115610a9f57610a9f615865565b600181815b80851115615c29578160001904821115615c0f57615c0f615865565b80851615615c1c57918102915b93841c9390800290615bf3565b509250929050565b600082615c4057506001610a9f565b81615c4d57506000610a9f565b8160018114615c635760028114615c6d57615c89565b6001915050610a9f565b60ff841115615c7e57615c7e615865565b50506001821b610a9f565b5060208310610133831016604e8410600b8410161715615cac575081810a610a9f565b615cb68383615bee565b8060001904821115615cca57615cca615865565b029392505050565b600061443360ff841683615c31565b600082615cfe57634e487b7160e01b600052601260045260246000fd5b500490565b8082028115828204841417610a9f57610a9f615865565b600060208284031215615d2c57600080fd5b8151801515811461443357600080fd5b600082516158d0818460208701614e6356fe360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800cd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033009b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f00a49807205ce4d355092ef5a8a18f56e8913cf4a201fbe287825b095693c21775a2646970667358221220eb56e97f8b55060fe82b7f668c16f2c9da96d2307f5eb23c6934e721c9efdaa964736f6c63430008160033",
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

// MakeWithdrawMessage is a free data retrieval call binding the contract method 0x9adc25d8.
//
// Solidity: function makeWithdrawMessage(address user, address destination, address token, uint256 amount, uint256 chainId, uint64 nonce) pure returns(bytes32)
func (_BridgeHub *BridgeHubCaller) MakeWithdrawMessage(opts *bind.CallOpts, user common.Address, destination common.Address, token common.Address, amount *big.Int, chainId *big.Int, nonce uint64) ([32]byte, error) {
	var out []interface{}
	err := _BridgeHub.contract.Call(opts, &out, "makeWithdrawMessage", user, destination, token, amount, chainId, nonce)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MakeWithdrawMessage is a free data retrieval call binding the contract method 0x9adc25d8.
//
// Solidity: function makeWithdrawMessage(address user, address destination, address token, uint256 amount, uint256 chainId, uint64 nonce) pure returns(bytes32)
func (_BridgeHub *BridgeHubSession) MakeWithdrawMessage(user common.Address, destination common.Address, token common.Address, amount *big.Int, chainId *big.Int, nonce uint64) ([32]byte, error) {
	return _BridgeHub.Contract.MakeWithdrawMessage(&_BridgeHub.CallOpts, user, destination, token, amount, chainId, nonce)
}

// MakeWithdrawMessage is a free data retrieval call binding the contract method 0x9adc25d8.
//
// Solidity: function makeWithdrawMessage(address user, address destination, address token, uint256 amount, uint256 chainId, uint64 nonce) pure returns(bytes32)
func (_BridgeHub *BridgeHubCallerSession) MakeWithdrawMessage(user common.Address, destination common.Address, token common.Address, amount *big.Int, chainId *big.Int, nonce uint64) ([32]byte, error) {
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

// Withdraw is a paid mutator transaction binding the contract method 0x7bfe950c.
//
// Solidity: function withdraw(address destination, address token, uint256 amount, uint256 chainId) returns()
func (_BridgeHub *BridgeHubTransactor) Withdraw(opts *bind.TransactOpts, destination common.Address, token common.Address, amount *big.Int, chainId *big.Int) (*types.Transaction, error) {
	return _BridgeHub.contract.Transact(opts, "withdraw", destination, token, amount, chainId)
}

// Withdraw is a paid mutator transaction binding the contract method 0x7bfe950c.
//
// Solidity: function withdraw(address destination, address token, uint256 amount, uint256 chainId) returns()
func (_BridgeHub *BridgeHubSession) Withdraw(destination common.Address, token common.Address, amount *big.Int, chainId *big.Int) (*types.Transaction, error) {
	return _BridgeHub.Contract.Withdraw(&_BridgeHub.TransactOpts, destination, token, amount, chainId)
}

// Withdraw is a paid mutator transaction binding the contract method 0x7bfe950c.
//
// Solidity: function withdraw(address destination, address token, uint256 amount, uint256 chainId) returns()
func (_BridgeHub *BridgeHubTransactorSession) Withdraw(destination common.Address, token common.Address, amount *big.Int, chainId *big.Int) (*types.Transaction, error) {
	return _BridgeHub.Contract.Withdraw(&_BridgeHub.TransactOpts, destination, token, amount, chainId)
}

// WithdrawBatchWithPermit is a paid mutator transaction binding the contract method 0x16a8dfee.
//
// Solidity: function withdrawBatchWithPermit((address,address,address,uint256,uint256,uint64,(uint256,uint256,uint8))[] withdraws) returns()
func (_BridgeHub *BridgeHubTransactor) WithdrawBatchWithPermit(opts *bind.TransactOpts, withdraws []WithdrawWithPermit) (*types.Transaction, error) {
	return _BridgeHub.contract.Transact(opts, "withdrawBatchWithPermit", withdraws)
}

// WithdrawBatchWithPermit is a paid mutator transaction binding the contract method 0x16a8dfee.
//
// Solidity: function withdrawBatchWithPermit((address,address,address,uint256,uint256,uint64,(uint256,uint256,uint8))[] withdraws) returns()
func (_BridgeHub *BridgeHubSession) WithdrawBatchWithPermit(withdraws []WithdrawWithPermit) (*types.Transaction, error) {
	return _BridgeHub.Contract.WithdrawBatchWithPermit(&_BridgeHub.TransactOpts, withdraws)
}

// WithdrawBatchWithPermit is a paid mutator transaction binding the contract method 0x16a8dfee.
//
// Solidity: function withdrawBatchWithPermit((address,address,address,uint256,uint256,uint64,(uint256,uint256,uint8))[] withdraws) returns()
func (_BridgeHub *BridgeHubTransactorSession) WithdrawBatchWithPermit(withdraws []WithdrawWithPermit) (*types.Transaction, error) {
	return _BridgeHub.Contract.WithdrawBatchWithPermit(&_BridgeHub.TransactOpts, withdraws)
}

// WithdrawConfirm is a paid mutator transaction binding the contract method 0xc8921965.
//
// Solidity: function withdrawConfirm((address,address,address,uint256,uint256,uint64,(uint256,uint256,uint8))[] withdrawConfirms) returns()
func (_BridgeHub *BridgeHubTransactor) WithdrawConfirm(opts *bind.TransactOpts, withdrawConfirms []WithdrawConfirm) (*types.Transaction, error) {
	return _BridgeHub.contract.Transact(opts, "withdrawConfirm", withdrawConfirms)
}

// WithdrawConfirm is a paid mutator transaction binding the contract method 0xc8921965.
//
// Solidity: function withdrawConfirm((address,address,address,uint256,uint256,uint64,(uint256,uint256,uint8))[] withdrawConfirms) returns()
func (_BridgeHub *BridgeHubSession) WithdrawConfirm(withdrawConfirms []WithdrawConfirm) (*types.Transaction, error) {
	return _BridgeHub.Contract.WithdrawConfirm(&_BridgeHub.TransactOpts, withdrawConfirms)
}

// WithdrawConfirm is a paid mutator transaction binding the contract method 0xc8921965.
//
// Solidity: function withdrawConfirm((address,address,address,uint256,uint256,uint64,(uint256,uint256,uint8))[] withdrawConfirms) returns()
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
	Destination common.Address
	Token       common.Address
	Amount      *big.Int
	ChainId     *big.Int
	Nonce       uint64
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x7ba97cc0a03c9de2b8d972e5a42d8867d1c3483652b679f5b8b8b5b62a4ef576.
//
// Solidity: event Withdraw(bytes32 indexed message, address indexed user, address destination, address indexed token, uint256 amount, uint256 chainId, uint64 nonce)
func (_BridgeHub *BridgeHubFilterer) FilterWithdraw(opts *bind.FilterOpts, message [][32]byte, user []common.Address, token []common.Address) (*BridgeHubWithdrawIterator, error) {

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

// WatchWithdraw is a free log subscription operation binding the contract event 0x7ba97cc0a03c9de2b8d972e5a42d8867d1c3483652b679f5b8b8b5b62a4ef576.
//
// Solidity: event Withdraw(bytes32 indexed message, address indexed user, address destination, address indexed token, uint256 amount, uint256 chainId, uint64 nonce)
func (_BridgeHub *BridgeHubFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *BridgeHubWithdraw, message [][32]byte, user []common.Address, token []common.Address) (event.Subscription, error) {

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

// ParseWithdraw is a log parse operation binding the contract event 0x7ba97cc0a03c9de2b8d972e5a42d8867d1c3483652b679f5b8b8b5b62a4ef576.
//
// Solidity: event Withdraw(bytes32 indexed message, address indexed user, address destination, address indexed token, uint256 amount, uint256 chainId, uint64 nonce)
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
