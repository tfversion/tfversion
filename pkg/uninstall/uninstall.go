package uninstall

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/tfversion/tfversion/pkg/download"
)

func deleteVersionFromDownloadLocation(version string) error {
	installLocation := download.GetInstallLocation(version)
	err := os.RemoveAll(installLocation)
	if err != nil {
		return fmt.Errorf("deleting Terraform version failed: %s", err)
	}
	return nil
}

func Uninstall(version string) error {
	if !checkIfBinaryIsPresent(version) {
		fmt.Printf("Terraform version %s is not installed\n", color.RedString(version))
		os.Exit(1)
	}

	err := deleteVersionFromDownloadLocation(version)
	if err != nil {
		return err
	}
	return nil
}

func checkIfBinaryIsPresent(version string) bool {
	return download.IsAlreadyDownloaded(version)
}
