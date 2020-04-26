package tfvm

import (
	"bufio"
	"os"
	"path"
	"path/filepath"
	"strings"
)

// A .tfvmrc configuration
type Configuration struct {
	version string
	file    string
}

var configFileNames = []string{".tfvmrc", ".terraform-version"}

// Get a terraform version by walking through directory structure up to the root
// and looking for .tfvmrc files.
func GetConfiguration() (*Configuration, error) {
	configFile, err := getNearestConfigFileFromCwd()
	if err != nil {
		if os.IsNotExist(err) {
			return nil, nil
		}

		return nil, err
	}

	file, err := os.Open(configFile)
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

	return &Configuration{version: tfVersion, file: configFile}, nil
}

func getNearestConfigFileFromCwd() (string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	return getNearestConfigFile(cwd), nil
}

func getNearestConfigFile(workingDir string) string {
	currentDir := workingDir
	for currentDir != "/" {
		for _, currentConfigFileName := range configFileNames {
			currentConfigFile := path.Join(currentDir, currentConfigFileName)
			if _, err := os.Stat(currentConfigFile); err == nil {
				return currentConfigFile
			}
		}
		currentDir = filepath.Dir(currentDir)
	}
	return ""
}
