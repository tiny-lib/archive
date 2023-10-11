// Package tarxz implements the Archive interface providing tar.xz archiving
// and compression.
package tarxz

import (
	"github.com/tiny-lib/archive/pkg/archiveFile"
	"github.com/tiny-lib/archive/pkg/tar"
	"io"

	"github.com/ulikunitz/xz"
)

// Archive as tar.xz.
type Archive struct {
	xzw *xz.Writer
	tw  *tar.Archive
}

// New tar.xz archive.
func New(target io.Writer) Archive {
	xzw, _ := xz.WriterConfig{DictCap: 16 * 1024 * 1024}.NewWriter(target)
	tw := tar.New(xzw)
	return Archive{
		xzw: xzw,
		tw:  &tw,
	}
}

// Close all closeables.
func (a Archive) Close() error {
	if err := a.tw.Close(); err != nil {
		return err
	}
	return a.xzw.Close()
}

// Add file to the archive.
func (a Archive) Add(f archiveFile.Entry) error {
	return a.tw.Add(f)
}