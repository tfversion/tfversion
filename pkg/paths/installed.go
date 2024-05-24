package paths

import (
	"os"
	"path/filepath"

	"github.com/tfversion/tfversion/pkg/helpers"
)

// GetInstallLocation returns the directory where tfversion downloads Terraform releases to.
func GetInstallLocation() string {
	downloadLocation := filepath.Join(GetApplicationLocation(), VersionsDir)
	err := EnsureDirExists(downloadLocation)
	if err != nil {
		helpers.ExitWithError("creating download directory", err)
	}
	return downloadLocation
}

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

// IsInstalled checks if the given Terraform version is already installed.
func IsInstalled(version string) bool {
	binaryPath := GetBinaryLocation(version)
	_, err := os.Stat(binaryPath)
	return !os.IsNotExist(err)
}

// GetInstalledVersionLocation returns the directory where a specific Terraform version is installed to.
func GetInstalledVersionLocation(version string) string {
	return filepath.Join(GetInstallLocation(), version)
}

// GetBinaryLocation returns the path to the Terraform binary for the given version.
func GetBinaryLocation(version string) string {
	return filepath.Join(GetInstalledVersionLocation(version), TerraformBinaryName)
}

// GetActiveBinaryLocation returns the path to the currently active Terraform binary.
func GetActiveBinaryLocation() string {
	return filepath.Join(GetUseLocation(), TerraformBinaryName)
}

// GetUseLocation returns the directory where tfversion stores the symlink to the currently used Terraform version.
func GetUseLocation() string {
	useLocation := filepath.Join(GetApplicationLocation(), UseDir)
	err := EnsureDirExists(useLocation)
	if err != nil {
		helpers.ExitWithError("creating use directory", err)
	}
	return useLocation
}
