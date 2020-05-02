package tfvm

import "runtime"

func printVersion() {
	buildInfo := GetBuildInfo()
	Print("Tfvm version is %s, commitish=%s, built on %s for %s/%s. Runnning on %s/%s.", buildInfo.version, buildInfo.commitish, buildInfo.buildTime, buildInfo.os, buildInfo.arch, runtime.GOOS, runtime.GOARCH)
}
