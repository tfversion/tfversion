package client

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"tfversion/internal/helpers"
	"tfversion/internal/store"
)

// Download downloads the Terraform release zip file for the given version, OS and architecture.
func Download(version string) (string, error) {
	downloadLocation := store.GetInstallLocation()

	// construct the file name based on the version, OS and architecture
	fileName := fmt.Sprintf("terraform_%s_%s_%s.zip", version, runtime.GOOS, runtime.GOARCH)
	downloadURL := fmt.Sprintf("%s/%s/%s", TerraformReleasesUrl, version, fileName)

	var err error
	var attempt int
	for attempt = 1; attempt <= MaxRetries; attempt++ {
		if err = downloadWithRetry(downloadURL, downloadLocation, fileName); err == nil {
			fmt.Printf("Terraform version %s downloaded successfully\n", helpers.ColoredVersion(version))
			return fmt.Sprintf("%s/%s", downloadLocation, fileName), nil
		}
		if strings.Contains(err.Error(), "404 Not Found") {
			if attempt == 1 && runtime.GOOS == "darwin" && runtime.GOARCH == "arm64" {
				fmt.Printf("Attempt %d failed: %s\n", attempt, err)
				fmt.Printf("ARM build not found, retrying with amd64 build\n")
				fileName = fmt.Sprintf("terraform_%s_%s_amd64.zip", version, runtime.GOOS)
				downloadURL = fmt.Sprintf("%s/%s/%s", TerraformReleasesUrl, version, fileName)
				continue
			}
			break
		}
		fmt.Printf("Attempt %d failed: %s\n", attempt, err)
		time.Sleep(time.Second * RetryTimeInSeconds)
	}

	// if we got here, we failed to download Terraform after MaxRetries attempts
	return "", fmt.Errorf("failed to download Terraform from %s after %d attempts: %s", downloadURL, attempt, err)
}

func downloadWithRetry(downloadURL, downloadLocation, fileName string) error {
	resp, err := http.Get(downloadURL)
	if err != nil {
		return fmt.Errorf("failed to download Terraform %s: %s", fileName, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to download Terraform %s: %s", fileName, resp.Status)
	}

	filePath := filepath.Join(downloadLocation, fileName)
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create file %s: %s", fileName, err)
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return fmt.Errorf("failed to write to file %s: %s", fileName, err)
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
