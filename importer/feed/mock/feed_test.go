package mock

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/Roverr/go-course/importer/feed/mock/model"
	"github.com/Roverr/go-course/importer/store/memory"
	"github.com/stretchr/testify/assert"
)

var (
	testURL      = "http://boston-ace.com"
	testKey      = "mock-feed"
	testResponse = []byte(`{"horses":[{"id":1,"name":"Johnsy","lineID":1}],"odds":[{"id":1,"odds":0.999}]}`)
)

func newFeedWithAPI() (*httptest.Server, *Feed, error) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write(testResponse)
	}))
	memoryStore, err := memory.New()
	if err != nil {
		return nil, nil, err
	}
	feed, err := New(memoryStore)
	feed.URL = server.URL
	return server, feed, err
}

func TestNewFeed(t *testing.T) {
	err := os.Setenv("FEED_MOCK_URL", testURL)
	assert.Nil(t, err)
	feed, err := New(nil)
	assert.Nil(t, err)
	assert.NotNil(t, feed)
	assert.Equal(t, feed.URL, testURL)
}

func checkIfHorsesAreTheSame(t *testing.T, a []model.Horse, b []model.Horse) {
	for i, value := range a {
		assert.Equal(t, value.LineID, b[i].LineID)
		assert.Equal(t, value.Name, b[i].Name)
	}
}

func checkIfOddsAreTheSame(t *testing.T, a []model.Odds, b []model.Odds) {
	for i, value := range a {
		assert.Equal(t, value.ID, b[i].ID)
		assert.Equal(t, value.Odd, b[i].Odd)
	}
}
func TestRequest(t *testing.T) {
	t.Run("Should work correctly with mock data", func(t *testing.T) {
		server, feed, err := newFeedWithAPI()
		assert.Nil(t, err)
		assert.NotNil(t, feed)
		assert.NotNil(t, server)
		response, err := feed.request()
		assert.Nil(t, err)
		assert.NotNil(t, response)

		assert.NotEmpty(t, response.Horses)
		assert.NotEmpty(t, response.Odds)
		var parsed model.APIResponse
		err = json.Unmarshal(testResponse, &parsed)
		assert.Nil(t, err)
		assert.Equal(t, len(parsed.Horses), len(response.Horses))
		assert.Equal(t, len(parsed.Odds), len(response.Odds))
		checkIfHorsesAreTheSame(t, response.Horses, parsed.Horses)
		checkIfOddsAreTheSame(t, response.Odds, parsed.Odds)
	})
}
