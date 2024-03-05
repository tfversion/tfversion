package helpers

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/fatih/color"
	"github.com/hashicorp/go-version"
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

// ExitWithError prints an error message and exits with status code 1
func ExitWithError(message string, err error) {
	fmt.Printf("%s %s: %s\n", color.HiRedString("error:"), message, err)
	os.Exit(1)
}

// ErorWithHelp returns an error with a help message
func ErorWithHelp(help string) error {
	return fmt.Errorf("see %s for help and examples", color.CyanString("`%s`", help))
}

// IsPreReleaseVersion checks if the given version is a Terraform pre-release version
func IsPreReleaseVersion(version string) bool {
	return strings.Contains(version, "-alpha") || strings.Contains(version, "-beta") || strings.Contains(version, "-rc")
}

// FindRequiredVersionInFile finds the required Terraform version in a given .tf file (using required_version = ">= x.x.x")
func FindRequiredVersionInFile(filepath string, availableVersions []string) (string, error) {
	bytes, err := os.ReadFile(filepath)
	if err != nil {
		return "", fmt.Errorf("unable to read file: %s", filepath)
	}

	re := regexp.MustCompile(`required_version\s?=\s?"([^"]+)"`)
	match := re.FindStringSubmatch(string(bytes))

	// no required version found, but no real error either
	if len(match) == 0 {
		return "", nil
	}

	// find and return the highest supported version
	for _, v := range availableVersions {
		testVersion, err := version.NewVersion(v)
		if err != nil {
			return "", err
		}

		constraints, err := version.NewConstraint(match[1])
		if err != nil {
			return "", err
		}

		if constraints.Check(testVersion) {
			return v, nil
		}
	}

	// no required version found, but no real error either
	return "", nil
}

// FindTerraformFiles finds all .tf files in the current directory (module)
func FindTerraformFiles() ([]string, error) {
	var files []string
	err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		if strings.HasSuffix(path, ".tf") {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		return files, fmt.Errorf("error finding Terraform files: %s", err)
	}
	return files, nil
}
