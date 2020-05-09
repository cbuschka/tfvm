package tfvm

func RunTfvmWhichCommand(args []string) error {

	inventory, err := GetInventory()
	if err != nil {
		return err
	}

	err = inventory.Update()
	if err != nil {
		return err
	}

	config, err := GetConfiguration()
	if err != nil {
		if IsNoConfigExists(err) {
			Print("No terraform version configured.")
			return nil
		}

		return err
	}

	tfRelease, err := inventory.GetTerraformRelease(config.versionSpec)
	if err != nil {
		if IsNoSuchTerraformRelease(err) {
			Die(1, "Configured terraform version %s (%s) is not known.", config.versionSpec, config.file)
			return err
		}

		return err
	}

	Print("Configured terraform version is %s (%s).", tfRelease.Version, config.file)
	return nil
}
