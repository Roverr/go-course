package feed

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	testURL = "http://boston-ace.com"
)

func TestInitialize(t *testing.T) {
	feeds, err := Initialize()
	assert.Nil(t, err)
	assert.NotEmpty(t, feeds)
}
