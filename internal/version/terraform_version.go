package version

import (
	"errors"
	goversion "github.com/hashicorp/go-version"
	"strconv"
)

const noSuchTerraformReleaseMsg = "no such terraform release"

// IsNoSuchTerraformRelease answer if an error describes an unknown terraform version.
func IsNoSuchTerraformRelease(err error) bool {
	return err.Error() == noSuchTerraformReleaseMsg
}

// NewNoSuchTerraformRelease describes an unknown terraform release version.
func NewNoSuchTerraformRelease() error {
	return errors.New(noSuchTerraformReleaseMsg)
}

// String gives string representation of version.
func (release *TerraformVersion) String() string {
	return release.Version.String()
}

// TerraformVersion is a representation of a terraform version.
type TerraformVersion struct {
	Version *goversion.Version
}

func SafeNewTerraformVersion(version string) *TerraformVersion {
	v, err := NewTerraformVersion(version)
	if err != nil {
		panic(err)
	}
	return v
}

func NewTerraformVersion(version string) (*TerraformVersion, error) {
	tfReleaseVersion, err := goversion.NewVersion(version)
	if err != nil {
		return nil, err
	}
	return &TerraformVersion{Version: tfReleaseVersion}, err
}

func (tfVersion *TerraformVersion) UnmarshalJSON(data []byte) error {

	unquoted, err := strconv.Unquote(string(data))
	if err != nil {
		return err
	}
	version, err := goversion.NewVersion(unquoted)
	if err != nil {
		return err
	}
	tfVersion.Version = version
	return nil
}

func (tfVersion *TerraformVersion) MarshalJSON() ([]byte, error) {

	unquoted := strconv.Quote(tfVersion.String())
	return []byte(unquoted), nil
}
