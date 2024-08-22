package uninstall

import (
	"fmt"

	"tfversion/internal/helpers"
	"tfversion/internal/store"
)

func Uninstall(version string) {
	if !store.IsInstalled(version) {
		err := fmt.Errorf("terraform version %s is not installed", helpers.ColoredVersion(version))
		helpers.ExitWithError("uninstalling", err)
	}

	installLocation := store.GetInstalledVersionLocation(version)
	err := store.RemoveDir(installLocation)
	if err != nil {
		helpers.ExitWithError("deleting Terraform version", err)
	}
}
