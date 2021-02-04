package chat

import "github.com/gorilla/websocket"

// Message object, defines attributes that a message will have
type Message struct {
	Username  string          `json:"username"`
	Message   string          `json:"message"`
	Avatar    string          `json:"avatar"`
	Encrypted bool            `json:"encrypted"`
	Leaving   bool            `json:"leaving"`
	UUID      string          `json:"uuid"`
	WS        *websocket.Conn `json:"-"`
	Users     map[string]User `json:"users"`
}
