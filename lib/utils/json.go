package utils

import (
	"encoding/json"
	"net/http"
)

//WriteJSONTo writes any interface to http writer as json
func WriteJSONTo(w http.ResponseWriter, status int, response interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(response)
}
