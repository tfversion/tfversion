package install

import (
	"fmt"
	"os"
	"runtime"

	"github.com/fatih/color"
	"github.com/tfversion/tfversion/pkg/download"
	"github.com/tfversion/tfversion/pkg/helpers"
	"github.com/tfversion/tfversion/pkg/list"
)

// InstallVersion installs the specified Terraform version or one of the latest versions
func InstallVersion(version string) {
	if download.IsAlreadyDownloaded(version) {
		if helpers.IsPreReleaseVersion(version) {
			fmt.Printf("Terraform version %s is already installed\n", color.YellowString(version))
		} else {
			fmt.Printf("Terraform version %s is already installed\n", color.BlueString(version))
		}
		os.Exit(0)
	}

	// Download the Terraform release
	zipFile, err := download.Download(version, runtime.GOOS, runtime.GOARCH)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Unzip the downloaded Terraform release
	installLocation := download.GetInstallLocation(version)
	err = download.UnzipRelease(zipFile, installLocation)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Clean up the downloaded zip file after unzipping
	err = download.DeleteDownloadedRelease(zipFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// InstallLatestVersion installs the latest Terraform version
func InstallLatestVersion(preRelease bool) {
	version := list.FindLatestVersion(preRelease)
	InstallVersion(version)
}

// InstallRequiredVersion installs the required Terraform version from the .tf files in the current directory
func InstallRequiredVersion() {
	terraformFiles := helpers.FindTerraformFiles()
	if len(terraformFiles) == 0 {
		fmt.Println("error: no Terraform files found in current directory")
		os.Exit(1)
	}

	var foundVersion string
	availableVersions := list.GetAvailableVersions()
	for _, file := range terraformFiles {
		requiredVersion := helpers.FindRequiredVersionInFile(file, availableVersions)
		if requiredVersion != "" {
			foundVersion = requiredVersion
		}
	}

	if len(foundVersion) == 0 {
		fmt.Println("error: no required version found in current directory")
		os.Exit(1)
	}

	InstallVersion(foundVersion)
}
