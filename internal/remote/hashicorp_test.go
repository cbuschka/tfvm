package remote

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestExtractReleases(t *testing.T) {

	html := "foo <a>terraform_1.0.0</a> bar\n<a>terraform_2.0.0-abc </a> stuff"

	releases, err := extractReleases(html)
	if err != nil {
		t.Errorf("Cannot extract.")
	}

	assert.Equal(t, 2, len(releases))
	assert.Equal(t, "1.0.0", releases[0].String())
	assert.Equal(t, "https://releases.hashicorp.com/terraform/1.0.0/terraform_1.0.0_linux_amd64.zip", GetUrl(releases[0]))
	assert.Equal(t, "2.0.0-abc", releases[1].String())
	assert.Equal(t, "https://releases.hashicorp.com/terraform/2.0.0-abc/terraform_2.0.0-abc_linux_amd64.zip", GetUrl(releases[1]))
}
