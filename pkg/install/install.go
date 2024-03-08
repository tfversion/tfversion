package install

import (
	"fmt"
	"runtime"
	"slices"

	"github.com/tfversion/tfversion/pkg/download"
	"github.com/tfversion/tfversion/pkg/helpers"
	"github.com/tfversion/tfversion/pkg/list"
)

// InstallVersion installs the specified Terraform version or one of the latest versions
func InstallVersion(version string) {
	if download.IsAlreadyDownloaded(version) {
		err := fmt.Errorf("terraform version %s is already installed", helpers.ColoredVersion(version))
		helpers.ExitWithError("installing", err)
	}

	// Check if the version exists
	availableVersions := list.GetAvailableVersions()
	if !slices.Contains(availableVersions, version) {
		err := fmt.Errorf("terraform version %s does not exist, please run %s to check available versions", helpers.ColoredVersion(version), helpers.ColoredListHelper())
		helpers.ExitWithError("installing", err)
	}

	// Download the Terraform release
	zipFile, err := download.Download(version, runtime.GOOS, runtime.GOARCH)
	if err != nil {
		helpers.ExitWithError("downloading", err)
	}

	// Unzip the downloaded Terraform release
	installLocation := download.GetInstallLocation(version)
	err = download.UnzipRelease(zipFile, installLocation)
	if err != nil {
		helpers.ExitWithError("unzipping", err)
	}

	// Clean up the downloaded zip file after unzipping
	err = download.DeleteDownloadedRelease(zipFile)
	if err != nil {
		helpers.ExitWithError("cleaning up", err)
	}
}

// InstallLatestVersion installs the latest Terraform version
func InstallLatestVersion(preRelease bool) {
	version := list.FindLatestVersion(preRelease)
	InstallVersion(version)
}

// InstallRequiredVersion installs the required Terraform version from the .tf files in the current directory
func InstallRequiredVersion() {
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

	InstallVersion(foundVersion)
}
