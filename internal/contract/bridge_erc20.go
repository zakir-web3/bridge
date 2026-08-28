package contract

import (
	"math/big"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
)

// EvmAddressToBytes32 matches BridgeHub._addressToBytes32 (left-padded EVM address).
func EvmAddressToBytes32(addr common.Address) [32]byte {
	var out [32]byte
	copy(out[12:], addr.Bytes())
	return out
}

func (transfer *BridgeERC20Transfer) ToTypedData(srcChainId, chainId *big.Int, verifyingContract common.Address) apitypes.TypedData {
	user := EvmAddressToBytes32(transfer.From)
	token := EvmAddressToBytes32(transfer.Raw.Address)
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
			"user":        common.BytesToHash(user[:]).Hex(),
			"destination": transfer.To.Hex(),
			"token":       common.BytesToHash(token[:]).Hex(),
			"amount":      transfer.Value.String(),
			"chainId":     srcChainId.String(),
			"blockNumber": strconv.FormatUint(transfer.Raw.BlockNumber, 10),
			"txHash":      transfer.Raw.TxHash.Hex(),
			"index":       strconv.FormatUint(uint64(transfer.Raw.Index), 10),
		},
	}
}

func (transfer *BridgeERC20Transfer) ToDepositConfirm(srcChainId *big.Int, signature Signature) DepositConfirm {
	return DepositConfirm{
		User:        EvmAddressToBytes32(transfer.From),
		Destination: transfer.To,
		Amount:      transfer.Value,
		Token:       EvmAddressToBytes32(transfer.Raw.Address),
		ChainId:     srcChainId,
		BlockNumber: transfer.Raw.BlockNumber,
		TxHash:      transfer.Raw.TxHash,
		Index:       uint32(transfer.Raw.Index),
		Signature:   signature,
	}
}
