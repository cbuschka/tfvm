package commands

import (
	"github.com/cbuschka/tfvm/internal/util"
)

// RunTfvmCommand runs tfvm commands.
func RunTfvmCommand(args []string) error {

	if len(args) == 0 {
		printUsage()
		return nil
	} else if args[0] == "list" {
		return RunTfvmListCommand(args[1:])
	} else if args[0] == "uptodate" {
		return RunTfvmUptodateCommand(args[1:])
	} else if args[0] == "install" {
		return RunTfvmInstallCommand(args[1:])
	} else if args[0] == "which" {
		return RunTfvmWhichCommand(args[1:])
	} else if args[0] == "help" {
		printUsage()
		return nil
	} else if args[0] == "info" {
		return printInfo()
	} else if args[0] == "use" {
		return RunTfvmUseCommand(args[1:])
	} else if args[0] == "version" {
		printVersion()
		return nil
	} else if args[0] == "terraform" {
		return RunTerraformCommand(args[1:])
	} else {
		util.Print("Unsupported command: '%s'", args[0])
		printUsage()
		util.Die(1, "")
		return nil
	}
}
