package inventory

import (
	"fmt"
	"github.com/cbuschka/tfvm/internal/remote"
	"github.com/cbuschka/tfvm/internal/version"
	"github.com/mitchellh/go-homedir"
	"os"
	"path/filepath"
	"runtime"
	"sort"
	"time"
)

type Inventory struct {
	lastUpdateTime       time.Time
	terraformReleasesAsc []*version.TerraformVersion
	cacheDir             string
}

func GetInventory() (*Inventory, error) {
	inventory := Inventory{lastUpdateTime: time.Now(), terraformReleasesAsc: make([]*version.TerraformVersion, 0)}
	err := inventory.initInventory()
	if err != nil {
		return nil, err
	}

	return &inventory, nil
}

func (inventory *Inventory) GetCacheDir() string {
	return inventory.cacheDir
}

func (inventory *Inventory) GetTerraformReleasesAsc() []*version.TerraformVersion {
	return inventory.terraformReleasesAsc[:]
}

func (inventory *Inventory) Update() error {
	tfReleases, err := remote.ListTerraformReleases()

	sort.Sort(version.Collection(tfReleases))

	if err == nil && len(tfReleases) > 0 {
		inventory.terraformReleasesAsc = tfReleases
		inventory.lastUpdateTime = time.Now()
		err = inventory.saveState()
		if err != nil {
			return err
		}
	}

	return nil
}

func (inventory *Inventory) GetLatestRelease() *version.TerraformVersion {
	if len(inventory.terraformReleasesAsc) == 0 {
		panic("Inventory has no terraform releases.")
	}

	return inventory.terraformReleasesAsc[len(inventory.terraformReleasesAsc)-1]
}

func (inventory *Inventory) GetTerraformRelease(versionSpec *version.TerraformVersionSpec) (*version.TerraformVersion, error) {
	latestTfRelease := inventory.GetLatestRelease()

	for index := range inventory.terraformReleasesAsc {
		tfRelease := inventory.terraformReleasesAsc[len(inventory.terraformReleasesAsc)-index-1]
		if versionSpec.Matches(tfRelease, latestTfRelease) {
			return tfRelease, nil
		}
	}

	return nil, version.NewNoSuchTerraformRelease()
}

func (inventory *Inventory) GetTerraform(tfRelease *version.TerraformVersion) (*Terraform, error) {
	tfPath, err := inventory.getTerraformPath(tfRelease)
	if err != nil {
		return nil, err
	}

	if _, err := os.Stat(tfPath); err != nil {
		return nil, err
	}

	return newTerraform(tfRelease.Version.String(), tfPath), nil
}

func (inventory *Inventory) GetTerraformBasePath(tfRelease *version.TerraformVersion) (string, error) {
	inventoryDir, err := getInventoryDir()
	if err != nil {
		return "", err
	}

	versionedTfPath := filepath.Join(inventoryDir, "v1", "installed", tfRelease.Version.String(),
		fmt.Sprintf("%s_%s", runtime.GOOS, runtime.GOARCH))
	return versionedTfPath, nil
}

func (inventory *Inventory) getTerraformPath(tfRelease *version.TerraformVersion) (string, error) {
	basePath, err := inventory.GetTerraformBasePath(tfRelease)
	if err != nil {
		return "", err
	}

	var terraformPath string
	if runtime.GOOS == "windows" {
		terraformPath = filepath.Join(basePath, "terraform.exe")
	} else {
		terraformPath = filepath.Join(basePath, "terraform")
	}

	return terraformPath, nil
}

func (inventory *Inventory) IsTerraformInstalled(tfRelease *version.TerraformVersion) (bool, error) {
	versionedTfPath, err := inventory.getTerraformPath(tfRelease)
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

func (inventory *Inventory) initInventory() error {
	inventoryDir, err := getInventoryDir()
	if err != nil {
		return err
	}

	inventory.cacheDir = inventoryDir

	err = os.MkdirAll(inventoryDir, 0755)
	if err != nil {
		return err
	}

	err = inventory.loadState()

	return err
}

func getInventoryDir() (string, error) {
	invDirFromEnv := os.Getenv("TFVM_DIR")
	if invDirFromEnv != "" {
		return invDirFromEnv, nil
	}

	return getDefaultInventoryDir()
}

func existsDir(path string) (bool, error) {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}

		return false, err
	}

	return true, nil
}

func getDefaultInventoryDir() (string, error) {
	homeDir, err := homedir.Dir()
	if err != nil {
		return "", err
	}

	oldInventoryDir := filepath.Join(homeDir, ".tfvm")
	oldInventoryDirExists, err := existsDir(oldInventoryDir)
	if err != nil {
		return "", err
	}

	if oldInventoryDirExists {
		return oldInventoryDir, nil
	}

	dotCacheDir := os.Getenv("XDG_CACHE_HOME")
	if dotCacheDir == "" {
		dotCacheDir = filepath.Join(homeDir, ".cache")
	}

	dotCacheTfvmDir := filepath.Join(dotCacheDir, "tfvm")
	return dotCacheTfvmDir, nil
}
