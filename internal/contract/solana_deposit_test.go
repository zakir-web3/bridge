package contract

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSolanaDeposit_ToTypedData(t *testing.T) {
	var user, token [32]byte
	copy(user[:], []byte{1, 2, 3})
	copy(token[31:], []byte{0xab})

	deposit := &SolanaDeposit{
		User:        user,
		Destination: common.HexToAddress("0x1234567890123456789012345678901234567890"),
		Token:       token,
		Amount:      big.NewInt(100_000),
		ChainID:     big.NewInt(900001),
		BlockNumber: 12345,
		TxHash:      common.HexToHash("0xaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"),
		Index:       2,
	}

	typedData := deposit.ToTypedData(deposit.ChainID, big.NewInt(1337), common.HexToAddress("0x3000000000000000000000000000000000000003"))
	require.Equal(t, "Deposit", typedData.PrimaryType)

	hash, err := typedData.HashStruct(typedData.PrimaryType, typedData.Message)
	require.NoError(t, err)
	assert.NotEmpty(t, hash)
}

func TestSolanaDeposit_ToDepositConfirm(t *testing.T) {
	t.Parallel()

	var user, token [32]byte
	user[0] = 1
	token[31] = 0xab

	deposit := &SolanaDeposit{
		User:        user,
		Destination: common.HexToAddress("0x1234567890123456789012345678901234567890"),
		Token:       token,
		Amount:      big.NewInt(100_000),
		ChainID:     big.NewInt(900001),
		BlockNumber: 42,
		TxHash:      common.HexToHash("0xaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"),
		Index:       3,
	}

	confirm := deposit.ToDepositConfirm(Signature{V: 27})
	require.Equal(t, deposit.User, confirm.User)
	require.Equal(t, deposit.Destination, confirm.Destination)
	require.Equal(t, deposit.Token, confirm.Token)
	require.Equal(t, deposit.Amount, confirm.Amount)
	require.Equal(t, deposit.ChainID, confirm.ChainId)
	require.Equal(t, deposit.BlockNumber, confirm.BlockNumber)
	require.Equal(t, deposit.TxHash, confirm.TxHash)
	require.Equal(t, deposit.Index, confirm.Index)
	require.Equal(t, Signature{V: 27}, confirm.Signature)
}
