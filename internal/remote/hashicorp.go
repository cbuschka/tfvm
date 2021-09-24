package remote

import (
	"fmt"
	"github.com/cbuschka/tfvm/internal/util"
	"github.com/cbuschka/tfvm/internal/version"
	goversion "github.com/hashicorp/go-version"
	"io/ioutil"
	"net/http"
	"regexp"
	"runtime"
	"strings"
)

// TerraformBuild is a specific build of a terraform release
type TerraformBuild struct {
	Os           string
	Arch         string
	DownloadPath string
}

// ListTerraformReleases lists terraform versions from the hashicorp website.
func ListTerraformReleases() ([]*version.TerraformVersion, error) {

	releasesPage, err := downloadReleasesPage()
	if err != nil {
		return nil, err
	}

	return extractReleases(releasesPage)
}

// ListTerraformBuilds lists builds of a particular release
func ListTerraformBuilds(release *version.TerraformVersion) ([]*TerraformBuild, error) {
	buildsPage, err := downloadBuildsPage(release)
	if err != nil {
		return nil, err
	}

	return extractBuilds(buildsPage)
}

func downloadReleasesPage() (string, error) {
	url := fmt.Sprintf("%s/index.html", getReleasesBaseURL())
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

func downloadBuildsPage(release *version.TerraformVersion) (string, error) {
	url := fmt.Sprintf("%s/index.html", getBuildBaseURL(release))
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

func extractReleases(releasePage string) ([]*version.TerraformVersion, error) {
	re, err := regexp.Compile(">terraform_([^<]+)</a>")
	if err != nil {
		return nil, err
	}

	matchSets := re.FindAllStringSubmatch(releasePage, -1)

	releases := make([]*version.TerraformVersion, len(matchSets))
	for index, matchSet := range matchSets {
		semVersionStr := strings.TrimSpace(matchSet[1])
		semVersion, err := goversion.NewVersion(semVersionStr)
		if err != nil {
			return nil, err
		}
		releases[index] = &version.TerraformVersion{Version: semVersion}
	}

	return releases, nil
}

// GetURL gives the remote url to a particular terraform version on the hashicorp site.
func GetURL(release *version.TerraformVersion) string {
	tfArch := util.GetFirstEnv("TFVM_TERRAFORM_ARCH", "TERRAFORM_ARCH")
	if tfArch == "" {
		tfArch = runtime.GOARCH
	}

	tfOs := util.GetFirstEnv("TFVM_TERRAFORM_OS", "TERRAFORM_OS")
	if tfOs == "" {
		tfOs = runtime.GOOS
	}

	return fmt.Sprintf("%s/%s/terraform_%s_%s_%s.zip", getReleasesBaseURL(),
		release.Version.String(), release.Version, tfOs, tfArch)
}

func getReleasesBaseURL() string {
	baseURL := util.GetFirstEnv("TFVM_TERRAFORM_RELEASES_BASE_URL")
	if baseURL != "" {
		return baseURL
	}

	return "https://releases.hashicorp.com/terraform"
}

func getBuildBaseURL(release *version.TerraformVersion) string {
	return fmt.Sprintf("%s/%s", getReleasesBaseURL(), release.String())
}

func extractBuilds(buildsPage string) ([]*TerraformBuild, error) {
	re, err := regexp.Compile("<a\\s+data-product=\"terraform\"\\s+data-version=\"([^\"]+)\"\\s+data-os=\"([^\"]+)\"\\s+data-arch=\"([^\"]+)\"\\s+href=\"([^\"]+)\"[^>]*>[^<]+</a>")
	if err != nil {
		return nil, err
	}

	matchSets := re.FindAllStringSubmatch(buildsPage, -1)

	builds := make([]*TerraformBuild, 0)
	for _, matchSet := range matchSets {
		_ = strings.TrimSpace(matchSet[1])
		os := strings.TrimSpace(matchSet[2])
		arch := strings.TrimSpace(matchSet[3])
		path := strings.TrimSpace(matchSet[4])
		if err != nil {
			return nil, err
		}
		builds = append(builds, &TerraformBuild{Os: os, Arch: arch, DownloadPath: path})
	}

	return builds, nil
}
