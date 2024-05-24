package unalias

import (
	"fmt"
	"path/filepath"

	"github.com/tfversion/tfversion/pkg/helpers"
	"github.com/tfversion/tfversion/pkg/paths"
)

// Unalias removes the symlink for the specified alias.
func Unalias(aliasName string) {
	if !paths.IsAlias(aliasName) {
		return
	}

	aliasPath := filepath.Join(paths.GetAliasLocation(), aliasName)
	err := paths.RemoveSymlink(aliasPath)
	if err != nil {
		helpers.ExitWithError("removing symlink", err)
	}

	fmt.Printf("Removed alias %s\n", helpers.ColoredVersion(aliasName))
}
