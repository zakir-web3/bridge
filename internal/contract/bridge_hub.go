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

func NewBridgeHubDomain(chainId *big.Int, verifyingContract common.Address) apitypes.TypedDataDomain {
	return apitypes.TypedDataDomain{
		Name:              "BridgeHub",
		Version:           "1",
		ChainId:           math.NewHexOrDecimal256(chainId.Int64()),
		VerifyingContract: verifyingContract.Hex(),
	}
}

func (v *BridgeHubRequestedValidatorSetUpdate) ToTypedData(_, chainId *big.Int, verifyingContract common.Address) apitypes.TypedData {
	hot := make([]interface{}, len(v.HotAddresses))
	for i, addr := range v.HotAddresses {
		hot[i] = addr.Hex()
	}
	cold := make([]interface{}, len(v.ColdAddresses))
	for i, addr := range v.ColdAddresses {
		cold[i] = addr.Hex()
	}
	powers := make([]interface{}, len(v.Powers))
	for i, power := range v.Powers {
		powers[i] = strconv.FormatUint(power, 10)
	}
	return apitypes.TypedData{
		Domain: NewBridgeHubDomain(chainId, verifyingContract),
		Types: apitypes.Types{
			"EIP712Domain": {
				{Name: "name", Type: "string"},
				{Name: "version", Type: "string"},
				{Name: "chainId", Type: "uint256"},
				{Name: "verifyingContract", Type: "address"},
			},
			"UpdateValidatorSet": {
				{Name: "epoch", Type: "uint64"},
				{Name: "hotAddresses", Type: "address[]"},
				{Name: "coldAddresses", Type: "address[]"},
				{Name: "powers", Type: "uint64[]"},
			},
		},
		PrimaryType: "UpdateValidatorSet",
		Message: apitypes.TypedDataMessage{
			"epoch":         strconv.FormatUint(v.NewEpoch, 10),
			"hotAddresses":  hot,
			"coldAddresses": cold,
			"powers":        powers,
		},
	}
}

func (w *BridgeHubWithdraw) ToTypedData(_, chainId *big.Int, verifyingContract common.Address) apitypes.TypedData {
	return apitypes.TypedData{
		Domain: NewBridgeDomain(chainId, verifyingContract),
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
			"chainId":     chainId.String(),
			"nonce":       strconv.FormatUint(w.Nonce, 10),
		},
	}
}

func (v *ValidatorSet) IsValidator(validator common.Address) bool {
	for i := 0; i < len(v.Validators); i++ {
		if v.Validators[i] == validator {
			return true
		}
	}
	return false
}

func (v *ValidatorSet) ValidatorsStr() []string {
	validators := make([]string, len(v.Validators))
	for i, v := range v.Validators {
		validators[i] = v.Hex()
	}
	return validators
}

type MessageSignature struct {
	TotalPower uint64
	Signers    []common.Address
	Signatures []Signature
	RawData    []byte
}

