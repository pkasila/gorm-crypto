package gormcrypto

import (
	"crypto/rsa"
	"github.com/pkosilo/gorm-crypto/algorithms"
)

var Algorithm algorithms.Algorithm

// Init initializes library with specified algorithm
func Init(algorithm algorithms.Algorithm) {
	Algorithm = algorithm
}

// Deprecated: InitFromKeyPair initializes library with RSA algorithm with specified private and public keys
func InitFromKeyPair(private *rsa.PrivateKey, public *rsa.PublicKey) {
	Init(algorithms.NewRSA(private, public))
}