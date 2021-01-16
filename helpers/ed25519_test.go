package helpers

import "testing"

func TestEd25519KeysConversion(t *testing.T) {
	// Generate key pair
	privateKey, publicKey, err := Ed25519GenerateKeyPair()
	if err != nil {
		t.Fatalf("Failed to generate key pair: %s\n", err.Error())
	}

	// Convert to bytes
	privBytes := Ed25519PrivateKeyToBytes(&privateKey)
	pubBytes := Ed25519PublicKeyToBytes(&publicKey)

	// Convert back to rsa.PrivateKey and rsa.PublicKey
	_, err = Ed25519BytesToPrivateKey(privBytes)
	if err != nil {
		t.Fatalf("Failed to convert bytes to private key: %s\n", err.Error())
	}
	_ = Ed25519BytesToPublicKey(pubBytes)
}