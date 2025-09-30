package cmd

import (
	"fmt"
	"mdlook/source/internal/mdrepo"
	"mdlook/source/internal/servers"
	"path/filepath"

	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use:   "serve [path]",
	Short: "Serve the application from the specified folder",
	Args:  cobra.ExactArgs(1), // One argument expected
	Run: func(cmd *cobra.Command, args []string) {
		servePath := args[0] // The first argument is the serve path
		fmt.Println("Serving application from:", servePath)

		// Convert the servePath to an absolute path
		absoluteServePath, err := filepath.Abs(servePath)
		if err != nil {
			fmt.Println("Failed to get absolute path:", err)
			return
		}

		// Create a new instance of MDLookManager with the specified directory
		mdManager := mdrepo.NewMDLookManager(absoluteServePath)

		server := servers.NewWebServer("127.0.0.1", "8080", mdManager, false)

		// Start the server with both the primary and secondary paths (absolute paths)
		server.Start()
	},
}

func InitCommandServe(rootCmd *cobra.Command) {
	// Add the serveCmd to the root command
	rootCmd.AddCommand(serveCmd)
}
