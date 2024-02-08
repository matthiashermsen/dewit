package appconfiguration

import (
	"github.com/matthiashermsen/dewit/appconfiguration/api"
	"github.com/matthiashermsen/dewit/appconfiguration/logging"
)

type Configuration struct {
	API     api.Configuration
	Logging logging.Configuration
}

func New() (Configuration, error) {
	apiConfiguration := api.New()
	loggingConfiguration, err := logging.New()

	if err != nil {
		return Configuration{}, err
	}

	appConfiguration := Configuration{
		API:     apiConfiguration,
		Logging: loggingConfiguration,
	}

	return appConfiguration, err
}
