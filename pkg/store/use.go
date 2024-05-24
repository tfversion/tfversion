package store

import (
	"path/filepath"

	"tfversion/pkg/helpers"
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
