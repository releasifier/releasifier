package utils_test

import (
	"testing"

	"github.com/alinz/releasifier/lib/utils"
	"github.com/stretchr/testify/assert"
)

func TestVersion(t *testing.T) {

	type TestMessage struct {
		Message *string `json:"message,required"`
	}

	//value := "Hello World!"
	testMessage := TestMessage{}

	err := utils.JSONValidation(testMessage)

	assert.NotEqual(t, nil, err, "should be an error")
}
