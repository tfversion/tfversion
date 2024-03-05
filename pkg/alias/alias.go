package alias

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/fatih/color"
	"github.com/tfversion/tfversion/pkg/download"
	"github.com/tfversion/tfversion/pkg/helpers"
)

// AliasVersion creates a symlink to the specified Terraform version.
func AliasVersion(alias string, version string) {
	if !download.IsAlreadyDownloaded(version) {
		if helpers.IsPreReleaseVersion(version) {
			fmt.Printf("Terraform version %s not found, run %s to install\n", color.YellowString(version), color.CyanString(fmt.Sprintf("`tfversion install %s`", version)))
		} else {
			fmt.Printf("Terraform version %s not found, run %s to install\n", color.CyanString(version), color.CyanString(fmt.Sprintf("`tfversion install %s`", version)))
		}
		os.Exit(0)
	}

	aliasLocation := GetAliasLocation()

	// delete existing alias symlink, we consider it non-destructive anyways since you can easily restore it
	aliasPath := filepath.Join(aliasLocation, alias)
	_, err := os.Lstat(aliasPath)
	if err == nil {
		err = os.RemoveAll(aliasPath)
		if err != nil {
			fmt.Printf("error removing symlink: %v\n", err)
			os.Exit(1)
		}
	}

	// create the symlink
	binaryVersionPath := download.GetInstallLocation(version)
	err = os.Symlink(binaryVersionPath, aliasPath)
	if err != nil {
		fmt.Printf("error creating symlink: %v\n", err)
		os.Exit(1)
	}

	if helpers.IsPreReleaseVersion(version) {
		fmt.Printf("Aliased Terraform version %s as %s\n", color.YellowString(version), color.YellowString(alias))
	} else {
		fmt.Printf("Aliased Terraform version %s as %s\n", color.CyanString(version), color.CyanString(alias))
	}
}

// GetAliasLocation returns the directory where tfversion stores the aliases.
func GetAliasLocation() string {
	user, err := os.UserHomeDir()
	if err != nil {
		fmt.Printf("error getting user home directory: %s", err)
		os.Exit(1)
	}

	aliasLocation := filepath.Join(user, download.ApplicationDir, download.AliasesDir)
	if _, err := os.Stat(aliasLocation); os.IsNotExist(err) {
		err := os.Mkdir(aliasLocation, 0755)
		if err != nil {
			fmt.Printf("error creating alias directory: %s", err)
			os.Exit(1)
		}
	}

	return aliasLocation
}

// IsAlias checks if the given alias is valid.
func IsAlias(alias string) bool {
	aliasPath := filepath.Join(GetAliasLocation(), alias)
	_, err := os.Stat(aliasPath)
	return !os.IsNotExist(err)
}

// GetVersion returns the Terraform version for the given alias.
func GetVersion(alias string) string {
	aliasLocation := GetAliasLocation()
	resolvePath, _ := filepath.EvalSymlinks(filepath.Join(aliasLocation, alias))
	_, targetVersion := filepath.Split(resolvePath)
	return targetVersion
}
