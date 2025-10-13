package workstation

import (
	"os"
	"path/filepath"
	"strings"
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

// ListMdFiles returns a slice of paths for all .md files within the docs directory
func (workstation *Workstation) ListMdFiles() ([]string, error) {
	var mdFilePaths []string

	docsFolderPath := workstation.GetDocsDir()

	// Walk through the docs folder and collect all .md file paths
	err := filepath.Walk(docsFolderPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Check if the file has a .md extension and add it to the list
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".md") {
			mdFilePaths = append(mdFilePaths, path)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return mdFilePaths, nil
}
