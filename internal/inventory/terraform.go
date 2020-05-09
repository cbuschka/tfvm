package inventory

// A version of terraform installed into the local inventory
type Terraform struct {
	version string
	path    string
}

func newTerraform(version string, path string) *Terraform {
	return &Terraform{version: version, path: path}
}
