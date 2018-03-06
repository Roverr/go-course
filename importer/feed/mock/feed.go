package mock

import (
	"fmt"

	"github.com/Roverr/go-course/importer/config"
	"github.com/kelseyhightower/envconfig"
)

// Feed represents a mock feed
type Feed struct {
	URL     string `envconfig:"MOCK_URL" default:"http://localhost:8000"`
	healthy bool
}

// New returns a new feed
func New() (*Feed, error) {
	var f Feed
	err := envconfig.Process("FEED", &f)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &f, nil
}

// IsHealthy gives back if the feed is healthy or not
func (f *Feed) IsHealthy() bool {
	return f.healthy
}

// Crawl is for making a request to the feed and import the data
func (f *Feed) Crawl() error {
	return nil
}

var _ config.Feed = (*Feed)(nil)
