package paths

import (
	"os"
	"path/filepath"

	"github.com/tfversion/tfversion/pkg/helpers"
)

// GetUseLocation returns the directory where tfversion stores the symlink to the currently used Terraform version.
func GetUseLocation() string {
	useLocation := filepath.Join(GetApplicationLocation(), UseDir)
	err := EnsureDirExists(useLocation)
	if err != nil {
		helpers.ExitWithError("creating use directory", err)
	}
	return useLocation
}

// GetActiveBinaryLocation returns the path to the currently active Terraform binary.
func GetActiveBinaryLocation() string {
	return filepath.Join(GetUseLocation(), TerraformBinaryName)
}

// GetUseVersion returns the currently active Terraform version.
func GetUseVersion() string {
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
