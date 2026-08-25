package contract

import (
	"context"
	"math/big"
	"strconv"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
	"github.com/pkg/errors"
)

func NewBridgeDomain(chainId *big.Int, verifyingContract common.Address) apitypes.TypedDataDomain {
	return apitypes.TypedDataDomain{
		Name:              "Bridge",
		Version:           "1",
		ChainId:           math.NewHexOrDecimal256(chainId.Int64()),
		VerifyingContract: verifyingContract.Hex(),
	}
}

func (w *BridgeFinalizedWithdrawal) ToTypedData(srcChainId, chainId *big.Int, verifyingContract common.Address) apitypes.TypedData {
	return apitypes.TypedData{
		Domain: NewBridgeHubDomain(chainId, verifyingContract),
		Types: apitypes.Types{
			"EIP712Domain": {
				{Name: "name", Type: "string"},
				{Name: "version", Type: "string"},
				{Name: "chainId", Type: "uint256"},
				{Name: "verifyingContract", Type: "address"},
			},
			"Withdraw": {
				{Name: "user", Type: "address"},
				{Name: "destination", Type: "address"},
				{Name: "token", Type: "address"},
				{Name: "amount", Type: "uint256"},
				{Name: "chainId", Type: "uint256"},
				{Name: "nonce", Type: "uint64"},
			},
		},
		PrimaryType: "Withdraw",
		Message: apitypes.TypedDataMessage{
			"user":        w.User.String(),
			"destination": w.Destination.String(),
			"token":       w.Token.String(),
			"amount":      w.Amount.String(),
			"chainId":     srcChainId.String(),
			"nonce":       strconv.FormatUint(w.Nonce, 10),
		},
	}
}

func (w *BridgeFinalizedWithdrawal) ToWithdrawConfirm(chainId *big.Int, signature Signature) WithdrawConfirm {
	return WithdrawConfirm{
		User:        w.User,
		Destination: w.Destination,
		Token:       w.Token,
		Amount:      w.Amount,
		ChainId:     chainId,
		Nonce:       w.Nonce,
		Signature:   signature,
	}
}

func (b *BridgeCaller) GetDomainSeparator(ctx context.Context) (common.Hash, error) {
	hash, err := b.DomainSeparator(&bind.CallOpts{Context: ctx})
	if err != nil {
		return common.Hash{}, errors.Wrap(err, "get domain separator")
	}
	return hash, nil
}

func (b *BridgeCaller) GetHotValidatorSetHash(ctx context.Context) (common.Hash, error) {
	hash, err := b.HotValidatorSetHash(&bind.CallOpts{Context: ctx})
	if err != nil {
		return common.Hash{}, errors.Wrap(err, "get hot validator set hash")
	}
	return hash, nil
}

func (b *BridgeCaller) GetRequestedWithdrawals(ctx context.Context, messageHash [32]byte) (*BridgeRequestedWithdrawal, error) {
	requestedWithdrawals, err := b.RequestedWithdrawals(&bind.CallOpts{Context: ctx}, messageHash)
	if err != nil {
		return nil, errors.Wrap(err, "get requested withdrawals")
	}
	return &BridgeRequestedWithdrawal{
		User:                 requestedWithdrawals.User,
		Destination:          requestedWithdrawals.Destination,
		Token:                requestedWithdrawals.Token,
		Amount:               requestedWithdrawals.Amount,
		Nonce:                requestedWithdrawals.Nonce,
		RequestedTime:        requestedWithdrawals.RequestedTime,
		RequestedBlockNumber: requestedWithdrawals.RequestedBlockNumber,
		Message:              requestedWithdrawals.Message,
	}, nil
}

func (b *BridgeCaller) GetDisputePeriod(ctx context.Context) (uint64, error) {
	result, err := b.DisputePeriodSeconds(&bind.CallOpts{Context: ctx})
	if err != nil {
		return 0, errors.Wrap(err, "get dispute period")
	}
	return result, nil
}

func (b *BridgeCaller) IsFinalizeWithdrawals(ctx context.Context, messageHash [32]byte) (bool, error) {
	finalizedWithdrawal, err := b.FinalizedWithdrawals(&bind.CallOpts{Context: ctx}, messageHash)
	if err != nil {
		return false, errors.Wrap(err, "get finalized withdrawals")
	}
	return finalizedWithdrawal, nil
}
