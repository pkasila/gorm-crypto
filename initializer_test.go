package gormcrypto

import (
	"github.com/pkosilo/gorm-crypto/algorithms"
	"github.com/pkosilo/gorm-crypto/helpers"
	"reflect"
	"testing"
)

func TestInitFromKeyPair(t *testing.T) {
	privateKey, publicKey, err := helpers.RSAGenerateKeyPair(4096)

	if err != nil {
		t.Fatalf("Failed to generate key pair: %s\n", err.Error())
	}

	InitFromKeyPair(privateKey, publicKey)

	if reflect.TypeOf(Algorithm) != reflect.TypeOf(&algorithms.RSA{}) {
		t.Fatalf("Algorithm is not of type algorithms.RSA: %s != %s", reflect.TypeOf(Algorithm), reflect.TypeOf(&algorithms.RSA{}))
	}
}
