package workspace

import (
	"github.com/cbuschka/tfvm/internal/log"
	"github.com/cbuschka/tfvm/internal/util"
	"os"
	"path/filepath"
)

func getRootDir() (string, error) {

	log.Trace("Finding workplace root dir...")

	rootDir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	log.Tracef("CWD: %s", rootDir)

	log.Tracef("Walking up the tree, starting with %s...", rootDir)
	err = walkUp(rootDir, func(path string) (bool, error) {
		dotGitDir := filepath.Join(path, ".git")
		isDir, err := util.IsDir(dotGitDir)
		if err != nil {
			return false, err
		}

		log.Tracef("Checked if git dir exists: %s => %t", dotGitDir, isDir)

		if isDir {
			log.Trace("Git dir found.")

			rootDir = path
			return false, nil
		}

		return true, nil
	})
	if err != nil {
		return "", err
	}

	log.Debugf("Root dir: %s", rootDir)

	return rootDir, err
}
