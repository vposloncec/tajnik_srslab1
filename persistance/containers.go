package persistance

// Length a password is "normalized" to.
// This means it is possible for an attacker to know password length up to normLen accuracy
const normLen = 32

// Used to normalize password length so as to hide real password length
// When saving a password, but using the PassWithPadding func
type PasswordContainer struct {
	// Should be set using the Save method, not directly
	Password string `json:"p"`

	PaddingBytes []byte `json:"r"`
}

const paddingRune = 'A'

// Creates PasswordContainer using a set password.
// This container is used to wrap passwords in order to hide length
func PassWithPadding(password string) PasswordContainer {
	c := PasswordContainer{}
	c.Save(password)
	return c
}

func (c *PasswordContainer) Save(password string) {
	c.Password = password
	padLength := normLen - (len(password) % normLen)
	c.PaddingBytes = createPadding(padLength)
}

func createPadding(length int) []byte {
	out := make([]byte, length)
	for i := range out {
		out[i] = paddingRune
	}
	return out
}

func (c *PasswordContainer) Get() string {
	return c.Password
}
