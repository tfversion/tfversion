package unalias

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/tfversion/tfversion/pkg/helpers"
	"github.com/tfversion/tfversion/pkg/paths"
)

// Unalias removes the symlink for the specified alias.
func Unalias(aliasName string) {
	aliasPath := filepath.Join(paths.GetAliasLocation(), aliasName)
	_, err := os.Lstat(aliasPath)
	if err != nil {
		helpers.ExitWithError("removing symlink", err)
	}

	err = os.RemoveAll(aliasPath)
	if err != nil {
		helpers.ExitWithError("removing symlink", err)
	}

	fmt.Printf("Removed alias %s\n", helpers.ColoredVersion(aliasName))
}
