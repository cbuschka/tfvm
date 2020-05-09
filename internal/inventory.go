package tfvm

import (
	"fmt"
	"github.com/mitchellh/go-homedir"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

type Inventory struct {
	LastUpdateTime    time.Time
	TerraformReleases []TerraformVersion
}

func GetInventory() (*Inventory, error) {
	inventory := Inventory{LastUpdateTime: time.Now(), TerraformReleases: make([]TerraformVersion, 0)}
	err := inventory.initInventory()
	if err != nil {
		return nil, err
	}

	return &inventory, nil
}

func (inventory *Inventory) GetTerraformReleases() []TerraformVersion {
	return inventory.TerraformReleases
}

func (inventory *Inventory) Update() error {
	tfReleases, err := ListTerraformReleases()
	if err == nil && len(tfReleases) > 0 {
		inventory.TerraformReleases = tfReleases
		inventory.LastUpdateTime = time.Now()
		err = inventory.saveState()
		if err != nil {
			return err
		}
	}

	return nil
}

func (inventory *Inventory) GetLatestRelease() *TerraformVersion {
	if len(inventory.TerraformReleases) == 0 {
		panic("Inventory has no terraform releases.")
	}

	return &inventory.TerraformReleases[len(inventory.TerraformReleases)-1]
}

func (inventory *Inventory) GetTerraformRelease(versionSpec *TerraformVersionSpec) (*TerraformVersion, error) {
	latestTfRelease := inventory.GetLatestRelease()

	for _, tfRelease := range inventory.TerraformReleases {
		if versionSpec.Matches(&tfRelease, latestTfRelease) {
			return &tfRelease, nil
		}
	}

	return nil, newNoSuchTerraformRelease()
}

func (inventory *Inventory) GetTerraform(tfRelease *TerraformVersion) (*Terraform, error) {
	tfPath, err := inventory.getTerraformPath(tfRelease)
	if err != nil {
		return nil, err
	}

	if _, err := os.Stat(tfPath); err != nil {
		return nil, err
	}

	return &Terraform{version: tfRelease.Version.String(), path: tfPath}, nil
}

func (inventory *Inventory) GetTerraformBasePath(tfRelease *TerraformVersion) (string, error) {
	inventoryDir, err := getInventoryDir()
	if err != nil {
		return "", err
	}

	versionedTfPath := filepath.Join(inventoryDir, "v1", "installed", tfRelease.Version.String(),
		fmt.Sprintf("%s_%s", runtime.GOOS, runtime.GOARCH))
	return versionedTfPath, nil
}

func (inventory *Inventory) getTerraformPath(tfRelease *TerraformVersion) (string, error) {
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

func (inventory *Inventory) IsTerraformInstalled(tfRelease *TerraformVersion) (bool, error) {
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

	homeDir, err := homedir.Dir()
	if err != nil {
		return "", err
	}

	return filepath.Join(homeDir, ".tfvm"), nil
}
