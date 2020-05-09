package tfvm

import (
	"fmt"
)

func RunTfvmListCommand(args []string) error {
	inventory, err := GetInventory()
	if err != nil {
		return err
	}

	err = inventory.Update()
	if err != nil {
		return err
	}

	config, err := GetConfiguration()
	if err != nil && !IsNoConfigExists(err) {
		return err
	}

	latestTfRelease := inventory.GetLatestRelease()
	for _, tfRelease := range inventory.GetTerraformReleases() {
		installed, err := inventory.IsTerraformInstalled(tfRelease)
		if err != nil {
			return err
		}

		current := " "
		notes := ""
		if config != nil && (config.versionSpec == tfRelease.Version || config.versionSpec == "latest" && latestTfRelease.Version == tfRelease.Version) {
			notes = fmt.Sprintf(" (selected via %s)", config.file)
			current = "*"
		}
		version := tfRelease.Version
		status := ""
		if installed {
			status = "installed"
		}

		Print("%s %s\t\t- %s%s", current, version, status, notes)
	}

	return nil
}
