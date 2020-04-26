package tfvm

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestExtractReleases(t *testing.T) {

	html := "foo <a>terraform_1.0</a> bar\n<a>terraform_2.0abc </a> stuff"

	releases, err := extractReleases(html)
	if err != nil {
		t.Errorf("Cannot extract.")
	}

	assert.Equal(t, 2, len(releases))
	assert.Equal(t, "1.0", releases[0].version)
	assert.Equal(t, "https://releases.hashicorp.com/terraform/1.0/terraform_1.0_linux_amd64.zip", releases[0].url)
	assert.Equal(t, "2.0abc", releases[1].version)
	assert.Equal(t, "https://releases.hashicorp.com/terraform/2.0abc/terraform_2.0abc_linux_amd64.zip", releases[1].url)
}
