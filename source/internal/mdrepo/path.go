package mdrepo

import (
	"os"
	"path/filepath"
)

// GetMdlookExeDir returns the folder path of the currently running executable (ovacli).
func (mdlook *MDLookManager) GetMdlookExeDir() string {
	exePath, err := os.Executable()
	if err != nil {
		return ""
	}
	exeDir := filepath.Dir(exePath)
	return exeDir
}

// GetWebUIPath returns the folder path of the web UI assets.
// If a "web" folder exists in the executable's directory, returns its path; otherwise, returns the executable folder path.
func (mdlook *MDLookManager) GetWebTemplateDir() string {
	exeDir := mdlook.GetMdlookExeDir()
	webDir := filepath.Join(exeDir, "web")
	if stat, err := os.Stat(webDir); err == nil && stat.IsDir() {
		return webDir
	}
	return exeDir
}

func (mdlook *MDLookManager) GetIndexHtmlPath() string {
	webDir := mdlook.GetWebTemplateDir()
	indexPath := filepath.Join(webDir, "index.html")
	return indexPath
}




