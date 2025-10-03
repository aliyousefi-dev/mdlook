package mdrepo

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

// CopyFile copies a single file from src to dst.
func CopyFile(src, dst string) error {
	// Open source file
	sourceFile, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("failed to open source file %s: %v", src, err)
	}
	defer sourceFile.Close()

	// Create destination file
	destinationFile, err := os.Create(dst)
	if err != nil {
		return fmt.Errorf("failed to create destination file %s: %v", dst, err)
	}
	defer destinationFile.Close()

	// Copy contents
	_, err = io.Copy(destinationFile, sourceFile)
	if err != nil {
		return fmt.Errorf("failed to copy file contents from %s to %s: %v", src, dst, err)
	}

	return nil
}

// CopyDir recursively copies the contents of a directory from src to dst.
func CopyDir(src, dst string) error {
	// Open source directory
	sourceDir, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("failed to open source directory %s: %v", src, err)
	}
	defer sourceDir.Close()

	// Get list of entries in the directory
	entries, err := sourceDir.Readdir(-1)
	if err != nil {
		return fmt.Errorf("failed to read directory %s: %v", src, err)
	}

	// Ensure destination directory exists
	if err := os.MkdirAll(dst, os.ModePerm); err != nil {
		return fmt.Errorf("failed to create destination directory %s: %v", dst, err)
	}

	// Copy each entry (file or directory)
	for _, entry := range entries {
		sourcePath := filepath.Join(src, entry.Name())
		destPath := filepath.Join(dst, entry.Name())

		if entry.IsDir() {
			// Recursively copy subdirectories
			if err := CopyDir(sourcePath, destPath); err != nil {
				return err
			}
		} else {
			// Copy files
			if err := CopyFile(sourcePath, destPath); err != nil {
				return err
			}
		}
	}

	return nil
}

func (mdlook *MDLookManager) Export() {
	webUIPath := mdlook.GetWebUIFolderPath()
	workstationDir := mdlook.GetWorkstationDir()

	// Define the export directory path
	exportDir := filepath.Join(workstationDir, "export")

	// Copy everything from webUIPath to exportDir
	if err := CopyDir(webUIPath, exportDir); err != nil {
		log.Fatalf("failed to copy files from webUIPath to exportDir: %v", err)
	}

	// List of specific files and directories to copy from workstationDir
	filesToCopy := []string{
		filepath.Join(workstationDir, "docs"),
		filepath.Join(workstationDir, "asset"),
		filepath.Join(workstationDir, "nav.md"),
		filepath.Join(workstationDir, "config.json"),
	}

	// Copy each specified file/folder to exportDir
	for _, filePath := range filesToCopy {
		destPath := filepath.Join(exportDir, filepath.Base(filePath))
		if stat, err := os.Stat(filePath); err == nil {
			if stat.IsDir() {
				// Copy directories
				if err := CopyDir(filePath, destPath); err != nil {
					log.Printf("failed to copy directory %s: %v", filePath, err)
				}
			} else {
				// Copy individual files
				if err := CopyFile(filePath, destPath); err != nil {
					log.Printf("failed to copy file %s: %v", filePath, err)
				}
			}
		}
	}
}
