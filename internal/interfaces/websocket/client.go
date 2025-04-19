package websocket

import (
	"bytes"
	"encoding/json"
	"gochat/pkg/logger"
	"time"

	"github.com/gorilla/websocket"
)

const (
	writeWait      = 10 * time.Second    // Time allowed to write a message to the peer.
	pongWait       = 60 * time.Second    // Time allowed to read the next pong message from the peer.
	pingPeriod     = (pongWait * 9) / 10 // Send pings to peer with this period. Must be less than pongWait.
	maxMessageSize = 1024 * 4            // Maximum message size allowed from peer.
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

// Client is a middleman between the websocket connection and the hub.
type Client struct {
	hub    *Hub
	conn   *websocket.Conn
	send   chan []byte
	UserID string
	Groups map[string]bool
}

// Message struct to represent incoming and outgoing messages
type Message struct {
	Type            string `json:"type"`
	TargetUserID    string `json:"target_user_id,omitempty"`
	GroupID         string `json:"group_id,omitempty"`
	Content         string `json:"content,omitempty"`
	RequesterUserID string `json:"requester_user_id,omitempty"`
	SenderUserID    string `json:"sender_user_id,omitempty"` // Useful for group messages
}

func (c *Client) readPump() {
	defer func() {
		c.hub.unregister <- c
		c.conn.Close()
		logger.Infof("WebSocket readPump closed for UserID: %s, RemoteAddr: %s", c.UserID, c.conn.RemoteAddr())
	}()
	c.conn.SetReadLimit(maxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error { c.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		_, payload, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				logger.Errorf("WebSocket read error for UserID: %s, RemoteAddr: %s: %v", c.UserID, c.conn.RemoteAddr(), err)
			} else {
				logger.Warnf("WebSocket connection closed normally for UserID: %s, RemoteAddr: %s: %v", c.UserID, c.conn.RemoteAddr(), err)
			}
			break
		}
		payload = bytes.TrimSpace(bytes.Replace(payload, newline, space, -1))

		var msg Message
		if err := json.Unmarshal(payload, &msg); err != nil {
			logger.Errorf("Error unmarshalling message from UserID: %s, RemoteAddr: %s: %v, Payload: %s", c.UserID, c.conn.RemoteAddr(), err, string(payload))
			continue // Skip processing invalid JSON
		}

		// c.hub.broadcast <- payload
		c.processMessage(msg)
	}
}

func (c *Client) writePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.conn.Close()
		logger.Infof("WebSocket writePump closed for UserID: %s, RemoteAddr: %s", c.UserID, c.conn.RemoteAddr())
	}()
	for {
		select {
		case messageBytes, ok := <-c.send:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				logger.Errorf("WebSocket writer error for UserID: %s, RemoteAddr: %s: %v", c.UserID, c.conn.RemoteAddr(), err)
				return
			}
			_, err = w.Write(messageBytes)
			if err != nil {
				logger.Errorf("WebSocket write error for UserID: %s, RemoteAddr: %s: %v", c.UserID, c.conn.RemoteAddr(), err)
				return
			}

			n := len(c.send)
			for i := 0; i < n; i++ {
				_, err = w.Write(newline)
				if err != nil {
					logger.Errorf("WebSocket write (queued newline) error for UserID: %s, RemoteAddr: %s: %v", c.UserID, c.conn.RemoteAddr(), err)
					return
				}
				queuedMsgBytes := <-c.send
				_, err = w.Write(queuedMsgBytes)
				if err != nil {
					logger.Errorf("WebSocket write (queued message) error for UserID: %s, RemoteAddr: %s: %v", c.UserID, c.conn.RemoteAddr(), err)
					return
				}
			}

			if err := w.Close(); err != nil {
				logger.Errorf("WebSocket writer close error for UserID: %s, RemoteAddr: %s: %v", c.UserID, c.conn.RemoteAddr(), err)
				return
			}
		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				logger.Errorf("WebSocket ping error for UserID: %s, RemoteAddr: %s: %v", c.UserID, c.conn.RemoteAddr(), err)
				return
			}
		}
	}
}

func (c *Client) processMessage(msg Message) {
	logger.Debugf("Received message from UserID: %s, RemoteAddr: %s: %+v", c.UserID, c.conn.RemoteAddr(), msg)

	switch msg.Type {
	case "private_message":
		if msg.TargetUserID != "" && msg.Content != "" {
			c.hub.sendMessage(c, msg.TargetUserID, []byte(msg.Content))
		} else {
			logger.Warnf("Invalid private message format from UserID: %s", c.UserID)
		}
	case "join_group":
		if msg.GroupID != "" {
			c.hub.joinGroup(c, msg.GroupID)
		} else {
			logger.Warnf("Invalid join group format from UserID: %s", c.UserID)
		}
	case "group_message":
		if msg.GroupID != "" && msg.Content != "" {
			c.hub.sendMessageToGroup(c, msg.GroupID, []byte(msg.Content))
		} else {
			logger.Warnf("Invalid group message format from UserID: %s", c.UserID)
		}
	case "friend_request":
		if msg.TargetUserID != "" {
			c.hub.sendFriendRequest(c, msg.TargetUserID)
		} else {
			logger.Warnf("Invalid friend request format from UserID: %s", c.UserID)
		}
	case "accept_friend_request":
		if msg.RequesterUserID != "" {
			c.hub.acceptFriendRequest(c, msg.RequesterUserID)
		} else {
			logger.Warnf("Invalid accept friend request format from UserID: %s", c.UserID)
		}
	case "add_to_group":
		if msg.GroupID != "" && msg.TargetUserID != "" {
			c.hub.addUserToGroup(c, msg.GroupID, msg.TargetUserID)
		} else {
			logger.Warnf("Invalid add to group format from UserID: %s", c.UserID)
		}
	case "remove_from_group":
		if msg.GroupID != "" && msg.TargetUserID != "" {
			c.hub.removeUserFromGroup(c, msg.GroupID, msg.TargetUserID)
		} else {
			logger.Warnf("Invalid remove from group format from UserID: %s", c.UserID)
		}
	default:
		logger.Warnf("Unknown message type from UserID: %s: %s", c.UserID, msg.Type)
		c.hub.broadcast <- []byte(msg.Content)
	}
}
