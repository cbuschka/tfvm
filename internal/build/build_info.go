package build

import "strings"

// Build info of tfvm
type BuildInfo struct {
	Version   string
	Commitish string
	Repo      string
	Author    string
	BuildTime string
	OS        string
	Arch      string
}

var (
	buildInfoVersion string
	buildInfoCommitish string
	buildInfoBuildTime string
	buildInfoOs string
	buildInfoArch string
)

func GetBuildInfo() BuildInfo {
	return BuildInfo{Version: buildInfoVersion,
		Repo:      "https://github.com/cbuschka/tfvm",
		Author:    "Cornelius Buschka <cbuschka@gmail.com>, https://github.com/cbuschka",
		Commitish: buildInfoCommitish,
		BuildTime: strings.Replace(buildInfoBuildTime, "_", " ", -1),
		Arch:      buildInfoArch,
		OS:        buildInfoOs}
}
