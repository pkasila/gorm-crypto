package gormcrypto

import (
	"github.com/pkosilo/gorm-crypto/algorithms"
	"github.com/pkosilo/gorm-crypto/helpers"
	"github.com/pkosilo/gorm-crypto/serialization"
	"testing"
)

func TestInit(t *testing.T) {
	privateKey, publicKey, err := helpers.RSAGenerateKeyPair(4096)

	if err != nil {
		t.Fatalf("Failed to generate key pair: %s\n", err.Error())
	}

	Init(algorithms.NewRSA(privateKey, publicKey), serialization.NewJSON())
}
