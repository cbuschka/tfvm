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

		if inventory.mustBeUpdated(tfReleaseState) {
			err := inventory.UpdateTerraformRelease(tfReleaseState)
			if err != nil {
				return err
			}
		}
	}

	inventory.LastUpdateTime = time.Now()

	log.Infof("Updated inventory.")

	return nil
}

func (inventory *Inventory) mustBeUpdated(tfReleaseState *state.TerraformReleaseState) bool {

	if tfReleaseState.Builds == nil || len(tfReleaseState.Builds) == 0 {
		return true
	}

	for _, build := range tfReleaseState.Builds {
		if build == nil || build.SHA256Checksum == "" {
			return true
		}
	}

	return false
}

// UpdateTerraformRelease updates the build meta data of a terraform release.
func (inventory *Inventory) UpdateTerraformRelease(tfReleaseState *state.TerraformReleaseState) error {

	err := inventory.updateMissingBuilds(tfReleaseState)
	if err != nil {
		return err
	}

	return inventory.updateMissingChecksums(tfReleaseState)
}

func (inventory *Inventory) updateMissingBuilds(tfReleaseState *state.TerraformReleaseState) error {

	log.Debugf("Updating builds for %s...", tfReleaseState.Version.String())

	builds, err := inventory.remoteProvider.ListTerraformBuilds(tfReleaseState.Version)
	if err != nil {
		return err
	}

	err = tfReleaseState.AddMissingBuilds(builds)
	if err != nil {
		return err
	}

	log.Infof("Updated build for %s.", tfReleaseState.Version.String())
	return nil
}

func (inventory *Inventory) updateMissingChecksums(tfReleaseState *state.TerraformReleaseState) error {
	log.Debugf("Updating checksums for %s...", tfReleaseState.Version.String())

	checksumsByPlatform, err := inventory.remoteProvider.ListChecksums(tfReleaseState.Version)
	if err != nil {
		return err
	}

	err = tfReleaseState.AddMissingChecksums(checksumsByPlatform)
	if err != nil {
		return err
	}

	log.Infof("Updated checksums for %s.", tfReleaseState.Version.String())
	return nil
}
