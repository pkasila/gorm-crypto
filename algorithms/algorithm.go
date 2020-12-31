package algorithms

type Algorithm interface {
	Encrypt(msg []byte) ([]byte, error)
	Decrypt(ciphertext []byte) ([]byte, error)
}
