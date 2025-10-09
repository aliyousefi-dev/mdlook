package servers

import (
	"fmt"
	"log"
	"mdlook/source/internal/mdrepo"
	"mdlook/source/internal/watcher"
)

// Server struct to hold server-related data and methods.
type DevServer struct {
	Address     string
	Port        string
	DisableSync bool
	mdlook      *mdrepo.MDLookManager
	WSServer    *WebsocketServer
	WebServer   *WebServer
	Watcher     *watcher.WorkstationWatcher
}

// NewServer creates a new instance of Server and initializes data storage.
func NewDevServer(addr string, port string, mdManager *mdrepo.MDLookManager, disableSync bool) *DevServer {
	searchPathes := []string{mdManager.Workstation.GetNavFilePath()}

	dirPaths, err := mdManager.Workstation.ListDocDirs()
	if err != nil {
		log.Println("Error getting folder paths:", err)
		return nil
	}

	searchPathes = append(searchPathes, dirPaths...)

	return &DevServer{
		Address:     addr,
		Port:        port,
		DisableSync: disableSync,
		mdlook:      mdManager,
		WSServer:    NewWebsocketServer(addr, IncrementPort(port)),
		WebServer:   NewWebServer(addr, port, mdManager, true),
		Watcher:     watcher.NewWatcher(searchPathes),
	}
}

// NewDevServerDefault creates a new DevServer with disableSync defaulting to false.
func NewDevServerDefault(addr string, port string, mdManager *mdrepo.MDLookManager) *DevServer {
	return NewDevServer(addr, port, mdManager, false)
}

func (s *DevServer) Start() {
	s.Watcher.SetCallback(s.SendReloadSignal)
	go s.WebServer.Start()
	go s.WSServer.Start()
	s.Watcher.Run()
}

func (s *DevServer) SendReloadSignal() {
	fmt.Println("Sending reload signal to WebSocket clients...")

	if !s.DisableSync {
		s.mdlook.SyncNav()
	}
	s.WSServer.BroadcastReload()
}
