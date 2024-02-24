package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/tfversion/tfversion/pkg/install"
)

const (
	installExample = "# Install a specific Terraform version\n" +
		"tfversion install 1.7.3" +
		"\n" +
		"\n" +
		"# Install the latest stable Terraform version\n" +
		"tfversion install --latest"
)

var (
	latest     bool
	preRelease bool
	installCmd = &cobra.Command{
		Use:     "install",
		Short:   "Installs a given Terraform version",
		Example: installExample,
		Run: func(cmd *cobra.Command, args []string) {
			if latest {
				if len(args) != 0 {
					fmt.Println("error: 'latest' flag does not require specifying a Terraform version")
					fmt.Println("See 'tfversion install -h' for help and examples")
					os.Exit(1)
				}
				install.InstallVersion("", latest, preRelease)
			} else {
				if len(args) != 1 {
					fmt.Println("error: provide a Terraform version to install")
					fmt.Println("See 'tfversion install -h' for help and examples")
					os.Exit(1)
				}
				install.InstallVersion(args[0], latest, preRelease)
			}
		},
	}
)

func init() {
	rootCmd.AddCommand(installCmd)
	installCmd.Flags().BoolVar(&latest, "latest", false, "install the latest stable Terraform version")
	installCmd.Flags().BoolVar(&preRelease, "pre-release", false, "When used with --latest, install the latest pre-release version")
}
