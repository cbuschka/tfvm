package commands

import (
	"github.com/cbuschka/tfvm/internal/build"
	inventoryPkg "github.com/cbuschka/tfvm/internal/inventory"
	"github.com/cbuschka/tfvm/internal/util"
	workspacePkg "github.com/cbuschka/tfvm/internal/workspace"
	"runtime"
)

func printInfo() error {
	inventory, err := inventoryPkg.GetInventory()
	if err != nil {
		return err
	}

	workspace, err := workspacePkg.GetWorkspace()
	if err != nil {
		return err
	}

	buildInfo := build.GetBuildInfo()
	util.Print(`Tfvm version is %s.
Built from commitish %s, on %s for %s/%s.
OS is %s, arch is %s.
Cache dir is %s.
Workspace dir is %s.
Config loaded from %v.`,
		buildInfo.Version,
		buildInfo.Commitish, buildInfo.BuildTime, buildInfo.OS, buildInfo.Arch,
		runtime.GOOS, runtime.GOARCH,
		inventory.GetCacheDir(),
		workspace.RootDir,
		workspace.Config.Sources)

	return nil
}
