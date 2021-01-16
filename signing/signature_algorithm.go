package signing

type SignatureAlgorithm interface {
	Sign(msg []byte) ([]byte, error)
	Verify(msg []byte, signature []byte) (bool, error)
}