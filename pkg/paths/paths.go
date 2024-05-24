package paths

import (
	"os"
	"path/filepath"

	"github.com/tfversion/tfversion/pkg/helpers"
)

// ApplicationDir is the name of the directory where tfversion stores its data.
func GetApplicationLocation() string {
	user, err := os.UserHomeDir()
	if err != nil {
		helpers.ExitWithError("error getting user home directory", err)
	}
	applicationLocation := filepath.Join(user, ApplicationDir)
	err = EnsureDirExists(applicationLocation)
	if err != nil {
		helpers.ExitWithError("creating application directory", err)
	}
	return applicationLocation
}

// GetDownloadLocation returns the directory where tfversion downloads Terraform releases to.
func GetDownloadLocation() string {
	downloadLocation := filepath.Join(GetApplicationLocation(), VersionsDir)
	err := EnsureDirExists(downloadLocation)
	if err != nil {
		helpers.ExitWithError("creating download directory", err)
	}
	return downloadLocation
}

// EnsureDirExists checks if the given directory exists, and creates it if it doesn't.
func EnsureDirExists(path string) error {
	_, err := os.Stat(path)
	if !os.IsNotExist(err) {
		return err
	}
	err = os.MkdirAll(path, 0755)
	return err
}
