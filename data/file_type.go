package data

import (
	"encoding/json"
	"fmt"
	"strings"
)

//Type represents type of bundle store
type FileType int

const (
	//CODE represents source code in JS
	CODE FileType = iota
	//IMAGE represents picture and image types
	IMAGE
)

var (
	typeNameToValue = map[string]FileType{
		"CODE":  CODE,
		"IMAGE": IMAGE,
	}

	typeValueToName = map[FileType]string{
		CODE:  "CODE",
		IMAGE: "IMAGE",
	}
)

//MarshalJSON for type FileType
func (a FileType) MarshalJSON() ([]byte, error) {
	if s, ok := interface{}(a).(fmt.Stringer); ok {
		return json.Marshal(s.String())
	}
	s, ok := typeValueToName[a]
	if !ok {
		return nil, fmt.Errorf("invalid Type: %d", a)
	}
	return json.Marshal(strings.ToLower(s))
}

//UnmarshalJSON for type Type
func (a *FileType) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("Type should be a string, got %s", data)
	}
	v, ok := typeNameToValue[strings.ToUpper(s)]
	if !ok {
		return fmt.Errorf("invalid Type %q", s)
	}
	*a = v
	return nil
}
