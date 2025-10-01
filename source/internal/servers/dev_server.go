package servers

import (
	"fmt"
	"log"
	"mdlook/source/internal/mdrepo"
	"mdlook/source/internal/watcher"
)

// Server struct to hold server-related data and methods.
type DevServer struct {
	Address   string
	Port      string
	mdlook    *mdrepo.MDLookManager
	WSServer  *WebsocketServer
	WebServer *WebServer
	Watcher   *watcher.WorkstationWatcher
}

// NewServer creates a new instance of Server and initializes data storage.
func NewDevServer(addr string, port string, mdManager *mdrepo.MDLookManager) *DevServer {
	searchPathes := []string{mdManager.GetNavFilePath()}

	dirPaths, err := mdManager.GetAllFolderPaths()
	if err != nil {
		log.Println("Error getting folder paths:", err)
		return nil
	}

	searchPathes = append(searchPathes, dirPaths...)

	return &DevServer{
		Address:   addr,
		Port:      port,
		mdlook:    mdManager,
		WSServer:  NewWebsocketServer(addr, IncrementPort(port)),
		WebServer: NewWebServer(addr, port, mdManager, true),
		Watcher:   watcher.NewWatcher(searchPathes),
	}
}

func (s *DevServer) Start() {
	s.Watcher.SetCallback(s.SendReloadSignal)
	go s.WebServer.Start()
	go s.WSServer.Start()
	s.Watcher.Run()
}

func (s *DevServer) SendReloadSignal() {
	fmt.Println("Sending reload signal to WebSocket clients...")
	s.mdlook.SyncNav()
	s.WSServer.BroadcastReload()
}
