package inventory

import (
	"bufio"
	"github.com/cbuschka/tfvm/internal/atomio"
	statePkg "github.com/cbuschka/tfvm/internal/inventory/state"
	"github.com/cbuschka/tfvm/internal/log"
	"io"
	"os"
	"path/filepath"
)

// WriteTo writes the inventory state to writer w.
func (inventory *Inventory) WriteTo(w io.Writer) (int64, error) {
	terraformReleases := make([]*statePkg.TerraformReleaseState, 0, len(inventory.TerraformReleases))
	for _, terraformRelease := range inventory.TerraformReleases {
		terraformReleases = append(terraformReleases, terraformRelease)
	}
	state := statePkg.State{LastUpdateTime: inventory.LastUpdateTime, TerraformReleases: terraformReleases}
	data, err := state.Marshall()
	if err != nil {
		return -1, err
	}

	size, err := w.Write(data)
	return int64(size), err
}

// Save saves the inventory state into the state.json file in an atomic way.
func (inventory *Inventory) Save() error {

	log.Debug("Saving inventory...")

	stateFilePath, err := inventory.getStateFilePath()
	if err != nil {
		return err
	}

	log.Debugf("State file path: '%s'", stateFilePath)

	if err := os.MkdirAll(filepath.Dir(stateFilePath), os.ModePerm); err != nil {
		return err
	}
	if err != nil {
		return err
	}

	file, err := atomio.OpenTempFile(stateFilePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	defer atomio.CloseAndReplaceFile(file, stateFilePath)
	if err != nil {
		return err
	}
	w := bufio.NewWriter(file)
	_, err = inventory.WriteTo(w)
	if err != nil {
		return err
	}

	log.Info("Inventory successfully saved.")

	return nil
}
