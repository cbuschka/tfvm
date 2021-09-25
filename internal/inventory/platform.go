package inventory

import "github.com/cbuschka/tfvm/internal/platform"

// GetPrimaryPlatform retrieves the primary platform.
func (inventory *Inventory) GetPrimaryPlatform() platform.Platform {
	return inventory.platforms[0]
}
