package commands

import (
	"github.com/cbuschka/tfvm/internal/build"
	"github.com/cbuschka/tfvm/internal/util"
)

// RunTfvmUptodateCommand runs tfvm uptodate check.
func RunTfvmUptodateCommand(args []string) error {

	updateStatus := build.GetUpdateStatus()
	if updateStatus.UpdateAvailable() {
		util.Print("This is tfvm %s. There is a newer tfvm release (%s) available.", updateStatus.CurrentVersion(), updateStatus.LatestVersion())
	} else {
		util.Print("This tfvm (%s) is up to date.", updateStatus.CurrentVersion())
	}

	return nil
}
