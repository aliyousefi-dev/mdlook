package cmd

import (
	"fmt"
	"mdlook/source/internal/mdrepo"
	"os"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init [path]",
	Short: "Initialize an ova repository in the specified folder",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

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

		mdmanager.Workstation.CreateWorkspace()

		fmt.Println("Initialized empty ova repository.")
	},
}

func InitCommandInit(rootCmd *cobra.Command) {
	// Add the initCmd to the root command
	rootCmd.AddCommand(initCmd)
}
