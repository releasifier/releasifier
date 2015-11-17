package data

import (
	"encoding/json"
	"fmt"

	"github.com/alinz/releasifier/lib/crypto"
)

var _secureIDKey []byte

//SecureID is an int64 type which does encrypt and decrypt the value
type SecureID int64

//MarshalJSON for type SecureID for encrypting id value
func (id SecureID) MarshalJSON() ([]byte, error) {
	value, err := crypto.EncryptSecureInt64AsBase64(int64(id), _secureIDKey)
	if err != nil {
		return nil, err
	}

	//remove '=' from last value. it makes the id look better :D
	return json.Marshal(value[:len(value)-1])
}

//UnmarshalJSON for type SecureID for decrypting the id
func (id *SecureID) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("ID should be a string, got %s", data)
	}

	//before we decrypt the id, we need to append '=' to the end of s.
	v, err := crypto.DecryptSecureInt64FromBase64(s+"=", _secureIDKey)

	if err != nil {
		return err
	}

	*id = SecureID(v)
	return nil
}

//SetSecureIDKey we need to set this value inside our main.
func SetSecureIDKey(secureKey string) {
	_secureIDKey = []byte(secureKey)
}
