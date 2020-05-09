package build

import (
	"github.com/tcnksm/go-latest"
)

type UpdateStatus struct {
	updateAvailable bool
	latestVersion   string
	currentVersion  string
}

func (updateStatus *UpdateStatus) UpdateAvailable() bool {
	return updateStatus.updateAvailable
}

func (updateStatus *UpdateStatus) CurrentVersion() string {
	return updateStatus.currentVersion
}

func (updateStatus *UpdateStatus) LatestVersion() string {
	return updateStatus.latestVersion
}

func GetUpdateStatus() *UpdateStatus {
	githubTag := &latest.GithubTag{
		Owner:             "cbuschka",
		Repository:        "tfvm",
		FixVersionStrFunc: latest.DeleteFrontV(),
	}

	buildInfo := GetBuildInfo()
	result, err := latest.Check(githubTag, buildInfo.Version)
	if err != nil || result == nil {
		return &UpdateStatus{updateAvailable: false, latestVersion: buildInfo.Version, currentVersion: buildInfo.Version}
	}
	return &UpdateStatus{updateAvailable: result.Outdated, latestVersion: result.Current, currentVersion: buildInfo.Version}
}
