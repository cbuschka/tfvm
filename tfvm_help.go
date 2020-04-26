package tfvm

import (
	"fmt"
)

func printUsage() {
	fmt.Printf(`
Usage: tfvm <sub command> or terraform <terraform opts*>

 terraform <terraform opts*> - Run terraform configured via .tfvmrc file.
 list - List terraform versions.
 install - Install terraform version.
 which - Print selected terraform version.
 help - Print this usage information.
 version - Print tfvm version.

For binaries, issues and source code visit https://github.com/cbuschka/tfvm.

`)
}
