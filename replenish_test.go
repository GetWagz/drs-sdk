package drs

import (
	"github.com/stretchr/testify/assert"
	"testing"
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
	assert.Equal(t, 400, err.Code)

}
