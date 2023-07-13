package remote

import (
	"fmt"
	"github.com/cbuschka/tfvm/internal/build"
	"github.com/cbuschka/tfvm/internal/log"
	platformPkg "github.com/cbuschka/tfvm/internal/platform"
	"github.com/cbuschka/tfvm/internal/util"
	versionPkg "github.com/cbuschka/tfvm/internal/version"
	"io/ioutil"
	"math/big"
	"net/http"
	"regexp"
	"strings"
)

// TerraformBuild is a specific build of a terraform release
type TerraformBuild struct {
	Os           string
	Arch         string
	DownloadPath string
}

type hashicorpRemote struct {
}

// ListTerraformReleases lists terraform versions from the hashicorp website.
func (hashicorp *hashicorpRemote) ListTerraformReleases() ([]*versionPkg.TerraformVersion, error) {

	releasesPage, err := downloadReleasesPage()
	if err != nil {
		return nil, err
	}

	return extractReleases(releasesPage)
}

// ListTerraformBuilds lists builds of a particular release
func (hashicorp *hashicorpRemote) ListTerraformBuilds(release *versionPkg.TerraformVersion) ([]*TerraformBuild, error) {
	buildsPage, err := downloadBuildsPage(release)
	if err != nil {
		return nil, err
	}

	return extractBuilds(buildsPage)
}

func getHTMLContentsFrom(url string) (string, error) {

	log.Debugf("Executing HTTP GET for %s...", url)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}
	client := &http.Client{}
	userAgent := fmt.Sprintf("tfvm/%s (https://github.com/cbuschka/tfvm)", build.GetBuildInfo().Version)
	req.Header.Set("User-Agent", userAgent)
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != 200 {
		log.Debugf("Response to GET for '%s': %d/%s", url, resp.StatusCode, resp.Status)
		return "", fmt.Errorf("downloading %s failed", url)
	}

	defer resp.Body.Close()
	html, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	log.Tracef("Response to GET for '%s': %d\n%s", url, resp.StatusCode, html)

	return string(html), nil
}

func (hashicorp *hashicorpRemote) ListChecksums(release *versionPkg.TerraformVersion) (map[string]big.Int, error) {

	log.Debug("Listing checksums...")

	url := fmt.Sprintf("%s/%s/terraform_%s_SHA256SUMS", getReleasesBaseURL(), release.String(), release.String())
	contents, err := getHTMLContentsFrom(url)
	if err != nil {
		return nil, err
	}

	checksums, err := extractChecksums(contents)
	if err != nil {
		return nil, err
	}

	log.Tracef("Checksums found: %v", checksums)

	log.Info("Got checksums")

	return checksums, nil
}

func extractChecksums(contents string) (map[string]big.Int, error) {

	re, err := regexp.Compile("\\s*([0-9a-f]+)\\s+terraform_[^_]+_([^_]+)_([^_]+).zip")
	if err != nil {
		return nil, err
	}

	matchSets := re.FindAllStringSubmatch(contents, -1)

	checksumsByFilename := map[string]big.Int{}
	for _, matchSet := range matchSets {
		sha256ChecksumStr := strings.TrimSpace(matchSet[1])
		sha256Checksum := new(big.Int)
		sha256Checksum, _ = sha256Checksum.SetString(sha256ChecksumStr, 16)
		os := strings.TrimSpace(matchSet[2])
		arch := strings.TrimSpace(matchSet[3])
		platformStr := fmt.Sprintf("%s/%s", os, arch)
		if err != nil {
			return nil, err
		}
		checksumsByFilename[platformStr] = *sha256Checksum
	}

	return checksumsByFilename, nil
}

func downloadReleasesPage() (string, error) {
	url := fmt.Sprintf("%s/index.html", getReleasesBaseURL())

	log.Infof("Download URL for releases page: %s", url)

	return getHTMLContentsFrom(url)
}

func downloadBuildsPage(release *versionPkg.TerraformVersion) (string, error) {
	url := fmt.Sprintf("%s/index.html", getBuildBaseURL(release))

	log.Infof("Downloads URL for builds page: %s", url)

	return getHTMLContentsFrom(url)
}

func extractReleases(releasePage string) ([]*versionPkg.TerraformVersion, error) {
	re, err := regexp.Compile(">terraform_([^<]+)</a>")
	if err != nil {
		return nil, err
	}

	matchSets := re.FindAllStringSubmatch(releasePage, -1)

	releases := make([]*versionPkg.TerraformVersion, len(matchSets))
	for index, matchSet := range matchSets {
		semVersionStr := strings.TrimSpace(matchSet[1])
		version, err := versionPkg.NewTerraformVersion(semVersionStr)
		if err != nil {
			return nil, err
		}
		releases[index] = version
	}

	return releases, nil
}

// GetURL gives the remote url to a particular terraform version on the hashicorp site.
func GetURL(release *versionPkg.TerraformVersion, platform platformPkg.Platform) string {

	url := fmt.Sprintf("%s/%s/terraform_%s_%s_%s.zip", getReleasesBaseURL(),
		release.String(), release.String(), platform.Os, platform.Arch)

	log.Debugf("Download URL for terraform %s on %s: ", release.String(), platform)

	return url
}

func getReleasesBaseURL() string {
	baseURL := util.GetFirstEnv("TFVM_TERRAFORM_RELEASES_BASE_URL")

	log.Debugf("Environment var TFVM_TERRAFORM_RELEASES_BASE_URL: '%s'", baseURL)

	if baseURL != "" {
		return baseURL
	}

	baseURL = "https://releases.hashicorp.com/terraform"

	log.Infof("Download base URL: %s", baseURL)

	return baseURL
}

func getBuildBaseURL(release *versionPkg.TerraformVersion) string {
	url := fmt.Sprintf("%s/%s", getReleasesBaseURL(), release.String())

	return url
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
