package store

import (
	"github.com/Roverr/go-course/importer/pkg"
	"github.com/Roverr/go-course/importer/store/memory"
)

// Initialize runs through the stores, initializes and returns them
func Initialize() (*pkg.Store, error) {
	var store pkg.Store
	var err error
	store, err = memory.New()
	return &store, err
}
