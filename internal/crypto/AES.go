package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"
	"os"
)

type AESCrypto struct{}

func getKey() ([]byte, error) {
	b := os.Getenv("MASTER_KEY")
	if b == "" {
		return nil, errors.New("MASTER_KEY not set")
	}
	k, err := base64.StdEncoding.DecodeString(b)
	if err != nil {
		return nil, err
	}
	if len(k) != 32 {
		return nil, errors.New("MASTER_KEY must decode to 32 bytes (AES-256)")
	}
	return k, nil
}

func (m *AESCrypto) Encrypt(plaintext []byte) (ciphertext []byte, nonce []byte, err error) {
	key, err := getKey()
	if err != nil {
		return nil, nil, err
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, nil, err
	}
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, nil, err
	}
	nonce = make([]byte, aesgcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, nil, err
	}
	ciphertext = aesgcm.Seal(nil, nonce, plaintext, nil)
	return ciphertext, nonce, nil
}

func (m *AESCrypto) Decrypt(ciphertext []byte, nonce []byte) ([]byte, error) {
	key, err := getKey()
	if err != nil {
		return nil, err
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	pt, err := aesgcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}
	return pt, nil
}

func NewAESCrypto() *AESCrypto {
	return &AESCrypto{}
}
