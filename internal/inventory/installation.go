package inventory

import (
	"archive/zip"
	"crypto/sha256"
	"fmt"
	"github.com/cbuschka/tfvm/internal/inventory/state"
	"github.com/cbuschka/tfvm/internal/log"
	platformPkg "github.com/cbuschka/tfvm/internal/platform"
	"github.com/cbuschka/tfvm/internal/remote"
	"github.com/cbuschka/tfvm/internal/util"
	"github.com/cbuschka/tfvm/internal/version"
	"io"
	"io/ioutil"
	"math/big"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// GetInstalledTerraform installs a terraform version by downloading and extracting it. GetInstalledTerraform respects
// the platform and fallbacks.
func (inventory *Inventory) GetInstalledTerraform(terraformVersion *version.TerraformVersion) (*Terraform, error) {

	log.Debugf("Request for installed terraform %s", terraformVersion.String())

	for _, platform := range inventory.platforms {
		terraform, err := inventory.InstallTerraform(terraformVersion, platform)
		if err != nil {
			if !IsNoSuchTerraformReleaseBuild(err) {
				return nil, err
			}
		} else {
			return terraform, nil
		}
	}

	return nil, NewOsArchUnsupportedByTerraform()
}

// InstallTerraform installs a terraform version for an particular os and arch.
func (inventory *Inventory) InstallTerraform(terraformVersion *version.TerraformVersion, platform platformPkg.Platform) (*Terraform, error) {

	log.Debugf("Request to install terraform %s on %s", terraformVersion.String(), platform)

	installed, err := inventory.IsTerraformInstalled(terraformVersion, platform)
	if err != nil {
		return nil, err
	}

	if !installed {
		tfRelease, err := inventory.GetTerraformRelease(terraformVersion)
		if err != nil {
			util.Die(1, "Finding terraform release failed: %s", err.Error())
			return nil, err
		}

		if tfRelease.Builds == nil || len(tfRelease.Builds) == 0 {
			err = inventory.UpdateTerraformRelease(tfRelease)
			if err != nil {
				util.Die(1, "Updating release failed: %s", err.Error())
				return nil, err
			}
		}

		tfReleaseBuild, err := inventory.getBuild(tfRelease, platform)
		if err != nil {
			return nil, err
		}

		tmpfile, err := ioutil.TempFile("", "tfvm_terraform*.zip")
		if err != nil {
			return nil, err
		}
		defer os.Remove(tmpfile.Name())

		util.Print("Downloading terraform %s for %s/%s...", terraformVersion.String(), tfReleaseBuild.Os, tfReleaseBuild.Arch)
		url := remote.GetURL(terraformVersion, platform)
		err = downloadFile(url, tmpfile.Name())
		if err != nil {
			util.Die(1, "Download failed: %s", err.Error())
			return nil, err
		}

		log.Debugf("Checksum for terraform %s for %s/%s is %s.", terraformVersion.String(), tfReleaseBuild.Os, tfReleaseBuild.Arch, tfReleaseBuild.SHA256Checksum)
		if tfReleaseBuild.SHA256Checksum == "n/a" || tfReleaseBuild.SHA256Checksum == "" {
			util.Print("Cannot verify downloaded terraform version %s for %s.", terraformVersion.String(), platform)
		} else {
			sha256Checksum := new(big.Int)
			sha256Checksum, _ = sha256Checksum.SetString(tfReleaseBuild.SHA256Checksum, 16)
			err := verifyFile(tmpfile.Name(), sha256Checksum)
			if err != nil {
				return nil, err
			}
		}

		util.Print("Installing terraform %s for %s...", terraformVersion.String(), platform)
		basePath, err := inventory.GetTerraformBasePath(terraformVersion, platform)
		if err != nil {
			return nil, err
		}

		_, err = unzip(tmpfile.Name(), basePath)
		if err != nil {
			util.Die(1, "Unzipping failed: %s", err.Error())
			return nil, err
		}

		util.Print("Terraform %s for %s installed.", terraformVersion.String(), platform)
	}

	terraform, err := inventory.GetTerraform(terraformVersion, platform)
	if err != nil {
		return nil, err
	}

	return terraform, nil
}

func verifyFile(name string, expectedSHA256Checksum *big.Int) error {

	log.Debugf("Verifying file %s...", name)

	f, err := os.Open(name)
	if err != nil {
		return err
	}
	defer f.Close()

	hasher := sha256.New()
	if _, err := io.Copy(hasher, f); err != nil {
		return err
	}

	calculatedChecksumBytes := hasher.Sum(nil)
	calculatedChecksum := new(big.Int)
	calculatedChecksum = calculatedChecksum.SetBytes(calculatedChecksumBytes)

	if calculatedChecksum.Cmp(expectedSHA256Checksum) != 0 {
		return fmt.Errorf("downloaded terraform zip has not expected checksum")
	}

	log.Infof("Successfully verified file %s, calculated checksum: %s, expected checksum: %s...", name, calculatedChecksum.Text(16), expectedSHA256Checksum.Text(16))

	return nil
}

func downloadFile(url string, filepath string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}

func unzip(src string, dest string) ([]string, error) {

	var filenames []string

	r, err := zip.OpenReader(src)
	if err != nil {
		return filenames, err
	}
	defer r.Close()

	for _, f := range r.File {

		// Store filename/path for returning and using later on
		fpath := filepath.Join(dest, f.Name)

		if !strings.HasPrefix(fpath, filepath.Clean(dest)+string(os.PathSeparator)) {
			return filenames, fmt.Errorf("%s: illegal file path", fpath)
		}

		filenames = append(filenames, fpath)

		if f.FileInfo().IsDir() {
			os.MkdirAll(fpath, os.ModePerm)
			continue
		}

		if err = os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
			return filenames, err
		}

		outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			return filenames, err
		}

		rc, err := f.Open()
		if err != nil {
			return filenames, err
		}

		_, err = io.Copy(outFile, rc)

		outFile.Close()
		rc.Close()

		if err != nil {
			return filenames, err
		}
	}
	return filenames, nil
}

func (inventory *Inventory) getBuild(release *state.TerraformReleaseState, platform platformPkg.Platform) (*state.TerraformReleaseBuildState, error) {

	for _, build := range release.Builds {
		if build.Os == platform.Os && build.Arch == platform.Arch {
			return build, nil
		}
	}

	return nil, NewNoSuchTerraformReleaseBuild()
}
