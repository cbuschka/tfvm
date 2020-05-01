package tfvm

import (
	"fmt"
	"os"
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
		fmt.Printf("Expected version to install.\n")
		os.Exit(1)
	}
	version := args[0]

	tfRelease, err := inventory.GetTerraformRelease(version)
	if err != nil {
		if IsNoSuchTerraformRelease(err) {
			fmt.Printf("Terraform version %s is not known.\n", version)
			os.Exit(1)
		}
		return err
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
