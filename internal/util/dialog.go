package util

import (
	"fmt"
	"os"
)

// Die prints a message and exists the process with exitCode.
func Die(exitCode int, format string, a ...interface{}) {
	fmt.Printf("%s\n", fmt.Sprintf(format, a...))
	os.Exit(exitCode)
}

// Print prints a message to console.
func Print(format string, a ...interface{}) {
	fmt.Printf("%s\n", fmt.Sprintf(format, a...))
}
