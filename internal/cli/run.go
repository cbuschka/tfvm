package cli

import (
	"github.com/cbuschka/tfvm/internal/build"
	"github.com/cbuschka/tfvm/internal/commands"
	"github.com/cbuschka/tfvm/internal/log"
	"github.com/cbuschka/tfvm/internal/util"
	"os"
	"path/filepath"
)

// Run tfvm command based on command line args.
func Run(args []string) error {

	log.Debugf("Orig args: %v", os.Args)
	log.Infof("Args: %v", args)
	buildInfo := build.GetBuildInfo()
	log.Infof("Build info: version=%s, commitish=%s, build timestamp=%s, build os=%s, build arch=%s", buildInfo.Version, buildInfo.Commitish, buildInfo.BuildTime, buildInfo.OS, buildInfo.Arch)

	var err error
	if isProgramNameTerraform(args) {
		err = commands.RunTerraformCommand(args[1:])
	} else {
		err = commands.RunTfvmCommand(args[1:])
	}

	if err != nil {
		util.Die(1, err.Error())
	}

	log.Infof("Exiting with code 0.")
	return nil
}

func isProgramNameTerraform(args []string) bool {
	if len(args) == 0 {
		return false
	}

	programName := filepath.Base(args[0])
	log.Debugf("Program name: '%s'", programName)

	return programName == "terraform" || programName == "terraform.exe"
}
