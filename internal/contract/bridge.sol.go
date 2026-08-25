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

// DepositWithPermit is an auto generated low-level Go binding around an user-defined struct.
type DepositWithPermit struct {
	User        common.Address
	Destination common.Address
	Token       common.Address
	Amount      *big.Int
	Deadline    uint64
	Signature   Signature
}

// Signature is an auto generated low-level Go binding around an user-defined struct.
type Signature struct {
	R *big.Int
	S *big.Int
	V uint8
}

// ValidatorSet is an auto generated low-level Go binding around an user-defined struct.
type ValidatorSet struct {
	Epoch      uint64
	Validators []common.Address
	Powers     []uint64
}

// ValidatorSetUpdateRequest is an auto generated low-level Go binding around an user-defined struct.
type ValidatorSetUpdateRequest struct {
	Epoch         uint64
	HotAddresses  []common.Address
	ColdAddresses []common.Address
	Powers        []uint64
}

// Withdrawal is an auto generated low-level Go binding around an user-defined struct.
type Withdrawal struct {
	User                 common.Address
	Destination          common.Address
	Token                common.Address
	Amount               *big.Int
	ChainId              *big.Int
	Nonce                uint64
	RequestedTime        uint64
	RequestedBlockNumber uint64
	Message              [32]byte
}

// WithdrawalRequest is an auto generated low-level Go binding around an user-defined struct.
type WithdrawalRequest struct {
	User        common.Address
	Destination common.Address
	Token       common.Address
	Amount      *big.Int
	ChainId     *big.Int
	Nonce       uint64
	Signatures  []Signature
}

