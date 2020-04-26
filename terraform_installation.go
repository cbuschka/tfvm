package tfvm

import (
	"archive/zip"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func (inventory *Inventory) InstallTerraform(version string) error {

	installed, err := inventory.IsTerraformInstalled(version)
	if err != nil {
		return err
	}

	if installed {
		return nil
	}

	tfRelease := GetTerraformRelease(version)

	tmpfile, err := ioutil.TempFile("", "tfvm_terraform*.zip")
	if err != nil {
		return err
	}

	fmt.Printf("Downloading terraform %s...\n", version)
	err = downloadFile(tfRelease.url, tmpfile.Name())
	if err != nil {
		return err
	}

	fmt.Printf("Installing terraform %s...\n", version)
	basePath, err := inventory.GetTerraformBasePath(version)
	if err != nil {
		return err
	}

	_, err = unzip(tmpfile.Name(), basePath)
	if err != nil {
		return err
	}

	fmt.Printf("Terraform %s installed.\n", version)

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

		// Check for ZipSlip. More Info: http://bit.ly/2MsjAWE
		if !strings.HasPrefix(fpath, filepath.Clean(dest)+string(os.PathSeparator)) {
			return filenames, fmt.Errorf("%s: illegal file path", fpath)
		}

		filenames = append(filenames, fpath)

		if f.FileInfo().IsDir() {
			// Make Folder
			os.MkdirAll(fpath, os.ModePerm)
			continue
		}

		// Make File
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

		// Close the file without defer to close before next iteration of loop
		outFile.Close()
		rc.Close()

		if err != nil {
			return filenames, err
		}
	}
	return filenames, nil
}
