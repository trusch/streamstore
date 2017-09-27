package gzip

import (
	"compress/gzip"
	"io"

	"github.com/trusch/streamstore"
)

// NewWriter returns a new gzip Writer
// close will be propagated if available
func NewWriter(base io.Writer, level int) (io.WriteCloser, error) {
	w, err := gzip.NewWriterLevel(base, level)
	if err != nil {
		return nil, err
	}
	return streamstore.NewIOCoppler(w, base), nil
}
