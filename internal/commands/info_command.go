package commands

import (
	"github.com/cbuschka/tfvm/internal/build"
	inventoryPkg "github.com/cbuschka/tfvm/internal/inventory"
	"github.com/cbuschka/tfvm/internal/util"
	"runtime"
)

func printInfo() error {
	inventory, err := inventoryPkg.GetInventory()
	if err != nil {
		return err
	}

	buildInfo := build.GetBuildInfo()
	util.Print(`Tfvm version is %s.
Built from commitish %s, on %s for %s/%s.
OS is %s, arch is %s.
Cache dir is %s.`,
		buildInfo.Version,
		buildInfo.Commitish, buildInfo.BuildTime, buildInfo.OS, buildInfo.Arch,
		runtime.GOOS, runtime.GOARCH,
		inventory.GetCacheDir())

	return nil
}