// BridgeMetaData contains all meta data concerning the Bridge contract.
var BridgeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"hotAddresses\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"coldAddresses\",\"type\":\"address[]\"},{\"internalType\":\"uint64[]\",\"name\":\"powers\",\"type\":\"uint64[]\"},{\"internalType\":\"uint64\",\"name\":\"_disputePeriodSeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_blockDurationMillis\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_lockerThreshold\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"EnforcedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExpectedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"newBlockDurationMillis\",\"type\":\"uint64\"}],\"name\":\"ChangedBlockDurationMillis\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"newDisputePeriodSeconds\",\"type\":\"uint64\"}],\"name\":\"ChangedDisputePeriodSeconds\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"newLockerThreshold\",\"type\":\"uint64\"}],\"name\":\"ChangedLockerThreshold\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"destination\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"message\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"errorCode\",\"type\":\"uint32\"}],\"name\":\"FailedWithdrawal\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"epoch\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"hotValidatorSetHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"coldValidatorSetHash\",\"type\":\"bytes32\"}],\"name\":\"FinalizedValidatorSetUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"message\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"destination\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"name\":\"FinalizedWithdrawal\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"destination\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"requestedTime\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"requestedBlockNumber\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"message\",\"type\":\"bytes32\"}],\"indexed\":false,\"internalType\":\"structWithdrawal\",\"name\":\"withdrawal\",\"type\":\"tuple\"}],\"name\":\"InvalidatedWithdrawal\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"finalizer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isFinalizer\",\"type\":\"bool\"}],\"name\":\"ModifiedFinalizer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"locker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isLocker\",\"type\":\"bool\"}],\"name\":\"ModifiedLocker\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"epoch\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"hotValidatorSetHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"coldValidatorSetHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"updateTime\",\"type\":\"uint64\"}],\"name\":\"RequestedValidatorSetUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"message\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"destination\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"requestedTime\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"requestedBlockNumber\",\"type\":\"uint64\"}],\"name\":\"RequestedWithdrawal\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"destination\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"deadline\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"r\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"s\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"}],\"internalType\":\"structSignature\",\"name\":\"signature\",\"type\":\"tuple\"}],\"internalType\":\"structDepositWithPermit[]\",\"name\":\"deposits\",\"type\":\"tuple[]\"}],\"name\":\"batchedDepositWithPermit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"messages\",\"type\":\"bytes32[]\"}],\"name\":\"batchedFinalizeWithdrawals\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"destination\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"r\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"s\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"}],\"internalType\":\"structSignature[]\",\"name\":\"signatures\",\"type\":\"tuple[]\"}],\"internalType\":\"structWithdrawalRequest[]\",\"name\":\"withdrawalRequests\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"epoch\",\"type\":\"uint64\"},{\"internalType\":\"address[]\",\"name\":\"validators\",\"type\":\"address[]\"},{\"internalType\":\"uint64[]\",\"name\":\"powers\",\"type\":\"uint64[]\"}],\"internalType\":\"structValidatorSet\",\"name\":\"hotValidatorSet\",\"type\":\"tuple\"}],\"name\":\"batchedRequestWithdrawals\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"blockDurationMillis\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"newBlockDurationMillis\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"epoch\",\"type\":\"uint64\"},{\"internalType\":\"address[]\",\"name\":\"validators\",\"type\":\"address[]\"},{\"internalType\":\"uint64[]\",\"name\":\"powers\",\"type\":\"uint64[]\"}],\"internalType\":\"structValidatorSet\",\"name\":\"activeColdValidatorSet\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"r\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"s\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"}],\"internalType\":\"structSignature[]\",\"name\":\"signatures\",\"type\":\"tuple[]\"}],\"name\":\"changeBlockDurationMillis\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"newDisputePeriodSeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"epoch\",\"type\":\"uint64\"},{\"internalType\":\"address[]\",\"name\":\"validators\",\"type\":\"address[]\"},{\"internalType\":\"uint64[]\",\"name\":\"powers\",\"type\":\"uint64[]\"}],\"internalType\":\"structValidatorSet\",\"name\":\"activeColdValidatorSet\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"r\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"s\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"}],\"internalType\":\"structSignature[]\",\"name\":\"signatures\",\"type\":\"tuple[]\"}],\"name\":\"changeDisputePeriodSeconds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"newLockerThreshold\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"epoch\",\"type\":\"uint64\"},{\"internalType\":\"address[]\",\"name\":\"validators\",\"type\":\"address[]\"},{\"internalType\":\"uint64[]\",\"name\":\"powers\",\"type\":\"uint64[]\"}],\"internalType\":\"structValidatorSet\",\"name\":\"activeColdValidatorSet\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"r\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"s\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"}],\"internalType\":\"structSignature[]\",\"name\":\"signatures\",\"type\":\"tuple[]\"}],\"name\":\"changeLockerThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"coldValidatorSetHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"destination\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"disputePeriodSeconds\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"domainSeparator\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"epoch\",\"type\":\"uint64\"},{\"internalType\":\"address[]\",\"name\":\"hotAddresses\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"coldAddresses\",\"type\":\"address[]\"},{\"internalType\":\"uint64[]\",\"name\":\"powers\",\"type\":\"uint64[]\"}],\"internalType\":\"structValidatorSetUpdateRequest\",\"name\":\"newValidatorSet\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"epoch\",\"type\":\"uint64\"},{\"internalType\":\"address[]\",\"name\":\"validators\",\"type\":\"address[]\"},{\"internalType\":\"uint64[]\",\"name\":\"powers\",\"type\":\"uint64[]\"}],\"internalType\":\"structValidatorSet\",\"name\":\"activeColdValidatorSet\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"r\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"s\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"}],\"internalType\":\"structSignature[]\",\"name\":\"signatures\",\"type\":\"tuple[]\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"name\":\"emergencyUnlock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"epoch\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"finalizeValidatorSetUpdate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"finalizedWithdrawals\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"finalizers\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLockersVotingLock\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"hotValidatorSetHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"messages\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"epoch\",\"type\":\"uint64\"},{\"internalType\":\"address[]\",\"name\":\"validators\",\"type\":\"address[]\"},{\"internalType\":\"uint64[]\",\"name\":\"powers\",\"type\":\"uint64[]\"}],\"internalType\":\"structValidatorSet\",\"name\":\"activeColdValidatorSet\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"r\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"s\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"}],\"internalType\":\"structSignature[]\",\"name\":\"signatures\",\"type\":\"tuple[]\"}],\"name\":\"invalidateWithdrawals\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"locker\",\"type\":\"address\"}],\"name\":\"isVotingLock\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lockerThreshold\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"lockers\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"finalizer\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_isFinalizer\",\"type\":\"bool\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"epoch\",\"type\":\"uint64\"},{\"internalType\":\"address[]\",\"name\":\"validators\",\"type\":\"address[]\"},{\"internalType\":\"uint64[]\",\"name\":\"powers\",\"type\":\"uint64[]\"}],\"internalType\":\"structValidatorSet\",\"name\":\"activeValidatorSet\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"r\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"s\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"}],\"internalType\":\"structSignature[]\",\"name\":\"signatures\",\"type\":\"tuple[]\"}],\"name\":\"modifyFinalizer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"locker\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_isLocker\",\"type\":\"bool\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"epoch\",\"type\":\"uint64\"},{\"internalType\":\"address[]\",\"name\":\"validators\",\"type\":\"address[]\"},{\"internalType\":\"uint64[]\",\"name\":\"powers\",\"type\":\"uint64[]\"}],\"internalType\":\"structValidatorSet\",\"name\":\"activeValidatorSet\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"r\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"s\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"}],\"internalType\":\"structSignature[]\",\"name\":\"signatures\",\"type\":\"tuple[]\"}],\"name\":\"modifyLocker\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingValidatorSetUpdate\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"epoch\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"totalValidatorPower\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"updateTime\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"updateBlockNumber\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nValidators\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"hotValidatorSetHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"coldValidatorSetHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"requestedWithdrawals\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"destination\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"requestedTime\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"requestedBlockNumber\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"message\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalValidatorPower\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unvoteEmergencyLock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"epoch\",\"type\":\"uint64\"},{\"internalType\":\"address[]\",\"name\":\"hotAddresses\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"coldAddresses\",\"type\":\"address[]\"},{\"internalType\":\"uint64[]\",\"name\":\"powers\",\"type\":\"uint64[]\"}],\"internalType\":\"structValidatorSetUpdateRequest\",\"name\":\"newValidatorSet\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"epoch\",\"type\":\"uint64\"},{\"internalType\":\"address[]\",\"name\":\"validators\",\"type\":\"address[]\"},{\"internalType\":\"uint64[]\",\"name\":\"powers\",\"type\":\"uint64[]\"}],\"internalType\":\"structValidatorSet\",\"name\":\"activeHotValidatorSet\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"r\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"s\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"}],\"internalType\":\"structSignature[]\",\"name\":\"signatures\",\"type\":\"tuple[]\"}],\"name\":\"updateValidatorSet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"usedMessages\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"voteEmergencyLock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"withdrawalsInvalidated\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b5060405162004603380380620046038339810160408190526200003491620007e4565b60018055604080517f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f60208201527f7aa5ae620294318af92bf4e2b2a729646c932a80312a5fa630da993a2ef5cc10918101919091527fc89efdaa54c0f20c7adf612882df0950f5a951637e0307cdcb4c672f298b8bc660608201524660808201523060a082015260c00160408051601f198184030181529190528051602090910120601155620000e5846200041a565b600d60086101000a8154816001600160401b0302191690836001600160401b031602179055508451865114620001765760405162461bcd60e51b815260206004820152602b60248201527f486f7420616e6420636f6c642076616c696461746f722073657473206c656e6760448201526a0e8d040dad2e6dac2e8c6d60ab1b60648201526084015b60405180910390fd5b620001a4604051806060016040528060006001600160401b0316815260200160608152602001606081525090565b506040805160608101825260008082526020820189905291810186905290620001cd8262000505565b6002819055604080516060808201835260008083526020808401839052928401829052835191820184528082529181018b905291820189905291925090620002158262000505565b6003819055600d80546001600160801b0316600160801b6001600160401b038b8116919091026001600160c01b031691909117600160c01b8a83160217909155600b80546001600160401b0319169188169190911790559050620002798a62000636565b600254600354604080516000815260208101939093528201526001600160401b03421660608201527f420bbe99bd2c52ec500d33614359525f3ef7bb3358c0e07d1312db0941cbf2f49060800160405180910390a16040805160e0810182526000808252600d546801000000000000000090046001600160401b0316602083015291810191909152606081016200030d4390565b6001600160401b0390811682528c51811660208084019190915260025460408085018290526003546060958601819052865160048054898701518a8601518b8b0151948a166001600160801b03199093169290921768010000000000000000918a1691909102176001600160801b0316600160801b918916919091026001600160c01b031617600160c01b928816929092029190911790556080870151600580546001600160401b031916919096161790945560a086015160065560c09095015160075584516000815291820152928301527f87da17ff65d815d1e1c369cb3bbda9a11af181b92dc52681a2779419781c6270910160405180910390a15050505050505050505062000a0d565b60008060005b8351816001600160401b031610156200047d5783816001600160401b03168151811062000451576200045162000910565b6020026020010151826200046691906200093c565b915080620004748162000966565b91505062000420565b506000816001600160401b031611620004ff5760405162461bcd60e51b815260206004820152603460248201527f5375626d69747465642076616c696461746f7220706f77657273206d7573742060448201527f62652067726561746572207468616e207a65726f00000000000000000000000060648201526084016200016d565b92915050565b600081604001515182602001515114620005625760405162461bcd60e51b815260206004820152601760248201527f4d616c666f726d65642076616c696461746f722073657400000000000000000060448201526064016200016d565b60007fcf7a991d34f65202b9a5ebe03e28c3fd6f86e1f75fabbddd532864507554c66783600001518460200151604051602001620005a1919062000997565b604051602081830303815290604052805190602001208560400151604051602001620005ce9190620009d8565b604051602081830303815290604052805190602001206040516020016200061794939291909384526001600160401b039290921660208401526040830152606082015260800190565b60408051601f1981840301815291905280516020909101209392505050565b805160005b816001600160401b0316816001600160401b03161015620006c557600083826001600160401b03168151811062000676576200067662000910565b6020908102919091018101516001600160a01b031660009081526009825260408082208054600160ff199182168117909255600c9094529120805490921681179091559190910190506200063b565b505050565b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f191681016001600160401b03811182821017156200070b576200070b620006ca565b604052919050565b60006001600160401b038211156200072f576200072f620006ca565b5060051b60200190565b600082601f8301126200074b57600080fd5b81516020620007646200075e8362000713565b620006e0565b8083825260208201915060208460051b8701019350868411156200078757600080fd5b602086015b84811015620007bc5780516001600160a01b0381168114620007ae5760008081fd5b83529183019183016200078c565b509695505050505050565b80516001600160401b0381168114620007df57600080fd5b919050565b60008060008060008060c08789031215620007fe57600080fd5b86516001600160401b03808211156200081657600080fd5b620008248a838b0162000739565b97506020915081890151818111156200083c57600080fd5b6200084a8b828c0162000739565b9750506040890151818111156200086057600080fd5b89019050601f81018a136200087457600080fd5b8051620008856200075e8262000713565b81815260059190911b8201830190838101908c831115620008a557600080fd5b928401925b82841015620008ce57620008be84620007c7565b82529284019290840190620008aa565b8098505050505050620008e460608801620007c7565b9250620008f460808801620007c7565b91506200090460a08801620007c7565b90509295509295509295565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b6001600160401b038181168382160190808211156200095f576200095f62000926565b5092915050565b60006001600160401b038281166002600160401b031981016200098d576200098d62000926565b6001019392505050565b815160009082906020808601845b83811015620009cc5781516001600160a01b031685529382019390820190600101620009a5565b50929695505050505050565b815160009082906020808601845b83811015620009cc5781516001600160401b031685529382019390820190600101620009e6565b613be68062000a1d6000396000f3fe608060405234801561001057600080fd5b50600436106101fb5760003560e01c80638340f5491161011a578063b5f389f3116100ad578063e3e6c4411161007c578063e3e6c441146105b9578063e73ea41e146105cc578063f698da25146105df578063f8156a6e146105e8578063fc3f7ad31461060257600080fd5b8063b5f389f3146104e2578063c10ee9ae146104f5578063c5bdf3ca14610583578063cea75eb71461059657600080fd5b80639d5bc9e1116100e95780639d5bc9e114610494578063a14238e7146104ae578063b0801e54146104d1578063b091049c146104da57600080fd5b80638340f54914610448578063900cf0cf1461045b57806391ed13441461046e5780639770e2c81461048157600080fd5b806342082828116101925780635a028400116101615780635a0284001461032a5780635c975abb1461034d5780636c9fc7b2146103585780637694c6fa1461036b57600080fd5b806342082828146102d75780634878ee53146102fa5780634aad62101461030257806353f79ef41461031557600080fd5b80630fb61a2e116101ce5780630fb61a2e1461026b578063180f2e8c1461027e5780632c8e7a21146102915780633a37326e146102c457600080fd5b806305355e2314610200578063058731e5146102305780630756183b1461023a5780630f71143814610254575b600080fd5b600b54610213906001600160401b031681565b6040516001600160401b0390911681526020015b60405180910390f35b610238610615565b005b600d5461021390600160801b90046001600160401b031681565b61025d60035481565b604051908152602001610227565b61023861027936600461316b565b61073f565b61023861028c366004613273565b6108c2565b6102b461029f36600461330f565b60096020526000908152604090205460ff1681565b6040519015158152602001610227565b6102b46102d236600461330f565b610a0d565b6102b46102e5366004613331565b60106020526000908152604090205460ff1681565b610238610a8b565b61023861031036600461334a565b610b91565b61031d610c6b565b6040516102279190613393565b6102b4610338366004613331565b60086020526000908152604090205460ff1681565b60005460ff166102b4565b6102386103663660046133e0565b610ccd565b6103e4610379366004613331565b600e6020526000908152604090208054600182015460028301546003840154600485015460058601546006909601546001600160a01b03958616969486169590931693919290916001600160401b0380821692600160401b8304821692600160801b90049091169089565b604080516001600160a01b039a8b168152988a1660208a01529690981695870195909552606086019390935260808501919091526001600160401b0390811660a085015290811660c08401521660e082015261010081019190915261012001610227565b6102386104563660046134e4565b610e04565b600d54610213906001600160401b031681565b61023861047c36600461334a565b610f4c565b61023861048f3660046135df565b611020565b600d5461021390600160c01b90046001600160401b031681565b6102b46104bc366004613331565b600f6020526000908152604090205460ff1681565b61025d60025481565b6102386111b5565b6102386104f03660046136ae565b611264565b600454600554600654600754610538936001600160401b0380821694600160401b8304821694600160801b8404831694600160c01b909404831693929091169187565b604080516001600160401b0398891681529688166020880152948716948601949094529185166060850152909316608083015260a082019290925260c081019190915260e001610227565b610238610591366004613810565b6112fe565b6102b46105a436600461330f565b600c6020526000908152604090205460ff1681565b6102386105c7366004613884565b611367565b6102386105da366004613273565b6114fe565b61025d60115481565b600d5461021390600160401b90046001600160401b031681565b61023861061036600461334a565b6115f2565b61061d6116e8565b610625611712565b61062e33611736565b600454600160801b90046001600160401b03166000036106ac5760405162461bcd60e51b815260206004820152602e60248201527f50656e64696e672076616c696461746f72207365742075706461746520616c7260448201526d1958591e48199a5b985b1a5e995960921b60648201526084015b60405180910390fd5b6004546000906106d5906001600160401b03600160801b8204811691600160c01b90041661179e565b905063ffffffff81161561072b5760405162461bcd60e51b815260206004820152601760248201527f5374696c6c20696e206469737075746520706572696f6400000000000000000060448201526064016106a3565b610733611857565b5061073d60018055565b565b60007fa0675b98ae6c277eba9efba80fbbdcac49c582fec62ca7833048cb970321ad2e85604051602001610773919061390b565b60408051601f198184030181528282528051602091820120908301939093528101919091526001600160401b03851660608201526080016040516020818303038152906040528051906020012090506107cb816118f3565b6107d9818484600354611964565b845160005b816001600160401b0316816001600160401b031610156108b95760016010600089846001600160401b03168151811061081957610819613941565b6020026020010151815260200190815260200160002060006101000a81548160ff0219169083151502179055507f4e1a2aef00d7868e1f49c3784b1b802acad3fd7e4c7fe753694a51d9b46346c5600e600089846001600160401b03168151811061088657610886613941565b602002602001015181526020019081526020016000206040516108a99190613957565b60405180910390a16001016107de565b50505050505050565b604080517f2bb2a8361d9a37d6bd9173b7a98ef8abfb3224b8f0c01732fa686695b7973af060208201526001600160a01b0387169181019190915284151560608201526001600160401b038416608082015260009060a001604051602081830303815290604052805190602001209050600085156109435750600254610948565b506003545b610951826118f3565b6109658261095e866139ef565b8584611964565b6001600160a01b03871660009081526009602052604090205460ff16801561098b575085155b801561099a575060005460ff16155b156109a8576109a887611c46565b6001600160a01b038716600081815260096020908152604091829020805460ff19168a151590811790915591519182527f26690dc5c5a9d2aa7ac3efa2b7c515652e4621a3e075d267bcac51c16fb9753291015b60405180910390a250505050505050565b600a54600090815b816001600160401b0316816001600160401b03161015610a8157836001600160a01b0316600a826001600160401b031681548110610a5557610a55613941565b6000918252602090912001546001600160a01b031603610a79575060019392505050565b600101610a15565b5060009392505050565b3360009081526009602052604090205460ff16610aba5760405162461bcd60e51b81526004016106a3906139fb565b610ac333610a0d565b15610b205760405162461bcd60e51b815260206004820152602760248201527f4c6f636b657220616c726561647920766f74656420666f7220656d657267656e6044820152666379206c6f636b60c81b60648201526084016106a3565b600a805460018101825560008290527fc65a7bb8d6351c1cf70c95a316cc6a92839c986682d98bc35f958f4883f9d2a80180546001600160a01b03191633179055600b5490546001600160401b03918216911610801590610b84575060005460ff16155b1561073d5761073d611df8565b604080517f12320cdb7a65d3e471a2a6ae0db5e58f3654c7d8bd5e6a1a70b3b8f2fea2aed760208201526001600160401b03808716928201929092529084166060820152600090608001604051602081830303815290604052805190602001209050610bfc816118f3565b610c0a818484600354611964565b600d80546001600160c01b0316600160c01b6001600160401b038816908102919091179091556040519081527f0ef2da393c3832a8f08ce447e14948d21e84f864facf7327137387bd0596a563906020015b60405180910390a15050505050565b6060600a805480602002602001604051908101604052809291908181526020018280548015610cc357602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311610ca5575b5050505050905090565b610cd56116e8565b610cdd611712565b805160005b816001600160401b0316816001600160401b03161015610df657610dee83826001600160401b031681518110610d1a57610d1a613941565b60200260200101516000015184836001600160401b031681518110610d4157610d41613941565b60200260200101516020015185846001600160401b031681518110610d6857610d68613941565b60200260200101516040015186856001600160401b031681518110610d8f57610d8f613941565b60200260200101516060015187866001600160401b031681518110610db657610db6613941565b60200260200101516080015188876001600160401b031681518110610ddd57610ddd613941565b602002602001015160a00151611e4d565b600101610ce2565b5050610e0160018055565b50565b610e0c6116e8565b610e14611712565b60008111610e755760405162461bcd60e51b815260206004820152602860248201527f4465706f73697420616d6f756e74206d7573742062652067726561746572207460448201526768616e207a65726f60c01b60648201526084016106a3565b6001600160a01b038316610ecb5760405162461bcd60e51b815260206004820152601b60248201527f496e76616c69642064657374696e6174696f6e2061646472657373000000000060448201526064016106a3565b610ee06001600160a01b038316333084611f6b565b6001600160a01b0383163314610f3e57604080516001600160a01b0385811682528416602082015290810182905233907f7cfff908a4b583f36430b25d75964c458d8ede8a99bd61be750e97ee1b2f3a969060600160405180910390a25b610f4760018055565b505050565b604080517f3cafbb096e40bdb2551d91252c05f542adbb6effe4860d14431e6605a04fd5fa60208201526001600160401b03808716928201929092529084166060820152600090608001604051602081830303815290604052805190602001209050610fb7816118f3565b610fc5818484600354611964565b600d805467ffffffffffffffff60801b1916600160801b6001600160401b038816908102919091179091556040519081527f04edaf680108675f58d2ea70e9e7886c39ed38b66439622f8362d36595fe816990602001610c5c565b611028611fd2565b60007f255d6a497f94e8b9aff3182744306fd858435def24a6ff727653948eadf175aa866000015187602001516040516020016110659190613a4a565b6040516020818303038152906040528051906020012088604001516040516020016110909190613a4a565b6040516020818303038152906040528051906020012089606001516040516020016110bb9190613a7d565b60408051601f198184030181528282528051602091820120908301969096526001600160401b03948516908201526060810192909252608082015260a0810192909252831660c082015260e001604051602081830303815290604052805190602001209050611129816118f3565b61119186611136876139ef565b8686808060200260200160405190810160405280939291908181526020016000905b828210156111845761117560608302860136819003810190613ab0565b81526020019060010190611158565b5050505050846001611ff5565b611199611857565b6111a5600a6000612e17565b6111ad6123da565b505050505050565b6111bd611712565b3360009081526009602052604090205460ff166111ec5760405162461bcd60e51b81526004016106a3906139fb565b6111f533610a0d565b61125b5760405162461bcd60e51b815260206004820152603160248201527f4c6f636b6572206973206e6f742063757272656e746c7920766f74696e6720666044820152706f7220656d657267656e6379206c6f636b60781b60648201526084016106a3565b61073d33611c46565b61126c6116e8565b611274611712565b815160005b816001600160401b0316816001600160401b031610156112ef57600084826001600160401b0316815181106112b0576112b0613941565b602002602001015190506112e6816000015182602001518360400151846060015185608001518660a001518a8860c00151612413565b50600101611279565b50506112fa60018055565b5050565b6113066116e8565b61130e611712565b61131733611736565b8060005b816001600160401b0316816001600160401b031610156112ef5761135f8484836001600160401b031681811061135357611353613941565b9050602002013561273f565b60010161131b565b61136f611712565b60025461137b836129b3565b146113ee5760405162461bcd60e51b815260206004820152603d60248201527f537570706c696564206163746976652076616c696461746f727320616e64207060448201527f6f7765727320646f206e6f74206d6174636820636865636b706f696e7400000060648201526084016106a3565b60007fcd26826da4f5c0e82ef8057ecacd8931dfb36167a70c820505f10826298cd05e8460000151856020015160405160200161142b9190613a4a565b6040516020818303038152906040528051906020012086604001516040516020016114569190613a4a565b6040516020818303038152906040528051906020012087606001516040516020016114819190613a7d565b604051602081830303815290604052805190602001206040516020016114d29594939291909485526001600160401b0393909316602085015260408401919091526060830152608082015260a00190565b6040516020818303038152906040528051906020012090506114f8848484846000611ff5565b50505050565b604080517f6f884b20c1fc555c21adb663cf93411199935f7263773056bc4711ea7266763160208201526001600160a01b0387169181019190915284151560608201526001600160401b038416608082015260009060a0016040516020818303038152906040528051906020012090506000851561157f5750600254611584565b506003545b61158d826118f3565b61159a8261095e866139ef565b6001600160a01b0387166000818152600c6020908152604091829020805460ff19168a151590811790915591519182527f2526bb92d75e00cfad8c7c16cb75f3e1073c854339e49b16baaad3067c2ed65a91016109fc565b604080517fd9c9d5e1590244de839782c61326c04e1b0506d5bf4a4750df02026b6df3953d60208201526001600160401b0380871692820192909252908416606082015260009060800160405160208183030381529060405280519060200120905061165d816118f3565b61166b818484600354611964565b600b805467ffffffffffffffff19166001600160401b03878116918217909255600a54909116108015906116a2575060005460ff16155b156116af576116af611df8565b6040516001600160401b03861681527f2dbe453726b24b2cee427a7d6e2dcc9f353f16bee104f3d21480157a0ee409f790602001610c5c565b60026001540361170b57604051633ee5aeb560e01b815260040160405180910390fd5b6002600155565b60005460ff161561073d5760405163d93c066560e01b815260040160405180910390fd5b6001600160a01b0381166000908152600c602052604090205460ff16610e015760405162461bcd60e51b815260206004820152601960248201527f53656e646572206973206e6f7420612066696e616c697a65720000000000000060448201526064016106a3565b600d5460009081906117c090600160801b90046001600160401b031685613ae2565b6001600160401b031642119050806117dc576003915050611851565b600d54439060009061180090600160801b90046001600160401b03166103e8613b09565b600d546001600160401b0391821691600160c01b909104166118228785613b34565b61182c9190613b09565b6001600160401b0316119050806118495760049350505050611851565b600093505050505b92915050565b6006546002819055600754600381905560048054600d80546001600160401b038084166001600160801b031990921691909117600160401b808504831602179182905567ffffffffffffffff60801b19909216909255604080519290911682526020820193909352918201527f87da17ff65d815d1e1c369cb3bbda9a11af181b92dc52681a2779419781c6270906060015b60405180910390a1565b60008181526008602052604090205460ff16156119495760405162461bcd60e51b81526020600482015260146024820152731b595cdcd859d948185b1c9958591e481d5cd95960621b60448201526064016106a3565b6000908152600860205260409020805460ff19166001179055565b8061196e846129b3565b146119f25760405162461bcd60e51b815260206004820152604860248201527f537570706c696564206163746976652076616c696461746f727320616e64207060448201527f6f7765727320646f206e6f74206d6174636820746865206163746976652063686064820152671958dadc1bda5b9d60c21b608482015260a4016106a3565b81516001600160401b038116611a3a5760405162461bcd60e51b815260206004820152600d60248201526c5369676e65727320656d70747960981b60448201526064016106a3565b6020840151516000908190815b816001600160401b0316816001600160401b03161015611b8d576000611a928a89866001600160401b031681518110611a8257611a82613941565b6020026020010151601154612add565b90508860200151826001600160401b031681518110611ab357611ab3613941565b60200260200101516001600160a01b0316816001600160a01b031603611b845760008960400151836001600160401b031681518110611af457611af4613941565b602002602001015190508086611b0a9190613ae2565b600d54909650611b2b90600160401b90046001600160401b03166002613b09565b6001600160401b0316611b3f876003613b09565b6001600160401b03161115611b55575050611b8d565b611b60600186613ae2565b9450866001600160401b0316856001600160401b031610611b82575050611b8d565b505b50600101611a47565b50600d54611bac90600160401b90046001600160401b03166002613b09565b6001600160401b0316611bc0846003613b09565b6001600160401b031611611c3c5760405162461bcd60e51b815260206004820152603b60248201527f5375626d69747465642076616c696461746f7220736574207369676e6174757260448201527f657320646f206e6f74206861766520656e6f75676820706f776572000000000060648201526084016106a3565b5050505050505050565b611c4e611712565b6001600160a01b03811660009081526009602052604090205460ff16611ccf5760405162461bcd60e51b815260206004820152603060248201527f41646472657373206973206e6f7420617574686f72697a656420746f206c6f6360448201526f1ac81cdb585c9d0818dbdb9d1c9858dd60821b60648201526084016106a3565b600a5460005b816001600160401b0316816001600160401b03161015610f4757826001600160a01b0316600a826001600160401b031681548110611d1557611d15613941565b6000918252602090912001546001600160a01b031603611df057600a611d3c600184613b34565b6001600160401b031681548110611d5557611d55613941565b600091825260209091200154600a80546001600160a01b03909216916001600160401b038416908110611d8a57611d8a613941565b9060005260206000200160006101000a8154816001600160a01b0302191690836001600160a01b03160217905550600a805480611dc957611dc9613b54565b600082815260209020810160001990810180546001600160a01b0319169055019055505050565b600101611cd5565b611e00611712565b6000805460ff191660011790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258611e353390565b6040516001600160a01b0390911681526020016118e9565b60408181015182516020840151925163d505accf60e01b81526001600160a01b038a811660048301523060248301819052604483018990526001600160401b038816606484015260ff909416608483015260a482019290925260c481019390935290919086169063d505accf9060e401600060405180830381600087803b158015611ed757600080fd5b505af1158015611eeb573d6000803e3d6000fd5b50611f05925050506001600160a01b038616888387611f6b565b866001600160a01b0316866001600160a01b0316146108b957604080516001600160a01b0388811682528781166020830152918101869052908816907f7cfff908a4b583f36430b25d75964c458d8ede8a99bd61be750e97ee1b2f3a96906060016109fc565b6040516001600160a01b0384811660248301528381166044830152606482018390526114f89186918216906323b872dd906084015b604051602081830303815290604052915060e01b6020820180516001600160e01b038381831617835250505050612ca4565b60005460ff1661073d57604051638dfc202b60e01b815260040160405180910390fd5b846040015151856020015151146120665760405162461bcd60e51b815260206004820152602f60248201527f4e657720686f7420616e6420636f6c642076616c696461746f7220736574732060448201526e0d8cadccee8d040dad2e6dac2e8c6d608b1b60648201526084016106a3565b846060015151856020015151146120d45760405162461bcd60e51b815260206004820152602c60248201527f4e65772076616c696461746f722073657420616e6420706f77657273206c656e60448201526b0cee8d040dad2e6dac2e8c6d60a31b60648201526084016106a3565b835185516001600160401b039182169116116121585760405162461bcd60e51b815260206004820152603d60248201527f4e65772076616c696461746f72207365742065706f6368206d7573742062652060448201527f67726561746572207468616e20746865206163746976652065706f636800000060648201526084016106a3565b60006121678660600151612d15565b90506000821561217a575060035461217f565b506002545b61218b84878784611964565b6121b8604051806060016040528060006001600160401b0316815260200160608152602001606081525090565b604051806060016040528089600001516001600160401b03168152602001896020015181526020018960600151815250905060006121f5826129b3565b9050612224604051806060016040528060006001600160401b0316815260200160608152602001606081525090565b5060408051606080820183528b516001600160401b031682528b83015160208301528b015191810191909152600061225b826129b3565b905060004290506040518060e001604052808d600001516001600160401b03168152602001886001600160401b03168152602001826001600160401b031681526020016122a54390565b6001600160401b03908116825260208881015151821681840152604080840189905260609384018790528451600480549387015187840151968801518616600160c01b026001600160c01b03978716600160801b02979097166fffffffffffffffffffffffffffffffff918716600160401b026001600160801b0319909616938716939093179490941793909316179390931790556080830151600580549190921667ffffffffffffffff199190911617905560a082015160065560c0909101516007558c5190517f420bbe99bd2c52ec500d33614359525f3ef7bb3358c0e07d1312db0941cbf2f4916123c4918790869086906001600160401b03948516815260208101939093526040830191909152909116606082015260800190565b60405180910390a1505050505050505050505050565b6123e2611fd2565b6000805460ff191690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa33611e35565b604080517fed44fa5a448edcc9a97caee522159268a5089c8700fe54f678309d7a73c9f6ec60208201526001600160a01b03808b16928201929092528189166060820152908716608082015260a0810186905260c081018590526001600160401b03841660e0820152600090610100016040516020818303038152906040528051906020012090506124b58160009081526010602052604090205460ff161590565b6124e6576040805182815260056020820152600080516020613b91833981519152910160405180910390a150611c3c565b60006040518061012001604052808b6001600160a01b031681526020018a6001600160a01b03168152602001896001600160a01b03168152602001888152602001878152602001866001600160401b03168152602001426001600160401b031681526020016125524390565b6001600160401b03908116825260209182018590526000858152600e909252604090912060050154919250600160401b90910416156125b9576040805183815260006020820152600080516020613b91833981519152910160405180910390a15050611c3c565b6125cf826125c6866139ef565b85600254611964565b6000828152600e6020908152604091829020835181546001600160a01b039182166001600160a01b0319918216811784559386015160018401805482851690841617905585870151600285018054948216949093169390931790915560608601516003840181905560808701516004850181905560a088015160058601805460c08b015160e08c01516001600160401b03818116600160801b0267ffffffffffffffff60801b19848316600160401b026001600160801b0319909616928816929092179490941716929092179092556101008b01516006909801889055985197987f6188a1db6e3e44fa683d03383b49bbe96178887d5c9d0174718f03ea6193b5ad9861272b98979594939291909788526001600160a01b039687166020890152949095166040870152606086019290925260808501526001600160401b0390811660a085015291821660c08401521660e08201526101000190565b60405180910390a250505050505050505050565b60008181526010602052604090205460ff1615612781576040805182815260056020820152600080516020613b9183398151915291015b60405180910390a150565b6000818152600f602052604090205460ff16156127bc576040805182815260016020820152600080516020613b918339815191529101612776565b6000818152600e6020908152604080832081516101208101835281546001600160a01b0390811682526001830154811694820194909452600282015490931691830191909152600381015460608301526004810154608083015260058101546001600160401b0380821660a0850152600160401b8204811660c08501819052600160801b9092041660e0840152600690910154610100830152909103612887576040805183815260026020820152600080516020613b91833981519152910160405180910390a15050565b600061289b8260c001518360e0015161179e565b905063ffffffff8116156128da576040805184815263ffffffff83166020820152600080516020613b91833981519152910160405180910390a1505050565b6000838152600f6020908152604091829020805460ff19166001179055830151606084015191840151612919926001600160a01b039091169190612de6565b81600001516001600160a01b03167f04cafa25a7826c4415eac718e45fc84f69b5539748e0206adbae70f9198875488361010001518460200151856040015186606001518760a001516040516129a69594939291909485526001600160a01b03938416602086015291909216604084015260608301919091526001600160401b0316608082015260a00190565b60405180910390a2505050565b600081604001515182602001515114612a0e5760405162461bcd60e51b815260206004820152601760248201527f4d616c666f726d65642076616c696461746f722073657400000000000000000060448201526064016106a3565b60007fcf7a991d34f65202b9a5ebe03e28c3fd6f86e1f75fabbddd532864507554c66783600001518460200151604051602001612a4b9190613a4a565b604051602081830303815290604052805190602001208560400151604051602001612a769190613a7d565b60405160208183030381529060405280519060200120604051602001612abe94939291909384526001600160401b039290921660208401526040830152606082015260800190565b60408051601f1981840301815291905280516020909101209392505050565b81516000908103612b305760405162461bcd60e51b815260206004820152601b60248201527f496e76616c6964207369676e6174757265202772272076616c7565000000000060448201526064016106a3565b8260200151600003612b845760405162461bcd60e51b815260206004820152601b60248201527f496e76616c6964207369676e6174757265202773272076616c7565000000000060448201526064016106a3565b60405161190160f01b6020820152602281018390526042810185905260009060620160408051601f1981840301815282825280516020918201208783015188518984015160008088529690940194859052919550600193612c009387939193845260ff9290921660208401526040830152606082015260800190565b6020604051602081039080840390855afa158015612c22573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b038116612c9b5760405162461bcd60e51b815260206004820152602d60248201527f496e76616c6964207369676e61747572652c207265636f76657265642074686560448201526c207a65726f206164647265737360981b60648201526084016106a3565b95945050505050565b600080602060008451602086016000885af180612cc7576040513d6000823e3d81fd5b50506000513d91508115612cdf578060011415612cec565b6001600160a01b0384163b155b156114f857604051635274afe760e01b81526001600160a01b03851660048201526024016106a3565b60008060005b8351816001600160401b03161015612d6f5783816001600160401b031681518110612d4857612d48613941565b602002602001015182612d5b9190613ae2565b915080612d6781613b6a565b915050612d1b565b506000816001600160401b0316116118515760405162461bcd60e51b815260206004820152603460248201527f5375626d69747465642076616c696461746f7220706f77657273206d7573742060448201527362652067726561746572207468616e207a65726f60601b60648201526084016106a3565b6040516001600160a01b03838116602483015260448201839052610f4791859182169063a9059cbb90606401611fa0565b5080546000825590600052602060002090810190610e0191905b80821115612e455760008155600101612e31565b5090565b634e487b7160e01b600052604160045260246000fd5b604051606081016001600160401b0381118282101715612e8157612e81612e49565b60405290565b60405160c081016001600160401b0381118282101715612e8157612e81612e49565b60405160e081016001600160401b0381118282101715612e8157612e81612e49565b604051601f8201601f191681016001600160401b0381118282101715612ef357612ef3612e49565b604052919050565b60006001600160401b03821115612f1457612f14612e49565b5060051b60200190565b80356001600160401b0381168114612f3557600080fd5b919050565b80356001600160a01b0381168114612f3557600080fd5b600082601f830112612f6257600080fd5b81356020612f77612f7283612efb565b612ecb565b8083825260208201915060208460051b870101935086841115612f9957600080fd5b602086015b84811015612fbc57612faf81612f3a565b8352918301918301612f9e565b509695505050505050565b600082601f830112612fd857600080fd5b81356020612fe8612f7283612efb565b8083825260208201915060208460051b87010193508684111561300a57600080fd5b602086015b84811015612fbc5761302081612f1e565b835291830191830161300f565b60006060828403121561303f57600080fd5b613047612e5f565b905061305282612f1e565b815260208201356001600160401b038082111561306e57600080fd5b61307a85838601612f51565b6020840152604084013591508082111561309357600080fd5b506130a084828501612fc7565b60408301525092915050565b6000606082840312156130be57600080fd5b6130c6612e5f565b90508135815260208201356020820152604082013560ff811681146130ea57600080fd5b604082015292915050565b600082601f83011261310657600080fd5b81356020613116612f7283612efb565b80838252602082019150606060206060860288010194508785111561313a57600080fd5b602087015b8581101561315e5761315189826130ac565b845292840192810161313f565b5090979650505050505050565b6000806000806080858703121561318157600080fd5b84356001600160401b038082111561319857600080fd5b818701915087601f8301126131ac57600080fd5b813560206131bc612f7283612efb565b82815260059290921b8401810191818101908b8411156131db57600080fd5b948201945b838610156131f9578535825294820194908201906131e0565b98506132089050898201612f1e565b96505050604087013591508082111561322057600080fd5b61322c8883890161302d565b9350606087013591508082111561324257600080fd5b5061324f878288016130f5565b91505092959194509250565b60006060828403121561326d57600080fd5b50919050565b600080600080600060a0868803121561328b57600080fd5b61329486612f3a565b9450602086013580151581146132a957600080fd5b93506132b760408701612f1e565b925060608601356001600160401b03808211156132d357600080fd5b6132df89838a0161325b565b935060808801359150808211156132f557600080fd5b50613302888289016130f5565b9150509295509295909350565b60006020828403121561332157600080fd5b61332a82612f3a565b9392505050565b60006020828403121561334357600080fd5b5035919050565b6000806000806080858703121561336057600080fd5b61336985612f1e565b935061337760208601612f1e565b925060408501356001600160401b038082111561322057600080fd5b6020808252825182820181905260009190848201906040850190845b818110156133d45783516001600160a01b0316835292840192918401916001016133af565b50909695505050505050565b600060208083850312156133f357600080fd5b82356001600160401b0381111561340957600080fd5b8301601f8101851361341a57600080fd5b8035613428612f7282612efb565b81815260089190911b8201830190838101908783111561344757600080fd5b928401925b828410156134d95761010084890312156134665760008081fd5b61346e612e87565b61347785612f3a565b8152613484868601612f3a565b868201526040613495818701612f3a565b908201526060858101359082015260806134b0818701612f1e565b9082015260a06134c28a8783016130ac565b90820152825261010093909301929084019061344c565b979650505050505050565b6000806000606084860312156134f957600080fd5b61350284612f3a565b925061351060208501612f3a565b9150604084013590509250925092565b60006080828403121561353257600080fd5b604051608081016001600160401b03828210818311171561355557613555612e49565b8160405282935061356585612f1e565b8352602085013591508082111561357b57600080fd5b61358786838701612f51565b602084015260408501359150808211156135a057600080fd5b6135ac86838701612f51565b604084015260608501359150808211156135c557600080fd5b506135d285828601612fc7565b6060830152505092915050565b6000806000806000608086880312156135f757600080fd5b85356001600160401b038082111561360e57600080fd5b61361a89838a01613520565b9650602088013591508082111561363057600080fd5b61363c89838a0161325b565b9550604088013591508082111561365257600080fd5b818801915088601f83011261366657600080fd5b81358181111561367557600080fd5b89602060608302850101111561368a57600080fd5b6020830195508094505050506136a260608701612f1e565b90509295509295909350565b600080604083850312156136c157600080fd5b82356001600160401b03808211156136d857600080fd5b818501915085601f8301126136ec57600080fd5b813560206136fc612f7283612efb565b82815260059290921b8401810191818101908984111561371b57600080fd5b8286015b848110156137e25780358681111561373657600080fd5b870160e0818d03601f1901121561374d5760008081fd5b613755612ea9565b613760868301612f3a565b815261376e60408301612f3a565b86820152606061377f818401612f3a565b60408301526080808401358284015260a0915081840135818401525060c06137a8818501612f1e565b8284015260e08401359150898211156137c15760008081fd5b6137cf8f89848701016130f5565b908301525084525091830191830161371f565b50965050860135925050808211156137f957600080fd5b506138068582860161325b565b9150509250929050565b6000806020838503121561382357600080fd5b82356001600160401b038082111561383a57600080fd5b818501915085601f83011261384e57600080fd5b81358181111561385d57600080fd5b8660208260051b850101111561387257600080fd5b60209290920196919550909350505050565b60008060006060848603121561389957600080fd5b83356001600160401b03808211156138b057600080fd5b6138bc87838801613520565b945060208601359150808211156138d257600080fd5b6138de8783880161302d565b935060408601359150808211156138f457600080fd5b50613901868287016130f5565b9150509250925092565b815160009082906020808601845b8381101561393557815185529382019390820190600101613919565b50929695505050505050565b634e487b7160e01b600052603260045260246000fd5b81546001600160a01b03168152610120810160018301546001600160a01b039081166020840152600284015416604080840191909152600384015460608401526004840154608084015260058401546001600160401b0380821660a08601529181901c821660c0850152906139dc60e08501828460801c166001600160401b03169052565b5050600683015461010083015292915050565b6000611851368361302d565b6020808252602f908201527f53656e646572206973206e6f7420617574686f72697a656420746f206c6f636b60408201526e081cdb585c9d0818dbdb9d1c9858dd608a1b606082015260800190565b815160009082906020808601845b838110156139355781516001600160a01b031685529382019390820190600101613a58565b815160009082906020808601845b838110156139355781516001600160401b031685529382019390820190600101613a8b565b600060608284031215613ac257600080fd5b61332a83836130ac565b634e487b7160e01b600052601160045260246000fd5b6001600160401b03818116838216019080821115613b0257613b02613acc565b5092915050565b6001600160401b03818116838216028082169190828114613b2c57613b2c613acc565b505092915050565b6001600160401b03828116828216039080821115613b0257613b02613acc565b634e487b7160e01b600052603160045260246000fd5b60006001600160401b03808316818103613b8657613b86613acc565b600101939250505056fe686cb4bac974cd11b0f8a75fc7c7764ed12cc46faaec53110f807aa802a7acb4a2646970667358221220ea73aa4df1e14c8214e21314e4fb4ed3c9c57c24c482e39f8b5162ca7d5a85b664736f6c63430008160033",
}

