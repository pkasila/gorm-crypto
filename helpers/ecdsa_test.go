package helpers

import "testing"

func TestECDSAKeysConversion(t *testing.T) {
	// Generate key pair
	privateKey, publicKey, err := ECDSAGenerateKeyPair()
	if err != nil {
		t.Fatalf("Failed to generate key pair: %s\n", err.Error())
	}

	// Convert to bytes
	privBytes := ECDSAPrivateKeyToBytes(privateKey)
	pubBytes := ECDSAPublicKeyToBytes(publicKey)

	// Convert back to rsa.PrivateKey and rsa.PublicKey
	_, err = ECDSABytesToPrivateKey(privBytes)
	if err != nil {
		t.Fatalf("Failed to convert bytes to private key: %s\n", err.Error())
	}
	_ = ECDSABytesToPublicKey(pubBytes)
}