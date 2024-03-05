package list

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/tfversion/tfversion/pkg/alias"
	"github.com/tfversion/tfversion/pkg/download"
	"github.com/tfversion/tfversion/pkg/helpers"
	"golang.org/x/net/html"
)

// GetAliasedVersions returns the aliased Terraform versions from the `~/.tfversion/aliases` directory
func GetAliasedVersions() []string {
	aliasLocation := alias.GetAliasLocation()

	// find all aliases
	aliasedVersions, err := os.ReadDir(aliasLocation)
	if err != nil {
		fmt.Printf("error listing alias directory: %s", err)
		os.Exit(1)
	}

	// resolve the symlinks to get the target versions
	var versionNames []string
	for _, v := range aliasedVersions {
		resolvePath, _ := filepath.EvalSymlinks(filepath.Join(aliasLocation, v.Name()))
		_, targetVersion := filepath.Split(resolvePath)
		versionNames = append(versionNames, fmt.Sprintf("%s -> %s", v.Name(), targetVersion))
	}

	// check if there are any versions
	if len(versionNames) == 0 {
		fmt.Println("error listing installed versions: no versions found")
		os.Exit(1)
	}

	return versionNames
}

// GetInstalledVersions returns the installed Terraform versions from the `~/.tfversion/versions` directory
func GetInstalledVersions() []string {
	installLocation := download.GetDownloadLocation()
	installedVersions, err := os.ReadDir(installLocation)
	if err != nil {
		fmt.Printf("error listing versions directory: %s", err)
		os.Exit(1)
	}

	var versionNames []string
	for _, v := range installedVersions {
		versionNames = append(versionNames, v.Name())
	}

	// check if there are any versions
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
