package watcher

import (
	"fmt"
	"time"

	"github.com/fsnotify/fsnotify"
)

// HandleFileEvent now checks the time difference
func (w *WorkstationWatcher) HandleFileEvent(event fsnotify.Event) {
	// Check if the event occurred within 500 milliseconds of the last event
	if time.Since(w.lastEventTime) < 500*time.Millisecond {
		// Ignore event if it's too close to the last event
		return
	}

	// Define the event type (Create, Modify, Remove)
	var eventType string

	if event.Op&fsnotify.Write == fsnotify.Write {
		eventType = "modified"
		fmt.Println("File modified:", eventType)
	} else if event.Op&fsnotify.Create == fsnotify.Create {
		eventType = "created"
		fmt.Println("File created:", eventType)
	} else if event.Op&fsnotify.Remove == fsnotify.Remove {
		eventType = "removed"
		fmt.Println("File removed:", eventType)
	}

	// Update the last event time
	w.lastEventTime = time.Now()

	// Call the callback (if set)
	if w.callback != nil {
		w.callback()
	}
}
