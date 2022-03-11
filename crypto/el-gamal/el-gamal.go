package elgamal

import (
	"context"
	"crypto/rand"

	"golang.org/x/crypto/openpgp/elgamal"
)

func GenerateKeys(ctx context.Context) (*elgamal.PrivateKey, *elgamal.PublicKey, error) {
	privKey := &elgamal.PrivateKey{}
	if err := ElgamalGenerate(ctx, privKey, rand.Reader); err != nil {
		return nil, nil, err
	}
	pubKey := createElgamalPublicKey(privKey.Y.Bytes())

	return privKey, pubKey, nil
}

func Encrypt(ctx context.Context, pubKey *elgamal.PublicKey, data []byte) ([]byte, error) {
	var err error
	for i := 0; i < 3; i++ {
		enc, err := encrypt(ctx, pubKey, data)
		if err == nil {
			return enc, nil
		}
	}

	return nil, err
}

func encrypt(ctx context.Context, pubKey *elgamal.PublicKey, data []byte) ([]byte, error) {
	encrypter, err := createElgamalEncryption(ctx, pubKey, rand.Reader)
	if err != nil {
		panic(err.Error())
	}
	encrypted, err := encrypter.Encrypt(ctx, data)
	if err != nil {
		return nil, err
	}

	return encrypted, nil
}

func Decrypt(ctx context.Context, privKey *elgamal.PrivateKey, d []byte) ([]byte, error) {
	decrypter := &elgDecrypter{
		k: privKey,
	}
	decrypted, err := decrypter.Decrypt(ctx, d)
	if err != nil {
		return nil, err
	}
	return decrypted, nil
}
