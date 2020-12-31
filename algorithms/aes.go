package algorithms

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"fmt"
	"io"
)

type AES struct {
	Algorithm
	AEAD cipher.AEAD
}

func NewAES(key []byte) (*AES, error) {
	aesCipher, err := aes.NewCipher(key)

	if err != nil {
		return nil, err
	}

	aesGCM, err := cipher.NewGCM(aesCipher)

	if err != nil {
		return nil, err
	}

	return &AES {
		AEAD: aesGCM,
	}, nil
}

// Encrypt encrypts data with key
func (algo *AES) Encrypt(msg []byte) ([]byte, error) {
	nonce := make([]byte, algo.AEAD.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	return algo.AEAD.Seal(nonce, nonce, msg, nil), nil
}

// Decrypt decrypts data with key
func (algo *AES) Decrypt(ciphertext []byte) ([]byte, error) {
	nonceSize := algo.AEAD.NonceSize()
	if len(ciphertext) < nonceSize {
		fmt.Println(errors.New("CipherTextIsNotValid"))
	}

	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	return algo.AEAD.Open(nil, nonce, ciphertext, nil)
}