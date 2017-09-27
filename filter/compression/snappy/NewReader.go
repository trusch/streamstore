package snappy

import (
	"io"

	"github.com/golang/snappy"
	"github.com/trusch/streamstore"
)

// NewReader returns a new snappy reader
func NewReader(base io.Reader) (io.ReadCloser, error) {
	snappyReader := snappy.NewReader(base)
	return streamstore.NewIOCoppler(snappyReader, base), nil
}
