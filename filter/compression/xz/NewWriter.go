package xz

import (
	"io"

	"github.com/trusch/streamstore"
	"github.com/ulikunitz/xz"
)

// NewWriter returns a new xz writer
func NewWriter(base io.Writer) (io.WriteCloser, error) {
	writer, err := xz.NewWriter(base)
	if err != nil {
		return nil, err
	}
	return streamstore.NewIOCoppler(writer, base), nil
}
