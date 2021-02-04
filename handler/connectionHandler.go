package handler

import (
	"log"

	"github.com/AcidicNic/Ekko/chat"
	"github.com/labstack/echo/v4"
)

// HandleConnections maintains the connection for each user in the chat
func (h *Handler) HandleConnections(c echo.Context) error {
	// upgrade Get request to websocket
	socket, err := h.upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		log.Fatal("ConnectionHandler:", err)
	}

	// close when this function ends
	defer socket.Close()

	for {
		var msg chat.Message

		// read new msg in as JSON, and initialize is as a Message object
		err := socket.ReadJSON(&msg)
		if err != nil {
			break
		}

		if msg.UUID == "" {
			// if there is no uuid with the message just skip it
			continue
		}

		msg.WS = socket
		// send the msg to the broadcast channel, to be handled by handleMessage goroutine
		h.broadcast <- msg
	}
	return nil
}
