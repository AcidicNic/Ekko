package handler

import (
	"net/http"

	"github.com/gorilla/websocket"
)

// global variables
var (
	allUsers  AllUsers
	clients   = make(map[*websocket.Conn]bool) // client connections map
	broadcast = make(chan Message)             // Broadcast Channel
	// configure the upgrader
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

// Message object, defines attributes that a message will have
type Message struct {
	Username  string `json:"username"`
	Message   string `json:"message"`
	Avatar    string `json:"avatar"`
	Encrypted bool   `json:"encrypted"`
}

// User object, holds information for a specific user
type User struct {
	Username string `json:"user"`
	Avatar   string `json:"avatar"`
}

// AllUsers object,  holds all the users in a string array
type AllUsers struct {
	Users []User `json:"users"`
}
