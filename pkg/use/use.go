package use

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/fatih/color"
	"github.com/tfversion/tfversion/pkg/download"
	"github.com/tfversion/tfversion/pkg/helpers"
	"github.com/tfversion/tfversion/pkg/list"
)

// UseVersion activates the specified Terraform version or one of the latest versions
func UseVersion(version string) {
	if !download.IsAlreadyDownloaded(version) {
		if helpers.IsPreReleaseVersion(version) {
			fmt.Printf("Terraform version %s not found, run %s to install\n", color.YellowString(version), color.BlueString(fmt.Sprintf("`tfversion install %s`", version)))
		} else {
			fmt.Printf("Terraform version %s not found, run %s to install\n", color.BlueString(version), color.BlueString(fmt.Sprintf("`tfversion install %s`", version)))
		}
		os.Exit(0)
	}

	// create the bin directory if it doesn't exist
	targetPath := filepath.Join(download.GetDownloadLocation(), download.BinaryDir)
	_, err := os.Stat(targetPath)
	if os.IsNotExist(err) {
		err = os.Mkdir(targetPath, 0755)
		if err != nil {
			fmt.Printf("error creating directory: %v\n", err)
			os.Exit(1)
		}
	}

	// inform the user that they need to update their PATH
	path := os.Getenv("PATH")
	if !strings.Contains(path, targetPath) {
		fmt.Println("Error: tfversion not found in your shell PATH.")
		fmt.Printf("Please run %s to make this version available in your shell\n", color.BlueString("`export PATH=%s:$PATH`", targetPath))
		fmt.Println("Additionally, consider adding this line to your shell profile (e.g., .bashrc, .zshrc or fish config) for persistence.")
		os.Exit(1)
	}

	// ensure the symlink target is available
	binaryTargetPath := filepath.Join(targetPath, download.TerraformBinaryName)
	_, err = os.Lstat(binaryTargetPath)
	if err == nil {
		err = os.Remove(binaryTargetPath)
		if err != nil {
			fmt.Printf("error removing symlink: %v\n", err)
			os.Exit(1)
		}
	}

	// create the symlink
	binaryVersionPath := download.GetBinaryLocation(version)
	err = os.Symlink(binaryVersionPath, binaryTargetPath)
	if err != nil {
		fmt.Printf("error creating symlink: %v\n", err)
		os.Exit(1)
	}

	if helpers.IsPreReleaseVersion(version) {
		fmt.Printf("Activated Terraform version %s\n", color.YellowString(version))
	} else {
		fmt.Printf("Activated Terraform version %s\n", color.BlueString(version))
	}
}

// UseLatestVersion activates the latest Terraform version
func UseLatestVersion(preRelease bool) {
	version := list.FindLatestVersion(preRelease)
	UseVersion(version)
}

// UseRequiredVersion activates the required Terraform version from the .tf files in the current directory
func UseRequiredVersion() {
	terraformFiles := helpers.FindTerraformFiles()
	if len(terraformFiles) == 0 {
		fmt.Println("No Terraform files found in current directory")
		os.Exit(1)
	}

	availableVersions := list.GetAvailableVersions()
	for _, file := range terraformFiles {
		requiredVersion := helpers.FindRequiredVersionInFile(file, availableVersions)
		if requiredVersion != "" {
			UseVersion(requiredVersion)
			break
		}
	}
}
