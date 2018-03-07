package feed

import (
	"testing"

	"github.com/Roverr/go-course/importer/store"
	"github.com/stretchr/testify/assert"
)

var (
	testURL = "http://boston-ace.com"
)

func TestInitialize(t *testing.T) {
	store.Initialize()
	feeds, err := Initialize(nil)
	assert.Nil(t, err)
	assert.NotEmpty(t, feeds)
}
