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
	version := args[0]

	tfRelease, err := inventory.GetTerraformRelease(version)
	if err != nil {
		if IsNoSuchTerraformRelease(err) {
			Die(1, "Terraform version %s is not known.", version)
			return err
		}
		return err
	}

	installed, err := inventory.IsTerraformInstalled(tfRelease)
	if err != nil {
		return err
	}

	if installed {
		Print("Terraform %s is already installed.", version)
		return nil
	}

	err = inventory.InstallTerraform(tfRelease)
	if err != nil {
		return err
	}

	return nil
}
