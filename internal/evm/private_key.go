package evm

import (
	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/crypto"
)

// PrivateKeyConfigured reports whether key holds a usable ECDSA private key.
func PrivateKeyConfigured(key *ecdsa.PrivateKey) bool {
	if key == nil || key.Curve == nil {
		return false
	}
	return len(crypto.FromECDSA(key)) > 0
}
