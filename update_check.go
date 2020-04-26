package tfvm

import "github.com/tcnksm/go-latest"

type UpdateStatus struct {
	updateAvailable bool
	latestVersion string
	currentVersion string
}

func GetUpdateStatus() *UpdateStatus {
	githubTag := &latest.GithubTag{
		Owner: "cbuschka",
		Repository: "tfvm",
	}

	buildInfo := GetBuildInfo()
	result, err := latest.Check(githubTag, buildInfo.version)
	if err == nil || result == nil {
		return &UpdateStatus{updateAvailable: false, latestVersion: buildInfo.version, currentVersion: buildInfo.version}
	}
	return &UpdateStatus{updateAvailable: result.Outdated, latestVersion: result.Current, currentVersion: buildInfo.version}
}
