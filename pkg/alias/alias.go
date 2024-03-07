package alias

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/tfversion/tfversion/pkg/download"
	"github.com/tfversion/tfversion/pkg/helpers"
)

// AliasVersion creates a symlink to the specified Terraform version.
func AliasVersion(alias string, version string) {
	if !download.IsAlreadyDownloaded(version) {
		err := fmt.Errorf("terraform version %s not found, run %s to install", helpers.ColoredVersion(version), helpers.ColoredInstallHelper(version))
		helpers.ExitWithError("aliasing", err)
	}

	aliasLocation := GetAliasLocation()

	// delete existing alias symlink, we consider it non-destructive anyways since you can easily restore it
	aliasPath := filepath.Join(aliasLocation, alias)
	_, err := os.Lstat(aliasPath)
	if err == nil {
		err = os.RemoveAll(aliasPath)
		if err != nil {
			helpers.ExitWithError("removing symlink", err)
		}
	}

	// create the symlink
	binaryVersionPath := download.GetInstallLocation(version)
	err = os.Symlink(binaryVersionPath, aliasPath)
	if err != nil {
		helpers.ExitWithError("creating symlink", err)
	}

	fmt.Printf("Aliased Terraform version %s as %s\n", helpers.ColoredVersion(version), helpers.ColoredVersion(alias))
}

// GetAliasLocation returns the directory where tfversion stores the aliases.
func GetAliasLocation() string {
	user, err := os.UserHomeDir()
	if err != nil {
		helpers.ExitWithError("getting user home directory", err)
	}

	aliasLocation := filepath.Join(user, download.ApplicationDir, download.AliasesDir)
	if _, err := os.Stat(aliasLocation); os.IsNotExist(err) {
		err := os.MkdirAll(aliasLocation, 0755)
		if err != nil {
			helpers.ExitWithError("creating alias directory", err)
		}
	}

	return aliasLocation
}

// IsAlias checks if the given alias is valid.
func IsAlias(alias string) bool {
	aliasPath := filepath.Join(GetAliasLocation(), alias)
	_, err := os.Stat(aliasPath)
	return !os.IsNotExist(err)
}

// GetVersion returns the Terraform version for the given alias.
func GetVersion(alias string) string {
	aliasLocation := GetAliasLocation()
	resolvePath, _ := filepath.EvalSymlinks(filepath.Join(aliasLocation, alias))
	_, targetVersion := filepath.Split(resolvePath)
	return targetVersion
}
