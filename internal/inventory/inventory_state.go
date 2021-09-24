package inventory

import (
	"encoding/json"
	"github.com/cbuschka/tfvm/internal/remote"
	"github.com/cbuschka/tfvm/internal/version"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

// TerraformReleaseBuildState describes a single terraform release build state.
type TerraformReleaseBuildState struct {
	Os           string `json:"os"`
	Arch         string `json:"arch"`
	DownloadPath string `json:"download_path"`
}

// TerraformReleaseState describes a single terraform release state.
type TerraformReleaseState struct {
	Version *version.TerraformVersion     `json:"version"`
	Builds  []*TerraformReleaseBuildState `json:"builds"`
}

// State describes the on disk inventory state format.
type State struct {
	LastUpdateTime    time.Time                `json:"lastUpdateTime"`
	TerraformReleases []*TerraformReleaseState `json:"terraformReleases"`
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
		tfReleases[index] = tfReleaseState.Version
	}

	inventory.terraformReleasesAsc = tfReleases
	inventory.lastUpdateTime = state.LastUpdateTime
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

	if err := os.MkdirAll(filepath.Dir(statefilepath), os.ModePerm); err != nil {
		return err
	}

	state := State{}
	state.LastUpdateTime = inventory.lastUpdateTime

	tfReleaseStates := make([]*TerraformReleaseState, len(inventory.terraformReleasesAsc))
	for index, tfRelease := range inventory.terraformReleasesAsc {
		tfReleaseStates[index] = &TerraformReleaseState{Version: tfRelease, Builds: nil}
	}

	state.TerraformReleases = tfReleaseStates

	data, err := state.Marshall()
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(statefilepath, data, 0644)
	if err != nil {
		return err
	}

	return nil
}

// Fill fills state with terraform releases and updates state.LastUpdateTime.
func (state *State) Fill(terraformReleasesAsc []*version.TerraformVersion) {
	state.LastUpdateTime = time.Now()

	tfReleaseStates := make([]*TerraformReleaseState, len(terraformReleasesAsc))
	for index, tfRelease := range terraformReleasesAsc {
		tfReleaseStates[index] = &TerraformReleaseState{Version: tfRelease, Builds: nil}
	}

	state.TerraformReleases = tfReleaseStates

}

// Marshall converts state into json on disk format.
func (state *State) Marshall() ([]byte, error) {

	data, err := json.MarshalIndent(state, "", " ")
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (state *State) FillBuilds(release *version.TerraformVersion, builds []*remote.TerraformBuild) error {
	found := false
	for _, tfRelease := range state.TerraformReleases {
		if tfRelease.Version.String() == release.String() {
			found = true
			buildStates := make([]*TerraformReleaseBuildState, len(builds))
			for index, build := range builds {
				buildStates[index] = &TerraformReleaseBuildState{Os: build.Os, Arch: build.Arch, DownloadPath: build.DownloadPath}
			}
			tfRelease.Builds = buildStates
		}
	}

	if !found {
		return version.NewNoSuchTerraformRelease()
	}

	return nil
}
