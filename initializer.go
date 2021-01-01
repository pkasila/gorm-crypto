package gormcrypto

import (
	"github.com/pkosilo/gorm-crypto/algorithms"
	"github.com/pkosilo/gorm-crypto/serialization"
)

var Algorithms []algorithms.Algorithm
var Serializers []serialization.Serializer

// InitWithFallbacks initializes library with fallback methods
func InitWithFallbacks(algorithms []algorithms.Algorithm, serializers []serialization.Serializer) {
	if len(algorithms) == 0 || len(serializers) == 0 {
		panic("There are no algorithms and/or serializers passed!")
	}

	Algorithms = algorithms
	Serializers = serializers
}

// Init initializes library with specified algorithm
func Init(algorithm algorithms.Algorithm, serializer serialization.Serializer) {
	InitWithFallbacks([]algorithms.Algorithm{algorithm}, []serialization.Serializer{serializer})
}
