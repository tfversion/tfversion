package use

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/fatih/color"
	"github.com/tfversion/tfversion/pkg/alias"
	"github.com/tfversion/tfversion/pkg/download"
	"github.com/tfversion/tfversion/pkg/helpers"
	"github.com/tfversion/tfversion/pkg/install"
	"github.com/tfversion/tfversion/pkg/list"
)

// UseVersion activates the specified Terraform version or one of the latest versions
func UseVersion(versionOrAlias string, autoInstall bool) {

	// find the version (via alias or directly)
	var version string
	if alias.IsAlias(versionOrAlias) {
		version = alias.GetVersion(versionOrAlias)
	} else {
		version = versionOrAlias
	}

	// check if the version is installed
	if !download.IsAlreadyDownloaded(version) {
		if !autoInstall {
			err := fmt.Errorf("terraform version %s not found, run %s to install", helpers.ColoredVersion(version), helpers.ColoredInstallHelper(version))
			helpers.ExitWithError("using", err)
		}
		install.InstallVersion(version)
	}

	// inform the user that they need to update their PATH
	path := os.Getenv("PATH")
	useLocation := GetUseLocation()
	if !strings.Contains(path, useLocation) {
		fmt.Printf("%s not found in your shell PATH\n", color.CyanString(useLocation))
		fmt.Printf("Please run %s to make this version available in your shell\n", color.CyanString("`export PATH=%s:$PATH`", useLocation))
		fmt.Printf("Additionally, consider adding this line to your shell profile (e.g., .bashrc, .zshrc or fish config) for persistence.\n")
		os.Exit(1)
	}

	// ensure the symlink target is available
	binaryTargetPath := filepath.Join(useLocation, download.TerraformBinaryName)
	_, err := os.Lstat(binaryTargetPath)
	if err == nil {
		err = os.Remove(binaryTargetPath)
		if err != nil {
			helpers.ExitWithError("removing symlink", err)
		}
	}

	// create the symlink
	binaryVersionPath := download.GetBinaryLocation(version)
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

func GetUseLocation() string {
	user, err := os.UserHomeDir()
	if err != nil {
		helpers.ExitWithError("user home directory", err)
	}

	useLocation := filepath.Join(user, download.ApplicationDir, download.UseDir)
	if _, err := os.Stat(useLocation); os.IsNotExist(err) {
		err := os.MkdirAll(useLocation, 0755)
		if err != nil {
			helpers.ExitWithError("creating use directory", err)
		}
	}

	return useLocation
}
