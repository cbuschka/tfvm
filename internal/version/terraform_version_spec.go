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

// SafeParseTerraformVersionSpec parses a version spec string to an TerraformVersionSpec, or panics if fails.
func SafeParseTerraformVersionSpec(versionSpecStr string) *TerraformVersionSpec {
	versionSpec, err := ParseTerraformVersionSpec(versionSpecStr)
	if err != nil {
		panic(err)
	}

	return versionSpec
}

// ParseTerraformVersionSpec parses a version spec string to an TerraformVersionSpec.
func ParseTerraformVersionSpec(versionSpecStr string) (*TerraformVersionSpec, error) {

	if versionSpecStr == "latest" {
		return &TerraformVersionSpec{text: versionSpecStr, constraints: nil}, nil
	}

	constraints, err := goversion.NewConstraint(versionSpecStr)
	if err != nil {
		return nil, fmt.Errorf("Invalid version spec: '%s' (%s)", versionSpecStr, err.Error())
	}

	return &TerraformVersionSpec{text: versionSpecStr, constraints: constraints}, nil
}

// Matches checks if a terraform version is matched by a version spec.
func (spec *TerraformVersionSpec) Matches(tfRelease *TerraformVersion, latestTfRelease *TerraformVersion) bool {
	if spec.text == "latest" {
		return tfRelease.String() == latestTfRelease.String()
	}

	return spec.constraints.Check(tfRelease.Version())
}
