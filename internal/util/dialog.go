package util

import (
	"fmt"
	"os"
)

var verbose = false

// Die prints a message and exists the process with exitCode.
func Die(exitCode int, format string, a ...interface{}) {
	fmt.Printf("%s\n", fmt.Sprintf(format, a...))
	os.Exit(exitCode)
}

// Print prints a message to console.
func Print(format string, a ...interface{}) {
	fmt.Printf("%s\n", fmt.Sprintf(format, a...))
}

// Debug prints a debug message to console.
func Debug(format string, a ...interface{}) {
	if verbose {
		fmt.Printf("%s\n", fmt.Sprintf(format, a...))
	}
}