// BridgeABI is the input ABI used to generate the binding from.
// Deprecated: Use BridgeMetaData.ABI instead.
var BridgeABI = BridgeMetaData.ABI

// BridgeBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BridgeMetaData.Bin instead.
var BridgeBin = BridgeMetaData.Bin

// DeployBridge deploys a new Ethereum contract, binding an instance of Bridge to it.
func DeployBridge(auth *bind.TransactOpts, backend bind.ContractBackend, hotAddresses []common.Address, coldAddresses []common.Address, powers []uint64, _disputePeriodSeconds uint64, _blockDurationMillis uint64, _lockerThreshold uint64) (common.Address, *types.Transaction, *Bridge, error) {
	parsed, err := BridgeMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BridgeBin), backend, hotAddresses, coldAddresses, powers, _disputePeriodSeconds, _blockDurationMillis, _lockerThreshold)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Bridge{BridgeCaller: BridgeCaller{contract: contract}, BridgeTransactor: BridgeTransactor{contract: contract}, BridgeFilterer: BridgeFilterer{contract: contract}}, nil
}

// Bridge is an auto generated Go binding around an Ethereum contract.
type Bridge struct {
	BridgeCaller     // Read-only binding to the contract
	BridgeTransactor // Write-only binding to the contract
	BridgeFilterer   // Log filterer for contract events
}

