package elgamal

import (
	"context"
	"crypto/sha256"
	handlederror "i2pdgo/core/handled-error"
	"math/big"
)

// encrypts data
type Encrypter interface {
	// encrypt a block of data
	// return encrypted block or nil and error if an error happened
	Encrypt(ctx context.Context, data []byte) ([]byte, error)
}

type PublicEncryptionKey interface {

	// create a new encrypter to encrypt data to this public key
	NewEncrypter(ctx context.Context) (Encrypter, error)

	// length of this public key in bytes
	Len() int
}

type ElgamalEncryption struct {
	p, a, b1 *big.Int
}

func (elg *ElgamalEncryption) Encrypt(ctx context.Context, data []byte) ([]byte, error) {
	return elg.EncryptPadding(ctx, data, true)
}

func (elg *ElgamalEncryption) EncryptPadding(ctx context.Context, data []byte, zeroPadding bool) ([]byte, error) {
	if len(data) > 222 {
		return nil, handlederror.HandleInternalError(ctx, nil, "failed to encrypt data, too big for elgamal")
	}
	mbytes := make([]byte, 255)
	mbytes[0] = 0xFF
	copy(mbytes[33:], data)
	// do sha256 of payload
	d := sha256.Sum256(mbytes[33 : len(data)+33])
	copy(mbytes[1:], d[:])
	m := new(big.Int).SetBytes(mbytes)
	// do encryption
	b := new(big.Int).Mod(new(big.Int).Mul(elg.b1, m), elg.p).Bytes()

	var encrypted []byte
	if zeroPadding {
		encrypted = make([]byte, 514)
		copy(encrypted[1:], elg.a.Bytes())
		copy(encrypted[258:], b)
	} else {
		encrypted = make([]byte, 512)
		copy(encrypted, elg.a.Bytes())
		copy(encrypted[256:], b)
	}
	return encrypted, nil
}
