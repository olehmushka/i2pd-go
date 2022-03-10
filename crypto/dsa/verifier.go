package dsa

import (
	"context"
	"crypto/dsa"
	"crypto/sha1"
	handlederror "i2pdgo/core/handled-error"
	"math/big"
)

// type for verifying signatures
type Verifier interface {
	// verify hashed data with this signing key
	// return nil on valid signature otherwise error
	VerifyHash(ctx context.Context, h, sig []byte) error
	// verify an unhashed piece of data by hashing it and calling VerifyHash
	Verify(ctx context.Context, data, sig []byte) error
}

type DSAVerifier struct {
	k *dsa.PublicKey
}

type DSAPublicKey [128]byte

// create a new dsa verifier
func (k DSAPublicKey) NewVerifier() (Verifier, error) {
	return &DSAVerifier{
		k: createDSAPublicKey(new(big.Int).SetBytes(k[:])),
	}, nil
}

// verify data with a dsa public key
func (v *DSAVerifier) Verify(ctx context.Context, data, sig []byte) error {
	h := sha1.Sum(data)

	return v.VerifyHash(ctx, h[:], sig)
}

// verify hash of data with a dsa public key
func (v *DSAVerifier) VerifyHash(ctx context.Context, h, sig []byte) error {
	if len(sig) != 40 {
		return handlederror.HandleInternalError(ctx, nil, "bad dsa public key signature size")
	}

	r := new(big.Int).SetBytes(sig[:20])
	s := new(big.Int).SetBytes(sig[20:])
	if dsa.Verify(v.k, h, r, s) {
		return nil
	}

	return handlederror.HandleInternalError(ctx, nil, "invalid dsa public key signature")
}

func (k DSAPublicKey) Len() int {
	return len(k)
}
