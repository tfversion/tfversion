package store

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/tfversion/tfversion/pkg/helpers"
)

// GetInstalledVersions returns a list of all installed Terraform versions.
func GetInstalledVersions() []string {
	installLocation := GetInstallLocation()
	installedVersions, err := os.ReadDir(installLocation)
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
	aliasLocation := GetAliasLocation()
	aliasedVersions, err := os.ReadDir(aliasLocation)
	if err != nil {
		helpers.ExitWithError("listing alias directory", err)
	}
	var versionNames []string
	for _, v := range aliasedVersions {
		versionNames = append(versionNames, fmt.Sprintf("%s -> %s", v.Name(), GetAliasVersion(v.Name())))
	}
	return versionNames
}
