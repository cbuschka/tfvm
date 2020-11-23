package workspace_test

import (
	workspacePkg "github.com/cbuschka/tfvm/internal/workspace"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetConfigReturnsNotNil(t *testing.T) {

	workspace, err := workspacePkg.GetWorkspace()
	if err != nil {
		t.Fatal(err)
		return
	}

	config := workspace.GetConfig()

	assert.NotNil(t, config)
}
