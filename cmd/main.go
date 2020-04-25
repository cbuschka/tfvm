package main

import (
	"fmt"
	"github.com/cbuschka/tfvm"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		panic(fmt.Sprintf("Usage: ..."))
	} else if os.Args[0] == "terraform" {
		tfvm.RunTerraformSubCommand(os.Args[1:])
	} else if os.Args[0] != "terraform" && len(os.Args)>1 && os.Args[1] == "terraform" {
		tfvm.RunTerraformSubCommand(os.Args[2:])
	} else {
		panic(fmt.Sprintf("Unsupported command: %s", os.Args[1]))
	}
}
