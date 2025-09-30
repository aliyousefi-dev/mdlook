package servers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// WebsocketServer struct to hold server-related data and methods.
type WebsocketServer struct {
	Address string
	Port    string
	clients map[*websocket.Conn]bool // Map to store connected WebSocket clients
}

// NewWebsocket creates a new instance of WebsocketServer and initializes data storage.
func NewWebsocketServer(addr string, port string) *WebsocketServer {
	return &WebsocketServer{
		Address: addr,
		Port:    port,
		clients: make(map[*websocket.Conn]bool), // Initialize the clients map
	}
}

// upgrader for upgrading HTTP connections to WebSocket.
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Allow connections from any origin (can be restricted if needed)
	},
}

// handleWebSocket handles incoming WebSocket connections.
func (ws *WebsocketServer) HandleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("WebSocket Upgrade failed:", err)
		return
	}
	defer conn.Close()

	// Add the new connection to the clients map
	ws.clients[conn] = true
	log.Println("New WebSocket connection established")

	// Listen for incoming messages (optional) or just keep the connection open
	for {
		_, _, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error reading message:", err)
			delete(ws.clients, conn) // Remove the client from the map on disconnect
			break
		}
	}
}

// BroadcastReload sends a "reload" message to all connected WebSocket clients.
func (ws *WebsocketServer) BroadcastReload() {
	for client := range ws.clients {
		err := client.WriteMessage(websocket.TextMessage, []byte("reload"))
		if err != nil {
			log.Println("Error broadcasting reload:", err)
			client.Close()
			delete(ws.clients, client) // Clean up the client connection
		}
	}
}

// Start starts the WebSocket server.
func (ws *WebsocketServer) Start() {
	http.HandleFunc("/ws", ws.HandleWebSocket)

	address := fmt.Sprintf("%s:%s", ws.Address, ws.Port)
	log.Printf("WebSocket server started at ws://%s", address)

	if err := http.ListenAndServe(address, nil); err != nil {
		log.Fatal("Error starting WebSocket server:", err)
	}
}
