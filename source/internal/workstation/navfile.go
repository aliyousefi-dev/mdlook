package workstation

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func (workstation *Workstation) LoadNavFile() string {
	navfilePath := workstation.GetNavFilePath()

	// Read the content of the nav.md file
	content, err := os.ReadFile(navfilePath)
	if err != nil {
		log.Println("Error reading nav file:", err)
		return ""
	}

	return string(content)
}

func (workstation *Workstation) UpdateNavFile(navContent string) error {
	// Get the file path dynamically
	navFilePath := workstation.GetNavFilePath()

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

func (workstation *Workstation) CleanNavFile() string {
	navContent := workstation.LoadNavFile()
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
