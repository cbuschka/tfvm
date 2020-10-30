package workspace

import (
	"errors"
	"fmt"
	"github.com/cbuschka/tfvm/internal/util"
	"github.com/cbuschka/tfvm/internal/version"
	"io/ioutil"
	"os"
)

// Workspace represents a workspace, that possibly contains a
// .terraform-version file with a terraform version selection.
type Workspace struct {
}

// GetWorkspace returns an initialized workspace instance.
func GetWorkspace() (*Workspace, error) {
	return &Workspace{}, nil
}

// SelectionSourceType describes the method how a TerraformVersionSelection is selected.
type SelectionSourceType int

const (
	// File means the terraform version is selected via a .terraform-version file.
	File SelectionSourceType = 0
	// Env means the terraform version is selected via a env variable.
	Env SelectionSourceType = 1
)

// TerraformVersionSelection describes the result of version selection.
type TerraformVersionSelection struct {
	versionSpec *version.TerraformVersionSpec
	sourceName  string
	sourceType  SelectionSourceType
}

// VersionSpec returns the version specification in the TerraformVersionSelection.
func (config *TerraformVersionSelection) VersionSpec() *version.TerraformVersionSpec {
	return config.versionSpec
}

// Source gives the source of version section.
func (config *TerraformVersionSelection) Source() string {
	if config.sourceType == File {
		return fmt.Sprintf("file %s", config.sourceName)
	} else if config.sourceType == Env {
		return fmt.Sprintf("env var %s", config.sourceName)
	} else {
		panic("selection without sourceType type")
	}
}

const noTfVersionSelectedMsg = "no terraform version selected"

// IsNoTfVersionSelected answer if an error's cause is that no terraform version has been selected.
func IsNoTfVersionSelected(err error) bool {
	return err.Error() == noTfVersionSelectedMsg
}

func newNoTfVersionSelected() error {
	return errors.New(noTfVersionSelectedMsg)
}

// WriteTerraformVersionSelection writes a terraform selection to a .terraform-version file.
func (workspace *Workspace) WriteTerraformVersionSelection(tfVersionSelection string) error {
	configFile, err := getNearestConfigFileFromCwd()
	if err != nil {
		if os.IsNotExist(err) {
			return newNoConfigExists()
		}

		return err
	}

	err = ioutil.WriteFile(configFile, []byte(tfVersionSelection), 0644)
	if err != nil {
		return nil
	}

	util.Print("%s written.", configFile)

	return nil

}

// GetTerraformVersionSelection get the terraform version selection in a workspace.
func (workspace *Workspace) GetTerraformVersionSelection() (*TerraformVersionSelection, error) {

	tfVersionEnvVar := util.GetFirstEnv("TFVM_TERRAFORM_VERSION", "TERRAFORM_VERSION")
	if tfVersionEnvVar != "" {
		versionSpec, err := version.ParseTerraformVersionSpec(tfVersionEnvVar)
		if err != nil {
			util.Die(1, "Invalid version '%s' specified via env var TERRAFORM_VERSION.", tfVersionEnvVar)
			return nil, errors.New("unreachable code")
		}

		return &TerraformVersionSelection{versionSpec: versionSpec, sourceName: "TERRAFORM_VERSION", sourceType: Env}, nil
	}

	config, err := getConfiguration()
	if err != nil {
		if IsNoConfigExists(err) {
			return nil, newNoTfVersionSelected()
		}
		return nil, err
	}

	return config, nil
}
