package remote

import (
	"fmt"
	platformPkg "github.com/cbuschka/tfvm/internal/platform"
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
	expectedURL := fmt.Sprintf("https://releases.hashicorp.com/terraform/1.0.0/terraform_1.0.0_%s_%s.zip", "myos", "myarch")
	assert.Equal(t, expectedURL, GetURL(releases[0], platformPkg.Platform{Os: "myos", Arch: "myarch"}))

	assert.Equal(t, "2.0.0-abc", releases[1].String())
	expectedURL = fmt.Sprintf("https://releases.hashicorp.com/terraform/2.0.0-abc/terraform_2.0.0-abc_%s_%s.zip", "myos", "myarch")
	assert.Equal(t, expectedURL, GetURL(releases[1], platformPkg.Platform{Os: "myos", Arch: "myarch"}))
}

func TestExtractChecksums(t *testing.T) {

	contents := "80ae021d6143c7f7cbf4571f65595d154561a2a25fd934b7a8ccc1ebf3014b9b  terraform_1.0.7_darwin_amd64.zip\ncbab9aca5bc4e604565697355eed185bb699733811374761b92000cc188a7725  terraform_1.0.7_darwin_arm64.zip\nb5dd00d3e8a072da806e2223a49a9f4292e5c978c25ec5b0dc45d7e38904c43b  terraform_1.0.7_freebsd_386.zip\n330f815036288ee1f135e26d2dff4f262429f8039fc9945f9e081ab1b355daf3  terraform_1.0.7_freebsd_amd64.zip\n76e1d92c0580aef177618552e4bd3103ed351c876a8c157cb4737b29297524c4  terraform_1.0.7_freebsd_arm.zip\n1e171065bde9ce7758f75b3740674b42db1ff1f6ffc0692d9905813a9c381263  terraform_1.0.7_linux_386.zip\nbc79e47649e2529049a356f9e60e06b47462bf6743534a10a4c16594f443be7b  terraform_1.0.7_linux_amd64.zip\n4e71a9e759578020750be41e945c086e387affb58568db6d259d80d123ac80d3  terraform_1.0.7_linux_arm64.zip\nbe8d9de8c34e3e843ff8cbc713b41ce0c2bd97b0020c80e934443976b4429ae2  terraform_1.0.7_linux_arm.zip\nc813de2db415f5df576366f9dbbaeb62f0bea56d5b5ad0b9e47f469a18cb28a4  terraform_1.0.7_openbsd_386.zip\n7b97be6c3eb56fe0f0f1eaa8104dd24308cb86785c8404c2df566d3907dcc634  terraform_1.0.7_openbsd_amd64.zip\n257fb6e57c2e3353445b52b5ab9cb7caf5e1c6680cbe8f158300705ee7a6e215  terraform_1.0.7_solaris_amd64.zip\n64b930ce06f85c214d1830a213a7528a5d988bd09779729d5b33956b962ec708  terraform_1.0.7_windows_386.zip\n88b8a4ae66367d9662a75599421e0bccbb70e6af92d64c62e91aff883c9eccc6  terraform_1.0.7_windows_amd64.zip"

	releases, err := extractChecksums(contents)
	if err != nil {
		t.Errorf("Cannot extract.")
	}

	assert.Equal(t, 14, len(releases))
}