// BridgeCaller is an auto generated read-only Go binding around an Ethereum contract.
type BridgeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BridgeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BridgeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BridgeSession struct {
	Contract     *Bridge           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BridgeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BridgeCallerSession struct {
	Contract *BridgeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// BridgeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BridgeTransactorSession struct {
	Contract     *BridgeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BridgeRaw is an auto generated low-level Go binding around an Ethereum contract.
type BridgeRaw struct {
	Contract *Bridge // Generic contract binding to access the raw methods on
}

// BridgeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BridgeCallerRaw struct {
	Contract *BridgeCaller // Generic read-only contract binding to access the raw methods on
}

// BridgeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BridgeTransactorRaw struct {
	Contract *BridgeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBridge creates a new instance of Bridge, bound to a specific deployed contract.
func NewBridge(address common.Address, backend bind.ContractBackend) (*Bridge, error) {
	contract, err := bindBridge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Bridge{BridgeCaller: BridgeCaller{contract: contract}, BridgeTransactor: BridgeTransactor{contract: contract}, BridgeFilterer: BridgeFilterer{contract: contract}}, nil
}

// NewBridgeCaller creates a new read-only instance of Bridge, bound to a specific deployed contract.
func NewBridgeCaller(address common.Address, caller bind.ContractCaller) (*BridgeCaller, error) {
	contract, err := bindBridge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeCaller{contract: contract}, nil
}

// NewBridgeTransactor creates a new write-only instance of Bridge, bound to a specific deployed contract.
func NewBridgeTransactor(address common.Address, transactor bind.ContractTransactor) (*BridgeTransactor, error) {
	contract, err := bindBridge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeTransactor{contract: contract}, nil
}

// NewBridgeFilterer creates a new log filterer instance of Bridge, bound to a specific deployed contract.
func NewBridgeFilterer(address common.Address, filterer bind.ContractFilterer) (*BridgeFilterer, error) {
	contract, err := bindBridge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BridgeFilterer{contract: contract}, nil
}

// bindBridge binds a generic wrapper to an already deployed contract.
func bindBridge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BridgeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bridge *BridgeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bridge.Contract.BridgeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bridge *BridgeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridge.Contract.BridgeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bridge *BridgeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bridge.Contract.BridgeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bridge *BridgeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bridge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bridge *BridgeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bridge *BridgeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bridge.Contract.contract.Transact(opts, method, params...)
}

// BlockDurationMillis is a free data retrieval call binding the contract method 0x9d5bc9e1.
//
// Solidity: function blockDurationMillis() view returns(uint64)
func (_Bridge *BridgeCaller) BlockDurationMillis(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "blockDurationMillis")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// BlockDurationMillis is a free data retrieval call binding the contract method 0x9d5bc9e1.
//
// Solidity: function blockDurationMillis() view returns(uint64)
func (_Bridge *BridgeSession) BlockDurationMillis() (uint64, error) {
	return _Bridge.Contract.BlockDurationMillis(&_Bridge.CallOpts)
}

// BlockDurationMillis is a free data retrieval call binding the contract method 0x9d5bc9e1.
//
// Solidity: function blockDurationMillis() view returns(uint64)
func (_Bridge *BridgeCallerSession) BlockDurationMillis() (uint64, error) {
	return _Bridge.Contract.BlockDurationMillis(&_Bridge.CallOpts)
}

// ColdValidatorSetHash is a free data retrieval call binding the contract method 0x0f711438.
//
// Solidity: function coldValidatorSetHash() view returns(bytes32)
func (_Bridge *BridgeCaller) ColdValidatorSetHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "coldValidatorSetHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ColdValidatorSetHash is a free data retrieval call binding the contract method 0x0f711438.
//
// Solidity: function coldValidatorSetHash() view returns(bytes32)
func (_Bridge *BridgeSession) ColdValidatorSetHash() ([32]byte, error) {
	return _Bridge.Contract.ColdValidatorSetHash(&_Bridge.CallOpts)
}

// ColdValidatorSetHash is a free data retrieval call binding the contract method 0x0f711438.
//
// Solidity: function coldValidatorSetHash() view returns(bytes32)
func (_Bridge *BridgeCallerSession) ColdValidatorSetHash() ([32]byte, error) {
	return _Bridge.Contract.ColdValidatorSetHash(&_Bridge.CallOpts)
}

// DisputePeriodSeconds is a free data retrieval call binding the contract method 0x0756183b.
//
// Solidity: function disputePeriodSeconds() view returns(uint64)
func (_Bridge *BridgeCaller) DisputePeriodSeconds(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "disputePeriodSeconds")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// DisputePeriodSeconds is a free data retrieval call binding the contract method 0x0756183b.
//
// Solidity: function disputePeriodSeconds() view returns(uint64)
func (_Bridge *BridgeSession) DisputePeriodSeconds() (uint64, error) {
	return _Bridge.Contract.DisputePeriodSeconds(&_Bridge.CallOpts)
}

// DisputePeriodSeconds is a free data retrieval call binding the contract method 0x0756183b.
//
// Solidity: function disputePeriodSeconds() view returns(uint64)
func (_Bridge *BridgeCallerSession) DisputePeriodSeconds() (uint64, error) {
	return _Bridge.Contract.DisputePeriodSeconds(&_Bridge.CallOpts)
}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_Bridge *BridgeCaller) DomainSeparator(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "domainSeparator")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_Bridge *BridgeSession) DomainSeparator() ([32]byte, error) {
	return _Bridge.Contract.DomainSeparator(&_Bridge.CallOpts)
}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_Bridge *BridgeCallerSession) DomainSeparator() ([32]byte, error) {
	return _Bridge.Contract.DomainSeparator(&_Bridge.CallOpts)
}

// Epoch is a free data retrieval call binding the contract method 0x900cf0cf.
//
// Solidity: function epoch() view returns(uint64)
func (_Bridge *BridgeCaller) Epoch(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "epoch")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// Epoch is a free data retrieval call binding the contract method 0x900cf0cf.
//
// Solidity: function epoch() view returns(uint64)
func (_Bridge *BridgeSession) Epoch() (uint64, error) {
	return _Bridge.Contract.Epoch(&_Bridge.CallOpts)
}

// Epoch is a free data retrieval call binding the contract method 0x900cf0cf.
//
// Solidity: function epoch() view returns(uint64)
func (_Bridge *BridgeCallerSession) Epoch() (uint64, error) {
	return _Bridge.Contract.Epoch(&_Bridge.CallOpts)
}

