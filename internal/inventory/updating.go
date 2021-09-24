package inventory

import (
	"github.com/cbuschka/tfvm/internal/inventory/state"
	"time"
)

// Update updates the inventory state.
func (inventory *Inventory) Update() error {
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

	return nil
}

// UpdateBuilds updates the build mata data of release tfReleaseState.Version.
func (inventory *Inventory) UpdateBuilds(tfReleaseState *state.TerraformReleaseState) error {

	builds, err := inventory.remoteProvider.ListTerraformBuilds(tfReleaseState.Version)
	if err != nil {
		return err
	}

	return tfReleaseState.AddMissingBuilds(builds)
}
