package tfvm

import (
	"errors"
)

const noSuchTerraformReleaseMsg = "no such terraform release"

func IsNoSuchTerraformRelease(err error) bool {
	return err.Error() == noSuchTerraformReleaseMsg
}

func newNoSuchTerraformRelease() error {
	return errors.New(noSuchTerraformReleaseMsg)
}

type TerraformRelease struct {
	Version string
}
