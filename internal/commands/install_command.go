package commands

import (
	"errors"
	inventoryPkg "github.com/cbuschka/tfvm/internal/inventory"
	"github.com/cbuschka/tfvm/internal/util"
	"github.com/cbuschka/tfvm/internal/version"
	workspacePkg "github.com/cbuschka/tfvm/internal/workspace"
)

// RunTfvmInstallCommand runs tfvm install command.
func RunTfvmInstallCommand(args []string) error {
	inventory, err := inventoryPkg.GetInventory()
	if err != nil {
		return err
	}

	err = inventory.Update()
	if err != nil {
		return err
	}

	versionSpec, err := getTfVersionSpecToInstall(args)
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

func getTfVersionSpecToInstall(args []string) (*version.TerraformVersionSpec, error) {
	if len(args) > 1 {
		util.Die(1, "Only version to install allowed.")
		return nil, errors.New("unreachable code")
	}

	if len(args) == 1 {
		versionSpec, err := version.ParseTerraformVersionSpec(args[0])
		if err != nil {
			util.Die(1, "Invalid version '%s'.", args[0])
			return nil, errors.New("unreachable code")
		}

		return versionSpec, nil
	}

	workspace, err := workspacePkg.GetWorkspace()
	if err != nil {
		return nil, err
	}

	tfVersionSelection, err := workspace.GetTerraformVersionSelection()
	if err != nil {
		return nil, err
	}

	return tfVersionSelection.VersionSpec(), nil
}
