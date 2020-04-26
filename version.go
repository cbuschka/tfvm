package tfvm

import (
	"fmt"
)

// VERSION The current tfvm version.
const VERSION = "v0.1"

func printVersion() {
	fmt.Printf("Tfvm version %s.\n", VERSION)
}
