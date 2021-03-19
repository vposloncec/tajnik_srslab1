package persistance

import (
	"encoding/json"
	"github.com/vposloncec/lab1-srs/cripter"
	"io"
	"io/ioutil"
)

type Storage map[string]string

func (s *Storage) LoadDecrypt(passphrase string, r io.Reader) error {
	bCipher, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}
	b, err := cripter.Decrypt(bCipher, passphrase)
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
	bCipher, err := cripter.Encrypt(b, passphrase)
	if err != nil {
		return err
	}

	if _, err := w.Write(bCipher); err != nil{
		return err
	}
	return nil
}