// FinalizedWithdrawals is a free data retrieval call binding the contract method 0xa14238e7.
//
// Solidity: function finalizedWithdrawals(bytes32 ) view returns(bool)
func (_Bridge *BridgeCaller) FinalizedWithdrawals(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "finalizedWithdrawals", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// FinalizedWithdrawals is a free data retrieval call binding the contract method 0xa14238e7.
//
// Solidity: function finalizedWithdrawals(bytes32 ) view returns(bool)
func (_Bridge *BridgeSession) FinalizedWithdrawals(arg0 [32]byte) (bool, error) {
	return _Bridge.Contract.FinalizedWithdrawals(&_Bridge.CallOpts, arg0)
}

// FinalizedWithdrawals is a free data retrieval call binding the contract method 0xa14238e7.
//
// Solidity: function finalizedWithdrawals(bytes32 ) view returns(bool)
func (_Bridge *BridgeCallerSession) FinalizedWithdrawals(arg0 [32]byte) (bool, error) {
	return _Bridge.Contract.FinalizedWithdrawals(&_Bridge.CallOpts, arg0)
}

// Finalizers is a free data retrieval call binding the contract method 0xcea75eb7.
//
// Solidity: function finalizers(address ) view returns(bool)
func (_Bridge *BridgeCaller) Finalizers(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "finalizers", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Finalizers is a free data retrieval call binding the contract method 0xcea75eb7.
//
// Solidity: function finalizers(address ) view returns(bool)
func (_Bridge *BridgeSession) Finalizers(arg0 common.Address) (bool, error) {
	return _Bridge.Contract.Finalizers(&_Bridge.CallOpts, arg0)
}

// Finalizers is a free data retrieval call binding the contract method 0xcea75eb7.
//
// Solidity: function finalizers(address ) view returns(bool)
func (_Bridge *BridgeCallerSession) Finalizers(arg0 common.Address) (bool, error) {
	return _Bridge.Contract.Finalizers(&_Bridge.CallOpts, arg0)
}

// GetLockersVotingLock is a free data retrieval call binding the contract method 0x53f79ef4.
//
// Solidity: function getLockersVotingLock() view returns(address[])
func (_Bridge *BridgeCaller) GetLockersVotingLock(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "getLockersVotingLock")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetLockersVotingLock is a free data retrieval call binding the contract method 0x53f79ef4.
//
// Solidity: function getLockersVotingLock() view returns(address[])
func (_Bridge *BridgeSession) GetLockersVotingLock() ([]common.Address, error) {
	return _Bridge.Contract.GetLockersVotingLock(&_Bridge.CallOpts)
}

// GetLockersVotingLock is a free data retrieval call binding the contract method 0x53f79ef4.
//
// Solidity: function getLockersVotingLock() view returns(address[])
func (_Bridge *BridgeCallerSession) GetLockersVotingLock() ([]common.Address, error) {
	return _Bridge.Contract.GetLockersVotingLock(&_Bridge.CallOpts)
}

// HotValidatorSetHash is a free data retrieval call binding the contract method 0xb0801e54.
//
// Solidity: function hotValidatorSetHash() view returns(bytes32)
func (_Bridge *BridgeCaller) HotValidatorSetHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "hotValidatorSetHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// HotValidatorSetHash is a free data retrieval call binding the contract method 0xb0801e54.
//
// Solidity: function hotValidatorSetHash() view returns(bytes32)
func (_Bridge *BridgeSession) HotValidatorSetHash() ([32]byte, error) {
	return _Bridge.Contract.HotValidatorSetHash(&_Bridge.CallOpts)
}

// HotValidatorSetHash is a free data retrieval call binding the contract method 0xb0801e54.
//
// Solidity: function hotValidatorSetHash() view returns(bytes32)
func (_Bridge *BridgeCallerSession) HotValidatorSetHash() ([32]byte, error) {
	return _Bridge.Contract.HotValidatorSetHash(&_Bridge.CallOpts)
}

// IsVotingLock is a free data retrieval call binding the contract method 0x3a37326e.
//
// Solidity: function isVotingLock(address locker) view returns(bool)
func (_Bridge *BridgeCaller) IsVotingLock(opts *bind.CallOpts, locker common.Address) (bool, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "isVotingLock", locker)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsVotingLock is a free data retrieval call binding the contract method 0x3a37326e.
//
// Solidity: function isVotingLock(address locker) view returns(bool)
func (_Bridge *BridgeSession) IsVotingLock(locker common.Address) (bool, error) {
	return _Bridge.Contract.IsVotingLock(&_Bridge.CallOpts, locker)
}

// IsVotingLock is a free data retrieval call binding the contract method 0x3a37326e.
//
// Solidity: function isVotingLock(address locker) view returns(bool)
func (_Bridge *BridgeCallerSession) IsVotingLock(locker common.Address) (bool, error) {
	return _Bridge.Contract.IsVotingLock(&_Bridge.CallOpts, locker)
}

// LockerThreshold is a free data retrieval call binding the contract method 0x05355e23.
//
// Solidity: function lockerThreshold() view returns(uint64)
func (_Bridge *BridgeCaller) LockerThreshold(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "lockerThreshold")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// LockerThreshold is a free data retrieval call binding the contract method 0x05355e23.
//
// Solidity: function lockerThreshold() view returns(uint64)
func (_Bridge *BridgeSession) LockerThreshold() (uint64, error) {
	return _Bridge.Contract.LockerThreshold(&_Bridge.CallOpts)
}

// LockerThreshold is a free data retrieval call binding the contract method 0x05355e23.
//
// Solidity: function lockerThreshold() view returns(uint64)
func (_Bridge *BridgeCallerSession) LockerThreshold() (uint64, error) {
	return _Bridge.Contract.LockerThreshold(&_Bridge.CallOpts)
}

// Lockers is a free data retrieval call binding the contract method 0x2c8e7a21.
//
// Solidity: function lockers(address ) view returns(bool)
func (_Bridge *BridgeCaller) Lockers(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "lockers", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Lockers is a free data retrieval call binding the contract method 0x2c8e7a21.
//
// Solidity: function lockers(address ) view returns(bool)
func (_Bridge *BridgeSession) Lockers(arg0 common.Address) (bool, error) {
	return _Bridge.Contract.Lockers(&_Bridge.CallOpts, arg0)
}

// Lockers is a free data retrieval call binding the contract method 0x2c8e7a21.
//
// Solidity: function lockers(address ) view returns(bool)
func (_Bridge *BridgeCallerSession) Lockers(arg0 common.Address) (bool, error) {
	return _Bridge.Contract.Lockers(&_Bridge.CallOpts, arg0)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Bridge *BridgeCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Bridge *BridgeSession) Paused() (bool, error) {
	return _Bridge.Contract.Paused(&_Bridge.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Bridge *BridgeCallerSession) Paused() (bool, error) {
	return _Bridge.Contract.Paused(&_Bridge.CallOpts)
}

// PendingValidatorSetUpdate is a free data retrieval call binding the contract method 0xc10ee9ae.
//
// Solidity: function pendingValidatorSetUpdate() view returns(uint64 epoch, uint64 totalValidatorPower, uint64 updateTime, uint64 updateBlockNumber, uint64 nValidators, bytes32 hotValidatorSetHash, bytes32 coldValidatorSetHash)
func (_Bridge *BridgeCaller) PendingValidatorSetUpdate(opts *bind.CallOpts) (struct {
	Epoch                uint64
	TotalValidatorPower  uint64
	UpdateTime           uint64
	UpdateBlockNumber    uint64
	NValidators          uint64
	HotValidatorSetHash  [32]byte
	ColdValidatorSetHash [32]byte
}, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "pendingValidatorSetUpdate")

	outstruct := new(struct {
		Epoch                uint64
		TotalValidatorPower  uint64
		UpdateTime           uint64
		UpdateBlockNumber    uint64
		NValidators          uint64
		HotValidatorSetHash  [32]byte
		ColdValidatorSetHash [32]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Epoch = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.TotalValidatorPower = *abi.ConvertType(out[1], new(uint64)).(*uint64)
	outstruct.UpdateTime = *abi.ConvertType(out[2], new(uint64)).(*uint64)
	outstruct.UpdateBlockNumber = *abi.ConvertType(out[3], new(uint64)).(*uint64)
	outstruct.NValidators = *abi.ConvertType(out[4], new(uint64)).(*uint64)
	outstruct.HotValidatorSetHash = *abi.ConvertType(out[5], new([32]byte)).(*[32]byte)
	outstruct.ColdValidatorSetHash = *abi.ConvertType(out[6], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

// PendingValidatorSetUpdate is a free data retrieval call binding the contract method 0xc10ee9ae.
//
// Solidity: function pendingValidatorSetUpdate() view returns(uint64 epoch, uint64 totalValidatorPower, uint64 updateTime, uint64 updateBlockNumber, uint64 nValidators, bytes32 hotValidatorSetHash, bytes32 coldValidatorSetHash)
func (_Bridge *BridgeSession) PendingValidatorSetUpdate() (struct {
	Epoch                uint64
	TotalValidatorPower  uint64
	UpdateTime           uint64
	UpdateBlockNumber    uint64
	NValidators          uint64
	HotValidatorSetHash  [32]byte
	ColdValidatorSetHash [32]byte
}, error) {
	return _Bridge.Contract.PendingValidatorSetUpdate(&_Bridge.CallOpts)
}

// PendingValidatorSetUpdate is a free data retrieval call binding the contract method 0xc10ee9ae.
//
// Solidity: function pendingValidatorSetUpdate() view returns(uint64 epoch, uint64 totalValidatorPower, uint64 updateTime, uint64 updateBlockNumber, uint64 nValidators, bytes32 hotValidatorSetHash, bytes32 coldValidatorSetHash)
func (_Bridge *BridgeCallerSession) PendingValidatorSetUpdate() (struct {
	Epoch                uint64
	TotalValidatorPower  uint64
	UpdateTime           uint64
	UpdateBlockNumber    uint64
	NValidators          uint64
	HotValidatorSetHash  [32]byte
	ColdValidatorSetHash [32]byte
}, error) {
	return _Bridge.Contract.PendingValidatorSetUpdate(&_Bridge.CallOpts)
}

// RequestedWithdrawals is a free data retrieval call binding the contract method 0x7694c6fa.
//
// Solidity: function requestedWithdrawals(bytes32 ) view returns(address user, address destination, address token, uint256 amount, uint256 chainId, uint64 nonce, uint64 requestedTime, uint64 requestedBlockNumber, bytes32 message)
func (_Bridge *BridgeCaller) RequestedWithdrawals(opts *bind.CallOpts, arg0 [32]byte) (struct {
	User                 common.Address
	Destination          common.Address
	Token                common.Address
	Amount               *big.Int
	ChainId              *big.Int
	Nonce                uint64
	RequestedTime        uint64
	RequestedBlockNumber uint64
	Message              [32]byte
}, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "requestedWithdrawals", arg0)

	outstruct := new(struct {
		User                 common.Address
		Destination          common.Address
		Token                common.Address
		Amount               *big.Int
		ChainId              *big.Int
		Nonce                uint64
		RequestedTime        uint64
		RequestedBlockNumber uint64
		Message              [32]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.User = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Destination = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Token = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.Amount = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.ChainId = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.Nonce = *abi.ConvertType(out[5], new(uint64)).(*uint64)
	outstruct.RequestedTime = *abi.ConvertType(out[6], new(uint64)).(*uint64)
	outstruct.RequestedBlockNumber = *abi.ConvertType(out[7], new(uint64)).(*uint64)
	outstruct.Message = *abi.ConvertType(out[8], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

// RequestedWithdrawals is a free data retrieval call binding the contract method 0x7694c6fa.
//
// Solidity: function requestedWithdrawals(bytes32 ) view returns(address user, address destination, address token, uint256 amount, uint256 chainId, uint64 nonce, uint64 requestedTime, uint64 requestedBlockNumber, bytes32 message)
func (_Bridge *BridgeSession) RequestedWithdrawals(arg0 [32]byte) (struct {
	User                 common.Address
	Destination          common.Address
	Token                common.Address
	Amount               *big.Int
	ChainId              *big.Int
	Nonce                uint64
	RequestedTime        uint64
	RequestedBlockNumber uint64
	Message              [32]byte
}, error) {
	return _Bridge.Contract.RequestedWithdrawals(&_Bridge.CallOpts, arg0)
}

// RequestedWithdrawals is a free data retrieval call binding the contract method 0x7694c6fa.
//
// Solidity: function requestedWithdrawals(bytes32 ) view returns(address user, address destination, address token, uint256 amount, uint256 chainId, uint64 nonce, uint64 requestedTime, uint64 requestedBlockNumber, bytes32 message)
func (_Bridge *BridgeCallerSession) RequestedWithdrawals(arg0 [32]byte) (struct {
	User                 common.Address
	Destination          common.Address
	Token                common.Address
	Amount               *big.Int
	ChainId              *big.Int
	Nonce                uint64
	RequestedTime        uint64
	RequestedBlockNumber uint64
	Message              [32]byte
}, error) {
	return _Bridge.Contract.RequestedWithdrawals(&_Bridge.CallOpts, arg0)
}

// TotalValidatorPower is a free data retrieval call binding the contract method 0xf8156a6e.
//
// Solidity: function totalValidatorPower() view returns(uint64)
func (_Bridge *BridgeCaller) TotalValidatorPower(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "totalValidatorPower")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// TotalValidatorPower is a free data retrieval call binding the contract method 0xf8156a6e.
//
// Solidity: function totalValidatorPower() view returns(uint64)
func (_Bridge *BridgeSession) TotalValidatorPower() (uint64, error) {
	return _Bridge.Contract.TotalValidatorPower(&_Bridge.CallOpts)
}

// TotalValidatorPower is a free data retrieval call binding the contract method 0xf8156a6e.
//
// Solidity: function totalValidatorPower() view returns(uint64)
func (_Bridge *BridgeCallerSession) TotalValidatorPower() (uint64, error) {
	return _Bridge.Contract.TotalValidatorPower(&_Bridge.CallOpts)
}

// UsedMessages is a free data retrieval call binding the contract method 0x5a028400.
//
// Solidity: function usedMessages(bytes32 ) view returns(bool)
func (_Bridge *BridgeCaller) UsedMessages(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "usedMessages", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// UsedMessages is a free data retrieval call binding the contract method 0x5a028400.
//
// Solidity: function usedMessages(bytes32 ) view returns(bool)
func (_Bridge *BridgeSession) UsedMessages(arg0 [32]byte) (bool, error) {
	return _Bridge.Contract.UsedMessages(&_Bridge.CallOpts, arg0)
}

// UsedMessages is a free data retrieval call binding the contract method 0x5a028400.
//
// Solidity: function usedMessages(bytes32 ) view returns(bool)
func (_Bridge *BridgeCallerSession) UsedMessages(arg0 [32]byte) (bool, error) {
	return _Bridge.Contract.UsedMessages(&_Bridge.CallOpts, arg0)
}

// WithdrawalsInvalidated is a free data retrieval call binding the contract method 0x42082828.
//
// Solidity: function withdrawalsInvalidated(bytes32 ) view returns(bool)
func (_Bridge *BridgeCaller) WithdrawalsInvalidated(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "withdrawalsInvalidated", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// WithdrawalsInvalidated is a free data retrieval call binding the contract method 0x42082828.
//
// Solidity: function withdrawalsInvalidated(bytes32 ) view returns(bool)
func (_Bridge *BridgeSession) WithdrawalsInvalidated(arg0 [32]byte) (bool, error) {
	return _Bridge.Contract.WithdrawalsInvalidated(&_Bridge.CallOpts, arg0)
}

// WithdrawalsInvalidated is a free data retrieval call binding the contract method 0x42082828.
//
// Solidity: function withdrawalsInvalidated(bytes32 ) view returns(bool)
func (_Bridge *BridgeCallerSession) WithdrawalsInvalidated(arg0 [32]byte) (bool, error) {
	return _Bridge.Contract.WithdrawalsInvalidated(&_Bridge.CallOpts, arg0)
}

// BatchedDepositWithPermit is a paid mutator transaction binding the contract method 0x6c9fc7b2.
//
// Solidity: function batchedDepositWithPermit((address,address,address,uint256,uint64,(uint256,uint256,uint8))[] deposits) returns()
func (_Bridge *BridgeTransactor) BatchedDepositWithPermit(opts *bind.TransactOpts, deposits []DepositWithPermit) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "batchedDepositWithPermit", deposits)
}

// BatchedDepositWithPermit is a paid mutator transaction binding the contract method 0x6c9fc7b2.
//
// Solidity: function batchedDepositWithPermit((address,address,address,uint256,uint64,(uint256,uint256,uint8))[] deposits) returns()
func (_Bridge *BridgeSession) BatchedDepositWithPermit(deposits []DepositWithPermit) (*types.Transaction, error) {
	return _Bridge.Contract.BatchedDepositWithPermit(&_Bridge.TransactOpts, deposits)
}

// BatchedDepositWithPermit is a paid mutator transaction binding the contract method 0x6c9fc7b2.
//
// Solidity: function batchedDepositWithPermit((address,address,address,uint256,uint64,(uint256,uint256,uint8))[] deposits) returns()
func (_Bridge *BridgeTransactorSession) BatchedDepositWithPermit(deposits []DepositWithPermit) (*types.Transaction, error) {
	return _Bridge.Contract.BatchedDepositWithPermit(&_Bridge.TransactOpts, deposits)
}

// BatchedFinalizeWithdrawals is a paid mutator transaction binding the contract method 0xc5bdf3ca.
//
// Solidity: function batchedFinalizeWithdrawals(bytes32[] messages) returns()
func (_Bridge *BridgeTransactor) BatchedFinalizeWithdrawals(opts *bind.TransactOpts, messages [][32]byte) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "batchedFinalizeWithdrawals", messages)
}

// BatchedFinalizeWithdrawals is a paid mutator transaction binding the contract method 0xc5bdf3ca.
//
// Solidity: function batchedFinalizeWithdrawals(bytes32[] messages) returns()
func (_Bridge *BridgeSession) BatchedFinalizeWithdrawals(messages [][32]byte) (*types.Transaction, error) {
	return _Bridge.Contract.BatchedFinalizeWithdrawals(&_Bridge.TransactOpts, messages)
}

// BatchedFinalizeWithdrawals is a paid mutator transaction binding the contract method 0xc5bdf3ca.
//
// Solidity: function batchedFinalizeWithdrawals(bytes32[] messages) returns()
func (_Bridge *BridgeTransactorSession) BatchedFinalizeWithdrawals(messages [][32]byte) (*types.Transaction, error) {
	return _Bridge.Contract.BatchedFinalizeWithdrawals(&_Bridge.TransactOpts, messages)
}

// BatchedRequestWithdrawals is a paid mutator transaction binding the contract method 0xb5f389f3.
//
// Solidity: function batchedRequestWithdrawals((address,address,address,uint256,uint256,uint64,(uint256,uint256,uint8)[])[] withdrawalRequests, (uint64,address[],uint64[]) hotValidatorSet) returns()
func (_Bridge *BridgeTransactor) BatchedRequestWithdrawals(opts *bind.TransactOpts, withdrawalRequests []WithdrawalRequest, hotValidatorSet ValidatorSet) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "batchedRequestWithdrawals", withdrawalRequests, hotValidatorSet)
}

// BatchedRequestWithdrawals is a paid mutator transaction binding the contract method 0xb5f389f3.
//
// Solidity: function batchedRequestWithdrawals((address,address,address,uint256,uint256,uint64,(uint256,uint256,uint8)[])[] withdrawalRequests, (uint64,address[],uint64[]) hotValidatorSet) returns()
func (_Bridge *BridgeSession) BatchedRequestWithdrawals(withdrawalRequests []WithdrawalRequest, hotValidatorSet ValidatorSet) (*types.Transaction, error) {
	return _Bridge.Contract.BatchedRequestWithdrawals(&_Bridge.TransactOpts, withdrawalRequests, hotValidatorSet)
}

// BatchedRequestWithdrawals is a paid mutator transaction binding the contract method 0xb5f389f3.
//
// Solidity: function batchedRequestWithdrawals((address,address,address,uint256,uint256,uint64,(uint256,uint256,uint8)[])[] withdrawalRequests, (uint64,address[],uint64[]) hotValidatorSet) returns()
func (_Bridge *BridgeTransactorSession) BatchedRequestWithdrawals(withdrawalRequests []WithdrawalRequest, hotValidatorSet ValidatorSet) (*types.Transaction, error) {
	return _Bridge.Contract.BatchedRequestWithdrawals(&_Bridge.TransactOpts, withdrawalRequests, hotValidatorSet)
}

// ChangeBlockDurationMillis is a paid mutator transaction binding the contract method 0x4aad6210.
//
// Solidity: function changeBlockDurationMillis(uint64 newBlockDurationMillis, uint64 nonce, (uint64,address[],uint64[]) activeColdValidatorSet, (uint256,uint256,uint8)[] signatures) returns()
func (_Bridge *BridgeTransactor) ChangeBlockDurationMillis(opts *bind.TransactOpts, newBlockDurationMillis uint64, nonce uint64, activeColdValidatorSet ValidatorSet, signatures []Signature) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "changeBlockDurationMillis", newBlockDurationMillis, nonce, activeColdValidatorSet, signatures)
}

// ChangeBlockDurationMillis is a paid mutator transaction binding the contract method 0x4aad6210.
//
// Solidity: function changeBlockDurationMillis(uint64 newBlockDurationMillis, uint64 nonce, (uint64,address[],uint64[]) activeColdValidatorSet, (uint256,uint256,uint8)[] signatures) returns()
func (_Bridge *BridgeSession) ChangeBlockDurationMillis(newBlockDurationMillis uint64, nonce uint64, activeColdValidatorSet ValidatorSet, signatures []Signature) (*types.Transaction, error) {
	return _Bridge.Contract.ChangeBlockDurationMillis(&_Bridge.TransactOpts, newBlockDurationMillis, nonce, activeColdValidatorSet, signatures)
}

// ChangeBlockDurationMillis is a paid mutator transaction binding the contract method 0x4aad6210.
//
// Solidity: function changeBlockDurationMillis(uint64 newBlockDurationMillis, uint64 nonce, (uint64,address[],uint64[]) activeColdValidatorSet, (uint256,uint256,uint8)[] signatures) returns()
func (_Bridge *BridgeTransactorSession) ChangeBlockDurationMillis(newBlockDurationMillis uint64, nonce uint64, activeColdValidatorSet ValidatorSet, signatures []Signature) (*types.Transaction, error) {
	return _Bridge.Contract.ChangeBlockDurationMillis(&_Bridge.TransactOpts, newBlockDurationMillis, nonce, activeColdValidatorSet, signatures)
}

// ChangeDisputePeriodSeconds is a paid mutator transaction binding the contract method 0x91ed1344.
//
// Solidity: function changeDisputePeriodSeconds(uint64 newDisputePeriodSeconds, uint64 nonce, (uint64,address[],uint64[]) activeColdValidatorSet, (uint256,uint256,uint8)[] signatures) returns()
func (_Bridge *BridgeTransactor) ChangeDisputePeriodSeconds(opts *bind.TransactOpts, newDisputePeriodSeconds uint64, nonce uint64, activeColdValidatorSet ValidatorSet, signatures []Signature) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "changeDisputePeriodSeconds", newDisputePeriodSeconds, nonce, activeColdValidatorSet, signatures)
}

// ChangeDisputePeriodSeconds is a paid mutator transaction binding the contract method 0x91ed1344.
//
// Solidity: function changeDisputePeriodSeconds(uint64 newDisputePeriodSeconds, uint64 nonce, (uint64,address[],uint64[]) activeColdValidatorSet, (uint256,uint256,uint8)[] signatures) returns()
func (_Bridge *BridgeSession) ChangeDisputePeriodSeconds(newDisputePeriodSeconds uint64, nonce uint64, activeColdValidatorSet ValidatorSet, signatures []Signature) (*types.Transaction, error) {
	return _Bridge.Contract.ChangeDisputePeriodSeconds(&_Bridge.TransactOpts, newDisputePeriodSeconds, nonce, activeColdValidatorSet, signatures)
}

// ChangeDisputePeriodSeconds is a paid mutator transaction binding the contract method 0x91ed1344.
//
// Solidity: function changeDisputePeriodSeconds(uint64 newDisputePeriodSeconds, uint64 nonce, (uint64,address[],uint64[]) activeColdValidatorSet, (uint256,uint256,uint8)[] signatures) returns()
func (_Bridge *BridgeTransactorSession) ChangeDisputePeriodSeconds(newDisputePeriodSeconds uint64, nonce uint64, activeColdValidatorSet ValidatorSet, signatures []Signature) (*types.Transaction, error) {
	return _Bridge.Contract.ChangeDisputePeriodSeconds(&_Bridge.TransactOpts, newDisputePeriodSeconds, nonce, activeColdValidatorSet, signatures)
}

// ChangeLockerThreshold is a paid mutator transaction binding the contract method 0xfc3f7ad3.
//
// Solidity: function changeLockerThreshold(uint64 newLockerThreshold, uint64 nonce, (uint64,address[],uint64[]) activeColdValidatorSet, (uint256,uint256,uint8)[] signatures) returns()
func (_Bridge *BridgeTransactor) ChangeLockerThreshold(opts *bind.TransactOpts, newLockerThreshold uint64, nonce uint64, activeColdValidatorSet ValidatorSet, signatures []Signature) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "changeLockerThreshold", newLockerThreshold, nonce, activeColdValidatorSet, signatures)
}

// ChangeLockerThreshold is a paid mutator transaction binding the contract method 0xfc3f7ad3.
//
// Solidity: function changeLockerThreshold(uint64 newLockerThreshold, uint64 nonce, (uint64,address[],uint64[]) activeColdValidatorSet, (uint256,uint256,uint8)[] signatures) returns()
func (_Bridge *BridgeSession) ChangeLockerThreshold(newLockerThreshold uint64, nonce uint64, activeColdValidatorSet ValidatorSet, signatures []Signature) (*types.Transaction, error) {
	return _Bridge.Contract.ChangeLockerThreshold(&_Bridge.TransactOpts, newLockerThreshold, nonce, activeColdValidatorSet, signatures)
}

