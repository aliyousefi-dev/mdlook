package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var syncCmd = &cobra.Command{
	Use:   "sync",
	Short: "Synchronize files",
	Run: func(cmd *cobra.Command, args []string) {
		// Path to the docs folder
		docsPath := "docs"

		// Open or create nav.md file
		navFilePath := "nav.md"
		navFile, err := os.OpenFile(navFilePath, os.O_RDWR|os.O_CREATE, 0644)
		if err != nil {
			fmt.Println("Error opening nav.md:", err)
			return
		}
		defer navFile.Close()

		// Clear the existing content in nav.md
		err = navFile.Truncate(0)
		if err != nil {
			fmt.Println("Error clearing nav.md:", err)
			return
		}

		// Add header to nav.md
		navFile.WriteString("# Navigation\n\n")

		// Walk through the docs folder to find all .md files
		err = filepath.Walk(docsPath, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				fmt.Println("Error walking the path:", err)
				return err
			}

			// Only consider .md files
			if !info.IsDir() && strings.HasSuffix(info.Name(), ".md") {
				relativePath := strings.TrimPrefix(path, docsPath+"/") // relative path to the docs folder
				relativePath = filepath.ToSlash(relativePath)          // Convert backslashes to forward slashes
				// Remove the .md extension from the file name
				trimmedName := strings.TrimSuffix(info.Name(), ".md")
				// Write the file to nav.md in a markdown list format
				navFile.WriteString(fmt.Sprintf("- [%s](%s)\n", trimmedName, relativePath))
			}
			return nil
		})

		if err != nil {
			fmt.Println("Error walking the docs folder:", err)
			return
		}

		// Print success message
		fmt.Println("Synchronization complete! nav.md has been updated.")
	},
}

func InitCommandSync(rootCmd *cobra.Command) {
	// Add the syncCmd to the root command
	rootCmd.AddCommand(syncCmd)
}
