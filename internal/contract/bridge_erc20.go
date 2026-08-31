package contract

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
)

func (transfer *BridgeERC20Transfer) ToTypedData(srcChainId, hubChainId *big.Int, verifyingContract common.Address) apitypes.TypedData {
	return transfer.BridgeERC20TransferDepositTypedData(
		srcChainId,
		hubChainId,
		verifyingContract,
		transfer.Raw.BlockNumber,
		transfer.Raw.TxHash,
		uint32(transfer.Raw.Index),
	)
}

func (transfer *BridgeERC20Transfer) ToDepositConfirm(srcChainId *big.Int, signature Signature) DepositConfirm {
	return DepositConfirm{
		User:        addressToBytes32(transfer.From),
		Destination: transfer.To,
		Token:       addressToBytes32(transfer.Raw.Address),
		Amount:      transfer.Value,
		ChainId:     srcChainId,
		BlockNumber: transfer.Raw.BlockNumber,
		TxHash:      transfer.Raw.TxHash,
		Index:       uint32(transfer.Raw.Index),
		Signature:   signature,
	}
}

// BridgeERC20TransferDepositTypedData builds EIP-712 typed data for deposit confirmation.
func (transfer *BridgeERC20Transfer) BridgeERC20TransferDepositTypedData(
	srcChainId, hubChainId *big.Int,
	verifyingContract common.Address,
	blockNumber uint64,
	txHash common.Hash,
	index uint32,
) apitypes.TypedData {
	user := addressToBytes32(transfer.From)
	token := addressToBytes32(transfer.Raw.Address)
	return apitypes.TypedData{
		Domain: NewBridgeHubDomain(hubChainId, verifyingContract),
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
			"user":        bytes32Hex(user),
			"destination": transfer.To.Hex(),
			"token":       bytes32Hex(token),
			"amount":      transfer.Value.String(),
			"chainId":     srcChainId.String(),
			"blockNumber": formatUint64(blockNumber),
			"txHash":      txHash.Hex(),
			"index":       formatUint32(index),
		},
	}
}

func addressToBytes32(addr common.Address) [32]byte {
	var out [32]byte
	copy(out[12:], addr.Bytes())
	return out
}

func bytes32Hex(b [32]byte) string {
	return "0x" + common.Bytes2Hex(b[:])
}

func formatUint64(v uint64) string {
	return new(big.Int).SetUint64(v).String()
}

func formatUint32(v uint32) string {
	return new(big.Int).SetUint64(uint64(v)).String()
}
