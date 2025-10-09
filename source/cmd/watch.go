package cmd

import (
	"fmt"
	"mdlook/source/internal/mdrepo"
	"mdlook/source/internal/servers"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var disableSync bool

var watchCmd = &cobra.Command{
	Use:   "watch [path]",
	Short: "Watch for changes in the specified folder and auto-reload",
	Args:  cobra.ExactArgs(1), // One argument expected for the path
	Run: func(cmd *cobra.Command, args []string) {

		// Convert the servePath to an absolute path
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

		// Convert the backslashes to forward slashes
		targetPath = filepath.ToSlash(targetPath)

		// Create a new instance of MDLookManager with the specified directory and disablesync flag
		mdManager := mdrepo.NewMDLookManager(targetPath)

		devServer := servers.NewDevServer("localhost", "8080", mdManager, disableSync)

		devServer.Start()
	},
}

func InitCommandWatch(rootCmd *cobra.Command) {
	// Add disablesync flag
	watchCmd.Flags().BoolVarP(&disableSync, "disablesync", "d", false, "Disable sync feature")
	rootCmd.AddCommand(watchCmd)
}
