package cmd

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/tfversion/tfversion/pkg/install"
)

const (
	installExample = "# Install a specific Terraform version\n" +
		"tfversion install 1.7.3" +
		"\n" +
		"\n" +
		"# Install the latest stable Terraform version\n" +
		"tfversion install --latest" +
		"\n" +
		"\n" +
		"# Install the latest pre-release Terraform version\n" +
		"tfversion install --latest --pre-release\n" +
		"\n" +
		"\n" +
		"# Install the required Terraform version for your current directory\n" +
		"tfversion install --required"
)

var (
	latest     bool
	preRelease bool
	required   bool
	installCmd = &cobra.Command{
		Use:     "install",
		Short:   "Installs a given Terraform version",
		Example: installExample,
		PreRun: func(cmd *cobra.Command, args []string) {
			if preRelease && !latest {
				cmd.MarkFlagRequired("latest")
			}
		},
		Run: func(cmd *cobra.Command, args []string) {

			// install latest
			if latest {
				if len(args) != 0 {
					fmt.Println("error: `--latest` flag does not require specifying a Terraform version")
					fmt.Printf("See %s for help and examples\n", color.BlueString("`tfversion install -h`"))
					os.Exit(1)
				}
				install.InstallLatestVersion(preRelease)
				os.Exit(0)
			}

			// installed required version
			if required {
				if len(args) != 0 {
					fmt.Println("error: `--required` flag does not require specifying a Terraform version")
					fmt.Printf("See %s for help and examples\n", color.BlueString("`tfversion install -h`"))
					os.Exit(1)
				}
				install.InstallRequiredVersion()
				os.Exit(0)
			}

			// install specific version
			if len(args) != 1 {
				fmt.Println("error: provide a Terraform version to install")
				fmt.Printf("See %s for help and examples\n", color.BlueString("`tfversion install -h`"))
				os.Exit(1)
			}
			install.InstallVersion(args[0])
		},
	}
)

func init() {
	rootCmd.AddCommand(installCmd)
	installCmd.Flags().BoolVar(&latest, "latest", false, "install the latest stable Terraform version")
	installCmd.Flags().BoolVar(&preRelease, "pre-release", false, "When used with --latest, install the latest pre-release version")
	installCmd.Flags().BoolVar(&required, "required", false, "When used with --required, install the minimum required version for the current module")
}
