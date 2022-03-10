package sha256

import (
	"crypto/sha256"
)

func Hash(input []byte) [32]byte {
	return sha256.Sum256(input)
}
