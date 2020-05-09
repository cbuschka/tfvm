package version

import (
	"errors"
	"github.com/hashicorp/go-version"
)

const noSuchTerraformReleaseMsg = "no such terraform release"

func IsNoSuchTerraformRelease(err error) bool {
	return err.Error() == noSuchTerraformReleaseMsg
}

func NewNoSuchTerraformRelease() error {
	return errors.New(noSuchTerraformReleaseMsg)
}

func (release *TerraformVersion) String() string {
	return release.Version.String()
}

type TerraformVersion struct {
	Version *version.Version
}
