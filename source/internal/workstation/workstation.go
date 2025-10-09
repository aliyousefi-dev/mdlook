package workstation

type Workstation struct {
	workingpath string
}

// NewWorkstation creates a new instance of Workstation and initializes data storage.
func NewWorkstation(docDir string) *Workstation {
	workstation := &Workstation{
		workingpath: docDir,
	}
	return workstation
}

func (workstation *Workstation) CreateWorkspace() {
	workstation.GenerateAssetsFolder()
	workstation.GenerateDocsFolder()
	workstation.GenerateDefaultNavFile()
	workstation.GenerateConfigJsonFile()
}
