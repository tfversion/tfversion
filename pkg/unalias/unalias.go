package unalias

import (
	"fmt"

	"tfversion/pkg/helpers"
	"tfversion/pkg/store"
)

// Unalias removes the symlink for the specified alias.
func Unalias(aliasName string) {
	if !store.IsAlias(aliasName) {
		return
	}

	err := store.RemoveSymlink(store.GetAliasPath(aliasName))
	if err != nil {
		helpers.ExitWithError("removing symlink", err)
	}

	fmt.Printf("Removed alias %s\n", helpers.ColoredVersion(aliasName))
}
