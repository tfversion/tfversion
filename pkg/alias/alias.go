package alias

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/tfversion/tfversion/pkg/helpers"
	"github.com/tfversion/tfversion/pkg/paths"
)

// AliasVersion creates a symlink to the specified Terraform version.
func AliasVersion(alias string, version string) {
	if !paths.IsAlreadyDownloaded(version) {
		err := fmt.Errorf("terraform version %s not found, run %s to install", helpers.ColoredVersion(version), helpers.ColoredInstallHelper(version))
		helpers.ExitWithError("aliasing", err)
	}

	aliasLocation := GetAliasLocation()
	aliasPath := filepath.Join(aliasLocation, alias)

	// ensure the symlink target is available
	err := paths.RemoveSymlink(aliasPath)
	if err != nil {
		helpers.ExitWithError("removing symlink", err)
	}

	// create the symlink
	binaryVersionPath := paths.GetInstallLocation(version)
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

	aliasLocation := filepath.Join(user, paths.ApplicationDir, paths.AliasesDir)
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
	_, err := os.Lstat(aliasPath)
	return !os.IsNotExist(err)
}

// GetVersion returns the Terraform version for the given alias.
func GetVersion(alias string) string {
	aliasLocation := GetAliasLocation()
	resolvePath, err := filepath.EvalSymlinks(filepath.Join(aliasLocation, alias))
	if err != nil {
		helpers.ExitWithError("resolving symlink", err)
	}
	_, targetVersion := filepath.Split(resolvePath)
	return targetVersion
}
