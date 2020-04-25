package main

import (
	"fmt"
	"github.com/cbuschka/tfvm"
	"os"
	"path"
)

func main() {
	programName := path.Base(os.Args[0])
	if programName == "terraform" {
		tfvm.RunTerraformSubCommand(os.Args[1:])
	} else if programName != "terraform" && len(os.Args)>1 && os.Args[1] == "terraform" {
		tfvm.RunTerraformSubCommand(os.Args[2:])
	} else {
		panic(fmt.Sprintf("Unsupported command: %s", os.Args[1]))
	}
}
