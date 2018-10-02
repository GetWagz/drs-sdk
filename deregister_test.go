package drs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeregistering(t *testing.T) {
	ConfigSetup()
	_, err := DeregisterDevice("")
	assert.NotNil(t, err)
	result, err := DeregisterDevice("TEST")
	assert.Nil(t, err)
	assert.True(t, result)
	result, err = DeregisterDevice("BadData")
	assert.NotNil(t, err)
	assert.False(t, result)
}
