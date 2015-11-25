package releases

import "github.com/alinz/releasifier/data"

type createReleaseRequest struct {
	Platform data.Platform `json:"platform"`
	Note     string        `json:"note"`
	Version  data.Version  `json:"version"`
}

func createReleaseRequestBuilder() interface{} {
	return &createReleaseRequest{}
}
