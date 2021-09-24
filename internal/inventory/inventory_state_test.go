package inventory

import (
	"encoding/json"
	"github.com/cbuschka/tfvm/internal/version"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUnmarshallState(t *testing.T) {

	state := State{}
	file := []byte(`{
 "lastUpdateTime": "2021-09-14T23:19:31+02:00",
 "terraformReleases": [
  {
   "version": "0.1.0",
   "builds": [
    {
     "os": "darwin",
     "arch": "amd64",
     "download_path": "/terraform/0.1.0/terraform_0.1.0_darwin_amd64.zip"
    }
   ]
  }
 ]
}`)
	err := json.Unmarshal(file, &state)
	if err != nil {
		t.Fatal(err)
		return
	}

	assert.Equal(t, 1, len(state.TerraformReleases))
	assert.Equal(t, "0.1.0", state.TerraformReleases[0].Version.String())
	assert.Equal(t, 1, len(state.TerraformReleases[0].Builds))
	assert.Equal(t, "darwin", state.TerraformReleases[0].Builds[0].Os)
	assert.Equal(t, "amd64", state.TerraformReleases[0].Builds[0].Arch)
	assert.Equal(t, "/terraform/0.1.0/terraform_0.1.0_darwin_amd64.zip", state.TerraformReleases[0].Builds[0].DownloadPath)
}

func TestMarshallState(t *testing.T) {

	state := State{}
	state.TerraformReleases = make([]*TerraformReleaseState, 1)

	state.TerraformReleases[0] = &TerraformReleaseState{Version: version.SafeNewTerraformVersion("1.2.3")}
	state.TerraformReleases[0].Builds = make([]*TerraformReleaseBuildState, 1)
	state.TerraformReleases[0].Builds[0] = &TerraformReleaseBuildState{Os: "os", Arch: "arch", DownloadPath: "downloadPath"}
	marshalledState, err := json.Marshal(state)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "{\"lastUpdateTime\":\"\",\"terraformReleases\":[{\"version\":\"1.2.3\",\"builds\":[{\"os\":\"os\",\"arch\":\"arch\",\"download_path\":\"downloadPath\"}]}]}", string(marshalledState))
}

func TestUnmarshalledStateContainsBuilds(t *testing.T) {

	state := State{}
	file := []byte(defaultStateJSON)
	err := json.Unmarshal(file, &state)
	if err != nil {
		t.Fatal(err)
		return
	}

	assert.GreaterOrEqual(t, len(state.TerraformReleases), 0)
	assert.GreaterOrEqual(t, len(state.TerraformReleases[0].Builds), 0)
}
