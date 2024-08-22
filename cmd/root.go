package cmd

import (
	"strings"

	"github.com/fatih/color"
	"github.com/spf13/cobra"

	"tfversion/internal/helpers"
)

var (
	version string

	rootCmd = &cobra.Command{
		Use:     "tfversion",
		Short:   "tfversion - A simple tool to manage Terraform versions",
		Version: version, // the version is set during the build by making using of `go build -ldflags`
		Run: func(cmd *cobra.Command, args []string) {
			err := cmd.Help()
			if err != nil {
				helpers.ExitWithError("unable to display help", err)
			}
		},
	}
)

func setupCobraUsageTemplate() {
	cobra.AddTemplateFunc("StyleHeading", color.New(color.FgBlue).SprintFunc())
	usageTemplate := rootCmd.UsageTemplate()
	usageTemplate = strings.NewReplacer(
		`Usage:`, `{{StyleHeading "Usage:"}}`,
		`Examples:`, `{{StyleHeading "Examples:"}}`,
		`Available Commands:`, `{{StyleHeading "Available Commands:"}}`,
		`Flags:`, `{{StyleHeading "Flags:"}}`,
	).Replace(usageTemplate)
	rootCmd.SetUsageTemplate(usageTemplate)
}

func init() {
	setupCobraUsageTemplate()
	rootCmd.CompletionOptions.DisableDefaultCmd = true
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		helpers.ExitWithError("unable to execute command", err)
	}
}
