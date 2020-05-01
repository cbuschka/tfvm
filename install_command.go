package tfvm

import (
	"fmt"
)

func RunTfvmInstallCommand(args []string) error {
	inventory, err := GetInventory()
	if err != nil {
		return err
	}

	version := args[0]

	tfRelease, err := inventory.GetTerraformRelease(version)
	if err != nil {
		return nil
	}

	installed, err := inventory.IsTerraformInstalled(tfRelease)
	if err != nil {
		return err
	}

	if installed {
		fmt.Printf("Terraform %s is already installed.\n", version)
		return nil
	}

	err = inventory.InstallTerraform(tfRelease)
	if err != nil {
		return err
	}

	return nil
}
