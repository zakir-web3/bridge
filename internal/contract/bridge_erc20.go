package contract

import (
	"math/big"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
)

func (transfer *BridgeERC20Transfer) ToTypedData(srcChainId, chainId *big.Int, verifyingContract common.Address) apitypes.TypedData {
	return apitypes.TypedData{
		Domain: NewBridgeHubDomain(chainId, verifyingContract),
		Types: apitypes.Types{
			"EIP712Domain": {
				{Name: "name", Type: "string"},
				{Name: "version", Type: "string"},
				{Name: "chainId", Type: "uint256"},
				{Name: "verifyingContract", Type: "address"},
			},
			"Deposit": {
				{Name: "user", Type: "address"},
				{Name: "destination", Type: "address"},
				{Name: "token", Type: "address"},
				{Name: "amount", Type: "uint256"},
				{Name: "chainId", Type: "uint256"},
				{Name: "blockNumber", Type: "uint64"},
				{Name: "txHash", Type: "bytes32"},
				{Name: "logIndex", Type: "uint64"},
			},
		},
		PrimaryType: "Deposit",
		Message: apitypes.TypedDataMessage{
			"user":        transfer.From.String(),
			"destination": transfer.To.String(),
			"token":       transfer.Raw.Address.String(),
			"amount":      transfer.Value.String(),
			"chainId":     srcChainId.String(),
			"blockNumber": strconv.FormatUint(transfer.Raw.BlockNumber, 10),
			"txHash":      transfer.Raw.TxHash.String(),
			"logIndex":    strconv.FormatUint(uint64(transfer.Raw.Index), 10),
		},
	}
}

func (transfer *BridgeERC20Transfer) ToDepositConfirm(srcChainId *big.Int, signature Signature) DepositConfirm {
	return DepositConfirm{
		User:        transfer.From,
		Destination: transfer.To,
		Amount:      transfer.Value,
		Token:       transfer.Raw.Address,
		ChainId:     srcChainId,
		BlockNumber: transfer.Raw.BlockNumber,
		TxHash:      transfer.Raw.TxHash,
		LogIndex:    uint64(transfer.Raw.Index),
		Signature:   signature,
	}
}
