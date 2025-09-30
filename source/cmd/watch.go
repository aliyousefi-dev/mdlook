package cmd

import (
	"fmt"
	"mdlook/source/internal/mdrepo"
	"mdlook/source/internal/servers"
	"path/filepath"

	"github.com/spf13/cobra"
)

var watchCmd = &cobra.Command{
	Use:   "watch [path]",
	Short: "Watch for changes in the specified folder and auto-reload",
	Args:  cobra.ExactArgs(1), // One argument expected for the path
	Run: func(cmd *cobra.Command, args []string) {

		// Convert the servePath to an absolute path
		absoluteServePath, err := filepath.Abs(args[0])
		if err != nil {
			fmt.Println("Failed to get absolute path:", err)
			return
		}

		// Convert the backslashes to forward slashes
		absoluteServePath = filepath.ToSlash(absoluteServePath)

		// Create a new instance of MDLookManager with the specified directory
		mdManager := mdrepo.NewMDLookManager(absoluteServePath)

		devServer := servers.NewDevServer("localhost", "8080", mdManager)

		devServer.Start()
	},
}

func InitCommandWatch(rootCmd *cobra.Command) {
	rootCmd.AddCommand(watchCmd)
}
