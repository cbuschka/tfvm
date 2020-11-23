package util

import (
	"os"
)

// IsDir checks if there is a directory at path.
// Returns true if path is a directory, false if it is a normal file or does not exist.
func IsDir(path string) (bool, error) {

	stat, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false, nil
	}

	if err != nil {
		return false, err
	}

	return stat.IsDir(), nil
}
