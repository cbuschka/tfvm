package tfvm

func RunTfvmUptodateCommand(args []string) error {

	updateStatus := GetUpdateStatus()
	if updateStatus.updateAvailable {
		Print("This is tfvm %s. There is a newer tfvm release (%s) available.", updateStatus.currentVersion, updateStatus.latestVersion)
	} else {
		Print("This tfvm (%s) is up to date.", updateStatus.currentVersion)
	}

	return nil
}
