package download

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/tfversion/tfversion/pkg/helpers"
)

// IsAlreadyDownloaded checks if the given Terraform version is already downloaded and unzipped.
func IsAlreadyDownloaded(version string) bool {
	binaryPath := GetBinaryLocation(version)
	_, err := os.Stat(binaryPath)
	return !os.IsNotExist(err)
}

// GetDownloadLocation returns the directory where tfversion downloads Terraform releases to.
func GetDownloadLocation() string {
	user, err := os.UserHomeDir()
	if err != nil {
		helpers.ExitWithError("error getting user home directory", err)
	}

	downloadLocation := filepath.Join(user, ApplicationDir, VersionsDir)
	if _, err := os.Stat(downloadLocation); os.IsNotExist(err) {
		err := os.MkdirAll(downloadLocation, 0755)
		if err != nil {
			helpers.ExitWithError("error creating download directory", err)
		}
	}

	return downloadLocation
}

// GetInstallLocation returns the directory where a specific Terraform version is installed to.
func GetInstallLocation(version string) string {
	return filepath.Join(GetDownloadLocation(), version)
}

// GetBinaryLocation returns the path to the Terraform binary for the given version.
func GetBinaryLocation(version string) string {
	return filepath.Join(GetInstallLocation(version), TerraformBinaryName)
}

// Download downloads the Terraform release zip file for the given version, OS and architecture.
func Download(version, goos, goarch string) (string, error) {
	downloadLocation := GetDownloadLocation()

	// construct the download URL based on the version and the OS and architecture
	downloadURL := fmt.Sprintf("%s/%s/terraform_%s_%s_%s.zip", TerraformReleasesUrl, version, version, goos, goarch)

	var err error
	for attempt := 1; attempt <= MaxRetries; attempt++ {
		if err = downloadWithRetry(downloadURL, downloadLocation, version, goos, goarch); err == nil {
			fmt.Printf("Terraform version %s downloaded successfully\n", helpers.ColoredVersion(version))
			return fmt.Sprintf("%s/terraform_%s_%s_%s.zip", downloadLocation, version, goos, goarch), nil
		}

		fmt.Printf("Attempt %d failed: %s\n", attempt, err)
		time.Sleep(time.Second * RetryTimeInSeconds)
	}

	// if we got here, we failed to download Terraform after MaxRetries attempts
	return "", fmt.Errorf("failed to download Terraform after %d attempts: %s", MaxRetries, err)
}

func downloadWithRetry(downloadURL, downloadLocation, version, goos, goarch string) error {
	resp, err := http.Get(downloadURL)
	if err != nil {
		return fmt.Errorf("failed to download Terraform: %s", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to download Terraform: %s", resp.Status)
	}

	filePath := filepath.Join(downloadLocation, fmt.Sprintf("terraform_%s_%s_%s.zip", version, goos, goarch))
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create file: %s", err)
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return fmt.Errorf("failed to write to file: %s", err)
	}

	return nil
}

// DeleteDownloadedRelease deletes the downloaded Terraform release zip file to prevent cluttering the filesystem.
func DeleteDownloadedRelease(zipFile string) error {
	err := os.Remove(zipFile)
	if err != nil {
		return fmt.Errorf("failed to delete Terraform release: %s", err)
	}
	return nil
}
