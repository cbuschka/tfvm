package tfvm

import (
	"fmt"
)

// VERSION The current tfvm version.
const VERSION = "devel"

func printVersion() {
	fmt.Printf("Tfvm version %s.\n", VERSION)
}
