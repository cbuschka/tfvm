package tfvm

import (
	"fmt"
)

func RunTfvmWhichCommand(args []string) error {

	inventory, err := GetInventory()
	if err != nil {
		return nil
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
		return err
	}

	fmt.Printf("Configured terraform is %s (%s).\n", tfRelease.Version, config.file)
	return nil
}
