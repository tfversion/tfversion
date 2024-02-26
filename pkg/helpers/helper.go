package helpers

import "strings"

// IsPreReleaseVersion checks if the given version is a Terraform pre-release version
func IsPreReleaseVersion(version string) bool {
	return strings.Contains(version, "-alpha") || strings.Contains(version, "-beta") || strings.Contains(version, "-rc")
}
