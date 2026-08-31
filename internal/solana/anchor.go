package solana

import "crypto/sha256"

// anchorDiscriminator returns the first 8 bytes of SHA-256("namespace:name"),
// matching Anchor's instruction/event/account discriminator scheme.
func anchorDiscriminator(namespace, name string) []byte {
	sum := sha256.Sum256([]byte(namespace + ":" + name))
	return sum[:8]
}

var (
	depositInstructionDiscriminator = anchorDiscriminator("global", "deposit")
	depositEventDiscriminator       = anchorDiscriminator("event", "DepositEvent")
)
