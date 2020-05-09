package tfvm

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/mitchellh/go-homedir"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

const noSuchTerraformReleaseMsg = "no such terraform release"

func IsNoSuchTerraformRelease(err error) bool {
	return err.Error() == noSuchTerraformReleaseMsg
}

func newNoSuchTerraformRelease() error {
	return errors.New(noSuchTerraformReleaseMsg)
}

type TerraformRelease struct {
	Version string `json:"version"`
}

type StateStruct struct {
	LastUpdateTime    string             `json:"lastUpdateTime"`
	TerraformReleases []TerraformRelease `json:"terraformReleases"`
}

type Inventory struct {
	LastUpdateTime    time.Time
	TerraformReleases []TerraformRelease
}

func GetInventory() (*Inventory, error) {
	inventory := Inventory{LastUpdateTime: time.Now(), TerraformReleases: make([]TerraformRelease, 0)}
	err := inventory.initInventory()
	if err != nil {
		return nil, err
	}

	return &inventory, nil
}

func (inventory *Inventory) GetTerraformReleases() []TerraformRelease {
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

func (inventory *Inventory) GetLatestRelease() TerraformRelease {
	return inventory.TerraformReleases[len(inventory.TerraformReleases)-1]
}

func (inventory *Inventory) GetTerraformRelease(version string) (TerraformRelease, error) {
	if version == "latest" {
		return inventory.GetLatestRelease(), nil
	}

	for _, tfRelease := range inventory.TerraformReleases {
		if tfRelease.Version == version {
			return tfRelease, nil
		}
	}

	return TerraformRelease{}, newNoSuchTerraformRelease()
}

func (inventory *Inventory) GetTerraform(tfRelease TerraformRelease) (*Terraform, error) {
	tfPath, err := inventory.getTerraformPath(tfRelease)
	if err != nil {
		return nil, err
	}

	if _, err := os.Stat(tfPath); err != nil {
		return nil, err
	}

	return &Terraform{version: tfRelease.Version, path: tfPath}, nil
}

func (inventory *Inventory) GetTerraformBasePath(tfRelease TerraformRelease) (string, error) {
	inventoryDir, err := getInventoryDir()
	if err != nil {
		return "", err
	}

	versionedTfPath := filepath.Join(inventoryDir, "v1", "installed", tfRelease.Version, fmt.Sprintf("%s_%s", runtime.GOOS, runtime.GOARCH))
	return versionedTfPath, nil
}

func (inventory *Inventory) getTerraformPath(tfRelease TerraformRelease) (string, error) {
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

func (inventory *Inventory) IsTerraformInstalled(tfRelease TerraformRelease) (bool, error) {
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
	homeDir, err := homedir.Dir()
	if err != nil {
		return "", err
	}

	return filepath.Join(homeDir, ".tfvm"), nil
}

func (inventory *Inventory) loadState() error {
	statefilepath, err := getStateFilePath()
	if err != nil {
		return err
	}

	file, err := ioutil.ReadFile(statefilepath)
	if err != nil {
		if os.IsNotExist(err) {
			file = []byte(defaultStateJson)
		} else {
			return err
		}
	}

	state := StateStruct{}
	err = json.Unmarshal(file, &state)
	if err != nil {
		return err
	}

	inventory.TerraformReleases = state.TerraformReleases
	inventory.LastUpdateTime, err = time.Parse(time.RFC3339, state.LastUpdateTime)
	if err != nil {
		return err
	}

	return nil
}

func getStateFilePath() (string, error) {
	inventoryDir, err := getInventoryDir()
	if err != nil {
		return "", err
	}
	statefilepath := filepath.Join(inventoryDir, "v1", "state.json")
	return statefilepath, nil
}

func (inventory *Inventory) saveState() error {
	statefilepath, err := getStateFilePath()
	if err != nil {
		return err
	}

	if err = os.MkdirAll(filepath.Dir(statefilepath), os.ModePerm); err != nil {
		return err
	}

	state := StateStruct{}
	state.LastUpdateTime = inventory.LastUpdateTime.Format(time.RFC3339)
	state.TerraformReleases = inventory.TerraformReleases

	file, err := json.MarshalIndent(state, "", " ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(statefilepath, file, 0644)
	if err != nil {
		return err
	}

	return nil
}
