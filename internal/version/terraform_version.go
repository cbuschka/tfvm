package version

import (
	"errors"
	goversion "github.com/hashicorp/go-version"
	"strconv"
)

const noSuchTerraformReleaseMsg = "no such terraform release"
const noTerraformReleases = "inventory has no terraform releases"

// IsNoSuchTerraformRelease answer if an error describes an unknown terraform version.
func IsNoSuchTerraformRelease(err error) bool {
	return err.Error() == noSuchTerraformReleaseMsg
}

// NewNoSuchTerraformRelease describes an unknown terraform release version.
func NewNoSuchTerraformRelease() error {
	return errors.New(noSuchTerraformReleaseMsg)
}

// NewNoTerraformReleases describes that now terraform releases are available.
func NewNoTerraformReleases() error {
	return errors.New(noTerraformReleases)
}

// String gives string representation of version.
func (release *TerraformVersion) String() string {
	return release.version.String()
}

// Version retrieves semantic version of terraform version.
func (release *TerraformVersion) Version() *goversion.Version {
	return release.version
}

// TerraformVersion is a representation of a terraform version.
type TerraformVersion struct {
	version *goversion.Version
}

// SafeNewTerraformVersion parses version or panics.
func SafeNewTerraformVersion(version string) *TerraformVersion {
	v, err := NewTerraformVersion(version)
	if err != nil {
		panic(err)
	}
	return v
}

// NewTerraformVersion parses version.
func NewTerraformVersion(version string) (*TerraformVersion, error) {
	tfReleaseVersion, err := goversion.NewVersion(version)
	if err != nil {
		return nil, err
	}
	return &TerraformVersion{version: tfReleaseVersion}, err
}

// UnmarshalJSON unmarshalls terraform version from json.
func (release *TerraformVersion) UnmarshalJSON(data []byte) error {

	unquoted, err := strconv.Unquote(string(data))
	if err != nil {
		return err
	}
	version, err := goversion.NewVersion(unquoted)
	if err != nil {
		return err
	}
	release.version = version
	return nil
}

// MarshalJSON converts json into a version.
func (release *TerraformVersion) MarshalJSON() ([]byte, error) {

	unquoted := strconv.Quote(release.String())
	return []byte(unquoted), nil
}
