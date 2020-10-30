package build

import "strings"

// Info represents information about the current tfvm version.
type Info struct {
	Version   string
	Commitish string
	Repo      string
	Author    string
	BuildTime string
	OS        string
	Arch      string
}

var (
	buildInfoVersion   string
	buildInfoCommitish string
	buildInfoBuildTime string
	buildInfoOs        string
	buildInfoArch      string
)

// GetBuildInfo gives meta data about the current tfvm version.
func GetBuildInfo() Info {
	return Info{Version: buildInfoVersion,
		Repo:      "https://github.com/cbuschka/tfvm",
		Author:    "Cornelius Buschka <cbuschka@gmail.com>, https://github.com/cbuschka",
		Commitish: buildInfoCommitish,
		BuildTime: strings.Replace(buildInfoBuildTime, "_", " ", -1),
		Arch:      buildInfoArch,
		OS:        buildInfoOs}
}
