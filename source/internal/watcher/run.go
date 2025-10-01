package watcher

import (
	"log"

	"github.com/fsnotify/fsnotify"
)

// Watch files and trigger sync (No change needed here)
func (w *WorkstationWatcher) Run() {

	// ... (rest of the function is the same)
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	// Start listening for events.
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				w.HandleFileEvent(event)
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()

	for _, path := range w.SearchPath {
		// Watch the directory containing your markdown files (replace with your path)
		err = watcher.Add(path)
		if err != nil {
			log.Fatal(err)
		}
	}

	// Block main goroutine forever.
	<-make(chan struct{})
}
