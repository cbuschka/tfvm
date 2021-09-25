package main

import (
	"github.com/cbuschka/tfvm/internal/commands"
	"github.com/cbuschka/tfvm/internal/log"
	"github.com/cbuschka/tfvm/internal/util"
	"os"
	"path/filepath"
	"runtime"
)

func main() {

	args := os.Args
	verbosity, args := removeVerbosityFlags(args)
	log.SetVerbosity(verbosity)

	log.Debugf("Orig args: %v", os.Args)
	log.Infof("Args: %v", args)

	programName := filepath.Base(args[0])
	log.Debugf("Program name: '%s'", programName)

	var err error
	if programName == "terraform" || (runtime.GOOS == "windows" && programName == "terraform.exe") {
		err = commands.RunTerraformCommand(args[1:])
	} else if programName != "terraform" && len(args) > 1 && args[1] == "terraform" {
		err = commands.RunTerraformCommand(args[2:])
	} else {
		err = commands.RunTfvmCommand(args[1:])
	}

	if err != nil {
		util.Die(1, err.Error())
	}

	log.Infof("Exiting with code 0.")
	os.Exit(0)
}

func removeVerbosityFlags(args []string) (int, []string) {
	verbose := 0
	newArgs := []string{}
	for _, arg := range args {
		if arg == "-v" || arg == "-vv" || arg == "-vvv" {
			verbose += len(arg) - 1
		} else {
			newArgs = append(newArgs, arg)
		}
	}

	return verbose, newArgs
}
