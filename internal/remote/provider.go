package remote

import "github.com/cbuschka/tfvm/internal/version"

// Provider is a provider of terraform release meta data.
type Provider interface {
	ListTerraformReleases() ([]*version.TerraformVersion, error)
	ListTerraformBuilds(release *version.TerraformVersion) ([]*TerraformBuild, error)
}

// GetDefaultProvider retrieves the default hashicorp web site provider.
func GetDefaultProvider() Provider {
	return Provider(&hashicorpRemote{})
}
