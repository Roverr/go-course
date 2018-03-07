package main

import (
	"log"

	"github.com/Roverr/go-course/importer/feed"
	"github.com/Roverr/go-course/importer/store"
)

func main() {
	store, err := store.Initialize()
	if err != nil {
		log.Fatal(err)
	}
	_, err = feed.Initialize(*store)
	if err != nil {
		log.Fatal(err)
	}
}
