package main

import (
	"bytes"
	"fmt"
	inventoryPkg "github.com/cbuschka/tfvm/internal/inventory"
	"github.com/cbuschka/tfvm/internal/util"
	versionPkg "github.com/cbuschka/tfvm/internal/version"
	"os"
)

func main() {
	inventory, err := inventoryPkg.NewInventory()
	if err != nil {
		util.Die(1, "Creating inventory failed: %s", err.Error())
	}

	err = inventory.Update()
	if err != nil {
		util.Die(1, "Updating inventory failed: %s", err.Error())
	}

	versionsWithoutChecksums := []string{"0.12.0-alpha2", "0.12.0-alpha3", "0.12.0-alpha4"}
	for _, versionStr := range versionsWithoutChecksums {
		release, err := inventory.GetTerraformRelease(versionPkg.SafeNewTerraformVersion(versionStr))
		if err != nil {
			panic(err)
		}

		for _, build := range release.Builds {
			if build.SHA256Checksum == "" {
				build.SHA256Checksum = "n/a"
			}
		}
	}

	buffer := bytes.NewBuffer([]byte{})
	_, err = inventory.WriteTo(buffer)
	if err != nil {
		util.Die(1, "Marshalling failed: %s", err.Error())
	}

	_, err = fmt.Fprintf(os.Stdout, buffer.String())
	if err != nil {
		panic(err)
	}
}
