package data

import (
	"encoding/json"
	"fmt"
)

//Permission represents type of bundle store
type Permission int

const (
	OWNER Permission = iota
	//WRITE represents source code in JS
	WRITE
	//READ represents picture and image types
	READ
)

var (
	permissionNameToValue = map[string]Permission{
		"WRITE": WRITE,
		"READ":  READ,
	}

	permissionValueToName = map[Permission]string{
		WRITE: "WRITE",
		READ:  "READ",
	}
)

//MarshalJSON for type Permission
func (p Permission) MarshalJSON() ([]byte, error) {
	if s, ok := interface{}(p).(fmt.Stringer); ok {
		return json.Marshal(s.String())
	}
	s, ok := permissionValueToName[p]
	if !ok {
		return nil, fmt.Errorf("invalid Permission: %d", p)
	}
	return json.Marshal(s)
}

//UnmarshalJSON for type Permission
func (p *Permission) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("Permission should be a string, got %s", data)
	}
	v, ok := permissionNameToValue[s]
	if !ok {
		return fmt.Errorf("invalid Permission %q", s)
	}
	*p = v
	return nil
}
