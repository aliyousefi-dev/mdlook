package mdrepo

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// Function that walks through the docs folder and returns the formatted markdown string
func (mdlook *MDLookManager) GenerateNavFileContent(docsPath string) (string, error) {
	var navContent strings.Builder

	docFolderPath := mdlook.GetDocsFolderPath()

	// Walk through the docs folder to find all .md files
	err := filepath.Walk(docFolderPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Only consider .md files
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".md") {
			relativePath := strings.TrimPrefix(path, docFolderPath+"/") // relative path to the docs folder
			relativePath = filepath.ToSlash(relativePath)               // Convert backslashes to forward slashes
			// Remove the .md extension from the file name
			trimmedName := strings.TrimSuffix(info.Name(), ".md")
			// Append the file in markdown list format
			navContent.WriteString(fmt.Sprintf("- [%s](%s)\n", trimmedName, relativePath))
		}
		return nil
	})

	if err != nil {
		return "", fmt.Errorf("error walking the docs folder: %w", err)
	}

	return navContent.String(), nil
}

// WriteNavHeader writes the header to the nav.md file
func (mdlook *MDLookManager) WriteNavHeader(Title string) error {
	// Get the file path dynamically
	navFilePath := mdlook.GetNavFilePath()

	// Open or create the file
	navFile, err := os.OpenFile(navFilePath, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return fmt.Errorf("error opening or creating %s: %w", navFilePath, err)
	}
	defer navFile.Close()

	// Clear the existing content in nav.md
	err = navFile.Truncate(0)
	if err != nil {
		return fmt.Errorf("error clearing %s: %w", navFilePath, err)
	}

	// Write the header to nav.md
	_, err = navFile.WriteString(fmt.Sprintf("# %s\n\n", Title))
	if err != nil {
		return fmt.Errorf("error writing header to %s: %w", navFilePath, err)
	}

	return nil
}

func (mdlook *MDLookManager) LoadNav() string {
	navfilePath := mdlook.GetNavFilePath()

	// Read the content of the nav.md file
	content, err := os.ReadFile(navfilePath)
	if err != nil {
		log.Println("Error reading nav file:", err)
		return ""
	}

	return string(content)
}

func (mdlook *MDLookManager) WriteNav(navContent string) error {
	// Get the file path dynamically
	navFilePath := mdlook.GetNavFilePath()

	// Open or create the file
	navFile, err := os.OpenFile(navFilePath, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return fmt.Errorf("error opening or creating %s: %w", navFilePath, err)
	}
	defer navFile.Close()

	// Clear the existing content in nav.md
	err = navFile.Truncate(0)
	if err != nil {
		return fmt.Errorf("error clearing %s: %w", navFilePath, err)
	}

	// Write the new content to nav.md
	_, err = navFile.WriteString(navContent)
	if err != nil {
		return fmt.Errorf("error writing content to %s: %w", navFilePath, err)
	}

	return nil
}

func (mdlook *MDLookManager) CleanNav() string {
	navContent := mdlook.LoadNav()
	// Remove all lines that start with a dash (i.e., the nav items)
	lines := strings.Split(navContent, "\n")
	var cleanedLines []string
	for _, line := range lines {
		if !strings.HasPrefix(line, "- ") {
			cleanedLines = append(cleanedLines, line)
		}
	}
	return strings.Join(cleanedLines, "\n")
}
