package solana

import (
	"testing"

	"github.com/gagliardetto/solana-go"
	"github.com/stretchr/testify/require"
)

func TestEncodeWithdrawInstructionData(t *testing.T) {
	user := [20]byte{0x10}
	destination := solana.NewWallet().PublicKey()
	mint := solana.MustPublicKeyFromBase58("EPjFWdd5AufqSSqeM2qN1xzybapC8G4wEGGkZwyTDt1v")

	params := WithdrawParams{
		User:        user,
		Destination: destination,
		Mint:        mint,
		Amount:      1_000_000,
		Nonce:       42,
		Signatures: []EcdsaSignature{
			{RecoveryID: 0},
			{RecoveryID: 1},
		},
	}
	params.Signatures[0].Sig[0] = 0xab
	params.Signatures[1].Sig[0] = 0xcd

	data, err := EncodeWithdrawInstructionData(params)
	require.NoError(t, err)
	require.True(t, IsWithdrawInstruction(data))
	require.Equal(t, withdrawInstructionDiscriminator, data[:8])
	require.Equal(t, byte(2), data[8+20+32+8+8])
}

func TestBuildWithdrawInstructionAccountCount(t *testing.T) {
	programID := solana.MustPublicKeyFromBase58("C4YxxrnCKnE4hVdTPcmTZN6yuHp5U9xVXRs3VanEeYfq")
	payer := solana.NewWallet().PublicKey()
	mint := solana.MustPublicKeyFromBase58("EPjFWdd5AufqSSqeM2qN1xzybapC8G4wEGGkZwyTDt1v")

	ix, err := BuildWithdrawInstruction(programID, payer, WithdrawParams{
		User:        [20]byte{1},
		Destination: solana.NewWallet().PublicKey(),
		Mint:        mint,
		Amount:      100,
		Nonce:       0,
		Signatures:  []EcdsaSignature{{RecoveryID: 0}},
	})
	require.NoError(t, err)
	require.Len(t, ix.Accounts(), 13)
}

func TestComputeUnitLimitForWithdraw(t *testing.T) {
	require.Equal(t, uint32(150_000), ComputeUnitLimitForWithdraw(2))
	require.Equal(t, uint32(480_000), ComputeUnitLimitForWithdraw(13))
}
