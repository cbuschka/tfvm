package inventory

import (
	"bufio"
	statePkg "github.com/cbuschka/tfvm/internal/inventory/state"
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

// Save saves the inventory state into the state.json file.
func (inventory *Inventory) Save() error {

	stateFilePath, err := inventory.getStateFilePath()
	if err != nil {
		return err
	}

	if err := os.MkdirAll(filepath.Dir(stateFilePath), os.ModePerm); err != nil {
		return err
	}
	if err != nil {
		return err
	}

	file, err := os.OpenFile(stateFilePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	defer file.Close()
	if err != nil {
		return err
	}
	w := bufio.NewWriter(file)
	_, err = inventory.WriteTo(w)
	return err
}
