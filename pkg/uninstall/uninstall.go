package uninstall

import (
	"fmt"
	"os"

	"github.com/tfversion/tfversion/pkg/download"
	"github.com/tfversion/tfversion/pkg/helpers"
)

func Uninstall(version string) {
	if !download.IsAlreadyDownloaded(version) {
		err := fmt.Errorf("terraform version %s is not installed", helpers.ColoredVersion(version))
		helpers.ExitWithError("uninstalling", err)
	}

	installLocation := download.GetInstallLocation(version)
	err := os.RemoveAll(installLocation)
	if err != nil {
		helpers.ExitWithError("deleting Terraform version", err)
	}
}
