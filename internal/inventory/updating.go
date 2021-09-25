package inventory

import (
	"github.com/cbuschka/tfvm/internal/inventory/state"
	"github.com/cbuschka/tfvm/internal/log"
	"time"
)

// Update updates the inventory state.
func (inventory *Inventory) Update() error {
	log.Debugf("Updating inventory...")

	tfReleaseVersions, err := inventory.remoteProvider.ListTerraformReleases()
	if err != nil {
		return err
	}

	for _, tfReleaseVersion := range tfReleaseVersions {
		tfReleaseState, found := inventory.TerraformReleases[tfReleaseVersion.String()]
		if !found {
			tfReleaseState = &state.TerraformReleaseState{Version: tfReleaseVersion, Builds: []*state.TerraformReleaseBuildState{}}
			inventory.TerraformReleases[tfReleaseVersion.String()] = tfReleaseState
		}

		if tfReleaseState.Builds == nil || len(tfReleaseState.Builds) == 0 {
			err := inventory.UpdateBuilds(tfReleaseState)
			if err != nil {
				return err
			}
		}
	}

	inventory.LastUpdateTime = time.Now()

	log.Infof("Updated inventory.")

	return nil
}

// UpdateBuilds updates the build mata data of release tfReleaseState.Version.
func (inventory *Inventory) UpdateBuilds(tfReleaseState *state.TerraformReleaseState) error {

	log.Debugf("Updating builds for %s...", tfReleaseState.Version.String())

	builds, err := inventory.remoteProvider.ListTerraformBuilds(tfReleaseState.Version)
	if err != nil {
		return err
	}

	err = tfReleaseState.AddMissingBuilds(builds)
	if err != nil {
		return err
	}

	log.Infof("Updated inventory build for %s.", tfReleaseState.Version.String())

	return nil
}
