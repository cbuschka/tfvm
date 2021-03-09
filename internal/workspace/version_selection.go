package workspace

import (
	"errors"
	"os"
	"path/filepath"
)

const versionSelectionNotFoundMsg = "version spec not found"

// IsVersionSelectionNotFound answers if an error means that no .terraform-version file could be found.
func IsVersionSelectionNotFound(err error) bool {
	return err.Error() == versionSelectionNotFoundMsg
}

func newVersionSelectionNotFound() error {
	return errors.New(versionSelectionNotFoundMsg)
}

var versionSelectionFileNames = []string{".terraform-version"}

// Get a terraform version by walking through directory structure up to the root
// and looking for selection files.
func getVersionSelection() (*TerraformVersionSelection, error) {
	versionSelectionFile, err := getNearestVersionSelectionFileFromCwd()
	if err != nil {
		if os.IsNotExist(err) {
			return nil, newNoTfVersionSelected()
		}

		return nil, err
	}

	return readVersionSelectionFromFile(versionSelectionFile)
}

func getNearestVersionSelectionFileFromCwd() (string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	return getNearestVersionSelectionFile(cwd)
}

func getNearestVersionSelectionFile(workingDir string) (string, error) {
	currentDir := workingDir
	previousDir := ""
	for currentDir != previousDir {
		for _, currentVersionSelectionFileName := range versionSelectionFileNames {
			currentVersionSelectionFile := filepath.Join(currentDir, currentVersionSelectionFileName)
			if _, err := os.Stat(currentVersionSelectionFile); err == nil {
				return currentVersionSelectionFile, nil
			}
		}
		previousDir = currentDir
		currentDir = filepath.Dir(currentDir)
	}
	return "", newVersionSelectionNotFound()
}
