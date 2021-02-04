package handler

import (
	"net/http"

	"github.com/AcidicNic/Ekko/chat"
	"github.com/gorilla/websocket"
)

// Handler handles the requests
type Handler struct {
	Rooms     map[string]chat.Room
	broadcast chan chat.Message
	upgrader  websocket.Upgrader
}

// NewHandler instantiates a Handler object
func NewHandler() Handler {
	return Handler{
		Rooms:     make(map[string]chat.Room),
		broadcast: make(chan chat.Message),
		upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
	}
}
