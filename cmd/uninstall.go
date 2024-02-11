package cmd

import (
	"fmt"
	"os"

	"github.com/bschaatsbergen/tfversion/pkg/uninstall"
	"github.com/spf13/cobra"
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
		Short:   "uninstalls a given Terraform version",
		Example: uninstallExample,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) != 1 {
				fmt.Println("error: provide a Terraform version to uninstall")
				fmt.Println("See 'tfversion uninstall -h' for help and examples")
				os.Exit(1)
			}
			execUninstall(args[0])
		},
	}
)

func init() {
	rootCmd.AddCommand(uninstallCmd)
}

func execUninstall(version string) {
	err := uninstall.DeleteVersionFromDownloadLocation(version)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Terraform %s uninstalled successfully\n", version)
}
