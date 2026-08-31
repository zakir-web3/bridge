package contract

import (
	"encoding/hex"
	"math/big"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
)

// SolanaDeposit represents a parsed Solana bridge deposit for EIP-712 signing.
type SolanaDeposit struct {
	User                 [32]byte
	Destination          common.Address
	Token                [32]byte
	Amount               *big.Int
	ChainID              *big.Int
	BlockNumber          uint64
	TxHash               [32]byte
	Index                uint32
	TransactionSignature string
}

func (d *SolanaDeposit) ToTypedData(srcChainID, hubChainID *big.Int, verifyingContract common.Address) apitypes.TypedData {
	return apitypes.TypedData{
		Domain: NewBridgeHubDomain(hubChainID, verifyingContract),
		Types: apitypes.Types{
			"EIP712Domain": {
				{Name: "name", Type: "string"},
				{Name: "version", Type: "string"},
				{Name: "chainId", Type: "uint256"},
				{Name: "verifyingContract", Type: "address"},
			},
			"Deposit": {
				{Name: "user", Type: "bytes32"},
				{Name: "destination", Type: "address"},
				{Name: "token", Type: "bytes32"},
				{Name: "amount", Type: "uint256"},
				{Name: "chainId", Type: "uint256"},
				{Name: "blockNumber", Type: "uint64"},
				{Name: "txHash", Type: "bytes32"},
				{Name: "index", Type: "uint32"},
			},
		},
		PrimaryType: "Deposit",
		Message: apitypes.TypedDataMessage{
			"user":        "0x" + hex.EncodeToString(d.User[:]),
			"destination": d.Destination.Hex(),
			"token":       "0x" + hex.EncodeToString(d.Token[:]),
			"amount":      d.Amount.String(),
			"chainId":     srcChainID.String(),
			"blockNumber": strconv.FormatUint(d.BlockNumber, 10),
			"txHash":      "0x" + hex.EncodeToString(d.TxHash[:]),
			"index":       strconv.FormatUint(uint64(d.Index), 10),
		},
	}
}

func (d *SolanaDeposit) ToDepositConfirm(signature Signature) DepositConfirm {
	return DepositConfirm{
		User:        d.User,
		Destination: d.Destination,
		Token:       d.Token,
		Amount:      d.Amount,
		ChainId:     d.ChainID,
		BlockNumber: d.BlockNumber,
		TxHash:      d.TxHash,
		Index:       d.Index,
		Signature:   signature,
	}
}
