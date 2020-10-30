package version

import (
	"fmt"
	goversion "github.com/hashicorp/go-version"
)

// TerraformVersionSpec specifies a terraform version with a semver expression.
type TerraformVersionSpec struct {
	text        string
	constraints goversion.Constraints
}

// String gives string representation of version spec.
func (spec *TerraformVersionSpec) String() string {
	return spec.text
}

// ParseTerraformVersionSpec parses a version spec string to an TerraformVersionSpec.
func ParseTerraformVersionSpec(versionSpec string) (*TerraformVersionSpec, error) {

	if versionSpec == "latest" {
		return &TerraformVersionSpec{text: versionSpec, constraints: nil}, nil
	}

	constraints, err := goversion.NewConstraint(versionSpec)
	if err != nil {
		return nil, fmt.Errorf("Invalid version spec: '%s' (%s)", versionSpec, err.Error())
	}

	return &TerraformVersionSpec{text: versionSpec, constraints: constraints}, nil
}

// Matches checks if a terraform version is matched by a version spec.
func (spec *TerraformVersionSpec) Matches(tfRelease *TerraformVersion, latestTfRelease *TerraformVersion) bool {
	if spec.text == "latest" {
		return tfRelease.Version.String() == latestTfRelease.Version.String()
	}

	return spec.constraints.Check(tfRelease.Version)
}
