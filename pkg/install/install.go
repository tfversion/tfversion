package install

import (
	"fmt"
	"os"
	"runtime"

	"github.com/tfversion/tfversion/pkg/download"
	"github.com/tfversion/tfversion/pkg/list"
)

// InstallVersion installs the specified Terraform version or one of the latest versions
func InstallVersion(version string, latest bool, preRelease bool) {
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

	if download.IsAlreadyDownloaded(version) {
		fmt.Printf("Terraform version %s is already installed\n", version)
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
