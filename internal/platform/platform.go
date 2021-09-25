package platform

import (
	"fmt"
	"github.com/cbuschka/tfvm/internal/util"
	"runtime"
)

// Platform is terraform supported combination of OS and processor architecture.
type Platform struct {
	Os   string
	Arch string
}

// GetSupportedPlatforms retrieves all platforms binaries can be executed of.
func GetSupportedPlatforms() []Platform {

	os := getOs()
	arch := getArch()

	platforms := []Platform{{Os: os, Arch: arch}}
	if os == "darwin" && arch == "arm64" {
		platforms = append(platforms, Platform{Os: os, Arch: "amd64"})
	}

	return platforms
}

func (platform Platform) String() string {
	return fmt.Sprintf("%s/%s", platform.Os, platform.Arch)
}

func getArch() string {
	tfArch := util.GetFirstEnv("TFVM_TERRAFORM_ARCH", "TERRAFORM_ARCH")
	if tfArch == "" {
		tfArch = runtime.GOARCH
	}
	return tfArch
}

func getOs() string {
	tfOs := util.GetFirstEnv("TFVM_TERRAFORM_OS", "TERRAFORM_OS")
	if tfOs == "" {
		tfOs = runtime.GOOS
	}
	return tfOs
}
