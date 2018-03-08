package drs

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogging(t *testing.T) {
	ConfigSetup()
	logOutput := log("info", "test", "A test message", map[string]string{"code": "amzKVSS1234"})
	assert.Equal(t, "INFO: test", logOutput)
	logOutput = log("warning", "test", "A test message", map[string]string{"code": "amzKVSS1234"})
	assert.Equal(t, "WARNING: test", logOutput)
	logOutput = log("error", "test", "A test message", map[string]string{"code": "amzKVSS1234"})
	assert.Equal(t, "ERROR: test", logOutput)

	Config.Environment = "production"
	logOutput = log("info", "test", "A test message", map[string]string{"code": "amzKVSS1234"})
	assert.Equal(t, "", logOutput)
	logOutput = log("warning", "test", "A test message", map[string]string{"code": "amzKVSS1234"})
	assert.Equal(t, "", logOutput)
	logOutput = log("error", "test", "A test message", map[string]string{"code": "amzKVSS1234"})
	assert.Equal(t, "", logOutput)

	Config.Environment = "test"
}

func TestEnvironment(t *testing.T) {
	originalEnv := os.Getenv("DRS_SDK_ENV")
	os.Setenv("DRS_SDK_ENV", "production")
	os.Setenv("DRS_SDK_ROOT_URL", "https://dash-replenishment-service-na.amazon.com")
	ConfigSetup()
	assert.Equal(t, "production", Config.Environment)
	assert.Equal(t, "https://dash-replenishment-service-na.amazon.com/", Config.RootURL)
	os.Setenv("DRS_SDK_ENV", originalEnv)
	ConfigSetup()
}
