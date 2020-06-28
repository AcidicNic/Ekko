package handler

// Message object, defines attributes that a message will have
type Message struct {
	Username string `json:"username"`
	Message  string `json:"message"`
	Avatar   string `json:"avatar"`
}

// User object, holds information for a specific user
type User struct {
	Username string `json:"user"`
	Avatar   string `json:"avatar"`
}

// AllUsers object,  holds all the users in a string array
type AllUsers struct {
	Users []User `json:"users"`
}
