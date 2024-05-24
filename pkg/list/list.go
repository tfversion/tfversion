package list

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/tfversion/tfversion/pkg/client"
	"github.com/tfversion/tfversion/pkg/helpers"
	"github.com/tfversion/tfversion/pkg/paths"
)

// GetAliasedVersions returns the aliased Terraform versions.
func GetAliasedVersions() []string {
	aliasLocation := paths.GetAliasLocation()

	// find all aliases
	aliasedVersions, err := os.ReadDir(aliasLocation)
	if err != nil {
		helpers.ExitWithError("listing alias directory", err)
	}

	// resolve the symlinks to get the target versions
	var versionNames []string
	for _, v := range aliasedVersions {
		resolvePath, _ := filepath.EvalSymlinks(filepath.Join(aliasLocation, v.Name()))
		_, targetVersion := filepath.Split(resolvePath)
		versionNames = append(versionNames, fmt.Sprintf("%s -> %s", v.Name(), targetVersion))
	}

	// check if there are any versions
	if len(versionNames) == 0 {
		err := fmt.Errorf("no versions found")
		helpers.ExitWithError("listing alias directory", err)
	}

	return versionNames
}

// GetInstalledVersions returns the installed Terraform versions.
func GetInstalledVersions() []string {
	installLocation := paths.GetDownloadLocation()
	installedVersions, err := os.ReadDir(installLocation)
	if err != nil {
		helpers.ExitWithError("listing versions directory", err)
	}

	var versionNames []string
	for _, v := range installedVersions {
		versionNames = append(versionNames, v.Name())
	}

	// check if there are any versions
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
