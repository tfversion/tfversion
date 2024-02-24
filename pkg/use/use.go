package use

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/tfversion/tfversion/pkg/download"
	"github.com/tfversion/tfversion/pkg/list"
)

// UseVersion activates the specified Terraform version or one of the latest versions
func UseVersion(version string, latest bool, preRelease bool) {
	// Get the available Terraform versions
	versions := list.GetAvailableVersions()

	// Set the version to the latest stable version if the `latest` flag is set
	// or to the latest pre-release version if the `latest` and `pre-release` flags are set
	if latest {
		for _, v := range versions {
			if !preRelease && list.IsPreReleaseVersion(v) {
				continue
			}
			version = v
			break
		}
	}

	if !download.IsAlreadyDownloaded(version) {
		fmt.Printf("Terraform version %s is not installed\n", version)
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

	fmt.Printf("Activated Terraform version %s\n", version)

	// inform the user that they need to update their PATH
	path := os.Getenv("PATH")
	if !strings.Contains(path, targetPath) {
		fmt.Printf("Please run `export PATH=%s:$PATH` to make this version available in your shell\n", targetPath)
	}
}
