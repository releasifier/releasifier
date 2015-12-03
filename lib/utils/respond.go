package utils

import (
	"net/http"

	"github.com/alinz/releasifier/errors"
)

//Respond is a utility function which helps identifies error from other messages
func Respond(w http.ResponseWriter, status int, v interface{}) {
	if err, ok := v.(error); ok {
		message := map[string]interface{}{"error": err.Error()}
		WriteJSONTo(w, status, message)
		return
	}

	if v != nil {
		WriteJSONTo(w, status, v)
	} else {
		w.WriteHeader(status)
	}
}

//RespondEx integrated error status code
func RespondEx(w http.ResponseWriter, response interface{}, statusCode int, err error) {
	if err != nil {

		if statusCode == 0 {
			statusCode = errors.GetErrorStatusCode(err)
		} else {
			statusCode = 400
		}
		Respond(w, statusCode, err)
		return
	}

	if statusCode == 0 {
		statusCode = http.StatusOK
	}

	Respond(w, statusCode, response)
}
