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
		"tfversion list --aliases\n" +
		"\n" +
		"\n" +
		"# Include pre-release versions\n" +
		"tfversion list --pre-release"
)

var (
	installed                 bool
	aliases                   bool
	maxResults                int
	includePreReleaseVersions bool
	listCmd                   = &cobra.Command{
		Use:     "list",
		Short:   "Lists all Terraform versions",
		Example: listExample,
		PreRun: func(cmd *cobra.Command, args []string) {
			if maxResults < 0 {
				err := helpers.ErrorWithHelp("tfversion list -h")
				helpers.ExitWithError("--max-results cannot be negative", err)
			}
		},
		Run: func(cmd *cobra.Command, args []string) {

			// find the correct type of versions to list
			var versions []string
			if installed {
				versions = list.GetInstalledVersions()
			} else if aliases {
				versions = list.GetAliasedVersions()
			} else {
				versions = list.GetAvailableVersionsFromApi(maxResults)
			}

			// filter out pre-release versions if needed
			var finalList []string
			for _, version := range versions {
				if !helpers.IsPreReleaseVersion(version) || includePreReleaseVersions {
					finalList = append(finalList, version)
				}
			}

			// show the list taking into consideration the max results
			limit := min(maxResults, len(finalList))
			for _, version := range finalList[:limit] {
				fmt.Println(helpers.ColoredVersion(version))
			}
		},
	}
)

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().BoolVar(&installed, "installed", false, "list the installed Terraform versions")
	listCmd.Flags().BoolVar(&aliases, "aliases", false, "list the aliased Terraform versions")
	listCmd.Flags().IntVar(&maxResults, "max-results", 20, "maximum number of versions to list")
	listCmd.Flags().BoolVar(&includePreReleaseVersions, "pre-release", false, "include pre-release versions")
}
