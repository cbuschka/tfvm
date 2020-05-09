package tfvm

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/hashicorp/go-version"
	"os"
	"path/filepath"
	"strings"
)

// A .tfvmrc configuration
type Configuration struct {
	versionSpec string
	file        string
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

	versionSpec := ""
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if !strings.HasPrefix(line, "#") && line != "" {
			versionSpec = line
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	if versionSpec != "latest" {
		_, err = version.NewConstraint(versionSpec)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("Invalid version spec: '%s' (%s)", versionSpec, err.Error()))
		}
	}

	return &Configuration{versionSpec: versionSpec, file: configFile}, nil
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
