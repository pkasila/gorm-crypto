package encryption

import "testing"

func TestKeysBytes(t *testing.T) {
	privateKey, publicKey, err := GenerateKeyPair(4096)
	if err != nil {
		t.Fatalf("Failed to generate key pair: %s\n", err.Error())
	}

	privBytes := PrivateKeyToBytes(privateKey)
	pubBytes, err := PublicKeyToBytes(publicKey)
	if err != nil {
		t.Fatalf("Failed to encode public key to bytes: %s\n", err.Error())
	}

	_, err = BytesToPrivateKey(privBytes)
	if err != nil {
		t.Fatalf("Failed to decode bytes to private key: %s\n", err.Error())
	}
	_, err = BytesToPublicKey(pubBytes)
	if err != nil {
		t.Fatalf("Failed to decode bytes to public key: %s\n", err.Error())
	}
}