package workspace

import (
	"errors"
	"fmt"
	"github.com/cbuschka/tfvm/internal/util"
	"github.com/cbuschka/tfvm/internal/version"
)

type Workspace struct {
}

func GetWorkspace() (*Workspace, error) {
	return &Workspace{}, nil
}

type SelectionSourceType int

const (
	File SelectionSourceType = 0
	Env  SelectionSourceType = 1
)

// A .tfvmrc configuration
type TerraformVersionSelection struct {
	versionSpec *version.TerraformVersionSpec
	sourceName  string
	sourceType  SelectionSourceType
}

func (config *TerraformVersionSelection) VersionSpec() *version.TerraformVersionSpec {
	return config.versionSpec
}

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

func IsNoTfVersionSelected(err error) bool {
	return err.Error() == noTfVersionSelectedMsg
}

func newNoTfVersionSelected() error {
	return errors.New(noTfVersionSelectedMsg)
}

// Get a terraform version by walking through directory structure up to the root
// and looking for .tfvmrc files.
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
		if isNoConfigExists(err) {
			return nil, newNoTfVersionSelected()
		}
		return nil, err
	}

	return config, nil
}
