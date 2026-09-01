package solana

import (
	"math/big"
	"testing"
	"time"

	"github.com/gagliardetto/solana-go"
	"github.com/stretchr/testify/require"
)

func TestConfigValidate_DecodesPubkeys(t *testing.T) {
	cfg := &Config{
		SlotScannerConfig: SlotScannerConfig{
			Interval:     time.Second,
			SlotInterval: 100,
		},
		NodeURL:        "http://127.0.0.1:8899",
		ChainID:        big.NewInt(900001),
		ProgramIDStr:   "C4YxxrnCKnE4hVdTPcmTZN6yuHp5U9xVXRs3VanEeYfq",
		BridgeMintsStr: []string{"EPjFWdd5AufqSSqeM2qN1xzybapC8G4wEGGkZwyTDt1v"},
	}

	require.NoError(t, cfg.Validate())
	require.Equal(t, cfg.ProgramIDStr, cfg.ProgramID.String())
	require.Len(t, cfg.BridgeMints, 1)
	require.True(t, cfg.BridgeMints[0].Equals(solana.MustPublicKeyFromBase58(cfg.BridgeMintsStr[0])))
}
