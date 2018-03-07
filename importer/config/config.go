package config

import "log"

// Feed represents a given feed which should be imported
type Feed interface {
	IsHealthy() bool
	Crawl(chan<- error)
}

// Store represents an interface for the stores
type Store interface {
	Insert(key string, v interface{}) error
	Connect() error
}

// JobQueue represents a queue which is for getting jobs
type JobQueue interface {
	GetOne() []byte
	Close() error
}

// Settings represents context for the application to run
type Settings struct {
	feeds []*Feed
	store *Store
	queue *JobQueue
}

// NewSettings return a pointer of Settings
func NewSettings() *Settings {
	return &Settings{}
}

// AddFeed is for adding a new feed to the feed list
func (s *Settings) AddFeed(f *Feed) {
	if s == nil {
		return
	}
	s.feeds = append(s.feeds, f)
}

// SetQueue is for setting the queue for the application
func (s *Settings) SetQueue(queue *JobQueue) {
	if s == nil {
		return
	}
	if s.queue != nil {
		if err := (*s.queue).Close(); err != nil {
			log.Fatal(err)
		}
	}
	s.queue = queue
}
