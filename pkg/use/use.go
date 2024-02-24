package use

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/bschaatsbergen/tfversion/pkg/download"
)

func UseVersion(version string) {
	if !download.IsAlreadyDownloaded(version) {
		fmt.Printf("Terraform version %s is not installed\n", version)
		os.Exit(0)
	}

	// create the bin directory if it doesn't exist
	targetPath := filepath.Join(download.GetDownloadLocation(), download.BinaryDir)
	os.Mkdir(targetPath, 0755)

	// ensure the symlink target is available
	binaryTargetPath := filepath.Join(targetPath, download.TerraformBinaryName)
	_, err := os.Lstat(binaryTargetPath)
	if err == nil {
		os.Remove(binaryTargetPath)
	}

	// create the symlink
	binaryVersionPath := download.GetBinaryLocation(version)
	os.Symlink(binaryVersionPath, binaryTargetPath)

	fmt.Printf("Activated Terraform version %s\n", version)
	fmt.Printf("Please run `export PATH=%s:$PATH` to make this version available in your shell\n", targetPath)
}
