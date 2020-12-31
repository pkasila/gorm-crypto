package algorithms

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha512"
)

type RSA struct {
	Algorithm
	PrivateKey *rsa.PrivateKey
	PublicKey  *rsa.PublicKey
}

func NewRSA(privateKey *rsa.PrivateKey, publicKey *rsa.PublicKey) *RSA {
	return &RSA {
		PrivateKey: privateKey,
		PublicKey: publicKey,
	}
}

// Encrypt encrypts data with public key
func (algo *RSA) Encrypt(msg []byte) ([]byte, error) {
	hash := sha512.New()
	ciphertext, err := rsa.EncryptOAEP(hash, rand.Reader, algo.PublicKey, msg, nil)
	if err != nil {
		return nil, err
	}
	return ciphertext, nil
}

// Decrypt decrypts data with private key
func (algo *RSA) Decrypt(ciphertext []byte) ([]byte, error) {
	hash := sha512.New()
	plaintext, err := rsa.DecryptOAEP(hash, rand.Reader, algo.PrivateKey, ciphertext, nil)
	if err != nil {
		return nil, err
	}
	return plaintext, nil
}

