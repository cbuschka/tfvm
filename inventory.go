package tfvm

import (
	"os"
	"os/user"
	"path/filepath"
)

type Inventory struct {}

func GetInventory() (*Inventory, error) {
	inventory := Inventory{}
	err := inventory.initInventory()
	if err != nil {
		return nil, err
	}

	return &inventory, nil
}

func (this *Inventory) GetTerraform(version string) (*Terraform, error) {
	tfPath, err := getTerraformPath(version)
	if err != nil {
		return nil, err
	}

	if _, err := os.Stat(tfPath); err != nil {
		return nil, err
	}

	return &Terraform{version: version, path: tfPath}, nil
}

func getTerraformPath(version string) (string, error) {
	inventoryDir, err := getInventoryDir()
	if err != nil {
		return "", err
	}

	versionedTfPath := filepath.Join(inventoryDir, "installed", version, "terraform")
	return versionedTfPath, nil
}

func (this *Inventory) IsTerraformInstalled(version string) (bool, error) {
	versionedTfPath, err := getTerraformPath(version)
	if err != nil {
		return false, err
	}

	if _, err := os.Stat(versionedTfPath); err != nil {
		return false, err
	}

	return true, nil
}

func (this *Inventory) initInventory() error {
	inventoryDir, err := getInventoryDir()
	if err != nil {
		return err
	}

	err = os.MkdirAll(inventoryDir, 0755)
	return err
}

func getInventoryDir() (string, error) {
	homeDir, err := getUserHomeDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(homeDir, ".tfvm"), nil
}

func getUserHomeDir() (string, error) {
	usr, err := user.Current()
	if err != nil {
		return "", err
	}

	return usr.HomeDir, nil
}
