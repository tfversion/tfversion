package cmd

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/spf13/cobra"

	"github.com/tfversion/tfversion/pkg/alias"
)

const (
	aliasExample = "# Alias a Terraform version\n" +
		"tfversion alias default 1.7.4"
)

var (
	aliasCmd = &cobra.Command{
		Use:     "alias",
		Short:   "Alias a Terraform version",
		Example: aliasExample,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) != 2 {
				fmt.Println("error: provide an alias name and Terraform version")
				fmt.Printf("See %s for help and examples\n", color.CyanString("`tfversion alias -h`"))
				os.Exit(1)
			}
			alias.AliasVersion(args[0], args[1])
		},
	}
)

func init() {
	rootCmd.AddCommand(aliasCmd)
}
