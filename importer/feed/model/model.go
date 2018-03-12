package model

import (
	"errors"
	"fmt"

	"github.com/Roverr/go-course/importer/pkg"
)

// Type is a string alias
type Type = string

var (
	// Mock is a feed type
	Mock Type = "mock"

	// ErrUnkownType is an error for unknown types
	ErrUnkownType = errors.New("Unknown type from queue")
)

// ErrFeedNotFound is
type ErrFeedNotFound struct {
	which string
}

// Error is the implementation of error interface
func (e ErrFeedNotFound) Error() string {
	return fmt.Sprintf("%s feeds was missing from the map", e.which)
}

// NewErrFeedNotFound initializes a new error for the given feed key
func NewErrFeedNotFound(which string) ErrFeedNotFound {
	return ErrFeedNotFound{which}
}

// Type check for interface implementation
var _ error = (*ErrFeedNotFound)(nil)

// FeedMap describes the way we quickly between feeds
type FeedMap map[Type]pkg.Feed

// Request describes an inner message
type Request struct {
	Type    string      `json:"type"`
	Context interface{} `json:"context"`
}
