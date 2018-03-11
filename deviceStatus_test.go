package drs

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeviceStatus(t *testing.T) {
	ConfigSetup()
	result, err := UpdateDeviceStatus("", "")
	assert.NotNil(t, err)
	assert.False(t, result)
	if apiError, ok := err.(*APIError); ok {
		assert.Equal(t, http.StatusBadRequest, apiError.Code)
	}

	result, err = UpdateDeviceStatus("TEST", "2005-11-06 01:54:00")
	assert.NotNil(t, err)
	assert.False(t, result)
	if apiError, ok := err.(*APIError); ok {
		assert.Equal(t, http.StatusBadRequest, apiError.Code)
	}

	result, err = UpdateDeviceStatus("TEST", "2008-01-01T17:08:00Z")
	assert.Nil(t, err)
	assert.True(t, result)

	result, err = UpdateDeviceStatus("TEST", "")
	assert.Nil(t, err)
	assert.True(t, result)
}
