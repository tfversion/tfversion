package store

import (
	"os"
	"path/filepath"

	"tfversion/pkg/helpers"
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

// CreateSymlink creates a symlink at the given path pointing to the source path.
func CreateSymlink(sourcePath, targetPath string) error {
	return os.Symlink(sourcePath, targetPath)
}

// RemoveSymlink removes the symlink at the given path.
func RemoveSymlink(path string) error {
	_, err := os.Lstat(path)
	if os.IsNotExist(err) {
		return nil
	}
	if err != nil {
		return err
	}
	err = os.Remove(path)
	return err
}

// ListFiles returns a list of files in the given directory.
func ListFiles(path string) ([]os.DirEntry, error) {
	return os.ReadDir(path)
}
