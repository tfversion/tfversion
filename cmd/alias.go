package cmd

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/spf13/cobra"

	"github.com/tfversion/tfversion/pkg/alias"
	"github.com/tfversion/tfversion/pkg/helpers"
)

const (
	aliasExample = "# Alias a Terraform version\n" +
		"tfversion alias default 1.7.4\n" +
		"tfversion alias legacy 1.2.4"
)

var (
	aliasCmd = &cobra.Command{
		Use:     "alias",
		Short:   "Alias a Terraform version",
		Example: aliasExample,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) != 2 {
				err := fmt.Errorf("see %s for help and examples", color.CyanString("`tfversion alias -h`"))
				helpers.ExitWithError("provide an alias name and Terraform version", err)
			}
			alias.AliasVersion(args[0], args[1])
		},
	}
)

func init() {
	rootCmd.AddCommand(aliasCmd)
}
