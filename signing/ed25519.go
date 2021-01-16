package signing

import (
	"crypto/ed25519"
)

type Ed25519 struct {
	SignatureAlgorithm
	Private *ed25519.PrivateKey
	Public  *ed25519.PublicKey
}

func NewEd25519(privateKey *ed25519.PrivateKey, publicKey *ed25519.PublicKey) *Ed25519 {
	return &Ed25519{
		Private: privateKey,
		Public:  publicKey,
	}
}

func (a *Ed25519) Sign(msg []byte) ([]byte, error) {
	return ed25519.Sign(*a.Private, msg), nil
}

func (a *Ed25519) Verify(msg []byte, signature []byte) (bool, error) {
	return ed25519.Verify(*a.Public, msg, signature), nil
}