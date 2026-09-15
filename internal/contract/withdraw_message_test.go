package contract

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseWithdrawMessageRawData(t *testing.T) {
	user := common.HexToAddress("0x1000000000000000000000000000000000000001")
	withdraw := &BridgeHubWithdraw{
		User:        user,
		Destination: AddressToBytes32(user),
		Token:       AddressToBytes32(common.HexToAddress("0x2000000000000000000000000000000000000002")),
		Amount:      big.NewInt(1e18),
		Nonce:       12345,
	}
	typedData := withdraw.ToTypedData(nil, big.NewInt(900001), common.HexToAddress("0xdb94ec3d773ea0a5b9b89b4bf28ed21dec3f0f8f"))
	raw, err := typedData.EncodeData(typedData.PrimaryType, typedData.Message, 1)
	require.NoError(t, err)

	gotUser, dest, token, amount, chainID, nonce, err := ParseWithdrawMessageRawData(raw)
	require.NoError(t, err)
	assert.Equal(t, user, gotUser)
	assert.Equal(t, AddressToBytes32(user), dest)
	assert.Equal(t, AddressToBytes32(common.HexToAddress("0x2000000000000000000000000000000000000002")), token)
	assert.Equal(t, big.NewInt(1e18), amount)
	assert.Equal(t, big.NewInt(900001), chainID)
	assert.Equal(t, uint64(12345), nonce)
}

func TestNonEmptyValidatorSignatures(t *testing.T) {
	sig := Signature{R: big.NewInt(1), S: big.NewInt(2), V: 27}
	got := NonEmptyValidatorSignatures([]Signature{{}, sig, {}})
	require.Len(t, got, 1)
	assert.Equal(t, sig, got[0])
}
