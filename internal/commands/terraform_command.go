package commands

import (
	inventoryPkg "github.com/cbuschka/tfvm/internal/inventory"
	"github.com/cbuschka/tfvm/internal/util"
	"github.com/cbuschka/tfvm/internal/version"
	workspacePkg "github.com/cbuschka/tfvm/internal/workspace"
	"os"
)

func RunTerraformCommand(args []string) error {

	workspace, err := workspacePkg.GetWorkspace()
	if err != nil {
		return err
	}

	tfVersionSelection, err := workspace.GetTerraformVersionSelection()
	if err != nil {
		if workspacePkg.IsNoTfVersionSelected(err) {
			util.Die(1, "No terraform version configured. Place .terraform-version in current or parent dir.")
			return err
		}

		return err
	}

	inventory, err := inventoryPkg.GetInventory()
	if err != nil {
		return err
	}

	err = inventory.Update()
	if err != nil {
		return err
	}

	tfRelease, err := inventory.GetTerraformRelease(tfVersionSelection.VersionSpec())
	if err != nil {
		if version.IsNoSuchTerraformRelease(err) {
			util.Die(1, "Terraform version %s is not known.", tfVersionSelection.VersionSpec().String())
			return err
		}

		return err
	}

	installed, err := inventory.IsTerraformInstalled(tfRelease)
	if err != nil {
		return err
	}

	if !installed {
		err = inventory.InstallTerraform(tfRelease)
		if err != nil {
			return err
		}
	}

	tf, err := inventory.GetTerraform(tfRelease)
	if err != nil {
		return err
	}

	exitCode, err := tf.Run(args...)
	if err != nil {
		util.Die(exitCode, "Running terraform failed: %s, exitCode=%d.", err.Error(), exitCode)
		return err
	}

	os.Exit(exitCode)

	return nil
}
