package client

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/tfversion/tfversion/pkg/helpers"
	"github.com/tfversion/tfversion/pkg/paths"
)

// Download downloads the Terraform release zip file for the given version, OS and architecture.
func Download(version, goos, goarch string) (string, error) {
	downloadLocation := paths.GetDownloadLocation()

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
