package utils

import (
	"encoding/json"
	"io"
	"io/ioutil"
)

//StreamJSONToStruct converts stream of json to define struct
func StreamJSONToStruct(r io.Reader, v interface{}) error {
	data, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	return nil
}

//StreamJSONToStructWithLimit similar to StreamJSONToStruct but with limit
func StreamJSONToStructWithLimit(r io.Reader, v interface{}, limit int64) error {
	raw, err := ioutil.ReadAll(io.LimitReader(r, limit))

	if err != nil {
		return err
	}

	if err := json.Unmarshal(raw, &v); err != nil {
		return err
	}

	return nil
}
