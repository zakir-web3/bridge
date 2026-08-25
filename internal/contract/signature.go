package contract

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/pkg/errors"
)

func RecoverSignerAddress(domainSeparator, message common.Hash, signature Signature) (common.Address, error) {
	rawData := fmt.Sprintf("\x19\x01%s%s", string(domainSeparator.Bytes()), string(message.Bytes()))
	sig := make([]byte, 65)
	rBytes := signature.R.Bytes()
	copy(sig[32-len(rBytes):32], rBytes)
	sBytes := signature.S.Bytes()
	copy(sig[64-len(sBytes):64], sBytes)
	sig[64] = signature.V - 27 // EIP-155 compatibility
	pub, err := crypto.SigToPub(crypto.Keccak256([]byte(rawData)), sig)
	if err != nil {
		return [20]byte{}, errors.Wrap(err, "recover signer address")
	}
	return crypto.PubkeyToAddress(*pub), nil
}

func (s *Signature) IsEmpty() bool {
	return (s.R == nil || s.R.Sign() == 0) || (s.S == nil || s.S.Sign() == 0)
}
