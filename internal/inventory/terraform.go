package inventory

// Terraform is a version of terraform installed into the local inventory.
type Terraform struct {
	version string
	os      string
	arch    string
	path    string
}

func newTerraform(version string, os string, arch string, path string) *Terraform {
	return &Terraform{version: version, os: os, arch: arch, path: path}
}
