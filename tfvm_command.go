package tfvm

import (
	"fmt"
)

func RunTfvmCommand(args []string) error {

	if len(args) == 0 {
		printUsage()
		return nil
	} else if args[0] == "list" {
		return RunTfvmListCommand(args[1:])
	} else if args[0] == "help" {
		printUsage()
		return nil
	} else {
		panic(fmt.Sprintf("Unsupported command: %s", args))
	}

	return nil
}
