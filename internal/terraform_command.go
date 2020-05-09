package tfvm

import (
	"os"
)

func RunTerraformCommand(args []string) error {

	config, err := GetConfiguration()
	if err != nil {
		if IsNoConfigExists(err) {
			Die(1, "No terraform version configured. Place .tfvmrc or .terraform-version in current or parent dir.")
			return err
		}

		return err
	}

	inventory, err := GetInventory()
	if err != nil {
		return err
	}

	err = inventory.Update()
	if err != nil {
		return err
	}

	tfRelease, err := inventory.GetTerraformRelease(config.versionSpec)
	if err != nil {
		if IsNoSuchTerraformRelease(err) {
			Die(1, "Terraform version %s is not known.", config.versionSpec.String())
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
		Die(exitCode, "Running terraform failed: %s, exitCode=%d.", err.Error(), exitCode)
		return err
	}

	os.Exit(exitCode)

	return nil
}
