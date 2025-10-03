package mdrepo

import (
	"encoding/json"
	"log"
	"mdlook/source/internal/types"
	"os"
	"path/filepath"
)

// CreateWorkspace initializes the workspace by creating necessary folders and files with content
func (mdlook *MDLookManager) CreateWorkspace() {
	mdlook.GenerateAssetsFolder()
	mdlook.GenerateDocsFolder()
	mdlook.GenerateDefaultNavFile()
	mdlook.GenerateConfigJsonFile()
}

// GenerateDefaultNavFile creates a default nav.md file in the docs folder
func (mdlook *MDLookManager) GenerateDefaultNavFile() {
	navFilePath := mdlook.GetNavFilePath()

	// Check if the nav file already exists
	if _, err := os.Stat(navFilePath); os.IsNotExist(err) {
		// Create the file if it doesn't exist
		file, err := os.Create(navFilePath)
		if err != nil {
			log.Fatalf("error creating nav file: %v", err)
		}
		defer file.Close()

		// Write default content to the nav.md file
		_, err = file.WriteString("# Documentation\n\n- [introduction](docs/introduction.md)\n")
		if err != nil {
			log.Fatalf("error writing to nav file: %v", err)
		}

		log.Println("Default nav.md file created successfully.")
	} else {
		log.Println("nav.md file already exists.")
	}
}

// GenerateDocsFolder creates the docs folder and adds a default introduction.md file
func (mdlook *MDLookManager) GenerateDocsFolder() {
	docDirPath := mdlook.GetDocsFolderPath()

	// Check if the docs folder exists
	if _, err := os.Stat(docDirPath); os.IsNotExist(err) {
		// Create the docs folder if it doesn't exist
		err := os.MkdirAll(docDirPath, os.ModePerm)
		if err != nil {
			log.Fatalf("error creating docs folder: %v", err)
		}

		log.Println("Docs folder created successfully.")

		// Create a default introduction.md file
		introFilePath := filepath.Join(docDirPath, "introduction.md")
		err = os.WriteFile(introFilePath, []byte("# Introduction\n\nThis is the introduction to the documentation."), 0644)
		if err != nil {
			log.Fatalf("error creating introduction.md: %v", err)
		}

		log.Println("Default introduction.md file created successfully.")
	} else {
		log.Println("Docs folder already exists.")
	}
}

// GenerateAssetsFolder creates the assets folder and adds default placeholder files (e.g., image or CSS)
func (mdlook *MDLookManager) GenerateAssetsFolder() {
	assetsDirPath := mdlook.GetAssetsFolder()

	// Check if the assets folder exists
	if _, err := os.Stat(assetsDirPath); os.IsNotExist(err) {
		// Create the assets folder if it doesn't exist
		err := os.MkdirAll(assetsDirPath, os.ModePerm)
		if err != nil {
			log.Fatalf("error creating assets folder: %v", err)
		}

		log.Println("Assets folder created successfully.")

		// Create a placeholder image (e.g., "logo.png")
		assetsFilePath := filepath.Join(assetsDirPath, "logo.png")
		// Just creating an empty file as a placeholder (you can replace this with an actual image later)
		_, err = os.Create(assetsFilePath)
		if err != nil {
			log.Fatalf("error creating logo.png in assets folder: %v", err)
		}

		log.Println("Placeholder logo.png file created in assets folder.")
	} else {
		log.Println("Assets folder already exists.")
	}
}

func (mdlook *MDLookManager) GenerateConfigJsonFile() {
	configPath := mdlook.GetConfigFilePath()

	// Get the default config data
	defaultConfig := types.GetDefaultConfigData()

	// Marshal the config data into JSON
	configData, err := json.MarshalIndent(defaultConfig, "", "  ")
	if err != nil {
		log.Fatalf("error marshaling config data: %v", err)
	}

	// Write the JSON data to the file
	err = os.WriteFile(configPath, configData, 0644)
	if err != nil {
		log.Fatalf("error creating config.json: %v", err)
	}

	log.Println("Default config.json file created successfully.")
}
