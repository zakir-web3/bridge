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
	Destination          [32]byte
	Token                [32]byte
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
	Destination [32]byte
	Token       [32]byte
	Amount      *big.Int
	ChainId     *big.Int
	Nonce       uint64
	Signatures  []Signature
}

// BridgeMetaData contains all meta data concerning the Bridge contract.
var BridgeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"hotAddresses\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"coldAddresses\",\"type\":\"address[]\"},{\"internalType\":\"uint64[]\",\"name\":\"powers\",\"type\":\"uint64[]\"},{\"internalType\":\"uint64\",\"name\":\"_disputePeriodSeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_blockDurationMillis\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_lockerThreshold\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AddressInsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EnforcedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExpectedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedInnerCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"newBlockDurationMillis\",\"type\":\"uint64\"}],\"name\":\"ChangedBlockDurationMillis\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"newDisputePeriodSeconds\",\"type\":\"uint64\"}],\"name\":\"ChangedDisputePeriodSeconds\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"newLockerThreshold\",\"type\":\"uint64\"}],\"name\":\"ChangedLockerThreshold\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"destination\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"message\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"errorCode\",\"type\":\"uint32\"}],\"name\":\"FailedWithdrawal\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"epoch\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"hotValidatorSetHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"coldValidatorSetHash\",\"type\":\"bytes32\"}],\"name\":\"FinalizedValidatorSetUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"message\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"destination\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"token\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"name\":\"FinalizedWithdrawal\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"destination\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"token\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"requestedTime\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"requestedBlockNumber\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"message\",\"type\":\"bytes32\"}],\"indexed\":false,\"internalType\":\"structWithdrawal\",\"name\":\"withdrawal\",\"type\":\"tuple\"}],\"name\":\"InvalidatedWithdrawal\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"finalizer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isFinalizer\",\"type\":\"bool\"}],\"name\":\"ModifiedFinalizer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"locker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isLocker\",\"type\":\"bool\"}],\"name\":\"ModifiedLocker\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"epoch\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"hotValidatorSetHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"coldValidatorSetHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"updateTime\",\"type\":\"uint64\"}],\"name\":\"RequestedValidatorSetUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"message\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"destination\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"token\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"requestedTime\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"requestedBlockNumber\",\"type\":\"uint64\"}],\"name\":\"RequestedWithdrawal\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"destination\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"deadline\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"r\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"s\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"}],\"internalType\":\"structSignature\",\"name\":\"signature\",\"type\":\"tuple\"}],\"internalType\":\"structDepositWithPermit[]\",\"name\":\"deposits\",\"type\":\"tuple[]\"}],\"name\":\"batchedDepositWithPermit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"messages\",\"type\":\"bytes32[]\"}],\"name\":\"batchedFinalizeWithdrawals\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"destination\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"token\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"r\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"s\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"}],\"internalType\":\"structSignature[]\",\"name\":\"signatures\",\"type\":\"tuple[]\"}],\"internalType\":\"structWithdrawalRequest[]\",\"name\":\"withdrawalRequests\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"epoch\",\"type\":\"uint64\"},{\"internalType\":\"address[]\",\"name\":\"validators\",\"type\":\"address[]\"},{\"internalType\":\"uint64[]\",\"name\":\"powers\",\"type\":\"uint64[]\"}],\"internalType\":\"structValidatorSet\",\"name\":\"hotValidatorSet\",\"type\":\"tuple\"}],\"name\":\"batchedRequestWithdrawals\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"blockDurationMillis\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"newBlockDurationMillis\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"epoch\",\"type\":\"uint64\"},{\"internalType\":\"address[]\",\"name\":\"validators\",\"type\":\"address[]\"},{\"internalType\":\"uint64[]\",\"name\":\"powers\",\"type\":\"uint64[]\"}],\"internalType\":\"structValidatorSet\",\"name\":\"activeColdValidatorSet\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"r\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"s\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"}],\"internalType\":\"structSignature[]\",\"name\":\"signatures\",\"type\":\"tuple[]\"}],\"name\":\"changeBlockDurationMillis\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"newDisputePeriodSeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"epoch\",\"type\":\"uint64\"},{\"internalType\":\"address[]\",\"name\":\"validators\",\"type\":\"address[]\"},{\"internalType\":\"uint64[]\",\"name\":\"powers\",\"type\":\"uint64[]\"}],\"internalType\":\"structValidatorSet\",\"name\":\"activeColdValidatorSet\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"r\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"s\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"}],\"internalType\":\"structSignature[]\",\"name\":\"signatures\",\"type\":\"tuple[]\"}],\"name\":\"changeDisputePeriodSeconds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"newLockerThreshold\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"epoch\",\"type\":\"uint64\"},{\"internalType\":\"address[]\",\"name\":\"validators\",\"type\":\"address[]\"},{\"internalType\":\"uint64[]\",\"name\":\"powers\",\"type\":\"uint64[]\"}],\"internalType\":\"structValidatorSet\",\"name\":\"activeColdValidatorSet\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"r\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"s\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"}],\"internalType\":\"structSignature[]\",\"name\":\"signatures\",\"type\":\"tuple[]\"}],\"name\":\"changeLockerThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"coldValidatorSetHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"destination\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"disputePeriodSeconds\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"domainSeparator\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"epoch\",\"type\":\"uint64\"},{\"internalType\":\"address[]\",\"name\":\"hotAddresses\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"coldAddresses\",\"type\":\"address[]\"},{\"internalType\":\"uint64[]\",\"name\":\"powers\",\"type\":\"uint64[]\"}],\"internalType\":\"structValidatorSetUpdateRequest\",\"name\":\"newValidatorSet\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"epoch\",\"type\":\"uint64\"},{\"internalType\":\"address[]\",\"name\":\"validators\",\"type\":\"address[]\"},{\"internalType\":\"uint64[]\",\"name\":\"powers\",\"type\":\"uint64[]\"}],\"internalType\":\"structValidatorSet\",\"name\":\"activeColdValidatorSet\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"r\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"s\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"}],\"internalType\":\"structSignature[]\",\"name\":\"signatures\",\"type\":\"tuple[]\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"name\":\"emergencyUnlock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"epoch\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"finalizeValidatorSetUpdate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"finalizedWithdrawals\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"finalizers\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLockersVotingLock\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"hotValidatorSetHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"messages\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"epoch\",\"type\":\"uint64\"},{\"internalType\":\"address[]\",\"name\":\"validators\",\"type\":\"address[]\"},{\"internalType\":\"uint64[]\",\"name\":\"powers\",\"type\":\"uint64[]\"}],\"internalType\":\"structValidatorSet\",\"name\":\"activeColdValidatorSet\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"r\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"s\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"}],\"internalType\":\"structSignature[]\",\"name\":\"signatures\",\"type\":\"tuple[]\"}],\"name\":\"invalidateWithdrawals\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"locker\",\"type\":\"address\"}],\"name\":\"isVotingLock\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lockerThreshold\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"lockers\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"finalizer\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_isFinalizer\",\"type\":\"bool\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"epoch\",\"type\":\"uint64\"},{\"internalType\":\"address[]\",\"name\":\"validators\",\"type\":\"address[]\"},{\"internalType\":\"uint64[]\",\"name\":\"powers\",\"type\":\"uint64[]\"}],\"internalType\":\"structValidatorSet\",\"name\":\"activeValidatorSet\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"r\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"s\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"}],\"internalType\":\"structSignature[]\",\"name\":\"signatures\",\"type\":\"tuple[]\"}],\"name\":\"modifyFinalizer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"locker\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_isLocker\",\"type\":\"bool\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"epoch\",\"type\":\"uint64\"},{\"internalType\":\"address[]\",\"name\":\"validators\",\"type\":\"address[]\"},{\"internalType\":\"uint64[]\",\"name\":\"powers\",\"type\":\"uint64[]\"}],\"internalType\":\"structValidatorSet\",\"name\":\"activeValidatorSet\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"r\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"s\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"}],\"internalType\":\"structSignature[]\",\"name\":\"signatures\",\"type\":\"tuple[]\"}],\"name\":\"modifyLocker\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingValidatorSetUpdate\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"epoch\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"totalValidatorPower\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"updateTime\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"updateBlockNumber\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nValidators\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"hotValidatorSetHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"coldValidatorSetHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"requestedWithdrawals\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"destination\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"token\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"requestedTime\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"requestedBlockNumber\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"message\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalValidatorPower\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unvoteEmergencyLock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"epoch\",\"type\":\"uint64\"},{\"internalType\":\"address[]\",\"name\":\"hotAddresses\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"coldAddresses\",\"type\":\"address[]\"},{\"internalType\":\"uint64[]\",\"name\":\"powers\",\"type\":\"uint64[]\"}],\"internalType\":\"structValidatorSetUpdateRequest\",\"name\":\"newValidatorSet\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"epoch\",\"type\":\"uint64\"},{\"internalType\":\"address[]\",\"name\":\"validators\",\"type\":\"address[]\"},{\"internalType\":\"uint64[]\",\"name\":\"powers\",\"type\":\"uint64[]\"}],\"internalType\":\"structValidatorSet\",\"name\":\"activeHotValidatorSet\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"r\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"s\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"}],\"internalType\":\"structSignature[]\",\"name\":\"signatures\",\"type\":\"tuple[]\"}],\"name\":\"updateValidatorSet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"usedMessages\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"voteEmergencyLock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"withdrawalsInvalidated\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b5060405162004778380380620047788339810160408190526200003491620007ee565b6000805460ff1916905560018055604080517f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f60208201527f7aa5ae620294318af92bf4e2b2a729646c932a80312a5fa630da993a2ef5cc10918101919091527fc89efdaa54c0f20c7adf612882df0950f5a951637e0307cdcb4c672f298b8bc660608201524660808201523060a082015260c00160408051601f198184030181529190528051602090910120601155620000ef8462000424565b600d60086101000a8154816001600160401b0302191690836001600160401b031602179055508451865114620001805760405162461bcd60e51b815260206004820152602b60248201527f486f7420616e6420636f6c642076616c696461746f722073657473206c656e6760448201526a0e8d040dad2e6dac2e8c6d60ab1b60648201526084015b60405180910390fd5b620001ae604051806060016040528060006001600160401b0316815260200160608152602001606081525090565b506040805160608101825260008082526020820189905291810186905290620001d7826200050f565b6002819055604080516060808201835260008083526020808401839052928401829052835191820184528082529181018b9052918201899052919250906200021f826200050f565b6003819055600d80546001600160801b0316600160801b6001600160401b038b8116919091026001600160c01b031691909117600160c01b8a83160217909155600b80546001600160401b0319169188169190911790559050620002838a62000640565b600254600354604080516000815260208101939093528201526001600160401b03421660608201527f420bbe99bd2c52ec500d33614359525f3ef7bb3358c0e07d1312db0941cbf2f49060800160405180910390a16040805160e0810182526000808252600d546801000000000000000090046001600160401b031660208301529181019190915260608101620003174390565b6001600160401b0390811682528c51811660208084019190915260025460408085018290526003546060958601819052865160048054898701518a8601518b8b0151948a166001600160801b03199093169290921768010000000000000000918a1691909102176001600160801b0316600160801b918916919091026001600160c01b031617600160c01b928816929092029190911790556080870151600580546001600160401b031916919096161790945560a086015160065560c09095015160075584516000815291820152928301527f87da17ff65d815d1e1c369cb3bbda9a11af181b92dc52681a2779419781c6270910160405180910390a15050505050505050505062000a17565b60008060005b8351816001600160401b03161015620004875783816001600160401b0316815181106200045b576200045b6200091a565b60200260200101518262000470919062000946565b9150806200047e8162000970565b9150506200042a565b506000816001600160401b031611620005095760405162461bcd60e51b815260206004820152603460248201527f5375626d69747465642076616c696461746f7220706f77657273206d7573742060448201527f62652067726561746572207468616e207a65726f000000000000000000000000606482015260840162000177565b92915050565b6000816040015151826020015151146200056c5760405162461bcd60e51b815260206004820152601760248201527f4d616c666f726d65642076616c696461746f7220736574000000000000000000604482015260640162000177565b60007fcf7a991d34f65202b9a5ebe03e28c3fd6f86e1f75fabbddd532864507554c66783600001518460200151604051602001620005ab9190620009a1565b604051602081830303815290604052805190602001208560400151604051602001620005d89190620009e2565b604051602081830303815290604052805190602001206040516020016200062194939291909384526001600160401b039290921660208401526040830152606082015260800190565b60408051601f1981840301815291905280516020909101209392505050565b805160005b816001600160401b0316816001600160401b03161015620006cf57600083826001600160401b0316815181106200068057620006806200091a565b6020908102919091018101516001600160a01b031660009081526009825260408082208054600160ff199182168117909255600c90945291208054909216811790915591909101905062000645565b505050565b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f191681016001600160401b0381118282101715620007155762000715620006d4565b604052919050565b60006001600160401b03821115620007395762000739620006d4565b5060051b60200190565b600082601f8301126200075557600080fd5b815160206200076e62000768836200071d565b620006ea565b8083825260208201915060208460051b8701019350868411156200079157600080fd5b602086015b84811015620007c65780516001600160a01b0381168114620007b85760008081fd5b835291830191830162000796565b509695505050505050565b80516001600160401b0381168114620007e957600080fd5b919050565b60008060008060008060c087890312156200080857600080fd5b86516001600160401b03808211156200082057600080fd5b6200082e8a838b0162000743565b97506020915081890151818111156200084657600080fd5b620008548b828c0162000743565b9750506040890151818111156200086a57600080fd5b89019050601f81018a136200087e57600080fd5b80516200088f62000768826200071d565b81815260059190911b8201830190838101908c831115620008af57600080fd5b928401925b82841015620008d857620008c884620007d1565b82529284019290840190620008b4565b8098505050505050620008ee60608801620007d1565b9250620008fe60808801620007d1565b91506200090e60a08801620007d1565b90509295509295509295565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b6001600160401b0381811683821601908082111562000969576200096962000930565b5092915050565b60006001600160401b038281166002600160401b0319810162000997576200099762000930565b6001019392505050565b815160009082906020808601845b83811015620009d65781516001600160a01b031685529382019390820190600101620009af565b50929695505050505050565b815160009082906020808601845b83811015620009d65781516001600160401b031685529382019390820190600101620009f0565b613d518062000a276000396000f3fe608060405234801561001057600080fd5b50600436106101fb5760003560e01c80637694c6fa1161011a578063b091049c116100ad578063e3e6c4411161007c578063e3e6c441146105ae578063e73ea41e146105c1578063f698da25146105d4578063f8156a6e146105dd578063fc3f7ad3146105f757600080fd5b8063b091049c146104e2578063c10ee9ae146104ea578063c5bdf3ca14610578578063cea75eb71461058b57600080fd5b80639770e2c8116100e95780639770e2c8146104895780639d5bc9e11461049c578063a14238e7146104b6578063b0801e54146104d957600080fd5b80637694c6fa1461037e5780638340f54914610450578063900cf0cf1461046357806391ed13441461047657600080fd5b80633a37326e1161019257806353f79ef41161016157806353f79ef4146103285780635a0284001461033d5780635c975abb146103605780636c9fc7b21461036b57600080fd5b80633a37326e146102d757806342082828146102ea5780634878ee531461030d5780634aad62101461031557600080fd5b80630fb61a2e116101ce5780630fb61a2e1461026b578063180f2e8c1461027e57806324f0b6c7146102915780632c8e7a21146102a457600080fd5b806305355e2314610200578063058731e5146102305780630756183b1461023a5780630f71143814610254575b600080fd5b600b54610213906001600160401b031681565b6040516001600160401b0390911681526020015b60405180910390f35b61023861060a565b005b600d5461021390600160801b90046001600160401b031681565b61025d60035481565b604051908152602001610227565b610238610279366004613298565b610734565b61023861028c3660046133ae565b6108b7565b61023861029f366004613445565b610a02565b6102c76102b2366004613599565b60096020526000908152604090205460ff1681565b6040519015158152602001610227565b6102c76102e5366004613599565b610a9c565b6102c76102f83660046135b4565b60106020526000908152604090205460ff1681565b610238610b1a565b6102386103233660046135cd565b610c20565b610330610cfa565b6040516102279190613616565b6102c761034b3660046135b4565b60086020526000908152604090205460ff1681565b60005460ff166102c7565b610238610379366004613663565b610d5c565b6103f361038c3660046135b4565b600e6020526000908152604090208054600182015460028301546003840154600485015460058601546006909601546001600160a01b039095169593949293919290916001600160401b0380821692600160401b8304821692600160801b90049091169089565b604080516001600160a01b03909a168a5260208a019890985296880195909552606087019390935260808601919091526001600160401b0390811660a086015290811660c08501521660e083015261010082015261012001610227565b61023861045e366004613767565b610e93565b600d54610213906001600160401b031681565b6102386104843660046135cd565b610fdb565b610238610497366004613862565b6110af565b600d5461021390600160c01b90046001600160401b031681565b6102c76104c43660046135b4565b600f6020526000908152604090205460ff1681565b61025d60025481565b610238611244565b60045460055460065460075461052d936001600160401b0380821694600160401b8304821694600160801b8404831694600160c01b909404831693929091169187565b604080516001600160401b0398891681529688166020880152948716948601949094529185166060850152909316608083015260a082019290925260c081019190915260e001610227565b610238610586366004613931565b6112f3565b6102c7610599366004613599565b600c6020526000908152604090205460ff1681565b6102386105bc3660046139a5565b61135c565b6102386105cf3660046133ae565b6114f3565b61025d60115481565b600d5461021390600160401b90046001600160401b031681565b6102386106053660046135cd565b6115e7565b6106126116dd565b61061a611707565b6106233361172b565b600454600160801b90046001600160401b03166000036106a15760405162461bcd60e51b815260206004820152602e60248201527f50656e64696e672076616c696461746f72207365742075706461746520616c7260448201526d1958591e48199a5b985b1a5e995960921b60648201526084015b60405180910390fd5b6004546000906106ca906001600160401b03600160801b8204811691600160c01b900416611793565b905063ffffffff8116156107205760405162461bcd60e51b815260206004820152601760248201527f5374696c6c20696e206469737075746520706572696f640000000000000000006044820152606401610698565b61072861184c565b5061073260018055565b565b60007fa0675b98ae6c277eba9efba80fbbdcac49c582fec62ca7833048cb970321ad2e856040516020016107689190613a2c565b60408051601f198184030181528282528051602091820120908301939093528101919091526001600160401b03851660608201526080016040516020818303038152906040528051906020012090506107c0816118e8565b6107ce818484600354611959565b845160005b816001600160401b0316816001600160401b031610156108ae5760016010600089846001600160401b03168151811061080e5761080e613a62565b6020026020010151815260200190815260200160002060006101000a81548160ff0219169083151502179055507f5fe92156b011257af09c3fceff5d0f5ee2c781900477e1eadca143dfe07ab9fd600e600089846001600160401b03168151811061087b5761087b613a62565b6020026020010151815260200190815260200160002060405161089e9190613a78565b60405180910390a16001016107d3565b50505050505050565b604080517f2bb2a8361d9a37d6bd9173b7a98ef8abfb3224b8f0c01732fa686695b7973af060208201526001600160a01b0387169181019190915284151560608201526001600160401b038416608082015260009060a00160405160208183030381529060405280519060200120905060008515610938575060025461093d565b506003545b610946826118e8565b61095a8261095386613b0e565b8584611959565b6001600160a01b03871660009081526009602052604090205460ff168015610980575085155b801561098f575060005460ff16155b1561099d5761099d87611c3b565b6001600160a01b038716600081815260096020908152604091829020805460ff19168a151590811790915591519182527f26690dc5c5a9d2aa7ac3efa2b7c515652e4621a3e075d267bcac51c16fb9753291015b60405180910390a250505050505050565b610a0a6116dd565b610a12611707565b815160005b816001600160401b0316816001600160401b03161015610a8d57600084826001600160401b031681518110610a4e57610a4e613a62565b60200260200101519050610a84816000015182602001518360400151846060015185608001518660a001518a8860c00151611ded565b50600101610a17565b5050610a9860018055565b5050565b600a54600090815b816001600160401b0316816001600160401b03161015610b1057836001600160a01b0316600a826001600160401b031681548110610ae457610ae4613a62565b6000918252602090912001546001600160a01b031603610b08575060019392505050565b600101610aa4565b5060009392505050565b3360009081526009602052604090205460ff16610b495760405162461bcd60e51b815260040161069890613b1a565b610b5233610a9c565b15610baf5760405162461bcd60e51b815260206004820152602760248201527f4c6f636b657220616c726561647920766f74656420666f7220656d657267656e6044820152666379206c6f636b60c81b6064820152608401610698565b600a805460018101825560008290527fc65a7bb8d6351c1cf70c95a316cc6a92839c986682d98bc35f958f4883f9d2a80180546001600160a01b03191633179055600b5490546001600160401b03918216911610801590610c13575060005460ff16155b15610732576107326120e7565b604080517f12320cdb7a65d3e471a2a6ae0db5e58f3654c7d8bd5e6a1a70b3b8f2fea2aed760208201526001600160401b03808716928201929092529084166060820152600090608001604051602081830303815290604052805190602001209050610c8b816118e8565b610c99818484600354611959565b600d80546001600160c01b0316600160c01b6001600160401b038816908102919091179091556040519081527f0ef2da393c3832a8f08ce447e14948d21e84f864facf7327137387bd0596a563906020015b60405180910390a15050505050565b6060600a805480602002602001604051908101604052809291908181526020018280548015610d5257602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311610d34575b5050505050905090565b610d646116dd565b610d6c611707565b805160005b816001600160401b0316816001600160401b03161015610e8557610e7d83826001600160401b031681518110610da957610da9613a62565b60200260200101516000015184836001600160401b031681518110610dd057610dd0613a62565b60200260200101516020015185846001600160401b031681518110610df757610df7613a62565b60200260200101516040015186856001600160401b031681518110610e1e57610e1e613a62565b60200260200101516060015187866001600160401b031681518110610e4557610e45613a62565b60200260200101516080015188876001600160401b031681518110610e6c57610e6c613a62565b602002602001015160a0015161213c565b600101610d71565b5050610e9060018055565b50565b610e9b6116dd565b610ea3611707565b60008111610f045760405162461bcd60e51b815260206004820152602860248201527f4465706f73697420616d6f756e74206d7573742062652067726561746572207460448201526768616e207a65726f60c01b6064820152608401610698565b6001600160a01b038316610f5a5760405162461bcd60e51b815260206004820152601b60248201527f496e76616c69642064657374696e6174696f6e206164647265737300000000006044820152606401610698565b610f6f6001600160a01b03831633308461225a565b6001600160a01b0383163314610fcd57604080516001600160a01b0385811682528416602082015290810182905233907f7cfff908a4b583f36430b25d75964c458d8ede8a99bd61be750e97ee1b2f3a969060600160405180910390a25b610fd660018055565b505050565b604080517f3cafbb096e40bdb2551d91252c05f542adbb6effe4860d14431e6605a04fd5fa60208201526001600160401b03808716928201929092529084166060820152600090608001604051602081830303815290604052805190602001209050611046816118e8565b611054818484600354611959565b600d805467ffffffffffffffff60801b1916600160801b6001600160401b038816908102919091179091556040519081527f04edaf680108675f58d2ea70e9e7886c39ed38b66439622f8362d36595fe816990602001610ceb565b6110b76122c1565b60007f255d6a497f94e8b9aff3182744306fd858435def24a6ff727653948eadf175aa866000015187602001516040516020016110f49190613b69565b60405160208183030381529060405280519060200120886040015160405160200161111f9190613b69565b60405160208183030381529060405280519060200120896060015160405160200161114a9190613b9c565b60408051601f198184030181528282528051602091820120908301969096526001600160401b03948516908201526060810192909252608082015260a0810192909252831660c082015260e0016040516020818303038152906040528051906020012090506111b8816118e8565b611220866111c587613b0e565b8686808060200260200160405190810160405280939291908181526020016000905b828210156112135761120460608302860136819003810190613bcf565b815260200190600101906111e7565b50505050508460016122e4565b61122861184c565b611234600a6000612f44565b61123c6126c9565b505050505050565b61124c611707565b3360009081526009602052604090205460ff1661127b5760405162461bcd60e51b815260040161069890613b1a565b61128433610a9c565b6112ea5760405162461bcd60e51b815260206004820152603160248201527f4c6f636b6572206973206e6f742063757272656e746c7920766f74696e6720666044820152706f7220656d657267656e6379206c6f636b60781b6064820152608401610698565b61073233611c3b565b6112fb6116dd565b611303611707565b61130c3361172b565b8060005b816001600160401b0316816001600160401b03161015610a8d576113548484836001600160401b031681811061134857611348613a62565b90506020020135612702565b600101611310565b611364611707565b600254611370836129e5565b146113e35760405162461bcd60e51b815260206004820152603d60248201527f537570706c696564206163746976652076616c696461746f727320616e64207060448201527f6f7765727320646f206e6f74206d6174636820636865636b706f696e740000006064820152608401610698565b60007fcd26826da4f5c0e82ef8057ecacd8931dfb36167a70c820505f10826298cd05e846000015185602001516040516020016114209190613b69565b60405160208183030381529060405280519060200120866040015160405160200161144b9190613b69565b6040516020818303038152906040528051906020012087606001516040516020016114769190613b9c565b604051602081830303815290604052805190602001206040516020016114c79594939291909485526001600160401b0393909316602085015260408401919091526060830152608082015260a00190565b6040516020818303038152906040528051906020012090506114ed8484848460006122e4565b50505050565b604080517f6f884b20c1fc555c21adb663cf93411199935f7263773056bc4711ea7266763160208201526001600160a01b0387169181019190915284151560608201526001600160401b038416608082015260009060a001604051602081830303815290604052805190602001209050600085156115745750600254611579565b506003545b611582826118e8565b61158f8261095386613b0e565b6001600160a01b0387166000818152600c6020908152604091829020805460ff19168a151590811790915591519182527f2526bb92d75e00cfad8c7c16cb75f3e1073c854339e49b16baaad3067c2ed65a91016109f1565b604080517fd9c9d5e1590244de839782c61326c04e1b0506d5bf4a4750df02026b6df3953d60208201526001600160401b03808716928201929092529084166060820152600090608001604051602081830303815290604052805190602001209050611652816118e8565b611660818484600354611959565b600b805467ffffffffffffffff19166001600160401b03878116918217909255600a5490911610801590611697575060005460ff16155b156116a4576116a46120e7565b6040516001600160401b03861681527f2dbe453726b24b2cee427a7d6e2dcc9f353f16bee104f3d21480157a0ee409f790602001610ceb565b60026001540361170057604051633ee5aeb560e01b815260040160405180910390fd5b6002600155565b60005460ff16156107325760405163d93c066560e01b815260040160405180910390fd5b6001600160a01b0381166000908152600c602052604090205460ff16610e905760405162461bcd60e51b815260206004820152601960248201527f53656e646572206973206e6f7420612066696e616c697a6572000000000000006044820152606401610698565b600d5460009081906117b590600160801b90046001600160401b031685613c01565b6001600160401b031642119050806117d1576003915050611846565b600d5443906000906117f590600160801b90046001600160401b03166103e8613c28565b600d546001600160401b0391821691600160c01b909104166118178785613c53565b6118219190613c28565b6001600160401b03161190508061183e5760049350505050611846565b600093505050505b92915050565b6006546002819055600754600381905560048054600d80546001600160401b038084166001600160801b031990921691909117600160401b808504831602179182905567ffffffffffffffff60801b19909216909255604080519290911682526020820193909352918201527f87da17ff65d815d1e1c369cb3bbda9a11af181b92dc52681a2779419781c6270906060015b60405180910390a1565b60008181526008602052604090205460ff161561193e5760405162461bcd60e51b81526020600482015260146024820152731b595cdcd859d948185b1c9958591e481d5cd95960621b6044820152606401610698565b6000908152600860205260409020805460ff19166001179055565b80611963846129e5565b146119e75760405162461bcd60e51b815260206004820152604860248201527f537570706c696564206163746976652076616c696461746f727320616e64207060448201527f6f7765727320646f206e6f74206d6174636820746865206163746976652063686064820152671958dadc1bda5b9d60c21b608482015260a401610698565b81516001600160401b038116611a2f5760405162461bcd60e51b815260206004820152600d60248201526c5369676e65727320656d70747960981b6044820152606401610698565b6020840151516000908190815b816001600160401b0316816001600160401b03161015611b82576000611a878a89866001600160401b031681518110611a7757611a77613a62565b6020026020010151601154612b0f565b90508860200151826001600160401b031681518110611aa857611aa8613a62565b60200260200101516001600160a01b0316816001600160a01b031603611b795760008960400151836001600160401b031681518110611ae957611ae9613a62565b602002602001015190508086611aff9190613c01565b600d54909650611b2090600160401b90046001600160401b03166002613c28565b6001600160401b0316611b34876003613c28565b6001600160401b03161115611b4a575050611b82565b611b55600186613c01565b9450866001600160401b0316856001600160401b031610611b77575050611b82565b505b50600101611a3c565b50600d54611ba190600160401b90046001600160401b03166002613c28565b6001600160401b0316611bb5846003613c28565b6001600160401b031611611c315760405162461bcd60e51b815260206004820152603b60248201527f5375626d69747465642076616c696461746f7220736574207369676e6174757260448201527f657320646f206e6f74206861766520656e6f75676820706f77657200000000006064820152608401610698565b5050505050505050565b611c43611707565b6001600160a01b03811660009081526009602052604090205460ff16611cc45760405162461bcd60e51b815260206004820152603060248201527f41646472657373206973206e6f7420617574686f72697a656420746f206c6f6360448201526f1ac81cdb585c9d0818dbdb9d1c9858dd60821b6064820152608401610698565b600a5460005b816001600160401b0316816001600160401b03161015610fd657826001600160a01b0316600a826001600160401b031681548110611d0a57611d0a613a62565b6000918252602090912001546001600160a01b031603611de557600a611d31600184613c53565b6001600160401b031681548110611d4a57611d4a613a62565b600091825260209091200154600a80546001600160a01b03909216916001600160401b038416908110611d7f57611d7f613a62565b9060005260206000200160006101000a8154816001600160a01b0302191690836001600160a01b03160217905550600a805480611dbe57611dbe613c73565b600082815260209020810160001990810180546001600160a01b0319169055019055505050565b600101611cca565b604080517fae7dee9fe1cf9016b724a74908236e5f57a1789d05c09f2625fa9baee0cde49d60208201526001600160a01b038a1691810191909152606081018890526080810187905260a0810186905260c081018590526001600160401b03841660e082015260009061010001604051602081830303815290604052805190602001209050611e8c8160009081526010602052604090205460ff161590565b611ebd576040805182815260056020820152600080516020613cfc833981519152910160405180910390a150611c31565b60006040518061012001604052808b6001600160a01b031681526020018a8152602001898152602001888152602001878152602001866001600160401b03168152602001426001600160401b03168152602001611f174390565b6001600160401b03908116825260209182018590526000858152600e909252604090912060050154919250600160401b9091041615611f7e576040805183815260006020820152600080516020613cfc833981519152910160405180910390a15050611c31565b611f9482611f8b86613b0e565b85600254611959565b6000828152600e6020908152604091829020835181546001600160a01b039091166001600160a01b0319909116811782559184015160018201819055838501516002830181905560608601516003840181905560808701516004850181905560a088015160058601805460c08b015160e08c01516001600160401b03818116600160801b0267ffffffffffffffff60801b19848316600160401b026001600160801b0319909616928816929092179490941716929092179092556101008b01516006909801889055985197987f2fe552dd13681968cbb9b3bf73bae9ca6abce7b77eead2e0600e94a4812d999e986120d398979695949392919097885260208801969096526040870194909452606086019290925260808501526001600160401b0390811660a085015290811660c08401521660e08201526101000190565b60405180910390a250505050505050505050565b6120ef611707565b6000805460ff191660011790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a2586121243390565b6040516001600160a01b0390911681526020016118de565b60408181015182516020840151925163d505accf60e01b81526001600160a01b038a811660048301523060248301819052604483018990526001600160401b038816606484015260ff909416608483015260a482019290925260c481019390935290919086169063d505accf9060e401600060405180830381600087803b1580156121c657600080fd5b505af11580156121da573d6000803e3d6000fd5b506121f4925050506001600160a01b03861688838761225a565b866001600160a01b0316866001600160a01b0316146108ae57604080516001600160a01b0388811682528781166020830152918101869052908816907f7cfff908a4b583f36430b25d75964c458d8ede8a99bd61be750e97ee1b2f3a96906060016109f1565b6040516001600160a01b0384811660248301528381166044830152606482018390526114ed9186918216906323b872dd906084015b604051602081830303815290604052915060e01b6020820180516001600160e01b038381831617835250505050612cd8565b60005460ff1661073257604051638dfc202b60e01b815260040160405180910390fd5b846040015151856020015151146123555760405162461bcd60e51b815260206004820152602f60248201527f4e657720686f7420616e6420636f6c642076616c696461746f7220736574732060448201526e0d8cadccee8d040dad2e6dac2e8c6d608b1b6064820152608401610698565b846060015151856020015151146123c35760405162461bcd60e51b815260206004820152602c60248201527f4e65772076616c696461746f722073657420616e6420706f77657273206c656e60448201526b0cee8d040dad2e6dac2e8c6d60a31b6064820152608401610698565b835185516001600160401b039182169116116124475760405162461bcd60e51b815260206004820152603d60248201527f4e65772076616c696461746f72207365742065706f6368206d7573742062652060448201527f67726561746572207468616e20746865206163746976652065706f63680000006064820152608401610698565b60006124568660600151612d3b565b905060008215612469575060035461246e565b506002545b61247a84878784611959565b6124a7604051806060016040528060006001600160401b0316815260200160608152602001606081525090565b604051806060016040528089600001516001600160401b03168152602001896020015181526020018960600151815250905060006124e4826129e5565b9050612513604051806060016040528060006001600160401b0316815260200160608152602001606081525090565b5060408051606080820183528b516001600160401b031682528b83015160208301528b015191810191909152600061254a826129e5565b905060004290506040518060e001604052808d600001516001600160401b03168152602001886001600160401b03168152602001826001600160401b031681526020016125944390565b6001600160401b03908116825260208881015151821681840152604080840189905260609384018790528451600480549387015187840151968801518616600160c01b026001600160c01b03978716600160801b02979097166fffffffffffffffffffffffffffffffff918716600160401b026001600160801b0319909616938716939093179490941793909316179390931790556080830151600580549190921667ffffffffffffffff199190911617905560a082015160065560c0909101516007558c5190517f420bbe99bd2c52ec500d33614359525f3ef7bb3358c0e07d1312db0941cbf2f4916126b3918790869086906001600160401b03948516815260208101939093526040830191909152909116606082015260800190565b60405180910390a1505050505050505050505050565b6126d16122c1565b6000805460ff191690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa33612124565b60008181526010602052604090205460ff1615612744576040805182815260056020820152600080516020613cfc83398151915291015b60405180910390a150565b6000818152600f602052604090205460ff161561277f576040805182815260016020820152600080516020613cfc8339815191529101612739565b6000818152600e6020908152604080832081516101208101835281546001600160a01b03168152600182015493810193909352600281015491830191909152600381015460608301526004810154608083015260058101546001600160401b0380821660a0850152600160401b8204811660c08501819052600160801b9092041660e0840152600690910154610100830152909103612843576040805183815260026020820152600080516020613cfc833981519152910160405180910390a15050565b60006128578260c001518360e00151611793565b905063ffffffff811615612896576040805184815263ffffffff83166020820152600080516020613cfc833981519152910160405180910390a1505050565b6000838152600f60209081526040909120805460ff191660011790558201516128c6906001600160a01b03101590565b80156128dd575060408201516001600160a01b0310155b61293b5760405162461bcd60e51b815260206004820152602960248201527f496e76616c69642045564d20776974686472617720746f6b656e206f7220646560448201526839ba34b730ba34b7b760b91b6064820152608401610698565b612963612949836020015190565b606084015160408501516001600160a01b03169190612e0c565b815161010083015160208085015160408087015160608089015160a0808b0151855198895296880195909552928601919091528401526001600160401b0390911660808301526001600160a01b03909216917f070f33f8f3095f4b7388e8db49f1902a3005e9cc9cc4ada25316b429403db4d8910160405180910390a2505050565b600081604001515182602001515114612a405760405162461bcd60e51b815260206004820152601760248201527f4d616c666f726d65642076616c696461746f72207365740000000000000000006044820152606401610698565b60007fcf7a991d34f65202b9a5ebe03e28c3fd6f86e1f75fabbddd532864507554c66783600001518460200151604051602001612a7d9190613b69565b604051602081830303815290604052805190602001208560400151604051602001612aa89190613b9c565b60405160208183030381529060405280519060200120604051602001612af094939291909384526001600160401b039290921660208401526040830152606082015260800190565b60408051601f1981840301815291905280516020909101209392505050565b81516000908103612b625760405162461bcd60e51b815260206004820152601b60248201527f496e76616c6964207369676e6174757265202772272076616c756500000000006044820152606401610698565b8260200151600003612bb65760405162461bcd60e51b815260206004820152601b60248201527f496e76616c6964207369676e6174757265202773272076616c756500000000006044820152606401610698565b60405161190160f01b6020820152602281018390526042810185905260009060620160408051601f1981840301815282825280516020918201208783015188518984015160008088529690940194859052919550600193612c329387939193845260ff9290921660208401526040830152606082015260800190565b6020604051602081039080840390855afa158015612c54573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b038116612ccd5760405162461bcd60e51b815260206004820152602d60248201527f496e76616c6964207369676e61747572652c207265636f76657265642074686560448201526c207a65726f206164647265737360981b6064820152608401610698565b9150505b9392505050565b6000612ced6001600160a01b03841683612e3d565b90508051600014158015612d12575080806020019051810190612d109190613c89565b155b15610fd657604051635274afe760e01b81526001600160a01b0384166004820152602401610698565b60008060005b8351816001600160401b03161015612d955783816001600160401b031681518110612d6e57612d6e613a62565b602002602001015182612d819190613c01565b915080612d8d81613ca6565b915050612d41565b506000816001600160401b0316116118465760405162461bcd60e51b815260206004820152603460248201527f5375626d69747465642076616c696461746f7220706f77657273206d7573742060448201527362652067726561746572207468616e207a65726f60601b6064820152608401610698565b6040516001600160a01b03838116602483015260448201839052610fd691859182169063a9059cbb9060640161228f565b6060612cd18383600084600080856001600160a01b03168486604051612e639190613ccc565b60006040518083038185875af1925050503d8060008114612ea0576040519150601f19603f3d011682016040523d82523d6000602084013e612ea5565b606091505b5091509150612eb5868383612ebf565b9695505050505050565b606082612ed457612ecf82612f1b565b612cd1565b8151158015612eeb57506001600160a01b0384163b155b15612f1457604051639996b31560e01b81526001600160a01b0385166004820152602401610698565b5080612cd1565b805115612f2b5780518082602001fd5b604051630a12f52160e11b815260040160405180910390fd5b5080546000825590600052602060002090810190610e9091905b80821115612f725760008155600101612f5e565b5090565b634e487b7160e01b600052604160045260246000fd5b604051606081016001600160401b0381118282101715612fae57612fae612f76565b60405290565b60405160e081016001600160401b0381118282101715612fae57612fae612f76565b60405160c081016001600160401b0381118282101715612fae57612fae612f76565b604051601f8201601f191681016001600160401b038111828210171561302057613020612f76565b604052919050565b60006001600160401b0382111561304157613041612f76565b5060051b60200190565b80356001600160401b038116811461306257600080fd5b919050565b80356001600160a01b038116811461306257600080fd5b600082601f83011261308f57600080fd5b813560206130a461309f83613028565b612ff8565b8083825260208201915060208460051b8701019350868411156130c657600080fd5b602086015b848110156130e9576130dc81613067565b83529183019183016130cb565b509695505050505050565b600082601f83011261310557600080fd5b8135602061311561309f83613028565b8083825260208201915060208460051b87010193508684111561313757600080fd5b602086015b848110156130e95761314d8161304b565b835291830191830161313c565b60006060828403121561316c57600080fd5b613174612f8c565b905061317f8261304b565b815260208201356001600160401b038082111561319b57600080fd5b6131a78583860161307e565b602084015260408401359150808211156131c057600080fd5b506131cd848285016130f4565b60408301525092915050565b6000606082840312156131eb57600080fd5b6131f3612f8c565b90508135815260208201356020820152604082013560ff8116811461321757600080fd5b604082015292915050565b600082601f83011261323357600080fd5b8135602061324361309f83613028565b80838252602082019150606060206060860288010194508785111561326757600080fd5b602087015b8581101561328b5761327e89826131d9565b845292840192810161326c565b5090979650505050505050565b600080600080608085870312156132ae57600080fd5b84356001600160401b03808211156132c557600080fd5b818701915087601f8301126132d957600080fd5b813560206132e961309f83613028565b82815260059290921b8401810191818101908b84111561330857600080fd5b948201945b838610156133265785358252948201949082019061330d565b9850613335905089820161304b565b96505050604087013591508082111561334d57600080fd5b6133598883890161315a565b9350606087013591508082111561336f57600080fd5b5061337c87828801613222565b91505092959194509250565b8015158114610e9057600080fd5b6000606082840312156133a857600080fd5b50919050565b600080600080600060a086880312156133c657600080fd5b6133cf86613067565b945060208601356133df81613388565b93506133ed6040870161304b565b925060608601356001600160401b038082111561340957600080fd5b61341589838a01613396565b9350608088013591508082111561342b57600080fd5b5061343888828901613222565b9150509295509295909350565b6000806040838503121561345857600080fd5b82356001600160401b038082111561346f57600080fd5b818501915085601f83011261348357600080fd5b8135602061349361309f83613028565b82815260059290921b840181019181810190898411156134b257600080fd5b8286015b8481101561356b578035868111156134cd57600080fd5b870160e0818d03601f190112156134e45760008081fd5b6134ec612fb4565b6134f7868301613067565b815260408201358682015260608083013560408301526080808401358284015260a0915081840135818401525060c061353181850161304b565b8284015260e084013591508982111561354a5760008081fd5b6135588f8984870101613222565b90830152508452509183019183016134b6565b509650508601359250508082111561358257600080fd5b5061358f85828601613396565b9150509250929050565b6000602082840312156135ab57600080fd5b612cd182613067565b6000602082840312156135c657600080fd5b5035919050565b600080600080608085870312156135e357600080fd5b6135ec8561304b565b93506135fa6020860161304b565b925060408501356001600160401b038082111561334d57600080fd5b6020808252825182820181905260009190848201906040850190845b818110156136575783516001600160a01b031683529284019291840191600101613632565b50909695505050505050565b6000602080838503121561367657600080fd5b82356001600160401b0381111561368c57600080fd5b8301601f8101851361369d57600080fd5b80356136ab61309f82613028565b81815260089190911b820183019083810190878311156136ca57600080fd5b928401925b8284101561375c5761010084890312156136e95760008081fd5b6136f1612fd6565b6136fa85613067565b8152613707868601613067565b868201526040613718818701613067565b9082015260608581013590820152608061373381870161304b565b9082015260a06137458a8783016131d9565b9082015282526101009390930192908401906136cf565b979650505050505050565b60008060006060848603121561377c57600080fd5b61378584613067565b925061379360208501613067565b9150604084013590509250925092565b6000608082840312156137b557600080fd5b604051608081016001600160401b0382821081831117156137d8576137d8612f76565b816040528293506137e88561304b565b835260208501359150808211156137fe57600080fd5b61380a8683870161307e565b6020840152604085013591508082111561382357600080fd5b61382f8683870161307e565b6040840152606085013591508082111561384857600080fd5b50613855858286016130f4565b6060830152505092915050565b60008060008060006080868803121561387a57600080fd5b85356001600160401b038082111561389157600080fd5b61389d89838a016137a3565b965060208801359150808211156138b357600080fd5b6138bf89838a01613396565b955060408801359150808211156138d557600080fd5b818801915088601f8301126138e957600080fd5b8135818111156138f857600080fd5b89602060608302850101111561390d57600080fd5b6020830195508094505050506139256060870161304b565b90509295509295909350565b6000806020838503121561394457600080fd5b82356001600160401b038082111561395b57600080fd5b818501915085601f83011261396f57600080fd5b81358181111561397e57600080fd5b8660208260051b850101111561399357600080fd5b60209290920196919550909350505050565b6000806000606084860312156139ba57600080fd5b83356001600160401b03808211156139d157600080fd5b6139dd878388016137a3565b945060208601359150808211156139f357600080fd5b6139ff8783880161315a565b93506040860135915080821115613a1557600080fd5b50613a2286828701613222565b9150509250925092565b815160009082906020808601845b83811015613a5657815185529382019390820190600101613a3a565b50929695505050505050565b634e487b7160e01b600052603260045260246000fd5b81546001600160a01b0316815261012081016001830154602083015260028301546040830152600383015460608301526004830154608083015260058301546001600160401b0380821660a0850152613ae160c08501828460401c166001600160401b03169052565b613afb60e08501828460801c166001600160401b03169052565b5050600683015461010083015292915050565b6000611846368361315a565b6020808252602f908201527f53656e646572206973206e6f7420617574686f72697a656420746f206c6f636b60408201526e081cdb585c9d0818dbdb9d1c9858dd608a1b606082015260800190565b815160009082906020808601845b83811015613a565781516001600160a01b031685529382019390820190600101613b77565b815160009082906020808601845b83811015613a565781516001600160401b031685529382019390820190600101613baa565b600060608284031215613be157600080fd5b612cd183836131d9565b634e487b7160e01b600052601160045260246000fd5b6001600160401b03818116838216019080821115613c2157613c21613beb565b5092915050565b6001600160401b03818116838216028082169190828114613c4b57613c4b613beb565b505092915050565b6001600160401b03828116828216039080821115613c2157613c21613beb565b634e487b7160e01b600052603160045260246000fd5b600060208284031215613c9b57600080fd5b8151612cd181613388565b60006001600160401b03808316818103613cc257613cc2613beb565b6001019392505050565b6000825160005b81811015613ced5760208186018101518583015201613cd3565b50600092019182525091905056fe686cb4bac974cd11b0f8a75fc7c7764ed12cc46faaec53110f807aa802a7acb4a26469706673582212209dfcbf00e7026836fe540cd1a519400003ac6ac1cc5fe498cbd1602ae5d4383e64736f6c63430008160033",
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
// Solidity: function requestedWithdrawals(bytes32 ) view returns(address user, bytes32 destination, bytes32 token, uint256 amount, uint256 chainId, uint64 nonce, uint64 requestedTime, uint64 requestedBlockNumber, bytes32 message)
func (_Bridge *BridgeCaller) RequestedWithdrawals(opts *bind.CallOpts, arg0 [32]byte) (struct {
	User                 common.Address
	Destination          [32]byte
	Token                [32]byte
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
		Destination          [32]byte
		Token                [32]byte
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
	outstruct.Destination = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)
	outstruct.Token = *abi.ConvertType(out[2], new([32]byte)).(*[32]byte)
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
// Solidity: function requestedWithdrawals(bytes32 ) view returns(address user, bytes32 destination, bytes32 token, uint256 amount, uint256 chainId, uint64 nonce, uint64 requestedTime, uint64 requestedBlockNumber, bytes32 message)
func (_Bridge *BridgeSession) RequestedWithdrawals(arg0 [32]byte) (struct {
	User                 common.Address
	Destination          [32]byte
	Token                [32]byte
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
// Solidity: function requestedWithdrawals(bytes32 ) view returns(address user, bytes32 destination, bytes32 token, uint256 amount, uint256 chainId, uint64 nonce, uint64 requestedTime, uint64 requestedBlockNumber, bytes32 message)
func (_Bridge *BridgeCallerSession) RequestedWithdrawals(arg0 [32]byte) (struct {
	User                 common.Address
	Destination          [32]byte
	Token                [32]byte
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

// BatchedRequestWithdrawals is a paid mutator transaction binding the contract method 0x24f0b6c7.
//
// Solidity: function batchedRequestWithdrawals((address,bytes32,bytes32,uint256,uint256,uint64,(uint256,uint256,uint8)[])[] withdrawalRequests, (uint64,address[],uint64[]) hotValidatorSet) returns()
func (_Bridge *BridgeTransactor) BatchedRequestWithdrawals(opts *bind.TransactOpts, withdrawalRequests []WithdrawalRequest, hotValidatorSet ValidatorSet) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "batchedRequestWithdrawals", withdrawalRequests, hotValidatorSet)
}

// BatchedRequestWithdrawals is a paid mutator transaction binding the contract method 0x24f0b6c7.
//
// Solidity: function batchedRequestWithdrawals((address,bytes32,bytes32,uint256,uint256,uint64,(uint256,uint256,uint8)[])[] withdrawalRequests, (uint64,address[],uint64[]) hotValidatorSet) returns()
func (_Bridge *BridgeSession) BatchedRequestWithdrawals(withdrawalRequests []WithdrawalRequest, hotValidatorSet ValidatorSet) (*types.Transaction, error) {
	return _Bridge.Contract.BatchedRequestWithdrawals(&_Bridge.TransactOpts, withdrawalRequests, hotValidatorSet)
}

// BatchedRequestWithdrawals is a paid mutator transaction binding the contract method 0x24f0b6c7.
//
// Solidity: function batchedRequestWithdrawals((address,bytes32,bytes32,uint256,uint256,uint64,(uint256,uint256,uint8)[])[] withdrawalRequests, (uint64,address[],uint64[]) hotValidatorSet) returns()
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
	Destination [32]byte
	Token       [32]byte
	Amount      *big.Int
	Nonce       uint64
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterFinalizedWithdrawal is a free log retrieval operation binding the contract event 0x070f33f8f3095f4b7388e8db49f1902a3005e9cc9cc4ada25316b429403db4d8.
//
// Solidity: event FinalizedWithdrawal(bytes32 message, address indexed user, bytes32 destination, bytes32 token, uint256 amount, uint64 nonce)
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

// WatchFinalizedWithdrawal is a free log subscription operation binding the contract event 0x070f33f8f3095f4b7388e8db49f1902a3005e9cc9cc4ada25316b429403db4d8.
//
// Solidity: event FinalizedWithdrawal(bytes32 message, address indexed user, bytes32 destination, bytes32 token, uint256 amount, uint64 nonce)
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

// ParseFinalizedWithdrawal is a log parse operation binding the contract event 0x070f33f8f3095f4b7388e8db49f1902a3005e9cc9cc4ada25316b429403db4d8.
//
// Solidity: event FinalizedWithdrawal(bytes32 message, address indexed user, bytes32 destination, bytes32 token, uint256 amount, uint64 nonce)
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

// FilterInvalidatedWithdrawal is a free log retrieval operation binding the contract event 0x5fe92156b011257af09c3fceff5d0f5ee2c781900477e1eadca143dfe07ab9fd.
//
// Solidity: event InvalidatedWithdrawal((address,bytes32,bytes32,uint256,uint256,uint64,uint64,uint64,bytes32) withdrawal)
func (_Bridge *BridgeFilterer) FilterInvalidatedWithdrawal(opts *bind.FilterOpts) (*BridgeInvalidatedWithdrawalIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "InvalidatedWithdrawal")
	if err != nil {
		return nil, err
	}
	return &BridgeInvalidatedWithdrawalIterator{contract: _Bridge.contract, event: "InvalidatedWithdrawal", logs: logs, sub: sub}, nil
}

// WatchInvalidatedWithdrawal is a free log subscription operation binding the contract event 0x5fe92156b011257af09c3fceff5d0f5ee2c781900477e1eadca143dfe07ab9fd.
//
// Solidity: event InvalidatedWithdrawal((address,bytes32,bytes32,uint256,uint256,uint64,uint64,uint64,bytes32) withdrawal)
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

// ParseInvalidatedWithdrawal is a log parse operation binding the contract event 0x5fe92156b011257af09c3fceff5d0f5ee2c781900477e1eadca143dfe07ab9fd.
//
// Solidity: event InvalidatedWithdrawal((address,bytes32,bytes32,uint256,uint256,uint64,uint64,uint64,bytes32) withdrawal)
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
	Destination          [32]byte
	Token                [32]byte
	Amount               *big.Int
	ChainId              *big.Int
	Nonce                uint64
	RequestedTime        uint64
	RequestedBlockNumber uint64
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterRequestedWithdrawal is a free log retrieval operation binding the contract event 0x2fe552dd13681968cbb9b3bf73bae9ca6abce7b77eead2e0600e94a4812d999e.
//
// Solidity: event RequestedWithdrawal(bytes32 message, address indexed user, bytes32 destination, bytes32 token, uint256 amount, uint256 chainId, uint64 nonce, uint64 requestedTime, uint64 requestedBlockNumber)
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

// WatchRequestedWithdrawal is a free log subscription operation binding the contract event 0x2fe552dd13681968cbb9b3bf73bae9ca6abce7b77eead2e0600e94a4812d999e.
//
// Solidity: event RequestedWithdrawal(bytes32 message, address indexed user, bytes32 destination, bytes32 token, uint256 amount, uint256 chainId, uint64 nonce, uint64 requestedTime, uint64 requestedBlockNumber)
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

// ParseRequestedWithdrawal is a log parse operation binding the contract event 0x2fe552dd13681968cbb9b3bf73bae9ca6abce7b77eead2e0600e94a4812d999e.
//
// Solidity: event RequestedWithdrawal(bytes32 message, address indexed user, bytes32 destination, bytes32 token, uint256 amount, uint256 chainId, uint64 nonce, uint64 requestedTime, uint64 requestedBlockNumber)
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
