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
		"tfversion install 1.7.3"
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
			execCmd(args[0])
		},
	}
)

func init() {
	rootCmd.AddCommand(installCmd)
}

func execCmd(version string) {
	isAlreadyDownloaded := download.IsAlreadyDownloaded(version)
	//TODO: Print a message to the user that the version is already downloaded

	if !isAlreadyDownloaded {
		// Download the Terraform release
		zipFile, err := download.Download(version, runtime.GOOS, runtime.GOARCH)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Unzip the downloaded Terraform release
		err = download.UnzipRelease(zipFile, fmt.Sprintf("/home/bruno/.tfversion/%s", version))
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Clean up the downloaded zip file after unzipping
		err = download.DeleteDownloadedRelease(zipFile)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}
