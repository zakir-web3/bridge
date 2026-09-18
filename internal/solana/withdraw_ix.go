package solana

import (
	"encoding/binary"
	"math/big"

	"github.com/gagliardetto/solana-go"
	"github.com/pkg/errors"

	computebudget "github.com/gagliardetto/solana-go/programs/compute-budget"

	"github.com/zakir-web3/bridge/internal/contract"
)

const (
	defaultComputeUnitPrice = 1_000
	baseComputeUnits        = 90_000
	computeUnitsPerSig      = 30_000
)

// EcdsaSignature is the on-chain Anchor representation of a validator ECDSA signature.
type EcdsaSignature struct {
	Sig        [64]byte
	RecoveryID uint8
}

// WithdrawParams holds decoded withdraw message fields for Solana submission.
type WithdrawParams struct {
	User        [20]byte
	Destination solana.PublicKey
	Mint        solana.PublicKey
	Amount      uint64
	Nonce       uint64
	Signatures  []EcdsaSignature
}

// EncodeWithdrawInstructionData serializes the Anchor withdraw instruction arguments.
func EncodeWithdrawInstructionData(params WithdrawParams) ([]byte, error) {
	data := make([]byte, 0, 8+20+32+8+8+4+len(params.Signatures)*65)
	data = append(data, withdrawInstructionDiscriminator...)
	data = append(data, params.User[:]...)
	data = append(data, params.Destination.Bytes()...)

	amountBuf := make([]byte, 8)
	binary.LittleEndian.PutUint64(amountBuf, params.Amount)
	data = append(data, amountBuf...)

	nonceBuf := make([]byte, 8)
	binary.LittleEndian.PutUint64(nonceBuf, params.Nonce)
	data = append(data, nonceBuf...)

	sigCount := uint32(len(params.Signatures)) //nolint:gosec // G115: sig count bounded by validator set size
	sigCountBuf := make([]byte, 4)
	binary.LittleEndian.PutUint32(sigCountBuf, sigCount)
	data = append(data, sigCountBuf...)
	for _, sig := range params.Signatures {
		data = append(data, sig.Sig[:]...)
		data = append(data, sig.RecoveryID)
	}
	return data, nil
}

// SignatureToEcdsa converts an EVM-style validator signature into Solana program format.
func SignatureToEcdsa(sig contract.Signature) (EcdsaSignature, error) {
	if sig.IsEmpty() {
		return EcdsaSignature{}, errors.New("empty signature")
	}
	var out EcdsaSignature
	rBytes := sig.R.Bytes()
	copy(out.Sig[32-len(rBytes):32], rBytes)
	sBytes := sig.S.Bytes()
	copy(out.Sig[64-len(sBytes):64], sBytes)
	if sig.V < 27 {
		return EcdsaSignature{}, errors.Errorf("invalid signature v=%d", sig.V)
	}
	out.RecoveryID = sig.V - 27
	return out, nil
}

// BuildWithdrawInstruction assembles the withdraw instruction and required account metas.
func BuildWithdrawInstruction(
	programID solana.PublicKey,
	payer solana.PublicKey,
	params WithdrawParams,
) (solana.Instruction, error) {
	configPDA, _, err := findConfigPDA(programID)
	if err != nil {
		return nil, errors.Wrap(err, "config pda")
	}
	validatorSetPDA, _, err := findValidatorSetPDA(programID)
	if err != nil {
		return nil, errors.Wrap(err, "validator_set pda")
	}
	vaultStatePDA, _, err := findVaultStatePDA(params.Mint, programID)
	if err != nil {
		return nil, errors.Wrap(err, "vault_state pda")
	}
	vaultAuthorityPDA, _, err := findVaultAuthorityPDA(params.Mint, programID)
	if err != nil {
		return nil, errors.Wrap(err, "vault_authority pda")
	}
	vaultTokenAccount, err := findAssociatedTokenAddress(vaultAuthorityPDA, params.Mint)
	if err != nil {
		return nil, errors.Wrap(err, "vault token account")
	}
	destinationTokenAccount, err := findAssociatedTokenAddress(params.Destination, params.Mint)
	if err != nil {
		return nil, errors.Wrap(err, "destination token account")
	}
	noncePagePDA, _, err := findNoncePagePDA(params.Nonce, programID)
	if err != nil {
		return nil, errors.Wrap(err, "nonce_page pda")
	}

	data, err := EncodeWithdrawInstructionData(params)
	if err != nil {
		return nil, err
	}

	accounts := solana.AccountMetaSlice{
		solana.Meta(payer).WRITE().SIGNER(),
		solana.Meta(configPDA),
		solana.Meta(validatorSetPDA),
		solana.Meta(params.Mint),
		solana.Meta(vaultStatePDA),
		solana.Meta(vaultAuthorityPDA),
		solana.Meta(vaultTokenAccount).WRITE(),
		solana.Meta(params.Destination),
		solana.Meta(destinationTokenAccount).WRITE(),
		solana.Meta(noncePagePDA).WRITE(),
		solana.Meta(solana.TokenProgramID),
		solana.Meta(solana.SPLAssociatedTokenAccountProgramID),
		solana.Meta(solana.SystemProgramID),
	}

	return solana.NewInstruction(programID, accounts, data), nil
}

// ComputeUnitLimitForWithdraw estimates CU limit from validator signature count.
func ComputeUnitLimitForWithdraw(sigCount int) uint32 {
	if sigCount < 0 {
		sigCount = 0
	}
	limit := uint64(baseComputeUnits + computeUnitsPerSig*sigCount)
	if limit > uint64(^uint32(0)) {
		return ^uint32(0)
	}
	return uint32(limit)
}

// PrependComputeBudgetInstructions returns compute budget ixs for legacy/v0 transactions.
func PrependComputeBudgetInstructions(sigCount int, unitPriceMicroLamports uint64) []solana.Instruction {
	limit := ComputeUnitLimitForWithdraw(sigCount)
	price := unitPriceMicroLamports
	if price == 0 {
		price = defaultComputeUnitPrice
	}
	return []solana.Instruction{
		computebudget.NewSetComputeUnitLimitInstruction(limit).Build(),
		computebudget.NewSetComputeUnitPriceInstruction(price).Build(),
	}
}

// AmountToUint64 converts a hub withdraw amount to SPL u64, rejecting overflow.
func AmountToUint64(amount *big.Int) (uint64, error) {
	if amount == nil || amount.Sign() <= 0 {
		return 0, errors.New("amount must be positive")
	}
	if !amount.IsUint64() {
		return 0, errors.Errorf("amount %s exceeds uint64", amount.String())
	}
	return amount.Uint64(), nil
}

// TokenBytesToMint converts a Hub withdraw token bytes32 into an SPL mint pubkey.
func TokenBytesToMint(token [32]byte) solana.PublicKey {
	return solana.PublicKeyFromBytes(token[:])
}

// DestinationBytesToPubkey converts a Hub withdraw destination bytes32 into a Solana pubkey.
func DestinationBytesToPubkey(destination [32]byte) solana.PublicKey {
	return solana.PublicKeyFromBytes(destination[:])
}

// IsWithdrawInstruction reports whether instruction data targets the withdraw handler.
func IsWithdrawInstruction(data []byte) bool {
	return len(data) >= 8 && bytesEqual(data[:8], withdrawInstructionDiscriminator)
}

func bytesEqual(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
