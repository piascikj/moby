package driver

import (
	"os"

	"github.com/pkg/errors"
)

func (d *driver) Mknod(path string, mode os.FileMode, major, minor int) error {
	return errors.Wrap(ErrNotSupported, "cannot create device node on Windows")
}

func (d *driver) Mkfifo(path string, mode os.FileMode) error {
	return errors.Wrap(ErrNotSupported, "cannot create fifo on Windows")
}

// Lchmod changes the mode of an file not following symlinks.
func (d *driver) Lchmod(path string, mode os.FileMode) (err error) {
	// TODO: Use Window's equivalent id:279 gh:280
	return os.Chmod(path, mode)
}
