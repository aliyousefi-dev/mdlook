package mdrepo

// RepoManager handles video registration, thumbnails, previews, etc.
type MDLookManager struct {
	docDir string
}

// NewMDLookManager creates a new instance of MDLookManager and initializes data storage.
func NewMDLookManager(docDir string) (*MDLookManager, error) {
	mdManager := &MDLookManager{
		docDir: docDir,
	}
	return mdManager, nil
}
