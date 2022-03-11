package aes

import (
	"bytes"
	"context"
	"testing"
)

func TestAES(t *testing.T) {
	ctx := context.Background()

	key, err := GenerateKey(ctx)
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	msg := []byte("hello world")
	enc, err := Encrypt(ctx, key, msg)
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	dec, err := Decrypt(ctx, key, enc)
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	if !bytes.Equal(msg, dec) {
		t.Fail()
	}
}
