package version

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLatestCanPeParsed(t *testing.T) {

	spec, err := ParseTerraformVersionSpec("latest")
	if err != nil {
		t.Errorf("Cannot parse.")
	}

	assert.Equal(t, spec.String(), "latest")
}
