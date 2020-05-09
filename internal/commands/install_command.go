package commands

import (
	"errors"
	inventoryPkg "github.com/cbuschka/tfvm/internal/inventory"
	"github.com/cbuschka/tfvm/internal/util"
	"github.com/cbuschka/tfvm/internal/version"
)

func RunTfvmInstallCommand(args []string) error {
	inventory, err := inventoryPkg.GetInventory()
	if err != nil {
		return err
	}

	err = inventory.Update()
	if err != nil {
		return err
	}

	if len(args) < 1 {
		util.Die(1, "Expected version to install.")
		return errors.New("unreachable code")
	}
	
	versionSpec, err := version.ParseTerraformVersionSpec(args[0])
	if err != nil {
		return err
	}

	tfRelease, err := inventory.GetTerraformRelease(versionSpec)
	if err != nil {
		if version.IsNoSuchTerraformRelease(err) {
			util.Die(1, "No matching terraform version for %s.", versionSpec.String())
			return err
		}
		return err
	}

	installed, err := inventory.IsTerraformInstalled(tfRelease)
	if err != nil {
		return err
	}

	if installed {
		util.Print("Terraform %s is already installed.", tfRelease.Version.String())
		return nil
	}

	err = inventory.InstallTerraform(tfRelease)
	if err != nil {
		return err
	}

	return nil
}
