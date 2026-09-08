package bridgehub

import (
	"context"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/require"

	"github.com/zakir-web3/bridge/internal/contract"
)

type stubBridgeContract struct {
	withdrawCalled bool
}

func (s *stubBridgeContract) Withdraw(_ context.Context, _ *big.Int, _ ...*contract.BridgeHubWithdraw) error {
	s.withdrawCalled = true
	return nil
}

func (s *stubBridgeContract) BridgeSignatureSubmitted(_ context.Context, _ *big.Int, _ ...*contract.MessageSignature) error {
	return nil
}

func TestProcessLog_WithdrawSkipsWhenNoHandlerRegistered(t *testing.T) {
	t.Parallel()

	hub := &BridgeHub{
		cfg:      Config{BridgeHubAddress: common.HexToAddress("0xcccccccccccccccccccccccccccccccccccccccc")},
		bridges:  make(map[uint64]BridgeContract),
		contract: mustNewBridgeHubFilterer(t),
	}

	user := common.HexToAddress("0x1000000000000000000000000000000000000001")
	userTopic := common.BytesToHash(common.LeftPadBytes(user.Bytes(), 32))

	log := types.Log{
		Address: hub.cfg.BridgeHubAddress,
		Topics: []common.Hash{
			common.HexToHash(WithdrawEventHash),
			common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000001"),
			userTopic,
			common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000002"),
		},
		Data: common.Hex2Bytes(
			"0000000000000000000000000000000000000000000000000000000000000003" + // destination
				"0000000000000000000000000000000000000000000000000000000000000004" + // amount
				"0000000000000000000000000000000000000000000000000000000000000065" + // chainId = 101
				"000000000000000000000000000000000000000000000000000000000000007b", // nonce = 123
		),
	}

	err := hub.ProcessLog(context.Background(), log)
	require.NoError(t, err)
}

func TestProcessLog_WithdrawRoutesByEventChainId(t *testing.T) {
	t.Parallel()

	hub := &BridgeHub{
		cfg:      Config{BridgeHubAddress: common.HexToAddress("0xcccccccccccccccccccccccccccccccccccccccc")},
		bridges:  make(map[uint64]BridgeContract),
		contract: mustNewBridgeHubFilterer(t),
	}
	stub := &stubBridgeContract{}
	hub.RegisterBridgeContract(101, stub)

	user := common.HexToAddress("0x1000000000000000000000000000000000000001")
	userTopic := common.BytesToHash(common.LeftPadBytes(user.Bytes(), 32))

	log := types.Log{
		Address: hub.cfg.BridgeHubAddress,
		Topics: []common.Hash{
			common.HexToHash(WithdrawEventHash),
			common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000001"),
			userTopic,
			common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000002"),
		},
		Data: common.Hex2Bytes(
			"0000000000000000000000000000000000000000000000000000000000000003" +
				"0000000000000000000000000000000000000000000000000000000000000004" +
				"0000000000000000000000000000000000000000000000000000000000000065" +
				"000000000000000000000000000000000000000000000000000000000000007b",
		),
	}

	err := hub.ProcessLog(context.Background(), log)
	require.NoError(t, err)
	require.True(t, stub.withdrawCalled)
}

func mustNewBridgeHubFilterer(t *testing.T) *contract.BridgeHub {
	t.Helper()
	c, err := contract.NewBridgeHub(common.Address{}, nil)
	require.NoError(t, err)
	return c
}
