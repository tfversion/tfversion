package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/tfversion/tfversion/pkg/current"
	"github.com/tfversion/tfversion/pkg/helpers"
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
			err := current.CheckCurrentVersion()
			if err != nil {
				helpers.ExitWithError("could not determine active Terraform version", err)
				os.Exit(1)
			}
		},
	}
)

func init() {
	rootCmd.AddCommand(currentCmd)
}