// ChangeLockerThreshold is a paid mutator transaction binding the contract method 0xfc3f7ad3.
//
// Solidity: function changeLockerThreshold(uint64 newLockerThreshold, uint64 nonce, (uint64,address[],uint64[]) activeColdValidatorSet, (uint256,uint256,uint8)[] signatures) returns()
func (_Bridge *BridgeTransactorSession) ChangeLockerThreshold(newLockerThreshold uint64, nonce uint64, activeColdValidatorSet ValidatorSet, signatures []Signature) (*types.Transaction, error) {
	return _Bridge.Contract.ChangeLockerThreshold(&_Bridge.TransactOpts, newLockerThreshold, nonce, activeColdValidatorSet, signatures)
}

// Deposit is a paid mutator transaction binding the contract method 0x8340f549.
//
// Solidity: function deposit(address destination, address token, uint256 amount) returns()
func (_Bridge *BridgeTransactor) Deposit(opts *bind.TransactOpts, destination common.Address, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "deposit", destination, token, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x8340f549.
//
// Solidity: function deposit(address destination, address token, uint256 amount) returns()
func (_Bridge *BridgeSession) Deposit(destination common.Address, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.Deposit(&_Bridge.TransactOpts, destination, token, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x8340f549.
//
// Solidity: function deposit(address destination, address token, uint256 amount) returns()
func (_Bridge *BridgeTransactorSession) Deposit(destination common.Address, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.Deposit(&_Bridge.TransactOpts, destination, token, amount)
}

// EmergencyUnlock is a paid mutator transaction binding the contract method 0x9770e2c8.
//
// Solidity: function emergencyUnlock((uint64,address[],address[],uint64[]) newValidatorSet, (uint64,address[],uint64[]) activeColdValidatorSet, (uint256,uint256,uint8)[] signatures, uint64 nonce) returns()
func (_Bridge *BridgeTransactor) EmergencyUnlock(opts *bind.TransactOpts, newValidatorSet ValidatorSetUpdateRequest, activeColdValidatorSet ValidatorSet, signatures []Signature, nonce uint64) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "emergencyUnlock", newValidatorSet, activeColdValidatorSet, signatures, nonce)
}

// EmergencyUnlock is a paid mutator transaction binding the contract method 0x9770e2c8.
//
// Solidity: function emergencyUnlock((uint64,address[],address[],uint64[]) newValidatorSet, (uint64,address[],uint64[]) activeColdValidatorSet, (uint256,uint256,uint8)[] signatures, uint64 nonce) returns()
func (_Bridge *BridgeSession) EmergencyUnlock(newValidatorSet ValidatorSetUpdateRequest, activeColdValidatorSet ValidatorSet, signatures []Signature, nonce uint64) (*types.Transaction, error) {
	return _Bridge.Contract.EmergencyUnlock(&_Bridge.TransactOpts, newValidatorSet, activeColdValidatorSet, signatures, nonce)
}

// EmergencyUnlock is a paid mutator transaction binding the contract method 0x9770e2c8.
//
// Solidity: function emergencyUnlock((uint64,address[],address[],uint64[]) newValidatorSet, (uint64,address[],uint64[]) activeColdValidatorSet, (uint256,uint256,uint8)[] signatures, uint64 nonce) returns()
func (_Bridge *BridgeTransactorSession) EmergencyUnlock(newValidatorSet ValidatorSetUpdateRequest, activeColdValidatorSet ValidatorSet, signatures []Signature, nonce uint64) (*types.Transaction, error) {
	return _Bridge.Contract.EmergencyUnlock(&_Bridge.TransactOpts, newValidatorSet, activeColdValidatorSet, signatures, nonce)
}

// FinalizeValidatorSetUpdate is a paid mutator transaction binding the contract method 0x058731e5.
//
// Solidity: function finalizeValidatorSetUpdate() returns()
func (_Bridge *BridgeTransactor) FinalizeValidatorSetUpdate(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "finalizeValidatorSetUpdate")
}

// FinalizeValidatorSetUpdate is a paid mutator transaction binding the contract method 0x058731e5.
//
// Solidity: function finalizeValidatorSetUpdate() returns()
func (_Bridge *BridgeSession) FinalizeValidatorSetUpdate() (*types.Transaction, error) {
	return _Bridge.Contract.FinalizeValidatorSetUpdate(&_Bridge.TransactOpts)
}

// FinalizeValidatorSetUpdate is a paid mutator transaction binding the contract method 0x058731e5.
//
// Solidity: function finalizeValidatorSetUpdate() returns()
func (_Bridge *BridgeTransactorSession) FinalizeValidatorSetUpdate() (*types.Transaction, error) {
	return _Bridge.Contract.FinalizeValidatorSetUpdate(&_Bridge.TransactOpts)
}

// InvalidateWithdrawals is a paid mutator transaction binding the contract method 0x0fb61a2e.
//
// Solidity: function invalidateWithdrawals(bytes32[] messages, uint64 nonce, (uint64,address[],uint64[]) activeColdValidatorSet, (uint256,uint256,uint8)[] signatures) returns()
func (_Bridge *BridgeTransactor) InvalidateWithdrawals(opts *bind.TransactOpts, messages [][32]byte, nonce uint64, activeColdValidatorSet ValidatorSet, signatures []Signature) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "invalidateWithdrawals", messages, nonce, activeColdValidatorSet, signatures)
}

// InvalidateWithdrawals is a paid mutator transaction binding the contract method 0x0fb61a2e.
//
// Solidity: function invalidateWithdrawals(bytes32[] messages, uint64 nonce, (uint64,address[],uint64[]) activeColdValidatorSet, (uint256,uint256,uint8)[] signatures) returns()
func (_Bridge *BridgeSession) InvalidateWithdrawals(messages [][32]byte, nonce uint64, activeColdValidatorSet ValidatorSet, signatures []Signature) (*types.Transaction, error) {
	return _Bridge.Contract.InvalidateWithdrawals(&_Bridge.TransactOpts, messages, nonce, activeColdValidatorSet, signatures)
}

// InvalidateWithdrawals is a paid mutator transaction binding the contract method 0x0fb61a2e.
//
// Solidity: function invalidateWithdrawals(bytes32[] messages, uint64 nonce, (uint64,address[],uint64[]) activeColdValidatorSet, (uint256,uint256,uint8)[] signatures) returns()
func (_Bridge *BridgeTransactorSession) InvalidateWithdrawals(messages [][32]byte, nonce uint64, activeColdValidatorSet ValidatorSet, signatures []Signature) (*types.Transaction, error) {
	return _Bridge.Contract.InvalidateWithdrawals(&_Bridge.TransactOpts, messages, nonce, activeColdValidatorSet, signatures)
}

// ModifyFinalizer is a paid mutator transaction binding the contract method 0xe73ea41e.
//
// Solidity: function modifyFinalizer(address finalizer, bool _isFinalizer, uint64 nonce, (uint64,address[],uint64[]) activeValidatorSet, (uint256,uint256,uint8)[] signatures) returns()
func (_Bridge *BridgeTransactor) ModifyFinalizer(opts *bind.TransactOpts, finalizer common.Address, _isFinalizer bool, nonce uint64, activeValidatorSet ValidatorSet, signatures []Signature) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "modifyFinalizer", finalizer, _isFinalizer, nonce, activeValidatorSet, signatures)
}

// ModifyFinalizer is a paid mutator transaction binding the contract method 0xe73ea41e.
//
// Solidity: function modifyFinalizer(address finalizer, bool _isFinalizer, uint64 nonce, (uint64,address[],uint64[]) activeValidatorSet, (uint256,uint256,uint8)[] signatures) returns()
func (_Bridge *BridgeSession) ModifyFinalizer(finalizer common.Address, _isFinalizer bool, nonce uint64, activeValidatorSet ValidatorSet, signatures []Signature) (*types.Transaction, error) {
	return _Bridge.Contract.ModifyFinalizer(&_Bridge.TransactOpts, finalizer, _isFinalizer, nonce, activeValidatorSet, signatures)
}

// ModifyFinalizer is a paid mutator transaction binding the contract method 0xe73ea41e.
//
// Solidity: function modifyFinalizer(address finalizer, bool _isFinalizer, uint64 nonce, (uint64,address[],uint64[]) activeValidatorSet, (uint256,uint256,uint8)[] signatures) returns()
func (_Bridge *BridgeTransactorSession) ModifyFinalizer(finalizer common.Address, _isFinalizer bool, nonce uint64, activeValidatorSet ValidatorSet, signatures []Signature) (*types.Transaction, error) {
	return _Bridge.Contract.ModifyFinalizer(&_Bridge.TransactOpts, finalizer, _isFinalizer, nonce, activeValidatorSet, signatures)
}

// ModifyLocker is a paid mutator transaction binding the contract method 0x180f2e8c.
//
// Solidity: function modifyLocker(address locker, bool _isLocker, uint64 nonce, (uint64,address[],uint64[]) activeValidatorSet, (uint256,uint256,uint8)[] signatures) returns()
func (_Bridge *BridgeTransactor) ModifyLocker(opts *bind.TransactOpts, locker common.Address, _isLocker bool, nonce uint64, activeValidatorSet ValidatorSet, signatures []Signature) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "modifyLocker", locker, _isLocker, nonce, activeValidatorSet, signatures)
}

// ModifyLocker is a paid mutator transaction binding the contract method 0x180f2e8c.
//
// Solidity: function modifyLocker(address locker, bool _isLocker, uint64 nonce, (uint64,address[],uint64[]) activeValidatorSet, (uint256,uint256,uint8)[] signatures) returns()
func (_Bridge *BridgeSession) ModifyLocker(locker common.Address, _isLocker bool, nonce uint64, activeValidatorSet ValidatorSet, signatures []Signature) (*types.Transaction, error) {
	return _Bridge.Contract.ModifyLocker(&_Bridge.TransactOpts, locker, _isLocker, nonce, activeValidatorSet, signatures)
}

// ModifyLocker is a paid mutator transaction binding the contract method 0x180f2e8c.
//
// Solidity: function modifyLocker(address locker, bool _isLocker, uint64 nonce, (uint64,address[],uint64[]) activeValidatorSet, (uint256,uint256,uint8)[] signatures) returns()
func (_Bridge *BridgeTransactorSession) ModifyLocker(locker common.Address, _isLocker bool, nonce uint64, activeValidatorSet ValidatorSet, signatures []Signature) (*types.Transaction, error) {
	return _Bridge.Contract.ModifyLocker(&_Bridge.TransactOpts, locker, _isLocker, nonce, activeValidatorSet, signatures)
}

// UnvoteEmergencyLock is a paid mutator transaction binding the contract method 0xb091049c.
//
// Solidity: function unvoteEmergencyLock() returns()
func (_Bridge *BridgeTransactor) UnvoteEmergencyLock(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "unvoteEmergencyLock")
}

// UnvoteEmergencyLock is a paid mutator transaction binding the contract method 0xb091049c.
//
// Solidity: function unvoteEmergencyLock() returns()
func (_Bridge *BridgeSession) UnvoteEmergencyLock() (*types.Transaction, error) {
	return _Bridge.Contract.UnvoteEmergencyLock(&_Bridge.TransactOpts)
}

// UnvoteEmergencyLock is a paid mutator transaction binding the contract method 0xb091049c.
//
// Solidity: function unvoteEmergencyLock() returns()
func (_Bridge *BridgeTransactorSession) UnvoteEmergencyLock() (*types.Transaction, error) {
	return _Bridge.Contract.UnvoteEmergencyLock(&_Bridge.TransactOpts)
}

// UpdateValidatorSet is a paid mutator transaction binding the contract method 0xe3e6c441.
//
// Solidity: function updateValidatorSet((uint64,address[],address[],uint64[]) newValidatorSet, (uint64,address[],uint64[]) activeHotValidatorSet, (uint256,uint256,uint8)[] signatures) returns()
func (_Bridge *BridgeTransactor) UpdateValidatorSet(opts *bind.TransactOpts, newValidatorSet ValidatorSetUpdateRequest, activeHotValidatorSet ValidatorSet, signatures []Signature) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "updateValidatorSet", newValidatorSet, activeHotValidatorSet, signatures)
}

// UpdateValidatorSet is a paid mutator transaction binding the contract method 0xe3e6c441.
//
// Solidity: function updateValidatorSet((uint64,address[],address[],uint64[]) newValidatorSet, (uint64,address[],uint64[]) activeHotValidatorSet, (uint256,uint256,uint8)[] signatures) returns()
func (_Bridge *BridgeSession) UpdateValidatorSet(newValidatorSet ValidatorSetUpdateRequest, activeHotValidatorSet ValidatorSet, signatures []Signature) (*types.Transaction, error) {
	return _Bridge.Contract.UpdateValidatorSet(&_Bridge.TransactOpts, newValidatorSet, activeHotValidatorSet, signatures)
}

// UpdateValidatorSet is a paid mutator transaction binding the contract method 0xe3e6c441.
//
// Solidity: function updateValidatorSet((uint64,address[],address[],uint64[]) newValidatorSet, (uint64,address[],uint64[]) activeHotValidatorSet, (uint256,uint256,uint8)[] signatures) returns()
func (_Bridge *BridgeTransactorSession) UpdateValidatorSet(newValidatorSet ValidatorSetUpdateRequest, activeHotValidatorSet ValidatorSet, signatures []Signature) (*types.Transaction, error) {
	return _Bridge.Contract.UpdateValidatorSet(&_Bridge.TransactOpts, newValidatorSet, activeHotValidatorSet, signatures)
}

// VoteEmergencyLock is a paid mutator transaction binding the contract method 0x4878ee53.
//
// Solidity: function voteEmergencyLock() returns()
func (_Bridge *BridgeTransactor) VoteEmergencyLock(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "voteEmergencyLock")
}

// VoteEmergencyLock is a paid mutator transaction binding the contract method 0x4878ee53.
//
// Solidity: function voteEmergencyLock() returns()
func (_Bridge *BridgeSession) VoteEmergencyLock() (*types.Transaction, error) {
	return _Bridge.Contract.VoteEmergencyLock(&_Bridge.TransactOpts)
}

// VoteEmergencyLock is a paid mutator transaction binding the contract method 0x4878ee53.
//
// Solidity: function voteEmergencyLock() returns()
func (_Bridge *BridgeTransactorSession) VoteEmergencyLock() (*types.Transaction, error) {
	return _Bridge.Contract.VoteEmergencyLock(&_Bridge.TransactOpts)
}

