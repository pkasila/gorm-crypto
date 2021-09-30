package algorithms

import (
	"github.com/pkasila/gorm-crypto/helpers"
	"testing"
)

func TestRSAEncryptDecryptCycle(t *testing.T) {
	// Generate key pair
	privateKey, publicKey, err := helpers.RSAGenerateKeyPair(4096)
	if err != nil {
		t.Fatalf("Failed to generate key pair: %s\n", err.Error())
	}

	rsa := NewRSA(privateKey, publicKey)

	message := "A string... just for testing purposes :-)"

	encrypted, err := rsa.Encrypt([]byte(message))
	if err != nil {
		t.Fatalf("Failed encrypt data: %s\n", err.Error())
	}

	decrypted, err := rsa.Decrypt(encrypted)
	if err != nil {
		t.Fatalf("Failed decrypt data: %s\n", err.Error())
	}

	decryptedStr := string(decrypted)

	if decryptedStr != message {
		t.Fatalf("Source and decrypted messages are not equal: %s != %s\n", decryptedStr, message)
	}
}
