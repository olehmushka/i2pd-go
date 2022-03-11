package elgamal

import (
	"bytes"
	"context"
	"crypto/rand"
	"io"
	"testing"
)

func TestElGamal(t *testing.T) {
	ctx := context.Background()

	tCases := []struct {
		name  string
		times int
	}{
		{
			name:  "should be ok for 100 times",
			times: 100,
		},
	}

	for _, tc := range tCases {
		t.Run(tc.name, func(t *testing.T) {
			for i := 0; i < tc.times; i++ {
				privKey, pubKey, err := GenerateKeys(ctx)
				if err != nil {
					t.Log(err)
					t.Fail()
				}

				input := make([]byte, 222)
				_, _ = io.ReadFull(rand.Reader, input)
				encrypted, err := Encrypt(ctx, pubKey, input)
				if err != nil {
					t.Log(err)
					t.Fail()
				}

				output, err := Decrypt(ctx, privKey, encrypted)
				if err != nil {
					t.Log(err)
					t.Fail()
				}

				if !bytes.Equal(input, output) {
					t.Fail()
				}
			}
		})
	}
}
