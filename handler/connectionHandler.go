package handler

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var (
	clients   = make(map[*websocket.Conn]bool) // client connections map
	broadcast = make(chan Message)             // Broadcast Channel
	// configure the upgrader
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

func handleConnections(w http.ResponseWriter, r *http.Request) {
	// upgrade Get request to websocket
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal("ConnectionHandler:", err)
	}

	// close when this function ends
	defer ws.Close()

	// add client to global clients map
	clients[ws] = true

	for {
		var msg Message

		// read new msg in as JSON, and initialize is as a Message object
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Printf("ConnectionHandlerMsg: %v", err)
			delete(clients, ws)
			break
		}
		// send the msg to the broadcast channel, to be handled by handleMessage goroutine
		broadcast <- msg
	}
}
