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
	installed, err := inventory.IsTerraformInstalled(version)
	if err != nil {
		return err
	}

	if installed {
		fmt.Printf("Terraform %s is already installed.\n", version)
		return nil
	}

	err = inventory.InstallTerraform(version)
	if err != nil {
		return err
	}

	return nil
}
