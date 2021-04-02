package persistance

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"

	"github.com/vposloncec/lab1-srs/cripter"
)

type Storage map[string]PasswordContainer

func LoadDecrypt(passphrase string, r io.Reader) (Storage, error) {
	s := Storage{}
	if err := s.LoadDecrypt(passphrase, r); err != nil {
		return nil, err
	}
	return s, nil
}

func (s *Storage) LoadDecrypt(passphrase string, r io.Reader) error {
	bCipher, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}
	b, err := cripter.Decrypt(bCipher, passphrase)
	if err != nil {
		return fmt.Errorf("Master file could not be decrypted, perhaps the master password was wrong?")
	}
	// fmt.Println("Decrypted! ", string(b))
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	return nil
}

func (s *Storage) SaveEncrypt(passphrase string, w io.Writer) error {
	b, err := json.Marshal(&s)
	if err != nil {
		return err
	}
	// fmt.Println("saving: ", string(b))
	bCipher, err := cripter.Encrypt(b, passphrase)
	if err != nil {
		return err
	}

	if _, err := w.Write(bCipher); err != nil {
		return err
	}
	return nil
}
