package helpers

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
)

// ECDSAGenerateKeyPair generates a new key pair
func ECDSAGenerateKeyPair() (*ecdsa.PrivateKey, *ecdsa.PublicKey, error) {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return nil, nil, err
	}
	return privateKey, &privateKey.PublicKey, nil
}

// ECDSABytesToPrivateKey bytes to private key
func ECDSABytesToPrivateKey(priv []byte) (*ecdsa.PrivateKey, error) {
	block, _ := pem.Decode(priv)
	x509Encoded := block.Bytes
	privateKey, _ := x509.ParseECPrivateKey(x509Encoded)
	return privateKey, nil
}

// ECDSABytesToPublicKey bytes to private key
func ECDSABytesToPublicKey(pub []byte) *ecdsa.PublicKey {
	blockPub, _ := pem.Decode(pub)
	x509EncodedPub := blockPub.Bytes
	genericPublicKey, _ := x509.ParsePKIXPublicKey(x509EncodedPub)
	publicKey := genericPublicKey.(*ecdsa.PublicKey)
	return publicKey
}

// ECDSAPrivateKeyToBytes private key to bytes
func ECDSAPrivateKeyToBytes(priv *ecdsa.PrivateKey) []byte {
	x509Encoded, _ := x509.MarshalECPrivateKey(priv)
	pemEncoded := pem.EncodeToMemory(&pem.Block{Type: "PRIVATE KEY", Bytes: x509Encoded})
	return pemEncoded
}

// ECDSAPublicKeyToBytes private key to bytes
func ECDSAPublicKeyToBytes(pub *ecdsa.PublicKey) []byte {
	x509Encoded, _ := x509.MarshalPKIXPublicKey(pub)
	pemEncoded := pem.EncodeToMemory(&pem.Block{Type: "PUBLIC KEY", Bytes: x509Encoded})
	return pemEncoded
}