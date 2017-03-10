package config

import (
	"os"

	"github.com/cloudfoundry-incubator/candiedyaml"
)

// Config is the data structure used to return the API configuration
type Config struct {
	APIConfiguration APIConfiguration `yaml:"pagerduty_api"`
}

// APIConfiguration is the data structure used to parse the API configuration
type APIConfiguration struct {
	ServiceKey  string `yaml:"service_key"`
	Description string `yaml:"description"`
}

// ParseConfig parse config file
func ParseConfig(path string) (Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return Config{}, err
	}

	var decodedConfig Config
	if err := candiedyaml.NewDecoder(file).Decode(&decodedConfig); err != nil {
		return Config{}, err
	}
	return decodedConfig, nil
}
