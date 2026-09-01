package solana

import (
	"math/big"

	"github.com/gagliardetto/solana-go"
	"github.com/pkg/errors"

	"github.com/zakir-web3/bridge/internal/evm"
)

type Config struct {
	SlotScannerConfig `mapstructure:",squash"        toml:",squash"`
	NodeURL           string          `mapstructure:"node_url"        toml:"node_url"`
	RetryConfig       evm.RetryConfig `mapstructure:",squash"        toml:",squash"`
	ChainID           *big.Int        `mapstructure:"chain_id"        toml:"chain_id"`
	ProgramID         solana.PublicKey
	ProgramIDStr      string `mapstructure:"program_id"      toml:"program_id"`
	BridgeMints       []solana.PublicKey
	BridgeMintsStr    []string `mapstructure:"bridge_mints"    toml:"bridge_mints"`
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
	if err := c.SlotScannerConfig.Validate(); err != nil {
		return err
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
