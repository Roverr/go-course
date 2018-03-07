package mock

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/Roverr/go-course/importer/config"
	"github.com/Roverr/go-course/importer/feed/mock/model"
	"github.com/kelseyhightower/envconfig"
)

// Feed represents a mock feed
type Feed struct {
	URL     string `envconfig:"MOCK_URL" default:"http://localhost:8000"`
	healthy bool
	store   *config.Store
	client  *http.Client
}

// New returns a new feed
func New(store *config.Store) (*Feed, error) {
	var f Feed
	err := envconfig.Process("FEED", &f)
	if err != nil {
		return nil, err
	}
	f.store = store
	f.client = &http.Client{}
	return &f, nil
}

// IsHealthy gives back if the feed is healthy or not
func (f *Feed) IsHealthy() bool {
	if f == nil {
		return false
	}
	return f.healthy
}

// setHealth is for setting the health of the importer
func (f *Feed) setHealth(healthy *bool) {
	if healthy == nil || f == nil {
		return
	}
	f.healthy = *healthy
}

// Crawl is for making a request to the feed and import the data
func (f *Feed) Crawl(errCh chan<- error) {
	// Keep the health indicator to false until the ned of the function
	var healthy bool
	defer f.setHealth(&healthy)
	// Prepare request to URL without request body
	req, err := http.NewRequest("POST", f.URL, bytes.NewBuffer(nil))
	if err != nil {
		errCh <- err
		return
	}
	// Do request
	resp, err := f.client.Do(req)
	if err != nil {
		errCh <- err
		return
	}
	// Read response body to bytes
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		errCh <- err
		return
	}
	// Parse bytes
	var parsed model.APIResponse
	if err = json.Unmarshal(body, &parsed); err != nil {
		errCh <- err
		return
	}
	healthy = true
}

var _ config.Feed = (*Feed)(nil)
