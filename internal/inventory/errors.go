package inventory

import "fmt"

const noSuchTerraformReleaseBuild = "no such build"
const osArchUnsupportedByTerraform = "unsupported os/arch"

// NewNoSuchTerraformReleaseBuild creates a new error.
func NewNoSuchTerraformReleaseBuild() error {
	return fmt.Errorf(noSuchTerraformReleaseBuild)
}

// IsNoSuchTerraformReleaseBuild tests if err is an NoSuchTerraformReleaseBuild error.
func IsNoSuchTerraformReleaseBuild(err error) bool {
	return err.Error() == noSuchTerraformReleaseBuild
}

// NewOsArchUnsupportedByTerraform creates a new error.
func NewOsArchUnsupportedByTerraform() error {
	return fmt.Errorf(osArchUnsupportedByTerraform)
}

// IsOsArchUnsupportedByTerraform tests if err is an OsArchUnsupportedByTerraform error.
func IsOsArchUnsupportedByTerraform(err error) bool {
	return err.Error() == osArchUnsupportedByTerraform
}
