package errors

import (
	err "errors"
	"net/http"
)

var (
	ErrorBodyIssue            = err.New("request body has some issues") //400
	ErrorFilesTooBig          = err.New("files are too big")            //400
	ErrorReleaseAlreadyLocked = err.New("release is already locked")    //400
	ErrorAlreadyAcceessed     = err.New("already accessed to this app") //400
	ErrorAuthorizationDenied  = err.New("authorization denied")         //401
	ErrorAuthorizeAccess      = err.New("unauthorized access")          //401
	ErrorReleaseLocked        = err.New("release is locked")            //403
	ErrorAppNotFound          = err.New("app not found")                //404
	ErrorReleaseNotFound      = err.New("release not found")            //404
	ErrorDuplicateEmail       = err.New("duplicate email")              //409
	ErrorDuplicateVersion     = err.New("duplicate version")            //409
	ErrorDuplicateName        = err.New("duplicate name")               //409
)

//GetErrorStatusCode returns the proper http status code based on error
func GetErrorStatusCode(err error) int {
	switch err {
	case ErrorBodyIssue:
		return http.StatusBadRequest
	case ErrorFilesTooBig:
		return http.StatusBadRequest
	case ErrorReleaseAlreadyLocked:
		return http.StatusBadRequest
	case ErrorAlreadyAcceessed:
		return http.StatusBadRequest
	case ErrorAuthorizationDenied:
		return http.StatusUnauthorized
	case ErrorAuthorizeAccess:
		return http.StatusUnauthorized
	case ErrorReleaseLocked:
		return http.StatusForbidden
	case ErrorAppNotFound:
		return http.StatusNotFound
	case ErrorReleaseNotFound:
		return http.StatusNotFound
	case ErrorDuplicateEmail:
		return http.StatusConflict
	case ErrorDuplicateVersion:
		return http.StatusConflict
	case ErrorDuplicateName:
		return http.StatusConflict
	}

	return http.StatusBadRequest
}
