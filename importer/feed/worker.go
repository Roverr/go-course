package feed

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/Roverr/go-course/importer/feed/model"
)

// Worker describes a worker in the system
type Worker struct {
	feeds model.FeedMap
	id    int
}

// NewWorker creates a new instance of worker
func NewWorker(feeds model.FeedMap, id int) Worker {
	return Worker{feeds, id}
}

// Listen is a function which gets a job, and calls the correct parser
func (w *Worker) Listen(sh <-chan string) {
	for {
		reqString := <-sh
		fmt.Println(fmt.Sprintf("%d WORKER || Processing %s", w.id, reqString))
		var request model.Request
		json.Unmarshal([]byte(reqString), &request)
		switch request.Type {
		case model.Mock:
			feed, ok := w.feeds[model.Mock]
			if !ok {
				log.Fatal(model.NewErrFeedNotFound(model.Mock))
			}
			err := feed.Crawl()
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(fmt.Sprintf("%d WORKER || Processing finished correctly", w.id))
		default:
			log.Fatal(reqString)
		}

	}
}
