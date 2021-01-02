package algorithms

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"fmt"
	"io"
)

type AES256GCM struct {
	Algorithm
	AEAD cipher.AEAD
}

// Deprecated: NewAES creates instance of AES256GCM with passed key (should be replaced with NewAES256GCM)
func NewAES(key []byte) (*AES256GCM, error) {
	return NewAES256GCM(key)
}

// NewAES256GCM creates instance of AES256GCM with passed key
func NewAES256GCM(key []byte) (*AES256GCM, error) {
	aesCipher, err := aes.NewCipher(key)

	if err != nil {
		return nil, err
	}

	aesGCM, err := cipher.NewGCM(aesCipher)

	if err != nil {
		return nil, err
	}

	return &AES256GCM {
		AEAD: aesGCM,
	}, nil
}

// Encrypt encrypts data with key
func (algo *AES256GCM) Encrypt(msg []byte) ([]byte, error) {
	nonce := make([]byte, algo.AEAD.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	return algo.AEAD.Seal(nonce, nonce, msg, nil), nil
}

// Decrypt decrypts data with key
func (algo *AES256GCM) Decrypt(ciphertext []byte) ([]byte, error) {
	nonceSize := algo.AEAD.NonceSize()
	if len(ciphertext) < nonceSize {
		fmt.Println(errors.New("CipherTextIsNotValid"))
	}

	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	return algo.AEAD.Open(nil, nonce, ciphertext, nil)
}