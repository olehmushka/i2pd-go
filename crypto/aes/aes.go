package aes

import (
	"context"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	handlederror "i2pdgo/core/handled-error"
	"io"
)

func Encrypt(ctx context.Context, key, msg []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, handlederror.HandleInternalError(ctx, err, "can not create new aes cipher")
	}

	encrypted := make([]byte, aes.BlockSize+len(msg))
	iv := encrypted[:aes.BlockSize]
	if _, err = io.ReadFull(rand.Reader, iv); err != nil {
		return nil, handlederror.HandleInternalError(ctx, err, "can not encrypt with aes")
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(encrypted[aes.BlockSize:], msg)

	return encrypted, nil
}

func Decrypt(ctx context.Context, key, encrypted []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, handlederror.HandleInternalError(ctx, err, "can not create new aes cipher")
	}

	if len(encrypted) < aes.BlockSize {
		return nil, handlederror.HandleInternalError(ctx, err, "invalid aes ciphertext block size")
	}

	iv := encrypted[:aes.BlockSize]
	encrypted = encrypted[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(encrypted, encrypted)

	return encrypted, nil
}
