package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init [path]",
	Short: "Initialize an ova repository in the specified folder",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println("Initialized empty ova repository.")
	},
}

func InitCommandInit(rootCmd *cobra.Command) {
	// Add the initCmd to the root command
	rootCmd.AddCommand(initCmd)
}
