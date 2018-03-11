package queue

import (
	"log"

	"github.com/Roverr/go-course/importer/pkg"
	"github.com/Roverr/go-course/importer/scheduler/memory"
)

// Service describes the service function of the queue
type Service struct {
	out   chan string
	store pkg.Queue
}

// Initialize returns a scheduler
func Initialize() pkg.Scheduler {
	memory := memory.NewStore()
	service := Service{
		store: &memory,
		out:   make(chan string, 50),
	}
	go service.schedule()
	return &service
}

// AddListener calls the given worker function when new data arrives
func (s *Service) AddListener(f func(<-chan string)) {
	f(s.out)
}

// Close closes the channel and ends the connection
func (s *Service) Close() error {
	close(s.out)
	return nil
}

func (s *Service) schedule() {
	for {
		data, err := s.store.GetOne()
		if err != nil {
			log.Fatal(err)
		}
		if data != nil {
			s.out <- string(data)
		}
	}
}
