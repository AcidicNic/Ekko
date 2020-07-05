package handler

import (
	"fmt"
	"log"
)

// HandleMessages sends and receives messages from users.
func HandleMessages() {
	for {
		msg := <-broadcast

		if msg.Message == "" {
			usr := User{Username: msg.Username, Avatar: msg.Avatar}

			msg.Message = fmt.Sprintf("NEW USER JOINED:  %s", usr.Username)
			msg.Username = "ATTENTION!"
			msg.Avatar = "red-alert.png"
		}

		for client := range clients {
			// Is this where we put the encryption of the message?
			err := client.WriteJSON(msg)
			if err != nil {
				log.Printf("error: %v", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}
