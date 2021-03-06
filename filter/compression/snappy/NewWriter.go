package snappy

import (
	"io"

	"github.com/golang/snappy"
	"github.com/trusch/streamstore"
)

// NewWriter returns a new snappy writer
func NewWriter(base io.Writer) (io.WriteCloser, error) {
	snappyWriter := snappy.NewWriter(base)
	return streamstore.NewIOCoppler(snappyWriter, base), nil
}
