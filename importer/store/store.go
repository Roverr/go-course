package store

import (
	"github.com/Roverr/go-course/importer/config"
	"github.com/Roverr/go-course/importer/store/memory"
)

// Initialize runs through the stores, initializes and returns them
func Initialize() (*config.Store, error) {
	var store config.Store
	var err error
	store, err = memory.New()
	return &store, err
}
