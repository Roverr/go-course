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
	store, err := store.Initialize()
	assert.Nil(t, err)
	assert.NotNil(t, store)
	feeds, err := Initialize(*store)
	assert.Nil(t, err)
	assert.NotEmpty(t, feeds)
}
