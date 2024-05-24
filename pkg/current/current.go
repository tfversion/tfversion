package current

import (
	"fmt"

	"github.com/tfversion/tfversion/pkg/helpers"
	"github.com/tfversion/tfversion/pkg/paths"
)

// CheckCurrentVersion prints the current active version of Terraform.
func CheckCurrentVersion() {
	version := paths.GetActiveVersion()
	fmt.Printf("Current active Terraform version: %s\n", helpers.ColoredVersion(version))
}
