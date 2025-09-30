package cmd

import (
	"log"
	"mdlook/source/internal/servers"
	"time"

	"github.com/spf13/cobra"
)

// Define the socketCmd which will handle WebSocket functionality
var socketCmd = &cobra.Command{
	Use:   "socket <path>",
	Short: "Watch for changes in the specified folder and auto-reload",
	Run: func(cmd *cobra.Command, args []string) {

		// Create a new WebSocket server
		wsServer := servers.NewWebsocketServer("localhost", "8080")

		// Start the WebSocket server in a separate goroutine
		go func() {
			log.Println("Starting WebSocket server...")
			wsServer.Start()
		}()

		// Set up a ticker that ticks every 5 seconds
		ticker := time.NewTicker(5 * time.Second)
		defer ticker.Stop()

		// Use for-range to receive ticks from ticker
		for range ticker.C {
			log.Println("Broadcasting reload signal...")
			wsServer.BroadcastReload() // Broadcast reload to all connected clients
		}
	},
}

// InitCommandSocket initializes the socket command and adds it to the root command
func InitCommandSocket(rootCmd *cobra.Command) {
	rootCmd.AddCommand(socketCmd)
}
