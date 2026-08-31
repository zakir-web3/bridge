package solana

import (
	"encoding/base64"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gagliardetto/solana-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseDepositEvent(t *testing.T) {
	t.Parallel()

	// Fixture from solana e2e deposit (scripts/deposit.ts → deposit_event_b64).
	const eventB64 = "ePg9Ux+Oa5AKAKz8FumvQyIF6fqVdTDyt8qyW6XjkfwWA/jh6d77+DxEzd22qQD6K1hd0pngPRL6QpO8OyExv9MeZnwULDWNSc5PoghW8FaPUvrQrJGn/uEzYSNAQg8AAAAAAA=="
	data, err := base64.StdEncoding.DecodeString(eventB64)
	require.NoError(t, err)

	wantUser := solana.MustPublicKeyFromBase58("g3g8KT2YWpCiz34auNozvZRrqEF7ymgoEB7LoowTJsH")
	wantMint := solana.MustPublicKeyFromBase58("4ypQrxxfk8wCYWLSM7ppDxCaq8qAfg9f4PvM4z5M94E6")
	wantDestination := common.HexToAddress("0x3C44CdDdB6a900fa2b585dd299e03d12FA4293BC")
	const wantAmount uint64 = 1_000_000

	gotUser, gotDest, gotMint, amount, err := ParseDepositEvent(data)
	require.NoError(t, err)
	assert.True(t, gotUser.Equals(wantUser))
	assert.Equal(t, wantDestination, gotDest)
	assert.True(t, gotMint.Equals(wantMint))
	assert.Equal(t, wantAmount, amount)
}

func TestIsDepositInstruction(t *testing.T) {
	var dest [20]byte
	data := EncodeDepositInstructionData(dest, 1)
	assert.True(t, IsDepositInstruction(data))
}
