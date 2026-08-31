package bridge

import (
	"errors"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"

	"github.com/zakir-web3/bridge/internal/evm"
	"github.com/zakir-web3/bridge/internal/scanner"
)

type Config struct {
	scanner.Config              `mapstructure:",squash" toml:",squash"`
	BridgeAddress               common.Address   `mapstructure:"bridge_address" toml:"bridge_address"`
	BridgeTokens                []common.Address `mapstructure:"bridge_tokens"  toml:"bridge_tokens"`
	AccConfig                   evm.Config       `mapstructure:",squash"        toml:",squash"`
	SendFinalizeWithdrawals     bool             `mapstructure:"send_finalize_withdrawals" toml:"send_finalize_withdrawals"`
	FinalizeWithdrawalsInterval time.Duration    `mapstructure:"finalize_withdrawals_interval" toml:"finalize_withdrawals_interval"`
	MaxWithdrawAmount           *big.Int         `mapstructure:"max_withdraw_amount" toml:"max_withdraw_amount"`
	RequestWithdrawGasLimit     uint64           `mapstructure:"request_withdraw_gas_limit" toml:"request_withdraw_gas_limit"`
	EnableRequestWithdraw       bool             `mapstructure:"enable_request_withdraw" toml:"enable_request_withdraw"`
}

func (c Config) Validate() error {
	if !c.Enabled() {
		return nil
	}
	if err := c.Config.Validate(); err != nil {
		return err
	}
	if c.BridgeAddress == (common.Address{}) {
		return errors.New("bridge_address is required")
	}
	if len(c.BridgeTokens) == 0 {
		return errors.New("bridge_tokens is required")
	}
	for _, token := range c.BridgeTokens {
		if token == (common.Address{}) {
			return errors.New("bridge_tokens cannot contain empty address")
		}
	}
	if err := c.AccConfig.Validate(); err != nil {
		return err
	}
	if c.SendFinalizeWithdrawals {
		if c.FinalizeWithdrawalsInterval == 0 {
			return errors.New("finalize_withdrawals_interval must be set when send_finalize_withdrawals is true")
		}
		if c.MaxWithdrawAmount == nil || c.MaxWithdrawAmount.Sign() <= 0 {
			return errors.New("max_withdraw_amount must be set and greater than zero when send_finalize_withdrawals is true")
		}
	}
	return nil
}

func (c Config) Enabled() bool {
	return c.BridgeAddress != (common.Address{})
}

func (c Config) IsBridgeToken(token common.Address) bool {
	for _, t := range c.BridgeTokens {
		if t == token {
			return true
		}
	}
	return false
}
