package cripter

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"

	"golang.org/x/crypto/pbkdf2"
	"golang.org/x/crypto/sha3"
)

const saltSize = 32

// create 32bit long key from passphrase
func createKey(passphrase string, salt []byte) ([]byte, error) {
	if len(salt) <= 8 {
		return nil, fmt.Errorf("solt has to be atleast 8 bytes long")
	}

	return pbkdf2.Key([]byte(passphrase),
		salt,
		4096,
		32,
		sha3.New224), nil
}

func Encrypt(in []byte, passphrase string) ([]byte, error) {
	// Create salt
	salt := make([]byte, saltSize)
	if _, err := io.ReadFull(rand.Reader, salt); err != nil {
		return nil, err
	}

	aesgcm, err := newAesGCM(passphrase, salt)
	if err != nil {
		return nil, err
	}

	// Create nonce
	nonce := make([]byte, aesgcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	prependData := append(salt, nonce...)
	b, err := aesgcm.Seal(prependData, nonce, in, nil), nil
	return b, err
}

func Decrypt(in []byte, passphrase string) ([]byte, error) {
	salt, in := in[:saltSize], in[saltSize:]
	aesgcm, err := newAesGCM(passphrase, salt)
	if err != nil {
		return nil, err
	}

	// Nonce does not have to be secure (only unique) so it is stored on the beginning of the ciphertext after password salt
	nonce, ciphertext := in[:aesgcm.NonceSize()], in[aesgcm.NonceSize():]
	return aesgcm.Open(nil, nonce, ciphertext, nil)
}

func newAesGCM(passphrase string, salt []byte) (cipher.AEAD, error) {
	ckey, err := createKey(passphrase, salt)
	if err != nil {
		return nil, err
	}
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
