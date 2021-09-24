package inventory

import (
	statePkg "github.com/cbuschka/tfvm/internal/inventory/state"
	"os"
)

func (inventory *Inventory) Load() error {

	err := os.MkdirAll(inventory.cacheDir, 0755)
	if err != nil {
		return err
	}

	stateFilePath, err := getStateFilePath()
	if err != nil {
		return err
	}

	state := statePkg.NewEmptyState()

	err = state.LoadFromFile(stateFilePath)
	if err != nil {
		if !os.IsNotExist(err) {
			return err
		}
	} else {
		inventory.mergeInState(state)
	}

	return err
}

func (inventory *Inventory) mergeInState(state *statePkg.State) {

	for _, newTfRelease := range state.TerraformReleases {
		existingTfRelease, found := inventory.TerraformReleases[newTfRelease.Version.String()]
		if !found {
			inventory.TerraformReleases[newTfRelease.Version.String()] = newTfRelease
		} else {
			existingTfRelease.MergeIn(newTfRelease)
		}
	}

	if state.LastUpdateTime.After(inventory.LastUpdateTime) {
		inventory.LastUpdateTime = state.LastUpdateTime
	}
}
