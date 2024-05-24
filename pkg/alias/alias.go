package alias

import (
	"fmt"

	"github.com/tfversion/tfversion/pkg/helpers"
	"github.com/tfversion/tfversion/pkg/store"
)

// AliasVersion creates a symlink to the specified Terraform version.
func AliasVersion(alias string, version string) {
	if !store.IsInstalled(version) {
		err := fmt.Errorf("terraform version %s not found, run %s to install", helpers.ColoredVersion(version), helpers.ColoredInstallHelper(version))
		helpers.ExitWithError("aliasing", err)
	}

	aliasPath := store.GetAliasPath(alias)
	err := store.RemoveSymlink(aliasPath)
	if err != nil {
		helpers.ExitWithError("removing symlink", err)
	}

	binaryVersionPath := store.GetInstalledVersionLocation(version)
	err = store.CreateSymlink(binaryVersionPath, aliasPath)
	if err != nil {
		helpers.ExitWithError("creating symlink", err)
	}

	fmt.Printf("Aliased Terraform version %s as %s\n", helpers.ColoredVersion(version), helpers.ColoredVersion(alias))
}
