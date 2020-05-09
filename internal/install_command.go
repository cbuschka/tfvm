package tfvm

import (
	"errors"
)

func RunTfvmInstallCommand(args []string) error {
	inventory, err := GetInventory()
	if err != nil {
		return err
	}

	err = inventory.Update()
	if err != nil {
		return err
	}

	if len(args) < 1 {
		Die(1, "Expected version to install.")
		return errors.New("unreachable code")
	}
	versionSpec, err := ParseTerraformVersionSpec(args[0])
	if err != nil {
		return err
	}

	tfRelease, err := inventory.GetTerraformRelease(versionSpec)
	if err != nil {
		if IsNoSuchTerraformRelease(err) {
			Die(1, "No matching terraform version for %s.", versionSpec.text)
			return err
		}
		return err
	}

	installed, err := inventory.IsTerraformInstalled(tfRelease)
	if err != nil {
		return err
	}

	if installed {
		Print("Terraform %s is already installed.", tfRelease.Version.String())
		return nil
	}

	err = inventory.InstallTerraform(tfRelease)
	if err != nil {
		return err
	}

	return nil
}
