package tfvm

import (
	"fmt"
)

func RunTfvmWhichCommand(args []string) error {

	tfVersion, err := GetConfiguredVersion()
	if err != nil {
		return err
	}

	if tfVersion == "" {
		fmt.Printf("No terraform version configured.\n")
	} else {
		fmt.Printf("Configured terraform is %s.\n", tfVersion)
	}

	return nil
}
