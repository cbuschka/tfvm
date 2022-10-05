package commands

import (
	"github.com/cbuschka/tfvm/internal/util"
)

// RunTfenvCommand runs tfenv commands.
func RunTfenvCommand(args []string) error {

	if len(args) == 0 {
		printUsage()
		return nil
	} else if args[0] == "install" {
		return RunTfvmInstallCommand(args[1:])
	} else if args[0] == "exec" {
		return RunTerraformCommand(args[1:])
	} else {
		util.Print("Unsupported command: '%s'", args[0])
		printUsage()
		util.Die(1, "")
		return nil
	}
}