func (m *MessageSignature) ToWithdrawalRequest(validatorSet ValidatorSet) (WithdrawalRequest, error) {
	// RawData format: abi.encode(WITHDRAW_TYPEHASH, user, destination, token, amount, chainId, nonce)
	// Skip first 32 bytes (typehash) and decode the remaining data
	if len(m.RawData) < 32 {
		return WithdrawalRequest{}, errors.New("raw data too short")
	}

	// Skip the first 32 bytes (typehash) and decode the remaining data
	encodedData := m.RawData[32:]

	// ABI encoding rules:
	// - address: 32 bytes (padded to the right)
	// - uint256: 32 bytes
	// - uint64: 32 bytes (padded to the right)
	// Total expected length: 6 * 32 = 192 bytes

	if len(encodedData) < 192 {
		return WithdrawalRequest{}, errors.New("encoded data too short, expected 192 bytes")
	}

	// Parse user address (first 32 bytes)
	userBytes := encodedData[0:32]
	user := common.BytesToAddress(userBytes[12:]) // address is right-padded, take last 20 bytes

	// Parse destination address (next 32 bytes)
	destBytes := encodedData[32:64]
	destination := common.BytesToAddress(destBytes[12:]) // address is right-padded, take last 20 bytes

	// Parse token address (next 32 bytes)
	tokenBytes := encodedData[64:96]
	token := common.BytesToAddress(tokenBytes[12:]) // address is right-padded, take last 20 bytes

	// Parse amount (next 32 bytes)
	amountBytes := encodedData[96:128]
	amount := new(big.Int).SetBytes(amountBytes)

	// Parse chainId (next 32 bytes)
	chainIdBytes := encodedData[128:160]
	chainId := new(big.Int).SetBytes(chainIdBytes)

	// Parse nonce (last 32 bytes)
	nonceBytes := encodedData[160:192]
	nonce := new(big.Int).SetBytes(nonceBytes).Uint64()

	for i, signer := range m.Signers {
		if validatorSet.Validators[i] != signer {
			return WithdrawalRequest{}, errors.New("signers do not match validator set")
		}
	}
	signedIndex := -1
	for i := 0; i < len(m.Signatures); i++ {
		if !m.Signatures[i].IsEmpty() {
			signedIndex = i
			break
		}
	}
	for i := 0; i < len(m.Signatures); i++ {
		if m.Signatures[i].IsEmpty() {
			m.Signatures[i] = m.Signatures[signedIndex]
		}
	}

	return WithdrawalRequest{
		User:        user,
		Destination: destination,
		Token:       token,
		Amount:      amount,
		ChainId:     chainId,
		Nonce:       nonce,
		Signatures:  m.Signatures,
	}, nil
}

func (b *BridgeHubCaller) GetValidatorSet(ctx context.Context) (*ValidatorSet, error) {
	validators, err := b.GetValidators(&bind.CallOpts{Context: ctx})
	if err != nil {
		return nil, errors.Wrap(err, "get validators")
	}
	return &ValidatorSet{
		Validators: validators.Validators,
		Powers:     validators.Powers,
		Epoch:      validators.Epoch,
	}, nil
}

func (b *BridgeHubCaller) GetDomainSeparator(ctx context.Context) (common.Hash, error) {
	hash, err := b.DomainSeparator(&bind.CallOpts{Context: ctx})
	if err != nil {
		return common.Hash{}, errors.Wrap(err, "get domain separator")
	}
	return hash, nil
}

func (b *BridgeHubCaller) GetMsgSign(ctx context.Context, messageHash common.Hash, signer common.Address) (*Signature, error) {
	signed, err := b.GetValidatorSignature(&bind.CallOpts{Context: ctx}, messageHash, signer)
	if err != nil {
		return nil, errors.Wrap(err, "get message signature")
	}
	return &signed, nil
}

func (b *BridgeHubCaller) IsMsgProcessed(ctx context.Context, messageHash common.Hash) (bool, error) {
	processed, err := b.ProcessedMessages(&bind.CallOpts{Context: ctx}, messageHash)
	if err != nil {
		return false, errors.Wrap(err, "check if message is processed")
	}
	return processed, nil
}

func (b *BridgeHubCaller) GetBridgeMsgSign(ctx context.Context, messageHash common.Hash, signer common.Address) (*Signature, error) {
	messageSignature, err := b.GetBridgeValidatorSignature(&bind.CallOpts{Context: ctx}, messageHash, signer)
	if err != nil {
		return nil, errors.Wrap(err, "get bridge message signature")
	}
	return &messageSignature, nil
}

func (b *BridgeHubCaller) GetBridgeMsgSigns(ctx context.Context, messageHash common.Hash) (*MessageSignature, error) {
	messageSignature, err := b.GetBridgeMessageSignatures(&bind.CallOpts{Context: ctx}, messageHash)
	if err != nil {
		return nil, errors.Wrap(err, "get bridge message signatures")
	}
	return &MessageSignature{
		TotalPower: messageSignature.TotalPower,
		Signers:    messageSignature.Signers,
		Signatures: messageSignature.Signatures,
		RawData:    messageSignature.RawData,
	}, nil
}
