package aes

import (
	"context"
	"crypto/rand"
	handlederror "i2pdgo/core/handled-error"
)

func GenerateKey(ctx context.Context) ([]byte, error) {
	b := make([]byte, 32) //generate a random 32 byte key for AES-256
	if _, err := rand.Read(b); err != nil {
		return nil, handlederror.HandleInternalError(ctx, err, "can not generate aes key")
	}

	return b, nil
}
