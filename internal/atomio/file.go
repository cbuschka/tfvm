package atomio

import (
	"crypto/rand"
	"encoding/hex"
	"github.com/cbuschka/tfvm/internal/log"
	atomicFile "github.com/natefinch/atomic"
	"os"
)

func getTempSuffix() (string, error) {
	bytes := make([]byte, 16)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

// OpenTempFile opens a file named xxx as xxx.randomsuffix with the given flags and file mode.
func OpenTempFile(name string, flag int, perm os.FileMode) (*os.File, error) {
	tempSuffix, err := getTempSuffix()
	if err != nil {
		return nil, err
	}
	tempName := name + "." + tempSuffix

	log.Tracef("Opening temp file %s instead of %s.", tempName, name)

	file, err := os.OpenFile(tempName, flag, perm)
	if err != nil {
		return nil, err
	}

	return file, nil
}

// CloseAndReplaceFile closes file and moves it atomically to newPath.
func CloseAndReplaceFile(file *os.File, newPath string) error {
	err := file.Close()
	if err != nil {
		return err
	}

	currPath := file.Name()
	if currPath != newPath {
		log.Tracef("Replacing %s with file %s.", currPath, newPath)
		err = atomicFile.ReplaceFile(currPath, newPath)
		if err != nil {
			return err
		}
	}

	return nil
}
