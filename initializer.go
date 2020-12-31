package gormcrypto

import (
	"crypto/rsa"
	"github.com/pkosilo/gorm-crypto/algorithms"
	"github.com/pkosilo/gorm-crypto/serialization"
)

var Algorithm algorithms.Algorithm
var Serializer serialization.Serializer

// Init initializes library with specified algorithm
func Init(algorithm algorithms.Algorithm, serializer serialization.Serializer) {
	Algorithm = algorithm
	Serializer = serializer
}

// Deprecated: InitFromKeyPair initializes library with RSA algorithm with specified private and public keys
func InitFromKeyPair(private *rsa.PrivateKey, public *rsa.PublicKey) {
	Init(algorithms.NewRSA(private, public), serialization.NewJSON())
}