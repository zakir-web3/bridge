package main

import (
	"crypto/ecdsa"
	"math/big"
	"reflect"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/go-viper/mapstructure/v2"
	"github.com/pelletier/go-toml/v2"
	"github.com/pkg/errors"

	"github.com/zakir-web3/bridge/internal/bridge"
	"github.com/zakir-web3/bridge/internal/bridgehub"
	"github.com/zakir-web3/bridge/internal/evm"
)

type Config struct {
	LogLevel  string            `mapstructure:"log_level"  toml:"log_level"`
	LogFormat string            `mapstructure:"log_format" toml:"log_format"`
	Source    string            `mapstructure:"source"     toml:"source"`
	PrivKey   *ecdsa.PrivateKey `mapstructure:"priv_key"   toml:"priv_key"`
	Bridge    bridge.Config     `mapstructure:"bridge"     toml:"bridge"`
	BridgeHub bridgehub.Config  `mapstructure:"bridge_hub" toml:"bridge_hub"`
}

func (c *Config) String() string {
	copyCfg := *c
	if copyCfg.PrivKey != nil && copyCfg.PrivKey.D != nil {
		copyCfg.PrivKey = &ecdsa.PrivateKey{}
	}
	if copyCfg.Bridge.AccConfig.PrivKey != nil && copyCfg.Bridge.AccConfig.PrivKey.D != nil {
		copyCfg.Bridge.AccConfig.PrivKey = &ecdsa.PrivateKey{}
	}
	if copyCfg.BridgeHub.AccConfig.PrivKey != nil && copyCfg.BridgeHub.AccConfig.PrivKey.D != nil {
		copyCfg.BridgeHub.AccConfig.PrivKey = &ecdsa.PrivateKey{}
	}
	data, _ := toml.Marshal(copyCfg)
	return string(data)
}

func (c *Config) Validate() error {
	if c.LogLevel == "" {
		return errors.New("log_level is required")
	}
	if c.LogFormat == "" {
		return errors.New("log_format is required")
	}
	if c.Source == "" {
		return errors.New("source is required")
	}
	if err := c.Bridge.Validate(); err != nil {
		return errors.Wrap(err, "bridge")
	}
	if err := c.BridgeHub.Validate(); err != nil {
		return errors.Wrap(err, "bridge_hub")
	}
	return nil
}

func (c *Config) MergeConfig() {
	c.mergePrivKeyConfig(&c.Bridge.AccConfig)
	c.mergePrivKeyConfig(&c.BridgeHub.AccConfig)
}

func (c *Config) mergePrivKeyConfig(moduleConfig *evm.Config) {
	if (moduleConfig.PrivKey == nil || moduleConfig.PrivKey.D == nil) &&
		c.PrivKey != nil && c.PrivKey.D != nil {
		moduleConfig.PrivKey = c.PrivKey
	}
}

func viperDecoderOption(config *mapstructure.DecoderConfig) {
	config.DecodeHook = mapstructure.ComposeDecodeHookFunc(
		mapstructure.StringToTimeDurationHookFunc(),
		mapstructure.StringToSliceHookFunc(","),
		StringToBigIntHookFunc(),
		StringToAddressHookFunc(),
		StringToEcdsaPrivateKeyHookFunc(),
	)
}

func StringToBigIntHookFunc() mapstructure.DecodeHookFunc {
	return func(f, t reflect.Type, data any) (any, error) {
		if f.Kind() != reflect.String {
			return data, nil
		}
		if t != reflect.TypeOf(big.Int{}) {
			return data, nil
		}
		i, ok := new(big.Int).SetString(data.(string), 10)
		if !ok {
			return nil, errors.Errorf("can't convert %v to big.Int", data)
		}
		return i, nil
	}
}

func StringToAddressHookFunc() mapstructure.DecodeHookFunc {
	return func(f, t reflect.Type, data any) (any, error) {
		if f.Kind() != reflect.String {
			return data, nil
		}
		if t != reflect.TypeOf(common.Address{}) {
			return data, nil
		}
		return common.HexToAddress(data.(string)), nil
	}
}

func StringToEcdsaPrivateKeyHookFunc() mapstructure.DecodeHookFunc {
	return func(f, t reflect.Type, data any) (any, error) {
		if f.Kind() != reflect.String {
			return data, nil
		}
		if t != reflect.TypeOf(&ecdsa.PrivateKey{}) {
			return data, nil
		}
		keyStr := data.(string)
		if keyStr == "" {
			return &ecdsa.PrivateKey{}, nil
		}
		if len(keyStr) >= 2 && keyStr[:2] == "0x" {
			keyStr = keyStr[2:]
		}
		privateKey, err := crypto.HexToECDSA(keyStr)
		if err != nil {
			return nil, errors.Wrapf(err, "can't convert %v to ecdsa.PrivateKey", data)
		}
		return privateKey, nil
	}
}
