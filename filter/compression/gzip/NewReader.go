package gzip

import (
	"compress/gzip"
	"io"

	"github.com/trusch/streamstore"
)

// NewReader returns a new gzip reader
// close will be propagated if available
func NewReader(base io.Reader) (io.ReadCloser, error) {
	gzipReader, err := gzip.NewReader(base)
	if err != nil {
		return nil, err
	}
	return streamstore.NewIOCoppler(gzipReader, base), nil
}
