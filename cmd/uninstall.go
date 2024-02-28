package cmd

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
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
				fmt.Println("error: provide a Terraform version to uninstall")
				fmt.Printf("See %s for help and examples\n", color.CyanString("`tfversion install -h`"))
				os.Exit(1)
			}
			uninstall.Uninstall(args[0])
		},
	}
)

func init() {
	rootCmd.AddCommand(uninstallCmd)
}
