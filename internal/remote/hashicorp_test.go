package remote

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"runtime"
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
	expectedURL := fmt.Sprintf("https://releases.hashicorp.com/terraform/1.0.0/terraform_1.0.0_%s_%s.zip", runtime.GOOS, runtime.GOARCH);
	assert.Equal(t, expectedURL, GetURL(releases[0]))

	assert.Equal(t, "2.0.0-abc", releases[1].String())
	expectedURL = fmt.Sprintf("https://releases.hashicorp.com/terraform/2.0.0-abc/terraform_2.0.0-abc_%s_%s.zip", runtime.GOOS, runtime.GOARCH);
	assert.Equal(t, expectedURL, GetURL(releases[1]))
}
