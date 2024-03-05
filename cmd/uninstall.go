package cmd

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/tfversion/tfversion/pkg/helpers"
	"github.com/tfversion/tfversion/pkg/uninstall"
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
				err := fmt.Errorf("see %s for help and examples", color.CyanString("`tfversion uninstall -h`"))
				helpers.ExitWithError("provide a Terraform version to uninstall", err)
			}
			uninstall.Uninstall(args[0])
		},
	}
)

func init() {
	rootCmd.AddCommand(uninstallCmd)
}
