package commands

import (
	"github.com/cbuschka/tfvm/internal/build"
	"github.com/cbuschka/tfvm/internal/util"
)

func printVersion() {
	buildInfo := build.GetBuildInfo()
	util.Print("%s", buildInfo.Version)
}
