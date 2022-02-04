package workspace

import (
	"github.com/cbuschka/tfvm/internal/util"
	"github.com/mitchellh/go-homedir"
	"gopkg.in/ini.v1"
	"path/filepath"
)

// Config is the tfvm config.
type Config struct {
	Sources []string
	Verbose bool `ini:"verbose"`
}

func loadConfig(workspaceRootDir string) (*Config, error) {

	homeDir, err := homedir.Dir()
	if err != nil {
		return nil, err
	}

	globalTfvmRc := "/etc/tfvmrc"
	userSpecificTfvmRc := filepath.Join(homeDir, ".tfvmrc")
	workspaceTfvmRc := filepath.Join(workspaceRootDir, ".tfvmrc")
	sources := []string{globalTfvmRc, userSpecificTfvmRc, workspaceTfvmRc}

	config := Config{Verbose: false, Sources: []string{"defaults"}}
	for _, source := range sources {
		isFile, err := util.IsFile(source)
		if err != nil {
			return nil, err
		}

		if isFile {
			iniConfig, err := ini.Load(source)
			if err != nil {
				return nil, err
			}

			err = iniConfig.MapTo(&config)
			if err != nil {
				return nil, err
			}

			config.Sources = append(config.Sources, source)
		}
	}

	return &config, nil
}
