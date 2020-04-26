package tfvm

import (
	"bufio"
	"os"
	"path"
	"path/filepath"
	"strings"
)

type Configuration struct {
	// A .tfvmrc configuration
	version string
	file string
}

func GetConfiguration() (*Configuration, error) {
	// Get a terraform version by walking through directory structure up to the root
	// and looking for .tfvmrc files.
	dotTfvmRcFile, err := getNearestDotTfvmRcFileFromCwd()
	if err != nil {
		if os.IsNotExist(err) {
			return nil, nil
		}

		return nil, err
	}

	file, err := os.Open(dotTfvmRcFile)
	if err != nil {
		return nil, err
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
		return nil, err
	}

	return &Configuration{version: tfVersion, file: dotTfvmRcFile}, nil
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
