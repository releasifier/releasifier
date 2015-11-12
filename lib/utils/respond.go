package utils

import "net/http"

//Respond is a utility function which helps identifies error from other messages
func Respond(w http.ResponseWriter, status int, v interface{}) {
	if err, ok := v.(error); ok {
		WriteJSONTo(w, status, map[string]interface{}{"error": err.Error()})
		return
	}

	WriteJSONTo(w, status, v)
}
