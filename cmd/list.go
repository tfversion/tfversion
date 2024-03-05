package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tfversion/tfversion/pkg/helpers"
	"github.com/tfversion/tfversion/pkg/list"
)

const (
	listExample = "# List all available Terraform versions\n" +
		"tfversion list" +
		"\n" +
		"\n" +
		"# Limit the number of results\n" +
		"tfversion list --max-results=20\n" +
		"\n" +
		"\n" +
		"# List all installed Terraform versions\n" +
		"tfversion list --installed\n" +
		"\n" +
		"\n" +
		"# List all aliased Terraform versions\n" +
		"tfversion list --aliases"
)

var (
	installed  bool
	aliases    bool
	maxResults int
	listCmd    = &cobra.Command{
		Use:     "list",
		Short:   "Lists all Terraform versions",
		Example: listExample,
		Run: func(cmd *cobra.Command, args []string) {

			var versions []string
			if installed {
				versions = list.GetInstalledVersions()
			} else if aliases {
				versions = list.GetAliasedVersions()
			} else {
				versions = list.GetAvailableVersions()
			}

			limit := min(maxResults, len(versions))
			for _, version := range versions[:limit] {
				fmt.Println(helpers.ColoredVersion(version))
			}
		},
	}
)

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().BoolVar(&installed, "installed", false, "list the installed Terraform versions")
	listCmd.Flags().BoolVar(&aliases, "aliases", false, "list the aliased Terraform versions")
	listCmd.Flags().IntVar(&maxResults, "max-results", 500, "maximum number of versions to list")
}
