package elgamal

import (
	"context"
	"crypto/rand"
	handlederror "i2pdgo/core/handled-error"
	"io"
	"math/big"

	"golang.org/x/crypto/openpgp/elgamal"
)

// generate an elgamal key pair
func ElgamalGenerate(ctx context.Context, priv *elgamal.PrivateKey, rand io.Reader) (err error) {
	priv.P = elgp
	priv.G = elgg
	xBytes := make([]byte, priv.P.BitLen()/8)
	_, err = io.ReadFull(rand, xBytes)
	if err != nil {
		return handlederror.HandleInternalError(ctx, err, "can not read x bytes for elgamal generation")
	}
	// set private key
	priv.X = new(big.Int).SetBytes(xBytes)
	// compute public key
	priv.Y = new(big.Int).Exp(priv.G, priv.X, priv.P)

	return nil
}

// create an elgamal public key from byte slice
func createElgamalPublicKey(data []byte) *elgamal.PublicKey {
	if len(data) != 256 {
		return nil
	}

	return &elgamal.PublicKey{
		G: elgg,
		P: elgp,
		Y: new(big.Int).SetBytes(data),
	}
}

// create an elgamal private key from byte slice
func createElgamalPrivateKey(data []byte) *elgamal.PrivateKey {
	if len(data) != 256 {
		return nil
	}

	x := new(big.Int).SetBytes(data)
	y := new(big.Int).Exp(elgg, x, elgp)

	return &elgamal.PrivateKey{
		PublicKey: elgamal.PublicKey{
			Y: y,
			G: elgg,
			P: elgp,
		},
		X: x,
	}
}

// create a new elgamal encryption session
func createElgamalEncryption(ctx context.Context, pub *elgamal.PublicKey, rand io.Reader) (*ElgamalEncryption, error) {
	kbytes := make([]byte, 256)
	k := new(big.Int)
	var err error
	for err == nil {
		_, err = io.ReadFull(rand, kbytes)
		k = new(big.Int).SetBytes(kbytes)
		k = k.Mod(k, pub.P)
		if k.Sign() != 0 {
			break
		}
	}
	if err != nil {
		return nil, handlederror.HandleInternalError(ctx, err, "can not create elgamal bytes for enc")
	}

	return &ElgamalEncryption{
		p:  pub.P,
		a:  new(big.Int).Exp(pub.G, k, pub.P),
		b1: new(big.Int).Exp(pub.Y, k, pub.P),
	}, nil
}

type ElgPublicKey [256]byte
type ElgPrivateKey [256]byte

func (elg ElgPublicKey) Len() int {
	return len(elg)
}

func (elg ElgPublicKey) NewEncrypter(ctx context.Context) (Encrypter, error) {
	k := createElgamalPublicKey(elg[:])
	return createElgamalEncryption(ctx, k, rand.Reader)
}

func (elg ElgPrivateKey) Len() int {
	return len(elg)
}

func (elg ElgPrivateKey) NewDecrypter() (Decrypter, error) {
	return &elgDecrypter{
		k: createElgamalPrivateKey(elg[:]),
	}, nil
}
