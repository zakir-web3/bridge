package scanner

import (
	"time"

	"github.com/pkg/errors"

	"github.com/zakir-web3/bridge/internal/evm"
)

type Config struct {
	NodeURL       string          `mapstructure:"node_url"       toml:"node_url"`
	RetryConfig   evm.RetryConfig `mapstructure:",squash" toml:",squash"`
	Interval      time.Duration   `mapstructure:"interval"       toml:"interval"`
	StartBlock    uint64          `mapstructure:"start_block"    toml:"start_block"`
	BlockInterval uint64          `mapstructure:"block_interval" toml:"block_interval"`
	BlockDelay    uint64          `mapstructure:"block_delay"    toml:"block_delay"`
	ClearCache    bool            `mapstructure:"clear_cache"    toml:"clear_cache"`
}

func (c Config) Validate() error {
	if c.NodeURL == "" {
		return errors.New("node_url is required")
	}
	if err := c.RetryConfig.Validate(); err != nil {
		return errors.Wrap(err, "retry_config")
	}
	if c.Interval == 0 {
		return errors.New("interval is required")
	}
	if c.StartBlock == 0 {
		return errors.New("start_block is required")
	}
	if c.BlockInterval == 0 {
		return errors.New("block_interval is required")
	}
	return nil
}
