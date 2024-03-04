package alias

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/fatih/color"
	"github.com/tfversion/tfversion/pkg/download"
	"github.com/tfversion/tfversion/pkg/helpers"
)

func AliasVersion(alias string, version string) {
	if !download.IsAlreadyDownloaded(version) {
		if helpers.IsPreReleaseVersion(version) {
			fmt.Printf("Terraform version %s not found, run %s to install\n", color.YellowString(version), color.CyanString(fmt.Sprintf("`tfversion install %s`", version)))
		} else {
			fmt.Printf("Terraform version %s not found, run %s to install\n", color.CyanString(version), color.CyanString(fmt.Sprintf("`tfversion install %s`", version)))
		}
		os.Exit(0)
	}

	// delete existing alias symlink, we consider it non-destructive anyways since you can easily restore it
	aliasPath := filepath.Join(download.GetDownloadLocation(), alias)
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
