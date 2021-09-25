package state

import (
	"encoding/json"
	"github.com/cbuschka/tfvm/internal/remote"
	"github.com/cbuschka/tfvm/internal/version"
	"io"
	"io/ioutil"
	"os"
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

// ContainsBuild answers if a terraform contains build information about an build for os and arch.
func (terraformRelease *TerraformReleaseState) ContainsBuild(os string, arch string) bool {

	if terraformRelease.Builds == nil {
		return false
	}

	for _, curr := range terraformRelease.Builds {
		if curr != nil && curr.Os == os && curr.Arch == arch {
			return true
		}
	}

	return false
}

// MergeIn merges the given release state into the inventory state.
func (terraformRelease *TerraformReleaseState) MergeIn(other *TerraformReleaseState) {

	if other.Builds == nil {
		return
	}

	for _, currBuild := range other.Builds {
		if currBuild != nil && !terraformRelease.ContainsBuild(currBuild.Os, currBuild.Arch) {
			terraformRelease.Builds = append(terraformRelease.Builds, currBuild)
		}
	}
}

// GetDefaultState retrieves the default state.
func GetDefaultState() (*State, error) {
	state := State{}
	err := json.Unmarshal([]byte(defaultStateJSON), &state)
	if err != nil {
		return nil, err
	}

	return &state, nil
}

// NewEmptyState create a new empty state.
func NewEmptyState() *State {
	return &State{LastUpdateTime: time.Unix(0, 0), TerraformReleases: []*TerraformReleaseState{}}
}

// LoadFromFile loads the json from statefilepath into the state struct.
func (state *State) LoadFromFile(statefilepath string) error {

	_, err := os.Stat(statefilepath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}

		return err
	}

	file, err := ioutil.ReadFile(statefilepath)
	if err != nil {
		return err
	}

	err = json.Unmarshal(file, state)
	if err != nil {
		return err
	}

	return nil
}

// ToJSON converts state to json representation.
func (state *State) ToJSON() (string, error) {

	data, err := state.Marshall()
	if err != nil {
		return "", err
	}

	return string(data), nil
}

// WriteTo writes the state as json to writer w.
func (state *State) WriteTo(w io.Writer) (int64, error) {
	data, err := state.Marshall()
	if err != nil {
		return -1, err
	}

	size, err := w.Write(data)
	return int64(size), err
}

// Marshall converts state into json on disk format.
func (state *State) Marshall() ([]byte, error) {

	data, err := json.MarshalIndent(state, "", " ")
	if err != nil {
		return nil, err
	}

	return data, nil
}

// AddMissingBuilds adds missing builds.
func (terraformRelease *TerraformReleaseState) AddMissingBuilds(newBuilds []*remote.TerraformBuild) error {

	for _, newBuild := range newBuilds {
		if !terraformRelease.ContainsBuild(newBuild.Os, newBuild.Arch) {
			terraformRelease.Builds = append(terraformRelease.Builds, &TerraformReleaseBuildState{Os: newBuild.Os, Arch: newBuild.Arch, DownloadPath: newBuild.DownloadPath})
		}
	}

	return nil
}
