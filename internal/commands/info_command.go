package commands

import (
	"github.com/cbuschka/tfvm/internal/build"
	"github.com/cbuschka/tfvm/internal/util"
	"runtime"
)

func printInfo() {
	buildInfo := build.GetBuildInfo()
	util.Print(`Tfvm version is %s.
Built from commitish %s, on %s for %s/%s.
OS is %s, arch is %s.`,
		buildInfo.Version,
		buildInfo.Commitish, buildInfo.BuildTime, buildInfo.OS, buildInfo.Arch,
		runtime.GOOS, runtime.GOARCH)
}
