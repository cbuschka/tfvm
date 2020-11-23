package util

import (
	"os"
)

// IsFile checks if there is a file at path.
// Returns true if path is a file, false if it is a directory or does not exist.
func IsFile(path string) (bool, error) {

	stat, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false, nil
	}

	if err != nil {
		return false, err
	}

	return !stat.IsDir(), nil
}
