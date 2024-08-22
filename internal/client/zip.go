package client

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"tfversion/internal/store"
)

// UnzipRelease extracts the Terraform binary from the zip file to the specified destination.
func UnzipRelease(source, destination string) error {
	// Open the zip file
	reader, err := zip.OpenReader(source)
	if err != nil {
		return err
	}
	defer reader.Close()

	// Create the destination directory if it does not exist.
	err = os.MkdirAll(destination, 0755)
	if err != nil {
		return err
	}

	// Get the absolute destination path.
	destination, err = filepath.Abs(destination)
	if err != nil {
		return err
	}

	// Iterate over zip files inside the archive and unzip only the "terraform" binary.
	for _, f := range reader.File {
		if f.Name == store.GetTerraformBinaryName() {
			return unzipFile(f, destination)
		}
	}

	return fmt.Errorf("terraform binary not found in the zip file")
}

func unzipFile(f *zip.File, destination string) error {
	// Check if file store are not vulnerable to Zip Slip.
	filePath := filepath.Join(destination, f.Name)
	if !strings.HasPrefix(filePath, filepath.Clean(destination)+string(os.PathSeparator)) {
		return fmt.Errorf("invalid file path: %s", filePath)
	}

	// Create a destination file for unzipped content.
	destinationFile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
	if err != nil {
		return err
	}
	defer destinationFile.Close()

	// Unzip the content of a file and copy it to the destination file.
	zippedFile, err := f.Open()
	if err != nil {
		return err
	}
	defer zippedFile.Close()

	_, err = io.Copy(destinationFile, zippedFile)
	return err
}
