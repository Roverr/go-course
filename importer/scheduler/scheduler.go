package scheduler

import (
	"log"

	"github.com/Roverr/go-course/importer/pkg"
	"github.com/Roverr/go-course/importer/scheduler/memory"
)

// Service describes the service function of the queue
type Service struct {
	out   chan string
	queue pkg.Queue
}

// Initialize returns a scheduler
func Initialize() pkg.Scheduler {
	memory := memory.NewQueue()
	service := Service{
		queue: &memory,
		out:   make(chan string, 50),
	}
	go service.schedule()
	return &service
}

// AddListener calls the given worker function when new data arrives
func (s *Service) AddListener(f func(<-chan string)) {
	go f(s.out)
}

// Close closes the channel and ends the connection
func (s *Service) Close() error {
	close(s.out)
	return nil
}

// GetQueue is for accessing the queue of the scheduler
func (s *Service) GetQueue() pkg.Queue {
	if s == nil {
		return nil
	}
	return s.queue
}

func (s *Service) schedule() {
	for {
		data, err := s.queue.GetOne()
		if err != nil {
			log.Fatal(err)
		}
		if data != nil {
			s.out <- string(data)
		}
	}
}
