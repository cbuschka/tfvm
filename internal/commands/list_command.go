package commands

import (
	"fmt"
	inventoryPkg "github.com/cbuschka/tfvm/internal/inventory"
	"github.com/cbuschka/tfvm/internal/platform"
	"github.com/cbuschka/tfvm/internal/util"
	workspacePkg "github.com/cbuschka/tfvm/internal/workspace"
)

// RunTfvmListCommand runs tfvm list command.
func RunTfvmListCommand(args []string) error {
	inventory, err := inventoryPkg.GetInventory()
	if err != nil {
		return err
	}

	err = inventory.Update()
	if err != nil {
		return err
	}

	err = inventory.Save()
	if err != nil {
		return err
	}

	workspace, err := workspacePkg.GetWorkspace()
	if err != nil {
		return err
	}

	tfVersionSelection, err := workspace.GetTerraformVersionSelection()
	if err != nil && !workspacePkg.IsNoTfVersionSelected(err) {
		return err
	}

	latestTfRelease, err := inventory.GetLatestRelease()
	if err != nil {
		return err
	}

	primaryPlatform := platform.GetPrimaryPlatform()

	for _, tfRelease := range inventory.GetTerraformReleasesAsc() {
		installed, err := inventory.IsTerraformInstalled(tfRelease, primaryPlatform.Os, primaryPlatform.Arch)
		if err != nil {
			return err
		}

		current := " "
		notes := ""
		if tfVersionSelection != nil {
			isSelected := tfVersionSelection.VersionSpec().Matches(tfRelease, latestTfRelease)
			if isSelected {
				notes = fmt.Sprintf(" (selected via %s)", tfVersionSelection.Source())
				current = "*"
			}
		}

		version := tfRelease.String()
		status := ""
		if installed {
			status = "installed"
		}

		util.Print("%s %s\t\t- %s%s", current, version, status, notes)
	}

	return nil
}
