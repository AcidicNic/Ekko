package handler

import (
	"fmt"
	"log"
	"net/http"
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
		fmt.Println(msg.UUID)
		// send the msg to the broadcast channel, to be handled by handleMessage goroutine
		broadcast <- msg
	}
}
