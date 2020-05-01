package tfvm

import (
	"fmt"
)

func RunTfvmWhichCommand(args []string) error {

	config, err := GetConfiguration()
	if err != nil {
		if IsNoConfigExists(err) {
			fmt.Printf("No terraform version configured.\n")
			return nil
		}
		return err
	}

	fmt.Printf("Configured terraform is %s (%s).\n", config.version, config.file)
	return nil
}
