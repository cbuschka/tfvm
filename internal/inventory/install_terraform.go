package inventory

import (
	"archive/zip"
	"fmt"
	"github.com/cbuschka/tfvm/internal/remote"
	"github.com/cbuschka/tfvm/internal/util"
	"github.com/cbuschka/tfvm/internal/version"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// InstallTerraform installs a terraform version by downloading and extracting it.
func (inventory *Inventory) InstallTerraform(tfRelease *version.TerraformVersion) error {

	installed, err := inventory.IsTerraformInstalled(tfRelease)
	if err != nil {
		return err
	}

	if installed {
		return nil
	}

	tmpfile, err := ioutil.TempFile("", "tfvm_terraform*.zip")
	if err != nil {
		return err
	}
	defer os.Remove(tmpfile.Name())

	util.Print("Downloading terraform %s...", tfRelease.String())
	url := remote.GetURL(tfRelease)
	err = downloadFile(url, tmpfile.Name())
	if err != nil {
		util.Die(1, "Download failed: %s", err.Error())
		return err
	}

	util.Print("Installing terraform %s...", tfRelease.String())
	basePath, err := inventory.GetTerraformBasePath(tfRelease)
	if err != nil {
		return err
	}

	_, err = unzip(tmpfile.Name(), basePath)
	if err != nil {
		util.Die(1, "Unzipping failed: %s", err.Error())
		return err
	}

	util.Print("Terraform %s installed.", tfRelease.String())

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
