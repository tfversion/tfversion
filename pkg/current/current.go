package current

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/tfversion/tfversion/pkg/download"
	"github.com/tfversion/tfversion/pkg/helpers"
	"github.com/tfversion/tfversion/pkg/use"
)

// CheckCurrentVersion prints the current active version of Terraform.
func CheckCurrentVersion() {

	symlinkPath := filepath.Join(use.GetUseLocation(), download.TerraformBinaryName)
	_, err := os.Lstat(symlinkPath)
	if err != nil {
		helpers.ExitWithError("no current terraform version found", err)
	}

	realPath, err := filepath.EvalSymlinks(symlinkPath)
	if err != nil {
		helpers.ExitWithError("resolving symlink", err)
	}
	_, currentVersion := filepath.Split(filepath.Dir(realPath))

	fmt.Printf("Current active Terraform version: %s\n", helpers.ColoredVersion(currentVersion))
}
