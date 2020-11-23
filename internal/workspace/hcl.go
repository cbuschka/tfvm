package workspace

import (
	"fmt"
	"github.com/cbuschka/tfvm/internal/util"
	"github.com/cbuschka/tfvm/internal/version"
	"github.com/hashicorp/hcl"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// TerraformHCLObject is the struct for a terraform block in hcl.
type TerraformHCLObject struct {
	RequiredVersion string `hcl:"required_version"`
}

// RootHCLObject is the struct for tf file in hcl.
type RootHCLObject struct {
	Terraform TerraformHCLObject `hcl:"terraform"`
}

func (workspace *Workspace) findTerraformVersionSpecFromHclFiles() (*TerraformVersionSelection, error) {

	hclFiles, err := collectHclFiles(workspace.RootDir)
	if err != nil {
		return nil, err
	}

	for _, hclFile := range *hclFiles {
		fmt.Printf("Trying to get tf version from %s\n", hclFile)

		hclFileBytes, err := ioutil.ReadFile(hclFile)
		if err != nil {
			return nil, err
		}
		hclFileContent, err := hcl.ParseBytes(hclFileBytes)
		if err != nil {
			return nil, err
		}g

		root := RootHCLObject{}
		err = hcl.DecodeObject(&root, hclFileContent)
		if err != nil {
			return nil, err
		}

		if root.Terraform.RequiredVersion != "" {
			spec, err := version.ParseTerraformVersionSpec(root.Terraform.RequiredVersion)
			if err != nil {
				return nil, err
			}
			return &TerraformVersionSelection{versionSpec: spec, sourceName: hclFile, sourceType: HclFile}, nil
		}

	}

	return nil, nil
}

func collectHclFiles(startDir string) (*[]string, error) {

	var hclFiles []string

	err := filepath.Walk(startDir, func(path string, fileInfo os.FileInfo, err error) error {

		if err != nil {
			return nil
		}

		if !fileInfo.IsDir() && strings.HasSuffix(fileInfo.Name(), ".tf") {
			util.Debug("File seen: %s", path)
			hclFiles = append(hclFiles, path)
		}

		return err
	})

	if err != nil {
		return nil, err
	}

	return &hclFiles, nil
}
