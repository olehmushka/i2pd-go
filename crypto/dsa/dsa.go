package dsa

import (
	"context"
	handlederror "i2pdgo/core/handled-error"
)

func GenerateKeys(ctx context.Context) (DSAPrivateKey, DSAPublicKey, error) {
	var sk DSAPrivateKey
	sk, err := sk.Generate()
	if err != nil {
		return DSAPrivateKey{}, DSAPublicKey{}, err
	}

	zeros := 0
	for _, b := range sk {
		if b == 0 {
			zeros++
		}
	}
	if zeros == len(sk) {
		return DSAPrivateKey{}, DSAPublicKey{}, handlederror.HandleInternalError(ctx, nil, "dsa key generation yielded all zeros")
	}

	pk, err := sk.Public(ctx)
	if err != nil {
		return DSAPrivateKey{}, DSAPublicKey{}, err
	}

	return sk, pk, nil
}

func Sign(ctx context.Context, sk DSAPrivateKey, input []byte) ([]byte, error) {
	signer, err := sk.NewSigner()
	if err != nil {
		return nil, err
	}

	return signer.Sign(ctx, input)
}

func Verify(ctx context.Context, pk DSAPublicKey, input, signed []byte) error {
	verify, err := pk.NewVerifier()
	if err != nil {
		return err
	}

	return verify.Verify(ctx, input, signed)
}
