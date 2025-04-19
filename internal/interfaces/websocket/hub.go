package websocket

import (
	"gochat/pkg/logger"
)

type Hub struct {
	clients        map[*Client]bool
	groups         map[string]map[*Client]bool // Map of group ID to a set of clients in that group
	broadcast      chan []byte                 // General broadcast messages
	register       chan *Client
	unregister     chan *Client
	friendRequests map[string]map[string]bool // userID -> map of pending friend requests (targetUserID -> true)
}

func NewHub() *Hub {
	return &Hub{
		clients:        make(map[*Client]bool),
		groups:         make(map[string]map[*Client]bool),
		broadcast:      make(chan []byte),
		register:       make(chan *Client),
		unregister:     make(chan *Client),
		friendRequests: make(map[string]map[string]bool),
	}
}

func (h *Hub) Run() {
	logger.Info("WebSocket Hub started")
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
			logger.Infof("Client registered: UserID: %s, RemoteAddr: %s", client.UserID, client.conn.RemoteAddr())
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				h.removeFromAllGroups(client)
				close(client.send)
				logger.Infof("Client unregistered: UserID: %s, RemoteAddr: %s", client.UserID, client.conn.RemoteAddr())
			}
		case message := <-h.broadcast:
			for client := range h.clients {
				h.sendMessageToClient(client, message)
			}
		}
	}
}

func (h *Hub) sendMessageToClient(client *Client, message []byte) {
	select {
	case client.send <- message:
	default:
		close(client.send)
		delete(h.clients, client)
		h.removeFromAllGroups(client)
		logger.Warnf("Client send buffer full, closing connection for UserID: %s, RemoteAddr: %s", client.UserID, client.conn.RemoteAddr())
	}
}

func (h *Hub) sendMessage(_ *Client, targetUserID string, message []byte) {
	for client := range h.clients {
		if client.UserID == targetUserID {
			h.sendMessageToClient(client, message)
			return
		}
	}
	logger.Warnf("Target user not found: %s", targetUserID)
}

func (h *Hub) sendMessageToGroup(_ *Client, groupID string, message []byte) {
	if group, ok := h.groups[groupID]; ok {
		for client := range group {
			h.sendMessageToClient(client, message)
		}
	} else {
		logger.Warnf("Group not found: %s", groupID)
	}
}

func (h *Hub) sendFriendRequest(sender *Client, targetUserID string) {
	if _, ok := h.clients[sender]; !ok {
		logger.Warnf("Sender not registered: %s", sender.UserID)
		return
	}
	for client := range h.clients {
		if client.UserID == targetUserID {
			if h.friendRequests[targetUserID] == nil {
				h.friendRequests[targetUserID] = make(map[string]bool)
			}
			if !h.friendRequests[targetUserID][sender.UserID] {
				h.friendRequests[targetUserID][sender.UserID] = true
				notification := []byte(sender.UserID + " sent you a friend request.")
				h.sendMessageToClient(client, notification)
			}
			return
		}
	}
	logger.Warnf("Target user not found for friend request: %s", targetUserID)
}

func (h *Hub) acceptFriendRequest(accepter *Client, requesterUserID string) {
	if _, ok := h.clients[accepter]; !ok {
		logger.Warnf("Accepter not registered: %s", accepter.UserID)
		return
	}
	if _, ok := h.friendRequests[accepter.UserID][requesterUserID]; ok {
		delete(h.friendRequests[accepter.UserID], requesterUserID)
		notificationAccepter := []byte("You accepted " + requesterUserID + "'s friend request.")
		h.sendMessageToClient(accepter, notificationAccepter)
		notificationRequester := []byte(accepter.UserID + " accepted your friend request.")
		h.sendMessageToClient(h.getClientByUserID(requesterUserID), notificationRequester)
	} else {
		logger.Warnf("No pending friend request from %s to %s", requesterUserID, accepter.UserID)
	}
}

func (h *Hub) addUserToGroup(inviter *Client, groupID string, targetUserID string) {
	if _, ok := h.clients[inviter]; !ok {
		logger.Warnf("Inviter not registered: %s", inviter.UserID)
		return
	}
	targetClient := h.getClientByUserID(targetUserID)
	if targetClient == nil {
		logger.Warnf("Target user not found: %s", targetUserID)
		return
	}
	if h.groups[groupID] == nil {
		h.groups[groupID] = make(map[*Client]bool)
	}
	h.groups[groupID][targetClient] = true
	targetClient.Groups[groupID] = true
	notification := []byte(targetUserID + " has been added to group " + groupID + " by " + inviter.UserID)
	h.sendMessageToGroup(inviter, groupID, notification)
	h.sendMessageToClient(targetClient, []byte("You have been added to group "+groupID+" by "+inviter.UserID))
}

func (h *Hub) removeUserFromGroup(remover *Client, groupID string, targetUserID string) {
	if _, ok := h.clients[remover]; !ok {
		logger.Warnf("Remover not registered: %s", remover.UserID)
		return
	}
	targetClient := h.getClientByUserID(targetUserID)
	if targetClient == nil {
		logger.Warnf("Target user not found: %s", targetUserID)
		return
	}
	if group, ok := h.groups[groupID]; ok {
		if _, inGroup := group[targetClient]; inGroup {
			delete(group, targetClient)
			delete(targetClient.Groups, groupID)
			notification := []byte(targetUserID + " has been removed from group " + groupID + " by " + remover.UserID)
			h.sendMessageToGroup(remover, groupID, notification)
			h.sendMessageToClient(targetClient, []byte("You have been removed from group "+groupID+" by "+remover.UserID))
		} else {
			logger.Warnf("User %s is not in group %s", targetUserID, groupID)
		}
	} else {
		logger.Warnf("Group not found: %s", groupID)
	}
}

func (h *Hub) joinGroup(client *Client, groupID string) {
	if _, ok := h.clients[client]; !ok {
		logger.Warnf("Client not registered: %s", client.UserID)
		return
	}
	if h.groups[groupID] == nil {
		h.groups[groupID] = make(map[*Client]bool)
	}
	h.groups[groupID][client] = true
	client.Groups[groupID] = true
	logger.Infof("User %s joined group %s", client.UserID, groupID)
	notification := []byte("User " + client.UserID + " has joined group " + groupID)
	h.sendMessageToGroup(client, groupID, notification)
	h.sendMessageToClient(client, []byte("You have joined group "+groupID))
}

func (h *Hub) getClientByUserID(userID string) *Client {
	for client := range h.clients {
		if client.UserID == userID {
			return client
		}
	}
	return nil
}

func (h *Hub) removeFromAllGroups(client *Client) {
	for groupID := range client.Groups {
		if group, ok := h.groups[groupID]; ok {
			delete(group, client)
			if len(group) == 0 {
				delete(h.groups, groupID)
			}
		}
	}
}
