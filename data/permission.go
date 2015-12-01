package data

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

//Permission represents type of bundle store
type Permission int

const (
	//OWNER can delete the project and make other people OWNER
	OWNER Permission = iota
	//ADMIN can invite other users and publish a build
	ADMIN
	//MEMBER can publish a build only
	MEMBER
	//ANONYMOUSE nothing
	ANONYMOUSE
)

var (
	permissionNameToValue = map[string]Permission{
		"OWNER":  OWNER,
		"ADMIN":  ADMIN,
		"MEMBER": MEMBER,
	}

	permissionValueToName = map[Permission]string{
		OWNER:  "OWNER",
		ADMIN:  "ADMIN",
		MEMBER: "MEMBER",
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

//GetPermissionByName tries to conver string to permission value
func GetPermissionByName(name string) (Permission, error) {
	if v, ok := permissionNameToValue[name]; ok {
		return v, nil
	}
	return ANONYMOUSE, errors.New("Permission is incorrect")
}
