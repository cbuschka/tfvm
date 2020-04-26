package tfvm

import (
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func GetConfiguredVersion() (string, error) {
	dotTfvmRcFile, err := GetNearestDotTfvmRcFileFromCwd()
	if err != nil {
		if os.IsNotExist(err) {
			return "", nil
		}

		return "", err
	}

	tfVersionBytes, err := ioutil.ReadFile(dotTfvmRcFile)
	if err != nil {
		return "", err
	}
	tfVersion := string(tfVersionBytes)
	tfVersion = strings.TrimSpace(tfVersion)
	return tfVersion, nil
}

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
