package cmd

import (
	"github.com/spf13/cobra"

	"github.com/tfversion/tfversion/pkg/alias"
	"github.com/tfversion/tfversion/pkg/helpers"
)

const (
	unaliasExample = "# Un-alias a Terraform version\n" +
		"tfversion unalias default\n" +
		"tfversion unalias legacy"
)

var (
	unaliasCmd = &cobra.Command{
		Use:     "unalias",
		Short:   "Un-alias a Terraform version",
		Example: unaliasExample,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) != 1 {
				err := helpers.ErrorWithHelp("tfversion unalias -h")
				helpers.ExitWithError("provide an alias name", err)
			}
			alias.UnaliasVersion(args[0])
		},
	}
)

func init() {
	rootCmd.AddCommand(unaliasCmd)
}
