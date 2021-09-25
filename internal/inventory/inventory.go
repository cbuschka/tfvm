package inventory

import (
	statePkg "github.com/cbuschka/tfvm/internal/inventory/state"
	"github.com/cbuschka/tfvm/internal/log"
	"github.com/cbuschka/tfvm/internal/platform"
	"github.com/cbuschka/tfvm/internal/remote"
	"time"
)

// Inventory represents the tfvm inventory on disk. An inventory stores
// terraform versions and a statePkg.json inventory statePkg file.
type Inventory struct {
	LastUpdateTime    time.Time
	TerraformReleases map[string]*statePkg.TerraformReleaseState
	cacheDir          string
	remoteProvider    remote.Provider
	platforms         []platform.Platform
}

// GetInventory initializes an inventory instance representing the
// inventory of the current machine.
func GetInventory() (*Inventory, error) {
	inventory, err := NewInventory()
	if err != nil {
		return nil, err
	}

	err = inventory.Load()
	if err != nil {
		return nil, err
	}

	return inventory, nil
}

// NewInventory creates a new empty inventory instance.
func NewInventory() (*Inventory, error) {
	inventoryDir, err := getInventoryDir()
	if err != nil {
		return nil, err
	}

	log.Debugf("Inventory dir: '%s'", inventoryDir)

	inventory := Inventory{LastUpdateTime: time.Unix(0, 0), TerraformReleases: map[string]*statePkg.TerraformReleaseState{}, cacheDir: inventoryDir}
	if err != nil {
		return nil, err
	}

	defaultState, err := statePkg.GetDefaultState()
	if err != nil {
		return nil, err
	}

	if log.IsTraceEnabled() {
		stateJSON, err := defaultState.ToJSON()
		if err != nil {
			log.Tracef("state.ToJSON() failed: %v", err)
		} else {
			log.Tracef("Default state:\n%s", stateJSON)
		}
	}

	inventory.mergeInState(defaultState)

	inventory.remoteProvider = remote.GetDefaultProvider()
	inventory.platforms = platform.GetSupportedPlatforms()
	log.Infof("Supported platforms: %v", inventory.platforms)

	return &inventory, err
}
