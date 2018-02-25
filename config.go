package drs

import (
	"github.com/sirupsen/logrus"
	"os"
	"strings"
)

// Config is the global configuration object that holds global configuration settings
var Config *ConfigStruct

//ConfigStruct holds the various configuration options
type ConfigStruct struct {
	Environment string
	RootURL     string
}

//ConfigSetup sets up the config struct with data from the environment
func ConfigSetup() *ConfigStruct {
	//a big change as of Jan 2018: Version is included in the URL, so no string generation is needed in the calls
	c := new(ConfigStruct)

	c.Environment = strings.ToLower(os.Getenv("ENV"))
	if c.Environment == "prod" || c.Environment == "production" {
		c.Environment = "production"
	} else if c.Environment == "" || c.Environment == "dev" || c.Environment == "development" {
		c.Environment = "dev"
	}

	c.RootURL = strings.ToLower(os.Getenv("ROOT_URL"))
	if c.RootURL == "" {
		//we set to the default as of 20180224
		c.RootURL = "https://dash-replenishment-service-na.amazon.com/"
	}
	if !strings.HasSuffix(c.RootURL, "/") {
		c.RootURL += "/"
	}

	Config = c

	//setup now
	return c
}

//init is called when the host application starts up and sets the Configuration and logging settings
func init() {
	ConfigSetup()
	logrus.SetFormatter(&logrus.JSONFormatter{})
}

//log provides structured logging through logrus. We support info, warning, and error
func log(level, key, message string, data interface{}) {
	if Config.Environment != "production" {
		level = strings.ToLower(level)

		fields := logrus.Fields{
			"key":  key,
			"data": data,
		}

		switch level {
		case "info":
			logrus.WithFields(fields).Info(message)
		case "warning":
			logrus.WithFields(fields).Warning(message)
		case "error":
			logrus.WithFields(fields).Error(message)
		}
	}
}
