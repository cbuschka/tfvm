package tfvm

import (
	"bufio"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func GetConfiguredVersion() (string, error) {
	// Get a terraform version by walking through directory structure up to the root
	// and looking for .tfvmrc files.
	dotTfvmRcFile, err := getNearestDotTfvmRcFileFromCwd()
	if err != nil {
		if os.IsNotExist(err) {
			return "", nil
		}

		return "", err
	}

	file, err := os.Open(dotTfvmRcFile)
	if err != nil {
		return "", err
	}
	defer file.Close()

	tfVersion := ""
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if !strings.HasPrefix(line, "#") && line != "" {
			tfVersion = line
		}
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	return tfVersion, nil
}

func getNearestDotTfvmRcFileFromCwd() (string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	return getNearestDotTfvmRcFile(cwd), nil
}

func getNearestDotTfvmRcFile(workingDir string) string {
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
