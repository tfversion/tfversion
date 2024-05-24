package use

import (
	"fmt"

	"github.com/tfversion/tfversion/pkg/client"
	"github.com/tfversion/tfversion/pkg/helpers"
	"github.com/tfversion/tfversion/pkg/install"
	"github.com/tfversion/tfversion/pkg/store"
)

// UseVersion activates the specified Terraform version or one of the latest versions
func UseVersion(version string, autoInstall bool) {
	// find the version (via alias or directly)
	if store.IsAlias(version) {
		version = store.GetAliasVersion(version)
	}

	// check if the version is installed
	if !store.IsInstalled(version) && !autoInstall {
		err := fmt.Errorf("terraform version %s not found, run %s to install", helpers.ColoredVersion(version), helpers.ColoredInstallHelper(version))
		helpers.ExitWithError("using", err)
	}

	// install the version if requested
	if autoInstall {
		install.InstallVersion(version)
	}

	// check the PATH environment
	helpers.WarnIfNotInPath(store.GetUseLocation())

	// ensure the symlink target is available
	binaryTargetPath := store.GetActiveBinaryLocation()
	err := store.RemoveSymlink(binaryTargetPath)
	if err != nil {
		helpers.ExitWithError("removing symlink", err)
	}

	// create the symlink
	binaryVersionPath := store.GetBinaryLocation(version)
	err = store.CreateSymlink(binaryVersionPath, binaryTargetPath)
	if err != nil {
		helpers.ExitWithError("creating symlink", err)
	}

	fmt.Printf("Activated Terraform version %s\n", helpers.ColoredVersion(version))
}

// UseLatestVersion activates the latest Terraform version
func UseLatestVersion(preRelease bool, autoInstall bool) {
	version := client.FindLatestVersion(preRelease)
	UseVersion(version, autoInstall)
}

// UseRequiredVersion activates the required Terraform version from the .tf files in the current directory
func UseRequiredVersion(autoInstall bool) {
	terraformFiles, err := helpers.FindTerraformFiles()
	if err != nil {
		helpers.ExitWithError("finding Terraform files", err)
	}

	var foundVersion string
	availableVersions := client.ListAvailableVersions()
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
