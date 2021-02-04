package handler

import (
	"fmt"
	"log"

	"github.com/AcidicNic/Ekko/chat"
	"github.com/gorilla/websocket"
)

// HandleMessages sends and receives messages from users.
func (h *Handler) HandleMessages() {
	for {
		msg := <-h.broadcast

		// if a new user joins. then send a message
		if msg.Encrypted == false && msg.Leaving == false {
			usr := chat.NewUser(msg.Username, msg.Avatar)

			h.Rooms[msg.UUID].Users[usr.Username] = usr

			msg.Message = fmt.Sprintf("NEW USER JOINED:  %s", msg.Username)
			msg.Username = "ATTENTION!"
			msg.Avatar = "chat/red-alert.png"

			// Add them to the clients in the room
			h.Rooms[msg.UUID].Clients[msg.WS] = true

			if len(h.Rooms[msg.UUID].Clients) == 1 {
				first(msg.WS)
			}
		} else if msg.Encrypted == false && msg.Leaving == true {
			// if a user is leaving delete them and send a message
			delete(h.Rooms[msg.UUID].Users, msg.Username)

			msg.Message = fmt.Sprintf("USER LEAVING:  %s", msg.Username)
			msg.Username = "ATTENTION!"
			msg.Avatar = "chat/red-alert.png"

			// remove them from the clients map connected to the room
			delete(h.Rooms[msg.UUID].Clients, msg.WS)

			// If the last person leaves delete the room
			if len(h.Rooms[msg.UUID].Clients) == 0 {
				delete(h.Rooms, msg.UUID)
			}
		}

		msg.Users = h.Rooms[msg.UUID].Users

		for client := range h.Rooms[msg.UUID].Clients {
			err := client.WriteJSON(msg)
			if err != nil {
				log.Printf("error: %v", err)
				delete(h.Rooms[msg.UUID].Users, msg.Username)
				delete(h.Rooms[msg.UUID].Clients, client)
				client.Close()
			}
		}
	}
}

// helper function for sending a message to the first user of a room
func first(ws *websocket.Conn) {
	msg := chat.Message{}
	msg.Message = fmt.Sprintf("You are the first occupant of this room.")
	msg.Username = "ATTENTION!"
	msg.Avatar = "chat/red-alert.png"

	_ = ws.WriteJSON(msg)
}
