package drs

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAPIError(t *testing.T) {
	ConfigSetup()
	err := APIError{
		Code: 400,
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
	assert.Equal(t, code, 404)
	assert.NotNil(t, err)
}
