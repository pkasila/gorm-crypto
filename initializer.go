package gormcrypto

import (
	"github.com/pkasila/gorm-crypto/algorithms"
	"github.com/pkasila/gorm-crypto/serialization"
	"github.com/pkasila/gorm-crypto/signing"
)

var Algorithms []algorithms.Algorithm
var Serializers []serialization.Serializer
var SignatureAlgorithms []signing.SignatureAlgorithm
var SignatureSerializers []serialization.Serializer

// InitWithFallbacks initializes library with fallback methods
func InitWithFallbacks(algorithms []algorithms.Algorithm, serializers []serialization.Serializer) {
	if len(algorithms) == 0 || len(serializers) == 0 {
		panic("There are no algorithms and/or serializers passed!")
	}

	Algorithms = algorithms
	Serializers = serializers
}

// InitSigning initializes library's signing capabilities
func InitSigning(algorithms []signing.SignatureAlgorithm, serializers []serialization.Serializer) {
	if len(algorithms) == 0 || len(serializers) == 0 {
		panic("There are no algorithms and/or serializers passed!")
	}

	SignatureAlgorithms = algorithms
	SignatureSerializers = serializers
}

// Init initializes library with specified algorithm
func Init(algorithm algorithms.Algorithm, serializer serialization.Serializer) {
	InitWithFallbacks([]algorithms.Algorithm{algorithm}, []serialization.Serializer{serializer})
}
