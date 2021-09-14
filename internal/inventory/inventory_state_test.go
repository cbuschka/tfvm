package inventory

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

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
