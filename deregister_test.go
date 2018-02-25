package drs

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDeregistering(t *testing.T) {
	ConfigSetup()
	_, err := DeregisterUser("")
	assert.NotNil(t, err)
	result, err := DeregisterUser("TEST")
	assert.Nil(t, err)
	assert.True(t, result)
}
