package paths

import (
	"fmt"
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
	err := EnsureDirExists(aliasLocation)
	if err != nil {
		helpers.ExitWithError("creating alias directory", err)
	}
	return aliasLocation
}

func GetAliasPath(alias string) string {
	return filepath.Join(GetAliasLocation(), alias)
}

// isAlias checks if the given alias is valid.
func IsAlias(alias string) bool {
	_, err := os.Lstat(GetAliasPath(alias))
	return !os.IsNotExist(err)
}

// GetAliasVersion returns the Terraform version for the given alias.
func GetAliasVersion(alias string) string {
	resolvePath, err := filepath.EvalSymlinks(GetAliasPath(alias))
	if err != nil {
		helpers.ExitWithError("resolving symlink", err)
	}
	_, targetVersion := filepath.Split(resolvePath)
	return targetVersion
}

// GetAliasVersions returns a list of all aliases and their corresponding Terraform versions.
func GetAliasVersions() []string {
	aliasLocation := GetAliasLocation()
	aliasedVersions, err := os.ReadDir(aliasLocation)
	if err != nil {
		helpers.ExitWithError("listing alias directory", err)
	}
	var versionNames []string
	for _, v := range aliasedVersions {
		versionNames = append(versionNames, fmt.Sprintf("%s -> %s", v.Name(), GetAliasVersion(v.Name())))
	}
	return versionNames
}
