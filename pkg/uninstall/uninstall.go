package uninstall

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/tfversion/tfversion/pkg/download"
	"github.com/tfversion/tfversion/pkg/helpers"
)

func deleteVersionFromDownloadLocation(version string) error {
	installLocation := download.GetInstallLocation(version)
	err := os.RemoveAll(installLocation)
	if err != nil {
		return fmt.Errorf("deleting Terraform version failed: %s", err)
	}
	return nil
}

func Uninstall(version string) {
	if !checkIfBinaryIsPresent(version) {
		if helpers.IsPreReleaseVersion(version) {
			fmt.Printf("Terraform version %s is not installed\n", color.YellowString(version))
		} else {
			fmt.Printf("Terraform version %s is not installed\n", color.CyanString(version))
		}
		os.Exit(1)
	}

	err := deleteVersionFromDownloadLocation(version)
	if err != nil {
		fmt.Printf("error deleting Terraform version: %s\n", err)
		os.Exit(1)
	}
}

func checkIfBinaryIsPresent(version string) bool {
	return download.IsAlreadyDownloaded(version)
}
