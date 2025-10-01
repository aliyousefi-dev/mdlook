package mdrepo

// RepoManager handles video registration, thumbnails, previews, etc.
type MDLookManager struct {
	WorkstationDir string
}

// NewMDLookManager creates a new instance of MDLookManager and initializes data storage.
func NewMDLookManager(docDir string) *MDLookManager {
	mdManager := &MDLookManager{
		WorkstationDir: docDir,
	}
	return mdManager
}
