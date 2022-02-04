package util

import (
	"os"
)

// IsFile checks if there is an ordinary file at path.
// Returns true if path is a file.
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
