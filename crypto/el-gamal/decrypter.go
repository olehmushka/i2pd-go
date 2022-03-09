package elgamal

import (
	"context"
	"crypto/sha256"
	"crypto/subtle"
	handlederror "i2pdgo/core/handled-error"
	"math/big"

	"golang.org/x/crypto/openpgp/elgamal"
)

// decrypts data
type Decrypter interface {
	// decrypt a block of data
	// return decrypted block or nil and error if error happens
	Decrypt(ctx context.Context, data []byte) ([]byte, error)
}

type PrivateEncryptionKey interface {

	// create a new decryption object for this private key to decrypt data encrypted to our public key
	// returns decrypter or nil and error if the private key is in a bad format
	NewDecrypter() (Decrypter, error)
}

type elgDecrypter struct {
	k *elgamal.PrivateKey
}

func (elg *elgDecrypter) Decrypt(ctx context.Context, data []byte) ([]byte, error) {
	return elgamalDecrypt(ctx, elg.k, data, true) // TODO(psi): should this be true or false?
}

// decrypt an elgamal encrypted message, i2p style
func elgamalDecrypt(ctx context.Context, priv *elgamal.PrivateKey, data []byte, zeroPadding bool) ([]byte, error) {
	a := new(big.Int)
	b := new(big.Int)
	idx := 0
	if zeroPadding {
		idx++
	}
	a.SetBytes(data[idx : idx+256])
	if zeroPadding {
		idx++
	}
	b.SetBytes(data[idx+256:])

	// decrypt
	m := new(big.Int).Mod(new(big.Int).Mul(b, new(big.Int).Exp(a, new(big.Int).Sub(new(big.Int).Sub(priv.P, priv.X), one), priv.P)), priv.P).Bytes()

	// check digest
	d := sha256.Sum256(m[33:255])
	if subtle.ConstantTimeCompare(d[:], m[1:33]) == 1 {
		decrypted := make([]byte, 222)
		subtle.ConstantTimeCopy(1, decrypted, m[33:255])
		return decrypted, nil
	}

	return nil, handlederror.HandleInternalError(ctx, nil, "failed to decrypt elgamal encrypted data")
}
