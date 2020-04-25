package tfvm

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func RunTerraformSubCommand(args []string) error {

	dotTfvmRcFile, err := GetNearestDotTfvmRcFileFromCwd()
	if err != nil {
		panic(fmt.Sprintf("%v", err.Error()))
	}

	tfVersionBytes, err := ioutil.ReadFile(dotTfvmRcFile)
	if err != nil {
		panic(fmt.Sprintf("%v", err.Error()))
	}
	tfVersion := string(tfVersionBytes)
	tfVersion = strings.TrimSpace(tfVersion)

	inventory, err := GetInventory()
	if err != nil {
		panic(fmt.Sprintf("Getting inventory failed: %v", err.Error()))
	}

	tf, err := inventory.GetTerraform(tfVersion)
	if err != nil {
		panic(fmt.Sprintf("Terraform %s not found: %v", tfVersion, err))
	}

	exitCode, err := tf.Run(os.Args[1:]...)
	if err != nil || exitCode != 0 {
		panic(fmt.Sprintf("Running tf failed: %v exitCode=%d", err, exitCode))
	}

	return nil
}
