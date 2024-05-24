package alias

import (
	"fmt"

	"github.com/tfversion/tfversion/pkg/helpers"
	"github.com/tfversion/tfversion/pkg/paths"
)

// AliasVersion creates a symlink to the specified Terraform version.
func AliasVersion(alias string, version string) {
	if !paths.IsInstalled(version) {
		err := fmt.Errorf("terraform version %s not found, run %s to install", helpers.ColoredVersion(version), helpers.ColoredInstallHelper(version))
		helpers.ExitWithError("aliasing", err)
	}

	aliasPath := paths.GetAliasPath(alias)
	err := paths.RemoveSymlink(aliasPath)
	if err != nil {
		helpers.ExitWithError("removing symlink", err)
	}

	binaryVersionPath := paths.GetInstalledVersionLocation(version)
	err = paths.CreateSymlink(binaryVersionPath, aliasPath)
	if err != nil {
		helpers.ExitWithError("creating symlink", err)
	}

	fmt.Printf("Aliased Terraform version %s as %s\n", helpers.ColoredVersion(version), helpers.ColoredVersion(alias))
}
