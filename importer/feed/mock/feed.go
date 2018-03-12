package mock

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/Roverr/go-course/importer/feed/mock/model"
	"github.com/Roverr/go-course/importer/pkg"
	"github.com/kelseyhightower/envconfig"
)

// Feed represents a mock feed
type Feed struct {
	URL     string `envconfig:"MOCK_URL" default:"http://localhost:8000"`
	Key     string `default:"mock"`
	healthy bool
	store   pkg.Store
	client  *http.Client
}

// New returns a new feed
func New(store pkg.Store) (*Feed, error) {
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

// GetKey returns a key for the feed
func (f Feed) GetKey() string {
	return f.Key
}

func (f *Feed) request() (*model.APIResponse, error) {
	// Prepare request to URL without request body
	req, err := http.NewRequest("GET", f.URL, nil)
	if err != nil {
		return nil, err
	}
	// Do request
	resp, err := f.client.Do(req)
	if err != nil {
		return nil, err
	}
	// Read response body to bytes
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	// Parse bytes
	var parsed model.APIResponse
	if err = json.Unmarshal(body, &parsed); err != nil {
		return nil, err
	}
	return &parsed, err
}

// Crawl is for making a request to the feed and import the data
func (f *Feed) Crawl() (err error) {
	// Keep the health indicator to false until the ned of the function
	var healthy bool
	defer f.setHealth(&healthy)

	// Make the request
	var resp *model.APIResponse
	resp, err = f.request()
	if err != nil {
		return
	}
	if err = f.store.Insert(f.Key, *resp); err != nil {
		return
	}

	healthy = true
	return
}

var _ pkg.Feed = (*Feed)(nil)
