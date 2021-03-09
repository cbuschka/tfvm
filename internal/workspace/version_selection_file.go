package workspace

import (
	"bufio"
	"github.com/cbuschka/tfvm/internal/version"
	"io"
	"os"
	"strings"
)

func readVersionSelectionFromFile(versionSelectionFile string) (*TerraformVersionSelection, error) {
	fileReader, err := os.Open(versionSelectionFile)
	if err != nil {
		return nil, err
	}
	defer fileReader.Close()

	content, err := ReadSingleLineFrom(fileReader)
	if err != nil {
		return nil, err
	}

	versionSpec, err := version.ParseTerraformVersionSpec(content)
	if err != nil {
		return nil, err
	}

	return &TerraformVersionSelection{versionSpec: versionSpec, sourceName: versionSelectionFile, sourceType: File}, nil
}

// ReadSingleLineFrom reads a single line, in case of multiples the last one, with
// comment lines (starting with #) or empty ones being skipped
func ReadSingleLineFrom(fileReader io.Reader) (string, error) {

	content := ""
	scanner := bufio.NewScanner(fileReader)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if !strings.HasPrefix(line, "#") && line != "" {
			content = line
		}
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	return content, nil
}
