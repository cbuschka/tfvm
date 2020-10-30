package inventory

import (
	"encoding/json"
	"github.com/cbuschka/tfvm/internal/version"
	goversion "github.com/hashicorp/go-version"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

// TerraformReleaseState describes a single terraform release state.
type TerraformReleaseState struct {
	Version string `json:"version"`
}

// State describes the on disk inventory state format.
type State struct {
	LastUpdateTime    string                  `json:"lastUpdateTime"`
	TerraformReleases []TerraformReleaseState `json:"terraformReleases"`
}

func (inventory *Inventory) loadState() error {
	statefilepath, err := getStateFilePath()
	if err != nil {
		return err
	}

	file, err := ioutil.ReadFile(statefilepath)
	if err != nil {
		if os.IsNotExist(err) {
			file = []byte(defaultStateJSON)
		} else {
			return err
		}
	}

	state := State{}
	err = json.Unmarshal(file, &state)
	if err != nil {
		return err
	}

	tfReleases := make([]*version.TerraformVersion, len(state.TerraformReleases))
	for index, tfReleaseState := range state.TerraformReleases {
		tfReleaseVersion, err := goversion.NewVersion(tfReleaseState.Version)
		if err != nil {
			return err
		}
		tfReleases[index] = &version.TerraformVersion{Version: tfReleaseVersion}
	}

	inventory.terraformReleasesAsc = tfReleases
	inventory.lastUpdateTime, err = time.Parse(time.RFC3339, state.LastUpdateTime)
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

	state := State{}
	state.LastUpdateTime = inventory.lastUpdateTime.Format(time.RFC3339)

	tfReleaseStates := make([]TerraformReleaseState, len(inventory.terraformReleasesAsc))
	for index, tfRelease := range inventory.terraformReleasesAsc {
		tfReleaseStates[index] = TerraformReleaseState{Version: tfRelease.Version.String()}
	}

	state.TerraformReleases = tfReleaseStates

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
