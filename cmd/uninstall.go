package cmd

import (
	"github.com/spf13/cobra"

	"tfversion/pkg/helpers"
	"tfversion/pkg/uninstall"
)

const (
	uninstallExample = "# uninstall a specific Terraform version\n" +
		"tfversion uninstall 1.7.3" +
		"\n" +
		"\n" +
		"# uninstall the latest stable Terraform version\n" +
		"tfversion uninstall --latest"
)

var (
	uninstallCmd = &cobra.Command{
		Use:     "uninstall",
		Short:   "Uninstalls a given Terraform version",
		Example: uninstallExample,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) != 1 {
				err := helpers.ErrorWithHelp("tfversion uninstall -h")
				helpers.ExitWithError("provide a Terraform version to uninstall", err)
			}
			uninstall.Uninstall(args[0])
		},
	}
)

func init() {
	rootCmd.AddCommand(uninstallCmd)
}
