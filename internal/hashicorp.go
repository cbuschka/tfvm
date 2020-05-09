package tfvm

import (
	"fmt"
	"github.com/hashicorp/go-version"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"runtime"
	"sort"
	"strings"
)

func ListTerraformReleases() ([]TerraformVersion, error) {

	releasesPage, err := downloadReleasesPage()
	if err != nil {
		return nil, err
	}

	return extractReleases(releasesPage)
}

func downloadReleasesPage() (string, error) {
	url := fmt.Sprintf("%s/index.html", getReleasesBaseUrl())
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

func extractReleases(releasePage string) ([]TerraformVersion, error) {
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

	releases := make([]TerraformVersion, len(semVersions))
	for index, semVersion := range semVersions {
		releases[index] = TerraformVersion{Version: semVersion}
	}

	return releases, nil
}

func (release *TerraformVersion) GetUrl() string {
	return fmt.Sprintf("%s/%s/terraform_%s_%s_%s.zip", getReleasesBaseUrl(),
		release.Version.String(), release.Version, runtime.GOOS, runtime.GOARCH)
}

func getReleasesBaseUrl() string {
	baseUrl := os.Getenv("TFVM_RELEASES_BASE_URL")
	if baseUrl != "" {
		return baseUrl
	}

	return "https://releases.hashicorp.com/terraform"
}
