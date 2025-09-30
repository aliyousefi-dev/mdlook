package mdrepo

import (
	"sync"
	"time"
)

// RepoManager handles video registration, thumbnails, previews, etc.
type MDLookManager struct {
	WorkstationDir string
	lastEventTimes map[string]time.Time
	mu             sync.Mutex
}

// NewMDLookManager creates a new instance of MDLookManager and initializes data storage.
func NewMDLookManager(docDir string) *MDLookManager {
	mdManager := &MDLookManager{
		WorkstationDir: docDir,
	}
	return mdManager
}
