package mdrepo

import (
	"log"
	"mdlook/source/internal/types"
	"os"
	"path/filepath"
	"strings"
)

// ScanDirectory scans the docs folder, finds all .md files, and generates a NavRenderStruct
func (mdlook *MDLookManager) ScanDirectory() types.NavRenderStruct {
	docsFolderPath := mdlook.GetDocsFolderPath()

	var navItems []types.NavItem

	// Walk through the docs folder to find all .md files
	err := filepath.Walk(docsFolderPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Println("Error walking the path:", err)
			return err
		}

		// Only consider .md files
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".md") {
			// Get the relative path to the docs folder
			relativePath, err := filepath.Rel(mdlook.WorkstationDir, path)
			if err != nil {
				log.Println("Error getting relative path:", err)
				return err
			}

			// Convert backslashes to forward slashes for web compatibility
			relativePath = filepath.ToSlash(relativePath)

			// Remove the .md extension from the file name for the title
			title := strings.TrimSuffix(info.Name(), ".md")

			// Create NavItem using NewNavItem
			navItem := types.NewNavItem(title, relativePath)
			navItems = append(navItems, navItem)
		}
		return nil
	})

	if err != nil {
		log.Fatalf("Error scanning docs folder: %v", err)
	}

	// Create and return the NavRenderStruct with relative paths
	return types.NewNavRender("Table of Contents", navItems)
}

func (mdlook *MDLookManager) GetAllFolderPaths() ([]string, error) {
	var folderPaths []string

	docsFolderPath := mdlook.GetDocsFolderPath()

	// Walk through the docs folder and collect all directory paths
	err := filepath.Walk(docsFolderPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Add the path to the list if it is a directory
		if info.IsDir() {
			folderPaths = append(folderPaths, path)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return folderPaths, nil
}
