package commands

import (
	inventoryPkg "github.com/cbuschka/tfvm/internal/inventory"
	"github.com/cbuschka/tfvm/internal/util"
	"github.com/cbuschka/tfvm/internal/version"
	workspacePkg "github.com/cbuschka/tfvm/internal/workspace"
	"os"
)

// RunTerraformCommand runs tfvm terraform, resp. terraform command.
func RunTerraformCommand(args []string) error {

	workspace, err := workspacePkg.GetWorkspace()
	if err != nil {
		return err
	}

	if !util.IsOutputTerminal() {
		util.SuppressOutput()
	}

	tfVersionSelection, err := workspace.GetTerraformVersionSelection()
	if err != nil {
		if workspacePkg.IsNoTfVersionSelected(err) {
			util.Die(1, "No terraform version configured. Place .terraform-version in current dir or above. (Or set env var TERRAFORM_VERSION or TFVM_TERRAFORM_VERSION.)")
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

	err = inventory.Save()
	if err != nil {
		return err
	}

	tfRelease, err := inventory.GetMatchingTerraformRelease(tfVersionSelection.VersionSpec())
	if err != nil {
		if version.IsNoSuchTerraformRelease(err) {
			util.Die(1, "Terraform version %s is not known.", tfVersionSelection.VersionSpec().String())
			return err
		}

		return err
	}

	tf, err := inventory.GetInstalledTerraform(tfRelease.Version)
	if err != nil {
		util.Die(1, "Running terraform failed: %s.", err.Error())
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
