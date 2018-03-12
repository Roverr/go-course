package feed

import (
	"github.com/Roverr/go-course/importer/feed/mock"
	"github.com/Roverr/go-course/importer/feed/model"
	"github.com/Roverr/go-course/importer/pkg"
)

// Initialize gives back an array of feeds
func Initialize(store pkg.Store) (feeds model.FeedMap, err error) {
	var feed pkg.Feed
	feeds = model.FeedMap{}
	if feed, err = mock.New(store); err != nil {
		return
	}
	feeds[feed.GetKey()] = feed
	return feeds, nil
}
