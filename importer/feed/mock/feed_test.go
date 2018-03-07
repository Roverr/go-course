package mock

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testURL = "http://boston-ace.com"

func TestNewFeed(t *testing.T) {
	err := os.Setenv("FEED_MOCK_URL", testURL)
	assert.Nil(t, err)
	feed, err := New(nil)
	assert.Nil(t, err)
	assert.NotNil(t, feed)
	assert.Equal(t, feed.URL, testURL)
}
