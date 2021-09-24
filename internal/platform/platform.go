package platform

import (
	"github.com/cbuschka/tfvm/internal/util"
	"runtime"
)

// Platform is terraform supported combination of OS and processor architecture.
type Platform struct {
	Os   string
	Arch string
}

// GetPrimaryPlatform retrieves the primary platform.
func GetPrimaryPlatform() Platform {
	return GetSupportedPlatforms()[0]
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
