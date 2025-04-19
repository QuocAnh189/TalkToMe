package websocket

import (
	"gochat/pkg/logger"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// WsHandler handles WebSocket connection requests.
func WsHandler(hub *Hub, c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		logger.Errorf("Failed to upgrade connection: %v", err)
		return
	}
	logger.Infof("WebSocket connection established: %s", conn.RemoteAddr())

	userID := c.Query("user_id")
	if userID == "" {
		logger.Warnf("UserID not provided for connection from %s", conn.RemoteAddr())
		conn.Close()
		return
	}

	client := &Client{
		hub:    hub,
		conn:   conn,
		send:   make(chan []byte, 256),
		UserID: userID,
		Groups: make(map[string]bool),
	}
	client.hub.register <- client

	go client.writePump()
	go client.readPump()
}
