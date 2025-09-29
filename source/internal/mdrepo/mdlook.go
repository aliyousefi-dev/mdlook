package mdrepo

// RepoManager handles video registration, thumbnails, previews, etc.
type MDLookManager struct {
	workstationDir string
}

// NewMDLookManager creates a new instance of MDLookManager and initializes data storage.
func NewMDLookManager(docDir string) *MDLookManager {
	mdManager := &MDLookManager{
		workstationDir: docDir,
	}
	return mdManager
}
