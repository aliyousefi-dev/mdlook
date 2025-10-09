package workstation

import (
	"os"
	"path/filepath"
)

func (workstation *Workstation) ListDocDirs() ([]string, error) {
	var folderPaths []string

	docsFolderPath := workstation.GetDocsDir()

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
