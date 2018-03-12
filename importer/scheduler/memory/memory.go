package memory

import (
	"container/list"

	"github.com/Roverr/go-course/importer/pkg"
)

// Queue describes a memory store
type Queue struct {
	data *list.List
}

// NewQueue returns a new instance of Queue
func NewQueue() Queue {
	return Queue{
		data: list.New(),
	}
}

// GetOne returns an element from the store
func (s Queue) GetOne() ([]byte, error) {
	e := s.data.Front()
	if e == nil {
		return nil, nil
	}
	s.data.Remove(e)
	return e.Value.([]byte), nil
}

// Store function stores the the given array of byte
func (s *Queue) Store(data []byte) {
	if s == nil {
		return
	}
	s.data.PushBack(data)
}

// Type check
var _ pkg.Queue = (*Queue)(nil)
