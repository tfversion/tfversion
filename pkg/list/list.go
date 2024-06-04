package list

import (
	"fmt"

	"tfversion/pkg/client"
	"tfversion/pkg/helpers"
	"tfversion/pkg/store"
)

// GetAliasedVersions returns the aliased Terraform versions.
func GetAliasedVersions() []string {
	versionNames := store.GetAliasVersions()
	if len(versionNames) == 0 {
		err := fmt.Errorf("no versions found")
		helpers.ExitWithError("listing alias directory", err)
	}
	return versionNames
}

// GetInstalledVersions returns the installed Terraform versions.
func GetInstalledVersions() []string {
	versionNames := store.GetInstalledVersions()
	if len(versionNames) == 0 {
		err := fmt.Errorf("no versions found")
		helpers.ExitWithError("listing installed versions", err)
	}

	// Reverse the order of versionNames to show the latest version first
	reversedVersions := make([]string, len(versionNames))
	for i := 0; i < len(versionNames); i++ {
		reversedVersions[i] = versionNames[len(versionNames)-1-i]
	}

	return reversedVersions
}

// GetAvailableVersions returns the available Terraform versions from the official Terraform releases page
func GetAvailableVersions() []string {
	return client.ListAvailableVersions()
}

// FindLatestVersion finds the latest available Terraform version (or pre-release version)
func FindLatestVersion(preRelease bool) string {
	return client.FindLatestVersion(preRelease)
}
