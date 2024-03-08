package download

import (
	"os"
	"path/filepath"
	"testing"
)

func TestGetDownloadLocation(t *testing.T) {
	tempDir := t.TempDir()
	os.Setenv("HOME", tempDir)
	downloadLocation := GetDownloadLocation()

	expectedLocation := filepath.Join(tempDir, ".tfversion", "versions")
	if downloadLocation != expectedLocation {
		t.Errorf("Expected download location %s, but got %s", expectedLocation, downloadLocation)
	}

	_, err := os.Stat(downloadLocation)
	if os.IsNotExist(err) {
		t.Errorf("Download location directory does not exist: %s", downloadLocation)
	}
}
