package build

import "strings"

// Build info of tfvm
type BuildInfo struct {
	Version   string
	Commitish string
	BuildTime string
	OS        string
	Arch      string
}

var buildInfoVersion string
var buildInfoCommitish string
var buildInfoBuildTime string
var buildInfoOs string
var buildInfoArch string

func GetBuildInfo() BuildInfo {
	return BuildInfo{Version: buildInfoVersion, Commitish: buildInfoCommitish, BuildTime: strings.Replace(buildInfoBuildTime, "_", " ", -1), Arch: buildInfoArch, OS: buildInfoOs}
}
