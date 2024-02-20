package cmd

import (
	"fmt"
	"os"

	"github.com/bschaatsbergen/tfversion/pkg/download"
	"github.com/spf13/cobra"
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
	// TODO: parse content from download.TerraformReleasesUrl
}

func execListInstalled() {
	installLocation := download.GetDownloadLocation()
	installedVersions, err := os.ReadDir(installLocation)
	if err != nil {
		fmt.Printf("error listing installation directory: %s", err)
	}
	for _, v := range installedVersions {
		fmt.Println(v.Name())
	}
}
