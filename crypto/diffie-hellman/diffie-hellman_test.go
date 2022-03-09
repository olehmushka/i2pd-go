package diffiehellman

import (
	"context"
	"testing"

	"github.com/aead/ecdh"
)

func TestDiffieHellman(t *testing.T) {
	ctx := context.Background()

	tCases := []struct {
		name    string
		generic ecdh.KeyExchange
		times   int
	}{
		{
			name:    "should be ok for p256",
			generic: GenerateP256Generic(),
			times:   100,
		},
		{
			name:    "should be ok for x25519",
			generic: GenerateX25519Generic(),
			times:   100,
		},
	}

	for _, tc := range tCases {
		t.Run(tc.name, func(t *testing.T) {
			for i := 0; i < tc.times; i++ {
				keys, err := GenerateKeys(ctx, tc.generic)
				if err != nil {
					t.Fatalf("unexpected error: %+v", err)
				}

				if !IsSecretEqual(keys.FirstParticipant.Secret, keys.SecondParticipant.Secret) {
					t.Fail()
				}
			}
		})
	}

	/*
		p256 := Generic(elliptic.P256())

		privateAlice, publicAlice, err := p256.GenerateKey(rand.Reader)
		if err != nil {
			fmt.Printf("Failed to generate Alice's private/public key pair: %s\n", err)
		}
		privateBob, publicBob, err := p256.GenerateKey(rand.Reader)
		if err != nil {
			fmt.Printf("Failed to generate Bob's private/public key pair: %s\n", err)
		}

		if err := p256.Check(publicBob); err != nil {
			fmt.Printf("Bob's public key is not on the curve: %s\n", err)
		}
		secretAlice := p256.ComputeSecret(privateAlice, publicBob)

		if err := p256.Check(publicAlice); err != nil {
			fmt.Printf("Alice's public key is not on the curve: %s\n", err)
		}
		secretBob := p256.ComputeSecret(privateBob, publicAlice)

		if !bytes.Equal(secretAlice, secretBob) {
			fmt.Printf("key exchange failed - secret X coordinates not equal\n")
		}
	*/
}
