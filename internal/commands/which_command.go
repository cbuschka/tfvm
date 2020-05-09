package commands

import (
	inventoryPkg "github.com/cbuschka/tfvm/internal/inventory"
	"github.com/cbuschka/tfvm/internal/util"
	"github.com/cbuschka/tfvm/internal/version"
	workspacePkg "github.com/cbuschka/tfvm/internal/workspace"
)

func RunTfvmWhichCommand(args []string) error {

	inventory, err := inventoryPkg.GetInventory()
	if err != nil {
		return err
	}

	err = inventory.Update()
	if err != nil {
		return err
	}

	workspace, err := workspacePkg.GetWorkspace()
	if err != nil {
		return err
	}

	tfVersionSelection, err := workspace.GetTerraformVersionSelection()
	if err != nil {
		if workspacePkg.IsNoTfVersionSelected(err) {
			util.Print("No terraform version configured.")
			return nil
		}

		return err
	}

	tfRelease, err := inventory.GetTerraformRelease(tfVersionSelection.VersionSpec())
	if err != nil {
		if version.IsNoSuchTerraformRelease(err) {
			util.Die(1, "Configured terraform version %s (%s) is not known.", tfVersionSelection.VersionSpec().String(), tfVersionSelection.Source())
			return err
		}

		return err
	}

	util.Print("Configured terraform version is %s (%s).", tfRelease.Version, tfVersionSelection.Source())
	return nil
}
