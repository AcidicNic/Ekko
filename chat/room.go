package chat

import "github.com/gorilla/websocket"

// Room holds the information for an individual room
type Room struct {
	UUID    string
	Users   map[string]User
	Clients map[*websocket.Conn]bool
}

// NewRoom instantiates a new room object
func NewRoom(uuid string) Room {
	return Room{
		UUID:    uuid,
		Users:   make(map[string]User),
		Clients: make(map[*websocket.Conn]bool),
	}
}
