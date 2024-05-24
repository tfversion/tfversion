package paths

import (
	"os"
	"path/filepath"

	"github.com/tfversion/tfversion/pkg/helpers"
)

// CreateSymlink creates a symlink at the given path pointing to the source path.
func CreateSymlink(sourcePath, targetPath string) error {
	return os.Symlink(sourcePath, targetPath)
}

// RemoveSymlink removes the symlink at the given path.
func RemoveSymlink(path string) error {
	_, err := os.Lstat(path)
	if err != nil {
		return err
	}
	err = os.Remove(path)
	return err
}

// GetAliasLocation returns the directory where tfversion stores the aliases.
func GetAliasLocation() string {
	aliasLocation := filepath.Join(GetApplicationLocation(), AliasesDir)
	if _, err := os.Stat(aliasLocation); os.IsNotExist(err) {
		err := os.MkdirAll(aliasLocation, 0755)
		if err != nil {
			helpers.ExitWithError("creating alias directory", err)
		}
	}
	return aliasLocation
}

// isAlias checks if the given alias is valid.
func IsAlias(alias string) bool {
	aliasPath := filepath.Join(GetAliasLocation(), alias)
	_, err := os.Lstat(aliasPath)
	return !os.IsNotExist(err)
}

// getAliasVersion returns the Terraform version for the given alias.
func GetAliasVersion(alias string) string {
	aliasPath := filepath.Join(GetAliasLocation(), alias)
	resolvePath, err := filepath.EvalSymlinks(aliasPath)
	if err != nil {
		helpers.ExitWithError("resolving symlink", err)
	}
	_, targetVersion := filepath.Split(resolvePath)
	return targetVersion
}
