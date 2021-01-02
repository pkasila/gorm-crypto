package algorithms

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
)

type AES256CBC struct {
	Algorithm
	Encrypter cipher.BlockMode
	Decrypter cipher.BlockMode
}

// NewAES256CBC creates instance of AES256CBC with passed key and IV
func NewAES256CBC(key []byte, iv []byte) (*AES256CBC, error) {
	block, err := aes.NewCipher(key)

	if err != nil {
		return nil, err
	}

	encrypter := cipher.NewCBCEncrypter(block, iv)
	decrypter := cipher.NewCBCDecrypter(block, iv)

	return &AES256CBC {
		Encrypter: encrypter,
		Decrypter: decrypter,
	}, nil
}

// Encrypt encrypts data with key
func (algo *AES256CBC) Encrypt(msg []byte) ([]byte, error) {
	if f := len(msg)%aes.BlockSize; f != 0 {
		padding := make([]byte, aes.BlockSize-f)
		msg = append(msg, padding...)
	}
	encrypted := make([]byte, len(msg))
	algo.Encrypter.CryptBlocks(encrypted, msg)
	return encrypted, nil
}

// Decrypt decrypts data with key
func (algo *AES256CBC) Decrypt(ciphertext []byte) ([]byte, error) {
	decrypted := make([]byte, len(ciphertext))
	algo.Decrypter.CryptBlocks(decrypted, ciphertext)
	withoutPadding := bytes.ReplaceAll(decrypted, make([]byte, 1), []byte{})
	return withoutPadding, nil
}