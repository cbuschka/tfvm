package inventory

import (
	statePkg "github.com/cbuschka/tfvm/internal/inventory/state"
	"github.com/cbuschka/tfvm/internal/log"
	"os"
)

// Load loads the inventory state from an optionally present state file
// and merges it with the current state.
func (inventory *Inventory) Load() error {

	err := os.MkdirAll(inventory.cacheDir, 0755)
	if err != nil {
		return err
	}

	stateFilePath, err := inventory.getStateFilePath()
	if err != nil {
		return err
	}

	log.Debugf("State file path: '%s'", stateFilePath)

	state := statePkg.NewEmptyState()

	err = state.LoadFromFile(stateFilePath)
	if err != nil {
		if !os.IsNotExist(err) {
			return err
		} else {
			return nil
		}
	}

	if log.IsTraceEnabled() {
		stateJSON, err := state.ToJSON()
		if err != nil {
			log.Tracef("state.ToJSON() failed: %v", err)
		} else {
			log.Tracef("Inventory state after load:\n%s", stateJSON)
		}
	}

	inventory.mergeInState(state)

	return nil
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
