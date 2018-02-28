package drs

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDeregistering(t *testing.T) {
	ConfigSetup()
	_, err := DeregisterDevice("")
	assert.NotNil(t, err)
	result, err := DeregisterDevice("TEST")
	assert.Nil(t, err)
	assert.True(t, result)
}
