package workspace_test

import (
	workspacePkg "github.com/cbuschka/tfvm/internal/workspace"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestGetVersionSpecFromHcl(t *testing.T) {

	cwd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
		return
	}

	workspace := workspacePkg.Workspace{RootDir: cwd}
	selection, err := workspace.GetTerraformVersionSelection()
	if err != nil {
		t.Errorf("Cannot get tf version selection: %s", err.Error())
		return
	}

	assert.Regexp(t, "hcl file.*", selection.Source())
}
