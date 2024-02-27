package list

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/tfversion/tfversion/pkg/download"
	"github.com/tfversion/tfversion/pkg/helpers"
	"golang.org/x/net/html"
)

// GetInstalledVersions returns the installed Terraform versions from the `~/.tfversion` directory
func GetInstalledVersions() []string {
	installLocation := download.GetDownloadLocation()
	installedVersions, err := os.ReadDir(installLocation)
	if err != nil {
		fmt.Printf("error listing installation directory: %s", err)
		os.Exit(1)
	}

	var versionNames []string
	for _, v := range installedVersions {
		if v.Name() != download.BinaryDir {
			versionNames = append(versionNames, v.Name())
		}
	}

	// Check if there are any versions
	if len(versionNames) == 0 {
		fmt.Println("error listing installed versions: no versions found")
		os.Exit(1)
	}

	// Reverse the order of versionNames to show the latest version first
	reversedVersions := make([]string, len(versionNames))
	for i := 0; i < len(versionNames); i++ {
		reversedVersions[i] = versionNames[len(versionNames)-1-i]
	}

	return reversedVersions
}

// GetAvailableVersions returns the available Terraform versions from the official Terraform releases page
func GetAvailableVersions() []string {
	resp, err := http.Get(download.TerraformReleasesUrl)
	if err != nil {
		fmt.Printf("failed to download Terraform: %s", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		fmt.Printf("failed to parse available versions: %s", err)
		os.Exit(1)
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

// FindLatestVersion finds the latest available Terraform version (or pre-release version)
func FindLatestVersion(preRelease bool) string {
	versions := GetAvailableVersions()
	var foundVersion string
	for _, v := range versions {
		if !preRelease && helpers.IsPreReleaseVersion(v) {
			continue
		}
		foundVersion = v
		break
	}
	if len(foundVersion) == 0 {
		fmt.Println("No versions found")
		os.Exit(1)
	}
	return foundVersion
}
