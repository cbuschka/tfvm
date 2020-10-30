package build

import (
	"github.com/tcnksm/go-latest"
)

// UpdateStatus represents the current update status.
type UpdateStatus struct {
	updateAvailable bool
	latestVersion   string
	currentVersion  string
}

// UpdateAvailable answers if an update is available.
func (updateStatus *UpdateStatus) UpdateAvailable() bool {
	return updateStatus.updateAvailable
}

// CurrentVersion gives the current tfvm version.
func (updateStatus *UpdateStatus) CurrentVersion() string {
	return updateStatus.currentVersion
}

// LatestVersion gives the latest tfvm version.
func (updateStatus *UpdateStatus) LatestVersion() string {
	return updateStatus.latestVersion
}

// GetUpdateStatus creates an instance of UpdateStatus.
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
