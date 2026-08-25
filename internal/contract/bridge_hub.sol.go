// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

import (
	"errors"
	"math/big"
	"strings"

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
)

// CrossChainMessage is an auto generated low-level Go binding around an user-defined struct.
type CrossChainMessage struct {
	DomainSeparator [32]byte
	Signature       Signature
	MessageRawData  []byte
}

// DepositConfirm is an auto generated low-level Go binding around an user-defined struct.
type DepositConfirm struct {
	User        common.Address
	Destination common.Address
	Token       common.Address
	Amount      *big.Int
	ChainId     *big.Int
	BlockNumber uint64
	TxHash      [32]byte
	LogIndex    uint64
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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EnforcedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExpectedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"message\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"totalPower\",\"type\":\"uint64\"}],\"name\":\"BridgeSignatureSubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"message\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"destination\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"blockNumber\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"txHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"logIndex\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"message\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"errorCode\",\"type\":\"uint32\"}],\"name\":\"FailedDeposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"FeesClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"newEpoch\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"hotAddresses\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"coldAddresses\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint64[]\",\"name\":\"powers\",\"type\":\"uint64[]\"}],\"name\":\"FinalizedValidatorSetUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"message\",\"type\":\"bytes32\"}],\"name\":\"MessageStorageCleared\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"oldEpoch\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"hotAddresses\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"coldAddresses\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint64[]\",\"name\":\"powers\",\"type\":\"uint64[]\"}],\"name\":\"RemovedValidatorSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"newEpoch\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"hotAddresses\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"coldAddresses\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint64[]\",\"name\":\"powers\",\"type\":\"uint64[]\"}],\"name\":\"RequestedValidatorSetUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"bridgedToken\",\"type\":\"address\"}],\"name\":\"TokenPairSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"message\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"destination\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"message\",\"type\":\"bytes32\"}],\"name\":\"WithdrawCompleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"WithdrawFeeSet\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PAUSER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"claimFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"message\",\"type\":\"bytes32\"}],\"name\":\"clearMessageStorage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"coldValidatorList\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"destination\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"blockNumber\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"txHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"logIndex\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"r\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"s\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"}],\"internalType\":\"structSignature\",\"name\":\"signature\",\"type\":\"tuple\"}],\"internalType\":\"structDepositConfirm[]\",\"name\":\"deposits\",\"type\":\"tuple[]\"}],\"name\":\"depositConfirm\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"depositNonce\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"domainSeparator\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"epoch\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"message\",\"type\":\"bytes32\"}],\"name\":\"getBridgeMessageSignatures\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"totalPower\",\"type\":\"uint64\"},{\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"r\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"s\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"}],\"internalType\":\"structSignature[]\",\"name\":\"signatures\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"rawData\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"message\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"getBridgeValidatorSignature\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"r\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"s\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"}],\"internalType\":\"structSignature\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getColdValidators\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getHotValidators\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"message\",\"type\":\"bytes32\"}],\"name\":\"getMessageSignatures\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"totalPower\",\"type\":\"uint64\"},{\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"r\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"s\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"}],\"internalType\":\"structSignature[]\",\"name\":\"signatures\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPendingMessages\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPendingValidatorSetUpdate\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"_epoch\",\"type\":\"uint64\"},{\"internalType\":\"address[]\",\"name\":\"hotAddresses\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"coldAddresses\",\"type\":\"address[]\"},{\"internalType\":\"uint64[]\",\"name\":\"powers\",\"type\":\"uint64[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"message\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"getValidatorSignature\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"r\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"s\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"}],\"internalType\":\"structSignature\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getValidators\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"_epoch\",\"type\":\"uint64\"},{\"internalType\":\"address[]\",\"name\":\"validators\",\"type\":\"address[]\"},{\"internalType\":\"uint64[]\",\"name\":\"powers\",\"type\":\"uint64[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"hotValidatorList\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"hotAddresses\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"coldAddresses\",\"type\":\"address[]\"},{\"internalType\":\"uint64[]\",\"name\":\"powers\",\"type\":\"uint64[]\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"destination\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"blockNumber\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"txHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"logIndex\",\"type\":\"uint64\"}],\"name\":\"makeDepositMessage\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"newEpoch\",\"type\":\"uint64\"},{\"internalType\":\"address[]\",\"name\":\"hotAddresses\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"coldAddresses\",\"type\":\"address[]\"},{\"internalType\":\"uint64[]\",\"name\":\"powers\",\"type\":\"uint64[]\"}],\"name\":\"makeUpdateValidatorSetMessage\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"destination\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"name\":\"makeWithdrawMessage\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"pendingMessages\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"processedMessages\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"bridgedToken\",\"type\":\"address\"}],\"name\":\"setTokenPair\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"setWithdrawFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"domainSeparator\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"r\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"s\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"}],\"internalType\":\"structSignature\",\"name\":\"signature\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"messageRawData\",\"type\":\"bytes\"}],\"internalType\":\"structCrossChainMessage[]\",\"name\":\"items\",\"type\":\"tuple[]\"}],\"name\":\"submitBridgeSignatures\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"tokenPair\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"tokenWithdrawFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalValidatorPower\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"newEpoch\",\"type\":\"uint64\"},{\"internalType\":\"address[]\",\"name\":\"hotAddresses\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"coldAddresses\",\"type\":\"address[]\"},{\"internalType\":\"uint64[]\",\"name\":\"powers\",\"type\":\"uint64[]\"}],\"name\":\"updateValidatorSet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"epoch\",\"type\":\"uint64\"},{\"internalType\":\"address[]\",\"name\":\"hotAddresses\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"coldAddresses\",\"type\":\"address[]\"},{\"internalType\":\"uint64[]\",\"name\":\"powers\",\"type\":\"uint64[]\"}],\"internalType\":\"structValidatorSetUpdate\",\"name\":\"validatorSet\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"r\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"s\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"}],\"internalType\":\"structSignature\",\"name\":\"signature\",\"type\":\"tuple\"}],\"name\":\"updateValidatorSetConfirm\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"validatorPowers\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"destination\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"destination\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"deadline\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"r\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"s\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"}],\"internalType\":\"structSignature\",\"name\":\"signature\",\"type\":\"tuple\"}],\"internalType\":\"structWithdrawWithPermit[]\",\"name\":\"withdraws\",\"type\":\"tuple[]\"}],\"name\":\"withdrawBatchWithPermit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"destination\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"r\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"s\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"}],\"internalType\":\"structSignature\",\"name\":\"signature\",\"type\":\"tuple\"}],\"internalType\":\"structWithdrawConfirm[]\",\"name\":\"withdrawConfirms\",\"type\":\"tuple[]\"}],\"name\":\"withdrawConfirm\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawNonce\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60a0604052306080523480156200001557600080fd5b506200002062000026565b620000da565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000900460ff1615620000775760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160401b0390811614620000d75780546001600160401b0319166001600160401b0390811782556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50565b60805161580562000104600039600081816135b5015281816135de015261372f01526158056000f3fe6080604052600436106102e45760003560e01c80637bfe950c11610190578063c4366158116100dc578063de35f5cb11610095578063f61aa09d1161006f578063f61aa09d14610983578063f698da25146109b9578063f8156a6e146109cf578063fa734b13146109ef57600080fd5b8063de35f5cb1461091f578063e63ab1e91461093f578063ee0791b71461096157600080fd5b8063c43661581461085b578063c7bd5b2e14610870578063c89219651461089f578063ceca23f1146108bf578063d547741f146108df578063dcf70f3b146108ff57600080fd5b8063a217fddf11610149578063b73124b211610123578063b73124b2146107ce578063b7ab4db5146107f0578063b8a4e15114610814578063ba14c40d1461083b57600080fd5b8063a217fddf1461075b578063a8ff007514610770578063ad3cb1cc1461079057600080fd5b80637bfe950c146106775780638456cb591461069757806388ba16ab146106ac578063900cf0cf146106dc57806391d148541461071b5780639adc25d81461073b57600080fd5b80633ba9613b1161024f57806352d1902d116102085780635c975abb116101e25780635c975abb146105f057806364f8b391146106155780636e4bc0aa1461063557806375b238fc1461065557600080fd5b806352d1902d1461059b57806354e0fb21146105b05780635a53db10146105d057600080fd5b80633ba9613b146104e35780633f4ba83a14610503578063446c2c9a14610518578063456b07f914610548578063490b105c146105685780634f1ef2861461058857600080fd5b8063248a9ca3116102a1578063248a9ca3146103f85780632922e6e5146104185780632b1a7b58146104385780632b90c338146104765780632f2ff15d146104a357806336568abe146104c357600080fd5b806301ffc9a7146102e957806303d3f5ee1461031e57806316a8dfee1461035957806317c6365d1461037b5780631c0950201461039b5780631dd82536146103d3575b600080fd5b3480156102f557600080fd5b50610309610304366004614545565b610a0f565b60405190151581526020015b60405180910390f35b34801561032a57600080fd5b5061034b61033936600461458b565b600a6020526000908152604090205481565b604051908152602001610315565b34801561036557600080fd5b506103796103743660046145f2565b610a46565b005b34801561038757600080fd5b50610379610396366004614633565b610b1d565b3480156103a757600080fd5b506103bb6103b6366004614691565b610d30565b6040516001600160a01b039091168152602001610315565b3480156103df57600080fd5b506103e8610d5a565b6040516103159493929190614729565b34801561040457600080fd5b5061034b610413366004614691565b610eee565b34801561042457600080fd5b5061034b610433366004614691565b610f10565b34801561044457600080fd5b506103bb61045336600461477c565b60006020818152928152604080822090935290815220546001600160a01b031681565b34801561048257600080fd5b5061049661049136600461477c565b610f31565b60405161031591906147a8565b3480156104af57600080fd5b506103796104be36600461477c565b610fab565b3480156104cf57600080fd5b506103796104de36600461477c565b610fcd565b3480156104ef57600080fd5b506103796104fe366004614827565b611005565b34801561050f57600080fd5b5061037961111f565b34801561052457600080fd5b50610538610533366004614691565b611142565b6040516103159493929190614973565b34801561055457600080fd5b50610379610563366004614b29565b611217565b34801561057457600080fd5b50610379610583366004614bb0565b611465565b610379610596366004614c25565b611523565b3480156105a757600080fd5b5061034b61153e565b3480156105bc57600080fd5b506103796105cb366004614cca565b61155b565b3480156105dc57600080fd5b5061034b6105eb366004614cf4565b6115f3565b3480156105fc57600080fd5b506000805160206157708339815191525460ff16610309565b34801561062157600080fd5b5061034b610630366004614d76565b611690565b34801561064157600080fd5b5061049661065036600461477c565b611784565b34801561066157600080fd5b5061034b6000805160206157b083398151915281565b34801561068357600080fd5b50610379610692366004614e0e565b6117fe565b3480156106a357600080fd5b50610379611832565b3480156106b857600080fd5b506103096106c7366004614691565b600d6020526000908152604090205460ff1681565b3480156106e857600080fd5b5060045461070390600160401b90046001600160401b031681565b6040516001600160401b039091168152602001610315565b34801561072757600080fd5b5061030961073636600461477c565b611852565b34801561074757600080fd5b5061034b610756366004614e50565b61188a565b34801561076757600080fd5b5061034b600081565b34801561077c57600080fd5b5061037961078b366004614691565b611912565b34801561079c57600080fd5b506107c1604051806040016040528060058152602001640352e302e360dc1b81525081565b6040516103159190614eb6565b3480156107da57600080fd5b506107e3611933565b6040516103159190614ec9565b3480156107fc57600080fd5b50610805611995565b60405161031593929190614edc565b34801561082057600080fd5b5060095461070390600160401b90046001600160401b031681565b34801561084757600080fd5b506103bb610856366004614691565b611aea565b34801561086757600080fd5b506107e3611afa565b34801561087c57600080fd5b5061089061088b366004614691565b611b5a565b60405161031593929190614f1a565b3480156108ab57600080fd5b506103796108ba3660046145f2565b611b8d565b3480156108cb57600080fd5b506103796108da366004614f4e565b611c3c565b3480156108eb57600080fd5b506103796108fa36600461477c565b611d5a565b34801561090b57600080fd5b5061037961091a366004614f8a565b611d76565b34801561092b57600080fd5b50600954610703906001600160401b031681565b34801561094b57600080fd5b5061034b60008051602061573083398151915281565b34801561096d57600080fd5b50610976611e9d565b6040516103159190614fc6565b34801561098f57600080fd5b5061070361099e36600461458b565b6003602052600090815260409020546001600160401b031681565b3480156109c557600080fd5b5061034b600f5481565b3480156109db57600080fd5b50600454610703906001600160401b031681565b3480156109fb57600080fd5b50610379610a0a36600461500a565b611ef4565b60006001600160e01b03198216637965db0b60e01b1480610a4057506301ffc9a760e01b6001600160e01b03198316145b92915050565b610a4e611fec565b610a5661201f565b80610aa05760405162461bcd60e51b8152602060048201526015602482015274456d7074792077697468647261777320617272617960581b60448201526064015b60405180910390fd5b8060005b816001600160401b0316816001600160401b03161015610b0057610af88484836001600160401b0316818110610adc57610adc61503f565b90506101200201803603810190610af3919061516f565b612057565b600101610aa4565b5050610b19600160008051602061579083398151915255565b5050565b610b25611fec565b610b2d61201f565b610b363361213d565b6000610c0a610b48602085018561518c565b610b5560208601866151a7565b80806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250610b949250505060408701876151a7565b80806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250610bd39250505060608801886151a7565b8080602002602001604051908101604052809392919081815260200183836020028082843760009201919091525061169092505050565b90506000610c17826121ab565b90506000610c3383610c2e368790038701876151f0565b6121fd565b90508015610d1657610c44826123f2565b610d16610c54602087018761518c565b610c6160208801886151a7565b80806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250610ca09250505060408901896151a7565b80806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250610cdf9250505060608a018a6151a7565b808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152506124dc92505050565b505050610b19600160008051602061579083398151915255565b60018181548110610d4057600080fd5b6000918252602090912001546001600160a01b0316905081565b60408051608081018252600580546001600160401b03168252600680548451602082810282018101909652818152600095606095869586958995929491938086019390830182828015610dd657602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311610db8575b5050505050815260200160028201805480602002602001604051908101604052809291908181526020018280548015610e3857602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311610e1a575b5050505050815260200160038201805480602002602001604051908101604052809291908181526020018280548015610ec257602002820191906000526020600020906000905b82829054906101000a90046001600160401b03166001600160401b031681526020019060080190602082600701049283019260010382029150808411610e7f5790505b505050919092525050815160208301516040840151606090940151919990985092965094509092505050565b6000908152600080516020615750833981519152602052604090206001015490565b600e8181548110610f2057600080fd5b600091825260209091200154905081565b610f5860405180606001604052806000815260200160008152602001600060ff1681525090565b506000828152600c602090815260408083206001600160a01b03851684526001908101835292819020815160608101835281548152938101549284019290925260029091015460ff169082015292915050565b610fb482610eee565b610fbd816128b1565b610fc783836128bb565b50505050565b6001600160a01b0381163314610ff65760405163334bd91960e11b815260040160405180910390fd5b6110008282612967565b505050565b6000805160206157b083398151915261101d816128b1565b6000886001600160401b0316116110765760405162461bcd60e51b815260206004820152601a60248201527f4e65772065706f6368206d75737420626520706f7369746976650000000000006044820152606401610a97565b6111158888888080602002602001604051908101604052809392919081815260200183836020028082843760009201919091525050604080516020808c0282810182019093528b82529093508b92508a91829185019084908082843760009201919091525050604080516020808b0282810182019093528a82529093508a9250899182918501908490808284376000920191909152506129e392505050565b5050505050505050565b600080516020615730833981519152611137816128b1565b61113f61301b565b50565b6000818152600b60205260408120606090819081906111608161307b565b82546002840180549397509195506001600160401b0316918691869181906111879061520c565b80601f01602080910402602001604051908101604052809291908181526020018280546111b39061520c565b80156112005780601f106111d557610100808354040283529160200191611200565b820191906000526020600020905b8154815290600101906020018083116111e357829003601f168201915b505050505090509450945094509450509193509193565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a008054600160401b810460ff1615906001600160401b031660008115801561125c5750825b90506000826001600160401b031660011480156112785750303b155b905081158015611286575080155b156112a45760405163f92ee8a960e01b815260040160405180910390fd5b845467ffffffffffffffff1916600117855583156112ce57845460ff60401b1916600160401b1785555b6112d661322f565b6112de61322f565b6112e6613237565b6112ee613247565b6112f96000336128bb565b506113126000805160206157b0833981519152336128bb565b5061132b600080516020615730833981519152336128bb565b5061133960008989896129e3565b61134660008989896124dc565b6040805180820182526009815268213934b233b2a43ab160b91b6020918201528151808301835260018152603160f81b9082015281517f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f818301527f4c4f9c18a13e136e05f5178f806d2f407f435220d27e3fd1cf60052217fef7e4818401527fc89efdaa54c0f20c7adf612882df0950f5a951637e0307cdcb4c672f298b8bc660608201524660808201523060a0808301919091528351808303909101815260c09091019092528151910120600f55831561111557845460ff60401b19168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15050505050505050565b61146d611fec565b61147561201f565b61147e3361213d565b806114c25760405162461bcd60e51b8152602060048201526014602482015273456d707479206465706f7369747320617272617960601b6044820152606401610a97565b60005b8181101561150b57368383838181106114e0576114e061503f565b905061016002019050611502818036038101906114fd9190615246565b613257565b506001016114c5565b50610b19600160008051602061579083398151915255565b61152b6135aa565b6115348261364f565b610b198282613667565b6000611548613724565b5060008051602061571083398151915290565b6000805160206157b0833981519152611573816128b1565b6001600160a01b0383166115995760405162461bcd60e51b8152600401610a97906152e8565b6001600160a01b0383166000818152600a602052604090819020849055517f64f4bc88c432868175af31fe623ff706067afe070c40f869f28e3daed63f8243906115e69085815260200190565b60405180910390a2505050565b604080517f0418a33254eb82955db214bbef77af34b0921778a526c3613eb9e56319477c1b6020808301919091526001600160a01b039a8b1682840152988a16606082015296909816608087015260a086019490945260c08501929092526001600160401b0390811660e0850152610100840191909152166101208083019190915283518083039091018152610140909101909252815191012090565b60007fcd26826da4f5c0e82ef8057ecacd8931dfb36167a70c820505f10826298cd05e85856040516020016116c59190615317565b60405160208183030381529060405280519060200120856040516020016116ec9190615317565b60405160208183030381529060405280519060200120856040516020016117139190615356565b604051602081830303815290604052805190602001206040516020016117649594939291909485526001600160401b0393909316602085015260408401919091526060830152608082015260a00190565b604051602081830303815290604052805190602001209050949350505050565b6117ab60405180606001604052806000815260200160008152602001600060ff1681525090565b506000828152600b602090815260408083206001600160a01b03851684526001908101835292819020815160608101835281548152938101549284019290925260029091015460ff169082015292915050565b611806611fec565b61180e61201f565b61181b338585858561376d565b610fc7600160008051602061579083398151915255565b60008051602061573083398151915261184a816128b1565b61113f613a5e565b6000918252600080516020615750833981519152602090815260408084206001600160a01b0393909316845291905290205460ff1690565b604080517fed44fa5a448edcc9a97caee522159268a5089c8700fe54f678309d7a73c9f6ec6020808301919091526001600160a01b0398891682840152968816606082015294909616608085015260a084019290925260c08301526001600160401b031660e08083019190915283518083039091018152610100909101909252815191012090565b6000805160206157b083398151915261192a816128b1565b610b1982613aa7565b6060600280548060200260200160405190810160405280929190818152602001828054801561198b57602002820191906000526020600020905b81546001600160a01b0316815260019091019060200180831161196d575b5050505050905090565b6000606080600460089054906101000a90046001600160401b031692506002805480602002602001604051908101604052809291908181526020018280548015611a0857602002820191906000526020600020905b81546001600160a01b031681526001909101906020018083116119ea575b5050505050915081516001600160401b03811115611a2857611a286149bb565b604051908082528060200260200182016040528015611a51578160200160208202803683370190505b50905060005b8251811015611ae45760036000848381518110611a7657611a7661503f565b60200260200101516001600160a01b03166001600160a01b0316815260200190815260200160002060009054906101000a90046001600160401b0316828281518110611ac457611ac461503f565b6001600160401b0390921660209283029190910190910152600101611a57565b50909192565b60028181548110610d4057600080fd5b6060600180548060200260200160405190810160405280929190818152602001828054801561198b576020028201919060005260206000209081546001600160a01b0316815260019091019060200180831161196d575050505050905090565b6000818152600c602052604081206060908190611b768161307b565b91546001600160401b031696909550909350915050565b611b95611fec565b611b9d61201f565b611ba63361213d565b80611bf35760405162461bcd60e51b815260206004820152601d60248201527f456d70747920776974686472617720636f6e6669726d732061727261790000006044820152606401610a97565b60005b8181101561150b5736838383818110611c1157611c1161503f565b905061012002019050611c3381803603810190611c2e919061516f565b613b2f565b50600101611bf6565b6000805160206157b0833981519152611c54816128b1565b6001600160a01b038316611caa5760405162461bcd60e51b815260206004820152601960248201527f496e76616c696420726563697069656e742061646472657373000000000000006044820152606401610a97565b60008211611cfa5760405162461bcd60e51b815260206004820152601760248201527f416d6f756e74206d75737420626520706f7369746976650000000000000000006044820152606401610a97565b611d0e6001600160a01b0385168484613bc7565b604080516001600160a01b038581168252602082018590528616917ffe3464cd748424446c37877c28ce5b700222c5bc9f90d908afcc4e5cb22707ff910160405180910390a250505050565b611d6382610eee565b611d6c816128b1565b610fc78383612967565b6000805160206157b0833981519152611d8e816128b1565b83600003611dae5760405162461bcd60e51b8152600401610a9790615389565b6001600160a01b038316611dd45760405162461bcd60e51b8152600401610a97906152e8565b6001600160a01b038216611e225760405162461bcd60e51b815260206004820152601560248201527424b73b30b634b210313934b233b2b2103a37b5b2b760591b6044820152606401610a97565b6000848152602081815260408083206001600160a01b038781168086529190935281842080546001600160a01b031990811694881694851790915583855282852080549091168217905590519192909187917f5478d55b1b0642b4b219976302fbc1f304d09602ef98dcacbeb94121af8a2a5491a450505050565b6060600e80548060200260200160405190810160405280929190818152602001828054801561198b57602002820191906000526020600020905b815481526020019060010190808311611ed7575050505050905090565b611efc611fec565b611f0461201f565b611f0d3361213d565b80611f535760405162461bcd60e51b8152602060048201526016602482015275456d707479207369676e61747572657320617272617960501b6044820152606401610a97565b60005b8181101561150b57611fe4838383818110611f7357611f7361503f565b9050602002810190611f8591906153b2565b602001848484818110611f9a57611f9a61503f565b9050602002810190611fac91906153b2565b35858585818110611fbf57611fbf61503f565b9050602002810190611fd191906153b2565b611fdf9060808101906153d2565b613c26565b600101611f56565b6000805160206157708339815191525460ff161561201d5760405163d93c066560e01b815260040160405180910390fd5b565b60008051602061579083398151915280546001190161205157604051633ee5aeb560e01b815260040160405180910390fd5b60029055565b6040818101518251606084015160a085015160c0860151808601518151602090920151965163d505accf60e01b81526001600160a01b03958616600482015230602482015260448101949094526001600160401b03909216606484015260ff909116608483015260a482015260c4810193909352169063d505accf9060e401600060405180830381600087803b1580156120f057600080fd5b505af1158015612104573d6000803e3d6000fd5b5050505061113f8160000151826020015183604001518460600151856080015161376d565b600160008051602061579083398151915255565b6001600160a01b0381166000908152600360205260409020546001600160401b031661113f5760405162461bcd60e51b815260206004820152601960248201527f5369676e6572206973206e6f7420612076616c696461746f72000000000000006044820152606401610a97565b6000806121b783613e1f565b905060008111610a405760405162461bcd60e51b815260206004820152601160248201527013595cdcd859d9481b9bdd08199bdd5b99607a1b6044820152606401610a97565b6000828152600d602052604081205460ff16156122505760405162461bcd60e51b8152602060048201526011602482015270105b1c9958591e481c1c9bd8d95cdcd959607a1b6044820152606401610a97565b600061225f8484600f54613e78565b905061226a8161213d565b6000848152600c602090815260408083206001600160a01b0385168452600101909152902054156122eb5760405162461bcd60e51b815260206004820152602560248201527f56616c696461746f7220616c7265616479207369676e65642074686973206d65604482015264737361676560d81b6064820152608401610a97565b6000848152600c602081815260408084206001600160a01b038616855260018082018452828620895181558985015191810191909155888301516002909101805460ff191660ff90921691909117905560038352908420548885529290915254612361916001600160401b03908116911661542e565b6000868152600c60205260409020805467ffffffffffffffff19166001600160401b03838116919091179091556004549192506123a09116600261544e565b6001600160401b03166123b482600361544e565b6001600160401b031611156123e7575050506000828152600d60205260409020805460ff19166001908117909155610a40565b506000949350505050565b6000811180156124045750600e548111155b6124405760405162461bcd60e51b815260206004820152600d60248201526c092dcecc2d8d2c840d2dcc8caf609b1b6044820152606401610a97565b600061244d600183615479565b600e5490915061245f90600190615479565b8110156124b157600e805461247690600190615479565b815481106124865761248661503f565b9060005260206000200154600e82815481106124a4576124a461503f565b6000918252602090912001555b600e8054806124c2576124c261548c565b600190038181906000526020600020016000905590555050565b6002546000906001600160401b038111156124f9576124f96149bb565b604051908082528060200260200182016040528015612522578160200160208202803683370190505b50905060005b6002548110156125fe576000600282815481106125475761254761503f565b60009182526020808320909101546001600160a01b031680835260039091526040909120549091506001600160401b0316156125f5576001600160a01b03811660009081526003602052604090205483516001600160401b03909116908490849081106125b6576125b661503f565b6001600160401b039092166020928302919091018201526001600160a01b0382166000908152600390915260409020805467ffffffffffffffff191690555b50600101612528565b50600460089054906101000a90046001600160401b03166001600160401b03167f7078dd7d8f69886fa49f17529fbd86d6b24846959cd85dcfb4e990205537dc5d6002600184604051612653939291906154e1565b60405180910390a2825161266e9060019060208601906143e2565b5083516126829060029060208701906143e2565b506000805b85518110156127195760008682815181106126a4576126a461503f565b6020026020010151905060008583815181106126c2576126c261503f565b6020908102919091018101516001600160a01b038416600090815260039092526040909120805467ffffffffffffffff19166001600160401b038316179055905061270d818561542e565b93505050600101612687565b50600480546001600160401b03888116600160401b026fffffffffffffffffffffffffffffffff1990921690841617179055604080516080810190915260008082526020820190604051908082528060200260200182016040528015612789578160200160208202803683370190505b50815260200160006040519080825280602002602001820160405280156127ba578160200160208202803683370190505b50815260200160006040519080825280602002602001820160405280156127eb578160200160208202803683370190505b50905280516005805467ffffffffffffffff19166001600160401b0390921691909117815560208083015180516128269260069201906143e2565b50604082015180516128429160028401916020909101906143e2565b506060820151805161285e916003840191602090910190614447565b50905050856001600160401b03167f7fb1406cb8c05384e1f39b879d591d98e7be9ffa67cf31cdbaa10b3aebf7157c60026001866040516128a1939291906154e1565b60405180910390a2505050505050565b61113f8133614041565b60006000805160206157508339815191526128d68484611852565b612956576000848152602082815260408083206001600160a01b03871684529091529020805460ff1916600117905561290c3390565b6001600160a01b0316836001600160a01b0316857f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a46001915050610a40565b6000915050610a40565b5092915050565b60006000805160206157508339815191526129828484611852565b15612956576000848152602082815260408083206001600160a01b0387168085529252808320805460ff1916905551339287917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a46001915050610a40565b6000835111612a345760405162461bcd60e51b815260206004820152601960248201527f456d70747920686f7420616464726573736573206172726179000000000000006044820152606401610a97565b8151835114612a945760405162461bcd60e51b815260206004820152602660248201527f486f7420616e6420636f6c6420616464726573736573206c656e677468206d696044820152650e6dac2e8c6d60d31b6064820152608401610a97565b8051835114612af65760405162461bcd60e51b815260206004820152602860248201527f486f742061646472657373657320616e6420706f77657273206c656e677468206044820152670dad2e6dac2e8c6d60c31b6064820152608401610a97565b600454600160401b90046001600160401b0316151580612b1e57506001600160401b03841615155b15612b9d576004546001600160401b03600160401b909104811690851611612b9d5760405162461bcd60e51b815260206004820152602c60248201527f4e65772065706f6368206d7573742062652067726561746572207468616e206360448201526b0eae4e4cadce840cae0dec6d60a31b6064820152608401610a97565b6000805b8451811015612e9d576000858281518110612bbe57612bbe61503f565b602002602001015190506000858381518110612bdc57612bdc61503f565b602002602001015190506000858481518110612bfa57612bfa61503f565b6020026020010151905060006001600160a01b0316836001600160a01b031603612c665760405162461bcd60e51b815260206004820152601d60248201527f5a65726f206164647265737320696e20686f74206164647265737365730000006044820152606401610a97565b6001600160a01b038216612cbc5760405162461bcd60e51b815260206004820152601e60248201527f5a65726f206164647265737320696e20636f6c642061646472657373657300006044820152606401610a97565b6000816001600160401b031611612d0e5760405162461bcd60e51b8152602060048201526016602482015275506f776572206d75737420626520706f73697469766560501b6044820152606401610a97565b816001600160a01b0316836001600160a01b031603612d7a5760405162461bcd60e51b815260206004820152602260248201527f486f7420616e6420636f6c6420616464726573736573206d757374206469666660448201526132b960f11b6064820152608401610a97565b60005b84811015612e8157888181518110612d9757612d9761503f565b60200260200101516001600160a01b0316846001600160a01b031603612dff5760405162461bcd60e51b815260206004820152601760248201527f4475706c696361746520686f74206164647265737365730000000000000000006044820152606401610a97565b878181518110612e1157612e1161503f565b60200260200101516001600160a01b0316836001600160a01b031603612e795760405162461bcd60e51b815260206004820152601860248201527f4475706c696361746520636f6c642061646472657373657300000000000000006044820152606401610a97565b600101612d7d565b50612e8c818661542e565b94505060019092019150612ba19050565b506000816001600160401b031611612ef75760405162461bcd60e51b815260206004820152601c60248201527f546f74616c20706f776572206d75737420626520706f736974697665000000006044820152606401610a97565b6000612f0586868686611690565b90506001600160401b03861615612f4c57600e80546001810182556000919091527fbb7b4a454dc3493923482f07822329ed19e8244eff582cc204f8554c3620c3fd018190555b604080516080810182526001600160401b0388168082526020808301899052928201879052606082018690526005805467ffffffffffffffff19169091178155875191929091612fa291600691908a01906143e2565b5060408201518051612fbe9160028401916020909101906143e2565b5060608201518051612fda916003840191602090910190614447565b50905050856001600160401b03167ff389db8d301520921c1c0eeb8eaf6790791f8916b65e8e93424a7ec831edd69e8686866040516128a19392919061551a565b61302361407a565b600080516020615770833981519152805460ff191681557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b6040516001600160a01b03909116815260200160405180910390a150565b60025460609081906000816001600160401b0381111561309d5761309d6149bb565b6040519080825280602002602001820160405280156130c6578160200160208202803683370190505b5090506000826001600160401b038111156130e3576130e36149bb565b60405190808252806020026020018201604052801561313b57816020015b61312860405180606001604052806000815260200160008152602001600060ff1681525090565b8152602001906001900390816131015790505b5090506000805b8481101561321e5760006002828154811061315f5761315f61503f565b60009182526020808320909101546001600160a01b031680835260018c810183526040938490208451606081018652815480825292820154948101949094526002015460ff16938301939093529250901561321457818685815181106131c7576131c761503f565b60200260200101906001600160a01b031690816001600160a01b031681525050808585815181106131fa576131fa61503f565b602002602001018190525083806132109061553f565b9450505b5050600101613142565b508083528152909590945092505050565b61201d6140aa565b61323f6140aa565b61201d6140f3565b61324f6140aa565b61201d614114565b80516001600160a01b03166132a55760405162461bcd60e51b8152602060048201526014602482015273496e76616c69642075736572206164647265737360601b6044820152606401610a97565b60208101516001600160a01b03166132ff5760405162461bcd60e51b815260206004820152601b60248201527f496e76616c69642064657374696e6174696f6e206164647265737300000000006044820152606401610a97565b60408101516001600160a01b03166133295760405162461bcd60e51b8152600401610a97906152e8565b60008160600151116133715760405162461bcd60e51b81526020600482015260116024820152700416d6f756e74206d757374206265203e3607c1b6044820152606401610a97565b80608001516000036133955760405162461bcd60e51b8152600401610a9790615389565b6080810151600090815260208181526040808320818501516001600160a01b039081168552925290912054166133ca8161411c565b60006133fc836000015184602001518560400151866060015187608001518860a001518960c001518a60e001516115f3565b9050600061340f828561010001516121fd565b90508015610fc757602084015160608501516040516340c10f1960e01b81526001600160a01b038616926340c10f199261345f926004016001600160a01b03929092168252602082015260400190565b600060405180830381600087803b15801561347957600080fd5b505af115801561348d573d6000803e3d6000fd5b505060098054600193509091506000906134b19084906001600160401b031661542e565b92506101000a8154816001600160401b0302191690836001600160401b0316021790555083604001516001600160a01b031684600001516001600160a01b0316837f63b86360efb9dda5b052fea4b2df3ca28fb97895a7714ee17dc371a5a472c6858760200151886060015189608001518a60a001518b60c001518c60e00151600960009054906101000a90046001600160401b031660405161359c97969594939291906001600160a01b03979097168752602087019590955260408601939093526001600160401b039182166060860152608085015290811660a08401521660c082015260e00190565b60405180910390a450505050565b306001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016148061363157507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316613625600080516020615710833981519152546001600160a01b031690565b6001600160a01b031614155b1561201d5760405163703e46dd60e11b815260040160405180910390fd5b6000805160206157b0833981519152610b19816128b1565b816001600160a01b03166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa9250505080156136c1575060408051601f3d908101601f191682019092526136be91810190615558565b60015b6136e957604051634c9c8ce360e01b81526001600160a01b0383166004820152602401610a97565b600080516020615710833981519152811461371a57604051632a87526960e21b815260048101829052602401610a97565b6110008383614164565b306001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000161461201d5760405163703e46dd60e11b815260040160405180910390fd5b6001600160a01b0383166000908152600a60205260409020548083116137ce5760405162461bcd60e51b8152602060048201526016602482015275416d6f756e74206d757374206578636565642066656560501b6044820152606401610a97565b6001600160a01b03861661381b5760405162461bcd60e51b8152602060048201526014602482015273496e76616c69642075736572206164647265737360601b6044820152606401610a97565b6001600160a01b0385166138715760405162461bcd60e51b815260206004820152601b60248201527f496e76616c69642064657374696e6174696f6e206164647265737300000000006044820152606401610a97565b6001600160a01b0384166138975760405162461bcd60e51b8152600401610a97906152e8565b816000036138b75760405162461bcd60e51b8152600401610a9790615389565b6000828152602081815260408083206001600160a01b038089168552925290912054166138e38161411c565b6138f86001600160a01b0386168830876141ba565b6139028285615479565b604051630852cd8d60e31b8152600481018290529094506001600160a01b038616906342966c6890602401600060405180830381600087803b15801561394757600080fd5b505af115801561395b573d6000803e3d6000fd5b5050505060006009600881819054906101000a90046001600160401b031661398290615571565b91906101000a8154816001600160401b0302191690836001600160401b031602179055905060006139b789898589898761188a565b600e80546001810182556000919091527fbb7b4a454dc3493923482f07822329ed19e8244eff582cc204f8554c3620c3fd01819055604080516001600160a01b038b81168252602082018a90529181018890526001600160401b038516606082015291925084811691908b169083907f7ba97cc0a03c9de2b8d972e5a42d8867d1c3483652b679f5b8b8b5b62a4ef5769060800160405180910390a4505050505050505050565b613a66611fec565b600080516020615770833981519152805460ff191660011781557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a2583361305d565b6000818152600c60205260408120805467ffffffffffffffff1916815590613ad260028301826144fa565b50506000818152600b60205260408120805467ffffffffffffffff1916815590613aff60028301826144fa565b505060405181907fa5e9cb0cafdc12b93fa2ff5814a2a2f6ef3e98f1097c8073778b085111058f0090600090a250565b6000613b57826000015183602001518460400151856060015186608001518760a0015161188a565b90506000613b64826121ab565b90506000613b76838560c001516121fd565b90508015610fc757613b87826123f2565b60a084015160405184916001600160401b0316907fda5f7a28c1ede043b7194dba8741d889157152b30393f445fac51654abe02fae90600090a350505050565b6040516001600160a01b0383811660248301526044820183905261100091859182169063a9059cbb906064015b604051602081830303815290604052915060e01b6020820180516001600160e01b0383818316178352505050506141f3565b600f548303613c685760405162461bcd60e51b815260206004820152600e60248201526d24b73b30b634b2103237b6b0b4b760911b6044820152606401610a97565b60008282604051613c7a929190615597565b60405180910390209050613c8d816121ab565b506000613ca982613ca3368990038901896151f0565b87613e78565b9050613cb48161213d565b6000828152600b602090815260408083206001600160a01b038516845260010190915290205415613d185760405162461bcd60e51b815260206004820152600e60248201526d105b1c9958591e481cda59db995960921b6044820152606401610a97565b6000828152600b602090815260408083206001600160a01b038516845260010190915290208690613d4982826155a7565b50506000828152600b60205260409020600201613d6784868361562a565b506001600160a01b038116600090815260036020908152604080832054858452600b909252822054613da5916001600160401b03908116911661542e565b6000848152600b6020908152604091829020805467ffffffffffffffff19166001600160401b03851690811790915582516001600160a01b03871681529182015291925084917fc36d3c025c0f565ea8451b940f66ac343971b6b6d515ea860aed580c6408af2b910160405180910390a250505050505050565b600e54600090815b81811015613e6e5783600e8281548110613e4357613e4361503f565b906000526020600020015403613e6657613e5e8160016156ea565b949350505050565b600101613e27565b5060009392505050565b81516000908103613ecb5760405162461bcd60e51b815260206004820152601b60248201527f496e76616c6964207369676e6174757265202772272076616c756500000000006044820152606401610a97565b8260200151600003613f1f5760405162461bcd60e51b815260206004820152601b60248201527f496e76616c6964207369676e6174757265202773272076616c756500000000006044820152606401610a97565b60405161190160f01b6020820152602281018390526042810185905260009060620160408051601f1981840301815282825280516020918201208783015188518984015160008088529690940194859052919550600193613f9b9387939193845260ff9290921660208401526040830152606082015260800190565b6020604051602081039080840390855afa158015613fbd573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b0381166140365760405162461bcd60e51b815260206004820152602d60248201527f496e76616c6964207369676e61747572652c207265636f76657265642074686560448201526c207a65726f206164647265737360981b6064820152608401610a97565b9150505b9392505050565b61404b8282611852565b610b195760405163e2517d3f60e01b81526001600160a01b038216600482015260248101839052604401610a97565b6000805160206157708339815191525460ff1661201d57604051638dfc202b60e01b815260040160405180910390fd5b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0054600160401b900460ff1661201d57604051631afcd79f60e31b815260040160405180910390fd5b6140fb6140aa565b600080516020615770833981519152805460ff19169055565b6121296140aa565b6001600160a01b03811661113f5760405162461bcd60e51b815260206004820152600f60248201526e151bdad95b881b9bdd08199bdd5b99608a1b6044820152606401610a97565b61416d82614264565b6040516001600160a01b038316907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b90600090a28051156141b25761100082826142c9565b610b1961433f565b6040516001600160a01b038481166024830152838116604483015260648201839052610fc79186918216906323b872dd90608401613bf4565b600080602060008451602086016000885af180614216576040513d6000823e3d81fd5b50506000513d9150811561422e57806001141561423b565b6001600160a01b0384163b155b15610fc757604051635274afe760e01b81526001600160a01b0385166004820152602401610a97565b806001600160a01b03163b60000361429a57604051634c9c8ce360e01b81526001600160a01b0382166004820152602401610a97565b60008051602061571083398151915280546001600160a01b0319166001600160a01b0392909216919091179055565b6060600080846001600160a01b0316846040516142e691906156fd565b600060405180830381855af49150503d8060008114614321576040519150601f19603f3d011682016040523d82523d6000602084013e614326565b606091505b509150915061433685838361435e565b95945050505050565b341561201d5760405163b398979f60e01b815260040160405180910390fd5b6060826143735761436e826143ba565b61403a565b815115801561438a57506001600160a01b0384163b155b156143b357604051639996b31560e01b81526001600160a01b0385166004820152602401610a97565b508061403a565b8051156143c957805160208201fd5b60405163d6bda27560e01b815260040160405180910390fd5b828054828255906000526020600020908101928215614437579160200282015b8281111561443757825182546001600160a01b0319166001600160a01b03909116178255602090920191600190910190614402565b50614443929150614530565b5090565b828054828255906000526020600020906003016004900481019282156144375791602002820160005b838211156144ba57835183826101000a8154816001600160401b0302191690836001600160401b031602179055509260200192600801602081600701049283019260010302614470565b80156144ed5782816101000a8154906001600160401b0302191690556008016020816007010492830192600103026144ba565b5050614443929150614530565b5080546145069061520c565b6000825580601f10614516575050565b601f01602090049060005260206000209081019061113f91905b5b808211156144435760008155600101614531565b60006020828403121561455757600080fd5b81356001600160e01b03198116811461403a57600080fd5b80356001600160a01b038116811461458657600080fd5b919050565b60006020828403121561459d57600080fd5b61403a8261456f565b60008083601f8401126145b857600080fd5b5081356001600160401b038111156145cf57600080fd5b602083019150836020610120830285010111156145eb57600080fd5b9250929050565b6000806020838503121561460557600080fd5b82356001600160401b0381111561461b57600080fd5b614627858286016145a6565b90969095509350505050565b600080828403608081121561464757600080fd5b83356001600160401b0381111561465d57600080fd5b84016080818703121561466f57600080fd5b92506060601f198201121561468357600080fd5b506020830190509250929050565b6000602082840312156146a357600080fd5b5035919050565b60008151808452602080850194506020840160005b838110156146e45781516001600160a01b0316875295820195908201906001016146bf565b509495945050505050565b60008151808452602080850194506020840160005b838110156146e45781516001600160401b031687529582019590820190600101614704565b6001600160401b038516815260806020820152600061474b60808301866146aa565b828103604084015261475d81866146aa565b9050828103606084015261477181856146ef565b979650505050505050565b6000806040838503121561478f57600080fd5b8235915061479f6020840161456f565b90509250929050565b815181526020808301519082015260408083015160ff169082015260608101610a40565b80356001600160401b038116811461458657600080fd5b60008083601f8401126147f557600080fd5b5081356001600160401b0381111561480c57600080fd5b6020830191508360208260051b85010111156145eb57600080fd5b60008060008060008060006080888a03121561484257600080fd5b61484b886147cc565b965060208801356001600160401b038082111561486757600080fd5b6148738b838c016147e3565b909850965060408a013591508082111561488c57600080fd5b6148988b838c016147e3565b909650945060608a01359150808211156148b157600080fd5b506148be8a828b016147e3565b989b979a50959850939692959293505050565b60008151808452602080850194506020840160005b838110156146e457614910878351805182526020808201519083015260409081015160ff16910152565b60609690960195908201906001016148e6565b60005b8381101561493e578181015183820152602001614926565b50506000910152565b6000815180845261495f816020860160208601614923565b601f01601f19169290920160200192915050565b6001600160401b038516815260806020820152600061499560808301866146aa565b82810360408401526149a781866148d1565b905082810360608401526147718185614947565b634e487b7160e01b600052604160045260246000fd5b60405161012081016001600160401b03811182821017156149f4576149f46149bb565b60405290565b604051601f8201601f191681016001600160401b0381118282101715614a2257614a226149bb565b604052919050565b60006001600160401b03821115614a4357614a436149bb565b5060051b60200190565b600082601f830112614a5e57600080fd5b81356020614a73614a6e83614a2a565b6149fa565b8083825260208201915060208460051b870101935086841115614a9557600080fd5b602086015b84811015614ab857614aab8161456f565b8352918301918301614a9a565b509695505050505050565b600082601f830112614ad457600080fd5b81356020614ae4614a6e83614a2a565b8083825260208201915060208460051b870101935086841115614b0657600080fd5b602086015b84811015614ab857614b1c816147cc565b8352918301918301614b0b565b600080600060608486031215614b3e57600080fd5b83356001600160401b0380821115614b5557600080fd5b614b6187838801614a4d565b94506020860135915080821115614b7757600080fd5b614b8387838801614a4d565b93506040860135915080821115614b9957600080fd5b50614ba686828701614ac3565b9150509250925092565b60008060208385031215614bc357600080fd5b82356001600160401b0380821115614bda57600080fd5b818501915085601f830112614bee57600080fd5b813581811115614bfd57600080fd5b86602061016083028501011115614c1357600080fd5b60209290920196919550909350505050565b60008060408385031215614c3857600080fd5b614c418361456f565b91506020808401356001600160401b0380821115614c5e57600080fd5b818601915086601f830112614c7257600080fd5b813581811115614c8457614c846149bb565b614c96601f8201601f191685016149fa565b91508082528784828501011115614cac57600080fd5b80848401858401376000848284010152508093505050509250929050565b60008060408385031215614cdd57600080fd5b614ce68361456f565b946020939093013593505050565b600080600080600080600080610100898b031215614d1157600080fd5b614d1a8961456f565b9750614d2860208a0161456f565b9650614d3660408a0161456f565b95506060890135945060808901359350614d5260a08a016147cc565b925060c08901359150614d6760e08a016147cc565b90509295985092959890939650565b60008060008060808587031215614d8c57600080fd5b614d95856147cc565b935060208501356001600160401b0380821115614db157600080fd5b614dbd88838901614a4d565b94506040870135915080821115614dd357600080fd5b614ddf88838901614a4d565b93506060870135915080821115614df557600080fd5b50614e0287828801614ac3565b91505092959194509250565b60008060008060808587031215614e2457600080fd5b614e2d8561456f565b9350614e3b6020860161456f565b93969395505050506040820135916060013590565b60008060008060008060c08789031215614e6957600080fd5b614e728761456f565b9550614e806020880161456f565b9450614e8e6040880161456f565b93506060870135925060808701359150614eaa60a088016147cc565b90509295509295509295565b60208152600061403a6020830184614947565b60208152600061403a60208301846146aa565b6001600160401b0384168152606060208201526000614efe60608301856146aa565b8281036040840152614f1081856146ef565b9695505050505050565b6001600160401b0384168152606060208201526000614f3c60608301856146aa565b8281036040840152614f1081856148d1565b600080600060608486031215614f6357600080fd5b614f6c8461456f565b9250614f7a6020850161456f565b9150604084013590509250925092565b600080600060608486031215614f9f57600080fd5b83359250614faf6020850161456f565b9150614fbd6040850161456f565b90509250925092565b6020808252825182820181905260009190848201906040850190845b81811015614ffe57835183529284019291840191600101614fe2565b50909695505050505050565b6000806020838503121561501d57600080fd5b82356001600160401b0381111561503357600080fd5b614627858286016147e3565b634e487b7160e01b600052603260045260246000fd5b60ff8116811461113f57600080fd5b60006060828403121561507657600080fd5b604051606081018181106001600160401b0382111715615098576150986149bb565b8060405250809150823581526020830135602082015260408301356150bc81615055565b6040919091015292915050565b600061012082840312156150dc57600080fd5b60405160e081018181106001600160401b03821117156150fe576150fe6149bb565b60405290508061510d8361456f565b815261511b6020840161456f565b602082015261512c6040840161456f565b6040820152606083013560608201526080830135608082015261515160a084016147cc565b60a08201526151638460c08501615064565b60c08201525092915050565b6000610120828403121561518257600080fd5b61403a83836150c9565b60006020828403121561519e57600080fd5b61403a826147cc565b6000808335601e198436030181126151be57600080fd5b8301803591506001600160401b038211156151d857600080fd5b6020019150600581901b36038213156145eb57600080fd5b60006060828403121561520257600080fd5b61403a8383615064565b600181811c9082168061522057607f821691505b60208210810361524057634e487b7160e01b600052602260045260246000fd5b50919050565b6000610160828403121561525957600080fd5b6152616149d1565b61526a8361456f565b81526152786020840161456f565b60208201526152896040840161456f565b604082015260608301356060820152608083013560808201526152ae60a084016147cc565b60a082015260c083013560c08201526152c960e084016147cc565b60e08201526101006152dd85828601615064565b908201529392505050565b602080825260159082015274496e76616c696420746f6b656e206164647265737360581b604082015260600190565b815160009082906020808601845b8381101561534a5781516001600160a01b031685529382019390820190600101615325565b50929695505050505050565b815160009082906020808601845b8381101561534a5781516001600160401b031685529382019390820190600101615364565b6020808252600f908201526e125b9d985b1a590818da185a5b9259608a1b604082015260600190565b60008235609e198336030181126153c857600080fd5b9190910192915050565b6000808335601e198436030181126153e957600080fd5b8301803591506001600160401b0382111561540357600080fd5b6020019150368190038213156145eb57600080fd5b634e487b7160e01b600052601160045260246000fd5b6001600160401b0381811683821601908082111561296057612960615418565b6001600160401b0381811683821602808216919082811461547157615471615418565b505092915050565b81810381811115610a4057610a40615418565b634e487b7160e01b600052603160045260246000fd5b600081548084526020808501945083600052602060002060005b838110156146e45781546001600160a01b0316875295820195600191820191016154bc565b6060815260006154f460608301866154a2565b828103602084015261550681866154a2565b90508281036040840152614f1081856146ef565b60608152600061552d60608301866146aa565b828103602084015261550681866146aa565b60006001820161555157615551615418565b5060010190565b60006020828403121561556a57600080fd5b5051919050565b60006001600160401b0380831681810361558d5761558d615418565b6001019392505050565b8183823760009101908152919050565b81358155602082013560018201556002810160408301356155c781615055565b815460ff191660ff919091161790555050565b601f821115611000576000816000526020600020601f850160051c810160208610156156035750805b601f850160051c820191505b818110156156225782815560010161560f565b505050505050565b6001600160401b03831115615641576156416149bb565b6156558361564f835461520c565b836155da565b6000601f84116001811461568957600085156156715750838201355b600019600387901b1c1916600186901b1783556156e3565b600083815260209020601f19861690835b828110156156ba578685013582556020948501946001909201910161569a565b50868210156156d75760001960f88860031b161c19848701351681555b505060018560011b0183555b5050505050565b80820180821115610a4057610a40615418565b600082516153c881846020870161492356fe360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800cd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033009b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f00a49807205ce4d355092ef5a8a18f56e8913cf4a201fbe287825b095693c21775a264697066735822122062927f601566df1973e8a192a1f60b1743d84ffe56879e3e4cf6e940180c470864736f6c63430008160033",
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

