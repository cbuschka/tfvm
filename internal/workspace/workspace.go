package workspace

import (
	"errors"
	"github.com/cbuschka/tfvm/internal/version"
)

type Workspace struct {
}

func GetWorkspace() (*Workspace, error) {
	return &Workspace{}, nil
}

// A .tfvmrc configuration
type TerraformVersionSelection struct {
	versionSpec *version.TerraformVersionSpec
	file        string
}

func (config *TerraformVersionSelection) VersionSpec() *version.TerraformVersionSpec {
	return config.versionSpec
}

func (config *TerraformVersionSelection) File() string {
	return config.file
}

const noTfVersionSelectedMsg = "no terraform version selected"

func IsNoTfVersionSelected(err error) bool {
	return err.Error() == noTfVersionSelectedMsg
}

func newNoTfVersionSelected() error {
	return errors.New(noTfVersionSelectedMsg)
}

// Get a terraform version by walking through directory structure up to the root
// and looking for .tfvmrc files.
func (workspace *Workspace) GetTerraformVersionSelection() (*TerraformVersionSelection, error) {

	config, err := getConfiguration()
	if err != nil {
		if isNoConfigExists(err) {
			return nil, newNoTfVersionSelected()
		}
		return nil, err
	}

	return config, nil
}
