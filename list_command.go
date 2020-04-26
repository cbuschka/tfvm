package tfvm

import (
	"fmt"
)

func RunTfvmListCommand(args []string) error {
	inventory, err := GetInventory()
	if err != nil {
		return err
	}

	tfReleases, err := ListTerraformReleases()
	if err != nil {
		return err
	}

	config, err := GetConfiguration()
	if err != nil {
		return err
	}

	for _, tfRelease := range tfReleases {
		installed, err := inventory.IsTerraformInstalled(tfRelease.version)
		if err != nil {
			return err
		}

		current := " "
		notes := ""
		if config != nil && config.version == tfRelease.version {
			notes = fmt.Sprintf(" (selected via %s)", config.file)
			current = "*"
		}
		version := tfRelease.version
		status := ""
		if installed {
			status = "installed"
		}

		fmt.Printf("%s %s\t\t- %s%s\n", current, version, status, notes)
	}

	return nil
}
