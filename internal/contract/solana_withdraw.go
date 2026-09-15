package contract

import (
	"math/big"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
)

// SolanaWithdrawFinalized represents a parsed Solana WithdrawFinalized event for hub confirm.
type SolanaWithdrawFinalized struct {
	Message     [32]byte
	User        common.Address
	Destination [32]byte
	Token       [32]byte
	Amount      *big.Int
	Nonce       uint64
}

// ToBridgeFinalizedWithdrawal converts the Solana event into the shared hub-confirm shape.
func (w *SolanaWithdrawFinalized) ToBridgeFinalizedWithdrawal() *BridgeFinalizedWithdrawal {
	return &BridgeFinalizedWithdrawal{
		Message:     w.Message,
		User:        w.User,
		Destination: w.Destination,
		Token:       w.Token,
		Amount:      w.Amount,
		Nonce:       w.Nonce,
	}
}

// ToTypedData builds BridgeHub-domain Withdraw typed data for withdrawConfirm.
func (w *SolanaWithdrawFinalized) ToTypedData(srcChainID, hubChainID *big.Int, verifyingContract common.Address) apitypes.TypedData {
	return apitypes.TypedData{
		Domain: NewBridgeHubDomain(hubChainID, verifyingContract),
		Types: apitypes.Types{
			"EIP712Domain": {
				{Name: "name", Type: "string"},
				{Name: "version", Type: "string"},
				{Name: "chainId", Type: "uint256"},
				{Name: "verifyingContract", Type: "address"},
			},
			"Withdraw": {
				{Name: "user", Type: "address"},
				{Name: "destination", Type: "bytes32"},
				{Name: "token", Type: "bytes32"},
				{Name: "amount", Type: "uint256"},
				{Name: "chainId", Type: "uint256"},
				{Name: "nonce", Type: "uint64"},
			},
		},
		PrimaryType: "Withdraw",
		Message: apitypes.TypedDataMessage{
			"user":        w.User.Hex(),
			"destination": Bytes32Hex(w.Destination),
			"token":       Bytes32Hex(w.Token),
			"amount":      w.Amount.String(),
			"chainId":     srcChainID.String(),
			"nonce":       strconv.FormatUint(w.Nonce, 10),
		},
	}
}

// ParseWithdrawMessageRawData decodes Bridge-domain withdraw message bytes without applying
// EVM empty-signature padding (required for Solana quorum verification).
func ParseWithdrawMessageRawData(rawData []byte) (
	user common.Address,
	destination [32]byte,
	token [32]byte,
	amount *big.Int,
	chainID *big.Int,
	nonce uint64,
	err error,
) {
	if len(rawData) < 32+192 {
		err = ErrWithdrawRawDataTooShort
		return
	}
	encoded := rawData[32:]

	user = common.BytesToAddress(encoded[0:32][12:])
	copy(destination[:], encoded[32:64])
	copy(token[:], encoded[64:96])
	amount = new(big.Int).SetBytes(encoded[96:128])
	chainID = new(big.Int).SetBytes(encoded[128:160])
	nonce = new(big.Int).SetBytes(encoded[160:192]).Uint64()
	return
}

// NonEmptyValidatorSignatures returns only non-empty signatures from a hub quorum payload.
// Solana withdraw must not reuse ToWithdrawalRequest padding semantics.
func NonEmptyValidatorSignatures(signatures []Signature) []Signature {
	out := make([]Signature, 0, len(signatures))
	for _, sig := range signatures {
		if !sig.IsEmpty() {
			out = append(out, sig)
		}
	}
	return out
}
