package gorm_crypto

import "crypto/rsa"

var PrivateKey *rsa.PrivateKey
var PublicKey *rsa.PublicKey

func InitFromKeyPair(private *rsa.PrivateKey, public *rsa.PublicKey) {
	PrivateKey = private
	PublicKey = public
}
