package helpers

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// IsPreReleaseVersion checks if the given version is a Terraform pre-release version
func IsPreReleaseVersion(version string) bool {
	return strings.Contains(version, "-alpha") || strings.Contains(version, "-beta") || strings.Contains(version, "-rc")
}

// FindRequiredVersionInFile finds the required Terraform version in a given .tf file (using required_version = ">= x.x.x")
// TODO: improve logic by using go-version/constraint?
func FindRequiredVersionInFile(filepath string) string {
	bytes, err := os.ReadFile(filepath)
	if err != nil {
		fmt.Printf("Unable to find version number in file: %s", filepath)
		return ""
	}
	re := regexp.MustCompile(`required_version\s*=\s*">=? ?(\d+\.\d+\.\d+)"`)
	match := re.FindStringSubmatch(string(bytes))
	if len(match) == 0 {
		return ""
	}
	return match[1]
}

// FindTerraformFiles finds all .tf files in the current directory (module)
func FindTerraformFiles() []string {
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
		fmt.Println("No Terraform files found in current directory:", err)
		os.Exit(1)
	}
	return files
}
