package evm

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/require"
)

func validConfig(t *testing.T) *Config {
	t.Helper()
	key, err := crypto.GenerateKey()
	require.NoError(t, err)

	return &Config{
		ChainId:                     big.NewInt(1),
		PrivKey:                     key,
		FeeHistoryBlockCount:        1,
		FeeHistoryRewardPercentiles: []float64{50},
		MaxGasTipCap:                big.NewInt(1),
		MaxGasFeeCap:                big.NewInt(2),
		IncreasePercentile:          big.NewInt(120),
	}
}

func TestConfigValidate_RequireIncreasePercentile(t *testing.T) {
	t.Parallel()

	cfg := validConfig(t)
	cfg.IncreasePercentile = nil

	err := cfg.Validate()
	require.Error(t, err)
	require.Contains(t, err.Error(), "increase_percentile is required")
}
