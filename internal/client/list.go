package client

import (
	"fmt"
	"net/http"
	"slices"
	"strings"

	"golang.org/x/net/html"

	"tfversion/internal/helpers"
)

// GetAvailableVersions returns the available Terraform versions from the official Terraform releases page
func ListAvailableVersions() []string {
	resp, err := http.Get(TerraformReleasesUrl)
	if err != nil {
		helpers.ExitWithError("getting Terraform releases page", err)
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		helpers.ExitWithError("parsing HTML", err)
	}

	availableVersions := parseAvailableVersions(doc)
	return availableVersions
}

func parseAvailableVersions(n *html.Node) []string {
	var availableVersions []string

	// find available verions in <a> elements
	if n.Type == html.ElementNode && n.Data == "a" && strings.Contains(n.FirstChild.Data, "terraform") {
		availableVersions = append(availableVersions, strings.Split(n.FirstChild.Data, "_")[1])
	}

	// recursively parse DOM elements
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		availableVersions = append(availableVersions, parseAvailableVersions(c)...)
	}

	return availableVersions
}

// IsAvailableVersion checks if a version is available in the official Terraform releases page
func IsAvailableVersion(version string) bool {
	availableVersions := ListAvailableVersions()
	return slices.Contains(availableVersions, version)
}

// FindLatestVersion finds the latest available Terraform version (or pre-release version)
func FindLatestVersion(preRelease bool) string {
	versions := ListAvailableVersions()
	var foundVersion string
	for _, v := range versions {
		if !preRelease && helpers.IsPreReleaseVersion(v) {
			continue
		}
		foundVersion = v
		break
	}

	if len(foundVersion) == 0 {
		err := fmt.Errorf("no versions found")
		helpers.ExitWithError("finding latest version", err)
	}

	return foundVersion
}
