package data

import (
	"encoding/json"
	"fmt"
	"strings"
)

//Permission represents type of bundle store
type Permission int

const (
	//OWNER who ever creates the app first
	OWNER Permission = iota
	//WRITE represents source code in JS
	WRITE
	//READ represents picture and image types
	READ
)

var (
	permissionNameToValue = map[string]Permission{
		"OWNER": OWNER,
		"WRITE": WRITE,
		"READ":  READ,
	}

	permissionValueToName = map[Permission]string{
		OWNER: "OWNER",
		WRITE: "WRITE",
		READ:  "READ",
	}
)

//MarshalJSON for type Permission
func (p Permission) MarshalJSON() ([]byte, error) {
	if s, ok := interface{}(p).(fmt.Stringer); ok {
		return json.Marshal(strings.ToLower(s.String()))
	}
	s, ok := permissionValueToName[p]
	if !ok {
		return nil, fmt.Errorf("invalid Permission: %d", p)
	}

	return json.Marshal(strings.ToLower(s))
}

//UnmarshalJSON for type Permission
func (p *Permission) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("Permission should be a string, got %s", data)
	}
	v, ok := permissionNameToValue[strings.ToUpper(s)]
	if !ok {
		return fmt.Errorf("invalid Permission %q", s)
	}
	*p = v
	return nil
}
