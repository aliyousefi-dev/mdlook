package watcher

func (w *WorkstationWatcher) SetCallback(callback func()) {
	w.callback = callback
}
