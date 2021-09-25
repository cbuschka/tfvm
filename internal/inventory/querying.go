package inventory

import (
	"fmt"
	"github.com/cbuschka/tfvm/internal/inventory/state"
	"github.com/cbuschka/tfvm/internal/log"
	"github.com/cbuschka/tfvm/internal/version"
	"os"
	"sort"
)

// GetMatchingTerraformRelease returns the terraform version for a version specification.
func (inventory *Inventory) GetMatchingTerraformRelease(versionSpec *version.TerraformVersionSpec) (*state.TerraformReleaseState, error) {

	terraformReleasesAsc := inventory.GetTerraformReleasesAsc()
	if len(terraformReleasesAsc) == 0 {
		return nil, version.NewNoTerraformReleases()
	}

	latestTfRelease := terraformReleasesAsc[len(terraformReleasesAsc)-1]

	for index := range terraformReleasesAsc {
		tfReleaseVersion := terraformReleasesAsc[len(terraformReleasesAsc)-index-1]
		if versionSpec.Matches(tfReleaseVersion, latestTfRelease) {
			tfRelease, found := inventory.TerraformReleases[tfReleaseVersion.String()]
			if !found {
				return nil, fmt.Errorf("terraform release of best version not found")
			}
			return tfRelease, nil
		}
	}

	return nil, version.NewNoSuchTerraformRelease()
}

// GetLatestRelease returns the newest terraform version known.
func (inventory *Inventory) GetLatestRelease() (*version.TerraformVersion, error) {

	terraformReleasesAsc := inventory.GetTerraformReleasesAsc()

	if len(terraformReleasesAsc) == 0 {
		return nil, version.NewNoTerraformReleases()
	}

	return terraformReleasesAsc[len(terraformReleasesAsc)-1], nil
}

// GetTerraformReleasesAsc lists terraform versions known in ascending order.
func (inventory *Inventory) GetTerraformReleasesAsc() []*version.TerraformVersion {

	tfReleasesAsc := make([]*version.TerraformVersion, 0, len(inventory.TerraformReleases))
	for _, tfReleaseState := range inventory.TerraformReleases {
		tfReleasesAsc = append(tfReleasesAsc, tfReleaseState.Version)
	}

	sort.Sort(version.Collection(tfReleasesAsc))

	return tfReleasesAsc
}

// GetTerraformRelease returns the terraform version.
func (inventory *Inventory) GetTerraformRelease(tfReleaseVersion *version.TerraformVersion) (*state.TerraformReleaseState, error) {
	tfRelease, found := inventory.TerraformReleases[tfReleaseVersion.String()]
	if !found {
		return nil, version.NewNoSuchTerraformRelease()
	}

	return tfRelease, nil
}

// GetTerraform get reference to a terraform installation of a given version.
func (inventory *Inventory) GetTerraform(tfRelease *version.TerraformVersion, osName string, arch string) (*Terraform, error) {

	log.Debugf("Looking up terraform %s on %s/%s...", tfRelease.String(), osName, arch)

	tfPath, err := inventory.getTerraformPath(tfRelease, osName, arch)
	if err != nil {
		return nil, err
	}

	if _, err := os.Stat(tfPath); err != nil {
		return nil, err
	}

	return newTerraform(tfRelease.String(), osName, arch, tfPath), nil
}

// IsTerraformInstalled answers if a particular terraform version is already installed locally.
func (inventory *Inventory) IsTerraformInstalled(tfRelease *version.TerraformVersion, osName string, arch string) (bool, error) {
	versionedTfPath, err := inventory.getTerraformPath(tfRelease, osName, arch)
	if err != nil {
		return false, err
	}

	if _, err := os.Stat(versionedTfPath); err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}

		return false, err
	}

	return true, nil
}
