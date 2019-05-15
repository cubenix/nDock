package services

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/gauravgahlot/dockerdoodle/types"
)

// ConfigReader represents the service that reads the configuration,
// either from config.json or database (future)
type ConfigReader interface {
	ReadConfig() (types.Config, error)
}

// JSONConfigReader reads the configuration from config.json
type JSONConfigReader struct{}

// ReadConfig reads confiuration for Docker hosts
func (jcr JSONConfigReader) ReadConfig() (types.Config, error) {
	var config types.Config

	data, err := ioutil.ReadFile("config.json")
	if err != nil {
		log.Fatalln(err)
		return config, err
	}

	jsonErr := json.Unmarshal(data, &config)
	if jsonErr != nil {
		log.Fatalln("Invalid JSON file:", err)
		return config, err
	}
	return config, nil
}
