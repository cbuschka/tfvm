package tfvm

import "strings"

type BuildInfo struct {
	// Build info of tfvm
	version string
	commitish string
	buildTime string
}

var version string
var commitish string
var buildTime string

func GetBuildInfo() BuildInfo {
	return BuildInfo{version: version, commitish: commitish, buildTime: strings.Replace(buildTime, "_", " ", -1)}
}
