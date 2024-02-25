package uninstall

import (
	"fmt"
	"os"

	"github.com/tfversion/tfversion/pkg/download"
)

// DeleteVersion deletes the given Terraform version from directory that tfversion downloads Terraform releases to.
func DeleteVersionFromDownloadLocation(version string) error {
	installLocation := download.GetInstallLocation(version)
	//TODO: Check whether the path exists
	err := os.RemoveAll(installLocation)
	if err != nil {
		return fmt.Errorf("deleting Terraform version failed: %s", err)
	}
	return nil
}
