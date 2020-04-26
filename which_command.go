package tfvm

import (
	"fmt"
)

func RunTfvmWhichCommand(args []string) error {

	config, err := GetConfiguration()
	if err != nil {
		return err
	}

	if config == nil {
		fmt.Printf("No terraform version configured.\n")
	} else {
		fmt.Printf("Configured terraform is %s (%s).\n", config.version, config.file)
	}

	return nil
}
