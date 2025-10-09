package cmd

import (
	"fmt"
	"mdlook/source/internal/mdrepo"
	"os"

	"github.com/spf13/cobra"
)

var testCmd = &cobra.Command{
	Use:   "test [path]",
	Short: "Run tests",
	Args:  cobra.MaximumNArgs(1), // Allow 0 or 1 argument (the path)
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

		// Print the navigation tree with dashes
		str := mdmanager.Renderer.JsonRender()

		fmt.Println(str)

	},
}

func InitCommandTest(rootCmd *cobra.Command) {
	// Add the testCmd to the root command
	rootCmd.AddCommand(testCmd)
}
