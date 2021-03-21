package cripter

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"

	"golang.org/x/crypto/pbkdf2"
	"golang.org/x/crypto/sha3"
)

// create 32bit long key from passphrase
func createKey(passphrase string) []byte {
	return pbkdf2.Key([]byte(passphrase),
		[]byte("Ča Je Život Vengo Fantažija"),
		4096,
		32,
		sha3.New224)
}

func Encrypt(in []byte, passphrase string) ([]byte, error) {
	aesgcm, err := newAesGCM(passphrase)
	if err != nil {
		return nil, err
	}
	nonce := make([]byte, aesgcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}
	b, err := aesgcm.Seal(nonce, nonce, in, nil), nil
	return b, err
}

func Decrypt(in []byte, passphrase string) ([]byte, error) {
	aesgcm, err := newAesGCM(passphrase)
	if err != nil {
		return nil, err
	}

	// Nonce does not have to be secure (only unique) so it is stored on the beginning of the ciphertext
	nonce, ciphertext := in[:aesgcm.NonceSize()], in[aesgcm.NonceSize():]

	return aesgcm.Open(nil, nonce, ciphertext, nil)
}

func newAesGCM(passphrase string) (cipher.AEAD, error) {
	ckey := createKey(passphrase)
	c, err := aes.NewCipher(ckey)
	if err != nil {
		return nil, err
	}
	aesgcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, err
	}
	return aesgcm, nil
}
