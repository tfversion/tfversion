package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/tfversion/tfversion/pkg/helpers"
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
				_ = cmd.MarkFlagRequired("latest")
			}
		},
		Run: func(cmd *cobra.Command, args []string) {

			// install latest
			if latest {
				if len(args) != 0 {
					err := helpers.ErorWithHelp("tfversion install -h")
					helpers.ExitWithError("`--latest` flag does not require specifying a Terraform version", err)
				}
				install.InstallLatestVersion(preRelease)
				os.Exit(0)
			}

			// installed required version
			if required {
				if len(args) != 0 {
					err := helpers.ErorWithHelp("tfversion install -h")
					helpers.ExitWithError("`--required` flag does not require specifying a Terraform version", err)
				}
				install.InstallRequiredVersion()
				os.Exit(0)
			}

			// install specific version
			if len(args) != 1 {
				err := helpers.ErorWithHelp("tfversion install -h")
				helpers.ExitWithError("provide a Terraform version to install", err)
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
