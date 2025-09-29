package cmd

import (
	"fmt"
	"log"
	"mdlook/source/internal/mdrepo"
	"path/filepath"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/cobra"
)

var watchCmd = &cobra.Command{
	Use:   "watch [path]",
	Short: "Watch for changes in the specified folder and auto-reload",
	Args:  cobra.ExactArgs(1), // One argument expected for the path
	Run: func(cmd *cobra.Command, args []string) {
		servePath := args[0]
		fmt.Println("Watching for changes in:", servePath)

		// Convert the servePath to an absolute path
		absoluteServePath, err := filepath.Abs(servePath)
		if err != nil {
			fmt.Println("Failed to get absolute path:", err)
			return
		}

		// Initialize the MDLookManager
		mdManager := mdrepo.NewMDLookManager(absoluteServePath)

		// Start the server in a separate goroutine
		go mdManager.StartServer("127.0.0.1:8080", absoluteServePath)

		// Create a new file watcher
		watcher, err := fsnotify.NewWatcher()
		if err != nil {
			log.Fatal("Error creating file watcher:", err)
		}
		defer watcher.Close()

		// Add the servePath and all subdirectories to the watcher
		err = watcher.Add(absoluteServePath)
		if err != nil {
			log.Fatal("Error watching the path:", err)
		}

		// Start watching for changes
		fmt.Println("Watching for changes...")

		for {
			select {
			case event := <-watcher.Events:
				if event.Op&fsnotify.Write == fsnotify.Write {
					fmt.Printf("Change detected: %s\n", event.Name)
					// Trigger rebuild or refresh logic here
					// For example, you can restart the server or just notify the browser to reload
					triggerReload(mdManager)
				}

			case err := <-watcher.Errors:
				fmt.Println("Watcher error:", err)
			}
		}
	},
}

func triggerReload(mdManager *mdrepo.MDLookManager) {
	// Here you can trigger a reload of the content or notify the server to reload the page.
	// For example, you can use WebSockets, signal a full page reload, or rebuild the content.
	fmt.Println("Triggering reload or rebuild...")
	// You can call a method here to trigger a refresh in the server or frontend.
	// mdManager.ReloadContent() or any other method
}

func InitCommandWatch(rootCmd *cobra.Command) {
	// Add the watchCmd to the root command
	rootCmd.AddCommand(watchCmd)
}
