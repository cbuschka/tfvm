package inventory

import (
	"fmt"
	"github.com/cbuschka/tfvm/internal/inventory/state"
	"github.com/cbuschka/tfvm/internal/log"
	platformPkg "github.com/cbuschka/tfvm/internal/platform"
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
func (inventory *Inventory) GetTerraform(tfRelease *version.TerraformVersion, platform platformPkg.Platform) (*Terraform, error) {

	log.Debugf("Looking up terraform %s on %s...", tfRelease.String(), platform)

	tfPath, err := inventory.getTerraformPath(tfRelease, platform)
	if err != nil {
		return nil, err
	}

	if _, err := os.Stat(tfPath); err != nil {
		return nil, err
	}

	return newTerraform(tfRelease.String(), platform.Os, platform.Arch, tfPath), nil
}

// IsTerraformInstalledOnAnyPlatform answers if a particular terraform version is already installed locally.
func (inventory *Inventory) IsTerraformInstalledOnAnyPlatform(tfRelease *version.TerraformVersion) (bool, error) {
	for _, platform := range inventory.platforms {

		log.Debugf("Checking if terraform %s is installed on %s", tfRelease.String(), platform)

		installed, err := inventory.IsTerraformInstalled(tfRelease, platform)
		if err != nil {
			return false, err
		}

		if installed {
			log.Infof("Terraform %s is installed on %s", tfRelease.String(), platform)

			return true, nil
		} else {
			log.Debugf("Terraform %s not installed on %s", tfRelease.String(), platform)
		}
	}

	log.Infof("Terraform %s not installed on any platform", tfRelease.String())

	return false, nil
}

// IsTerraformInstalled answers if a particular terraform version is already installed locally.
func (inventory *Inventory) IsTerraformInstalled(tfRelease *version.TerraformVersion, platform platformPkg.Platform) (bool, error) {
	versionedTfPath, err := inventory.getTerraformPath(tfRelease, platform)
	if err != nil {
		return false, err
	}

	log.Debugf("Terraform path for %s on %s: %s", tfRelease.String(), platform, versionedTfPath)

	if _, err := os.Stat(versionedTfPath); err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}

		return false, err
	}

	log.Infof("Terraform %s on %s found at %s", tfRelease.String(), platform, versionedTfPath)

	return true, nil
}
