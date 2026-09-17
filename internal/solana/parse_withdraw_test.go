package solana

import (
	"encoding/base64"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gagliardetto/solana-go"
	"github.com/stretchr/testify/require"
)

func TestParseWithdrawFinalizedEvent(t *testing.T) {
	user := common.HexToAddress("0x1000000000000000000000000000000000000001")
	destination := solana.NewWallet().PublicKey()
	mint := solana.MustPublicKeyFromBase58("EPjFWdd5AufqSSqeM2qN1xzybapC8G4wEGGkZwyTDt1v")

	raw := make([]byte, 0, 140)
	raw = append(raw, withdrawFinalizedEventDiscriminator...)
	var message [32]byte
	message[0] = 0xaa
	raw = append(raw, message[:]...)
	raw = append(raw, user.Bytes()...)
	raw = append(raw, destination.Bytes()...)
	raw = append(raw, mint.Bytes()...)
	amountLE := []byte{0x00, 0x10, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00} // 4096 LE
	raw = append(raw, amountLE...)
	raw = append(raw, 0x2a, 0, 0, 0, 0, 0, 0, 0) // nonce 42 LE

	parsed, err := ParseWithdrawFinalizedEvent(raw)
	require.NoError(t, err)
	require.Equal(t, message, parsed.Message)
	require.Equal(t, user, parsed.User)
	require.Equal(t, destination.Bytes(), parsed.Destination[:])
	require.Equal(t, mint.Bytes(), parsed.Token[:])
	require.Equal(t, "4096", parsed.Amount.String())
	require.Equal(t, uint64(42), parsed.Nonce)

	logLine := "Program data: " + base64.StdEncoding.EncodeToString(raw)
	fromLogs, err := ParseWithdrawFinalizedFromLogs([]string{logLine})
	require.NoError(t, err)
	require.Equal(t, parsed.Nonce, fromLogs.Nonce)
}
