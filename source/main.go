package main

import (
	"mdlook/source/cmd"

	"github.com/spf13/cobra"
)

func main() {

	var rootCmd = &cobra.Command{
		Use:   "mdlook",
		Short: "MDLook — a command-line tool to create HTML docs from Markdown",
		Long: `MDLook provides a command-line interface for converting Markdown files to HTML documentation,
with additional features for organizing, processing, and managing your documentation workflow.`,
	}

	// common commands
	cmd.InitCommandInit(rootCmd)
	cmd.InitCommandServe(rootCmd)
	cmd.InitCommandVersion(rootCmd)
	cmd.InitCommandSync(rootCmd)
	cmd.InitCommandWatch(rootCmd)
	cmd.InitCommandSocket(rootCmd)

	rootCmd.Execute()
	// Initialize the root command and add subcommands
}
