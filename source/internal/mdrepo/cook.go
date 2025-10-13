package mdrepo

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func (mdlook *MDLookManager) Cook() {
	mdlook.generateSearchIndex()
}

// generateSearchIndex generates a search index from all markdown files and saves it to a file
func (mdlook *MDLookManager) generateSearchIndex() error {
	// Get all markdown file paths
	mdPaths, err := mdlook.Workstation.ListMdFiles()
	if err != nil {
		return fmt.Errorf("failed to list markdown files: %w", err)
	}

	// Prepare a slice to hold the indexed data
	var searchIndex []SearchIndex

	// Process each markdown file
	for _, path := range mdPaths {
		content, err := os.ReadFile(path) // Updated to os.ReadFile
		if err != nil {
			return fmt.Errorf("failed to read file %s: %w", path, err)
		}

		// Convert path to use forward slashes and remove the ".md" extension
		normalizedPath := filepath.ToSlash(path)
		normalizedPath = strings.TrimSuffix(normalizedPath, ".md")
		normalizedPath = "/" + normalizedPath

		// Extract title and content for the search index
		// Here, we're just using the filename as the title and the whole content as the content
		// You can modify this to extract more specific information if needed (e.g., front matter, headings).
		indexData := SearchIndex{
			Title:   filepath.Base(normalizedPath),
			Content: string(content),
			Path:    normalizedPath,
		}

		// Add to the search index slice
		searchIndex = append(searchIndex, indexData)
	}

	// Save the index to a file
	searchIndexFilePath := mdlook.Workstation.GetSearchIndexFilePath()
	err = mdlook.saveSearchIndexToFile(searchIndexFilePath, searchIndex)
	if err != nil {
		return fmt.Errorf("failed to save search index to file: %w", err)
	}

	return nil
}

// SearchIndex represents the structure for each document in the search index
type SearchIndex struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	Path    string `json:"path"`
}

// saveSearchIndexToFile serializes the search index and writes it to a file
func (mdlook *MDLookManager) saveSearchIndexToFile(filePath string, searchIndex []SearchIndex) error {
	// Marshal the search index to JSON format
	indexData, err := json.MarshalIndent(searchIndex, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal search index: %w", err)
	}

	// Write the JSON data to the file
	err = os.WriteFile(filePath, indexData, 0644) // Updated to os.WriteFile
	if err != nil {
		return fmt.Errorf("failed to write search index to file: %w", err)
	}

	fmt.Printf("Search index saved to: %s\n", filePath)
	return nil
}
