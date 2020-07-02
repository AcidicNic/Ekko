package handler

import (
	"log"
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

// HandleConnections maintains the connection for each user in the chat
func HandleConnections(w http.ResponseWriter, r *http.Request) {
	// upgrade Get request to websocket
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal("ConnectionHandler:", err)
	}

	// close when this function ends
	defer ws.Close()

	// create a map that holds all the server instances by uuid,
	// key = uuid, value = clients map

	// add client to global clients map
	clients[ws] = true

	for {
		// initialize msg variable of Message type
		var msg Message

		// read new msg in as JSON, and initialize is as a Message object
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Printf("ConnectionHandlerMsg: %v", err)

			if websocket.IsCloseError(err, 1001) {
				broadcast <- Message{Username: "ATTENTION!", Avatar: "red-alert.png", Message: "Someone left the chat."}
			}

			delete(clients, ws)
			break
		}
		// send the msg to the broadcast channel, to be handled by handleMessage goroutine
		broadcast <- msg
	}
}
