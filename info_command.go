package tfvm

import "runtime"

func printInfo() {
	buildInfo := GetBuildInfo()
	Print(`Tfvm version is %s.
Built from commitish %s, on %s for %s/%s.
OS is %s, arch is %s.`,
		buildInfo.version,
		buildInfo.commitish, buildInfo.buildTime, buildInfo.os, buildInfo.arch,
		runtime.GOOS, runtime.GOARCH)
}
