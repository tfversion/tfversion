package list

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/tfversion/tfversion/pkg/alias"
	"github.com/tfversion/tfversion/pkg/api"
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
		helpers.ExitWithError("listing alias directory", err)
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
		err := fmt.Errorf("no versions found")
		helpers.ExitWithError("listing alias directory", err)
	}

	return versionNames
}

// GetInstalledVersions returns the installed Terraform versions from the `~/.tfversion/versions` directory
func GetInstalledVersions() []string {
	installLocation := download.GetDownloadLocation()
	installedVersions, err := os.ReadDir(installLocation)
	if err != nil {
		helpers.ExitWithError("listing versions directory", err)
	}

	var versionNames []string
	for _, v := range installedVersions {
		versionNames = append(versionNames, v.Name())
	}

	// check if there are any versions
	if len(versionNames) == 0 {
		err := fmt.Errorf("no versions found")
		helpers.ExitWithError("listing installed versions", err)
	}

	// Reverse the order of versionNames to show the latest version first
	reversedVersions := make([]string, len(versionNames))
	for i := 0; i < len(versionNames); i++ {
		reversedVersions[i] = versionNames[len(versionNames)-1-i]
	}

	return reversedVersions
}

func GetAvailableVersionsFromApi(maxResults int) []string {
	return api.ListVersions(maxResults)
}

// GetAvailableVersions returns the available Terraform versions from the official Terraform releases page
func GetAvailableVersions() []string {
	resp, err := http.Get(download.TerraformReleasesUrl)
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
		err := fmt.Errorf("no versions found")
		helpers.ExitWithError("finding latest version", err)
	}

	return foundVersion
}
