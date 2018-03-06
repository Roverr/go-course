package store

import (
	"github.com/Roverr/go-course/importer/config"
	"github.com/Roverr/go-course/importer/store/memory"
)

// Initialize runs through the stores, initializes and returns them
func Initialize() (stores []*config.Store, err error) {
	var store config.Store
	if store, err = memory.New(); err != nil {
		return
	}
	stores = append(stores, &store)
	return stores, nil
}
