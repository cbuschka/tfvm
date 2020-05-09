package main

import (
	"github.com/cbuschka/tfvm/internal/commands"
	"github.com/cbuschka/tfvm/internal/util"
	"os"
	"path/filepath"
	"runtime"
)

func main() {
	programName := filepath.Base(os.Args[0])
	var err error
	if programName == "terraform" || (runtime.GOOS == "windows" && programName == "terraform.exe") {
		err = commands.RunTerraformCommand(os.Args[1:])
	} else if programName != "terraform" && len(os.Args) > 1 && os.Args[1] == "terraform" {
		err = commands.RunTerraformCommand(os.Args[2:])
	} else {
		err = commands.RunTfvmCommand(os.Args[1:])
	}

	if err != nil {
		util.Die(1, err.Error())
	}
}
