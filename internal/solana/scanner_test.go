package solana

import (
	"testing"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/stretchr/testify/require"

	"github.com/zakir-web3/bridge/internal/cache"
)

func TestShouldSkipSignature(t *testing.T) {
	var sigA, sigB solana.Signature
	sigA[0] = 1
	sigB[0] = 2

	checkpoint := cache.SolanaCheckpoint{Slot: 100, Signature: sigA}

	require.True(t, shouldSkipSignature(&rpc.TransactionSignature{Slot: 99, Signature: sigB}, checkpoint))
	require.True(t, shouldSkipSignature(&rpc.TransactionSignature{Slot: 100, Signature: sigA}, checkpoint))
	require.False(t, shouldSkipSignature(&rpc.TransactionSignature{Slot: 100, Signature: sigB}, checkpoint))
	require.False(t, shouldSkipSignature(&rpc.TransactionSignature{Slot: 101, Signature: sigA}, checkpoint))
}

func TestSlotScannerConfigApplyDefaults(t *testing.T) {
	cfg := SlotScannerConfig{Interval: 1, SlotInterval: 1}
	cfg.applyDefaults()
	require.Equal(t, uint64(defaultSlotLookback), cfg.SlotLookback)
	require.Equal(t, uint64(defaultRescanSlots), cfg.RescanSlots)
	require.Equal(t, defaultSignaturesPageLimit, cfg.SignaturesPageLimit)
	require.Equal(t, defaultTxFetchRetries, cfg.TxFetchRetries)
}
