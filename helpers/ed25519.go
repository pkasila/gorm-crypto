package helpers

import (
	"crypto/ed25519"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
)

// Ed25519GenerateKeyPair generates a new key pair
func Ed25519GenerateKeyPair() (ed25519.PrivateKey, ed25519.PublicKey, error) {
	privateKey, publicKey, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		return nil, nil, err
	}
	return publicKey, privateKey, nil
}

// Ed25519PrivateKeyToBytes private key to bytes
func Ed25519PrivateKeyToBytes(priv *ed25519.PrivateKey) []byte {
	x509Encoded, _ := x509.MarshalPKCS8PrivateKey(*priv)
	pemEncoded := pem.EncodeToMemory(&pem.Block{Type: "PRIVATE KEY", Bytes: x509Encoded})
	return pemEncoded
}

// Ed25519PublicKeyToBytes private key to bytes
func Ed25519PublicKeyToBytes(pub *ed25519.PublicKey) []byte {
	x509Encoded, _ := x509.MarshalPKIXPublicKey(*pub)
	pemEncoded := pem.EncodeToMemory(&pem.Block{Type: "PUBLIC KEY", Bytes: x509Encoded})
	return pemEncoded
}

// Ed25519BytesToPrivateKey bytes to private key
func Ed25519BytesToPrivateKey(priv []byte) (*ed25519.PrivateKey, error) {
	block, _ := pem.Decode(priv)
	x509Encoded := block.Bytes
	genericPrivateKey, _ := x509.ParsePKCS8PrivateKey(x509Encoded)
	privateKey := genericPrivateKey.(ed25519.PrivateKey)
	return &privateKey, nil
}

// Ed25519BytesToPublicKey bytes to private key
func Ed25519BytesToPublicKey(pub []byte) *ed25519.PublicKey {
	blockPub, _ := pem.Decode(pub)
	x509EncodedPub := blockPub.Bytes
	genericPublicKey, _ := x509.ParsePKIXPublicKey(x509EncodedPub)
	publicKey := genericPublicKey.(ed25519.PublicKey)
	return &publicKey
}