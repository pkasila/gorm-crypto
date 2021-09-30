package signing

import (
	"github.com/pkasila/gorm-crypto/helpers"
	"testing"
)

func TestEd25519EncryptDecryptCycle(t *testing.T) {
	// Generate key pair
	privateKey, publicKey, err := helpers.Ed25519GenerateKeyPair()
	if err != nil {
		t.Fatalf("Failed to generate key pair: %s\n", err.Error())
	}

	algo := NewEd25519(&privateKey, &publicKey)

	message := "A string... just for testing purposes :-)"

	sig, err := algo.Sign([]byte(message))
	if err != nil {
		t.Fatalf("Failed sign data: %s\n", err.Error())
	}

	valid, err := algo.Verify([]byte(message), sig)
	if err != nil || !valid {
		t.Fatalf("Failed verify data: %s\n", err.Error())
	}
}
