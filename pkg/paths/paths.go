package paths

import (
	"os"
	"path/filepath"

	"github.com/tfversion/tfversion/pkg/helpers"
)

// ApplicationDir is the name of the directory where tfversion stores its data.
func GetApplicationLocation() string {
	user, err := os.UserHomeDir()
	if err != nil {
		helpers.ExitWithError("error getting user home directory", err)
	}
	applicationLocation := filepath.Join(user, ApplicationDir)
	err = EnsureDirExists(applicationLocation)
	if err != nil {
		helpers.ExitWithError("creating application directory", err)
	}
	return applicationLocation
}

// GetDownloadLocation returns the directory where tfversion downloads Terraform releases to.
func GetDownloadLocation() string {
	downloadLocation := filepath.Join(GetApplicationLocation(), VersionsDir)
	err := EnsureDirExists(downloadLocation)
	if err != nil {
		helpers.ExitWithError("creating download directory", err)
	}
	return downloadLocation
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

// GetInstallLocation returns the directory where a specific Terraform version is installed to.
func GetInstallLocation(version string) string {
	return filepath.Join(GetDownloadLocation(), version)
}

// GetBinaryLocation returns the path to the Terraform binary for the given version.
func GetBinaryLocation(version string) string {
	return filepath.Join(GetInstallLocation(version), TerraformBinaryName)
}

// GetActiveBinaryLocation returns the path to the currently active Terraform binary.
func GetActiveBinaryLocation() string {
	return filepath.Join(GetUseLocation(), TerraformBinaryName)
}

// GetInstalledVersions returns a list of all installed Terraform versions.
func GetInstalledVersions() []string {
	installLocation := GetDownloadLocation()
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

// IsAlreadyDownloaded checks if the given Terraform version is already downloaded and unzipped.
func IsAlreadyDownloaded(version string) bool {
	binaryPath := GetBinaryLocation(version)
	_, err := os.Stat(binaryPath)
	return !os.IsNotExist(err)
}

// EnsureDirExists checks if the given directory exists, and creates it if it doesn't.
func EnsureDirExists(path string) error {
	_, err := os.Stat(path)
	if !os.IsNotExist(err) {
		return err
	}
	err = os.MkdirAll(path, 0755)
	return err
}
