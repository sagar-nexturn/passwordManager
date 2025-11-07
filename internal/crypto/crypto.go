package crypto

type Crypto interface {
	Encrypt(plaintext []byte) (ciphertext []byte, nonce []byte, err error)
	Decrypt(ciphertext []byte, nonce []byte) ([]byte, error)
}
