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

	terraformPath := filepath.Join(basePath, "terraform")
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

const defaultStateJson = `
{
 "lastUpdateTime": "2020-05-01T16:35:41+02:00",
 "terraformReleases": [
  {
   "version": "0.1.0"
  },
  {
   "version": "0.1.1"
  },
  {
   "version": "0.2.0"
  },
  {
   "version": "0.2.1"
  },
  {
   "version": "0.2.2"
  },
  {
   "version": "0.3.0"
  },
  {
   "version": "0.3.1"
  },
  {
   "version": "0.3.5"
  },
  {
   "version": "0.3.6"
  },
  {
   "version": "0.3.7"
  },
  {
   "version": "0.4.0"
  },
  {
   "version": "0.4.1"
  },
  {
   "version": "0.4.2"
  },
  {
   "version": "0.5.0"
  },
  {
   "version": "0.5.1"
  },
  {
   "version": "0.5.3"
  },
  {
   "version": "0.6.0"
  },
  {
   "version": "0.6.1"
  },
  {
   "version": "0.6.2"
  },
  {
   "version": "0.6.3"
  },
  {
   "version": "0.6.4"
  },
  {
   "version": "0.6.5"
  },
  {
   "version": "0.6.6"
  },
  {
   "version": "0.6.7"
  },
  {
   "version": "0.6.8"
  },
  {
   "version": "0.6.9"
  },
  {
   "version": "0.6.10"
  },
  {
   "version": "0.6.11"
  },
  {
   "version": "0.6.12"
  },
  {
   "version": "0.6.13"
  },
  {
   "version": "0.6.14"
  },
  {
   "version": "0.6.15"
  },
  {
   "version": "0.6.16"
  },
  {
   "version": "0.7.0"
  },
  {
   "version": "0.7.1"
  },
  {
   "version": "0.7.2"
  },
  {
   "version": "0.7.3"
  },
  {
   "version": "0.7.4"
  },
  {
   "version": "0.7.5"
  },
  {
   "version": "0.7.6"
  },
  {
   "version": "0.7.7"
  },
  {
   "version": "0.7.8"
  },
  {
   "version": "0.7.9"
  },
  {
   "version": "0.7.10"
  },
  {
   "version": "0.7.11"
  },
  {
   "version": "0.7.12"
  },
  {
   "version": "0.7.13"
  },
  {
   "version": "0.8.0"
  },
  {
   "version": "0.8.1"
  },
  {
   "version": "0.8.2"
  },
  {
   "version": "0.8.3"
  },
  {
   "version": "0.8.4"
  },
  {
   "version": "0.8.5"
  },
  {
   "version": "0.8.6"
  },
  {
   "version": "0.8.7"
  },
  {
   "version": "0.8.8"
  },
  {
   "version": "0.9.0"
  },
  {
   "version": "0.9.1"
  },
  {
   "version": "0.9.2"
  },
  {
   "version": "0.9.3"
  },
  {
   "version": "0.9.4"
  },
  {
   "version": "0.9.5"
  },
  {
   "version": "0.9.6"
  },
  {
   "version": "0.9.7"
  },
  {
   "version": "0.9.8"
  },
  {
   "version": "0.9.9"
  },
  {
   "version": "0.9.10"
  },
  {
   "version": "0.9.11"
  },
  {
   "version": "0.10.0-beta1"
  },
  {
   "version": "0.10.0-beta2"
  },
  {
   "version": "0.10.0-rc1"
  },
  {
   "version": "0.10.0"
  },
  {
   "version": "0.10.1"
  },
  {
   "version": "0.10.2"
  },
  {
   "version": "0.10.3"
  },
  {
   "version": "0.10.4"
  },
  {
   "version": "0.10.5"
  },
  {
   "version": "0.10.6"
  },
  {
   "version": "0.10.7"
  },
  {
   "version": "0.10.8"
  },
  {
   "version": "0.11.0-beta1"
  },
  {
   "version": "0.11.0-rc1"
  },
  {
   "version": "0.11.0"
  },
  {
   "version": "0.11.1"
  },
  {
   "version": "0.11.2"
  },
  {
   "version": "0.11.3"
  },
  {
   "version": "0.11.4"
  },
  {
   "version": "0.11.5"
  },
  {
   "version": "0.11.6"
  },
  {
   "version": "0.11.7"
  },
  {
   "version": "0.11.8"
  },
  {
   "version": "0.11.9-beta1"
  },
  {
   "version": "0.11.9"
  },
  {
   "version": "0.11.10"
  },
  {
   "version": "0.11.11"
  },
  {
   "version": "0.11.12-beta1"
  },
  {
   "version": "0.11.12"
  },
  {
   "version": "0.11.13"
  },
  {
   "version": "0.11.14"
  },
  {
   "version": "0.11.15-oci"
  },
  {
   "version": "0.12.0-alpha1"
  },
  {
   "version": "0.12.0-alpha2"
  },
  {
   "version": "0.12.0-alpha3"
  },
  {
   "version": "0.12.0-alpha4"
  },
  {
   "version": "0.12.0-beta1"
  },
  {
   "version": "0.12.0-beta2"
  },
  {
   "version": "0.12.0-rc1"
  },
  {
   "version": "0.12.0"
  },
  {
   "version": "0.12.1"
  },
  {
   "version": "0.12.2"
  },
  {
   "version": "0.12.3"
  },
  {
   "version": "0.12.4"
  },
  {
   "version": "0.12.5"
  },
  {
   "version": "0.12.6"
  },
  {
   "version": "0.12.7"
  },
  {
   "version": "0.12.8"
  },
  {
   "version": "0.12.9"
  },
  {
   "version": "0.12.10"
  },
  {
   "version": "0.12.11"
  },
  {
   "version": "0.12.12"
  },
  {
   "version": "0.12.13"
  },
  {
   "version": "0.12.14"
  },
  {
   "version": "0.12.15"
  },
  {
   "version": "0.12.16"
  },
  {
   "version": "0.12.17"
  },
  {
   "version": "0.12.18"
  },
  {
   "version": "0.12.19"
  },
  {
   "version": "0.12.20"
  },
  {
   "version": "0.12.21"
  },
  {
   "version": "0.12.22"
  },
  {
   "version": "0.12.23"
  },
  {
   "version": "0.12.24"
  }
 ]
}`
