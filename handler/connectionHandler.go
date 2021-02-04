package handler

import (
	"log"

	"github.com/labstack/echo/v4"
)

// HandleConnections maintains the connection for each user in the chat
func (h *Handler) HandleConnections(c echo.Context) error {
	// upgrade Get request to websocket
	socket, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		log.Fatal("ConnectionHandler:", err)
	}

	// close when this function ends
	defer socket.Close()

	// add client to global clients map
	clients[socket] = true

	for {
		var msg Message

		// read new msg in as JSON, and initialize is as a Message object
		err := socket.ReadJSON(&msg)
		if err != nil {
			log.Printf("ConnectionHandlerMsg: %v", err)
			delete(clients, socket)
			break
		}

		msg.ws = socket
		// send the msg to the broadcast channel, to be handled by handleMessage goroutine
		broadcast <- msg
	}
	return nil
}
