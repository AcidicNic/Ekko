package handler

import (
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

// HandleMessages sends and receives messages from users.
func HandleMessages() {
	for {
		msg := <-broadcast

		if msg.UUID == "" {
			// if there is no uuid with the message just skip it
			continue
		}

		// if a new user joins. then send a message
		if msg.Encrypted == false && msg.Leaving == false {

			msg.Message = fmt.Sprintf("NEW USER JOINED:  %s", msg.Username)
			msg.Username = "ATTENTION!"
			msg.Avatar = "red-alert.png"

			// Add them to the room
			rooms[msg.UUID][msg.ws] = true

			if len(rooms[msg.UUID]) == 1 {
				first(msg.ws)
			}
		} else if msg.Encrypted == false && msg.Leaving == true {
			// if a user is leaving, send a message
			msg.Message = fmt.Sprintf("USER LEAVING:  %s", msg.Username)
			msg.Username = "ATTENTION!"
			msg.Avatar = "red-alert.png"

			// remove them from the clients map connected to the room
			delete(rooms[msg.UUID], msg.ws)

			// If the last person leaves delete the room
			if len(rooms[msg.UUID]) == 0 {
				delete(rooms, msg.UUID)
			}
		}

		for client := range rooms[msg.UUID] {
			msg.ws = nil
			err := client.WriteJSON(msg)
			if err != nil {
				log.Printf("error: %v", err)
				client.Close()
				delete(rooms[msg.UUID], client)
			}
		}
	}
}

// helper function for sending a message to the first user of a room
func first(ws *websocket.Conn) {
	msg := Message{}
	msg.Message = fmt.Sprintf("You are the first occupant of this room.")
	msg.Username = "ATTENTION!"
	msg.Avatar = "red-alert.png"

	_ = ws.WriteJSON(msg)
}
