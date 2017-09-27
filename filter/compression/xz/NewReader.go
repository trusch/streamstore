package xz

import (
	"io"

	"github.com/trusch/streamstore"
	"github.com/ulikunitz/xz"
)

// NewReader returns a new xz reader
func NewReader(base io.Reader) (io.ReadCloser, error) {
	xzReader, err := xz.NewReader(base)
	if err != nil {
		return nil, err
	}
	return storage.NewIOCoppler(xzReader, base), nil
}
