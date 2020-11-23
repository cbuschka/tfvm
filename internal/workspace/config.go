package workspace

import (
	"github.com/mitchellh/go-homedir"
	"gopkg.in/ini.v1"
	"path/filepath"
)

// Config is the tfvm config.
type Config struct {
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

	iniConfig, err := ini.LooseLoad(globalTfvmRc, userSpecificTfvmRc, workspaceTfvmRc)
	if err != nil {
		return nil, err
	}

	config := Config{Verbose: false}
	err = iniConfig.MapTo(&config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
