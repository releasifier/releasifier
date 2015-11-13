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
	return json.Marshal(value)
}

//UnmarshalJSON for type SecureID for decrypting the id
func (id *SecureID) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("ID should be a string, got %s", data)
	}

	v, err := crypto.DecryptSecureInt64FromBase64(s, _secureIDKey)

	if err != nil {
		return err
	}

	*id = SecureID(v)
	return nil
}

//MarshalDB converts SecureID to int64 so it can be store properly to db
func (id SecureID) MarshalDB() (interface{}, error) {
	return int64(id), nil
}

//UnmarshalDB converts int64 to SecureID
func (id *SecureID) UnmarshalDB(v interface{}) error {
	val, ok := v.(int64)
	if !ok {
		return fmt.Errorf("id is not int64")
	}
	*id = SecureID(val)
	return nil
}

//SetSecureIDKey we need to set this value inside our main.
func SetSecureIDKey(secureKey string) {
	_secureIDKey = []byte(secureKey)
}
