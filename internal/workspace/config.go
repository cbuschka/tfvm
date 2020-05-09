package workspace

import (
	"bufio"
	"errors"
	"github.com/cbuschka/tfvm/internal/version"
	"os"
	"path/filepath"
	"strings"
)

const noConfigExistsMsg = "config not exists"

func isNoConfigExists(err error) bool {
	return err.Error() == noConfigExistsMsg
}

func newNoConfigExists() error {
	return errors.New(noConfigExistsMsg)
}

var configFileNames = []string{".tfvmrc", ".terraform-version"}

// Get a terraform version by walking through directory structure up to the root
// and looking for .tfvmrc files.
func getConfiguration() (*TerraformVersionSelection, error) {
	configFile, err := getNearestConfigFileFromCwd()
	if err != nil {
		if os.IsNotExist(err) {
			return nil, newNoTfVersionSelected()
		}

		return nil, err
	}

	return readConfiguration(configFile)
}

func readConfiguration(configFile string) (*TerraformVersionSelection, error) {
	file, err := os.Open(configFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	versionSpecStr := ""
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if !strings.HasPrefix(line, "#") && line != "" {
			versionSpecStr = line
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	versionSpec, err := version.ParseTerraformVersionSpec(versionSpecStr)
	if err != nil {
		return nil, err
	}

	return &TerraformVersionSelection{versionSpec: versionSpec, file: configFile}, nil
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
