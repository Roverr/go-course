package memory

import "github.com/Roverr/go-course/importer/pkg"

// Store represents a memory store
type Store struct {
	records map[string]interface{}
}

// New returns a new Store instance
func New() (*Store, error) {
	db := Store{}
	return &db, nil
}

// Insert represents a function which inserts a value interface to a key section
func (d *Store) Insert(key string, v interface{}) error {
	d.records[key] = v
	return nil
}

// Connect represents the process of connection
func (d *Store) Connect() error {
	if d == nil {
		return nil
	}
	d.records = make(map[string]interface{})
	return nil
}

// Interface validation
var _ pkg.Store = (*Store)(nil)
