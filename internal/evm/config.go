package evm

import (
	"crypto/ecdsa"
	"math/big"

	"github.com/pkg/errors"
)

type Config struct {
	ChainId *big.Int          `mapstructure:"chain_id"                       toml:"chain_id"`
	PrivKey *ecdsa.PrivateKey `mapstructure:"priv_key"                       toml:"priv_key"`

	FeeHistoryBlockCount        uint64    `mapstructure:"fee_history_block_count"        toml:"fee_history_block_count"`
	FeeHistoryRewardPercentiles []float64 `mapstructure:"fee_history_reward_percentiles" toml:"fee_history_reward_percentiles"`
	MaxGasTipCap                *big.Int  `mapstructure:"max_gas_tip_cap"                toml:"max_gas_tip_cap"`
	MaxGasFeeCap                *big.Int  `mapstructure:"max_gas_fee_cap"                toml:"max_gas_fee_cap"`
	IncreasePercentile          *big.Int  `mapstructure:"increase_percentile"            toml:"increase_percentile"`

	MaxGasLimit uint64 `mapstructure:"max_gas_limit"                  toml:"max_gas_limit"`
	NoSend      bool   `mapstructure:"no_send"                        toml:"no_send"`

	GasTipCap *big.Int `mapstructure:"gas_tip_cap" toml:"gas_tip_cap"`
	GasFeeCap *big.Int `mapstructure:"gas_fee_cap" toml:"gas_fee_cap"`
}

func (c *Config) Validate() error {
	if c.ChainId == nil {
		return errors.New("chain_id is required")
	}
	if c.PrivKey == nil || c.PrivKey.D == nil {
		return errors.New("priv_key is required")
	}
	if c.FeeHistoryBlockCount < 1 || c.FeeHistoryBlockCount > 1024 {
		return errors.New("fee_history_block_count must be between 1 and 1024")
	}
	if len(c.FeeHistoryRewardPercentiles) == 0 {
		return errors.New("fee_history_reward_percentiles is required")
	}
	for _, percentile := range c.FeeHistoryRewardPercentiles {
		if percentile < 0 || percentile > 100 {
			return errors.New("fee_history_reward_percentiles must be between 0 and 100")
		}
	}
	if c.MaxGasTipCap == nil {
		return errors.New("max_gas_tip_cap is required")
	}
	if c.MaxGasFeeCap == nil {
		return errors.New("max_gas_fee_cap is required")
	}
	if c.MaxGasFeeCap.Cmp(c.MaxGasTipCap) < 0 {
		return errors.New("max_gas_fee_cap must be greater than or equal to max_gas_tip_cap")
	}
	if c.IncreasePercentile == nil {
		return errors.New("increase_percentile is required")
	}
	if c.IncreasePercentile.Cmp(big.NewInt(100)) < 0 ||
		c.IncreasePercentile.Cmp(big.NewInt(200)) > 0 {
		return errors.New("increase_percentile must be between 100 and 200")
	}
	return nil
}
