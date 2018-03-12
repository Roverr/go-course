package config

import "github.com/kelseyhightower/envconfig"

// Settings describes the global configuration of the application
type Settings struct {
	NumberOfWorkers int    `envconfig:"NUMBER_OF_WORKERS" default:"3"`
	ProfilingPort   string `envconfig:"PROFILING_PORT" default:":8080"`
}

// NewSettings returns a new instance of settings
func NewSettings() (settings Settings, err error) {
	err = envconfig.Process("IMPORTER", &settings)
	return
}
