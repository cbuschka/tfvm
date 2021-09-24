package inventory

import "fmt"

const noSuchTerraformReleaseBuild = "no such build"
const osArchUnsupportedByTerraform = "unsupported os/arch"

func NewNoSuchTerraformReleaseBuild() error {
	return fmt.Errorf(noSuchTerraformReleaseBuild)
}

func IsNoSuchTerraformReleaseBuild(err error) bool {
	return err.Error() == noSuchTerraformReleaseBuild
}

func NewOsArchUnsupportedByTerraform() error {
	return fmt.Errorf(osArchUnsupportedByTerraform)
}

func IsOsArchUnsupportedByTerraform(err error) bool {
	return err.Error() == osArchUnsupportedByTerraform
}
