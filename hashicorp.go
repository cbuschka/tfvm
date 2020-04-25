package tfvm

import (
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

type TerraformRelease struct {
	version string
	baseUrl string
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
	var releases []TerraformRelease
	for _, matchSet := range matchSets {
		version := strings.TrimSpace(matchSet[1])
		release := TerraformRelease{version: version}
		releases = append(releases, release)
	}

	return releases, nil
}
