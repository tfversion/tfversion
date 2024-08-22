package store

import (
	"os"
	"path/filepath"

	"tfversion/internal/helpers"
)

// GetAliasLocation returns the directory where tfversion stores the aliases.
func GetAliasLocation() string {
	aliasLocation := filepath.Join(GetApplicationLocation(), AliasesDir)
	err := EnsureDirExists(aliasLocation)
	if err != nil {
		helpers.ExitWithError("creating alias directory", err)
	}
	return aliasLocation
}

// GetAliasPath returns the path to the symlink for the given alias.
func GetAliasPath(alias string) string {
	return filepath.Join(GetAliasLocation(), alias)
}

// isAlias checks if the given alias is valid.
func IsAlias(alias string) bool {
	_, err := os.Lstat(GetAliasPath(alias))
	if err != nil {
		return false
	}
	_, err = filepath.EvalSymlinks(GetAliasPath(alias))
	return err == nil
}
