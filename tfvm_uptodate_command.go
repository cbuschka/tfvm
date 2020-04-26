package tfvm

import (
	"fmt"
)

func RunTfvmUptodateCommand(args []string) error {

	updateStatus := GetUpdateStatus()
	if updateStatus.updateAvailable {
		fmt.Printf("This is tfvm %s. There is a newer tfvm release (%s) available.\n", updateStatus.currentVersion, updateStatus.latestVersion)
	} else {
		fmt.Printf("This tfvm (%s) is up to date.\n", updateStatus.currentVersion)
	}

	return nil
}
