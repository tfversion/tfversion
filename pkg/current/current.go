package current

import (
	"fmt"

	"tfversion/pkg/helpers"
	"tfversion/pkg/store"
)

// CheckCurrentVersion prints the current active version of Terraform.
func CheckCurrentVersion() {
	version := store.GetActiveVersion()
	fmt.Printf("Current active Terraform version: %s\n", helpers.ColoredVersion(version))
}
