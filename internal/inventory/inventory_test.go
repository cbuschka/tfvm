package inventory

import (
	"github.com/cbuschka/tfvm/internal/inventory/state"
	"github.com/cbuschka/tfvm/internal/remote"
	"github.com/cbuschka/tfvm/internal/version"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewInventoryContainsDefaults(t *testing.T) {

	inventory, err := NewInventory()
	if err != nil {
		t.Fatal(err)
		return
	}

	assert.GreaterOrEqual(t, len(inventory.GetTerraformReleasesAsc()), 1)
}

type MockRemoteProvider struct{}

func (d *MockRemoteProvider) ListTerraformReleases() ([]*version.TerraformVersion, error) {
	return []*version.TerraformVersion{version.SafeNewTerraformVersion("999.0")}, nil
}
func (d *MockRemoteProvider) ListTerraformBuilds(release *version.TerraformVersion) ([]*remote.TerraformBuild, error) {
	return []*remote.TerraformBuild{{Os: "MyOs", Arch: "MyArch", DownloadPath: "MyDownloadPath"}}, nil
}

func TestMergeInAddsMissingVersion(t *testing.T) {

	inventory, err := NewInventory()
	if err != nil {
		t.Fatal(err)
		return
	}
	inventory.remoteProvider = remote.Provider(&MockRemoteProvider{})
	err = inventory.Update()
	if err != nil {
		t.Fatal(err)
		return
	}

	tfRelease999, err := inventory.GetMatchingTerraformRelease(version.SafeParseTerraformVersionSpec("999.0"))
	if err != nil {
		t.Fatal(err)
		return
	}
	assert.Equal(t, "999.0.0", tfRelease999.Version.String())
	assert.Equal(t, []*state.TerraformReleaseBuildState{{Os: "MyOs", Arch: "MyArch", DownloadPath: "MyDownloadPath"}}, tfRelease999.Builds)
}
