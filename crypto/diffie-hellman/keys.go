package diffiehellman

import (
	"bytes"
	"context"
	"crypto"
	"crypto/elliptic"
	"crypto/rand"
	handlederror "i2pdgo/core/handled-error"

	"github.com/aead/ecdh"
)

func GenerateP256Generic() ecdh.KeyExchange {
	return ecdh.Generic(elliptic.P256())
}

func GenerateX25519Generic() ecdh.KeyExchange {
	return ecdh.X25519()
}

type Key struct {
	PrivateKey crypto.PrivateKey
	PublicKey  crypto.PublicKey

	Secret []byte
}

type Keys struct {
	FirstParticipant  *Key
	SecondParticipant *Key
}

func GenerateKeys(ctx context.Context, generic ecdh.KeyExchange) (*Keys, error) {
	firstPrivateKey, firstPublicKey, err := generic.GenerateKey(rand.Reader)
	if err != nil {
		return nil, handlederror.HandleInternalError(ctx, err, "failed to generate first private/public key pair")
	}
	if err = generic.Check(firstPublicKey); err != nil {
		return nil, handlederror.HandleInternalError(ctx, err, "first public key is not on the curve")
	}

	secondPrivateKey, secondPublicKey, err := generic.GenerateKey(rand.Reader)
	if err != nil {
		return nil, handlederror.HandleInternalError(ctx, err, "failed to generate second private/public key pair")
	}
	if err = generic.Check(firstPublicKey); err != nil {
		return nil, handlederror.HandleInternalError(ctx, err, "second public key is not on the curve")
	}

	return &Keys{
		FirstParticipant: &Key{
			PrivateKey: firstPrivateKey,
			PublicKey:  firstPublicKey,
			Secret:     generic.ComputeSecret(firstPrivateKey, secondPublicKey),
		},
		SecondParticipant: &Key{
			PrivateKey: secondPrivateKey,
			PublicKey:  secondPublicKey,
			Secret:     generic.ComputeSecret(secondPrivateKey, firstPublicKey),
		},
	}, nil
}

func IsSecretEqual(firstSecret, secondSecret []byte) bool {
	return bytes.Equal(firstSecret, secondSecret)
}
