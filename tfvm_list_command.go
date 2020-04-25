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

	for _, tfRelease := range tfReleases {
		installed, err := inventory.IsTerraformInstalled(tfRelease.version)
		if err != nil {
			return err
		}

		if installed {
			fmt.Printf("%s (installed)\n", tfRelease.version)
		} else {
			fmt.Printf("%s\n", tfRelease.version)
		}
	}

	return nil
}
