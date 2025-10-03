package cmd

import (
	"fmt"
	"mdlook/source/internal/mdrepo"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var exportCmd = &cobra.Command{
	Use:   "export",
	Short: "Export the project",
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

		// Create a new instance of MDLookManager with the specified directory
		mdManager := mdrepo.NewMDLookManager(targetPath)

		mdManager.Export()
	},
}

func InitCommandExport(rootCmd *cobra.Command) {
	// Add the exportCmd to the root command
	rootCmd.AddCommand(exportCmd)
}
