package platform

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsMacOS(t *testing.T) {

	darwinArm64 := Platform{Os: "darwin", Arch: "arm64"}
	assert.True(t, darwinArm64.IsMacOS())

	darwinAmd64 := Platform{Os: "darwin", Arch: "amd64"}
	assert.True(t, darwinAmd64.IsMacOS())

	windows386 := Platform{Os: "windows", Arch: "386"}
	assert.False(t, windows386.IsMacOS())
}

func TestIsWindows(t *testing.T) {

	windows386 := Platform{Os: "windows", Arch: "386"}
	assert.True(t, windows386.IsWindows())

	darwinAmd64 := Platform{Os: "darwin", Arch: "amd64"}
	assert.False(t, darwinAmd64.IsWindows())
}
