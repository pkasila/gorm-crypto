package helpers

import "testing"

func TestRSAKeysConversion(t *testing.T) {
	// Generate key pair
	privateKey, publicKey, err := RSAGenerateKeyPair(4096)
	if err != nil {
		t.Fatalf("Failed to generate key pair: %s\n", err.Error())
	}

	// Convert to bytes
	privBytes := RSAPrivateKeyToBytes(privateKey)
	pubBytes, err := RSAPublicKeyToBytes(publicKey)
	if err != nil {
		t.Fatalf("Failed to convert public key to bytes: %s\n", err.Error())
	}

	// Convert back to rsa.PrivateKey and rsa.PublicKey
	privateKey, err = RSABytesToPrivateKey(privBytes)
	if err != nil {
		t.Fatalf("Failed to convert bytes to private key: %s\n", err.Error())
	}
	publicKey, err = RSABytesToPublicKey(pubBytes)
	if err != nil {
		t.Fatalf("Failed to convert bytes to public key: %s\n", err.Error())
	}
}