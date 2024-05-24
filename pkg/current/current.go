package current

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/tfversion/tfversion/pkg/helpers"
	"github.com/tfversion/tfversion/pkg/paths"
)

// CheckCurrentVersion prints the current active version of Terraform.
func CheckCurrentVersion() {
	binaryPath := paths.GetActiveBinaryLocation()
	_, err := os.Lstat(binaryPath)
	if err != nil {
		helpers.ExitWithError("no current terraform version found", err)
	}

	realPath, err := filepath.EvalSymlinks(binaryPath)
	if err != nil {
		helpers.ExitWithError("resolving symlink", err)
	}
	_, currentVersion := filepath.Split(filepath.Dir(realPath))

	fmt.Printf("Current active Terraform version: %s\n", helpers.ColoredVersion(currentVersion))
}
