package servers

import (
	"fmt"
	"log"
	"mdlook/source/internal/mdrepo"
	"net/http"
	"os"
	"path/filepath"
)

// Server struct to hold server-related data and methods.
type WebServer struct {
	Address   string
	Port      string
	mdlook    *mdrepo.MDLookManager
	server    *http.Server
	isRunning bool
	isDevMode bool
}

// NewWebServer creates a new instance of WebServer and initializes data storage.
func NewWebServer(addr string, port string, mdManager *mdrepo.MDLookManager, isDevMode bool) *WebServer {
	return &WebServer{
		Address:   addr,
		Port:      port,
		isRunning: false,
		mdlook:    mdManager,
		isDevMode: isDevMode,
	}
}

// noCacheHandler wraps the HTTP handler with cache control headers.
func noCacheHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
		w.Header().Set("Pragma", "no-cache")
		w.Header().Set("Expires", "0")
		h.ServeHTTP(w, r)
	})
}

// Start begins serving the HTTP server.
func (s *WebServer) Start() {

	// Prevent the server from starting if it is already running
	if s.isRunning {
		log.Println("Server is already running!")
		return
	}

	// Get paths
	webUIPath := s.mdlook.GetWebTemplateDir()
	indexHtmlPath := s.mdlook.GetIndexHtmlPath()
	workspacePath := s.mdlook.WorkstationDir

	// Set up file servers
	primaryFileServer := noCacheHandler(http.FileServer(http.Dir(webUIPath)))
	secondaryFileServer := noCacheHandler(http.FileServer(http.Dir(workspacePath)))

	// Set up HTTP routes
	mux := http.NewServeMux()

	rootHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path

		primaryFilePath := filepath.Join(webUIPath, path)
		if _, err := os.Stat(primaryFilePath); err == nil {
			primaryFileServer.ServeHTTP(w, r)
			return
		}

		secondaryFilePath := filepath.Join(s.mdlook.WorkstationDir, path)
		if _, err := os.Stat(secondaryFilePath); err == nil {
			secondaryFileServer.ServeHTTP(w, r)
			return
		}

		// Inject WebSocket code if dev mode is enabled
		if s.isDevMode {
			socketCode := s.GetInjectSocketCode(s.Address, IncrementPort(s.Port))
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
			w.Write([]byte("<script>" + socketCode + "</script>"))
			http.ServeFile(w, r, indexHtmlPath)

			// Inject WebSocket script at the end of the body
			return
		}

		w.WriteHeader(http.StatusOK)
		http.ServeFile(w, r, indexHtmlPath)
	})

	// Register the root handler
	mux.Handle("/", rootHandler)

	// Initialize the HTTP server
	s.server = &http.Server{
		Addr:    fmt.Sprintf("%s:%s", s.Address, s.Port),
		Handler: mux,
	}

	// Start the server
	fmt.Printf("Starting server on %s:%s\n", s.Address, s.Port)
	s.isRunning = true
	// IMPORTANT: ListenAndServe blocks. When s.Stop() calls Close(), it returns http.ErrServerClosed.
	if err := s.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		// Log a fatal error only if it's NOT the expected server closed error
		log.Printf("Error starting server: %v", err)
	}
}

func (s *WebServer) GetInjectSocketCode(host, port string) string {
	// JavaScript code to inject for WebSocket functionality with dynamic host and port
	socketCode := fmt.Sprintf(`
			const socket = new WebSocket("ws://%s:%s/ws");

			socket.onopen = function () {
				console.log("WebSocket connection established.");
			};

			socket.onmessage = function (event) {
				if (event.data === "reload") {
					console.log("Received reload signal, refreshing the page.");
					window.location.reload();
				}
			};

			socket.onclose = function () {
				console.log("WebSocket connection closed.");
			};
	`, host, port)

	return socketCode
}

// Stop gracefully shuts down the server.
func (s *WebServer) Stop() {
	// Check if server is running before attempting to stop
	if s.isRunning {
		// Use Close() for immediate shutdown during a watch/reload cycle
		if err := s.server.Close(); err != nil {
			fmt.Printf("Error stopping server: %v\n", err)
		}
		s.server = nil // Mark server as stopped
	} else {
		fmt.Println("Server is not running.")
	}
}

// Reload stops and restarts the server.
func (s *WebServer) Reload() {
	fmt.Println("Reloading server...")
	s.Stop()
	s.Start()
}
