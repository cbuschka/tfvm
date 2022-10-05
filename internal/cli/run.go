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

	programName := getProgramNameFrom(args)
	log.Debugf("Program name: '%s'", programName)

	var err error
	if isProgramNameTerraform(programName) {
		err = commands.RunTerraformCommand(args[1:])
	} else if isProgramNameTfenv(programName) {
		err = commands.RunTfenvCommand(args[1:])
	} else {
		err = commands.RunTfvmCommand(args[1:])
	}

	if err != nil {
		util.Die(1, err.Error())
	}

	log.Infof("Exiting with code 0.")
	return nil
}

func getProgramNameFrom(args []string) string {
	return filepath.Base(args[0])
}

func isProgramNameTerraform(programName string) bool {
	return programName == "terraform" || programName == "terraform.exe"
}

func isProgramNameTfenv(programName string) bool {
	return programName == "tfenv"
}
