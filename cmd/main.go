package main

import (
	"fmt"
	"github.com/cbuschka/tfvm"
	"os"
	"path"
)

func main() {
	programName := path.Base(os.Args[0])
	var err error
	if programName == "terraform" {
		err = tfvm.RunTerraformCommand(os.Args[1:])
	} else if programName != "terraform" && len(os.Args) > 1 && os.Args[1] == "terraform" {
		err = tfvm.RunTerraformCommand(os.Args[2:])
	} else {
		err = tfvm.RunTfvmCommand(os.Args[1:])
	}

	if err != nil {
		panic(fmt.Sprintf("%v", err.Error()))
	}
}
