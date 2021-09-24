package inventory

import (
	"fmt"
	"github.com/cbuschka/tfvm/internal/version"
	"github.com/mitchellh/go-homedir"
	"os"
	"path/filepath"
)

// GetTerraformBasePath returns the base path for a terraform installation.
func (inventory *Inventory) GetTerraformBasePath(tfRelease *version.TerraformVersion, os string, arch string) (string, error) {
	inventoryDir, err := getInventoryDir()
	if err != nil {
		return "", err
	}

	versionedTfPath := filepath.Join(inventoryDir, "v1", "installed", tfRelease.String(),
		fmt.Sprintf("%s_%s", os, arch))
	return versionedTfPath, nil
}

func (inventory *Inventory) getTerraformPath(tfRelease *version.TerraformVersion, os string, arch string) (string, error) {
	basePath, err := inventory.GetTerraformBasePath(tfRelease, os, arch)
	if err != nil {
		return "", err
	}

	var terraformPath string
	if os == "windows" {
		terraformPath = filepath.Join(basePath, "terraform.exe")
	} else {
		terraformPath = filepath.Join(basePath, "terraform")
	}

	return terraformPath, nil
}

// GetCacheDir gives the cache dir path of the inventory.
func (inventory *Inventory) GetCacheDir() string {
	return inventory.cacheDir
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

func getStateFilePath() (string, error) {
	inventoryDir, err := getInventoryDir()
	if err != nil {
		return "", err
	}
	statefilepath := filepath.Join(inventoryDir, "v1", "state.json")
	return statefilepath, nil
}
