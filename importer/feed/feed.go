package feed

import (
	"github.com/Roverr/go-course/importer/config"
	"github.com/Roverr/go-course/importer/feed/mock"
)

// Initialize gives back an array of feeds
func Initialize() (feeds []*config.Feed, err error) {
	var feed config.Feed
	if feed, err = mock.New(); err != nil {
		return
	}
	feeds = append(feeds, &feed)
	return feeds, nil
}