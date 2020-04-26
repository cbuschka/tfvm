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

		notes := ""
		if config != nil && config.version == tfRelease.version {
			notes = fmt.Sprintf(" (selected, configured in %s)", config.file)
		}
		version := tfRelease.version
		status := ""
		if installed {
			status = "installed"
		}

		fmt.Printf("%s - %s%s\n", version, status, notes)
	}

	return nil
}
