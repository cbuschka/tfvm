package workspace

import (
	"os"
	"path/filepath"
)

// WalkUpFunc is the callback called for walking up the filesystem. Thre traversal is aborted if it returns false or an error.
type walkUpFunc func(path string) (bool, error)

// WalkUp walks up the filesystem starting with path. If path is a file and not a directory the traversal
// starts at the parent of path. For every dir found callback is called.
func walkUp(path string, callback walkUpFunc) error {

	stat, err := os.Stat(path)
	if err != nil {
		return err
	}

	if !stat.IsDir() {
		path = filepath.Dir(path)
		stat, err = os.Stat(path)
		if err != nil {
			return err
		}
	}

	for {
		cont, err := callback(path)
		if err != nil {
			return err
		}
		if !cont {
			break
		}

		newpath := filepath.Dir(path)
		if newpath == path {
			break
		}

		path = newpath
	}

	return nil
}
