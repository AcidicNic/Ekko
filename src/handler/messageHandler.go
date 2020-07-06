package handler

import (
	"fmt"
	"log"
)

// HandleMessages sends and receives messages from users.
func HandleMessages() {
	for {
		msg := <-broadcast

		// if a new user joins. then send a message
		if msg.Encrypted == false && msg.Leaving == false {
			usr := User{Username: msg.Username, Avatar: msg.Avatar}

			msg.Message = fmt.Sprintf("NEW USER JOINED:  %s", usr.Username)
			msg.Username = "ATTENTION!"
			msg.Avatar = "red-alert.png"
		}
		// if a user is leaving, send a message
		if msg.Encrypted == false && msg.Leaving == true {
			msg.Message = fmt.Sprintf("USER LEAVING:  %s", msg.Username)
			msg.Username = "ATTENTION!"
			msg.Avatar = "red-alert.png"
		}

		for client := range clients {
			err := client.WriteJSON(msg)
			if err != nil {
				log.Printf("error: %v", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}
