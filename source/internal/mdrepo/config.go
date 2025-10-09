package mdrepo

import (
	"encoding/json"
	"log"
	"mdlook/source/internal/types"
	"os"
)

func (mdlook *MDLookManager) LoadConfig() types.MdLookConfig {
	// Load the configuration from the config.json file
	configFilePath := mdlook.Workstation.GetConfigFilePath()
	file, err := os.Open(configFilePath)
	if err != nil {
		log.Printf("Error opening config file: %v", err)
		return types.MdLookConfig{} // Return empty config on error
	}
	defer file.Close()

	var config types.MdLookConfig
	decoder := json.NewDecoder(file)

	// Decode the JSON file into the config struct
	if err := decoder.Decode(&config); err != nil {
		log.Printf("Error decoding config file: %v", err)
		return types.MdLookConfig{} // Return empty config on error
	}

	return config
}
