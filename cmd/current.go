package cmd

import (
	"github.com/spf13/cobra"

	"tfversion/pkg/current"
)

const (
	currentExample = "# Print the current active version of Terraform\n" +
		"tfversion current"
)

var (
	currentCmd = &cobra.Command{
		Use:     "current",
		Short:   "Print the current active version of Terraform",
		Example: currentExample,
		Run: func(cmd *cobra.Command, args []string) {
			current.CheckCurrentVersion()
		},
	}
)

func init() {
	rootCmd.AddCommand(currentCmd)
}
