package uninstall

import (
	"fmt"
	"os"

	"github.com/tfversion/tfversion/pkg/helpers"
	"github.com/tfversion/tfversion/pkg/paths"
)

func Uninstall(version string) {
	if !paths.IsAlreadyDownloaded(version) {
		err := fmt.Errorf("terraform version %s is not installed", helpers.ColoredVersion(version))
		helpers.ExitWithError("uninstalling", err)
	}

	installLocation := paths.GetInstallLocation(version)
	err := os.RemoveAll(installLocation)
	if err != nil {
		helpers.ExitWithError("deleting Terraform version", err)
	}
}
