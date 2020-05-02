package main

import (
	"github.com/cbuschka/tfvm/internal"
	"os"
	"path/filepath"
	"runtime"
)

func main() {
	programName := filepath.Base(os.Args[0])
	var err error
	if programName == "terraform" || (runtime.GOOS == "windows" && programName == "terraform.exe") {
		err = tfvm.RunTerraformCommand(os.Args[1:])
	} else if programName != "terraform" && len(os.Args) > 1 && os.Args[1] == "terraform" {
		err = tfvm.RunTerraformCommand(os.Args[2:])
	} else {
		err = tfvm.RunTfvmCommand(os.Args[1:])
	}

	if err != nil {
		tfvm.Die(1, err.Error())
	}
}
