package inventory

import (
	state "github.com/cbuschka/tfvm/internal/inventory/state"
	"github.com/cbuschka/tfvm/internal/remote"
	"time"
)

// Inventory represents the tfvm inventory on disk. An inventory stores
// terraform versions and a state.json inventory state file.
type Inventory struct {
	LastUpdateTime    time.Time
	TerraformReleases map[string]*state.TerraformReleaseState
	cacheDir          string
	remoteProvider    remote.Provider
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

	inventory := Inventory{LastUpdateTime: time.Unix(0, 0), TerraformReleases: map[string]*state.TerraformReleaseState{}, cacheDir: inventoryDir}
	if err != nil {
		return nil, err
	}

	defaultState, err := state.GetDefaultState()
	if err != nil {
		return nil, err
	}
	inventory.mergeInState(defaultState)

	inventory.remoteProvider = remote.GetDefaultProvider()

	return &inventory, err
}
