package algorithms

import (
	"testing"
)

func TestAES256GCMEncryptDecryptCycle(t *testing.T) {
	aes, err := NewAES256GCM([]byte("passphrasewhichneedstobe32bytes!"))
	if err != nil {
		t.Fatalf("Failed to setup AES: %s\n", err.Error())
	}

	message := "A string... just for testing purposes :-)"

	encrypted, err := aes.Encrypt([]byte(message))
	if err != nil {
		t.Fatalf("Failed encrypt data: %s\n", err.Error())
	}

	decrypted, err := aes.Decrypt(encrypted)
	if err != nil {
		t.Fatalf("Failed decrypt data: %s\n", err.Error())
	}

	decryptedStr := string(decrypted)

	if decryptedStr != message {
		t.Fatalf("Source and decrypted messages are not equal: %s != %s\n", decryptedStr, message)
	}
}

func TestAES256CBCEncryptDecryptCycle(t *testing.T) {
	aes, err := NewAES256CBC([]byte("passphrasewhichneedstobe32bytes!"), []byte("0000000000000000"))
	if err != nil {
		t.Fatalf("Failed to setup AES: %s\n", err.Error())
	}

	message := []byte("A string... just for testing purposes :-)")

	encrypted, err := aes.Encrypt(message)
	if err != nil {
		t.Fatalf("Failed encrypt data: %s\n", err.Error())
	}

	decrypted, err := aes.Decrypt(encrypted)
	if err != nil {
		t.Fatalf("Failed decrypt data: %s\n", err.Error())
	}

	decryptedStr := string(decrypted)
	messageStr := string(message)

	if decryptedStr != messageStr {
		t.Fatalf("Source and decrypted messages are not equal: %s != %s\n", decryptedStr, message)
	}
}
