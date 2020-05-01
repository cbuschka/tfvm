package tfvm

import (
	"fmt"
	"os"
)

func RunTerraformCommand(args []string) error {

	config, err := GetConfiguration()
	if err != nil {
		if IsNoConfigExists(err) {
			fmt.Printf("No terraform version configured. Place .tfvmrc or .terraform-version in current or parent dir.\n")
			os.Exit(1)
		}

		return err
	}

	inventory, err := GetInventory()
	if err != nil {
		return err
	}

	installed, err := inventory.IsTerraformInstalled(config.version)
	if err != nil {
		return err
	}

	if !installed {
		err = inventory.InstallTerraform(config.version)
		if err != nil {
			return nil
		}
	}

	tf, err := inventory.GetTerraform(config.version)
	if err != nil {
		return err
	}

	exitCode, err := tf.Run(args...)
	if err != nil {
		return fmt.Errorf("running tf failed: %v exitCode=%d", err, exitCode)
	}

	os.Exit(exitCode)

	return nil
}
