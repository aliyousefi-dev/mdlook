package mdrepo

import (
	"os"
	"path/filepath"
)

// GetExecutableFolderPath returns the folder path of the currently running executable (ovacli).
func (mdlook *MDLookManager) GetExecutableFolderPath() string {
	exePath, err := os.Executable()
	if err != nil {
		return ""
	}
	exeDir := filepath.Dir(exePath)
	return exeDir
}

// GetWebUIPath returns the folder path of the web UI assets.
// If a "web" folder exists in the executable's directory, returns its path; otherwise, returns the executable folder path.
func (mdlook *MDLookManager) GetWebUIFolderPath() string {
	exeDir := mdlook.GetExecutableFolderPath()
	webDir := filepath.Join(exeDir, "web")
	if stat, err := os.Stat(webDir); err == nil && stat.IsDir() {
		return webDir
	}
	return exeDir
}

func (mdlook *MDLookManager) GetIndexHtmlPath() string {
	webDir := mdlook.GetWebUIFolderPath()
	indexPath := filepath.Join(webDir, "index.html")
	return indexPath
}

func (mdlook *MDLookManager) GetNavFilePath() string {
	workstatonDir := mdlook.WorkstationDir
	navFilePath := filepath.Join(workstatonDir, "nav.md")
	return navFilePath
}

func (mdlook *MDLookManager) GetDocsFolderPath() string {
	workstationDir := mdlook.WorkstationDir
	docsFolder := filepath.Join(workstationDir, "docs")
	return docsFolder
}

func (mdlook *MDLookManager) GetAssetsFolder() string {
	workstationDir := mdlook.WorkstationDir
	assetsFolder := filepath.Join(workstationDir, "assets")
	return assetsFolder
}

func (mdlook *MDLookManager) GetConfigFilePath() string {
	workstationDir := mdlook.WorkstationDir
	configFilePath := filepath.Join(workstationDir, "config.json")
	return configFilePath
}
