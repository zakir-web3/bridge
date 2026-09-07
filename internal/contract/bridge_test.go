package contract

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
	"github.com/stretchr/testify/assert"
)

func TestBridgeFinalizedWithdrawal_ToTypedData(t *testing.T) {
	userAddr := common.HexToAddress("0x1000000000000000000000000000000000000001")
	destinationAddr := common.HexToAddress("0x1000000000000000000000000000000000000001")
	tokenAddr := common.HexToAddress("0x2000000000000000000000000000000000000002")
	verifyingContract := common.HexToAddress("0x3000000000000000000000000000000000000003")

	srcChainId := big.NewInt(1337)
	chainId := big.NewInt(97)

	withdrawal := &BridgeFinalizedWithdrawal{
		User:        userAddr,
		Destination: AddressToBytes32(destinationAddr),
		Token:       AddressToBytes32(tokenAddr),
		Amount:      big.NewInt(1e18),
		Nonce:       12345,
		Raw: types.Log{
			Address:     tokenAddr,
			BlockNumber: uint64(1),
			TxHash: common.HexToHash(
				"0x0000000000000000000000000000000000000000000000000000000000000000",
			),
			Index: uint(2),
		},
	}

	typedData := withdrawal.ToTypedData(srcChainId, chainId, verifyingContract)
	assert.NotNil(t, typedData)

	domainSeparator, err := typedData.HashStruct("EIP712Domain", typedData.Domain.Map())
	assert.NoError(t, err)
	assert.Equal(t, "0x25e1c2d18a3106cd706593b1aa9b6e16121734a2b98f8fa3e8bc690576e2da35", domainSeparator.String(), "Domain separator should not be empty")

	typedDataHash, err := typedData.HashStruct(typedData.PrimaryType, typedData.Message)
	assert.NoError(t, err)
	assert.Equal(t, "0x2486f3baf331176f06d4eda1e971483466ca196931705258bc5ef971b81d4ad2", typedDataHash.String(), "Typed data hash should not be empty")

	hash, _, err := apitypes.TypedDataAndHash(typedData)
	assert.NoError(t, err)
	assert.Equal(t, "0x39e51aa182071d8a9efef2a23f2f2dffbcb80509e75a14e2002bfdd09dbebecb", common.BytesToHash(hash).String(), "Complete typed data hash should not be empty")
}

func TestBridgeFinalizedWithdrawal_ToWithdrawConfirm(t *testing.T) {
	userAddr := common.HexToAddress("0x1000000000000000000000000000000000000001")
	destAddr := common.HexToAddress("0x1000000000000000000000000000000000000009")
	tokenAddr := common.HexToAddress("0x2000000000000000000000000000000000000002")
	chainId := big.NewInt(97)
	sig := Signature{R: big.NewInt(1), S: big.NewInt(2), V: 28}

	withdrawal := &BridgeFinalizedWithdrawal{
		User:        userAddr,
		Destination: AddressToBytes32(destAddr),
		Token:       AddressToBytes32(tokenAddr),
		Amount:      big.NewInt(1e18),
		Nonce:       12345,
	}

	got := withdrawal.ToWithdrawConfirm(chainId, sig)
	assert.Equal(t, userAddr, got.User)
	assert.Equal(t, AddressToBytes32(destAddr), got.Destination)
	assert.Equal(t, AddressToBytes32(tokenAddr), got.Token)
	assert.Equal(t, big.NewInt(1e18), got.Amount)
	assert.Equal(t, chainId, got.ChainId)
	assert.Equal(t, uint64(12345), got.Nonce)
	assert.Equal(t, sig, got.Signature)
}
