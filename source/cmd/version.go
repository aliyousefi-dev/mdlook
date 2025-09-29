package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number",
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println("Version 1.0.0")
	},
}

func InitCommandVersion(rootCmd *cobra.Command) {
	// Add the versionCmd to the root command
	rootCmd.AddCommand(versionCmd)
}
