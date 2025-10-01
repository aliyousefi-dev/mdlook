package watcher

import "time"

// Watcher struct to hold watcher-related data and methods.
type WorkstationWatcher struct {
	SearchPath    []string
	lastEventTime time.Time // New field to store the last event time
	callback      func()    // Store the callback function
}

// NewWatcher creates a new instance of Watcher and initializes data storage.
func NewWatcher(searchPath []string) *WorkstationWatcher {
	return &WorkstationWatcher{
		SearchPath: searchPath,
	}
}
