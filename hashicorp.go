package tfvm

import (
	"fmt"
	"github.com/hashicorp/go-version"
	"io/ioutil"
	"net/http"
	"regexp"
	"runtime"
	"sort"
	"strings"
)

type TerraformRelease struct {
	version string
	url     string
}

func GetTerraformRelease(version string) TerraformRelease {
	release := createTerraformRelease(version)
	return release
}

func ListTerraformReleases() ([]TerraformRelease, error) {

	releasesPage, err := downloadReleasesPage()
	if err != nil {
		return nil, err
	}

	return extractReleases(releasesPage)
}

func downloadReleasesPage() (string, error) {
	url := "https://releases.hashicorp.com/terraform/index.html"
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	html, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(html), nil
}

func extractReleases(releasePage string) ([]TerraformRelease, error) {
	re, err := regexp.Compile(">terraform_([^<]+)</a>")
	if err != nil {
		return nil, err
	}

	matchSets := re.FindAllStringSubmatch(releasePage, -1)

	semVersions := make([]*version.Version, len(matchSets))
	for index, matchSet := range matchSets {
		semVersionStr := strings.TrimSpace(matchSet[1])
		semVersion, err := version.NewVersion(semVersionStr)
		if err != nil {
			return nil, err
		}
		semVersions[index] = semVersion
	}

	sort.Sort(version.Collection(semVersions))

	releases := make([]TerraformRelease, len(semVersions))
	for index, semVersion := range semVersions {
		releases[index] = createTerraformRelease(semVersion.String())
	}

	return releases, nil
}

func createTerraformRelease(version string) TerraformRelease {
	url := fmt.Sprintf("https://releases.hashicorp.com/terraform/%s/terraform_%s_%s_%s.zip", version, version, runtime.GOOS, runtime.GOARCH)
	return TerraformRelease{version: version, url: url}
}
