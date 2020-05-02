package tfvm

import (
	"bufio"
	"errors"
	"os"
	"path/filepath"
	"strings"
)

// A .tfvmrc configuration
type Configuration struct {
	version string
	file    string
}

const noConfigExistsMsg = "config not exists"

func IsNoConfigExists(err error) bool {
	return err.Error() == noConfigExistsMsg
}

func newNoConfigExists() error {
	return errors.New(noConfigExistsMsg)
}

var configFileNames = []string{".tfvmrc", ".terraform-version"}

// Get a terraform version by walking through directory structure up to the root
// and looking for .tfvmrc files.
func GetConfiguration() (*Configuration, error) {
	configFile, err := getNearestConfigFileFromCwd()
	if err != nil {
		if os.IsNotExist(err) {
			return nil, newNoConfigExists()
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

	return getNearestConfigFile(cwd)
}

func getNearestConfigFile(workingDir string) (string, error) {
	currentDir := workingDir
	previousDir := ""
	for currentDir != previousDir {
		for _, currentConfigFileName := range configFileNames {
			currentConfigFile := filepath.Join(currentDir, currentConfigFileName)
			if _, err := os.Stat(currentConfigFile); err == nil {
				return currentConfigFile, nil
			}
		}
		previousDir = currentDir
		currentDir = filepath.Dir(currentDir)
	}
	return "", newNoConfigExists()
}
