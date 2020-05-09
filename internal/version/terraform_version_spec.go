package version

import (
	"fmt"
	goversion "github.com/hashicorp/go-version"
)

type TerraformVersionSpec struct {
	text        string
	constraints goversion.Constraints
}

func (spec *TerraformVersionSpec) String() string {
	return spec.text
}

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

func (spec *TerraformVersionSpec) Matches(tfRelease *TerraformVersion, latestTfRelease *TerraformVersion) bool {
	if spec.text == "latest" {
		return tfRelease.Version.String() == latestTfRelease.Version.String()
	}

	return spec.constraints.Check(tfRelease.Version)
}
