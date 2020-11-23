package commands

import (
	"errors"
	inventoryPkg "github.com/cbuschka/tfvm/internal/inventory"
	"github.com/cbuschka/tfvm/internal/util"
	"github.com/cbuschka/tfvm/internal/version"
	workspacePkg "github.com/cbuschka/tfvm/internal/workspace"
)

// RunTfvmUseCommand runs the tfvm use command.
func RunTfvmUseCommand(args []string) error {
	inventory, err := inventoryPkg.GetInventory()
	if err != nil {
		return err
	}

	err = inventory.Update()
	if err != nil {
		return err
	}

	versionSpec, err := getTfVersionSpecToUse(args)
	if err != nil {
		return err
	}

	_, err = inventory.GetTerraformRelease(versionSpec)
	if err != nil {
		if version.IsNoSuchTerraformRelease(err) {
			util.Die(1, "No matching terraform version for %s.", versionSpec.String())
			return err
		}
		return err
	}

	workspace, err := workspacePkg.GetWorkspace()
	if err != nil {
		return err
	}

	err = workspace.WriteTerraformVersionSelection(versionSpec.String())
	if err != nil {
		if workspacePkg.IsVersionSelectionNotFound(err) {
			util.Die(1, "No .terraform-version found.")
		}
		return err
	}

	return nil
}

func getTfVersionSpecToUse(args []string) (*version.TerraformVersionSpec, error) {
	if len(args) != 1 {
		util.Die(1, "Version to use required.")
		return nil, errors.New("unreachable code")
	}

	versionSpec, err := version.ParseTerraformVersionSpec(args[0])
	if err != nil {
		util.Die(1, "Invalid version '%s'.", args[0])
		return nil, errors.New("unreachable code")
	}

	return versionSpec, nil
}
