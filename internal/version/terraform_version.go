package version

import (
	"errors"
	"github.com/hashicorp/go-version"
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
	Version *version.Version
}
