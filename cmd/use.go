package cmd

import (
	"fmt"
	"os"

	"github.com/bschaatsbergen/tfversion/pkg/use"
	"github.com/spf13/cobra"
)

const (
	useExample = "# Activate a specific Terraform version\n" +
		"tfversion use 1.7.3"
)

var (
	useCmd = &cobra.Command{
		Use:     "use",
		Short:   "Activates a given Terraform version",
		Example: useExample,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) != 1 {
				fmt.Println("error: provide a Terraform version to install")
				fmt.Println("See 'tfversion install -h' for help and examples")
				os.Exit(1)
			}
			use.UseVersion(args[0])
		},
	}
)

func init() {
	rootCmd.AddCommand(useCmd)
}
