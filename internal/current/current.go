package current

import (
	"fmt"

	"tfversion/internal/helpers"
	"tfversion/internal/store"
)

// CheckCurrentVersion prints the current active version of Terraform.
func CheckCurrentVersion() {
	version := store.GetActiveVersion()
	fmt.Printf("Current active Terraform version: %s\n", helpers.ColoredVersion(version))
}
