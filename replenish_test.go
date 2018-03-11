package drs

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReplenishment(t *testing.T) {
	ConfigSetup()
	badRes, badErr := ReplenishSlot("", "")
	assert.Nil(t, badRes)
	assert.NotNil(t, badErr)

	result, err := ReplenishSlot("TEST", "1234")
	assert.Nil(t, err)
	assert.Equal(t, "STANDARD_ORDER_PLACED", result.DetailCode)
	assert.Equal(t, "SOME_EVENT_INSTANCE", result.EventInstanceID)

	result, err = ReplenishSlot("BadToken", "1234")
	assert.Nil(t, result)
	assert.NotNil(t, err)

	if apiError, ok := err.(*APIError); ok {
		assert.Equal(t, http.StatusBadRequest, apiError.Code)
	}
}
