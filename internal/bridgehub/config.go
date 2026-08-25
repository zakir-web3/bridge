package bridgehub

import (
	"errors"

	"github.com/ethereum/go-ethereum/common"

	"github.com/zakir-web3/bridge/internal/evm"
	"github.com/zakir-web3/bridge/internal/scanner"
)

type Config struct {
	scanner.Config   `mapstructure:",squash" toml:",squash"`
	BridgeHubAddress common.Address `mapstructure:"bridge_hub_address" toml:"bridge_hub_address"`
	AccConfig        evm.Config     `mapstructure:",squash"            toml:",squash"`
}

func (c Config) Validate() error {
	if err := c.Config.Validate(); err != nil {
		return err
	}
	if c.BridgeHubAddress == (common.Address{}) {
		return errors.New("bridge_hub_address is required")
	}
	if err := c.AccConfig.Validate(); err != nil {
		return err
	}
	return nil
}
