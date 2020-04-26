package tfvm

import "strings"

// Build info of tfvm
type BuildInfo struct {
	version   string
	commitish string
	buildTime string
}

var buildInfoVersion string
var buildInfoCommitish string
var buildInfoBuildTime string

func GetBuildInfo() BuildInfo {
	return BuildInfo{version: buildInfoVersion, commitish: buildInfoCommitish, buildTime: strings.Replace(buildInfoBuildTime, "_", " ", -1)}
}
