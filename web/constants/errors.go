package constants

import "errors"

var (
	ErrorBodyIssue            = errors.New("request body has some issues") //400
	ErrorFilesTooBig          = errors.New("files are too big")            //400
	ErrorReleaseAlreadyLocked = errors.New("release is already locked")    //400
	ErrorAuthorizeAccess      = errors.New("unauthorized access")          //401
	ErrorReleaseLocked        = errors.New("release is locked")            //403
	ErrorAppNotFound          = errors.New("app not found")                //404
	ErrorReleaseNotFound      = errors.New("release not found")            //404
	ErrorDuplicateVersion     = errors.New("duplicate version")            //409
	ErrorDuplicateName        = errors.New("duplicate name")               //409
)