// BridgeChangedBlockDurationMillisIterator is returned from FilterChangedBlockDurationMillis and is used to iterate over the raw logs and unpacked data for ChangedBlockDurationMillis events raised by the Bridge contract.
type BridgeChangedBlockDurationMillisIterator struct {
	Event *BridgeChangedBlockDurationMillis // Event containing the contract specifics and raw log

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
func (it *BridgeChangedBlockDurationMillisIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeChangedBlockDurationMillis)
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
		it.Event = new(BridgeChangedBlockDurationMillis)
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
func (it *BridgeChangedBlockDurationMillisIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeChangedBlockDurationMillisIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeChangedBlockDurationMillis represents a ChangedBlockDurationMillis event raised by the Bridge contract.
type BridgeChangedBlockDurationMillis struct {
	NewBlockDurationMillis uint64
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterChangedBlockDurationMillis is a free log retrieval operation binding the contract event 0x0ef2da393c3832a8f08ce447e14948d21e84f864facf7327137387bd0596a563.
//
// Solidity: event ChangedBlockDurationMillis(uint64 newBlockDurationMillis)
func (_Bridge *BridgeFilterer) FilterChangedBlockDurationMillis(opts *bind.FilterOpts) (*BridgeChangedBlockDurationMillisIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "ChangedBlockDurationMillis")
	if err != nil {
		return nil, err
	}
	return &BridgeChangedBlockDurationMillisIterator{contract: _Bridge.contract, event: "ChangedBlockDurationMillis", logs: logs, sub: sub}, nil
}

// WatchChangedBlockDurationMillis is a free log subscription operation binding the contract event 0x0ef2da393c3832a8f08ce447e14948d21e84f864facf7327137387bd0596a563.
//
// Solidity: event ChangedBlockDurationMillis(uint64 newBlockDurationMillis)
func (_Bridge *BridgeFilterer) WatchChangedBlockDurationMillis(opts *bind.WatchOpts, sink chan<- *BridgeChangedBlockDurationMillis) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "ChangedBlockDurationMillis")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeChangedBlockDurationMillis)
				if err := _Bridge.contract.UnpackLog(event, "ChangedBlockDurationMillis", log); err != nil {
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

// ParseChangedBlockDurationMillis is a log parse operation binding the contract event 0x0ef2da393c3832a8f08ce447e14948d21e84f864facf7327137387bd0596a563.
//
// Solidity: event ChangedBlockDurationMillis(uint64 newBlockDurationMillis)
func (_Bridge *BridgeFilterer) ParseChangedBlockDurationMillis(log types.Log) (*BridgeChangedBlockDurationMillis, error) {
	event := new(BridgeChangedBlockDurationMillis)
	if err := _Bridge.contract.UnpackLog(event, "ChangedBlockDurationMillis", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeChangedDisputePeriodSecondsIterator is returned from FilterChangedDisputePeriodSeconds and is used to iterate over the raw logs and unpacked data for ChangedDisputePeriodSeconds events raised by the Bridge contract.
type BridgeChangedDisputePeriodSecondsIterator struct {
	Event *BridgeChangedDisputePeriodSeconds // Event containing the contract specifics and raw log

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
func (it *BridgeChangedDisputePeriodSecondsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeChangedDisputePeriodSeconds)
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
		it.Event = new(BridgeChangedDisputePeriodSeconds)
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
func (it *BridgeChangedDisputePeriodSecondsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeChangedDisputePeriodSecondsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeChangedDisputePeriodSeconds represents a ChangedDisputePeriodSeconds event raised by the Bridge contract.
type BridgeChangedDisputePeriodSeconds struct {
	NewDisputePeriodSeconds uint64
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterChangedDisputePeriodSeconds is a free log retrieval operation binding the contract event 0x04edaf680108675f58d2ea70e9e7886c39ed38b66439622f8362d36595fe8169.
//
// Solidity: event ChangedDisputePeriodSeconds(uint64 newDisputePeriodSeconds)
func (_Bridge *BridgeFilterer) FilterChangedDisputePeriodSeconds(opts *bind.FilterOpts) (*BridgeChangedDisputePeriodSecondsIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "ChangedDisputePeriodSeconds")
	if err != nil {
		return nil, err
	}
	return &BridgeChangedDisputePeriodSecondsIterator{contract: _Bridge.contract, event: "ChangedDisputePeriodSeconds", logs: logs, sub: sub}, nil
}

// WatchChangedDisputePeriodSeconds is a free log subscription operation binding the contract event 0x04edaf680108675f58d2ea70e9e7886c39ed38b66439622f8362d36595fe8169.
//
// Solidity: event ChangedDisputePeriodSeconds(uint64 newDisputePeriodSeconds)
func (_Bridge *BridgeFilterer) WatchChangedDisputePeriodSeconds(opts *bind.WatchOpts, sink chan<- *BridgeChangedDisputePeriodSeconds) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "ChangedDisputePeriodSeconds")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeChangedDisputePeriodSeconds)
				if err := _Bridge.contract.UnpackLog(event, "ChangedDisputePeriodSeconds", log); err != nil {
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

// ParseChangedDisputePeriodSeconds is a log parse operation binding the contract event 0x04edaf680108675f58d2ea70e9e7886c39ed38b66439622f8362d36595fe8169.
//
// Solidity: event ChangedDisputePeriodSeconds(uint64 newDisputePeriodSeconds)
func (_Bridge *BridgeFilterer) ParseChangedDisputePeriodSeconds(log types.Log) (*BridgeChangedDisputePeriodSeconds, error) {
	event := new(BridgeChangedDisputePeriodSeconds)
	if err := _Bridge.contract.UnpackLog(event, "ChangedDisputePeriodSeconds", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeChangedLockerThresholdIterator is returned from FilterChangedLockerThreshold and is used to iterate over the raw logs and unpacked data for ChangedLockerThreshold events raised by the Bridge contract.
type BridgeChangedLockerThresholdIterator struct {
	Event *BridgeChangedLockerThreshold // Event containing the contract specifics and raw log

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
func (it *BridgeChangedLockerThresholdIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeChangedLockerThreshold)
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
		it.Event = new(BridgeChangedLockerThreshold)
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
func (it *BridgeChangedLockerThresholdIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeChangedLockerThresholdIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeChangedLockerThreshold represents a ChangedLockerThreshold event raised by the Bridge contract.
type BridgeChangedLockerThreshold struct {
	NewLockerThreshold uint64
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterChangedLockerThreshold is a free log retrieval operation binding the contract event 0x2dbe453726b24b2cee427a7d6e2dcc9f353f16bee104f3d21480157a0ee409f7.
//
// Solidity: event ChangedLockerThreshold(uint64 newLockerThreshold)
func (_Bridge *BridgeFilterer) FilterChangedLockerThreshold(opts *bind.FilterOpts) (*BridgeChangedLockerThresholdIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "ChangedLockerThreshold")
	if err != nil {
		return nil, err
	}
	return &BridgeChangedLockerThresholdIterator{contract: _Bridge.contract, event: "ChangedLockerThreshold", logs: logs, sub: sub}, nil
}

// WatchChangedLockerThreshold is a free log subscription operation binding the contract event 0x2dbe453726b24b2cee427a7d6e2dcc9f353f16bee104f3d21480157a0ee409f7.
//
// Solidity: event ChangedLockerThreshold(uint64 newLockerThreshold)
func (_Bridge *BridgeFilterer) WatchChangedLockerThreshold(opts *bind.WatchOpts, sink chan<- *BridgeChangedLockerThreshold) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "ChangedLockerThreshold")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeChangedLockerThreshold)
				if err := _Bridge.contract.UnpackLog(event, "ChangedLockerThreshold", log); err != nil {
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

// ParseChangedLockerThreshold is a log parse operation binding the contract event 0x2dbe453726b24b2cee427a7d6e2dcc9f353f16bee104f3d21480157a0ee409f7.
//
// Solidity: event ChangedLockerThreshold(uint64 newLockerThreshold)
func (_Bridge *BridgeFilterer) ParseChangedLockerThreshold(log types.Log) (*BridgeChangedLockerThreshold, error) {
	event := new(BridgeChangedLockerThreshold)
	if err := _Bridge.contract.UnpackLog(event, "ChangedLockerThreshold", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the Bridge contract.
type BridgeDepositIterator struct {
	Event *BridgeDeposit // Event containing the contract specifics and raw log

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
func (it *BridgeDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeDeposit)
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
		it.Event = new(BridgeDeposit)
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
func (it *BridgeDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeDeposit represents a Deposit event raised by the Bridge contract.
type BridgeDeposit struct {
	User        common.Address
	Destination common.Address
	Token       common.Address
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0x7cfff908a4b583f36430b25d75964c458d8ede8a99bd61be750e97ee1b2f3a96.
//
// Solidity: event Deposit(address indexed user, address destination, address token, uint256 amount)
func (_Bridge *BridgeFilterer) FilterDeposit(opts *bind.FilterOpts, user []common.Address) (*BridgeDepositIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "Deposit", userRule)
	if err != nil {
		return nil, err
	}
	return &BridgeDepositIterator{contract: _Bridge.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0x7cfff908a4b583f36430b25d75964c458d8ede8a99bd61be750e97ee1b2f3a96.
//
// Solidity: event Deposit(address indexed user, address destination, address token, uint256 amount)
func (_Bridge *BridgeFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *BridgeDeposit, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "Deposit", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeDeposit)
				if err := _Bridge.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0x7cfff908a4b583f36430b25d75964c458d8ede8a99bd61be750e97ee1b2f3a96.
//
// Solidity: event Deposit(address indexed user, address destination, address token, uint256 amount)
func (_Bridge *BridgeFilterer) ParseDeposit(log types.Log) (*BridgeDeposit, error) {
	event := new(BridgeDeposit)
	if err := _Bridge.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeFailedWithdrawalIterator is returned from FilterFailedWithdrawal and is used to iterate over the raw logs and unpacked data for FailedWithdrawal events raised by the Bridge contract.
type BridgeFailedWithdrawalIterator struct {
	Event *BridgeFailedWithdrawal // Event containing the contract specifics and raw log

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
func (it *BridgeFailedWithdrawalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeFailedWithdrawal)
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
		it.Event = new(BridgeFailedWithdrawal)
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
func (it *BridgeFailedWithdrawalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeFailedWithdrawalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeFailedWithdrawal represents a FailedWithdrawal event raised by the Bridge contract.
type BridgeFailedWithdrawal struct {
	Message   [32]byte
	ErrorCode uint32
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterFailedWithdrawal is a free log retrieval operation binding the contract event 0x686cb4bac974cd11b0f8a75fc7c7764ed12cc46faaec53110f807aa802a7acb4.
//
// Solidity: event FailedWithdrawal(bytes32 message, uint32 errorCode)
func (_Bridge *BridgeFilterer) FilterFailedWithdrawal(opts *bind.FilterOpts) (*BridgeFailedWithdrawalIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "FailedWithdrawal")
	if err != nil {
		return nil, err
	}
	return &BridgeFailedWithdrawalIterator{contract: _Bridge.contract, event: "FailedWithdrawal", logs: logs, sub: sub}, nil
}

// WatchFailedWithdrawal is a free log subscription operation binding the contract event 0x686cb4bac974cd11b0f8a75fc7c7764ed12cc46faaec53110f807aa802a7acb4.
//
// Solidity: event FailedWithdrawal(bytes32 message, uint32 errorCode)
func (_Bridge *BridgeFilterer) WatchFailedWithdrawal(opts *bind.WatchOpts, sink chan<- *BridgeFailedWithdrawal) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "FailedWithdrawal")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeFailedWithdrawal)
				if err := _Bridge.contract.UnpackLog(event, "FailedWithdrawal", log); err != nil {
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

// ParseFailedWithdrawal is a log parse operation binding the contract event 0x686cb4bac974cd11b0f8a75fc7c7764ed12cc46faaec53110f807aa802a7acb4.
//
// Solidity: event FailedWithdrawal(bytes32 message, uint32 errorCode)
func (_Bridge *BridgeFilterer) ParseFailedWithdrawal(log types.Log) (*BridgeFailedWithdrawal, error) {
	event := new(BridgeFailedWithdrawal)
	if err := _Bridge.contract.UnpackLog(event, "FailedWithdrawal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeFinalizedValidatorSetUpdateIterator is returned from FilterFinalizedValidatorSetUpdate and is used to iterate over the raw logs and unpacked data for FinalizedValidatorSetUpdate events raised by the Bridge contract.
type BridgeFinalizedValidatorSetUpdateIterator struct {
	Event *BridgeFinalizedValidatorSetUpdate // Event containing the contract specifics and raw log

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
func (it *BridgeFinalizedValidatorSetUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeFinalizedValidatorSetUpdate)
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
		it.Event = new(BridgeFinalizedValidatorSetUpdate)
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
func (it *BridgeFinalizedValidatorSetUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeFinalizedValidatorSetUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeFinalizedValidatorSetUpdate represents a FinalizedValidatorSetUpdate event raised by the Bridge contract.
type BridgeFinalizedValidatorSetUpdate struct {
	Epoch                uint64
	HotValidatorSetHash  [32]byte
	ColdValidatorSetHash [32]byte
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterFinalizedValidatorSetUpdate is a free log retrieval operation binding the contract event 0x87da17ff65d815d1e1c369cb3bbda9a11af181b92dc52681a2779419781c6270.
//
// Solidity: event FinalizedValidatorSetUpdate(uint64 epoch, bytes32 hotValidatorSetHash, bytes32 coldValidatorSetHash)
func (_Bridge *BridgeFilterer) FilterFinalizedValidatorSetUpdate(opts *bind.FilterOpts) (*BridgeFinalizedValidatorSetUpdateIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "FinalizedValidatorSetUpdate")
	if err != nil {
		return nil, err
	}
	return &BridgeFinalizedValidatorSetUpdateIterator{contract: _Bridge.contract, event: "FinalizedValidatorSetUpdate", logs: logs, sub: sub}, nil
}

// WatchFinalizedValidatorSetUpdate is a free log subscription operation binding the contract event 0x87da17ff65d815d1e1c369cb3bbda9a11af181b92dc52681a2779419781c6270.
//
// Solidity: event FinalizedValidatorSetUpdate(uint64 epoch, bytes32 hotValidatorSetHash, bytes32 coldValidatorSetHash)
func (_Bridge *BridgeFilterer) WatchFinalizedValidatorSetUpdate(opts *bind.WatchOpts, sink chan<- *BridgeFinalizedValidatorSetUpdate) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "FinalizedValidatorSetUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeFinalizedValidatorSetUpdate)
				if err := _Bridge.contract.UnpackLog(event, "FinalizedValidatorSetUpdate", log); err != nil {
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

// ParseFinalizedValidatorSetUpdate is a log parse operation binding the contract event 0x87da17ff65d815d1e1c369cb3bbda9a11af181b92dc52681a2779419781c6270.
//
// Solidity: event FinalizedValidatorSetUpdate(uint64 epoch, bytes32 hotValidatorSetHash, bytes32 coldValidatorSetHash)
func (_Bridge *BridgeFilterer) ParseFinalizedValidatorSetUpdate(log types.Log) (*BridgeFinalizedValidatorSetUpdate, error) {
	event := new(BridgeFinalizedValidatorSetUpdate)
	if err := _Bridge.contract.UnpackLog(event, "FinalizedValidatorSetUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeFinalizedWithdrawalIterator is returned from FilterFinalizedWithdrawal and is used to iterate over the raw logs and unpacked data for FinalizedWithdrawal events raised by the Bridge contract.
type BridgeFinalizedWithdrawalIterator struct {
	Event *BridgeFinalizedWithdrawal // Event containing the contract specifics and raw log

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
func (it *BridgeFinalizedWithdrawalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeFinalizedWithdrawal)
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
		it.Event = new(BridgeFinalizedWithdrawal)
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
func (it *BridgeFinalizedWithdrawalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeFinalizedWithdrawalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeFinalizedWithdrawal represents a FinalizedWithdrawal event raised by the Bridge contract.
type BridgeFinalizedWithdrawal struct {
	Message     [32]byte
	User        common.Address
	Destination common.Address
	Token       common.Address
	Amount      *big.Int
	Nonce       uint64
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterFinalizedWithdrawal is a free log retrieval operation binding the contract event 0x04cafa25a7826c4415eac718e45fc84f69b5539748e0206adbae70f919887548.
//
// Solidity: event FinalizedWithdrawal(bytes32 message, address indexed user, address destination, address token, uint256 amount, uint64 nonce)
func (_Bridge *BridgeFilterer) FilterFinalizedWithdrawal(opts *bind.FilterOpts, user []common.Address) (*BridgeFinalizedWithdrawalIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "FinalizedWithdrawal", userRule)
	if err != nil {
		return nil, err
	}
	return &BridgeFinalizedWithdrawalIterator{contract: _Bridge.contract, event: "FinalizedWithdrawal", logs: logs, sub: sub}, nil
}

// WatchFinalizedWithdrawal is a free log subscription operation binding the contract event 0x04cafa25a7826c4415eac718e45fc84f69b5539748e0206adbae70f919887548.
//
// Solidity: event FinalizedWithdrawal(bytes32 message, address indexed user, address destination, address token, uint256 amount, uint64 nonce)
func (_Bridge *BridgeFilterer) WatchFinalizedWithdrawal(opts *bind.WatchOpts, sink chan<- *BridgeFinalizedWithdrawal, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "FinalizedWithdrawal", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeFinalizedWithdrawal)
				if err := _Bridge.contract.UnpackLog(event, "FinalizedWithdrawal", log); err != nil {
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

// ParseFinalizedWithdrawal is a log parse operation binding the contract event 0x04cafa25a7826c4415eac718e45fc84f69b5539748e0206adbae70f919887548.
//
// Solidity: event FinalizedWithdrawal(bytes32 message, address indexed user, address destination, address token, uint256 amount, uint64 nonce)
func (_Bridge *BridgeFilterer) ParseFinalizedWithdrawal(log types.Log) (*BridgeFinalizedWithdrawal, error) {
	event := new(BridgeFinalizedWithdrawal)
	if err := _Bridge.contract.UnpackLog(event, "FinalizedWithdrawal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeInvalidatedWithdrawalIterator is returned from FilterInvalidatedWithdrawal and is used to iterate over the raw logs and unpacked data for InvalidatedWithdrawal events raised by the Bridge contract.
type BridgeInvalidatedWithdrawalIterator struct {
	Event *BridgeInvalidatedWithdrawal // Event containing the contract specifics and raw log

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
func (it *BridgeInvalidatedWithdrawalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeInvalidatedWithdrawal)
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
		it.Event = new(BridgeInvalidatedWithdrawal)
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
func (it *BridgeInvalidatedWithdrawalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeInvalidatedWithdrawalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeInvalidatedWithdrawal represents a InvalidatedWithdrawal event raised by the Bridge contract.
type BridgeInvalidatedWithdrawal struct {
	Withdrawal Withdrawal
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterInvalidatedWithdrawal is a free log retrieval operation binding the contract event 0x4e1a2aef00d7868e1f49c3784b1b802acad3fd7e4c7fe753694a51d9b46346c5.
//
// Solidity: event InvalidatedWithdrawal((address,address,address,uint256,uint256,uint64,uint64,uint64,bytes32) withdrawal)
func (_Bridge *BridgeFilterer) FilterInvalidatedWithdrawal(opts *bind.FilterOpts) (*BridgeInvalidatedWithdrawalIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "InvalidatedWithdrawal")
	if err != nil {
		return nil, err
	}
	return &BridgeInvalidatedWithdrawalIterator{contract: _Bridge.contract, event: "InvalidatedWithdrawal", logs: logs, sub: sub}, nil
}

// WatchInvalidatedWithdrawal is a free log subscription operation binding the contract event 0x4e1a2aef00d7868e1f49c3784b1b802acad3fd7e4c7fe753694a51d9b46346c5.
//
// Solidity: event InvalidatedWithdrawal((address,address,address,uint256,uint256,uint64,uint64,uint64,bytes32) withdrawal)
func (_Bridge *BridgeFilterer) WatchInvalidatedWithdrawal(opts *bind.WatchOpts, sink chan<- *BridgeInvalidatedWithdrawal) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "InvalidatedWithdrawal")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeInvalidatedWithdrawal)
				if err := _Bridge.contract.UnpackLog(event, "InvalidatedWithdrawal", log); err != nil {
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

// ParseInvalidatedWithdrawal is a log parse operation binding the contract event 0x4e1a2aef00d7868e1f49c3784b1b802acad3fd7e4c7fe753694a51d9b46346c5.
//
// Solidity: event InvalidatedWithdrawal((address,address,address,uint256,uint256,uint64,uint64,uint64,bytes32) withdrawal)
func (_Bridge *BridgeFilterer) ParseInvalidatedWithdrawal(log types.Log) (*BridgeInvalidatedWithdrawal, error) {
	event := new(BridgeInvalidatedWithdrawal)
	if err := _Bridge.contract.UnpackLog(event, "InvalidatedWithdrawal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeModifiedFinalizerIterator is returned from FilterModifiedFinalizer and is used to iterate over the raw logs and unpacked data for ModifiedFinalizer events raised by the Bridge contract.
type BridgeModifiedFinalizerIterator struct {
	Event *BridgeModifiedFinalizer // Event containing the contract specifics and raw log

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
func (it *BridgeModifiedFinalizerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeModifiedFinalizer)
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
		it.Event = new(BridgeModifiedFinalizer)
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
func (it *BridgeModifiedFinalizerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeModifiedFinalizerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeModifiedFinalizer represents a ModifiedFinalizer event raised by the Bridge contract.
type BridgeModifiedFinalizer struct {
	Finalizer   common.Address
	IsFinalizer bool
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterModifiedFinalizer is a free log retrieval operation binding the contract event 0x2526bb92d75e00cfad8c7c16cb75f3e1073c854339e49b16baaad3067c2ed65a.
//
// Solidity: event ModifiedFinalizer(address indexed finalizer, bool isFinalizer)
func (_Bridge *BridgeFilterer) FilterModifiedFinalizer(opts *bind.FilterOpts, finalizer []common.Address) (*BridgeModifiedFinalizerIterator, error) {

	var finalizerRule []interface{}
	for _, finalizerItem := range finalizer {
		finalizerRule = append(finalizerRule, finalizerItem)
	}

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "ModifiedFinalizer", finalizerRule)
	if err != nil {
		return nil, err
	}
	return &BridgeModifiedFinalizerIterator{contract: _Bridge.contract, event: "ModifiedFinalizer", logs: logs, sub: sub}, nil
}

// WatchModifiedFinalizer is a free log subscription operation binding the contract event 0x2526bb92d75e00cfad8c7c16cb75f3e1073c854339e49b16baaad3067c2ed65a.
//
// Solidity: event ModifiedFinalizer(address indexed finalizer, bool isFinalizer)
func (_Bridge *BridgeFilterer) WatchModifiedFinalizer(opts *bind.WatchOpts, sink chan<- *BridgeModifiedFinalizer, finalizer []common.Address) (event.Subscription, error) {

	var finalizerRule []interface{}
	for _, finalizerItem := range finalizer {
		finalizerRule = append(finalizerRule, finalizerItem)
	}

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "ModifiedFinalizer", finalizerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeModifiedFinalizer)
				if err := _Bridge.contract.UnpackLog(event, "ModifiedFinalizer", log); err != nil {
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

// ParseModifiedFinalizer is a log parse operation binding the contract event 0x2526bb92d75e00cfad8c7c16cb75f3e1073c854339e49b16baaad3067c2ed65a.
//
// Solidity: event ModifiedFinalizer(address indexed finalizer, bool isFinalizer)
func (_Bridge *BridgeFilterer) ParseModifiedFinalizer(log types.Log) (*BridgeModifiedFinalizer, error) {
	event := new(BridgeModifiedFinalizer)
	if err := _Bridge.contract.UnpackLog(event, "ModifiedFinalizer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeModifiedLockerIterator is returned from FilterModifiedLocker and is used to iterate over the raw logs and unpacked data for ModifiedLocker events raised by the Bridge contract.
type BridgeModifiedLockerIterator struct {
	Event *BridgeModifiedLocker // Event containing the contract specifics and raw log

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
func (it *BridgeModifiedLockerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeModifiedLocker)
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
		it.Event = new(BridgeModifiedLocker)
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
func (it *BridgeModifiedLockerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeModifiedLockerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeModifiedLocker represents a ModifiedLocker event raised by the Bridge contract.
type BridgeModifiedLocker struct {
	Locker   common.Address
	IsLocker bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterModifiedLocker is a free log retrieval operation binding the contract event 0x26690dc5c5a9d2aa7ac3efa2b7c515652e4621a3e075d267bcac51c16fb97532.
//
// Solidity: event ModifiedLocker(address indexed locker, bool isLocker)
func (_Bridge *BridgeFilterer) FilterModifiedLocker(opts *bind.FilterOpts, locker []common.Address) (*BridgeModifiedLockerIterator, error) {

	var lockerRule []interface{}
	for _, lockerItem := range locker {
		lockerRule = append(lockerRule, lockerItem)
	}

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "ModifiedLocker", lockerRule)
	if err != nil {
		return nil, err
	}
	return &BridgeModifiedLockerIterator{contract: _Bridge.contract, event: "ModifiedLocker", logs: logs, sub: sub}, nil
}

// WatchModifiedLocker is a free log subscription operation binding the contract event 0x26690dc5c5a9d2aa7ac3efa2b7c515652e4621a3e075d267bcac51c16fb97532.
//
// Solidity: event ModifiedLocker(address indexed locker, bool isLocker)
func (_Bridge *BridgeFilterer) WatchModifiedLocker(opts *bind.WatchOpts, sink chan<- *BridgeModifiedLocker, locker []common.Address) (event.Subscription, error) {

	var lockerRule []interface{}
	for _, lockerItem := range locker {
		lockerRule = append(lockerRule, lockerItem)
	}

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "ModifiedLocker", lockerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeModifiedLocker)
				if err := _Bridge.contract.UnpackLog(event, "ModifiedLocker", log); err != nil {
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

// ParseModifiedLocker is a log parse operation binding the contract event 0x26690dc5c5a9d2aa7ac3efa2b7c515652e4621a3e075d267bcac51c16fb97532.
//
// Solidity: event ModifiedLocker(address indexed locker, bool isLocker)
func (_Bridge *BridgeFilterer) ParseModifiedLocker(log types.Log) (*BridgeModifiedLocker, error) {
	event := new(BridgeModifiedLocker)
	if err := _Bridge.contract.UnpackLog(event, "ModifiedLocker", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgePausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Bridge contract.
type BridgePausedIterator struct {
	Event *BridgePaused // Event containing the contract specifics and raw log

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
func (it *BridgePausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgePaused)
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
		it.Event = new(BridgePaused)
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
func (it *BridgePausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgePausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgePaused represents a Paused event raised by the Bridge contract.
type BridgePaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Bridge *BridgeFilterer) FilterPaused(opts *bind.FilterOpts) (*BridgePausedIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &BridgePausedIterator{contract: _Bridge.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Bridge *BridgeFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *BridgePaused) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgePaused)
				if err := _Bridge.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_Bridge *BridgeFilterer) ParsePaused(log types.Log) (*BridgePaused, error) {
	event := new(BridgePaused)
	if err := _Bridge.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeRequestedValidatorSetUpdateIterator is returned from FilterRequestedValidatorSetUpdate and is used to iterate over the raw logs and unpacked data for RequestedValidatorSetUpdate events raised by the Bridge contract.
type BridgeRequestedValidatorSetUpdateIterator struct {
	Event *BridgeRequestedValidatorSetUpdate // Event containing the contract specifics and raw log

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
func (it *BridgeRequestedValidatorSetUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeRequestedValidatorSetUpdate)
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
		it.Event = new(BridgeRequestedValidatorSetUpdate)
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
func (it *BridgeRequestedValidatorSetUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeRequestedValidatorSetUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeRequestedValidatorSetUpdate represents a RequestedValidatorSetUpdate event raised by the Bridge contract.
type BridgeRequestedValidatorSetUpdate struct {
	Epoch                uint64
	HotValidatorSetHash  [32]byte
	ColdValidatorSetHash [32]byte
	UpdateTime           uint64
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterRequestedValidatorSetUpdate is a free log retrieval operation binding the contract event 0x420bbe99bd2c52ec500d33614359525f3ef7bb3358c0e07d1312db0941cbf2f4.
//
// Solidity: event RequestedValidatorSetUpdate(uint64 epoch, bytes32 hotValidatorSetHash, bytes32 coldValidatorSetHash, uint64 updateTime)
func (_Bridge *BridgeFilterer) FilterRequestedValidatorSetUpdate(opts *bind.FilterOpts) (*BridgeRequestedValidatorSetUpdateIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "RequestedValidatorSetUpdate")
	if err != nil {
		return nil, err
	}
	return &BridgeRequestedValidatorSetUpdateIterator{contract: _Bridge.contract, event: "RequestedValidatorSetUpdate", logs: logs, sub: sub}, nil
}

// WatchRequestedValidatorSetUpdate is a free log subscription operation binding the contract event 0x420bbe99bd2c52ec500d33614359525f3ef7bb3358c0e07d1312db0941cbf2f4.
//
// Solidity: event RequestedValidatorSetUpdate(uint64 epoch, bytes32 hotValidatorSetHash, bytes32 coldValidatorSetHash, uint64 updateTime)
func (_Bridge *BridgeFilterer) WatchRequestedValidatorSetUpdate(opts *bind.WatchOpts, sink chan<- *BridgeRequestedValidatorSetUpdate) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "RequestedValidatorSetUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeRequestedValidatorSetUpdate)
				if err := _Bridge.contract.UnpackLog(event, "RequestedValidatorSetUpdate", log); err != nil {
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

// ParseRequestedValidatorSetUpdate is a log parse operation binding the contract event 0x420bbe99bd2c52ec500d33614359525f3ef7bb3358c0e07d1312db0941cbf2f4.
//
// Solidity: event RequestedValidatorSetUpdate(uint64 epoch, bytes32 hotValidatorSetHash, bytes32 coldValidatorSetHash, uint64 updateTime)
func (_Bridge *BridgeFilterer) ParseRequestedValidatorSetUpdate(log types.Log) (*BridgeRequestedValidatorSetUpdate, error) {
	event := new(BridgeRequestedValidatorSetUpdate)
	if err := _Bridge.contract.UnpackLog(event, "RequestedValidatorSetUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeRequestedWithdrawalIterator is returned from FilterRequestedWithdrawal and is used to iterate over the raw logs and unpacked data for RequestedWithdrawal events raised by the Bridge contract.
type BridgeRequestedWithdrawalIterator struct {
	Event *BridgeRequestedWithdrawal // Event containing the contract specifics and raw log

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
func (it *BridgeRequestedWithdrawalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeRequestedWithdrawal)
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
		it.Event = new(BridgeRequestedWithdrawal)
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
func (it *BridgeRequestedWithdrawalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeRequestedWithdrawalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeRequestedWithdrawal represents a RequestedWithdrawal event raised by the Bridge contract.
type BridgeRequestedWithdrawal struct {
	Message              [32]byte
	User                 common.Address
	Destination          common.Address
	Token                common.Address
	Amount               *big.Int
	ChainId              *big.Int
	Nonce                uint64
	RequestedTime        uint64
	RequestedBlockNumber uint64
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterRequestedWithdrawal is a free log retrieval operation binding the contract event 0x6188a1db6e3e44fa683d03383b49bbe96178887d5c9d0174718f03ea6193b5ad.
//
// Solidity: event RequestedWithdrawal(bytes32 message, address indexed user, address destination, address token, uint256 amount, uint256 chainId, uint64 nonce, uint64 requestedTime, uint64 requestedBlockNumber)
func (_Bridge *BridgeFilterer) FilterRequestedWithdrawal(opts *bind.FilterOpts, user []common.Address) (*BridgeRequestedWithdrawalIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "RequestedWithdrawal", userRule)
	if err != nil {
		return nil, err
	}
	return &BridgeRequestedWithdrawalIterator{contract: _Bridge.contract, event: "RequestedWithdrawal", logs: logs, sub: sub}, nil
}

// WatchRequestedWithdrawal is a free log subscription operation binding the contract event 0x6188a1db6e3e44fa683d03383b49bbe96178887d5c9d0174718f03ea6193b5ad.
//
// Solidity: event RequestedWithdrawal(bytes32 message, address indexed user, address destination, address token, uint256 amount, uint256 chainId, uint64 nonce, uint64 requestedTime, uint64 requestedBlockNumber)
func (_Bridge *BridgeFilterer) WatchRequestedWithdrawal(opts *bind.WatchOpts, sink chan<- *BridgeRequestedWithdrawal, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "RequestedWithdrawal", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeRequestedWithdrawal)
				if err := _Bridge.contract.UnpackLog(event, "RequestedWithdrawal", log); err != nil {
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

// ParseRequestedWithdrawal is a log parse operation binding the contract event 0x6188a1db6e3e44fa683d03383b49bbe96178887d5c9d0174718f03ea6193b5ad.
//
// Solidity: event RequestedWithdrawal(bytes32 message, address indexed user, address destination, address token, uint256 amount, uint256 chainId, uint64 nonce, uint64 requestedTime, uint64 requestedBlockNumber)
func (_Bridge *BridgeFilterer) ParseRequestedWithdrawal(log types.Log) (*BridgeRequestedWithdrawal, error) {
	event := new(BridgeRequestedWithdrawal)
	if err := _Bridge.contract.UnpackLog(event, "RequestedWithdrawal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Bridge contract.
type BridgeUnpausedIterator struct {
	Event *BridgeUnpaused // Event containing the contract specifics and raw log

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
func (it *BridgeUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeUnpaused)
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
		it.Event = new(BridgeUnpaused)
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
func (it *BridgeUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeUnpaused represents a Unpaused event raised by the Bridge contract.
type BridgeUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Bridge *BridgeFilterer) FilterUnpaused(opts *bind.FilterOpts) (*BridgeUnpausedIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &BridgeUnpausedIterator{contract: _Bridge.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Bridge *BridgeFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *BridgeUnpaused) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeUnpaused)
				if err := _Bridge.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_Bridge *BridgeFilterer) ParseUnpaused(log types.Log) (*BridgeUnpaused, error) {
	event := new(BridgeUnpaused)
	if err := _Bridge.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
