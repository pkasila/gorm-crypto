package algorithms

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"io"
)

type AES256CBC struct {
	Algorithm
	Block cipher.Block
}

// NewAES256CBC creates instance of AES256CBC with passed key
func NewAES256CBC(key []byte) (*AES256CBC, error) {
	block, err := aes.NewCipher(key)

	if err != nil {
		return nil, err
	}

	return &AES256CBC {
		Block: block,
	}, nil
}

// Encrypt encrypts data with key
func (algo *AES256CBC) Encrypt(msg []byte) ([]byte, error) {
	if f := len(msg)%aes.BlockSize; f != 0 {
		padding := make([]byte, aes.BlockSize-f)
		msg = append(msg, padding...)
	}

	ciphertext := make([]byte, aes.BlockSize+len(msg))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	mode := cipher.NewCBCEncrypter(algo.Block, iv)
	mode.CryptBlocks(ciphertext[aes.BlockSize:], msg)
	return ciphertext, nil
}

// Decrypt decrypts data with key
func (algo *AES256CBC) Decrypt(ciphertext []byte) ([]byte, error) {
	if len(ciphertext) < aes.BlockSize {
		return nil, errors.New("CipherTextTooShort")
	}

	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	if len(ciphertext)%aes.BlockSize != 0 {
		return nil, errors.New("CipherTextIsNotMultipleOfBlockSize")
	}

	mode := cipher.NewCBCDecrypter(algo.Block, iv)
	mode.CryptBlocks(ciphertext, ciphertext)
	withoutPadding := bytes.ReplaceAll(ciphertext, make([]byte, 1), []byte{})
	return withoutPadding, nil
}