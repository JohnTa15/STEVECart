package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"

	"steve-api/initializers"
	"steve-api/models"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// upgrader upgrades HTTP connections to WebSocket
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	// Allow all origins during development
	CheckOrigin: func(r *http.Request) bool { return true },
}

// clients holds all currently connected WebSocket connections
var (
	clients   = make(map[*websocket.Conn]bool)
	clientsMu sync.Mutex
)

// HandleMinimapWS is the exported Gin handler — register it in main.go
// GET /ws/minimap
func HandleMinimapWS(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Printf("WebSocket upgrade error: %v", err)
		return
	}

	// Register the new client
	clientsMu.Lock()
	clients[conn] = true
	clientsMu.Unlock()
	fmt.Printf("Minimap client connected. Total clients: %d\n", len(clients))

	// Cleanup when the client disconnects
	defer func() {
		clientsMu.Lock()
		delete(clients, conn)
		clientsMu.Unlock()
		conn.Close()
		fmt.Println("Minimap client disconnected")
	}()

	// Read loop — necessary to detect disconnection (ping/pong frames)
	for {
		if _, _, err := conn.ReadMessage(); err != nil {
			break
		}
	}
}

// StartUWBBroadcaster listens on initializers.UWBBroadcast and fans out
// to every connected WebSocket client. Call this ONCE from main.go as a goroutine.
func StartUWBBroadcaster() {
	go func() {
		for uwbData := range initializers.UWBBroadcast {
			broadcastUWB(uwbData)
		}
	}()
}

// broadcastUWB sends a UWBData payload to all connected clients
func broadcastUWB(data models.UWBData) {
	payload, err := json.Marshal(data)
	if err != nil {
		log.Printf("UWB JSON marshal error: %v", err)
		return
	}

	clientsMu.Lock()
	defer clientsMu.Unlock()

	for conn := range clients {
		if err := conn.WriteMessage(websocket.TextMessage, payload); err != nil {
			log.Printf("WebSocket write error (client removed): %v", err)
			conn.Close()
			delete(clients, conn)
		}
	}
}