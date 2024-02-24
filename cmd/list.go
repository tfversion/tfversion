package cmd

import (
	"github.com/spf13/cobra"
	"github.com/tfversion/tfversion/pkg/list"
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
				list.ListInstalledVersions()
			} else {
				list.ListAvailableVersions()
			}
		},
	}
)

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().BoolVar(&installed, "installed", false, "list the installed Terraform versions")
}
