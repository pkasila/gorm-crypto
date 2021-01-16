package signing

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"errors"
)

type ECDSA struct {
	SignatureAlgorithm
	Private *ecdsa.PrivateKey
	Public  *ecdsa.PublicKey
}

func NewECDSA(privateKey *ecdsa.PrivateKey, publicKey *ecdsa.PublicKey) *ECDSA {
	return &ECDSA{
		Private: privateKey,
		Public:  publicKey,
	}
}

func (a *ECDSA) Sign(msg []byte) ([]byte, error) {
	hash := sha256.Sum256(msg)
	return ecdsa.SignASN1(rand.Reader, a.Private, hash[:])
}

func (a *ECDSA) Verify(msg []byte, signature []byte) (bool, error) {
	hash := sha256.Sum256(msg)
	valid := ecdsa.VerifyASN1(a.Public, hash[:], signature)
	var err error
	if !valid {
		err = errors.New("InvalidSignature")
	}
	return valid, err
}