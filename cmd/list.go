package cmd

import (
	"fmt"
	"net/http"
	"os"

	"github.com/bschaatsbergen/tfversion/pkg/download"
	"github.com/spf13/cobra"
	"golang.org/x/net/html"
)

const (
	listExample = "# List all available Terraform versions\n" +
		"tfversion list" +
		"\n" +
		"\n" +
		"# List all installed Terraform versions\n" +
		"tfversion list --installed"
)

var (
	installed bool
	listCmd   = &cobra.Command{
		Use:     "list",
		Short:   "Lists all Terraform versions",
		Example: listExample,
		Run: func(cmd *cobra.Command, args []string) {
			if installed {
				execListInstalled()
			} else {
				execList()
			}
		},
	}
)

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().BoolVar(&installed, "installed", false, "list the installed Terraform versions")
}

func execList() {
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

	var availableVersions []string
	var processAllVersions func(*html.Node)
	processAllVersions = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			availableVersions = append(availableVersions, n.FirstChild.Data)
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			processAllVersions(c)
		}
	}

	processAllVersions(doc)

	for _, v := range availableVersions {
		fmt.Println(v)
	}
}

func execListInstalled() {
	installLocation := download.GetDownloadLocation()
	installedVersions, err := os.ReadDir(installLocation)
	if err != nil {
		fmt.Printf("error listing installation directory: %s", err)
		os.Exit(1)
	}
	for _, v := range installedVersions {
		fmt.Println(v.Name())
	}
}
