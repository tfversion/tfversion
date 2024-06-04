package store

import (
	"fmt"
	"os"
	"path/filepath"

	"tfversion/pkg/helpers"
)

// GetInstalledVersions returns a list of all installed Terraform versions.
func GetInstalledVersions() []string {
	installedVersions, err := ListFiles(GetInstallLocation())
	if err != nil {
		helpers.ExitWithError("listing versions directory", err)
	}
	var versionNames []string
	for _, v := range installedVersions {
		versionNames = append(versionNames, v.Name())
	}
	return versionNames
}

// GetActiveVersion returns the currently active Terraform version.
func GetActiveVersion() string {
	binaryPath := GetActiveBinaryLocation()
	_, err := os.Lstat(binaryPath)
	if err != nil {
		helpers.ExitWithError("no current terraform version found", err)
	}
	realPath, err := filepath.EvalSymlinks(binaryPath)
	if err != nil {
		helpers.ExitWithError("resolving symlink", err)
	}
	_, version := filepath.Split(filepath.Dir(realPath))
	return version
}

// GetAliasVersion returns the Terraform version for the given alias.
func GetAliasVersion(alias string) string {
	resolvePath, err := filepath.EvalSymlinks(GetAliasPath(alias))
	if err != nil {
		helpers.ExitWithError("resolving symlink", err)
	}
	_, targetVersion := filepath.Split(resolvePath)
	return targetVersion
}

// GetAliasVersions returns a list of all aliases and their corresponding Terraform versions.
func GetAliasVersions() []string {
	aliasedVersions, err := ListFiles(GetAliasLocation())
	if err != nil {
		helpers.ExitWithError("listing alias directory", err)
	}
	var versionNames []string
	for _, v := range aliasedVersions {
		versionNames = append(versionNames, formatAliasListItem(v.Name()))
	}
	return versionNames
}

func formatAliasListItem(alias string) string {
	return fmt.Sprintf("%s -> %s", alias, formatAliasVersion(alias))
}

func formatAliasVersion(alias string) string {
	if IsAlias(alias) {
		return GetAliasVersion(alias)
	}
	return helpers.ColoredUnavailableVersion()
}
