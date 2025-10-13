package workstation

import "path/filepath"

func (workstation *Workstation) GetNavFilePath() string {
	workstatonDir := workstation.workingpath
	navFilePath := filepath.Join(workstatonDir, "nav.md")
	return navFilePath
}

func (workstation *Workstation) GetDocsDir() string {
	workstationDir := workstation.workingpath
	docsFolder := filepath.Join(workstationDir, "docs")
	return docsFolder
}

func (workstation *Workstation) GetWorkstationDir() string {
	return workstation.workingpath
}

func (workstation *Workstation) GetAssetsDir() string {
	workstationDir := workstation.workingpath
	assetsFolder := filepath.Join(workstationDir, "assets")
	return assetsFolder
}

func (workstation *Workstation) GetConfigFilePath() string {
	workstationDir := workstation.workingpath
	configFilePath := filepath.Join(workstationDir, "config.json")
	return configFilePath
}

func (workstation *Workstation) GetSearchIndexFilePath() string {
	workstationDir := workstation.workingpath
	searchIndexFilePath := filepath.Join(workstationDir, "search_index.json")
	return searchIndexFilePath
}
