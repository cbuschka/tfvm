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

	for _, tfRelease := range tfReleases {
		builds, err := remote.ListTerraformBuilds(tfRelease)
		if err != nil {
			util.Die(1, "Getting terraform builds for release %s failed: %s", tfRelease.String(), err.Error())
		}

		err = state.FillBuilds(tfRelease, builds)
		if err != nil {
			util.Die(1, "Filling builds for release %s failed: %s", tfRelease.String(), err.Error())
		}
	}

	data, err := state.Marshall()
	if err != nil {
		util.Die(1, "Marshalling failed: %s", err.Error())
	}

	_, err = fmt.Fprintf(os.Stdout, string(data))
	if err != nil {
		panic(err)
	}
}
