package drs

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAPIError(t *testing.T) {
	ConfigSetup()
	err := APIError{
		Code: http.StatusBadRequest,
		Data: map[string]string{
			"message": "Testing the error message",
		},
	}
	assert.NotNil(t, err.Data)
	parsed := err.Error()
	assert.Equal(t, "400", parsed)
}

func TestMakeCall(t *testing.T) {
	ConfigSetup()

	//test an endpoint that doesn't exist
	code, _, err := makeCall("/platypus", nil, "", nil)
	assert.Equal(t, http.StatusNotFound, code)
	assert.NotNil(t, err)

	code, _, err = makeCall("testDelete", nil, "", nil)
	assert.Equal(t, http.StatusNotFound, code)
	assert.NotNil(t, err)
}
