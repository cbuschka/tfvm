package main

import (
	"fmt"
	inventoryPkg "github.com/cbuschka/tfvm/internal/inventory"
	"github.com/cbuschka/tfvm/internal/remote"
	"github.com/cbuschka/tfvm/internal/util"
	"github.com/cbuschka/tfvm/internal/version"
	"os"
	"sort"
)

func main() {
	state := inventoryPkg.State{}

	tfReleases, err := remote.ListTerraformReleases()
	if err != nil {
		util.Die(1, "Getting terraform releases failed: %s", err.Error())
	}

	sort.Sort(version.Collection(tfReleases))

	state.Fill(tfReleases)

	data, err := state.Marshall()
	if err != nil {
		util.Die(1, "Marshalling failed: %s", err.Error())
	}

	_, err = fmt.Fprintf(os.Stdout, string(data))
	if err != nil {
		panic(err)
	}
}
