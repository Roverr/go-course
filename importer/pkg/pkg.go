package pkg

// Feed represents a given feed which should be imported
type Feed interface {
	IsHealthy() bool
	Crawl(chan<- error)
	GetKey() string
}

// Store represents an interface for the stores
type Store interface {
	Insert(key string, v interface{}) error
	Connect() error
}

// Scheduler represents a master which is going to schedule the jobs
// for the other workers
type Scheduler interface {
	AddListener(func(<-chan string))
	Close() error
}

// Queue represents the underlying Store of a Scheduler
type Queue interface {
	GetOne() ([]byte, error)
	Store(data []byte)
}

// Settings represents context for the application to run
type Settings struct {
	feeds     []*Feed
	store     *Store
	scheduler *Scheduler
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
