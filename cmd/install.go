package cmd

import (
	"fmt"
	"os"
	"runtime"

	"github.com/bschaatsbergen/tfversion/pkg/download"
	"github.com/spf13/cobra"
)

const (
	installExample = "# Install a specific Terraform version\n" +
		"tfversion install 1.7.1"
)

var (
	installCmd = &cobra.Command{
		Use:     "install",
		Short:   "Installs a given Terraform version",
		Example: installExample,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) != 1 {
				fmt.Println("error: provide a Terraform version to install")
				fmt.Println("See 'tfversion install -h' for help and examples")
				os.Exit(1)
			}
			install(args[0])
		},
	}
)

func init() {
	rootCmd.AddCommand(installCmd)
}

func install(version string) {
	download.DownloadTerraform(version, runtime.GOOS, runtime.GOARCH)
}
