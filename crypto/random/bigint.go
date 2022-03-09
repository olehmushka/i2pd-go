package random

import (
	"context"
	handlederror "i2pdgo/core/handled-error"

	"crypto/rand"
	"math/big"
)

func GetRandomBigInt(ctx context.Context) (*big.Int, error) {
	max := new(big.Int)
	max.Exp(big.NewInt(2), big.NewInt(130), nil).Sub(max, big.NewInt(1))

	n, err := rand.Int(rand.Reader, max)
	if err != nil {
		return nil, handlederror.HandleInternalError(ctx, err, "can not generate cryptographically bigint")
	}

	return n, nil
}
