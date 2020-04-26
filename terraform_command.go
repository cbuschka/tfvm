package tfvm

import (
	"errors"
	"fmt"
	"os"
)

func RunTerraformCommand(args []string) error {

	config, err := GetConfiguration()
	if err != nil {
		return err
	}

	if config == nil {
		return errors.New("no terraform version configured")
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
