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
		t.Log(err)
		t.Fail()
	}

	data := make([]byte, 512)
	_, err = io.ReadFull(rand.Reader, data)
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	signed, err := Sign(ctx, sk, data)
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	if err = Verify(ctx, pk, data, signed); err != nil {
		t.Log(err)
		t.Fail()
	}
}
