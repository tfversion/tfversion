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
			fmt.Printf("Terraform version %s not found, run %s to install\n", color.YellowString(version), color.CyanString(fmt.Sprintf("`tfversion install %s`", version)))
		} else {
			fmt.Printf("Terraform version %s not found, run %s to install\n", color.CyanString(version), color.CyanString(fmt.Sprintf("`tfversion install %s`", version)))
		}
		os.Exit(0)
	}

	useLocation := getUseLocation()

	// inform the user that they need to update their PATH
	path := os.Getenv("PATH")
	if !strings.Contains(path, useLocation) {
		fmt.Println("Error: tfversion not found in your shell PATH.")
		fmt.Printf("Please run %s to make this version available in your shell\n", color.CyanString("`export PATH=%s:$PATH`", useLocation))
		fmt.Println("Additionally, consider adding this line to your shell profile (e.g., .bashrc, .zshrc or fish config) for persistence.")
		os.Exit(1)
	}

	// ensure the symlink target is available
	binaryTargetPath := filepath.Join(useLocation, download.TerraformBinaryName)
	_, err := os.Lstat(binaryTargetPath)
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
		fmt.Printf("Activated Terraform version %s\n", color.CyanString(version))
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

func getUseLocation() string {
	user, err := os.UserHomeDir()
	if err != nil {
		fmt.Printf("error getting user home directory: %s", err)
		os.Exit(1)
	}

	useLocation := filepath.Join(user, download.ApplicationDir, download.UseDir)
	if _, err := os.Stat(useLocation); os.IsNotExist(err) {
		err := os.Mkdir(useLocation, 0755)
		if err != nil {
			fmt.Printf("error creating use directory: %s", err)
			os.Exit(1)
		}
	}

	return useLocation
}
