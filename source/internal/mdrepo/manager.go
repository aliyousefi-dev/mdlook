package mdrepo

import (
	"mdlook/source/internal/renderer"
	"mdlook/source/internal/workstation"
)

// RepoManager handles video registration, thumbnails, previews, etc.
type MDLookManager struct {
	WorkstationDir string
	Workstation    *workstation.Workstation
	Renderer       *renderer.Renderer
}

// NewMDLookManager creates a new instance of MDLookManager and initializes data storage.
func NewMDLookManager(docDir string) *MDLookManager {
	mdManager := &MDLookManager{
		WorkstationDir: docDir,
		Workstation:    workstation.NewWorkstation(docDir),
	}

	mdManager.Renderer = renderer.NewRenderer(mdManager.Workstation)

	return mdManager
}
