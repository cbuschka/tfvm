package inventory

import (
	"fmt"
	"github.com/cbuschka/tfvm/internal/log"
	platformPkg "github.com/cbuschka/tfvm/internal/platform"
	"github.com/cbuschka/tfvm/internal/util"
	"github.com/cbuschka/tfvm/internal/version"
	"github.com/mitchellh/go-homedir"
	"os"
	"path/filepath"
)

type fs interface {
	IsDir(path string) (bool, error)
	GetHomeDir() (string, error)
}

type realFs struct {
}

func getRealFs() fs {
	return fs(&realFs{})
}

func (d *realFs) IsDir(path string) (bool, error) {
	return util.IsDir(path)
}

func (d *realFs) GetHomeDir() (string, error) {
	return homedir.Dir()
}

// GetTerraformBasePath returns the base path for a terraform installation.
func (inventory *Inventory) GetTerraformBasePath(tfRelease *version.TerraformVersion, platform platformPkg.Platform) (string, error) {
	inventoryDir := inventory.getInventoryDir()

	versionedTfPath := filepath.Join(inventoryDir, "v1", "installed", tfRelease.String(),
		fmt.Sprintf("%s_%s", platform.Os, platform.Arch))
	return versionedTfPath, nil
}

func (inventory *Inventory) getTerraformPath(tfRelease *version.TerraformVersion, platform platformPkg.Platform) (string, error) {
	basePath, err := inventory.GetTerraformBasePath(tfRelease, platform)
	if err != nil {
		return "", err
	}

	var terraformPath string
	if platform.IsWindows() {
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

func (inventory *Inventory) getInventoryDir() string {
	return inventory.cacheDir
}

func getInventoryDir(fs fs, platform platformPkg.Platform) (string, error) {
	invDirFromEnv := os.Getenv("TFVM_DIR")
	log.Debugf("Environment var TFVM_DIR: '%s'", invDirFromEnv)
	if invDirFromEnv != "" {
		return invDirFromEnv, nil
	}

	return getDefaultInventoryDir(fs, platform)
}

func getDefaultInventoryDir(fs fs, platform platformPkg.Platform) (string, error) {
	homeDir, err := fs.GetHomeDir()
	if err != nil {
		return "", err
	}

	oldInventoryDir := filepath.Join(homeDir, ".tfvm")
	oldInventoryDirExists, err := fs.IsDir(oldInventoryDir)
	if err != nil {
		return "", err
	}

	if oldInventoryDirExists {
		return oldInventoryDir, nil
	}

	if platform.IsMacOS() {
		return getDefaultInventoryDirForMacOS(homeDir), nil
	}

	return getDefaultInventoryDirForXDG(homeDir), nil
}

func getDefaultInventoryDirForMacOS(homeDir string) string {
	return filepath.Join(homeDir, "Library", "Caches", "tfvm")
}

func getDefaultInventoryDirForXDG(homeDir string) string {
	dotCacheDir := os.Getenv("XDG_CACHE_HOME")
	if dotCacheDir == "" {
		dotCacheDir = filepath.Join(homeDir, ".cache")
	}

	dotCacheTfvmDir := filepath.Join(dotCacheDir, "tfvm")
	return dotCacheTfvmDir

}

func (inventory *Inventory) getStateFilePath() (string, error) {
	inventoryDir := inventory.getInventoryDir()
	statefilepath := filepath.Join(inventoryDir, "v1", "state.json")
	return statefilepath, nil
}
