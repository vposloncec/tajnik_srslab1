package cripter

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
    "golang.org/x/crypto/bcrypt"
	"io"
)

func Encrypt(in []byte, passphrase string) ([]byte, error){
	aesgcm, err := newAesGCM(passphrase)
	if err != nil {
		return nil, err
	}
	nonce := make([]byte, aesgcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}
	return aesgcm.Seal(nonce, nonce, in, nil), nil
}

func Decrypt(in []byte, passphrase string)([]byte, error){
	aesgcm, err := newAesGCM(passphrase)
	if err != nil {
		return nil, err
	}

	// Nonce does not have to be secure (only unique) so it is stored on the beginning of the ciphertext
	nonce, ciphertext := in[:aesgcm.NonceSize()], in[aesgcm.NonceSize():]

	return aesgcm.Open(nil, nonce, ciphertext, nil)
}

func newAesGCM(passphrase string) (cipher.AEAD, error){
	ckey, err := bcrypt.GenerateFromPassword([]byte(passphrase), 15)
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