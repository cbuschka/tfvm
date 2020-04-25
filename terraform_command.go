package tfvm

import (
	"errors"
	"fmt"
	"os"
)

func RunTerraformCommand(args []string) error {

        tfVersion, err := GetConfiguredVersion()
        if err != nil {
                return err
        }

	if tfVersion == "" {
		return errors.New("No terraform version configured.")
	}

	inventory, err := GetInventory()
	if err != nil {
		return err
	}

	installed, err := inventory.IsTerraformInstalled(tfVersion)
	if err != nil {
		return err
	}

	if !installed {
		err = inventory.InstallTerraform(tfVersion)
		if err != nil {
			return nil
		}
	}

	tf, err := inventory.GetTerraform(tfVersion)
	if err != nil {
		return err
	}

	exitCode, err := tf.Run(args...)
	if err != nil {
		return fmt.Errorf("Running tf failed: %v exitCode=%d", err, exitCode)
	}

	os.Exit(exitCode)

	return nil
}
