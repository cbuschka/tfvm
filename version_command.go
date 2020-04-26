package tfvm

import (
	"fmt"
)

func printVersion() {
	buildInfo := GetBuildInfo()
	fmt.Printf("Tfvm version is %s, commitish=%s, built on %s.\n", buildInfo.version, buildInfo.commitish, buildInfo.buildTime)
}
