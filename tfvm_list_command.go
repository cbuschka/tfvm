package tfvm

import (
	"fmt"
)

func RunTfvmListCommand(args []string) error {
	tfReleases, err := ListTerraformReleases()
	if err != nil {
		return err
	}

	for _, tfRelease := range tfReleases {
		fmt.Printf("%s\n", tfRelease.version)
	}

	return nil
}
