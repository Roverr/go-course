package config

import "github.com/kelseyhightower/envconfig"

// Settings describes the global configuration of the application
type Settings struct {
	NumberOfWorkers int `envonfig:"NUMBER_OF_WORKERS" default:"3"`
}

// NewSettings returns a new instance of settings
func NewSettings() (settings Settings, err error) {
	err = envconfig.Process("IMPORTER", &settings)
	return
}
