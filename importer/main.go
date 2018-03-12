package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Roverr/go-course/importer/config"
	"github.com/Roverr/go-course/importer/feed"
	"github.com/Roverr/go-course/importer/scheduler"
	"github.com/Roverr/go-course/importer/store"
	"github.com/Roverr/go-course/importer/tools"
)

func main() {
	// Initalize store
	store, err := store.Initialize()
	if err != nil {
		log.Fatal(err)
	}

	// Initialize feeds
	feeds, err := feed.Initialize(*store)
	if err != nil {
		log.Fatal(err)
	}

	// Initialize scheduler
	scheduler := scheduler.Initialize()
	queue := scheduler.GetQueue()
	if queue == nil {
		log.Fatal(err)
	}

	// Initialize mock pusher
	pusher := tools.NewPusher(queue)
	go pusher.MockPushJobs()

	// Get application settings
	settings, err := config.NewSettings()

	workers := make([]int, settings.NumberOfWorkers)
	for i := range workers {
		worker := feed.NewWorker(feeds, i)
		scheduler.AddListener(worker.Listen)
	}

	// Start profilling function
	go func() {
		router := initProfiling()
		fmt.Println(fmt.Sprintf("Profiling listening on: %s", settings.ProfilingPort))
		http.ListenAndServe(settings.ProfilingPort, router)
	}()
	<-time.After(325235325235)
}
