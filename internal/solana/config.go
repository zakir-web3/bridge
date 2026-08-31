package solana

import (
	"math/big"
	"time"

	"github.com/gagliardetto/solana-go"
	"github.com/pkg/errors"

	"github.com/zakir-web3/bridge/internal/evm"
)

type Config struct {
	NodeURL        string          `mapstructure:"node_url"        toml:"node_url"`
	RetryConfig    evm.RetryConfig `mapstructure:",squash"        toml:",squash"`
	Interval       time.Duration   `mapstructure:"interval"        toml:"interval"`
	ChainID        *big.Int        `mapstructure:"chain_id"        toml:"chain_id"`
	ProgramID      solana.PublicKey
	ProgramIDStr   string `mapstructure:"program_id"      toml:"program_id"`
	BridgeMints    []solana.PublicKey
	BridgeMintsStr []string `mapstructure:"bridge_mints"    toml:"bridge_mints"`
	StartSlot      uint64   `mapstructure:"start_slot"      toml:"start_slot"`
	SlotInterval   uint64   `mapstructure:"slot_interval"   toml:"slot_interval"`
	SlotDelay      uint64   `mapstructure:"slot_delay"      toml:"slot_delay"`
	ClearCache     bool     `mapstructure:"clear_cache"     toml:"clear_cache"`
}

func (c *Config) DecodePubkeys() error {
	if c.ProgramIDStr != "" {
		pub, err := solana.PublicKeyFromBase58(c.ProgramIDStr)
		if err != nil {
			return errors.Wrap(err, "program_id")
		}
		c.ProgramID = pub
	}
	c.BridgeMints = make([]solana.PublicKey, 0, len(c.BridgeMintsStr))
	for i, mint := range c.BridgeMintsStr {
		pub, err := solana.PublicKeyFromBase58(mint)
		if err != nil {
			return errors.Wrapf(err, "bridge_mints[%d]", i)
		}
		c.BridgeMints = append(c.BridgeMints, pub)
	}
	return nil
}

func (c Config) Enabled() bool {
	return c.ProgramIDStr != ""
}

func (c *Config) Validate() error {
	if !c.Enabled() {
		return nil
	}
	if c.NodeURL == "" {
		return errors.New("node_url is required")
	}
	if err := c.RetryConfig.Validate(); err != nil {
		return errors.Wrap(err, "retry_config")
	}
	if c.Interval == 0 {
		return errors.New("interval is required")
	}
	if c.ChainID == nil || c.ChainID.Sign() <= 0 {
		return errors.New("chain_id is required")
	}
	if c.ProgramIDStr == "" {
		return errors.New("program_id is required")
	}
	if len(c.BridgeMintsStr) == 0 {
		return errors.New("bridge_mints is required")
	}
	if c.SlotInterval == 0 {
		return errors.New("slot_interval is required")
	}
	return c.DecodePubkeys()
}

func (c Config) IsBridgeMint(mint solana.PublicKey) bool {
	for _, m := range c.BridgeMints {
		if m.Equals(mint) {
			return true
		}
	}
	return false
}
