package use

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/fatih/color"
	"github.com/tfversion/tfversion/pkg/helpers"
	"github.com/tfversion/tfversion/pkg/install"
	"github.com/tfversion/tfversion/pkg/list"
	"github.com/tfversion/tfversion/pkg/paths"
)

// UseVersion activates the specified Terraform version or one of the latest versions
func UseVersion(versionOrAlias string, autoInstall bool) {

	// find the version (via alias or directly)
	var version string
	if isAlias(versionOrAlias) {
		version = getAliasVersion(versionOrAlias)
	} else {
		version = versionOrAlias
	}

	// check if the version is installed
	if !paths.IsAlreadyDownloaded(version) {
		if !autoInstall {
			err := fmt.Errorf("terraform version %s not found, run %s to install", helpers.ColoredVersion(version), helpers.ColoredInstallHelper(version))
			helpers.ExitWithError("using", err)
		}
		install.InstallVersion(version)
	}

	// inform the user that they need to update their PATH
	path := os.Getenv("PATH")
	useLocation := paths.GetUseLocation()
	if !strings.Contains(path, useLocation) {
		fmt.Printf("%s not found in your shell PATH\n", color.CyanString(useLocation))
		fmt.Printf("Please run %s to make this version available in your shell\n", color.CyanString("`export PATH=%s:$PATH`", useLocation))
		fmt.Printf("Additionally, consider adding this line to your shell profile (e.g., .bashrc, .zshrc or fish config) for persistence.\n")
		os.Exit(1)
	}

	binaryTargetPath := filepath.Join(useLocation, paths.TerraformBinaryName)

	// ensure the symlink target is available
	err := paths.RemoveSymlink(binaryTargetPath)
	if err != nil {
		helpers.ExitWithError("removing symlink", err)
	}

	// create the symlink
	binaryVersionPath := paths.GetBinaryLocation(version)
	err = os.Symlink(binaryVersionPath, binaryTargetPath)
	if err != nil {
		helpers.ExitWithError("creating symlink", err)
	}

	fmt.Printf("Activated Terraform version %s\n", helpers.ColoredVersion(version))
}

// UseLatestVersion activates the latest Terraform version
func UseLatestVersion(preRelease bool, autoInstall bool) {
	version := list.FindLatestVersion(preRelease)
	UseVersion(version, autoInstall)
}

// UseRequiredVersion activates the required Terraform version from the .tf files in the current directory
func UseRequiredVersion(autoInstall bool) {
	terraformFiles, err := helpers.FindTerraformFiles()
	if err != nil {
		helpers.ExitWithError("finding Terraform files", err)
	}

	var foundVersion string
	availableVersions := list.GetAvailableVersions()
	for _, file := range terraformFiles {
		requiredVersion, err := helpers.FindRequiredVersionInFile(file, availableVersions)
		if err != nil {
			helpers.ExitWithError("finding required version", err)
		}

		if requiredVersion != "" {
			foundVersion = requiredVersion
			break
		}
	}

	if len(foundVersion) == 0 {
		err := fmt.Errorf("no required version found in current directory")
		helpers.ExitWithError("installing required version", err)
	}

	UseVersion(foundVersion, autoInstall)
}

// isAlias checks if the given alias is valid.
func isAlias(alias string) bool {
	aliasPath := paths.GetAliasLocation(alias)
	_, err := os.Lstat(aliasPath)
	return !os.IsNotExist(err)
}

// getAliasVersion returns the Terraform version for the given alias.
func getAliasVersion(alias string) string {
	aliasPath := paths.GetAliasLocation(alias)
	resolvePath, err := filepath.EvalSymlinks(aliasPath)
	if err != nil {
		helpers.ExitWithError("resolving symlink", err)
	}
	_, targetVersion := filepath.Split(resolvePath)
	return targetVersion
}
