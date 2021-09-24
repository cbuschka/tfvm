package main

import (
	"bytes"
	"fmt"
	inventoryPkg "github.com/cbuschka/tfvm/internal/inventory"
	"github.com/cbuschka/tfvm/internal/util"
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
