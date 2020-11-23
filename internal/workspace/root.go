package workspace

import (
	"github.com/cbuschka/tfvm/internal/util"
	"os"
	"path/filepath"
)

func getRootDir() (string, error) {

	cwd, err := os.Getwd()
	rootDir := cwd
	err = util.WalkUp(cwd, func(path string, info os.FileInfo) (bool, error) {
		dotGitDir := filepath.Join(path, ".git")
		isDir, err := util.IsDir(dotGitDir)
		if err != nil {
			return false, err
		}

		if isDir {
			rootDir = path
			return false, nil
		}

		return true, nil
	})
	if err != nil {
		return "", err
	}

	return rootDir, err
}
