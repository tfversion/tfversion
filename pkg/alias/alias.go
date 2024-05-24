package alias

import (
	"fmt"
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

	aliasPath := filepath.Join(paths.GetAliasLocation(), alias)
	err := paths.RemoveSymlink(aliasPath)
	if err != nil {
		helpers.ExitWithError("removing symlink", err)
	}

	binaryVersionPath := paths.GetInstallLocation(version)
	err = paths.CreateSymlink(binaryVersionPath, aliasPath)
	if err != nil {
		helpers.ExitWithError("creating symlink", err)
	}

	fmt.Printf("Aliased Terraform version %s as %s\n", helpers.ColoredVersion(version), helpers.ColoredVersion(alias))
}
