package tfvm

import (
	"fmt"
	"os"
)

func RunTfvmWhichCommand(args []string) error {

	inventory, err := GetInventory()
	if err != nil {
		return err
	}

	err = inventory.Update()
	if err != nil {
		return err
	}

	config, err := GetConfiguration()
	if err != nil {
		if IsNoConfigExists(err) {
			fmt.Printf("No terraform version configured.\n")
			return nil
		}

		return err
	}

	tfRelease, err := inventory.GetTerraformRelease(config.version)
	if err != nil {
		if IsNoSuchTerraformRelease(err) {
			fmt.Printf("Configured terraform version %s (%s) is not known.\n", config.version, config.file)
			os.Exit(1)
		}

		return err
	}

	fmt.Printf("Configured terraform version is %s (%s).\n", tfRelease.Version, config.file)
	return nil
}
