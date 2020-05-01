package tfvm

import (
	"fmt"
	"time"
)

func RunTfvmListCommand(args []string) error {
	inventory, err := GetInventory()
	if err != nil {
		return err
	}

	tfReleases, err := ListTerraformReleases()
	if err != nil && len(inventory.TerraformReleases) == 0 {
		return err
	}

	if tfReleases != nil {
		inventory.TerraformReleases = tfReleases
		inventory.LastUpdateTime = time.Now()
		err = inventory.Save()
		if err != nil {
			return err
		}
	} else {
		tfReleases = inventory.TerraformReleases
	}

	config, err := GetConfiguration()
	if err != nil && !IsNoConfigExists(err) {
		return err
	}

	for _, tfRelease := range tfReleases {
		installed, err := inventory.IsTerraformInstalled(tfRelease.Version)
		if err != nil {
			return err
		}

		current := " "
		notes := ""
		if config != nil && config.version == tfRelease.Version {
			notes = fmt.Sprintf(" (selected via %s)", config.file)
			current = "*"
		}
		version := tfRelease.Version
		status := ""
		if installed {
			status = "installed"
		}

		fmt.Printf("%s %s\t\t- %s%s\n", current, version, status, notes)
	}

	return nil
}
