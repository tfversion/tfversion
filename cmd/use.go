package cmd

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/tfversion/tfversion/pkg/use"
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
		"tfversion use --required"
)

var (
	useCmd = &cobra.Command{
		Use:     "use",
		Short:   "Activates a given Terraform version",
		Example: useExample,
		PreRun: func(cmd *cobra.Command, args []string) {
			if preRelease && !latest {
				cmd.MarkFlagRequired("latest")
			}
		},
		Run: func(cmd *cobra.Command, args []string) {

			// use latest version
			if latest {
				if len(args) != 0 {
					fmt.Println("error: `--latest` flag does not require specifying a Terraform version")
					fmt.Printf("See %s for help and examples\n", color.BlueString("`tfversion install -h`"))
					os.Exit(1)
				}
				use.UseLatestVersion(preRelease)
				os.Exit(0)
			}

			// use required version
			if required {
				if len(args) != 0 {
					fmt.Println("error: `--required` flag does not require specifying a Terraform version")
					fmt.Printf("See %s for help and examples\n", color.BlueString("`tfversion install -h`"))
					os.Exit(1)
				}
				use.UseRequiredVersion()
				os.Exit(0)
			}

			// use specific version
			if len(args) != 1 {
				fmt.Println("error: provide a Terraform version to activate")
				fmt.Printf("See %s for help and examples\n", color.BlueString("`tfversion install -h`"))
				os.Exit(1)
			}
			use.UseVersion(args[0])
		},
	}
)

func init() {
	rootCmd.AddCommand(useCmd)
	useCmd.Flags().BoolVar(&latest, "latest", false, "use the latest stable Terraform version")
	useCmd.Flags().BoolVar(&preRelease, "pre-release", false, "When used with --latest, use the latest pre-release version")
	useCmd.Flags().BoolVar(&required, "required", false, "use the required Terraform version for your current directory")
}
