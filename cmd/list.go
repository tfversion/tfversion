package cmd

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/tfversion/tfversion/pkg/list"
)

const (
	listExample = "# List all available Terraform versions\n" +
		"tfversion list" +
		"\n" +
		"\n" +
		"# List more available versions\n" +
		"tfversion list --max-results=20\n" +
		"\n" +
		"\n" +
		"# List all installed Terraform versions\n" +
		"tfversion list --installed"
)

var (
	installed  bool
	maxResults int
	listCmd    = &cobra.Command{
		Use:     "list",
		Short:   "Lists all Terraform versions",
		Example: listExample,
		Run: func(cmd *cobra.Command, args []string) {
			if installed {
				installedVersions := list.GetInstalledVersions()
				for _, version := range installedVersions {
					fmt.Println(color.BlueString(version))
				}
			} else {
				availableVersions := list.GetAvailableVersions()
				for _, v := range availableVersions[:maxResults] {
					fmt.Println(color.BlueString(v))
				}
			}
		},
	}
)

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().BoolVar(&installed, "installed", false, "list the installed Terraform versions")
	listCmd.Flags().IntVar(&maxResults, "max-results", 500, "maximum number of versions to list")
}
