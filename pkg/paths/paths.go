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

// EnsureDirExists checks if the given directory exists, and creates it if it doesn't.
func EnsureDirExists(path string) error {
	_, err := os.Stat(path)
	if !os.IsNotExist(err) {
		return err
	}
	err = os.MkdirAll(path, 0755)
	return err
}

// RemoveDir removes the given directory and all its contents.
func RemoveDir(path string) error {
	return os.RemoveAll(path)
}