// MakeDepositMessage is a free data retrieval call binding the contract method 0x5a53db10.
//
// Solidity: function makeDepositMessage(address user, address destination, address token, uint256 amount, uint256 chainId, uint64 blockNumber, bytes32 txHash, uint64 logIndex) pure returns(bytes32)
func (_BridgeHub *BridgeHubCaller) MakeDepositMessage(opts *bind.CallOpts, user common.Address, destination common.Address, token common.Address, amount *big.Int, chainId *big.Int, blockNumber uint64, txHash [32]byte, logIndex uint64) ([32]byte, error) {
	var out []interface{}
	err := _BridgeHub.contract.Call(opts, &out, "makeDepositMessage", user, destination, token, amount, chainId, blockNumber, txHash, logIndex)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MakeDepositMessage is a free data retrieval call binding the contract method 0x5a53db10.
//
// Solidity: function makeDepositMessage(address user, address destination, address token, uint256 amount, uint256 chainId, uint64 blockNumber, bytes32 txHash, uint64 logIndex) pure returns(bytes32)
func (_BridgeHub *BridgeHubSession) MakeDepositMessage(user common.Address, destination common.Address, token common.Address, amount *big.Int, chainId *big.Int, blockNumber uint64, txHash [32]byte, logIndex uint64) ([32]byte, error) {
	return _BridgeHub.Contract.MakeDepositMessage(&_BridgeHub.CallOpts, user, destination, token, amount, chainId, blockNumber, txHash, logIndex)
}

// MakeDepositMessage is a free data retrieval call binding the contract method 0x5a53db10.
//
// Solidity: function makeDepositMessage(address user, address destination, address token, uint256 amount, uint256 chainId, uint64 blockNumber, bytes32 txHash, uint64 logIndex) pure returns(bytes32)
func (_BridgeHub *BridgeHubCallerSession) MakeDepositMessage(user common.Address, destination common.Address, token common.Address, amount *big.Int, chainId *big.Int, blockNumber uint64, txHash [32]byte, logIndex uint64) ([32]byte, error) {
	return _BridgeHub.Contract.MakeDepositMessage(&_BridgeHub.CallOpts, user, destination, token, amount, chainId, blockNumber, txHash, logIndex)
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

// TokenPair is a free data retrieval call binding the contract method 0x2b1a7b58.
//
// Solidity: function tokenPair(uint256 , address ) view returns(address)
func (_BridgeHub *BridgeHubCaller) TokenPair(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (common.Address, error) {
	var out []interface{}
	err := _BridgeHub.contract.Call(opts, &out, "tokenPair", arg0, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenPair is a free data retrieval call binding the contract method 0x2b1a7b58.
//
// Solidity: function tokenPair(uint256 , address ) view returns(address)
func (_BridgeHub *BridgeHubSession) TokenPair(arg0 *big.Int, arg1 common.Address) (common.Address, error) {
	return _BridgeHub.Contract.TokenPair(&_BridgeHub.CallOpts, arg0, arg1)
}

// TokenPair is a free data retrieval call binding the contract method 0x2b1a7b58.
//
// Solidity: function tokenPair(uint256 , address ) view returns(address)
func (_BridgeHub *BridgeHubCallerSession) TokenPair(arg0 *big.Int, arg1 common.Address) (common.Address, error) {
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

// DepositConfirm is a paid mutator transaction binding the contract method 0x490b105c.
//
// Solidity: function depositConfirm((address,address,address,uint256,uint256,uint64,bytes32,uint64,(uint256,uint256,uint8))[] deposits) returns()
func (_BridgeHub *BridgeHubTransactor) DepositConfirm(opts *bind.TransactOpts, deposits []DepositConfirm) (*types.Transaction, error) {
	return _BridgeHub.contract.Transact(opts, "depositConfirm", deposits)
}

// DepositConfirm is a paid mutator transaction binding the contract method 0x490b105c.
//
// Solidity: function depositConfirm((address,address,address,uint256,uint256,uint64,bytes32,uint64,(uint256,uint256,uint8))[] deposits) returns()
func (_BridgeHub *BridgeHubSession) DepositConfirm(deposits []DepositConfirm) (*types.Transaction, error) {
	return _BridgeHub.Contract.DepositConfirm(&_BridgeHub.TransactOpts, deposits)
}

// DepositConfirm is a paid mutator transaction binding the contract method 0x490b105c.
//
// Solidity: function depositConfirm((address,address,address,uint256,uint256,uint64,bytes32,uint64,(uint256,uint256,uint8))[] deposits) returns()
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

// SetTokenPair is a paid mutator transaction binding the contract method 0xdcf70f3b.
//
// Solidity: function setTokenPair(uint256 chainId, address token, address bridgedToken) returns()
func (_BridgeHub *BridgeHubTransactor) SetTokenPair(opts *bind.TransactOpts, chainId *big.Int, token common.Address, bridgedToken common.Address) (*types.Transaction, error) {
	return _BridgeHub.contract.Transact(opts, "setTokenPair", chainId, token, bridgedToken)
}

// SetTokenPair is a paid mutator transaction binding the contract method 0xdcf70f3b.
//
// Solidity: function setTokenPair(uint256 chainId, address token, address bridgedToken) returns()
func (_BridgeHub *BridgeHubSession) SetTokenPair(chainId *big.Int, token common.Address, bridgedToken common.Address) (*types.Transaction, error) {
	return _BridgeHub.Contract.SetTokenPair(&_BridgeHub.TransactOpts, chainId, token, bridgedToken)
}

// SetTokenPair is a paid mutator transaction binding the contract method 0xdcf70f3b.
//
// Solidity: function setTokenPair(uint256 chainId, address token, address bridgedToken) returns()
func (_BridgeHub *BridgeHubTransactorSession) SetTokenPair(chainId *big.Int, token common.Address, bridgedToken common.Address) (*types.Transaction, error) {
	return _BridgeHub.Contract.SetTokenPair(&_BridgeHub.TransactOpts, chainId, token, bridgedToken)
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
	User        common.Address
	Destination common.Address
	Token       common.Address
	Amount      *big.Int
	ChainId     *big.Int
	BlockNumber uint64
	TxHash      [32]byte
	LogIndex    uint64
	Nonce       uint64
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0x63b86360efb9dda5b052fea4b2df3ca28fb97895a7714ee17dc371a5a472c685.
//
// Solidity: event Deposit(bytes32 indexed message, address indexed user, address destination, address indexed token, uint256 amount, uint256 chainId, uint64 blockNumber, bytes32 txHash, uint64 logIndex, uint64 nonce)
func (_BridgeHub *BridgeHubFilterer) FilterDeposit(opts *bind.FilterOpts, message [][32]byte, user []common.Address, token []common.Address) (*BridgeHubDepositIterator, error) {

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

// WatchDeposit is a free log subscription operation binding the contract event 0x63b86360efb9dda5b052fea4b2df3ca28fb97895a7714ee17dc371a5a472c685.
//
// Solidity: event Deposit(bytes32 indexed message, address indexed user, address destination, address indexed token, uint256 amount, uint256 chainId, uint64 blockNumber, bytes32 txHash, uint64 logIndex, uint64 nonce)
func (_BridgeHub *BridgeHubFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *BridgeHubDeposit, message [][32]byte, user []common.Address, token []common.Address) (event.Subscription, error) {

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

// ParseDeposit is a log parse operation binding the contract event 0x63b86360efb9dda5b052fea4b2df3ca28fb97895a7714ee17dc371a5a472c685.
//
// Solidity: event Deposit(bytes32 indexed message, address indexed user, address destination, address indexed token, uint256 amount, uint256 chainId, uint64 blockNumber, bytes32 txHash, uint64 logIndex, uint64 nonce)
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
	ChainId      *big.Int
	Token        common.Address
	BridgedToken common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterTokenPairSet is a free log retrieval operation binding the contract event 0x5478d55b1b0642b4b219976302fbc1f304d09602ef98dcacbeb94121af8a2a54.
//
// Solidity: event TokenPairSet(uint256 indexed chainId, address indexed token, address indexed bridgedToken)
func (_BridgeHub *BridgeHubFilterer) FilterTokenPairSet(opts *bind.FilterOpts, chainId []*big.Int, token []common.Address, bridgedToken []common.Address) (*BridgeHubTokenPairSetIterator, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var bridgedTokenRule []interface{}
	for _, bridgedTokenItem := range bridgedToken {
		bridgedTokenRule = append(bridgedTokenRule, bridgedTokenItem)
	}

	logs, sub, err := _BridgeHub.contract.FilterLogs(opts, "TokenPairSet", chainIdRule, tokenRule, bridgedTokenRule)
	if err != nil {
		return nil, err
	}
	return &BridgeHubTokenPairSetIterator{contract: _BridgeHub.contract, event: "TokenPairSet", logs: logs, sub: sub}, nil
}

// WatchTokenPairSet is a free log subscription operation binding the contract event 0x5478d55b1b0642b4b219976302fbc1f304d09602ef98dcacbeb94121af8a2a54.
//
// Solidity: event TokenPairSet(uint256 indexed chainId, address indexed token, address indexed bridgedToken)
func (_BridgeHub *BridgeHubFilterer) WatchTokenPairSet(opts *bind.WatchOpts, sink chan<- *BridgeHubTokenPairSet, chainId []*big.Int, token []common.Address, bridgedToken []common.Address) (event.Subscription, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var bridgedTokenRule []interface{}
	for _, bridgedTokenItem := range bridgedToken {
		bridgedTokenRule = append(bridgedTokenRule, bridgedTokenItem)
	}

	logs, sub, err := _BridgeHub.contract.WatchLogs(opts, "TokenPairSet", chainIdRule, tokenRule, bridgedTokenRule)
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

// ParseTokenPairSet is a log parse operation binding the contract event 0x5478d55b1b0642b4b219976302fbc1f304d09602ef98dcacbeb94121af8a2a54.
//
// Solidity: event TokenPairSet(uint256 indexed chainId, address indexed token, address indexed bridgedToken)
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
