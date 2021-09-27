package inventory

import (
	platformPkg "github.com/cbuschka/tfvm/internal/platform"
	"github.com/stretchr/testify/assert"
	"runtime"
	"strings"
	"testing"
)

type mockFs struct {
	existingDirs []string
}

func (m *mockFs) IsDir(path string) (bool, error) {
	for _, v := range m.existingDirs {
		if v == path {
			return true, nil
		}
	}
	return false, nil
}

func (m *mockFs) GetHomeDir() (string, error) {
	return adjustPathForRuntimeOS("/HOME"), nil
}

func TestInventoryDirOnMacOSIsDotTfvmIfExists(t *testing.T) {

	inventoryDir, err := getInventoryDir(fs(&mockFs{existingDirs: []string{adjustPathForRuntimeOS("/HOME/.tfvm")}}), platformPkg.Platform{Os: "darwin", Arch: "arm64"})
	if err != nil {
		t.Fatal(err)
		return
	}

	assert.Equal(t, adjustPathForRuntimeOS("/HOME/.tfvm"), inventoryDir)
}

func TestInventoryDirOnMacOSIsLibraryCachesTfvmIfDotTfvmNotPresent(t *testing.T) {

	inventoryDir, err := getInventoryDir(fs(&mockFs{existingDirs: []string{}}), platformPkg.Platform{Os: "darwin", Arch: "arm64"})
	if err != nil {
		t.Fatal(err)
		return
	}

	assert.Equal(t, adjustPathForRuntimeOS("/HOME/Library/Caches/tfvm"), inventoryDir)
}

func TestInventoryDirOnLinuxIsDotTfvmIfExists(t *testing.T) {

	inventoryDir, err := getInventoryDir(fs(&mockFs{existingDirs: []string{adjustPathForRuntimeOS("/HOME/.tfvm")}}), platformPkg.Platform{Os: "linux", Arch: "arm64"})
	if err != nil {
		t.Fatal(err)
		return
	}

	assert.Equal(t, adjustPathForRuntimeOS("/HOME/.tfvm"), inventoryDir)
}

func TestInventoryDirOnLinuxIsDotCachesTfvmIfDotTfvmNotPresent(t *testing.T) {

	inventoryDir, err := getInventoryDir(fs(&mockFs{existingDirs: []string{}}), platformPkg.Platform{Os: "linux", Arch: "arm64"})
	if err != nil {
		t.Fatal(err)
		return
	}

	assert.Equal(t, adjustPathForRuntimeOS("/HOME/.cache/tfvm"), inventoryDir)
}

func adjustPathForRuntimeOS(path string) string {
	if runtime.GOOS == "windows" {
		return strings.ReplaceAll(path, "/", "\\")
	}

	return path
}
