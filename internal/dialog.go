package tfvm

import (
	"fmt"
	"os"
)

func Die(exitCode int, format string, a ...interface{}) {
	fmt.Printf("%s\n", fmt.Sprintf(format, a...))
	os.Exit(exitCode)
}

func Print(format string, a ...interface{}) {
	fmt.Printf("%s\n", fmt.Sprintf(format, a...))
}
