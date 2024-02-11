package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

const (
	uninstallExample = "# Uninstall a specific Terraform version\n" +
		"tfversion uninstall 1.7.1"
)

var (
	uninstallCmd = &cobra.Command{
		Use:     "uninstall",
		Short:   "Uninstalls a given Terraform version",
		Example: uninstallExample,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) != 1 {
				fmt.Println("error: provide a Terraform version to uninstall")
				fmt.Println("See 'tfversion uninstall -h' for help and examples")
				os.Exit(1)
			}
			uninstall()
		},
	}
)

func init() {
	rootCmd.AddCommand(uninstallCmd)
}

func uninstall() {
}
