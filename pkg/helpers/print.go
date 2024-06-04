package helpers

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

// ColoredVersion returns a colored version string
func ColoredVersion(version string) string {
	if IsPreReleaseVersion(version) {
		return color.YellowString(version)
	}
	return color.CyanString(version)
}

// ColoredInstallHelper returns a colored install helper string
func ColoredInstallHelper(version string) string {
	return color.CyanString(fmt.Sprintf("`tfversion install %s`", version))
}

func ColoredListHelper() string {
	return color.CyanString("`tfversion list`")
}

func ColoredUnavailableVersion() string {
	return color.HiRedString("[removed]")
}

// ExitWithError prints an error message and exits with status code 1
func ExitWithError(message string, err error) {
	fmt.Printf("%s %s: %s\n", color.HiRedString("error:"), message, err)
	os.Exit(1)
}

// ErrorWithHelp returns an error with a help message
func ErrorWithHelp(help string) error {
	return fmt.Errorf("see %s for help and examples", color.CyanString("`%s`", help))
}
