package dsa

import (
	"context"
	"crypto/rand"
	"io"
	"testing"
)

func TestDSA(t *testing.T) {
	ctx := context.Background()
	sk, pk, err := GenerateKeys(ctx)
	if err != nil {
		t.Fail()
	}

	data := make([]byte, 512)
	io.ReadFull(rand.Reader, data)
	signed, err := Sign(ctx, sk, data)
	if err != nil {
		t.Fail()
	}

	if err := Verify(ctx, pk, data, signed); err != nil {
		t.Fail()
	}
}
