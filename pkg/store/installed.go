package store

import (
	"os"
	"path/filepath"

	"tfversion/pkg/helpers"
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
