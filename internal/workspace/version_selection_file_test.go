package workspace_test

import (
	workspacePkg "github.com/cbuschka/tfvm/internal/workspace"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestReadSingleLineFromWhenEmpty(t *testing.T) {

	content, _ := workspacePkg.ReadSingleLineFrom(strings.NewReader(""))

	assert.Equal(t, "", content)
}

func TestReadSingleLineFromWithSingleLine(t *testing.T) {

	content, _ := workspacePkg.ReadSingleLineFrom(strings.NewReader("\nline\n"))

	assert.Equal(t, "line", content)
}

func TestReadSingleLineFromWithSingleLineWithinComments(t *testing.T) {

	content, _ := workspacePkg.ReadSingleLineFrom(strings.NewReader("# comment\nline\n # comment 2"))

	assert.Equal(t, "line", content)
}

func TestReadSingleLineFromWithMultipleLines(t *testing.T) {

	content, _ := workspacePkg.ReadSingleLineFrom(strings.NewReader("# comment\nline\n # comment 2\nline 2\n"))

	assert.Equal(t, "line 2", content)
}
