package memory

import (
	"container/list"
)

// Store describes a memory store
type Store struct {
	data *list.List
}

// NewStore returns a new instance of Store
func NewStore() Store {
	return Store{
		data: list.New(),
	}
}

// GetOne returns an element from the store
func (s Store) GetOne() ([]byte, error) {
	e := s.data.Front()
	if e == nil {
		return nil, nil
	}
	return e.Value.([]byte), nil
}

// Store function stores the the given array of byte
func (s *Store) Store(data []byte) {
	if s == nil {
		return
	}
	s.data.PushBack(data)
}
