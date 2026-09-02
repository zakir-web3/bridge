package evm

import (
	"crypto/ecdsa"
	"testing"

	"github.com/ethereum/go-ethereum/crypto"
)

func TestPrivateKeyConfigured(t *testing.T) {
	valid, err := crypto.HexToECDSA("0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef")
	if err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		name string
		key  *ecdsa.PrivateKey
		want bool
	}{
		{"nil", nil, false},
		{"empty", &ecdsa.PrivateKey{}, false},
		{"valid", valid, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PrivateKeyConfigured(tt.key); got != tt.want {
				t.Fatalf("PrivateKeyConfigured() = %v, want %v", got, tt.want)
			}
		})
	}
}
