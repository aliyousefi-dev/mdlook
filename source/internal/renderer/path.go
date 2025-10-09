package renderer

import (
	"os"
	"path/filepath"
)

func (renderer *Renderer) GetAllPaths() ([]string, error) {
	var allFilesAndFolders []string
	docsFolderPath := renderer.workstation.GetDocsDir()

	// Walk through the docs folder and all subfolders
	err := filepath.Walk(docsFolderPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err // if there's an error accessing the file/folder, return it
		}

		targetPath := filepath.ToSlash(path)
		// Append the path to the list (both files and directories)
		allFilesAndFolders = append(allFilesAndFolders, targetPath)
		return nil
	})

	if err != nil {
		return nil, err // if an error occurred, return it
	}

	return allFilesAndFolders, nil
}
