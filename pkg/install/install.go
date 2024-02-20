package install

import (
	"fmt"
	"os"
	"runtime"

	"github.com/bschaatsbergen/tfversion/pkg/download"
)

func InstallVersion(version string) {
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
