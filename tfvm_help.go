package tfvm

import (
	"fmt"
)

func printUsage() {
	fmt.Printf(`
Usage: tfvm <sub command> or terraform <terraform opts*>

 terraform <terraform opts*> - Run terraform configured via .tfvmrc file.
 list - List terraform versions.
 help - Print this usage information.

`)
}
