package dsa

import (
	"context"
	"crypto/dsa"
	"crypto/rand"
	handlederror "i2pdgo/core/handled-error"
	"io"
	"math/big"
)

// generate a dsa keypair
func generateDSA(priv *dsa.PrivateKey, rand io.Reader) error {
	// put our paramters in
	priv.P = param.P
	priv.Q = param.Q
	priv.G = param.G
	// generate the keypair
	return dsa.GenerateKey(priv, rand)
}

// create i2p dsa public key given its public component
func createDSAPublicKey(Y *big.Int) *dsa.PublicKey {
	return &dsa.PublicKey{
		Parameters: param,
		Y:          Y,
	}
}

// createa i2p dsa private key given its public component
func createDSAPrivkey(X *big.Int) *dsa.PrivateKey {
	if X.Cmp(dsap) != -1 {
		return nil
	}

	Y := new(big.Int)
	Y.Exp(dsag, X, dsap)

	return &dsa.PrivateKey{
		PublicKey: dsa.PublicKey{
			Parameters: param,
			Y:          Y,
		},
		X: X,
	}
}

type DSAPrivateKey [20]byte

// create a new dsa signer
func (k DSAPrivateKey) NewSigner() (Signer, error) {
	return &DSASigner{
		k: createDSAPrivkey(new(big.Int).SetBytes(k[:])),
	}, nil
}

func (k DSAPrivateKey) Public(ctx context.Context) (DSAPublicKey, error) {
	p := createDSAPrivkey(new(big.Int).SetBytes(k[:]))
	if p == nil {
		return DSAPublicKey{}, handlederror.HandleInternalError(ctx, nil, "invalid dsa private key format")
	}

	var pk DSAPublicKey
	copy(pk[:], p.Y.Bytes())

	return pk, nil
}

func (k DSAPrivateKey) Generate() (s DSAPrivateKey, err error) {
	dk := new(dsa.PrivateKey)
	err = generateDSA(dk, rand.Reader)
	if err == nil {
		copy(k[:], dk.X.Bytes())
		s = k
	}
	return
}
