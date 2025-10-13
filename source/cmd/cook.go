package cmd

import (
	"fmt"
	"mdlook/source/internal/mdrepo"
	"os"

	"github.com/spf13/cobra"
)

var cookCmd = &cobra.Command{
	Use:   "cook [path]",
	Short: "Cook the application from the specified folder",
	Args:  cobra.ExactArgs(1), // One argument expected
	Run: func(cmd *cobra.Command, args []string) {
		// Determine the directory path to use
		var targetPath string
		if len(args) > 0 {
			// If a path argument is provided, use it
			targetPath = args[0]
		} else {
			// If no argument is provided, use the current working directory
			var err error
			targetPath, err = os.Getwd()
			if err != nil {
				fmt.Println("Failed to get working directory:", err)
				return
			}
		}

		// Initialize the MDLookManager with the target path
		mdmanager := mdrepo.NewMDLookManager(targetPath)

		mdmanager.Cook()

	},
}

func InitCommandCook(rootCmd *cobra.Command) {
	// Add the cookCmd to the root command
	rootCmd.AddCommand(cookCmd)
}
