package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var watchCmd = &cobra.Command{
	Use:   "watch",
	Short: "Watch for changes in the specified folder",
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println("Watching for changes...")
	},
}

func InitCommandWatch(rootCmd *cobra.Command) {
	// Add the watchCmd to the root command
	rootCmd.AddCommand(watchCmd)
}
