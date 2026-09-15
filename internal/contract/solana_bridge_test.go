package contract

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gagliardetto/solana-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestVerifyingContractFromProgramID(t *testing.T) {
	programID := solana.MustPublicKeyFromBase58("C4YxxrnCKnE4hVdTPcmTZN6yuHp5U9xVXRs3VanEeYfq")
	got := VerifyingContractFromProgramID(programID)
	assert.Equal(t, common.HexToAddress("0xdb94ec3d773ea0a5b9b89b4bf28ed21dec3f0f8f"), got)
}

func TestSolanaBridgeDomainSeparator(t *testing.T) {
	programID := solana.MustPublicKeyFromBase58("C4YxxrnCKnE4hVdTPcmTZN6yuHp5U9xVXRs3VanEeYfq")
	separator, err := SolanaBridgeDomainSeparator(big.NewInt(900001), programID)
	require.NoError(t, err)
	assert.Equal(t, "0xe038ca293e650b49e9781d6f45d165e6ac0f202e3e6d4e00c07072d8088c3633", separator.Hex())
}
