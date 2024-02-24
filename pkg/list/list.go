package list

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/bschaatsbergen/tfversion/pkg/download"
	"golang.org/x/net/html"
)

func ListInstalledVersions() {
	installLocation := download.GetDownloadLocation()
	installedVersions, err := os.ReadDir(installLocation)
	if err != nil {
		fmt.Printf("error listing installation directory: %s", err)
		os.Exit(1)
	}
	for _, v := range installedVersions {
		if v.Name() != download.BinaryDir {
			fmt.Println(v.Name())
		}
	}
}

func ListAvailableVersions() {
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
	for _, v := range availableVersions {
		fmt.Println(v)
	}
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
