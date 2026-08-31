package solana

import (
	"bytes"
	"encoding/base64"
	"encoding/binary"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/pkg/errors"

	"github.com/zakir-web3/bridge/internal/contract"
)

// ParsedDeposit is the parsed on-chain deposit event plus tx metadata.
type ParsedDeposit struct {
	contract.SolanaDeposit
	Signature solana.Signature
}

// TxSignatureHash returns keccak256 over the 64-byte ed25519 signature bytes.
func TxSignatureHash(sig solana.Signature) [32]byte {
	return crypto.Keccak256Hash(sig[:])
}

// ParseDepositEvent decodes Anchor DepositEvent bytes (after 8-byte discriminator).
func ParseDepositEvent(data []byte) (user solana.PublicKey, destination common.Address, mint solana.PublicKey, amount uint64, err error) {
	if len(data) < 8+32+20+32+8 {
		return solana.PublicKey{}, destination, solana.PublicKey{}, 0, errors.New("deposit event data too short")
	}
	if !bytes.Equal(data[:8], depositEventDiscriminator) {
		return solana.PublicKey{}, destination, solana.PublicKey{}, 0, errors.New("invalid deposit event discriminator")
	}
	offset := 8
	copy(user[:], data[offset:offset+32])
	offset += 32
	destination = common.BytesToAddress(data[offset : offset+20])
	offset += 20
	copy(mint[:], data[offset:offset+32])
	offset += 32
	amount = binary.LittleEndian.Uint64(data[offset : offset+8])
	return user, destination, mint, amount, nil
}

// ParseDepositFromLogs scans program logs for a DepositEvent.
func ParseDepositFromLogs(logs []string) (user solana.PublicKey, destination common.Address, mint solana.PublicKey, amount uint64, err error) {
	for _, line := range logs {
		const prefix = "Program data: "
		if len(line) < len(prefix) || line[:len(prefix)] != prefix {
			continue
		}
		raw, decErr := base64.StdEncoding.DecodeString(line[len(prefix):])
		if decErr != nil {
			continue
		}
		user, destination, mint, amount, err = ParseDepositEvent(raw)
		if err == nil {
			return user, destination, mint, amount, nil
		}
	}
	return solana.PublicKey{}, destination, solana.PublicKey{}, 0, errors.New("deposit event not found in logs")
}

// ParseDepositFromTransaction parses a finalized transaction for a bridge deposit.
func ParseDepositFromTransaction(
	programID solana.PublicKey,
	result *rpc.GetTransactionResult,
	signature solana.Signature,
) (*ParsedDeposit, error) {
	if result == nil || result.Meta == nil || result.Transaction == nil {
		return nil, errors.New("empty transaction result")
	}
	if result.Slot == 0 {
		return nil, errors.New("transaction slot missing")
	}

	tx, err := result.Transaction.GetTransaction()
	if err != nil {
		return nil, errors.Wrap(err, "decode transaction")
	}

	instructionIndex, err := findDepositInstructionIndex(tx, programID)
	if err != nil {
		return nil, err
	}

	user, destination, mint, amount, err := ParseDepositFromLogs(result.Meta.LogMessages)
	if err != nil {
		return nil, err
	}
	if destination == (common.Address{}) {
		return nil, errors.New("empty destination address")
	}

	parsed := &ParsedDeposit{
		SolanaDeposit: contract.SolanaDeposit{
			User:                 pubkeyToBytes32(user),
			Destination:          destination,
			Token:                pubkeyToBytes32(mint),
			Amount:               new(big.Int).SetUint64(amount),
			BlockNumber:          result.Slot,
			TxHash:               TxSignatureHash(signature),
			Index:                instructionIndex,
			TransactionSignature: signature.String(),
		},
		Signature: signature,
	}
	return parsed, nil
}

func findDepositInstructionIndex(tx *solana.Transaction, programID solana.PublicKey) (uint32, error) {
	accountKeys := tx.Message.AccountKeys
	for i, ix := range tx.Message.Instructions {
		programKey := accountKeys[ix.ProgramIDIndex]
		if programKey.Equals(programID) && IsDepositInstruction(ix.Data) {
			return uint32(i), nil
		}
	}
	return 0, errors.New("deposit instruction not found")
}

func pubkeyToBytes32(pub solana.PublicKey) [32]byte {
	var out [32]byte
	copy(out[:], pub.Bytes())
	return out
}
