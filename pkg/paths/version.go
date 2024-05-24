package paths

import (
	"os"
	"path/filepath"

	"github.com/tfversion/tfversion/pkg/helpers"
)

// GetInstalledVersions returns a list of all installed Terraform versions.
func GetInstalledVersions() []string {
	installLocation := GetInstallLocation()
	installedVersions, err := os.ReadDir(installLocation)
	if err != nil {
		helpers.ExitWithError("listing versions directory", err)
	}
	var versionNames []string
	for _, v := range installedVersions {
		versionNames = append(versionNames, v.Name())
	}
	return versionNames
}

// GetActiveVersion returns the currently active Terraform version.
func GetActiveVersion() string {
	binaryPath := GetActiveBinaryLocation()
	_, err := os.Lstat(binaryPath)
	if err != nil {
		helpers.ExitWithError("no current terraform version found", err)
	}
	realPath, err := filepath.EvalSymlinks(binaryPath)
	if err != nil {
		helpers.ExitWithError("resolving symlink", err)
	}
	_, version := filepath.Split(filepath.Dir(realPath))
	return version
}
