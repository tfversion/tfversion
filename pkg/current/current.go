package current

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"regexp"

	"github.com/tfversion/tfversion/pkg/download"
	"github.com/tfversion/tfversion/pkg/helpers"
	"github.com/tfversion/tfversion/pkg/use"
)

// CheckCurrentVersion prints the current active version of Terraform.
func CheckCurrentVersion() {

	symlinkPath := filepath.Join(use.GetUseLocation(), download.TerraformBinaryName)
	_, err := os.Lstat(symlinkPath)
	if err != nil {
		helpers.ExitWithError("could not determine active Terraform version", err)
	}

	realPath, err := filepath.EvalSymlinks(symlinkPath)
	if err != nil {
		helpers.ExitWithError("could not determine active Terraform version", err)
	}
	left := fmt.Sprintf("%s/", download.VersionsDir)
	right := fmt.Sprintf("/%s", download.TerraformBinaryName)
	rx := regexp.MustCompile(`(?s)` + regexp.QuoteMeta(left) + `(.*?)` + regexp.QuoteMeta(right))
	match := rx.FindStringSubmatch(realPath)

	if match == nil {
		// helpers.er
		helpers.ExitWithError("could not determine active Terraform version", errors.New("failed to match regex on path"))
	}
	fmt.Printf("Current active Terraform version: %s\n", helpers.ColoredVersion(match[1]))
}
