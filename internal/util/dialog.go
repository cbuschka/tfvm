package util

import (
	"fmt"
	"github.com/cbuschka/tfvm/internal/log"
	"os"
)

var forceQuiet = false

// Die prints a message and exists the process with exitCode.
func Die(exitCode int, format string, a ...interface{}) {
	_, _ = fmt.Fprintf(os.Stderr, "%s\n", fmt.Sprintf(format, a...))
	log.Infof("Exiting with code %d.", exitCode)
	os.Exit(exitCode)
}

// IsOutputTerminal answers if stdout is a terminal.
func IsOutputTerminal() bool {
	fileInfo, _ := os.Stdout.Stat()
	return (fileInfo.Mode() & os.ModeCharDevice) != 0
}

// SuppressOutput suppresses output messages.
func SuppressOutput() {
	forceQuiet = true
}

// Print prints a message to console.
func Print(format string, a ...interface{}) {
	if forceQuiet {
		return
	}

	fmt.Printf("%s\n", fmt.Sprintf(format, a...))
}
