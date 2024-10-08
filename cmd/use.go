package cmd

import (
	"os"

	"github.com/spf13/cobra"

	"tfversion/internal/helpers"
	"tfversion/internal/use"
)

const (
	useExample = "# Use a specific Terraform version\n" +
		"tfversion use 1.7.3" +
		"\n" +
		"\n" +
		"# Use the latest stable Terraform version\n" +
		"tfversion use --latest" +
		"\n" +
		"\n" +
		"# Use the latest pre-release Terraform version\n" +
		"tfversion use --latest --pre-release\n" +
		"\n" +
		"\n" +
		"# Use the required Terraform version for your current directory\n" +
		"tfversion use --required\n" +
		"\n" +
		"\n" +
		"# Automatically install the target version if it is missing\n" +
		"tfversion use 1.7.3 --install"
)

var (
	autoInstall bool
	useCmd      = &cobra.Command{
		Use:     "use",
		Short:   "Activates a given Terraform version",
		Example: useExample,
		PreRun: func(cmd *cobra.Command, args []string) {
			if preRelease && !latest {
				_ = cmd.MarkFlagRequired("latest")
			}
		},
		Run: func(cmd *cobra.Command, args []string) {

			// use latest version
			if latest {
				if len(args) != 0 {
					err := helpers.ErrorWithHelp("tfversion use -h")
					helpers.ExitWithError("`--latest` flag does not require specifying a Terraform version", err)
				}
				use.UseLatestVersion(preRelease, autoInstall)
				os.Exit(0)
			}

			// use required version
			if required {
				if len(args) != 0 {
					err := helpers.ErrorWithHelp("tfversion use -h")
					helpers.ExitWithError("`--required` flag does not require specifying a Terraform version", err)
				}
				use.UseRequiredVersion(autoInstall)
				os.Exit(0)
			}

			// use specific version
			if len(args) != 1 {
				err := helpers.ErrorWithHelp("tfversion use -h")
				helpers.ExitWithError("provide a Terraform version to activate", err)
			}
			use.UseVersion(args[0], autoInstall)
		},
	}
)

func init() {
	rootCmd.AddCommand(useCmd)
	useCmd.Flags().BoolVar(&latest, "latest", false, "use the latest stable Terraform version")
	useCmd.Flags().BoolVar(&preRelease, "pre-release", false, "When used with --latest, use the latest pre-release version")
	useCmd.Flags().BoolVar(&required, "required", false, "use the required Terraform version for your current directory")
	useCmd.Flags().BoolVar(&autoInstall, "install", false, "automatically install if the target version is missing")
}
