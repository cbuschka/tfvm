package remote

import "github.com/cbuschka/tfvm/internal/version"

type Provider interface {
	ListTerraformReleases() ([]*version.TerraformVersion, error)
	ListTerraformBuilds(release *version.TerraformVersion) ([]*TerraformBuild, error)
}

func GetDefaultProvider() Provider {
	return Provider(&HashicorpRemote{})
}
