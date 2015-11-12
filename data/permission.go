package data

import (
	"encoding/json"
	"fmt"
)

//Permission represents type of bundle store
type Permission int

const (
	//FULL represents source code in JS
	FULL Permission = iota
	//READONLY represents picture and image types
	READONLY
)

var (
	permissionNameToValue = map[string]Permission{
		"FULL":     FULL,
		"READONLY": READONLY,
	}

	permissionValueToName = map[Permission]string{
		FULL:     "FULL",
		READONLY: "READONLY",
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
