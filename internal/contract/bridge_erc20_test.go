package contract

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
	"github.com/stretchr/testify/assert"
)

func TestBridgeERC20Transfer_ToTypedData(t *testing.T) {
	fromAddr := common.HexToAddress("0x1000000000000000000000000000000000000001")
	toAddr := common.HexToAddress("0x1000000000000000000000000000000000000001")
	tokenAddr := common.HexToAddress("0x2000000000000000000000000000000000000002")
	verifyingContract := common.HexToAddress("0x3000000000000000000000000000000000000003")

	srcChainId := big.NewInt(97)
	chainId := big.NewInt(1337)
	transfer := &BridgeERC20Transfer{
		From:  fromAddr,
		To:    toAddr,
		Value: big.NewInt(1e18),
		Raw: types.Log{
			Address:     tokenAddr,
			BlockNumber: uint64(1),
			TxHash: common.HexToHash(
				"0x0000000000000000000000000000000000000000000000000000000000000000",
			),
			Index: uint(2),
		},
	}

	typedData := transfer.ToTypedData(srcChainId, chainId, verifyingContract)
	assert.NotNil(t, typedData)

	domainSeparator, err := typedData.HashStruct("EIP712Domain", typedData.Domain.Map())
	assert.NoError(t, err)
	assert.Equal(
		t,
		"0xf3d6b0a4558887bc1b333227f79535295b282ebe1eb2a1228e635f17fa1015a0",
		domainSeparator.String(),
	)

	typedDataHash, err := typedData.HashStruct(typedData.PrimaryType, typedData.Message)
	assert.NoError(t, err)
	assert.Equal(
		t,
		"0xd1d57c2e92c65952c77296afa21cffbe559486a20b1275d6196d8812b8a60e27",
		typedDataHash.String(),
	)

	hash, _, err := apitypes.TypedDataAndHash(typedData)
	assert.NoError(t, err)
	assert.Equal(
		t,
		"0xc9ca8abc30b3e71207f661b8292e83dbb7b1ef70119bac4aa409831ebde4b925",
		common.BytesToHash(hash).String(),
	)
}

func TestBridgeERC20Transfer_ToDepositConfirm(t *testing.T) {
	fromAddr := common.HexToAddress("0x1000000000000000000000000000000000000001")
	toAddr := common.HexToAddress("0x1000000000000000000000000000000000000009")
	tokenAddr := common.HexToAddress("0x2000000000000000000000000000000000000002")
	srcChainId := big.NewInt(97)
	sig := Signature{R: big.NewInt(1), S: big.NewInt(2), V: 28}

	transfer := &BridgeERC20Transfer{
		From:  fromAddr,
		To:    toAddr,
		Value: big.NewInt(1e18),
		Raw: types.Log{
			Address:     tokenAddr,
			BlockNumber: 42,
			TxHash:      common.HexToHash("0xaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"),
			Index:       3,
		},
	}

	got := transfer.ToDepositConfirm(srcChainId, sig)
	assert.Equal(t, EvmAddressToBytes32(fromAddr), got.User)
	assert.Equal(t, toAddr, got.Destination)
	assert.Equal(t, EvmAddressToBytes32(tokenAddr), got.Token)
	assert.Equal(t, big.NewInt(1e18), got.Amount)
	assert.Equal(t, srcChainId, got.ChainId)
	assert.Equal(t, uint64(42), got.BlockNumber)
	assert.Equal(t, transfer.Raw.TxHash, common.Hash(got.TxHash))
	assert.Equal(t, uint32(3), got.Index)
	assert.Equal(t, sig, got.Signature)
}
