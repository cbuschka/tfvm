package tfvm

import (
	"os"
	"path"
	"path/filepath"
)

func GetNearestDotTfvmRcFileFromCwd() (string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	return GetNearestDotTfvmRcFile(cwd), nil
}

func GetNearestDotTfvmRcFile(workingDir string) string {
	currentDir := workingDir
	for currentDir != "/" {
		currentRcFile := path.Join(currentDir, ".tfvmrc")
		if _, err := os.Stat(currentRcFile); err == nil {
			return currentRcFile
		}
		currentDir = filepath.Dir(currentDir)
	}
	return ""
}
