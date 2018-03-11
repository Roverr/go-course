package tools

import (
	"context"
	"log"

	"github.com/Roverr/go-course/importer/pkg"
	"golang.org/x/time/rate"
)

var mockRequest = []byte(`{"type":"mock", "context":{}}`)

// Pusher describes a struct which can push events into a store
type Pusher struct {
	queue   pkg.Queue
	limiter *rate.Limiter
}

// NewPusher gives back a new instance of Pusher
func NewPusher(queue pkg.Queue) Pusher {
	pusher := Pusher{
		queue:   queue,
		limiter: rate.NewLimiter(1, 3),
	}
	return pusher
}

// MockPushJobs is creating some jobs from time to time
func (p *Pusher) MockPushJobs() {
	for {
		err := p.limiter.Wait(context.Background())
		if err != nil {
			log.Fatal(err)
		}
		p.queue.Store(mockRequest)
	}
}
